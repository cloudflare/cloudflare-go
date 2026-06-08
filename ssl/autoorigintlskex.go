// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AutoOriginTLSKexService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutoOriginTLSKexService] method instead.
type AutoOriginTLSKexService struct {
	Options []option.RequestOption
}

// NewAutoOriginTLSKexService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAutoOriginTLSKexService(opts ...option.RequestOption) (r *AutoOriginTLSKexService) {
	r = &AutoOriginTLSKexService{}
	r.Options = opts
	return
}
