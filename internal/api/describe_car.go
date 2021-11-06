package api

import (
	"context"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
