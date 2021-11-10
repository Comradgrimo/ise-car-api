package retranslator

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/ise-car-api/internal/mocks"
	"github.com/ozonmp/ise-car-api/internal/model"
)

func TestStart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}
	var events []model.CarEvent

	for i := uint64(0); i < cfg.ConsumeSize; i++ {
		event := model.CarEvent{
			ID:     i,
			Type:   model.Created,
			Status: model.InProcess,
			Entity: &model.Car{ID: i, CarInfo: "Toyota"},
		}
		events = append(events, event)
	}

	locksCnt := uint64(0)
	repo.EXPECT().
		Lock(gomock.Any()).
		DoAndReturn(func(consumeSize uint64) ([]model.CarEvent, error) {
			atomic.AddUint64(&locksCnt, consumeSize)
			return events, nil
		}).
		AnyTimes()

	sendCnt := uint64(0)
	// successfully sent to kafka -> should remove from db
	sender.EXPECT().Send(gomock.Any()).DoAndReturn(func(event *model.CarEvent) error {
		atomic.AddUint64(&sendCnt, 1)
		return nil
	}).AnyTimes()

	removesCnt := uint64(0)

	repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(ids []uint64) error {
		atomic.AddUint64(&removesCnt, 1)
		return nil
	}).AnyTimes()
	retranslator := NewRetranslator(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConsumeTimeout)
	defer cancel()

	retranslator.Start(ctx)
	retranslator.Close()

	if !(locksCnt == sendCnt && sendCnt == removesCnt) {
		t.Errorf("Error: read from db %v messages, sent to kafka %v, removed from db %v records. "+
			"Expect them to be equal",
			locksCnt, sendCnt, removesCnt)
	}
}
