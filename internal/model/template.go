package model

// Car - car entity.
type Car struct {
	ID  uint64 `db:"id"`
	Foo uint64 `db:"foo"`
}
