//go:build dev
// +build dev

package main

import (
	"github.com/rs/zerolog"
	"os"
)

func zerologger() zerolog.Logger {
	log := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05",
	}).
		With().
		Timestamp().
		Logger().
		Level(zerolog.DebugLevel)

	log.Info().
		Str("env", "development").
		Str("lvl", log.GetLevel().String()).
		Msg("Logger initialized")

	return log
}
