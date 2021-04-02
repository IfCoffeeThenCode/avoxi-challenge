/*
Package geolite2 provides access to the https://dev.maxmind.com/geoip/geoip2/web-services/
IP location mapping service
*/
package geolite2

// Response is the top-level structure for an API response
type Response struct {
	City               *City               `json:"city"`
	Continent          *Continent          `json:"continent"`
	Country            *Country            `json:"country"`
	RegisteredCountry  *RegisteredCountry  `json:"registered_country"`
	RepresentedCountry *RepresentedCountry `json:"represented_country"`
}

// City holds the responce city data if provided
type City struct {
	Confidence *int `json:"confidence"`
	GeoNameID  int  `json:"geoname_id"`

	// Note: I dislike map[string]string (or worse, map[string]interface{})
	// due to lack of type safety. However, the API docs hint that not only
	// is it possible for nothing to be available here (and if anything it will
	// be English (en)), but they may add other languages later.
	Names map[string]string `json:"names"`
}

// Continent holds the responce continent data
type Continent struct {
	Code      string            `json:"code"`
	GeoNameID int               `json:"geoname_id"`
	Names     map[string]string `json:"names"`
}

// Country holds the data we are really looking for
type Country struct {
	Confidence      *int              `json:"confidence"`
	GeoNameID       int               `json:"geoname_id"`
	InEuropeanUnion bool              `json:"is_in_european_union"`
	ISOCode         string            `json:"iso_code"`
	Names           map[string]string `json:"names"`
}

// RegisteredCountry shows where the IP address is registered
type RegisteredCountry struct {
	GeoNameID       int               `json:"geoname_id"`
	InEuropeanUnion bool              `json:"is_in_european_union"`
	ISOCode         string            `json:"iso_code"`
	Names           map[string]string `json:"names"`
}

// RepresentedCountry shows if an IP address represents a country other than its
// geographic location would suggest (like a military base)
type RepresentedCountry struct {
	GeoNameID       int               `json:"geoname_id"`
	InEuropeanUnion bool              `json:"is_in_european_union"`
	ISOCode         string            `json:"iso_code"`
	Names           map[string]string `json:"names"`
	Type            string            `json:"type"`
}
