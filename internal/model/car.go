package model

// Car - car entity.
type Car struct {
	ID         uint64  `db:"id"`
	CarInfo    string  `db:"car_info"`
	UserID     uint64  `db:"user_id"`
	TotalPrice float64 `db:"total_price"`
	RiskRate   float64 `db:"risk_rate"`
	CircsLink  string  `db:"circs_link"`
}
// Cars - slice of cars
type Cars []Car

//go:generate stringer -linecomment -type=EventType
//go:generate stringer -linecomment -type=EventStatus

// EventType - describes types of events (created, updated, removed)
type EventType uint8

// EventStatus - describes statuses for events (available, in_process)
type EventStatus uint8

const (
	_       = EventType(iota)
	// Created - event for created car
	Created // created
	// Updated - event for updated car
	Updated // updated
	// Removed - event for removed car
	Removed // removed
)
const (
	_         = EventStatus(iota)
	// Available - consumer can process this event
	Available // available
	// InProcess - this event is already locked by consumer
	InProcess // in_process
)

// CarEvent - struct for CarEvent
type CarEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Car
}
