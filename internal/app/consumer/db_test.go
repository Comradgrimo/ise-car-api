package consumer

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/ise-car-api/internal/mocks"
	"github.com/ozonmp/ise-car-api/internal/model"
	"testing"
	"time"
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
	rep := mocks.NewMockEventRepo(ctrl)
	rep.EXPECT().Lock(uint64(len(events))).Return(events, nil).Times(1)

	eventChan := make(chan model.CarEvent)

	cons := NewDbConsumer(
		1,
		uint64(len(events)),
		2*time.Second,
		rep,
		eventChan,
		)
	cons.Start()

	idx := 0
	timeout := time.After(1*time.Second)
	loop:
		for {
			select {
			case event := <-eventChan:
				//fmt.Println(event)
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