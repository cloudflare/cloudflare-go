// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// CTService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCTService] method instead.
type CTService struct {
	Options  []option.RequestOption
	Alerting *CTAlertingService
}

// NewCTService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCTService(opts ...option.RequestOption) (r *CTService) {
	r = &CTService{}
	r.Options = opts
	r.Alerting = NewCTAlertingService(opts...)
	return
}
