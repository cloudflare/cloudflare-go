// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountIntelIndicatorFeedService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountIntelIndicatorFeedService] method instead.
type AccountIntelIndicatorFeedService struct {
	Options     []option.RequestOption
	Permissions *AccountIntelIndicatorFeedPermissionService
}

// NewAccountIntelIndicatorFeedService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountIntelIndicatorFeedService(opts ...option.RequestOption) (r *AccountIntelIndicatorFeedService) {
	r = &AccountIntelIndicatorFeedService{}
	r.Options = opts
	r.Permissions = NewAccountIntelIndicatorFeedPermissionService(opts...)
	return
}
