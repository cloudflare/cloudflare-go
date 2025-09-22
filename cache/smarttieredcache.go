// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// SmartTieredCacheService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmartTieredCacheService] method instead.
type SmartTieredCacheService struct {
	Options []option.RequestOption
}

// NewSmartTieredCacheService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSmartTieredCacheService(opts ...option.RequestOption) (r *SmartTieredCacheService) {
	r = &SmartTieredCacheService{}
	r.Options = opts
	return
}
