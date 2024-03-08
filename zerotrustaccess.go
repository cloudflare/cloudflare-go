// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustAccessService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustAccessService] method
// instead.
type ZeroTrustAccessService struct {
	Options       []option.RequestOption
	Applications  *ZeroTrustAccessApplicationService
	Certificates  *ZeroTrustAccessCertificateService
	Groups        *ZeroTrustAccessGroupService
	ServiceTokens *ZeroTrustAccessServiceTokenService
	Bookmarks     *ZeroTrustAccessBookmarkService
	Keys          *ZeroTrustAccessKeyService
	Logs          *ZeroTrustAccessLogService
	Users         *ZeroTrustAccessUserService
	CustomPages   *ZeroTrustAccessCustomPageService
	Tags          *ZeroTrustAccessTagService
}

// NewZeroTrustAccessService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZeroTrustAccessService(opts ...option.RequestOption) (r *ZeroTrustAccessService) {
	r = &ZeroTrustAccessService{}
	r.Options = opts
	r.Applications = NewZeroTrustAccessApplicationService(opts...)
	r.Certificates = NewZeroTrustAccessCertificateService(opts...)
	r.Groups = NewZeroTrustAccessGroupService(opts...)
	r.ServiceTokens = NewZeroTrustAccessServiceTokenService(opts...)
	r.Bookmarks = NewZeroTrustAccessBookmarkService(opts...)
	r.Keys = NewZeroTrustAccessKeyService(opts...)
	r.Logs = NewZeroTrustAccessLogService(opts...)
	r.Users = NewZeroTrustAccessUserService(opts...)
	r.CustomPages = NewZeroTrustAccessCustomPageService(opts...)
	r.Tags = NewZeroTrustAccessTagService(opts...)
	return
}
