package api

import (
	"context"
	"fmt"
	pb "github.com/ozonmp/ise-car-api/pkg/ise-car-api"
	"github.com/rs/zerolog/log"
)

func (o *carAPI) CreateCarV1(ctx context.Context, req *pb.CreateCarV1Request) (*pb.CreateCarV1Response, error) {
	log.Debug().Msg(fmt.Sprintf("CreateCarV1 called: title=%v", req.Title))
	return new(pb.CreateCarV1Response), nil
}
