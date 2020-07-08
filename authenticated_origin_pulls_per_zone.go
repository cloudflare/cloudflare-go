package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// PerZoneAOPSettingsResponse represents the settings for Per Zone AOP.
type PerZoneAOPSettings struct {
	Enabled bool `json:"enabled"`
}

// PerZoneAOPSettingsResponse represents the response from the Per Zone AOP settings endpoint.
type PerZoneAOPSettingsResponse struct {
	Response
	Result PerZoneAOPSettings `json:"result"`
}

// PerZoneAOPCertificateDetails represents the metadata for a Per Zone AOP client certificate.
type PerZoneAOPCertificateDetails struct {
	ID          string    `json:"id"`
	Certificate string    `json:"certificate"`
	Issuer      string    `json:"issuer"`
	Signature   string    `json:"signature"`
	ExpiresOn   time.Time `json:"expires_on"`
	Status      string    `json:"status"`
	UploadedOn  time.Time `json:"uploaded_on"`
}

// PerZoneAOPCertificateResponse represents the response from endpoints relating to creating and deleting a per zone AOP certificate.
type PerZoneAOPCertificateResponse struct {
	Response
	Result PerZoneAOPCertificateDetails `json:"result"`
}

// PerZoneAOPCertificatesResponse represents the response from the per zone AOP certificate list endpoint.
type PerZoneAOPCertificatesResponse struct {
	Response
	Result []PerZoneAOPCertificateDetails `json:"result"`
}

// PerZoneAOPCertificateParams represents the required data related to the client certificate being uploaded to be used in Per Zone AOP.
type PerZoneAOPCertificateParams struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

// GetPerZoneAOPStatus returns whether per zone AOP is enabled or not. It is false by default.
//
// API reference: https://api.cloudflare.com/#zone-level-authenticated-origin-pulls-get-enablement-setting-for-zone
func (api *API) GetPerZoneAOPStatus(zoneID string) (PerZoneAOPSettings, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/settings"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return PerZoneAOPSettings{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerZoneAOPSettingsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerZoneAOPSettings{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SetPerZoneAOPStatus will update whether Per Zone AOP is enabled for the zone.
//
// API reference: https://api.cloudflare.com/#zone-level-authenticated-origin-pulls-set-enablement-for-zone
func (api *API) SetPerZoneAOPStatus(zoneID string, enable bool) (PerZoneAOPSettings, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/settings"
	params := struct {
		Enabled bool `json:"enabled"`
	}{
		Enabled: enable,
	}
	res, err := api.makeRequest("PUT", uri, params)
	if err != nil {
		return PerZoneAOPSettings{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerZoneAOPSettingsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerZoneAOPSettings{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UploadPerZoneAOPCertificate will upload a provided client certificate and enable it to be used in all AOP requests for the zone.
//
// API reference: https://api.cloudflare.com/#zone-level-authenticated-origin-pulls-upload-certificate
func (api *API) UploadPerZoneAOPCertificate(zoneID string, params PerZoneAOPCertificateParams) (PerZoneAOPCertificateDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth"
	res, err := api.makeRequest("POST", uri, params)
	if err != nil {
		return PerZoneAOPCertificateDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerZoneAOPCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerZoneAOPCertificateDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListPerZoneAOPCertificates returns a list of all user uploaded client certificates to Per Zone AOP.
//
// API reference: https://api.cloudflare.com/#zone-level-authenticated-origin-pulls-list-certificates
func (api *API) ListPerZoneAOPCertificates(zoneID string) ([]PerZoneAOPCertificateDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []PerZoneAOPCertificateDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerZoneAOPCertificatesResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []PerZoneAOPCertificateDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// GetPerZoneAOPCertificateDetails returns the metadata associated with a user uploaded client certificate to Per Zone AOP.
//
// API reference: https://api.cloudflare.com/#zone-level-authenticated-origin-pulls-get-certificate-details
func (api *API) GetPerZoneAOPCertificateDetails(zoneID, certificateID string) (PerZoneAOPCertificateDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/" + certificateID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return PerZoneAOPCertificateDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerZoneAOPCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerZoneAOPCertificateDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeletePerZoneAOPCertificate removes the specified client certificate from the edge.
//
// API reference: https://api.cloudflare.com/#zone-level-authenticated-origin-pulls-delete-certificate
func (api *API) DeletePerZoneAOPCertificate(zoneID, certificateID string) (PerZoneAOPCertificateDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/" + certificateID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return PerZoneAOPCertificateDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerZoneAOPCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerZoneAOPCertificateDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
