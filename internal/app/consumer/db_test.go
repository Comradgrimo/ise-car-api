package consumer

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/ozonmp/ise-car-api/internal/mocks"
	"github.com/ozonmp/ise-car-api/internal/model"
)

func TestStart(t *testing.T) {
	events := []model.CarEvent{
		{
			ID:     0,
			Type:   model.Created,
			Status: model.InProcess,
			Entity: &model.Car{ID: 0, Title: "Toyota"},
		},
		{
			ID:     1,
			Type:   model.Created,
			Status: model.InProcess,
			Entity: &model.Car{ID: 1, Title: "Lexus"},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockEventRepo(ctrl)
	repo.EXPECT().Lock(uint64(len(events))).Return(events, nil).Times(1)

	eventChan := make(chan model.CarEvent)

	consTimeout := 2*time.Second
	cons := NewDbConsumer(
		1,
		uint64(len(events)),
		consTimeout,
		repo,
		eventChan,
	)
	ctx, cancel := context.WithTimeout(context.Background(), consTimeout)
	defer cancel()

	cons.Start(ctx)

	idx := 0
	timeout := time.After(1 * time.Second)
loop:
	for {
		select {
		case event := <-eventChan:
			if event != events[idx] {
				t.FailNow()
			}
			idx++
		case <-timeout:
			cons.Close()
			break loop
		}
	}
}
