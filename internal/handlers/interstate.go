package handlers

import (
	"bank24/internal/config"
	"bank24/internal/service"
	"context"
)

// May contain more than one service
// Or additional fields
type Interstate struct {
	service service.Service
}

func NewInterstate(ctx context.Context, config config.Config) *Interstate {
	service := service.NewService(ctx, config)
	return &Interstate{
		service: service,
	}
}
