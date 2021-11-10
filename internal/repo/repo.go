package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/ise-car-api/internal/model"
)

// Repo is DAO for Car
type Repo interface {
	Add(ctx context.Context, car *model.Car) (uint64, error)
	Get(ctx context.Context, carID uint64) (*model.Car, error)
	List(ctx context.Context,limit uint64, cursor uint64) (model.Cars, error)
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
	query := sq.Select("*").PlaceholderFormat(sq.Dollar).From("car").Where(sq.Eq{"id": carID})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res model.Car
	err = r.db.SelectContext(ctx, &res, s, args...)

	return &res, err
}

func (r *repo) Add(_ context.Context, car *model.Car) (uint64, error) {
	return 0, nil
}

func (r *repo) List(_ context.Context, limit uint64, cursor uint64) (model.Cars, error) {
	return nil, nil
}

func (r *repo) Remove(ctx context.Context, carID uint64) (bool, error) {
	query := sq.Delete("car").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": carID})
	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return true, err
}
