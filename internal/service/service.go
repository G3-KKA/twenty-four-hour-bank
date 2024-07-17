package service

import (
	"bank24/internal/config"
	"bank24/internal/database/requests"
	socketpool "bank24/internal/database/socket-pool"
	"context"
)

type Service interface {
	socketpool.DatabaseRequestHandler
}
type service struct {
	pool socketpool.SocketPool
}

// There may be multiple steps before sending request to pool
// Like try-read from cache, to avoid db req at all
func (s *service) Send(req requests.ValidRequest) requests.Response {
	return s.pool.Send(req)
}

func NewService(ctx context.Context, config config.Config) Service {
	pool := socketpool.NewSocketPool(ctx, config)
	return &service{pool: pool}
}
