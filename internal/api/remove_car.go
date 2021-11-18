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

func (o *carAPI) RemoveCarV1(ctx context.Context, req *pb.RemoveCarV1Request) (*pb.RemoveCarV1Response, error) {
	metr.EventsCUDTotal.Inc()

	ctx = setLogLevelFromHeader(ctx)
	logger.DebugKV(ctx, fmt.Sprintf("RemoveCarV1 called for carID=%v", req.GetCarId()))

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("RemoveCarV1 -- invalid argument carID=%v", req.CarId))

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	found, err := o.repo.Remove(ctx, req.GetCarId())
	if err != nil {
		logger.ErrorKV(ctx, "RemoveCarV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.RemoveCarV1Response{Found: found}, nil
}
