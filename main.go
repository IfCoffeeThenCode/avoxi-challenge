package main

import (
	"os"

	"github.com/IfCoffeeThenCode/avoxi-challenge/geolite2"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	accountID := os.Getenv("ACCOUNT_ID")
	licenseKey := os.Getenv("LICENSE_KEY")

	geoClient := geolite2.NewClient(accountID, licenseKey)

	myLocation, err := geoClient.Get("me")
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("could not obtain my IP address")
	}

	logger.Info().
		Str("Author", "Ben Durbin").
		Interface("location", myLocation).
		Msg("Hello AVOXI")
}
