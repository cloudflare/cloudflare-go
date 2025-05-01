// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// GatewayConfigurationCustomCertificateService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayConfigurationCustomCertificateService] method instead.
type GatewayConfigurationCustomCertificateService struct {
	Options []option.RequestOption
}

// NewGatewayConfigurationCustomCertificateService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewGatewayConfigurationCustomCertificateService(opts ...option.RequestOption) (r *GatewayConfigurationCustomCertificateService) {
	r = &GatewayConfigurationCustomCertificateService{}
	r.Options = opts
	return
}
