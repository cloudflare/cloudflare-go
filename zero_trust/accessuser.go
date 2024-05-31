// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AccessUserService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserService] method instead.
type AccessUserService struct {
	Options          []option.RequestOption
	ActiveSessions   *AccessUserActiveSessionService
	LastSeenIdentity *AccessUserLastSeenIdentityService
	FailedLogins     *AccessUserFailedLoginService
}

// NewAccessUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessUserService(opts ...option.RequestOption) (r *AccessUserService) {
	r = &AccessUserService{}
	r.Options = opts
	r.ActiveSessions = NewAccessUserActiveSessionService(opts...)
	r.LastSeenIdentity = NewAccessUserLastSeenIdentityService(opts...)
	r.FailedLogins = NewAccessUserFailedLoginService(opts...)
	return
}
