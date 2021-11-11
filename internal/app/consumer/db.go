package consumer

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/ozonmp/ise-car-api/internal/app/repo"
	"github.com/ozonmp/ise-car-api/internal/model"
)

type Consumer interface {
	Start(ctx context.Context)
	Close()
}

type consumer struct {
	n      uint64
	events chan<- model.CarEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	wg *sync.WaitGroup
}

// NewDbConsumer contains all consumer data
func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.CarEvent) Consumer {

	wg := &sync.WaitGroup{}

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        wg,
	}
}

func timeoutWithJitter(timeout time.Duration, deviation float64) time.Duration {
	absoluteDeviation := int64(float64(timeout.Nanoseconds()) * deviation)
	jitter := rand.Int63n(absoluteDeviation) - absoluteDeviation/2 //nolint
	return time.Duration(timeout.Nanoseconds()+jitter) * time.Nanosecond
}

func (c *consumer) Start(ctx context.Context) {
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, c.timeout)
	defer cancelFunction()

	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()

			ticker := time.NewTicker(timeoutWithJitter(c.timeout, 0.1))
			for {
				// return with block. gets from db events and in the same time locks them in db
				events, err := c.repo.Lock(ctxWithTimeout, c.batchSize)
				if err != nil {
					continue
				}
				for _, event := range events {
					if event.Type == model.Created {
						c.events <- event
					}
				}
				select {
				case <-ticker.C:
				case <-ctxWithTimeout.Done():
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	c.wg.Wait()
	close(c.events)
}
