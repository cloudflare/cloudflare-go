package cloudflare

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ZoneSSLVerification represents custom SSL Verfication metadata.
type ZoneSSLVerification struct {
	CertificateStatus string `json:"certificate_status"`
	CertPackUUID      string `json:"cert_pack_uuid"`
	ValidationMethod  string `json:"validation_method"`
	ValidationType    string `json:"validation_type"`
	VerificationInfo  struct {
		HTTPURL  string `json:"http_url"`
		HTTPBody string `json:"http_body"`
	} `json:"verification_info"`
	Hostname string `json:"hostname"`
}

type zoneSSLVerificationResponse struct {
	Response
	Result []ZoneSSLVerification `json:"result"`
}

// ListSSLVerification get SSL Verification Info for a Zone
//
// API reference: https://api.cloudflare.com/#ssl-verification-ssl-verification-details
func (api *API) ListSSLVerification(zoneID string) ([]ZoneSSLVerification, error) {
	uri := "/zones/" + zoneID + "/ssl/verification"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r zoneSSLVerificationResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
