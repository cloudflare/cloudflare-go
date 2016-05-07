package cloudflare

import (
	"github.com/pkg/errors"
	"encoding/json"
)

// CreateSSL allows you to add a custom SSL certificate to the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-create-ssl-configuration
// 	POST /zones/:zone_identifier/custom_certificates
func (c *API) CreateSSL() {
}

// ListSSL lists the custom certificates for the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-list-ssl-configurations
// 	GET /zones/:zone_identifier/custom_certificates
func (c *API) ListSSL() {
}

// SSLDetails returns the configuration details for a custom SSL certificate.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-ssl-configuration-details
// 	GET /zones/:zone_identifier/custom_certificates/:identifier
func (c *API) SSLDetails() {
}

// UpdateSSL updates (replaces) a custom SSL certificate.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-update-ssl-configuration
// 	PATCH /zones/:zone_identifier/custom_certificates/:identifier
func (c *API) UpdateSSL() {
}

// ReprioSSL allows you to change the priority (which is served for a given
// request) of custom SSL certificates associated with the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-re-prioritize-ssl-certificates
// 	PUT /zones/:zone_identifier/custom_certificates/prioritize
func (c *API) ReprioSSL() {
}

// DeleteSSL deletes a custom SSL certificate from the given zone.
// API reference:
// 	https://api.cloudflare.com/#custom-ssl-for-a-zone-delete-an-ssl-certificate
// 	DELETE /zones/:zone_identifier/custom_certificates/:identifier
func (c *API) DeleteSSL() {
}

// SSLVerification gets the Domain Control Validation records customers
// need to validate ownership of a domain for Universal SSL
//  GET /zones/:zone_identifier/ssl/verification"
func (c *API) SSLVerification(zoneID string) (SSLVerificationResponse, error) {
	uri := "/zones/" + zoneID + "/ssl/verification"
	res, err := c.makeRequest("GET", uri, nil)
	if err != nil {
		return SSLVerificationResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r SSLVerificationResponse
	if err = json.Unmarshal(res, &r); err != nil {
		return SSLVerificationResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}
