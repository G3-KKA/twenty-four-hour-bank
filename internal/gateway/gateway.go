package gateway

import (
	"bank24/internal/config"
	"bank24/internal/handlers"
	"bank24/internal/handlers/middleware"
	log "bank24/internal/logger"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type Gateway interface {
	ListenAndServe(ctx context.Context, config config.Config) error
}
type gateway struct {
	is *handlers.Interstate
}

func New(ctx context.Context, config config.Config) (Gateway, error) {

	is := handlers.NewInterstate(ctx, config)
	return &gateway{
		is: is,
	}, nil
}
func (gateway *gateway) ListenAndServe(ctx context.Context, config config.Config) error {
	r := chi.NewRouter()
	r.Use(middleware.Logging)
	r.Use(chiMiddleware.Recoverer)
	r.Post("/accounts", gateway.is.CreateAccount)
	r.Route("/accounts/{id}", func(r chi.Router) {
		r.Post("/deposit", gateway.is.Deposit)
		r.Post("/withdraw", gateway.is.Withdraw)
		r.Get("/balance", gateway.is.Balance)
	})

	log.Info("Listening on :8080")
	err := http.ListenAndServe(":8080", r)
	return err
}
