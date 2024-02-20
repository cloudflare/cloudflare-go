// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// CacheService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCacheService] method instead.
type CacheService struct {
	Options                  []option.RequestOption
	CacheReserves            *CacheCacheReserveService
	TieredCacheSmartTopology *CacheTieredCacheSmartTopologyService
	Variants                 *CacheVariantService
	RegionalTieredCache      *CacheRegionalTieredCacheService
}

// NewCacheService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCacheService(opts ...option.RequestOption) (r *CacheService) {
	r = &CacheService{}
	r.Options = opts
	r.CacheReserves = NewCacheCacheReserveService(opts...)
	r.TieredCacheSmartTopology = NewCacheTieredCacheSmartTopologyService(opts...)
	r.Variants = NewCacheVariantService(opts...)
	r.RegionalTieredCache = NewCacheRegionalTieredCacheService(opts...)
	return
}
