package main

import (
	"context"
	"flag"
	"github.com/rs/zerolog"
	"github.com/vasch3nko/in_reserv/app"
	"github.com/vasch3nko/in_reserv/config"
	"os"
	"os/signal"
)

// Path to config file in YAML, JSON, TOML or ENV format
var configPath string

func init() {
	flag.StringVar(
		&configPath,
		"config",
		"config.yaml",
		"path to config file (YAML, JSON, TOML or ENV)",
	)
	flag.Parse()
}

func main() {
	// Initializing context with signal
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	// Initializing logger by build tag
	log := zerologger()

	// Launching app
	if err := run(ctx, log); err != nil {
		log.Fatal().Err(err).Msg("App running error")
		os.Exit(1)
	}
}

func run(ctx context.Context, log zerolog.Logger) error {
	// Config initialization
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return err
	}
	log.Info().Str("path", configPath).Msg("Config initialized")

	// Starting application
	application := app.NewApp(ctx, *cfg, log)
	return application.Run()
}
