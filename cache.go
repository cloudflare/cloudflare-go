// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// CacheService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCacheService] method instead.
type CacheService struct {
	Options             []option.RequestOption
	CacheReserve        *CacheCacheReserveService
	SmartTieredCached   *CacheSmartTieredCachedService
	Variants            *CacheVariantService
	RegionalTieredCache *CacheRegionalTieredCacheService
}

// NewCacheService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCacheService(opts ...option.RequestOption) (r *CacheService) {
	r = &CacheService{}
	r.Options = opts
	r.CacheReserve = NewCacheCacheReserveService(opts...)
	r.SmartTieredCached = NewCacheSmartTieredCachedService(opts...)
	r.Variants = NewCacheVariantService(opts...)
	r.RegionalTieredCache = NewCacheRegionalTieredCacheService(opts...)
	return
}

// ### Purge All Cached Content
//
// Removes ALL files from Cloudflare's cache. All tiers can purge everything.
//
// ### Purge Cached Content by URL
//
// Granularly removes one or more files from Cloudflare's cache by specifying URLs.
// All tiers can purge by URL.
//
// To purge files with custom cache keys, include the headers used to compute the
// cache key as in the example. If you have a device type or geo in your cache key,
// you will need to include the CF-Device-Type or CF-IPCountry headers. If you have
// lang in your cache key, you will need to include the Accept-Language header.
//
// **NB:** When including the Origin header, be sure to include the **scheme** and
// **hostname**. The port number can be omitted if it is the default port (80 for
// http, 443 for https), but must be included otherwise.
//
// ### Purge Cached Content by Tag, Host or Prefix
//
// Granularly removes one or more files from Cloudflare's cache either by
// specifying the host, the associated Cache-Tag, or a Prefix. Only Enterprise
// customers are permitted to purge by Tag, Host or Prefix.
//
// **NB:** Cache-Tag, host, and prefix purging each have a rate limit of 30,000
// purge API calls in every 24 hour period. You may purge up to 30 tags, hosts, or
// prefixes in one API call. This rate limit can be raised for customers who need
// to purge at higher volume.
func (r *CacheService) Purge(ctx context.Context, params CachePurgeParams, opts ...option.RequestOption) (res *CachePurgeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachePurgeResponseEnvelope
	path := fmt.Sprintf("zones/%s/purge_cache", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CachePurgeResponse struct {
	// Identifier
	ID   string                 `json:"id,required"`
	JSON cachePurgeResponseJSON `json:"-"`
}

// cachePurgeResponseJSON contains the JSON metadata for the struct
// [CachePurgeResponse]
type cachePurgeResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeResponseJSON) RawJSON() string {
	return r.raw
}

type CachePurgeParams struct {
	ZoneID          param.Field[string]                 `path:"zone_id,required"`
	Files           param.Field[[]CachePurgeParamsFile] `json:"files"`
	Hosts           param.Field[[]string]               `json:"hosts"`
	Prefixes        param.Field[[]string]               `json:"prefixes"`
	PurgeEverything param.Field[bool]                   `json:"purge_everything"`
	Tags            param.Field[[]string]               `json:"tags"`
}

func (r CachePurgeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [CachePurgeParamsFilesCachePurgeURLAndHeaders].
type CachePurgeParamsFile interface {
	ImplementsCachePurgeParamsFile()
}

type CachePurgeParamsFilesCachePurgeURLAndHeaders struct {
	Headers param.Field[interface{}] `json:"headers"`
	URL     param.Field[string]      `json:"url"`
}

func (r CachePurgeParamsFilesCachePurgeURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsFilesCachePurgeURLAndHeaders) ImplementsCachePurgeParamsFile() {}

type CachePurgeResponseEnvelope struct {
	Errors   []CachePurgeResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachePurgeResponseEnvelopeMessages `json:"messages,required"`
	Result   CachePurgeResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success CachePurgeResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachePurgeResponseEnvelopeJSON    `json:"-"`
}

// cachePurgeResponseEnvelopeJSON contains the JSON metadata for the struct
// [CachePurgeResponseEnvelope]
type cachePurgeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CachePurgeResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    cachePurgeResponseEnvelopeErrorsJSON `json:"-"`
}

// cachePurgeResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CachePurgeResponseEnvelopeErrors]
type cachePurgeResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CachePurgeResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    cachePurgeResponseEnvelopeMessagesJSON `json:"-"`
}

// cachePurgeResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CachePurgeResponseEnvelopeMessages]
type cachePurgeResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CachePurgeResponseEnvelopeSuccess bool

const (
	CachePurgeResponseEnvelopeSuccessTrue CachePurgeResponseEnvelopeSuccess = true
)
