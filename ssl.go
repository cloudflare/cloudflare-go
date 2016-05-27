package cloudflare

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// CreateSSL allows you to add a custom SSL certificate to the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-create-ssl-configuration
// 	POST /zones/:zone_identifier/custom_certificates
func (api *API) CreateSSL(zoneID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates"
	res, err := api.makeRequest("POST", uri, options)
	if err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errMakeRequestError)
	}
	var r ZoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListSSL lists the custom certificates for the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-list-ssl-configurations
// 	GET /zones/:zone_identifier/custom_certificates
func (api *API) ListSSL(zoneID string) ([]ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r ZoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SSLDetails returns the configuration details for a custom SSL certificate.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-ssl-configuration-details
// 	GET /zones/:zone_identifier/custom_certificates/:identifier
func (api *API) SSLDetails(zoneID, certificateID string) (ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errMakeRequestError)
	}
	var r ZoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateSSL updates (replaces) a custom SSL certificate.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-update-ssl-configuration
// 	PATCH /zones/:zone_identifier/custom_certificates/:identifier
func (api *API) UpdateSSL(zoneID, certificateID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	res, err := api.makeRequest("PATCH", uri, options)
	if err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errMakeRequestError)
	}
	var r ZoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ReprioritizeSSL allows you to change the priority (which is served for a given
// request) of custom SSL certificates associated with the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-re-prioritize-ssl-certificates
// 	PUT /zones/:zone_identifier/custom_certificates/prioritize
func (api *API) ReprioritizeSSL(zoneID string, p []ZoneCustomSSLPriority) ([]ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates/prioritize"
	params := struct {
		Certificates []ZoneCustomSSLPriority `json:"certificates"`
	}{
		Certificates: p,
	}
	res, err := api.makeRequest("PATCH", uri, params)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r ZoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteSSL deletes a custom SSL certificate from the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-delete-an-ssl-certificate
// 	DELETE /zones/:zone_identifier/custom_certificates/:identifier
func (api *API) DeleteSSL(zoneID, certificateID string) error {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}
