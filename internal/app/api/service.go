package api

import (
	"context"
)

type Service struct {
	ctx context.Context
}

func New(ctx context.Context) *Service {
	service := &Service{
		ctx: ctx,
	}
	return service
}
