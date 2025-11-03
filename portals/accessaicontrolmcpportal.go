// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package portals

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// AccessAIControlMcpPortalService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessAIControlMcpPortalService] method instead.
type AccessAIControlMcpPortalService struct {
	Options []option.RequestOption
}

// NewAccessAIControlMcpPortalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessAIControlMcpPortalService(opts ...option.RequestOption) (r *AccessAIControlMcpPortalService) {
	r = &AccessAIControlMcpPortalService{}
	r.Options = opts
	return
}
