package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// PerHostnameAOPCertificateDetails represents the metadata for a Per Hostname AOP certificate.
type PerHostnameAOPCertificateDetails struct {
	ID           string    `json:"id"`
	Certificate  string    `json:"certificate"`
	Issuer       string    `json:"issuer"`
	Signature    string    `json:"signature"`
	SerialNumber string    `json:"serial_number"`
	ExpiresOn    time.Time `json:"expires_on"`
	Status       string    `json:"status"`
	UploadedOn   time.Time `json:"uploaded_on"`
}

// PerHostnameAOPCertificateResponse represents the response from endpoints relating to creating and deleting a Per Hostname AOP certificate.
type PerHostnameAOPCertificateResponse struct {
	Response
	Result PerHostnameAOPCertificateDetails `json:"result"`
}

// PerHostnameAOPDetails contains metadata about the Per Hostname AOP configuration on a hostname.
type PerHostnameAOPDetails struct {
	Hostname       string    `json:"hostname"`
	CertID         string    `json:"cert_id"`
	Enabled        bool      `json:"enabled"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CertStatus     string    `json:"cert_status"`
	Issuer         string    `json:"issuer"`
	Signature      string    `json:"signature"`
	SerialNumber   string    `json:"serial_number"`
	Certificate    string    `json:"certificate"`
	CertUploadedOn time.Time `json:"cert_uploaded_on"`
	CertUpdatedAt  time.Time `json:"cert_updated_at"`
	ExpiresOn      time.Time `json:"expires_on"`
}

// PerHostnamesAOPDetailsResponse represents Per Hostname AOP configuration metadata for a single hostname.
type PerHostnameAOPDetailsResponse struct {
	Response
	Result PerHostnameAOPDetails `json:"result"`
}

// PerHostnamesAOPDetailsResponse represents Per Hostname AOP configuration metadata for multiple hostnames.
type PerHostnamesAOPDetailsResponse struct {
	Response
	Result []PerHostnameAOPDetails `json:"result"`
}

// PerHostnameAOPCertificateParams represents the required data related to the client certificate being uploaded to be used in Per Hostname AOP.
type PerHostnameAOPCertificateParams struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

// PerHostnameAOPConfig represents the config state for Per Hostname AOP applied on a hostname.
type PerHostnameAOPConfig struct {
	Hostname string `json:"hostname"`
	CertID   string `json:"cert_id"`
	Enabled  bool   `json:"enabled"`
}

// UploadPerHostnameAOPCertificate will upload the provided certificate and private key to the edge under Per Hostname AOP.
//
// API reference: https://api.cloudflare.com/#per-hostname-authenticated-origin-pull-upload-a-hostname-client-certificate
func (api *API) UploadPerHostnameAOPCertificate(zoneID string, params PerHostnameAOPCertificateParams) (PerHostnameAOPCertificateDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/hostnames/certificates"
	res, err := api.makeRequest("POST", uri, params)
	if err != nil {
		return PerHostnameAOPCertificateDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerHostnameAOPCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerHostnameAOPCertificateDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// GetPerHostnameAOPCertificate retrieves certificate metadata about the requested Per Hostname certificate.
//
// API reference: https://api.cloudflare.com/#per-hostname-authenticated-origin-pull-get-the-hostname-client-certificate
func (api *API) GetPerHostnameAOPCertificate(zoneID, certificateID string) (PerHostnameAOPCertificateDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/hostnames/certificates/" + certificateID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return PerHostnameAOPCertificateDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerHostnameAOPCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerHostnameAOPCertificateDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeletePerHostnameAOPCertificate will remove the requested Per Hostname certificate from the edge.
//
// API reference: https://api.cloudflare.com/#per-hostname-authenticated-origin-pull-delete-hostname-client-certificate
func (api *API) DeletePerHostnameAOPCertificate(zoneID, certificateID string) (PerHostnameAOPCertificateDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/hostnames/certificates/" + certificateID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return PerHostnameAOPCertificateDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerHostnameAOPCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerHostnameAOPCertificateDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// EditPerHostnameAOPConfig applies the supplied Per Hostname AOP config onto a hostname(s) in the edge.
//
// API reference: https://api.cloudflare.com/#per-hostname-authenticated-origin-pull-enable-or-disable-a-hostname-for-client-authentication
func (api *API) EditPerHostnameAOPConfig(zoneID string, config []PerHostnameAOPConfig) ([]PerHostnameAOPDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/hostnames"
	res, err := api.makeRequest("PUT", uri, nil)
	if err != nil {
		return []PerHostnameAOPDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerHostnamesAOPDetailsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []PerHostnameAOPDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// GetPerHostnameAOPConfig returns the config state of Per Hostname AOP of the provided hostname within a zone.
//
// API reference: https://api.cloudflare.com/#per-hostname-authenticated-origin-pull-get-the-hostname-status-for-client-authentication
func (api *API) GetPerHostnameAOPConfig(zoneID, hostname string) (PerHostnameAOPDetails, error) {
	uri := "/zones/" + zoneID + "/origin_tls_client_auth/hostnames/" + hostname
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return PerHostnameAOPDetails{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PerHostnameAOPDetailsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return PerHostnameAOPDetails{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
