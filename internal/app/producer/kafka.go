package producer

import (
	"context"
	"github.com/ozonmp/ise-car-api/internal/app/repo"
	"github.com/ozonmp/ise-car-api/internal/app/sender"
	"github.com/ozonmp/ise-car-api/internal/model"
	"log"
	"sync"

	"github.com/gammazero/workerpool"
)

// Producer - interface for producer
type Producer interface {
	Start(ctx context.Context)
	Close()
}

type producer struct {
	n uint64

	sender sender.EventSender
	events <-chan model.CarEvent
	repo   repo.EventRepo

	workerPool *workerpool.WorkerPool

	wg *sync.WaitGroup
}

// NewKafkaProducer creates the producer
func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	events <-chan model.CarEvent,
	repo repo.EventRepo,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		repo:       repo,
		workerPool: workerPool,
		wg:         wg,
	}
}

func (p *producer) Start(ctx context.Context) {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event, ok := <-p.events:
					if !ok {
						return
					}
					// send event to kafka
					if err := p.sender.Send(&event); err != nil {
						p.unlockEvent(ctx, event)
					} else {
						// record processed, delete it from db event records
						p.deleteEventFromDB(ctx, event)
					}
				}
			}
		}()
	}
}

func (p *producer) deleteEventFromDB(ctx context.Context, event model.CarEvent) {
	p.wg.Add(1)
	p.workerPool.Submit(func() {
		defer p.wg.Done()
		_, err := p.repo.Remove(ctx, []uint64{event.ID})
		if err != nil {
			log.Printf("failed to remove event records from db, id=%v", event.ID)
		}
	})
}

func (p *producer) unlockEvent(ctx context.Context, event model.CarEvent) {
	p.wg.Add(1)
	p.workerPool.Submit(func() {
		defer p.wg.Done()
		//in case of error - unlock db records
		err := p.repo.Unlock(ctx, []uint64{event.ID})
		if err != nil {
			log.Printf("failed to unlock event ids %v", event.ID)
		}
	})
}

func (p *producer) Close() {
	p.wg.Wait()
}
