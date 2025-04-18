//go:build !dev

package main

import (
	"github.com/rs/zerolog"
	"os"
)

func zerologger() zerolog.Logger {
	log := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)

	log.Info().
		Str("env", "production").
		Str("lvl", log.GetLevel().String()).
		Msg("Logger initialized")

	return log
}
