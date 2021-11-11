package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/ise-car-api/internal/model"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
)

func (o *carAPI) AddCarV1(ctx context.Context, req *pb.AddCarV1Request) (*pb.AddCarV1Response, error) {
	log.Debug().Msg(fmt.Sprintf("AddCarV1 called: title=%v", req.CarInfo))

	car := model.Car{
		CarInfo:    req.GetCarInfo(),
		UserID:     req.GetUserId(),
		TotalPrice: float64(req.GetTotalPrice()),
		RiskRate:   float64(req.GetRiskRate()),
		CircsLink:  req.GetCircsLink(),
	}
	carID, err := o.repo.Add(ctx, &car)
	return &pb.AddCarV1Response{CarId: carID}, err
}
