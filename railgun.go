package cloudflare

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Railgun represents a Railgun's properties.
type Railgun struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Enabled        bool      `json:"enabled"`
	ZonesConnected int       `json:"zones_connected"`
	Build          string    `json:"build"`
	Version        string    `json:"version"`
	Revision       string    `json:"revision"`
	ActivationKey  string    `json:"activation_key"`
	ActivatedOn    time.Time `json:"activated_on"`
	CreatedOn      time.Time `json:"created_on"`
	ModifiedOn     time.Time `json:"modified_on"`
	UpgradeInfo    struct {
		LatestVersion string `json:"latest_version"`
		DownloadLink  string `json:"download_link"`
	} `json:"upgrade_info"`
}

// RailgunListOptions represents the parameters used to list railguns.
type RailgunListOptions struct {
	Direction string
}

// railgunResponse represents the response from the Create Railgun and the Railgun Details endpoints.
type railgunResponse struct {
	Response
	Result Railgun `json:"result"`
}

// railgunsResponse represents the response from the List Railguns endpoint.
type railgunsResponse struct {
	Response
	Result []Railgun `json:"result"`
}

// ZoneRailgun represents the status of a Railgun on a zone.
type ZoneRailgun struct {
	ID        string `json:"id"`
	Name      string `json:"string"`
	Enabled   bool   `json:"enabled"`
	Connected bool   `json:"connected"`
}

// zoneRailgunsResponse represents the response from the zone Railgun endpoint.
type zoneRailgunsResponse struct {
	Response
	Result []ZoneRailgun `json:"result"`
}

// CreateRailgun creates a new Railgun.
// API reference:
// 	https://api.cloudflare.com/#railgun-create-railgun
// 	POST /railguns
func (api *API) CreateRailgun(name string) (Railgun, error) {
	uri := "/railguns"
	params := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	res, err := api.makeRequest("POST", uri, params)
	if err != nil {
		return Railgun{}, errors.Wrap(err, errMakeRequestError)
	}
	var r railgunResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return Railgun{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListRailguns lists Railguns connected to an account.
// API reference:
//  https://api.cloudflare.com/#railgun-list-railguns
//  GET /railguns
func (api *API) ListRailguns(options RailgunListOptions) ([]Railgun, error) {
	v := url.Values{}
	if options.Direction != "" {
		v.Set("direction", options.Direction)
	}
	uri := "/railguns" + "?" + v.Encode()
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r railgunsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// RailgunDetails returns the details for a Railgun.
// API reference:
// 	https://api.cloudflare.com/#railgun-railgun-details
// 	GET /railguns/:identifier
func (api *API) RailgunDetails(railgunID string) (Railgun, error) {
	uri := "/railguns/" + railgunID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return Railgun{}, errors.Wrap(err, errMakeRequestError)
	}
	var r railgunResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return Railgun{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// RailgunZones returns the zones that are currently using a Railgun.
// API reference:
// 	https://api.cloudflare.com/#railgun-get-zones-connected-to-a-railgun
// 	GET /railguns/:identifier/zones
func (api *API) RailgunZones(railgunID string) ([]Zone, error) {
	uri := "/railguns/" + railgunID + "/zones"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r ZonesResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// EnableRailgun enables or disables a Railgun for all zones connected to it.
// API reference:
// 	https://api.cloudflare.com/#railgun-enable-or-disable-a-railgun
// 	PATCH /railguns/:identifier
func (api *API) EnableRailgun(railgunID string, enable bool) (Railgun, error) {
	uri := "/railguns/" + railgunID
	params := struct {
		Enabled bool `json:"enabled"`
	}{
		Enabled: enable,
	}
	res, err := api.makeRequest("PATCH", uri, params)
	if err != nil {
		return Railgun{}, errors.Wrap(err, errMakeRequestError)
	}
	var r railgunResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return Railgun{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteRailgun disables and deletes a Railgun.
// API reference:
// 	https://api.cloudflare.com/#railgun-delete-railgun
// 	DELETE /railguns/:identifier
func (api *API) DeleteRailgun(railgunID string) error {
	uri := "/railguns/" + railgunID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// 	Zone railgun info

// Railguns returns the available Railguns for a zone.
// API reference:
// 	https://api.cloudflare.com/#railguns-for-a-zone-get-available-railguns
// 	GET /zones/:zone_identifier/railguns
func (api *API) Railguns() {
}

// Railgun returns the configuration for a given Railgun.
// API reference:
// 	https://api.cloudflare.com/#railguns-for-a-zone-get-railgun-details
// 	GET /zones/:zone_identifier/railguns/:identifier
func (api *API) Railgun() {
}

// ZoneRailgun connects (true) or disconnects (false) a Railgun for a given zone.
// API reference:
// 	https://api.cloudflare.com/#railguns-for-a-zone-connect-or-disconnect-a-railgun
// 	PATCH /zones/:zone_identifier/railguns/:identifier
func (api *API) ZoneRailgun(connected bool) {
}
