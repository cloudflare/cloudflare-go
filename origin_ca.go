package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// OriginCACertificate resource
// This is the origin_ca resource definition as defined here:
// https://api.cloudflare.com/#cloudflare-ca
type OriginCACertificate struct {
	ID              string    `json:"id"`
	Certificate     string    `json:"certificate"`
	Hostnames       []string  `json:"hostnames"`
	ExpiresOn       time.Time `json:"expires_on"`
	RequestType     string    `json:"request_type"`
	RequestValidity string    `json:"requested_validity"`
	Csr             string    `json:"csr"`
}

// OriginCACertificateID is the resource sent back for a revoke action
type OriginCACertificateID struct {
	ID string `json:"id"`
}

// OriginCACertificateResponse is the APIv4 response envelop containing the OriginCA result
type OriginCACertificateResponse struct {
	Response
	Result OriginCACertificate `json:"result"`
}

// OriginCACertificateResponseList is the APIv4 response envelop containing a listof OriginCA result
type OriginCACertificateResponseList struct {
	Response
	Result     []OriginCACertificate `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// OriginCACertificateResponseRevoke is the APIv4 response envelop containing a revoked cert id
type OriginCACertificateResponseRevoke struct {
	Response
	Result OriginCACertificateID `json:"result"`
}

// CreateOriginCertificate will create an origin certificate for a User
// API reference: https://api.cloudflare.com/#cloudflare-ca-create-certificate
func (api *API) CreateOriginCertificate(certificate OriginCACertificate) (*OriginCACertificate, error) {
	uri := "/certificates"
	res, err := api.makeRequest("POST", uri, certificate)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCACertificateResponse

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil
}

// OriginCertificates will list all certificates owned by the users
// API reference: https://api.cloudflare.com/#cloudflare-ca-list-certificates
func (api *API) OriginCertificates() ([]OriginCACertificate, error) {
	uri := "/certificates"
	res, err := api.makeRequest("GET", uri, nil)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCACertificateResponseList

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return originResponse.Result, nil
}

// OriginCertificate will get the details for a given certificate
// API reference: https://api.cloudflare.com/#cloudflare-ca-certificate-details
func (api *API) OriginCertificate(certificateID string) (*OriginCACertificate, error) {
	uri := "/certificates/" + certificateID
	res, err := api.makeRequest("GET", uri, nil)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCACertificateResponse

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil
}

// RevokeOriginCertificate will revoke a given certificate
// API reference: https://api.cloudflare.com/#cloudflare-ca-revoke-certificate
func (api *API) RevokeOriginCertificate(certificateID string) (*OriginCACertificateID, error) {
	uri := "/certificates/" + certificateID
	res, err := api.makeRequest("DELETE", uri, nil)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCACertificateResponseRevoke

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil

}
