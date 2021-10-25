package sender

import (
	"github.com/ozonmp/ise-car-api/internal/model"
)

type EventSender interface {
	Send(car *model.CarEvent) error
}
