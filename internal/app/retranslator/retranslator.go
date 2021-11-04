package retranslator

import (
	"context"
	"time"

	"github.com/ozonmp/ise-car-api/internal/app/consumer"
	"github.com/ozonmp/ise-car-api/internal/app/producer"
	"github.com/ozonmp/ise-car-api/internal/app/repo"
	"github.com/ozonmp/ise-car-api/internal/app/sender"
	"github.com/ozonmp/ise-car-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Retranslator interface {
	Start(ctx context.Context)
	Close()
}

type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64 // batch size
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.CarEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.CarEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)
	producer := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.Sender,
		events,
		cfg.Repo,
		workerPool)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start(ctx context.Context) {
	r.producer.Start()
	r.consumer.Start(ctx)
}

func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
