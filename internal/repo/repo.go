package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/ise-car-api/internal/model"
)

// Repo is DAO for Car
type Repo interface {
	Add(ctx context.Context, car *model.Car) (uint64, error)
	Get(ctx context.Context, carID uint64) (*model.Car, error)
	List(ctx context.Context, cursor uint64, limit uint64) (model.Cars, error)
	Remove(ctx context.Context, carID uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) Get(ctx context.Context, carID uint64) (*model.Car, error) {
	query := sq.Select("id, car_info, user_id, total_price, risk_rate, circs_link").PlaceholderFormat(sq.Dollar).
		From("car").Where(sq.And{sq.Eq{"id": carID}, sq.Eq{"removed": false}})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res model.Car
	if err = r.db.GetContext(ctx, &res, s, args...); err != nil {
		return nil, err
	}
	return &res, err
}

func (r *repo) Add(ctx context.Context, car *model.Car) (uint64, error) {

	dummyUserId := 1
	dummyCircsLink := ""
	query := sq.Insert("car").PlaceholderFormat(sq.Dollar).Columns(
		"car_info", "user_id", "total_price", "risk_rate", "circs_link").
		Values(car.CarInfo, dummyUserId, car.TotalPrice, car.RiskRate, dummyCircsLink).
		Suffix("RETURNING id").
		RunWith(r.db)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return 0, err
	}

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, err
		}

		return id, nil
	} else {
		return 0, sql.ErrNoRows
	}
}

func (r *repo) List(ctx context.Context, cursor uint64, limit uint64) (model.Cars, error) {
	query := sq.Select("id, car_info, user_id, total_price, risk_rate, circs_link").
		From("car").
		Offset(cursor).Limit(limit)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res model.Cars

	err = r.db.SelectContext(ctx, &res, s, args...)

	return res, err

}

func (r *repo) Remove(ctx context.Context, carID uint64) (bool, error) {
	query := sq.Update("car").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id" : carID}).
		Set("removed", true)
	s, args, err := query.ToSql()

	_, err = r.db.ExecContext(ctx, s, args...)
	return true, err
}
