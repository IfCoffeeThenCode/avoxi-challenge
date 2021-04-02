package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().
		Str("Author", "Ben Durbin").
		Msg("Hello AVOXI")
}
