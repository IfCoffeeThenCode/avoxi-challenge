package geolite2

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

type Client interface {
	Get(ipAddress string) (*Response, error)
}

type clientActual struct {
	client     *http.Client
	accountID  string
	licenseKey string
}

func NewClient(accountID, licenseKey string) Client {
	return &clientActual{
		client:     new(http.Client),
		accountID:  accountID,
		licenseKey: licenseKey,
	}
}

func (ca *clientActual) Get(ipAddress string) (*Response, error) {
	// Create request URL
	geoipURL := &url.URL{
		Scheme: "https",
		Host:   "geoip.maxmind.com",
	}
	geoipURL.Path = path.Join("geoip/v2.1/country", ipAddress)

	// Perform request
	// TODO: add headers for auth
	response, err := ca.client.Get(geoipURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf("bad response code %s", response.Status)
	}

	// Decode/Unmarshall response
	output := new(Response)
	if err := json.NewDecoder(response.Body).Decode(output); err != nil {
		return nil, errors.Wrap(err, "failed payload decode")
	}

	return output, nil
}
