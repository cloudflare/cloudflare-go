// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ServiceEnvironmentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewServiceEnvironmentService] method
// instead.
type ServiceEnvironmentService struct {
	Options  []option.RequestOption
	Content  *ServiceEnvironmentContentService
	Settings *ServiceEnvironmentSettingService
}

// NewServiceEnvironmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewServiceEnvironmentService(opts ...option.RequestOption) (r *ServiceEnvironmentService) {
	r = &ServiceEnvironmentService{}
	r.Options = opts
	r.Content = NewServiceEnvironmentContentService(opts...)
	r.Settings = NewServiceEnvironmentSettingService(opts...)
	return
}
