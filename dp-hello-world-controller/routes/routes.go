package routes

import (
	"context"

	"github.com/ONSdigital/dp-hello-world-controller/config"
	"github.com/ONSdigital/dp-hello-world-controller/handlers"

	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

// Setup registers routes for the service
func Setup(ctx context.Context, r *mux.Router, cfg *config.Config, hc health.HealthCheck) {
	log.Info(ctx, "adding routes")
	r.StrictSlash(true).Path("/health").HandlerFunc(hc.Handler)
	r.StrictSlash(true).Path("/helloworld").Methods("GET").HandlerFunc(handlers.HelloWorld(*cfg))
}
