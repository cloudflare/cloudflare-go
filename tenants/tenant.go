// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// TenantService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options      []option.RequestOption
	AccountTypes *AccountTypeService
	Accounts     *AccountService
	Entitlements *EntitlementService
	Memberships  *MembershipService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r *TenantService) {
	r = &TenantService{}
	r.Options = opts
	r.AccountTypes = NewAccountTypeService(opts...)
	r.Accounts = NewAccountService(opts...)
	r.Entitlements = NewEntitlementService(opts...)
	r.Memberships = NewMembershipService(opts...)
	return
}
