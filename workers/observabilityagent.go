// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ObservabilityAgentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityAgentService] method instead.
type ObservabilityAgentService struct {
	Options []option.RequestOption
	Traces  *ObservabilityAgentTraceService
}

// NewObservabilityAgentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservabilityAgentService(opts ...option.RequestOption) (r *ObservabilityAgentService) {
	r = &ObservabilityAgentService{}
	r.Options = opts
	r.Traces = NewObservabilityAgentTraceService(opts...)
	return
}
