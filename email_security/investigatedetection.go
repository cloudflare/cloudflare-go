// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InvestigateDetectionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateDetectionService] method instead.
type InvestigateDetectionService struct {
	Options []option.RequestOption
}

// NewInvestigateDetectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateDetectionService(opts ...option.RequestOption) (r *InvestigateDetectionService) {
	r = &InvestigateDetectionService{}
	r.Options = opts
	return
}
