package service

import (
	"context"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-hello-world-api/api"
	"github.com/ONSdigital/dp-hello-world-api/config"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Service struct {
	Config      *config.Config
	server      *server.Server
	Router      *mux.Router
	API         *api.API
	HealthCheck *healthcheck.HealthCheck
}

// Run the service
func Run(buildTime, gitCommit, version string, svcErrors chan error) (*Service, error) {
	ctx := context.Background()
	log.Event(ctx, "running service", log.INFO)

	cfg, err := config.Get()
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve service configuration")
	}
	log.Event(ctx, "got service configuration", log.Data{"config": cfg}, log.INFO)

	r := mux.NewRouter()

	s := server.New(cfg.BindAddr, r)
	s.HandleOSSignals = false

	a := api.Setup(ctx, r)

	versionInfo, err := healthcheck.NewVersionInfo(
		buildTime,
		gitCommit,
		version,
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse version information")
	}
	hc := healthcheck.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)
	if err := registerCheckers(ctx, &hc); err != nil {
		return nil, errors.Wrap(err, "unable to register checkers")
	}
	r.StrictSlash(true).Path("/health").HandlerFunc(hc.Handler)
	hc.Start(ctx)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			svcErrors <- errors.Wrap(err, "failure in http listen and serve")
		}
	}()

	return &Service{
		Config:      cfg,
		Router:      r,
		API:         a,
		HealthCheck: &hc,
		server:      s,
	}, nil
}

// Gracefully shutdown the service
func (svc *Service) Close(ctx context.Context) {
	timeout := svc.Config.GracefulShutdownTimeout
	log.Event(ctx, "commencing graceful shutdown", log.Data{"graceful_shutdown_timeout": timeout}, log.INFO)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// stop any incoming requests before closing any outbound connections
	if err := svc.server.Shutdown(ctx); err != nil {
		log.Event(ctx, "failed to shutdown http server", log.Error(err), log.ERROR)
	}

	if err := svc.API.Close(ctx); err != nil {
		log.Event(ctx, "error closing API", log.Error(err), log.ERROR)
	}

	log.Event(ctx, "graceful shutdown complete", log.INFO)
}

func registerCheckers(ctx context.Context, hc *healthcheck.HealthCheck) (err error) {
	// TODO ADD HEALTH CHECKS HERE
	return
}
