package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/ozonmp/ise-car-api/internal/repo"

	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
)

var (
	totalCarNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ise_car_api_car_not_found_total",
		Help: "Total number of cars that were not found",
	})
)

type carAPI struct {
	pb.UnimplementedIseCarApiServiceServer
	repo repo.Repo
}

// NewCarAPI returns api of ise-car-api service
func NewCarAPI(r repo.Repo) pb.IseCarApiServiceServer {
	return &carAPI{repo: r}
}

