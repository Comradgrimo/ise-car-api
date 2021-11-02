package api

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (o *carAPI) DescribeCarV1(
	ctx context.Context,
	req *pb.DescribeCarV1Request,
) (*pb.DescribeCarV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeCarV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	car, err := o.repo.DescribeCar(ctx, req.CarId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeCarV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if car == nil {
		log.Debug().Uint64("carId", req.CarId).Msg("car not found")
		totalCarNotFound.Inc()

		return nil, status.Error(codes.NotFound, "car not found")
	}

	log.Debug().Msg("DescribeCarV1 - success")

	return &pb.DescribeCarV1Response{
		Value: &pb.Car{
			Id:    car.ID,
			Title: car.Title,
		},
	}, nil
}

func (o *carAPI) ListCarsV1(ctx context.Context, in *pb.ListCarsV1Request) (*pb.ListCarsV1Response, error) {
	log.Debug().Msg("ListCarsV1 called")
	return new(pb.ListCarsV1Response), nil
}

func (o *carAPI) CreateCarV1(ctx context.Context, in *pb.CreateCarV1Request) (*pb.CreateCarV1Response, error) {
	log.Debug().Msg(fmt.Sprintf("CreateCarV1 called: title=%v", in.Title))
	return new(pb.CreateCarV1Response), nil
}

func (o *carAPI) RemoveCarV1(ctx context.Context, in *pb.RemoveCarV1Request) (*pb.RemoveCarV1Response, error) {
	log.Debug().Msg(fmt.Sprintf("RemoveCarV1 called for %v", in.CarId))
	return new(pb.RemoveCarV1Response), nil
}
