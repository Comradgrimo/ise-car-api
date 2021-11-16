package api

import (
	"context"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *carAPI) GetCarV1(
	ctx context.Context,
	req *pb.GetCarV1Request,
) (*pb.GetCarV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("GetCarV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	car, err := o.repo.Get(ctx, req.GetCarId())
	if err != nil {
		log.Error().Err(err).Msg("GetCarV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if car == nil {
		log.Debug().Uint64("carId", req.CarId).Msg("car not found")
		return nil, err
	}

	log.Debug().Msg("GetCarV1 - success")
	return &pb.GetCarV1Response{
		Value: &pb.Car{
			Id:         car.ID,
			CarInfo:    car.CarInfo,
			UserId:     car.UserID,
			TotalPrice: car.TotalPrice,
			RiskRate:   float32(car.RiskRate),
			CircsLink:  car.CircsLink,
		},
	}, nil
}
