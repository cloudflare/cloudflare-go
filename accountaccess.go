// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessService] method
// instead.
type AccountAccessService struct {
	Options           []option.RequestOption
	Bookmarks         *AccountAccessBookmarkService
	Certificates      *AccountAccessCertificateService
	Groups            *AccountAccessGroupService
	IdentityProviders *AccountAccessIdentityProviderService
	Keys              *AccountAccessKeyService
	Logs              *AccountAccessLogService
	Organizations     *AccountAccessOrganizationService
	Seats             *AccountAccessSeatService
	ServiceTokens     *AccountAccessServiceTokenService
	Users             *AccountAccessUserService
}

// NewAccountAccessService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAccessService(opts ...option.RequestOption) (r *AccountAccessService) {
	r = &AccountAccessService{}
	r.Options = opts
	r.Bookmarks = NewAccountAccessBookmarkService(opts...)
	r.Certificates = NewAccountAccessCertificateService(opts...)
	r.Groups = NewAccountAccessGroupService(opts...)
	r.IdentityProviders = NewAccountAccessIdentityProviderService(opts...)
	r.Keys = NewAccountAccessKeyService(opts...)
	r.Logs = NewAccountAccessLogService(opts...)
	r.Organizations = NewAccountAccessOrganizationService(opts...)
	r.Seats = NewAccountAccessSeatService(opts...)
	r.ServiceTokens = NewAccountAccessServiceTokenService(opts...)
	r.Users = NewAccountAccessUserService(opts...)
	return
}
