// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// ThreatEventDatasetHealthService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventDatasetHealthService] method instead.
type ThreatEventDatasetHealthService struct {
	Options []option.RequestOption
}

// NewThreatEventDatasetHealthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventDatasetHealthService(opts ...option.RequestOption) (r *ThreatEventDatasetHealthService) {
	r = &ThreatEventDatasetHealthService{}
	r.Options = opts
	return
}
