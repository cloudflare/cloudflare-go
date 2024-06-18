// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TURNService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTURNService] method instead.
type TURNService struct {
	Options []option.RequestOption
	Keys    *TURNKeyService
}

// NewTURNService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTURNService(opts ...option.RequestOption) (r *TURNService) {
	r = &TURNService{}
	r.Options = opts
	r.Keys = NewTURNKeyService(opts...)
	return
}
