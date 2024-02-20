// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessService] method instead.
type AccessService struct {
	Options           []option.RequestOption
	Applications      *AccessApplicationService
	Certificates      *AccessCertificateService
	Groups            *AccessGroupService
	IdentityProviders *AccessIdentityProviderService
	Organizations     *AccessOrganizationService
	ServiceTokens     *AccessServiceTokenService
	Bookmarks         *AccessBookmarkService
	Keys              *AccessKeyService
	Logs              *AccessLogService
	Seats             *AccessSeatService
	Users             *AccessUserService
	CustomPages       *AccessCustomPageService
	Tags              *AccessTagService
}

// NewAccessService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccessService(opts ...option.RequestOption) (r *AccessService) {
	r = &AccessService{}
	r.Options = opts
	r.Applications = NewAccessApplicationService(opts...)
	r.Certificates = NewAccessCertificateService(opts...)
	r.Groups = NewAccessGroupService(opts...)
	r.IdentityProviders = NewAccessIdentityProviderService(opts...)
	r.Organizations = NewAccessOrganizationService(opts...)
	r.ServiceTokens = NewAccessServiceTokenService(opts...)
	r.Bookmarks = NewAccessBookmarkService(opts...)
	r.Keys = NewAccessKeyService(opts...)
	r.Logs = NewAccessLogService(opts...)
	r.Seats = NewAccessSeatService(opts...)
	r.Users = NewAccessUserService(opts...)
	r.CustomPages = NewAccessCustomPageService(opts...)
	r.Tags = NewAccessTagService(opts...)
	return
}
