package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-hello-world-controller/assets"
	"github.com/ONSdigital/dp-hello-world-controller/config"
	"github.com/ONSdigital/dp-hello-world-controller/service"
	render "github.com/ONSdigital/dp-renderer"
	"github.com/ONSdigital/log.go/v2/log"
)

func main() {
	log.Namespace = "dp-hello-world-controller"
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(ctx, "application unexpectedly failed", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run(ctx context.Context) error {
	// Create error channel for os signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Create service initialiser and an error channel for service errors
	svcList := service.NewServiceList(&service.Init{})
	svcErrors := make(chan error, 1)

	// Read config
	cfg, err := config.Get()
	if err != nil {
		log.Error(ctx, "unable to retrieve service configuration", err)
		return err
	}

	log.Info(ctx, "got service configuration", log.Data{"config": cfg})

	// Initialise render client, routes and initialise localisations bundles
	renderer := render.NewWithDefaultClient(assets.Asset, assets.AssetNames, cfg.PatternLibraryAssetsPath, cfg.SiteDomain)

	// Run service
	svc := service.New()
	if err := svc.Init(ctx, cfg, svcList, renderer); err != nil {
		log.Error(ctx, "failed to initialise service", err)
		return err
	}
	svc.Run(ctx, svcErrors)

	// Blocks until an os interrupt or a fatal error occurs
	select {
	case err := <-svcErrors:
		log.Error(ctx, "service error received", err)
	case sig := <-signals:
		log.Info(ctx, "os signal received", log.Data{"signal": sig})
	}

	return svc.Close(ctx)
}
