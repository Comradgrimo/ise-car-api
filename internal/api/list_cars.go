package api

import (
	"context"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
)

func (o *carAPI) ListCarsV1(ctx context.Context, in *pb.ListCarsV1Request) (*pb.ListCarsV1Response, error) {
	log.Debug().Msg("ListCarsV1 called")
	return new(pb.ListCarsV1Response), nil
}
