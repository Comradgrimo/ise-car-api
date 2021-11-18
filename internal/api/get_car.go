package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/ise-car-api/internal/logger"
	metr "github.com/ozonmp/ise-car-api/internal/metrics"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *carAPI) GetCarV1(
	ctx context.Context,
	req *pb.GetCarV1Request,
) (*pb.GetCarV1Response, error) {

	ctx = setLogLevelFromHeader(ctx)
	logger.InfoKV(ctx, fmt.Sprintf("GetCarV1 called: id=%v", req.GetCarId()))

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, "GetCarV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	car, err := o.repo.Get(ctx, req.GetCarId())
	if err != nil {
		logger.ErrorKV(ctx, "GetCarV1 -- failed to get car by id", "carID", req.CarId)
		metr.EventsNotFoundTotal.Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	if car == nil {
		logger.DebugKV(ctx, "GetCarV1 -- car not found", "carID", req.CarId)
		return nil, err
	}
	logger.DebugKV(ctx, "GetCarV1 -- success", "carID", req.CarId)

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
