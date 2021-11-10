package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/ise-car-api/internal/model"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
)

func (o *carAPI) AddCarV1(ctx context.Context, req *pb.AddCarV1Request) (*pb.AddCarV1Response, error) {
	log.Debug().Msg(fmt.Sprintf("CreateCarV1 called: title=%v", req.CarInfo))

	car := model.Car{
		CarInfo:    req.GetCarInfo(),
		TotalPrice: float64(req.GetTotalPrice()),
		RiskRate:   float64(req.GetRiskRate()),
	}
	carID, err := o.repo.Add(ctx, &car)
	return &pb.AddCarV1Response{CarId: carID}, err
}
