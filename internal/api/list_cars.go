package api

import (
	"context"
	"github.com/ozonmp/ise-car-api/internal/model"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
)

func (o *carAPI) ListCarsV1(ctx context.Context, req *pb.ListCarsV1Request) (*pb.ListCarsV1Response, error) {
	log.Debug().Msg("ListCarsV1 called")
	dbCars, err := o.repo.List(ctx, req.GetCursor(), req.GetLimit())
	if err != nil {
		return nil, err
	}

	cars := make([]*pb.Car, len(dbCars))
	for idx := range dbCars {
		cars[idx] = convertCarToProtobuf(&dbCars[idx])
	}

	return &pb.ListCarsV1Response{
		Cars: cars,
	}, nil
}

func convertCarToProtobuf(orig *model.Car) *pb.Car {
	return &pb.Car{
		Id:         orig.ID,
		CarInfo:    orig.CarInfo,
		UserId:     orig.UserID,
		TotalPrice: float32(orig.TotalPrice),
		RiskRate:   float32(orig.RiskRate),
		CircsLink:  orig.CircsLink,
	}
}
