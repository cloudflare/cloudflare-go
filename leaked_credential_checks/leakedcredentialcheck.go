// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_checks

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// LeakedCredentialCheckService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeakedCredentialCheckService] method instead.
type LeakedCredentialCheckService struct {
	Options    []option.RequestOption
	Detections *DetectionService
}

// NewLeakedCredentialCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLeakedCredentialCheckService(opts ...option.RequestOption) (r *LeakedCredentialCheckService) {
	r = &LeakedCredentialCheckService{}
	r.Options = opts
	r.Detections = NewDetectionService(opts...)
	return
}
