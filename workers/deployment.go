// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DeploymentService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDeploymentService] method instead.
type DeploymentService struct {
	Options   []option.RequestOption
	ByScripts *DeploymentByScriptService
}

// NewDeploymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeploymentService(opts ...option.RequestOption) (r *DeploymentService) {
	r = &DeploymentService{}
	r.Options = opts
	r.ByScripts = NewDeploymentByScriptService(opts...)
	return
}
