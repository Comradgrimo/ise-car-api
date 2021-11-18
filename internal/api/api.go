package api

import (
	"context"
	"github.com/ozonmp/ise-car-api/internal/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
	"strings"

	"github.com/ozonmp/ise-car-api/internal/repo"

	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
)

var (
	//nolint
	totalCarNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ise_car_api_car_not_found_total",
		Help: "Total number of cars that were not found",
	})
)

type carAPI struct {
	pb.UnimplementedIseCarApiServiceServer
	repo repo.Repo
}

// NewCarAPI returns api of ise-car-api service
func NewCarAPI(r repo.Repo) pb.IseCarApiServiceServer {
	return &carAPI{repo: r}
}

func parseLogLevel(str string) (zapcore.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zapcore.DebugLevel, true
	case "info":
		return zapcore.InfoLevel, true
	case "warn":
		return zapcore.WarnLevel, true
	case "error":
		return zapcore.ErrorLevel, true
	default:
		return zapcore.DebugLevel, false
	}
}

func setLogLevelFromHeader(ctx context.Context) context.Context {
	//pass log level through header in request
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		levels := md.Get("log-level")
		logger.DebugKV(ctx, "got log level", "levels", levels)
		if len(levels) > 0 {
			if parsedLevel, ok := parseLogLevel(levels[0]); ok {
				newLogger := logger.CloneWithLevel(ctx, parsedLevel)
				ctx = logger.AttachLogger(ctx, newLogger)
			}
		}
	}
	return ctx
}