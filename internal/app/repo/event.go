package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/ise-car-api/internal/model"
	"time"
)

// EventRepo - repo for events
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

func (r *eventRepo) Lock(ctx context.Context, n uint64) ([]model.CarEvent, error) {
	secondsInterval := 3 //todo jr fill it from config or something
	const query = `update car_event
				set status = :status, updated = now()
				where car_id in (
					select car_id from car_event
					where car_id not in
					  (    select distinct car_id
						   from car_event
						   where status = :status and now() - updated <= (:inter || 'sec')::interval
					  )
						and pg_try_advisory_xact_lock('car_event'::regclass::integer, car_id::integer)
					order by created
					limit :limit
				)
				returning car_event.*`

	rows, err := r.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"status": model.InProcess.String(),
		"limit":  n,
		"inter":  secondsInterval,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.CarEvent
	for rows.Next() {
		var event model.CarEvent
		err = rows.StructScan(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (r *eventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	query := sb.Update("car_event").
		Where(sq.Eq{"id": eventIDs}).
		Set("status", model.Available.String()).
		Set("updated", time.Now())
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return err
}

func (r *eventRepo) Remove(ctx context.Context, eventIDs []uint64) (bool, error) {
	query := sb.Delete("car_event").Where(sq.Eq{"id": eventIDs})
	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return true, err
}
