// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package google_tag_gateway

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// GoogleTagGatewayService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoogleTagGatewayService] method instead.
type GoogleTagGatewayService struct {
	Options []option.RequestOption
	Config  *ConfigService
}

// NewGoogleTagGatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGoogleTagGatewayService(opts ...option.RequestOption) (r *GoogleTagGatewayService) {
	r = &GoogleTagGatewayService{}
	r.Options = opts
	r.Config = NewConfigService(opts...)
	return
}
