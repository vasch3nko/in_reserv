package app

import (
	"context"
	"fmt"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/vasch3nko/in_reserv/config"
	"os"
)

// App is the prime struct of the project
type App struct {
	Ctx    context.Context
	Server *fiber.App
	Config *config.Config
	log    zerolog.Logger
}

// NewApp is the constructor of the App struct
func NewApp(ctx context.Context, cfg config.Config, log zerolog.Logger) *App {
	app := fiber.New(fiber.Config{
		AppName:               "InReserv API",
		IdleTimeout:           cfg.Server.IdleTimeout,
		ReadTimeout:           cfg.Server.ReadTimeout,
		WriteTimeout:          cfg.Server.WriteTimeout,
		DisableStartupMessage: true,
		StrictRouting:         true,
	})

	// Registering app hooks
	app.Hooks().OnListen(func(data fiber.ListenData) error {
		log.Info().
			Str("addr", func() string {
				return fmt.Sprintf("%s:%s", data.Host, data.Port)
			}()).
			Uint32("handlers", app.HandlersCount()).
			Int("pid", os.Getpid()).
			Str("tls", func() string {
				if data.TLS {
					return "enabled"
				}
				return "disabled"
			}()).
			Msg("App starting")

		return nil
	})

	// Using the zerolog middleware
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger:          &log,
		FieldsSnakeCase: true,
		Fields:          []string{"ip", "latency", "status", "method", "url", "error", "body"},
	}))

	return &App{Ctx: ctx, Server: app, Config: &cfg, log: log}
}

// Run starts the app
func (a *App) Run() error {
	go a.handleGracefulShutdown()
	return a.Server.Listen(a.Config.Server.Addr)
}

// handleGracefulShutdown handles
// the completion of the context
func (a *App) handleGracefulShutdown() {
	for {
		select {
		case <-a.Ctx.Done():
			a.log.Info().Msg("App shutting down")
			if err := a.Shutdown(); err != nil {
				a.log.Fatal().Err(err).Msg("App shutdown error")
			}
		}
	}
}

// Shutdown shuts down the app
func (a *App) Shutdown() error {
	return a.Server.Shutdown()
}
