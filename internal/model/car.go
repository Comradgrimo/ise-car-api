package model

// Car - car entity.
type Car struct {
	ID    uint64 `db:"id"`
	Title string `db:"title"`
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed
)
const (
	Available EventStatus = iota
	InProcess
)

type CarEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Car
}
