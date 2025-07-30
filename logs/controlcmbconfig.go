// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ControlCmbConfigService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewControlCmbConfigService] method instead.
type ControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewControlCmbConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewControlCmbConfigService(opts ...option.RequestOption) (r *ControlCmbConfigService) {
	r = &ControlCmbConfigService{}
	r.Options = opts
	return
}
