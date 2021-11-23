package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/ise-car-api/internal/logger"
	metr "github.com/ozonmp/ise-car-api/internal/metrics"
	"github.com/ozonmp/ise-car-api/internal/model"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
)

func (o *carAPI) AddCarV1(ctx context.Context, req *pb.AddCarV1Request) (*pb.AddCarV1Response, error) {
	metr.EventsCUDTotal.Inc()
	ctx = setLogLevelFromHeader(ctx)

	logger.InfoKV(ctx, fmt.Sprintf("AddCarV1 called: title=%v", req.CarInfo))

	car := model.Car{
		CarInfo:    req.GetCarInfo(),
		UserID:     req.GetUserId(),
		TotalPrice: req.GetTotalPrice(),
		RiskRate:   float64(req.GetRiskRate()),
		CircsLink:  req.GetCircsLink(),
	}
	carID, err := o.repo.Add(ctx, &car)
	return &pb.AddCarV1Response{CarId: carID}, err
}
