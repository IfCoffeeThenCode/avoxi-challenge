package main

import (
	"os"

	"github.com/IfCoffeeThenCode/avoxi-challenge/geolite2"
	"github.com/rs/zerolog"
)

func main() {
	myLocation := geolite2.Response{
		City: &geolite2.City{
			Names: map[string]string{
				"en": "Atlanta",
				"ru": "Атланта",
			},
		},
		Country: &geolite2.Country{
			ISOCode: "US",
			Names: map[string]string{
				"en": "United States",
				"ru": "США",
			},
		},
		Continent: &geolite2.Continent{
			Code: geolite2.NA,
			Names: map[string]string{
				"en": "North America",
				"ru": "Северная Америка",
			},
		},
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().
		Str("Author", "Ben Durbin").
		Interface("location", myLocation).
		Msg("Hello AVOXI")
}
