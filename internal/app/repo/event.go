package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/ise-car-api/internal/model"
)

type EventRepo interface {
	Lock(ctx context.Context, n uint64) ([]model.CarEvent, error) //blocks n records in DB
	Unlock(ctx context.Context, eventIDs []uint64) error
	Remove(ctx context.Context, eventIDs []uint64) (bool, error)
}

type eventRepo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewEventRepo returns EventRepo interface
func NewEventRepo(db *sqlx.DB, batchSize uint) EventRepo {
	return &eventRepo{db: db, batchSize: batchSize}
}

func (r *eventRepo)  Lock(_ context.Context, n uint64) ([]model.CarEvent, error) {
	//todo jr WIP
	return make([]model.CarEvent, 3), nil
}

func (r *eventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	query := sq.Update("car_event").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id" : eventIDs}).
		Set("status", "")
	s, args, err := query.ToSql()

	_, err = r.db.ExecContext(ctx, s, args...)
	return err
}

func (r *eventRepo) Remove(ctx context.Context, eventIDs []uint64) (bool, error) {
	query := sq.Delete("car_event").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": eventIDs})
	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return true, err
}





