package cloudflare

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// OriginCACertificate represents a Cloudflare-issued certificate.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca
type OriginCACertificate struct {
	ID              string    `json:"id,omitempty"`
	Certificate     string    `json:"certificate,omitempty"`
	Hostnames       []string  `json:"hostnames,omitempty"`
	ExpiresOn       time.Time `json:"expires_on,omitempty"`
	RequestType     string    `json:"request_type,omitempty"`
	RequestValidity int       `json:"requested_validity,omitempty"`
	RevokedAt       time.Time `json:"revoked_at,omitempty"`
	CSR             string    `json:"csr,omitempty"`
}

type originCACertificateResult struct {
	ID              string    `json:"id,omitempty"`
	Certificate     string    `json:"certificate,omitempty"`
	Hostnames       []string  `json:"hostnames,omitempty"`
	ExpiresOn       string    `json:"expires_on,omitempty"`
	RequestType     string    `json:"request_type,omitempty"`
	RequestValidity int       `json:"requested_validity,omitempty"`
	RevokedAt       time.Time `json:"revoked_at,omitempty"`
	CSR             string    `json:"csr,omitempty"`
}

// UnmarshalJSON handles custom parsing from an API response to an OriginCACertificate
func (c *OriginCACertificate) UnmarshalJSON(data []byte) error {
	var res originCACertificateResult
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	// This format comes from time.Time.String() source
	expires, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", res.ExpiresOn)
	if err != nil {
		expires, err = time.Parse(time.RFC3339, res.ExpiresOn)
	}
	if err != nil {
		return err
	}

	*c = OriginCACertificate{
		ID:              res.ID,
		Certificate:     res.Certificate,
		Hostnames:       res.Hostnames,
		ExpiresOn:       expires,
		RequestType:     res.RequestType,
		RequestValidity: res.RequestValidity,
		RevokedAt:       res.RevokedAt,
		CSR:             res.CSR,
	}
	return nil
}

// OriginCACertificateListOptions represents the parameters used to list Cloudflare-issued certificates.
type OriginCACertificateListOptions struct {
	ZoneID string
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

// originCACertificateResponseList represents the response from the List Certificates endpoint.
type originCACertificateResponseList struct {
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
	res, err := api.makeRequestWithAuthType(context.TODO(), "POST", uri, certificate, AuthUserService)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *originCACertificateResponse

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil
}

// OriginCertificates lists all Cloudflare-issued certificates.
//
// This function requires api.APIUserServiceKey be set to your Certificates API key.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca-list-certificates
func (api *API) OriginCertificates(options OriginCACertificateListOptions) ([]OriginCACertificate, error) {
	v := url.Values{}
	if options.ZoneID != "" {
		v.Set("zone_id", options.ZoneID)
	}
	uri := "/certificates" + "?" + v.Encode()
	res, err := api.makeRequestWithAuthType(context.TODO(), "GET", uri, nil, AuthUserService)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *originCACertificateResponseList

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return originResponse.Result, nil
}

// OriginCertificate returns the details for a Cloudflare-issued certificate.
//
// This function requires api.APIUserServiceKey be set to your Certificates API key.
//
// API reference: https://api.cloudflare.com/#cloudflare-ca-certificate-details
func (api *API) OriginCertificate(certificateID string) (*OriginCACertificate, error) {
	uri := "/certificates/" + certificateID
	res, err := api.makeRequestWithAuthType(context.TODO(), "GET", uri, nil, AuthUserService)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *originCACertificateResponse

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
	res, err := api.makeRequestWithAuthType(context.TODO(), "DELETE", uri, nil, AuthUserService)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *originCACertificateResponseRevoke

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil

}
