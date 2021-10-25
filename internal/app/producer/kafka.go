package producer

import (
	"log"
	"sync"
	"time"

	"github.com/ozonmp/ise-car-api/internal/app/repo"
	"github.com/ozonmp/ise-car-api/internal/app/sender"
	"github.com/ozonmp/ise-car-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	sender sender.EventSender
	events <-chan model.CarEvent
	repo      repo.EventRepo

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	events <-chan model.CarEvent,
	repo repo.EventRepo,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		repo:       repo,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					// send event to kafka
					if err := p.sender.Send(&event); err != nil {
						p.workerPool.Submit(func() {
							//in case of error - unlock db records
							err = p.repo.Unlock([]uint64{event.ID})
							if err != nil {
								log.Printf("failed to unlock event ids %v", event.ID)
							}
						})
					} else {
						// record processed, delete it from db event records
						// call cleaners
						p.workerPool.Submit(func() {
							err = p.repo.Remove([]uint64{event.ID})
							if err != nil {
								log.Printf("failed to remove event records from db, id=%v", event.ID)
							}
						})
					}
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
