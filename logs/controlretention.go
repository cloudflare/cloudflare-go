// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ControlRetentionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewControlRetentionService] method instead.
type ControlRetentionService struct {
	Options []option.RequestOption
}

// NewControlRetentionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewControlRetentionService(opts ...option.RequestOption) (r *ControlRetentionService) {
	r = &ControlRetentionService{}
	r.Options = opts
	return
}
