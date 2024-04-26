// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// CacheService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCacheService] method instead.
type CacheService struct {
	Options             []option.RequestOption
	CacheReserve        *CacheReserveService
	SmartTieredCache    *SmartTieredCacheService
	Variants            *VariantService
	RegionalTieredCache *RegionalTieredCacheService
}

// NewCacheService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCacheService(opts ...option.RequestOption) (r *CacheService) {
	r = &CacheService{}
	r.Options = opts
	r.CacheReserve = NewCacheReserveService(opts...)
	r.SmartTieredCache = NewSmartTieredCacheService(opts...)
	r.Variants = NewVariantService(opts...)
	r.RegionalTieredCache = NewRegionalTieredCacheService(opts...)
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
// http, 443 for https), but must be included otherwise. **NB:** For Zones on
// Free/Pro/Business plan, you may purge up to 30 URLs in one API call. For Zones
// on Enterprise plan, you may purge up to 500 URLs in one API call.
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
	ZoneID param.Field[string]       `path:"zone_id,required"`
	Body   CachePurgeParamsBodyUnion `json:"body,required"`
}

func (r CachePurgeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CachePurgeParamsBody struct {
	Tags            param.Field[interface{}] `json:"tags,required"`
	Hosts           param.Field[interface{}] `json:"hosts,required"`
	Prefixes        param.Field[interface{}] `json:"prefixes,required"`
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Files           param.Field[interface{}] `json:"files,required"`
}

func (r CachePurgeParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBody) implementsCacheCachePurgeParamsBodyUnion() {}

// Satisfied by [cache.CachePurgeParamsBodyCachePurgeTags],
// [cache.CachePurgeParamsBodyCachePurgeHosts],
// [cache.CachePurgeParamsBodyCachePurgePrefixes],
// [cache.CachePurgeParamsBodyCachePurgeEverything],
// [cache.CachePurgeParamsBodyCachePurgeFiles], [CachePurgeParamsBody].
type CachePurgeParamsBodyUnion interface {
	implementsCacheCachePurgeParamsBodyUnion()
}

type CachePurgeParamsBodyCachePurgeTags struct {
	Tags param.Field[[]string] `json:"tags"`
}

func (r CachePurgeParamsBodyCachePurgeTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeTags) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeHosts struct {
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r CachePurgeParamsBodyCachePurgeHosts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeHosts) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgePrefixes struct {
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r CachePurgeParamsBodyCachePurgePrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgePrefixes) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeEverything struct {
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r CachePurgeParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeEverything) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeFiles struct {
	Files param.Field[[]CachePurgeParamsBodyCachePurgeFilesFileUnion] `json:"files"`
}

func (r CachePurgeParamsBodyCachePurgeFiles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFiles) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeFilesFile struct {
	Headers param.Field[interface{}] `json:"headers,required"`
	URL     param.Field[string]      `json:"url"`
}

func (r CachePurgeParamsBodyCachePurgeFilesFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFilesFile) ImplementsCacheCachePurgeParamsBodyCachePurgeFilesFileUnion() {
}

// Satisfied by [shared.UnionString],
// [cache.CachePurgeParamsBodyCachePurgeFilesFilesCachePurgeURLAndHeaders],
// [CachePurgeParamsBodyCachePurgeFilesFile].
type CachePurgeParamsBodyCachePurgeFilesFileUnion interface {
	ImplementsCacheCachePurgeParamsBodyCachePurgeFilesFileUnion()
}

type CachePurgeParamsBodyCachePurgeFilesFilesCachePurgeURLAndHeaders struct {
	Headers param.Field[interface{}] `json:"headers"`
	URL     param.Field[string]      `json:"url"`
}

func (r CachePurgeParamsBodyCachePurgeFilesFilesCachePurgeURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFilesFilesCachePurgeURLAndHeaders) ImplementsCacheCachePurgeParamsBodyCachePurgeFilesFileUnion() {
}

type CachePurgeResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success CachePurgeResponseEnvelopeSuccess `json:"success,required"`
	Result  CachePurgeResponse                `json:"result,nullable"`
	JSON    cachePurgeResponseEnvelopeJSON    `json:"-"`
}

// cachePurgeResponseEnvelopeJSON contains the JSON metadata for the struct
// [CachePurgeResponseEnvelope]
type cachePurgeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CachePurgeResponseEnvelopeSuccess bool

const (
	CachePurgeResponseEnvelopeSuccessTrue CachePurgeResponseEnvelopeSuccess = true
)

func (r CachePurgeResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CachePurgeResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
