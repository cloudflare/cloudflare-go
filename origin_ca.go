package cloudflare

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// OriginCACertificate represents a Cloudflare-issued certificate.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca
type OriginCACertificate struct {
	ID              string    `json:"id"`
	Certificate     string    `json:"certificate"`
	Hostnames       []string  `json:"hostnames"`
	ExpiresOn       time.Time `json:"expires_on"`
	RequestType     string    `json:"request_type"`
	RequestValidity int       `json:"requested_validity"`
	CSR             string    `json:"csr"`
}

// OriginCACertificateID represents the ID of the revoked certificate from the Revoke Certificate endpoint.
type OriginCACertificateID struct {
	ID string `json:"id"`
}

// originCACertificateResponse represents the response from the Create Certificate and the Certificate Details endpoints.
type originCACertificateResponse struct {
	Response
	Result OriginCACertificate `json:"result"`
}

// OriginCACertificateFilter represents the object used to filter results.
type OriginCACertificateFilter struct {
	ZoneID string
}

// OriginCACertificateListResponse represents the response from the List Certificates endpoint.
type OriginCACertificateListResponse struct {
	Response
	Result     []OriginCACertificate `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// originCACertificateResponseRevoke represents the response from the Revoke Certificate endpoint.
type originCACertificateResponseRevoke struct {
	Response
	Result OriginCACertificateID `json:"result"`
}

// CreateOriginCertificate creates a Cloudflare-signed certificate.
//
// This function requires api.APIUserServiceKey be set to your Certificates API key.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca-create-certificate
func (api *API) CreateOriginCertificate(certificate OriginCACertificate) (*OriginCACertificate, error) {
	uri := "/certificates"
	res, err := api.makeRequestWithAuthType("POST", uri, certificate, AuthUserService)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse originCACertificateResponse

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil
}

// ListOriginCACertificates lists all Cloudflare-issued certificates.
//
// This function requires api.APIUserServiceKey be set to your Certificates API key.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca-list-certificates
func (api *API) ListOriginCACertificates(page int) (*OriginCACertificateListResponse, error) {
	return api.FilterOriginCACertificates(page, OriginCACertificateFilter{})
}

// FilterOriginCACertificates lists all Cloudflare-issued certificates, filtered by zone.
//
// This function requires api.APIUserServiceKey be set to your Certificates API key.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca-list-certificates
func (api *API) FilterOriginCACertificates(page int, filter OriginCACertificateFilter) (*OriginCACertificateListResponse, error) {
	v := url.Values{}
	if page <= 0 {
		page = 1
	}

	v.Set("page", strconv.Itoa(page))
	v.Set("per_page", strconv.Itoa(100))
	if filter.ZoneID != "" {
		v.Set("zone_id", filter.ZoneID)
	}
	query := "?" + v.Encode()

	uri := "/certificates" + query
	res, err := api.makeRequestWithAuthType("GET", uri, nil, AuthUserService)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	response := &OriginCACertificateListResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response, nil
}

// OriginCertificate returns the details for a Cloudflare-issued certificate.
//
// This function requires api.APIUserServiceKey be set to your Certificates API key.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca-certificate-details
func (api *API) OriginCertificate(certificateID string) (*OriginCACertificate, error) {
	uri := "/certificates/" + certificateID
	res, err := api.makeRequestWithAuthType("GET", uri, nil, AuthUserService)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse originCACertificateResponse

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil
}

// RevokeOriginCertificate revokes a created certificate for a zone.
//
// This function requires api.APIUserServiceKey be set to your Certificates API key.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca-revoke-certificate
func (api *API) RevokeOriginCertificate(certificateID string) (*OriginCACertificateID, error) {
	uri := "/certificates/" + certificateID
	res, err := api.makeRequestWithAuthType("DELETE", uri, nil, AuthUserService)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse originCACertificateResponseRevoke

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil

}
