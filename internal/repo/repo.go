package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/ise-car-api/internal/model"
)

// Repo is DAO for Car
type Repo interface {
	CreateCar(ctx context.Context, carTitle string) (uint64, error)
	DescribeCar(ctx context.Context, carID uint64) (*model.Car, error)
	ListCars(ctx context.Context) (model.Cars, error)
	RemoveCar(ctx context.Context, carID uint64) error
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeCar(_ context.Context, carID uint64) (*model.Car, error) {
	return nil, nil
}

func (r *repo) CreateCar(_ context.Context, carTitle string) (uint64, error) {
	return 0, nil
}

func (r *repo) ListCars(_ context.Context) (model.Cars, error) {
	return model.Cars{nil, nil}, nil
}

func (r *repo) RemoveCar(_ context.Context, carID uint64) error {
	return nil
}
