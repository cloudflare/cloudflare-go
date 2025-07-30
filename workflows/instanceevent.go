// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// InstanceEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceEventService] method instead.
type InstanceEventService struct {
	Options []option.RequestOption
}

// NewInstanceEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceEventService(opts ...option.RequestOption) (r *InstanceEventService) {
	r = &InstanceEventService{}
	r.Options = opts
	return
}
