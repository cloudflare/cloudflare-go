// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// PayloadService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayloadService] method instead.
type PayloadService struct {
	Options []option.RequestOption
}

// NewPayloadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPayloadService(opts ...option.RequestOption) (r *PayloadService) {
	r = &PayloadService{}
	r.Options = opts
	return
}
