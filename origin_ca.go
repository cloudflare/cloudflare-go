package cloudflare

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// OriginCA resource
// This is the origin_ca resource definition (as defined here:
// https://api.cloudflare.com/#origin-ca || https://api.cloudflare.com/#cloudflare-ca
type OriginCA struct {
	ID              string   `json:"id"`
	Certificate     string   `json:"certificate"`
	Hostnames       []string `json:"hostnames"`
	ExpiresOn       string   `json:"expires_on"`
	RequestType     string   `json:"request_type"`
	RequestValidity string   `json:"requested_validity"`
	Csr             string   `json:"csr"`
}

// OriginCertificateID is the resource sent back for a revoke action
type OriginCertificateID struct {
	ID string `json:"id"`
}

// OriginCAResponse is the APIv4 response envelop containing the OriginCA result
type OriginCAResponse struct {
	Response
	Result OriginCA `json:"result"`
}

// OriginCAResponseList is the APIv4 response envelop containing a listof OriginCA result
type OriginCAResponseList struct {
	Response
	Result     []OriginCA `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// OriginCAResponseRevoke is the APIv4 response envelop containing a revoked cert id
type OriginCAResponseRevoke struct {
	Response
	Result OriginCertificateID `json:"result"`
}

// CreateOriginCertificate will create an origin certificate for a User
// API reference: https://api.cloudflare.com/#origin-ca-create-certificate
func (api *API) CreateOriginCertificate(certificate OriginCA) (*OriginCA, error) {
	uri := "/certificates"
	res, err := api.makeRequest("POST", uri, certificate)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCAResponse

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	createdCert := originResponse.Result

	return &createdCert, nil
}

// OriginCertificates will list all certificates owned by the users
// API reference: https://api.cloudflare.com/#origin-ca-list-certificates
func (api *API) OriginCertificates() ([]OriginCA, error) {
	uri := "/certificates"
	res, err := api.makeRequest("GET", uri, nil)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCAResponseList

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
// API reference: https://api.cloudflare.com/#origin-ca-certificate-details
func (api *API) OriginCertificate(certificateID string) (*OriginCA, error) {
	uri := "/certificates/" + certificateID
	res, err := api.makeRequest("GET", uri, nil)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCAResponse

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
// API reference: https://api.cloudflare.com/#origin-ca-revoke-certificate
func (api *API) RevokeOriginCertificate(certificateID string) (*OriginCertificateID, error) {
	uri := "/certificates/" + certificateID
	res, err := api.makeRequest("DELETE", uri, nil)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var originResponse *OriginCAResponseRevoke

	err = json.Unmarshal(res, &originResponse)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !originResponse.Success {
		return nil, errors.New(errRequestNotSuccessful)
	}

	return &originResponse.Result, nil

}
