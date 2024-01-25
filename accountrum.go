// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRumService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountRumService] method instead.
type AccountRumService struct {
	Options  []option.RequestOption
	SiteInfo *AccountRumSiteInfoService
	V2       *AccountRumV2Service
}

// NewAccountRumService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRumService(opts ...option.RequestOption) (r *AccountRumService) {
	r = &AccountRumService{}
	r.Options = opts
	r.SiteInfo = NewAccountRumSiteInfoService(opts...)
	r.V2 = NewAccountRumV2Service(opts...)
	return
}
