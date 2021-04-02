package main

import (
	"os"

	"github.com/IfCoffeeThenCode/avoxi-challenge/geolite2"
	"github.com/IfCoffeeThenCode/avoxi-challenge/handler"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	accountID := os.Getenv("ACCOUNT_ID")
	licenseKey := os.Getenv("LICENSE_KEY")

	geoClient := geolite2.NewClient(accountID, licenseKey)

	handler := handler.NewHandler(geoClient)

	router := gin.Default()
	router.GET("/whitelist", handler.CheckIP)

	router.Run(os.Getenv("GEO_HOST_PORT"))
}
