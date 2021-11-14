package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/ise-car-api/internal/database"
	"time"

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

var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (r *repo) Get(ctx context.Context, carID uint64) (*model.Car, error) {
	query := sb.Select("id, car_info, user_id, total_price, risk_rate, circs_link").
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
	res, txErr := database.WithTx(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		id, err := addCar(ctx, car, tx)
		if err != nil {
			return id, err
		}
		payload, err := getJsonBytes(car)
		if err != nil {
			return nil, err
		}
		if err := insertEvent(ctx, id, model.Created, payload, tx); err != nil {
			return 0, err
		}
		return id, nil
	})
	if txErr != nil {
		return 0, txErr
	}

	return res.(uint64), nil
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
	_, txErr := database.WithTx(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		if err := removeCar(ctx, carID, tx); err != nil {
			return nil, err
		}
		if err := insertEvent(ctx, carID, model.Removed,"\"\"", tx); err != nil {
			return nil, err
		}
		return nil, nil
	})

	if txErr != nil {
		return false, txErr
	}
	return true, nil
}

func removeCar(ctx context.Context, carID uint64, tx *sqlx.Tx) error {
	//set removed flag for car
	query := sb.Update("car").
		Where(sq.Eq{"id": carID}).
		Set("removed", true).Set("updated", time.Now())
	s, args, err := query.ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, s, args...)
	return err
}

func addCar(ctx context.Context, car *model.Car, tx *sqlx.Tx) (uint64, error) {
	dummyUserId := 1
	dummyCircsLink := ""
	query := sb.Insert("car").Columns(
		"car_info", "user_id", "total_price", "risk_rate", "circs_link").
		Values(car.CarInfo, dummyUserId, car.TotalPrice, car.RiskRate, dummyCircsLink).
		Suffix("RETURNING id")
	s, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}
	rows, err := tx.QueryContext(ctx, s, args...)
	defer rows.Close()

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
		if rows.Err() != nil {
			return 0, rows.Err()
		}
		return 0, sql.ErrNoRows
	}
}
func insertEvent(ctx context.Context, carID uint64, eventType model.EventType, payload string, tx *sqlx.Tx) error {
	query := sb.Insert("car_event").
		Columns("car_id", "type", "payload").
		Values(carID, eventType.String(), payload)
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, s, args...)

	return err
}

func getJsonBytes(car *model.Car) (string, error) {
	jsonBytes, err := json.Marshal(car)

	if err != nil {
		return "", err
	}
	return string(jsonBytes), err
}