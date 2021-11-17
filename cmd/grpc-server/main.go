package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/ise-car-api/internal/logger"
	gelf "github.com/snovichkov/zap-gelf"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/ozonmp/ise-car-api/internal/config"
	"github.com/ozonmp/ise-car-api/internal/database"
	"github.com/ozonmp/ise-car-api/internal/server"
)

var (
	batchSize uint = 2
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	syncLogger := initLogger(ctx, cfg)
	defer syncLogger()

	closer := initTracer(ctx, cfg)
	defer closer.Close()

	logger.InfoKV(ctx,
		fmt.Sprintf("Starting service: %s", cfg.Project.Name),
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		logger.FatalKV(ctx, "Failed to init postgres", "err", err)
	}
	defer db.Close()

	if *migration {
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			logger.ErrorKV(ctx, "Migration failed", "err", err)

			return
		}
	}

	if err := server.NewGrpcServer(db, batchSize).Start(&cfg); err != nil {
		logger.ErrorKV(ctx, "Failed creating gRPC server", "err", err)

		return
	}
}

func initLogger(ctx context.Context, cfg config.Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		logger.FatalKV(ctx, "gelf error", "err", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	logger.SetLogger(sugaredLogger.With(
		"service", cfg.Project.ServiceName,
	))

	return func() {
		notSugaredLogger.Sync()
	}
}

func initTracer(ctx context.Context, cfg config.Config) (closer io.Closer) {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	jcfg := jaegercfg.Configuration{
		ServiceName: cfg.Project.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := jcfg.NewTracer()
	if err != nil {
		logger.FatalKV(ctx, "cfg.NewTracer()", "err", err)
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	return closer
}
