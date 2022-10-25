package services

import (
	"context"
	"github.com/gordrick/honk/pkg/pb/honk"
)

type HonkService struct {
	honk.UnimplementedHonkServiceServer
}

func (s *HonkService) Honk(ctx context.Context, req *honk.HonkRequest) (*honk.HonkResponse, error) {
	return &honk.HonkResponse{
		Message: "Honk!",
	}, nil
}
