package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// AOP represents global AOP (tls_client_auth) metadata.
type AOP struct {
	ID         string    `json:"id"`
	Value      string    `json:"value"`
	Editable   bool      `json:"editable"`
	ModifiedOn time.Time `json:"modified_on"`
}

// AOPResponse represents the response from the global AOP (tls_client_auth) details endpoint.
type AOPResponse struct {
	Response
	Result AOP `json:"result"`
}

// GetAOPStatus returns the configuration details for global AOP (tls_client_auth).
//
// API reference: https://api.cloudflare.com/#zone-settings-get-tls-client-auth-setting
func (api *API) GetAOPStatus(zoneID string) (AOP, error) {
	uri := "/zones/" + zoneID + "/settings/tls_client_auth"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return AOP{}, errors.Wrap(err, errMakeRequestError)
	}
	var r AOPResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AOP{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SetAOPStatus toggles whether global AOP is enabled for the zone.
//
// API reference: https://api.cloudflare.com/#zone-settings-change-tls-client-auth-setting
func (api *API) SetAOPStatus(zoneID string, enable bool) (AOP, error) {
	uri := "/zones/" + zoneID + "/settings/tls_client_auth"
	var val string
	if enable {
		val = "on"
	} else {
		val = "off"
	}
	params := struct {
		Value string `json:"value"`
	}{
		Value: val,
	}
	res, err := api.makeRequest("PATCH", uri, params)
	if err != nil {
		return AOP{}, errors.Wrap(err, errMakeRequestError)
	}
	var r AOPResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AOP{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
