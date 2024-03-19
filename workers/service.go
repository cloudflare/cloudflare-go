// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ServiceService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewServiceService] method instead.
type ServiceService struct {
	Options      []option.RequestOption
	Environments *ServiceEnvironmentService
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	r.Environments = NewServiceEnvironmentService(opts...)
	return
}
