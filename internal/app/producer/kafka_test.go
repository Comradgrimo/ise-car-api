package producer

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/gammazero/workerpool"

	"github.com/ozonmp/ise-car-api/internal/mocks"
	"github.com/ozonmp/ise-car-api/internal/model"
)

func TestStart(t *testing.T) {
	events := []model.CarEvent{
		{
			ID:     0,
			Type:   model.Created,
			Status: model.InProcess,
			Entity: &model.Car{ID: 0, CarInfo: "Toyota"},
		},
		{
			ID:     1,
			Type:   model.Created,
			Status: model.InProcess,
			Entity: &model.Car{ID: 1, CarInfo: "Lexus"},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockEventRepo(ctrl)
	eventSender := mocks.NewMockEventSender(ctrl)

	eventChan := make(chan model.CarEvent)
	workerPool := workerpool.New(2)

	producer := NewKafkaProducer(
		2,
		eventSender,
		eventChan,
		repo,
		workerPool,
	)
	ctx := context.Background()

	// successfully sent to kafka -> should remove from db
	eventSender.EXPECT().Send(&events[0]).Return(nil).Times(1)
	repo.EXPECT().Remove(ctx, []uint64{events[0].ID}).Return(true, nil).Times(1)

	// failed to send to kafka -> unlock event in db
	eventSender.EXPECT().Send(&events[1]).Return(errors.New("failed to pass event to kafka")).Times(1)
	repo.EXPECT().Unlock(ctx, []uint64{events[1].ID}).Return(nil).Times(1)

	producer.Start(ctx)
	defer producer.Close()

	timeout := time.After(time.Second)
loop:
	for _, event := range events {
		select {
		case eventChan <- event:
		case <-timeout:
			break loop
		}
	}
	close(eventChan)
}
