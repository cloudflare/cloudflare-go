package cloudflare

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// CustomHostnameSSL represents the SSL section in a given custom hostname.
type CustomHostnameSSL struct {
	Status      string `json:"status,omitempty"`
	Method      string `json:"method,omitempty"`
	Type        string `json:"type,omitempty"`
	CnameTarget string `json:"cname_target,omitempty"`
	CnameName   string `json:"cname_name,omitempty"`
}

// CustomHostname represents a custom hostname in a zone.
type CustomHostname struct {
	ID       string            `json:"id,omitempty"`
	Hostname string            `json:"hostname,omitempty"`
	SSL      CustomHostnameSSL `json:"ssl,omitempty"`
}

type CustomHostnameResponse struct {
	Result CustomHostname `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type CustomHostnameListResponse struct {
	Result []CustomHostname `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// CreateCustomHostname creaets a new custom hostname and requests that an SSL certificate be issued for it.
//
// API reference: https://api.cloudflare.com/#custom-hostname-for-a-zone-create-custom-hostname
func (api *API) CreateCustomHostname(zoneID string, ch CustomHostname) (*CustomHostnameResponse, error) {
	uri := "/zones/" + zoneID + "/custom_hostnames"
	res, err := api.makeRequest("POST", uri, ch)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var response *CustomHostnameResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response, nil
}

// CreateCustomHostnames lists custom hostnames for a given zone.
//
// API reference: https://api.cloudflare.com/#custom-hostname-for-a-zone-list-custom-hostnames
func (api *API) CustomHostnames(zoneID string) ([]CustomHostname, error) {
	uri := "/zones/" + zoneID + "/custom_hostnames"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []CustomHostname{}, errors.Wrap(err, errMakeRequestError)
	}

	var response CustomHostnameListResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return []CustomHostname{}, errors.Wrap(err, errUnmarshalError)
	}

	return response.Result, nil
}

// CreateCustomHostname inspects a given custom hostname in a given zone
//
// API reference: https://api.cloudflare.com/#custom-hostname-for-a-zone-custom-hostname-configuration-details
func (api *API) CustomHostname(zoneID string, customHostnameID string) (CustomHostname, error) {
	uri := "/zones/" + zoneID + "/custom_hostnames/" + customHostnameID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return CustomHostname{}, errors.Wrap(err, errMakeRequestError)
	}

	var response CustomHostnameResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CustomHostname{}, errors.Wrap(err, errUnmarshalError)
	}

	return response.Result, nil
}

// CustomHostnameIDByName retrieves a custom hostname's ID from the hostname.
func (api *API) CustomHostnameIDByName(zoneID string, hostname string) (string, error) {
	res, err := api.CustomHostnames(zoneID)
	if err != nil {
		return "", errors.Wrap(err, "CustomHostnames command failed")
	}
	for _, ch := range res {
		if ch.Hostname == hostname {
			return ch.ID, nil
		}
	}
	return "", errors.New("Zone could not be found")
}
