package repo

import (
	"github.com/ozonmp/ise-car-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.CarEvent, error) //blocks n records in DB
	Unlock(eventIDs []uint64) error

	//Add(event []model.CarEvent) error
	Remove(eventIDs []uint64) error
}
