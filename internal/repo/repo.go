package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/ise-car-api/internal/model"
)

// Repo is DAO for Car
type Repo interface {
	DescribeTemplate(ctx context.Context, templateID uint64) (*model.Car, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeTemplate(ctx context.Context, templateID uint64) (*model.Car, error) {
	return nil, nil
}
