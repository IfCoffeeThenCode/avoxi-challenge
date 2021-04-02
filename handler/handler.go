package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IfCoffeeThenCode/avoxi-challenge/geolite2"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CheckIP(c *gin.Context)
}

type HandlerActual struct {
	client geolite2.Client
}

func NewHandler(client geolite2.Client) Handler {
	return &HandlerActual{
		client: client,
	}
}

type request struct {
	IP        string   `json:"ip"`
	Countries []string `json:"countries"`
}

func (ha *HandlerActual) CheckIP(c *gin.Context) {
	countryRequest := new(request)
	if err := json.NewDecoder(c.Request.Body).Decode(countryRequest); err != nil {
		c.String(http.StatusInternalServerError, "failed payload decode: %s", err)
		return
	}

	// TODO: Check cache if we have seen this IP before
	// TODO: Determine cache solution

	countryResponse, err := ha.client.GetCountry(countryRequest.IP)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get country: %s", err)
		return
	}
	defer c.Request.Body.Close()

	// TODO: Massage country names; GeoIP provides ISO 3166-1 two-letter codes,
	// and I'm assuming the Avoxi request does likewise

	whitelistCountries := make(map[string]bool, len(countryRequest.Countries))

	for _, country := range countryRequest.Countries {
		whitelistCountries[country] = true
	}

	if permitted := whitelistCountries[countryResponse.Country.ISOCode]; permitted {
		c.String(http.StatusOK, "")
		return
	}

	c.String(http.StatusForbidden, "")
}
