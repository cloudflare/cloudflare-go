// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InvestigateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateService] method instead.
type InvestigateService struct {
	Options    []option.RequestOption
	Detections *InvestigateDetectionService
	Preview    *InvestigatePreviewService
	Raw        *InvestigateRawService
	Trace      *InvestigateTraceService
	Move       *InvestigateMoveService
	Reclassify *InvestigateReclassifyService
	Release    *InvestigateReleaseService
}

// NewInvestigateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateService(opts ...option.RequestOption) (r *InvestigateService) {
	r = &InvestigateService{}
	r.Options = opts
	r.Detections = NewInvestigateDetectionService(opts...)
	r.Preview = NewInvestigatePreviewService(opts...)
	r.Raw = NewInvestigateRawService(opts...)
	r.Trace = NewInvestigateTraceService(opts...)
	r.Move = NewInvestigateMoveService(opts...)
	r.Reclassify = NewInvestigateReclassifyService(opts...)
	r.Release = NewInvestigateReleaseService(opts...)
	return
}
