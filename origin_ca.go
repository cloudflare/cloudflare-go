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
	Success  bool     `json:"success"`
	Errors   []string `json:"errors"`
	Messages []string `json:"messages"`
	Result   OriginCA `json:"result"`
}

// OriginCAResponseList is the APIv4 response envelop containing a listof OriginCA result
type OriginCAResponseList struct {
	Success    bool       `json:"success"`
	Errors     []string   `json:"errors"`
	Messages   []string   `json:"messages"`
	Result     []OriginCA `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// OriginCAResponseRevoke is the APIv4 response envelop containing a revoked cert id
type OriginCAResponseRevoke struct {
	Success  bool                `json:"success"`
	Errors   []string            `json:"errors"`
	Messages []string            `json:"messages"`
	Result   OriginCertificateID `json:"result"`
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
		return nil, errors.New(errRequestNotSuccessfull)
	}

	createdCert := originResponse.Result

	return &createdCert, nil
}
