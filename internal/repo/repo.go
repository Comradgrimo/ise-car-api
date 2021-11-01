package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/ise-car-api/internal/model"
)

// Repo is DAO for Car
type Repo interface {
	DescribeCar(ctx context.Context, carID uint64) (*model.Car, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeCar(ctx context.Context, carID uint64) (*model.Car, error) {
	return nil, nil
}
