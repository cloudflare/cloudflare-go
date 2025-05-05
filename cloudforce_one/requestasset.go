// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// RequestAssetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestAssetService] method instead.
type RequestAssetService struct {
	Options []option.RequestOption
}

// NewRequestAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestAssetService(opts ...option.RequestOption) (r *RequestAssetService) {
	r = &RequestAssetService{}
	r.Options = opts
	return
}
