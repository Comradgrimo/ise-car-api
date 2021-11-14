package sender

import (
	"github.com/ozonmp/ise-car-api/internal/model"
)

// EventSender - interface for sending events
type EventSender interface {
	Send(car *model.CarEvent) error
}
