// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InstanceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceService] method instead.
type InstanceService struct {
	Options []option.RequestOption
	Status  *InstanceStatusService
	Events  *InstanceEventService
}

// NewInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceService(opts ...option.RequestOption) (r *InstanceService) {
	r = &InstanceService{}
	r.Options = opts
	r.Status = NewInstanceStatusService(opts...)
	r.Events = NewInstanceEventService(opts...)
	return
}
