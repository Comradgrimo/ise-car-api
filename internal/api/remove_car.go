package api

import (
	"context"
	"fmt"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *carAPI) RemoveCarV1(ctx context.Context, req *pb.RemoveCarV1Request) (*pb.RemoveCarV1Response, error) {
	log.Debug().Msg(fmt.Sprintf("RemoveCarV1 called for %v", req.CarId))

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveCarV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err := o.repo.Remove(ctx, req.GetCarId())
	if err != nil {
		log.Error().Err(err).Msg("RemoveCarV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}
	return new(pb.RemoveCarV1Response), nil
}
