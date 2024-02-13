// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// CachService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCachService] method instead.
type CachService struct {
	Options                         []option.RequestOption
	CacheReserves                   *CachCacheReserveService
	TieredCacheSmartTopologyEnables *CachTieredCacheSmartTopologyEnableService
	Variants                        *CachVariantService
}

// NewCachService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCachService(opts ...option.RequestOption) (r *CachService) {
	r = &CachService{}
	r.Options = opts
	r.CacheReserves = NewCachCacheReserveService(opts...)
	r.TieredCacheSmartTopologyEnables = NewCachTieredCacheSmartTopologyEnableService(opts...)
	r.Variants = NewCachVariantService(opts...)
	return
}
