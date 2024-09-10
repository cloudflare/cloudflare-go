// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/shared"
)

// CacheService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCacheService] method instead.
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
// ```
// {"purge_everything": true}
// ```
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
// **NB:** For Zones on Free/Pro/Business plan, you may purge up to 30 URLs in one
// API call. For Zones on Enterprise plan, you may purge up to 500 URLs in one API
// call.
//
// Single file purge example with files:
//
// ```
// {"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}
// ```
//
// Single file purge example with url and header pairs:
//
// ```
// {"files": [{url: "http://www.example.com/cat_picture.jpg", headers: { "CF-IPCountry": "US", "CF-Device-Type": "desktop", "Accept-Language": "zh-CN" }}, {url: "http://www.example.com/dog_picture.jpg", headers: { "CF-IPCountry": "EU", "CF-Device-Type": "mobile", "Accept-Language": "en-US" }}]}
// ```
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
//
// Flex purge with tags:
//
// ```
// {"tags": ["a-cache-tag", "another-cache-tag"]}
// ```
//
// Flex purge with hosts:
//
// ```
// {"hosts": ["www.example.com", "images.example.com"]}
// ```
//
// Flex purge with prefixes:
//
// ```
// {"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}
// ```
func (r *CacheService) Purge(ctx context.Context, params CachePurgeParams, opts ...option.RequestOption) (res *CachePurgeResponse, err error) {
	var env CachePurgeResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
	Tags     param.Field[interface{}] `json:"tags,required"`
	Hosts    param.Field[interface{}] `json:"hosts,required"`
	Prefixes param.Field[interface{}] `json:"prefixes,required"`
	// For more information, please refer to
	// [purge everything documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Files           param.Field[interface{}] `json:"files,required"`
}

func (r CachePurgeParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBody) implementsCacheCachePurgeParamsBodyUnion() {}

// Satisfied by [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByTags],
// [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames],
// [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes],
// [cache.CachePurgeParamsBodyCachePurgeEverything],
// [cache.CachePurgeParamsBodyCachePurgeSingleFile],
// [cache.CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders],
// [CachePurgeParamsBody].
type CachePurgeParamsBodyUnion interface {
	implementsCacheCachePurgeParamsBodyUnion()
}

type CachePurgeParamsBodyCachePurgeFlexPurgeByTags struct {
	// For more information on cache tags and purging by tags, please refer to
	// [purge by cache-tags documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/#purge-cache-by-cache-tags-enterprise-only).
	Tags param.Field[[]string] `json:"tags"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByTags) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames struct {
	// For more information purging by hostnames, please refer to
	// [purge by hostname documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames) implementsCacheCachePurgeParamsBodyUnion() {
}

type CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes struct {
	// For more information on purging by prefixes, please refer to
	// [purge by prefix documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes) implementsCacheCachePurgeParamsBodyUnion() {
}

type CachePurgeParamsBodyCachePurgeEverything struct {
	// For more information, please refer to
	// [purge everything documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r CachePurgeParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeEverything) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeSingleFile struct {
	// For more information on purging files, please refer to
	// [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]string] `json:"files"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeSingleFile) implementsCacheCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders struct {
	// For more information on purging files with URL and headers, please refer to
	// [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile] `json:"files"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders) implementsCacheCachePurgeParamsBodyUnion() {
}

type CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile struct {
	Headers param.Field[interface{}] `json:"headers"`
	URL     param.Field[string]      `json:"url"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
