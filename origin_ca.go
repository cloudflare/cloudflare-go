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
	RequestValidity string   `json:"request_validity"`
	Csr             string   `json:"csr"`
}

// CreateOriginCertificate will create an origin certificate for a User
// API reference: https://api.cloudflare.com/#origin-ca-create-certificate
func (api *API) CreateOriginCertificate(certificate OriginCA) (*OriginCA, error) {
	uri := "/certificates"
	res, err := api.makeRequest("POST", uri, certificate)

	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var createdCert *OriginCA

	err = json.Unmarshal(res, &createdCert)

	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return createdCert, nil
}
