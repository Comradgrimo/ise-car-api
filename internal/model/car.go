package model

// Car - car entity.
type Car struct {
	ID         uint64  `db:"id"`
	Title      string  `db:"title"`
	UserID     uint64  `db:"user_id"`
	CarInfo    string  `db:"car_info"`
	TotalPrice float64 `db:"total_price"`
	RiskRate   float64 `db:"risk_rate"`
	CircsLink  string  `db:"circs_link"`
}
type Cars []*Car

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
