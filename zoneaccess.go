// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAccessService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneAccessService] method instead.
type ZoneAccessService struct {
	Options           []option.RequestOption
	Certificates      *ZoneAccessCertificateService
	Groups            *ZoneAccessGroupService
	IdentityProviders *ZoneAccessIdentityProviderService
	Organizations     *ZoneAccessOrganizationService
	ServiceTokens     *ZoneAccessServiceTokenService
}

// NewZoneAccessService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAccessService(opts ...option.RequestOption) (r *ZoneAccessService) {
	r = &ZoneAccessService{}
	r.Options = opts
	r.Certificates = NewZoneAccessCertificateService(opts...)
	r.Groups = NewZoneAccessGroupService(opts...)
	r.IdentityProviders = NewZoneAccessIdentityProviderService(opts...)
	r.Organizations = NewZoneAccessOrganizationService(opts...)
	r.ServiceTokens = NewZoneAccessServiceTokenService(opts...)
	return
}
