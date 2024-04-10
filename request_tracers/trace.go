// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package request_tracers

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TraceService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTraceService] method instead.
type TraceService struct {
	Options []option.RequestOption
}

// NewTraceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTraceService(opts ...option.RequestOption) (r *TraceService) {
	r = &TraceService{}
	r.Options = opts
	return
}
