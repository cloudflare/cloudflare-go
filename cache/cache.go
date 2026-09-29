// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
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
	OriginCloudRegions  *OriginCloudRegionService
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
	r.OriginCloudRegions = NewOriginCloudRegionService(opts...)
	return
}

// Marks cached content as stale in every Cloudflare data center and cache tier,
// including Cache Reserve. The content stays in cache. The next request for it
// makes Cloudflare revalidate it with your origin, using the `ETag` and
// `Last-Modified` values it was cached with:
//
//   - If your origin answers `304 Not Modified`, Cloudflare serves the cached copy
//     without downloading it again, and `CF-Cache-Status` is `REVALIDATED`.
//   - If your origin sends a full response, Cloudflare serves and caches the new
//     content, and `CF-Cache-Status` is `EXPIRED`.
//
// With Tiered Cache, each tier revalidates with the tier above it, so a visitor
// can see `EXPIRED` even when your origin answered `304`.
//
// Until content is revalidated, your `stale-while-revalidate` and `stale-if-error`
// directives still apply, counted from the time you invalidated it. For example,
// if your origin fails during revalidation, Cloudflare can keep serving the stale
// copy for the `stale-if-error` window.
//
// ### Invalidate or purge?
//
//   - **Invalidate** when content may not have changed, for example after a deploy.
//     Unchanged content costs your origin a `304` instead of a full response. That
//     saving needs an origin that sends `ETag` or `Last-Modified` and answers
//     conditional requests. Otherwise, every revalidation downloads the full
//     response.
//   - **Purge**, with `POST /zones/{zone_id}/purge_cache`, when content must not be
//     served again, for example content you removed for legal or security reasons.
//
// Invalidating takes the same request bodies as purging, needs the same
// permission, and counts against the same rate limits. After a broad invalidation,
// such as `purge_everything`, expect more conditional requests to your origin
// while visitors request the invalidated content again.
//
// ### Choose what to invalidate
//
// Send one of these fields in the request body:
//
//   - `files`: specific URLs. If your cache key includes request headers, send each
//     URL with the header values it was cached with.
//   - `tags`: all content whose `Cache-Tag` response header contains one of the
//     tags.
//   - `hosts`: all content cached for the hostnames.
//   - `prefixes`: all content whose URL starts with one of the prefixes.
//   - `purge_everything`: all cached content in the zone.
//
// ### Check the result
//
// A `200` response with `success: true` means Cloudflare accepted the request. To
// check, request an invalidated URL and confirm that the `CF-Cache-Status`
// response header is `REVALIDATED` or `EXPIRED`.
//
// ### Availability and limits
//
// Rate limits and the number of items you can send in one request depend on your
// plan. See
// [Purge cache: availability and limits](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
func (r *CacheService) Invalidate(ctx context.Context, params CacheInvalidateParams, opts ...option.RequestOption) (res *CacheInvalidateResponse, err error) {
	var env CacheInvalidateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/invalidate_cache", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Marks cached content as stale for one environment of the zone. Content cached
// for the zone's other environments, including production, is not affected.
// Otherwise this works like `POST /zones/{zone_id}/invalidate_cache`: the next
// request for invalidated content makes Cloudflare revalidate it with your origin,
// and the request body takes the same fields.
//
// Environments are part of
// [Version Management](https://developers.cloudflare.com/version-management/). To
// delete the content instead, use
// `POST /zones/{zone_id}/environments/{environment_id}/purge_cache`.
//
// Invalidating by URL (`files`) does not work for environments that select
// requests by IP address, country, ASN, or threat score, and fails with error
// `1136`. Use `tags`, `hosts`, `prefixes`, or `purge_everything` for those
// environments.
//
// ### Availability and limits
//
// Rate limits and the number of items you can send in one request depend on your
// plan. See
// [Purge cache: availability and limits](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
func (r *CacheService) InvalidateEnvironment(ctx context.Context, environmentID string, params CacheInvalidateEnvironmentParams, opts ...option.RequestOption) (res *CacheInvalidateEnvironmentResponse, err error) {
	var env CacheInvalidateEnvironmentResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if environmentID == "" {
		err = errors.New("missing required environment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments/%s/invalidate_cache", params.ZoneID, environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes cached content in every Cloudflare data center and cache tier, including
// Cache Reserve. The next request for purged content is a cache `MISS`: Cloudflare
// fetches the full response from your origin and caches it again. Cloudflare does
// not serve purged content from cache again, even if your origin is unavailable.
//
// To keep content cached and have Cloudflare revalidate it with your origin
// instead, use `POST /zones/{zone_id}/invalidate_cache`.
//
// ### Choose what to purge
//
// Send one of these fields in the request body:
//
//   - `files`: specific URLs. If your cache key includes request headers, send each
//     URL with the header values it was cached with.
//   - `tags`: all content whose `Cache-Tag` response header contains one of the
//     tags.
//   - `hosts`: all content cached for the hostnames.
//   - `prefixes`: all content whose URL starts with one of the prefixes.
//   - `purge_everything`: all cached content in the zone.
//
// ### Check the result
//
// A `200` response with `success: true` means Cloudflare accepted the request. It
// does not confirm that any content was cached or removed. To check, request a
// purged URL and confirm that the `CF-Cache-Status` response header is `MISS`.
//
// ### Availability and limits
//
// Rate limits and the number of items you can send in one request depend on your
// plan. See
// [Purge cache: availability and limits](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
func (r *CacheService) Purge(ctx context.Context, params CachePurgeParams, opts ...option.RequestOption) (res *CachePurgeResponse, err error) {
	var env CachePurgeResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/purge_cache", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes cached content for one environment of the zone. Content cached for the
// zone's other environments, including production, is not affected. Otherwise this
// works like `POST /zones/{zone_id}/purge_cache`: the next request for purged
// content is a cache `MISS`, and the request body takes the same fields.
//
// Environments are part of
// [Version Management](https://developers.cloudflare.com/version-management/). To
// keep content cached and have Cloudflare revalidate it instead, use
// `POST /zones/{zone_id}/environments/{environment_id}/invalidate_cache`.
//
// Purging by URL (`files`) does not work for environments that select requests by
// IP address, country, ASN, or threat score, and fails with error `1136`. Use
// `tags`, `hosts`, `prefixes`, or `purge_everything` for those environments.
//
// ### Availability and limits
//
// Rate limits and the number of items you can send in one request depend on your
// plan. See
// [Purge cache: availability and limits](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
func (r *CacheService) PurgeEnvironment(ctx context.Context, environmentID string, params CachePurgeEnvironmentParams, opts ...option.RequestOption) (res *CachePurgeEnvironmentResponse, err error) {
	var env CachePurgeEnvironmentResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if environmentID == "" {
		err = errors.New("missing required environment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/environments/%s/purge_cache", params.ZoneID, environmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CacheInvalidateResponse struct {
	ID   string                      `json:"id" api:"required"`
	JSON cacheInvalidateResponseJSON `json:"-"`
}

// cacheInvalidateResponseJSON contains the JSON metadata for the struct
// [CacheInvalidateResponse]
type cacheInvalidateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheInvalidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheInvalidateResponseJSON) RawJSON() string {
	return r.raw
}

type CacheInvalidateEnvironmentResponse struct {
	ID   string                                 `json:"id" api:"required"`
	JSON cacheInvalidateEnvironmentResponseJSON `json:"-"`
}

// cacheInvalidateEnvironmentResponseJSON contains the JSON metadata for the struct
// [CacheInvalidateEnvironmentResponse]
type cacheInvalidateEnvironmentResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheInvalidateEnvironmentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheInvalidateEnvironmentResponseJSON) RawJSON() string {
	return r.raw
}

type CachePurgeResponse struct {
	ID   string                 `json:"id" api:"required"`
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

type CachePurgeEnvironmentResponse struct {
	ID   string                            `json:"id" api:"required"`
	JSON cachePurgeEnvironmentResponseJSON `json:"-"`
}

// cachePurgeEnvironmentResponseJSON contains the JSON metadata for the struct
// [CachePurgeEnvironmentResponse]
type cachePurgeEnvironmentResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeEnvironmentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeEnvironmentResponseJSON) RawJSON() string {
	return r.raw
}

type CacheInvalidateParams struct {
	ZoneID param.Field[string]            `path:"zone_id" api:"required"`
	Body   CacheInvalidateParamsBodyUnion `json:"body" api:"required"`
}

func (r CacheInvalidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CacheInvalidateParamsBody struct {
	Files    param.Field[interface{}] `json:"files"`
	Hosts    param.Field[interface{}] `json:"hosts"`
	Prefixes param.Field[interface{}] `json:"prefixes"`
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Tags            param.Field[interface{}] `json:"tags"`
}

func (r CacheInvalidateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateParamsBody) implementsCacheInvalidateParamsBodyUnion() {}

// Satisfied by [cache.CacheInvalidateParamsBodyCachePurgeFlexPurgeByTags],
// [cache.CacheInvalidateParamsBodyCachePurgeFlexPurgeByHostnames],
// [cache.CacheInvalidateParamsBodyCachePurgeFlexPurgeByPrefixes],
// [cache.CacheInvalidateParamsBodyCachePurgeEverything],
// [cache.CacheInvalidateParamsBodyCachePurgeSingleFile],
// [cache.CacheInvalidateParamsBodyCachePurgeSingleFileWithURLAndHeaders],
// [CacheInvalidateParamsBody].
type CacheInvalidateParamsBodyUnion interface {
	implementsCacheInvalidateParamsBodyUnion()
}

type CacheInvalidateParamsBodyCachePurgeFlexPurgeByTags struct {
	// Cache tags. Targets all content whose `Cache-Tag` response header contains at
	// least one of these tags. See
	// [Purge cache by cache-tags](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/).
	Tags param.Field[[]string] `json:"tags"`
}

func (r CacheInvalidateParamsBodyCachePurgeFlexPurgeByTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateParamsBodyCachePurgeFlexPurgeByTags) implementsCacheInvalidateParamsBodyUnion() {
}

type CacheInvalidateParamsBodyCachePurgeFlexPurgeByHostnames struct {
	// Hostnames, such as `www.example.com`. Targets all content cached for these
	// hostnames. See
	// [Purge cache by hostname](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r CacheInvalidateParamsBodyCachePurgeFlexPurgeByHostnames) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateParamsBodyCachePurgeFlexPurgeByHostnames) implementsCacheInvalidateParamsBodyUnion() {
}

type CacheInvalidateParamsBodyCachePurgeFlexPurgeByPrefixes struct {
	// URL prefixes, each a hostname followed by a path, such as
	// `www.example.com/blog/`. Targets all content whose URL starts with one of these
	// prefixes. Do not include a scheme, query string, or fragment. See
	// [Purge cache by prefix](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r CacheInvalidateParamsBodyCachePurgeFlexPurgeByPrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateParamsBodyCachePurgeFlexPurgeByPrefixes) implementsCacheInvalidateParamsBodyUnion() {
}

type CacheInvalidateParamsBodyCachePurgeEverything struct {
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r CacheInvalidateParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateParamsBodyCachePurgeEverything) implementsCacheInvalidateParamsBodyUnion() {}

type CacheInvalidateParamsBodyCachePurgeSingleFile struct {
	// Full URLs, such as `https://www.example.com/css/styles.css`. Targets the content
	// cached for each URL. If your cache key includes request headers, send objects
	// with `url` and `headers` instead. See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]string] `json:"files"`
}

func (r CacheInvalidateParamsBodyCachePurgeSingleFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateParamsBodyCachePurgeSingleFile) implementsCacheInvalidateParamsBodyUnion() {}

type CacheInvalidateParamsBodyCachePurgeSingleFileWithURLAndHeaders struct {
	// URLs with the request headers your cache key uses. Use this form when your cache
	// key includes request headers, or the visitor's device type, country, or
	// language: send the header values each URL was cached with, such as
	// `CF-Device-Type`, `CF-IPCountry`, or `Accept-Language`.
	//
	// When you send the `Origin` header, include the scheme and hostname. Include the
	// port unless it is the default for the scheme: 80 for `http`, 443 for `https`.
	//
	// See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]CacheInvalidateParamsBodyCachePurgeSingleFileWithURLAndHeadersFile] `json:"files"`
}

func (r CacheInvalidateParamsBodyCachePurgeSingleFileWithURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateParamsBodyCachePurgeSingleFileWithURLAndHeaders) implementsCacheInvalidateParamsBodyUnion() {
}

type CacheInvalidateParamsBodyCachePurgeSingleFileWithURLAndHeadersFile struct {
	// Request headers and the values the content was cached with.
	Headers param.Field[map[string]string] `json:"headers"`
	// Full URL of the content.
	URL param.Field[string] `json:"url"`
}

func (r CacheInvalidateParamsBodyCachePurgeSingleFileWithURLAndHeadersFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CacheInvalidateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Indicates the API call's success or failure.
	Success bool                                `json:"success" api:"required"`
	Result  CacheInvalidateResponse             `json:"result" api:"nullable"`
	JSON    cacheInvalidateResponseEnvelopeJSON `json:"-"`
}

// cacheInvalidateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheInvalidateResponseEnvelope]
type cacheInvalidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheInvalidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheInvalidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CacheInvalidateEnvironmentParams struct {
	ZoneID param.Field[string]                       `path:"zone_id" api:"required"`
	Body   CacheInvalidateEnvironmentParamsBodyUnion `json:"body" api:"required"`
}

func (r CacheInvalidateEnvironmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CacheInvalidateEnvironmentParamsBody struct {
	Files    param.Field[interface{}] `json:"files"`
	Hosts    param.Field[interface{}] `json:"hosts"`
	Prefixes param.Field[interface{}] `json:"prefixes"`
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Tags            param.Field[interface{}] `json:"tags"`
}

func (r CacheInvalidateEnvironmentParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateEnvironmentParamsBody) implementsCacheInvalidateEnvironmentParamsBodyUnion() {}

// Satisfied by
// [cache.CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByTags],
// [cache.CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames],
// [cache.CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes],
// [cache.CacheInvalidateEnvironmentParamsBodyCachePurgeEverything],
// [cache.CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFile],
// [cache.CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders],
// [CacheInvalidateEnvironmentParamsBody].
type CacheInvalidateEnvironmentParamsBodyUnion interface {
	implementsCacheInvalidateEnvironmentParamsBodyUnion()
}

type CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByTags struct {
	// Cache tags. Targets all content whose `Cache-Tag` response header contains at
	// least one of these tags. See
	// [Purge cache by cache-tags](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/).
	Tags param.Field[[]string] `json:"tags"`
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByTags) implementsCacheInvalidateEnvironmentParamsBodyUnion() {
}

type CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames struct {
	// Hostnames, such as `www.example.com`. Targets all content cached for these
	// hostnames. See
	// [Purge cache by hostname](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames) implementsCacheInvalidateEnvironmentParamsBodyUnion() {
}

type CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes struct {
	// URL prefixes, each a hostname followed by a path, such as
	// `www.example.com/blog/`. Targets all content whose URL starts with one of these
	// prefixes. Do not include a scheme, query string, or fragment. See
	// [Purge cache by prefix](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes) implementsCacheInvalidateEnvironmentParamsBodyUnion() {
}

type CacheInvalidateEnvironmentParamsBodyCachePurgeEverything struct {
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeEverything) implementsCacheInvalidateEnvironmentParamsBodyUnion() {
}

type CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFile struct {
	// Full URLs, such as `https://www.example.com/css/styles.css`. Targets the content
	// cached for each URL. If your cache key includes request headers, send objects
	// with `url` and `headers` instead. See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]string] `json:"files"`
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFile) implementsCacheInvalidateEnvironmentParamsBodyUnion() {
}

type CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders struct {
	// URLs with the request headers your cache key uses. Use this form when your cache
	// key includes request headers, or the visitor's device type, country, or
	// language: send the header values each URL was cached with, such as
	// `CF-Device-Type`, `CF-IPCountry`, or `Accept-Language`.
	//
	// When you send the `Origin` header, include the scheme and hostname. Include the
	// port unless it is the default for the scheme: 80 for `http`, 443 for `https`.
	//
	// See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeadersFile] `json:"files"`
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders) implementsCacheInvalidateEnvironmentParamsBodyUnion() {
}

type CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeadersFile struct {
	// Request headers and the values the content was cached with.
	Headers param.Field[map[string]string] `json:"headers"`
	// Full URL of the content.
	URL param.Field[string] `json:"url"`
}

func (r CacheInvalidateEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeadersFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CacheInvalidateEnvironmentResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Indicates the API call's success or failure.
	Success bool                                           `json:"success" api:"required"`
	Result  CacheInvalidateEnvironmentResponse             `json:"result" api:"nullable"`
	JSON    cacheInvalidateEnvironmentResponseEnvelopeJSON `json:"-"`
}

// cacheInvalidateEnvironmentResponseEnvelopeJSON contains the JSON metadata for
// the struct [CacheInvalidateEnvironmentResponseEnvelope]
type cacheInvalidateEnvironmentResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheInvalidateEnvironmentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheInvalidateEnvironmentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CachePurgeParams struct {
	ZoneID param.Field[string]       `path:"zone_id" api:"required"`
	Body   CachePurgeParamsBodyUnion `json:"body" api:"required"`
}

func (r CachePurgeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CachePurgeParamsBody struct {
	Files    param.Field[interface{}] `json:"files"`
	Hosts    param.Field[interface{}] `json:"hosts"`
	Prefixes param.Field[interface{}] `json:"prefixes"`
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Tags            param.Field[interface{}] `json:"tags"`
}

func (r CachePurgeParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBody) implementsCachePurgeParamsBodyUnion() {}

// Satisfied by [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByTags],
// [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames],
// [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes],
// [cache.CachePurgeParamsBodyCachePurgeEverything],
// [cache.CachePurgeParamsBodyCachePurgeSingleFile],
// [cache.CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders],
// [CachePurgeParamsBody].
type CachePurgeParamsBodyUnion interface {
	implementsCachePurgeParamsBodyUnion()
}

type CachePurgeParamsBodyCachePurgeFlexPurgeByTags struct {
	// Cache tags. Targets all content whose `Cache-Tag` response header contains at
	// least one of these tags. See
	// [Purge cache by cache-tags](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/).
	Tags param.Field[[]string] `json:"tags"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByTags) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames struct {
	// Hostnames, such as `www.example.com`. Targets all content cached for these
	// hostnames. See
	// [Purge cache by hostname](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes struct {
	// URL prefixes, each a hostname followed by a path, such as
	// `www.example.com/blog/`. Targets all content whose URL starts with one of these
	// prefixes. Do not include a scheme, query string, or fragment. See
	// [Purge cache by prefix](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeEverything struct {
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r CachePurgeParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeEverything) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeSingleFile struct {
	// Full URLs, such as `https://www.example.com/css/styles.css`. Targets the content
	// cached for each URL. If your cache key includes request headers, send objects
	// with `url` and `headers` instead. See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]string] `json:"files"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeSingleFile) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders struct {
	// URLs with the request headers your cache key uses. Use this form when your cache
	// key includes request headers, or the visitor's device type, country, or
	// language: send the header values each URL was cached with, such as
	// `CF-Device-Type`, `CF-IPCountry`, or `Accept-Language`.
	//
	// When you send the `Origin` header, include the scheme and hostname. Include the
	// port unless it is the default for the scheme: 80 for `http`, 443 for `https`.
	//
	// See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile] `json:"files"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders) implementsCachePurgeParamsBodyUnion() {
}

type CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile struct {
	// Request headers and the values the content was cached with.
	Headers param.Field[map[string]string] `json:"headers"`
	// Full URL of the content.
	URL param.Field[string] `json:"url"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CachePurgeResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Indicates the API call's success or failure.
	Success bool                           `json:"success" api:"required"`
	Result  CachePurgeResponse             `json:"result" api:"nullable"`
	JSON    cachePurgeResponseEnvelopeJSON `json:"-"`
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

type CachePurgeEnvironmentParams struct {
	ZoneID param.Field[string]                  `path:"zone_id" api:"required"`
	Body   CachePurgeEnvironmentParamsBodyUnion `json:"body" api:"required"`
}

func (r CachePurgeEnvironmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CachePurgeEnvironmentParamsBody struct {
	Files    param.Field[interface{}] `json:"files"`
	Hosts    param.Field[interface{}] `json:"hosts"`
	Prefixes param.Field[interface{}] `json:"prefixes"`
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Tags            param.Field[interface{}] `json:"tags"`
}

func (r CachePurgeEnvironmentParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeEnvironmentParamsBody) implementsCachePurgeEnvironmentParamsBodyUnion() {}

// Satisfied by [cache.CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByTags],
// [cache.CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames],
// [cache.CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes],
// [cache.CachePurgeEnvironmentParamsBodyCachePurgeEverything],
// [cache.CachePurgeEnvironmentParamsBodyCachePurgeSingleFile],
// [cache.CachePurgeEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders],
// [CachePurgeEnvironmentParamsBody].
type CachePurgeEnvironmentParamsBodyUnion interface {
	implementsCachePurgeEnvironmentParamsBodyUnion()
}

type CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByTags struct {
	// Cache tags. Targets all content whose `Cache-Tag` response header contains at
	// least one of these tags. See
	// [Purge cache by cache-tags](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/).
	Tags param.Field[[]string] `json:"tags"`
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByTags) implementsCachePurgeEnvironmentParamsBodyUnion() {
}

type CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames struct {
	// Hostnames, such as `www.example.com`. Targets all content cached for these
	// hostnames. See
	// [Purge cache by hostname](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByHostnames) implementsCachePurgeEnvironmentParamsBodyUnion() {
}

type CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes struct {
	// URL prefixes, each a hostname followed by a path, such as
	// `www.example.com/blog/`. Targets all content whose URL starts with one of these
	// prefixes. Do not include a scheme, query string, or fragment. See
	// [Purge cache by prefix](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeFlexPurgeByPrefixes) implementsCachePurgeEnvironmentParamsBodyUnion() {
}

type CachePurgeEnvironmentParamsBodyCachePurgeEverything struct {
	// Set to `true` to target all cached content in the zone, or in the environment
	// for the environment endpoints. Must be the only field in the request. See
	// [Purge everything](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeEverything) implementsCachePurgeEnvironmentParamsBodyUnion() {
}

type CachePurgeEnvironmentParamsBodyCachePurgeSingleFile struct {
	// Full URLs, such as `https://www.example.com/css/styles.css`. Targets the content
	// cached for each URL. If your cache key includes request headers, send objects
	// with `url` and `headers` instead. See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]string] `json:"files"`
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeSingleFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeSingleFile) implementsCachePurgeEnvironmentParamsBodyUnion() {
}

type CachePurgeEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders struct {
	// URLs with the request headers your cache key uses. Use this form when your cache
	// key includes request headers, or the visitor's device type, country, or
	// language: send the header values each URL was cached with, such as
	// `CF-Device-Type`, `CF-IPCountry`, or `Accept-Language`.
	//
	// When you send the `Origin` header, include the scheme and hostname. Include the
	// port unless it is the default for the scheme: 80 for `http`, 443 for `https`.
	//
	// See
	// [Purge by single-file](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]CachePurgeEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeadersFile] `json:"files"`
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeaders) implementsCachePurgeEnvironmentParamsBodyUnion() {
}

type CachePurgeEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeadersFile struct {
	// Request headers and the values the content was cached with.
	Headers param.Field[map[string]string] `json:"headers"`
	// Full URL of the content.
	URL param.Field[string] `json:"url"`
}

func (r CachePurgeEnvironmentParamsBodyCachePurgeSingleFileWithURLAndHeadersFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CachePurgeEnvironmentResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Indicates the API call's success or failure.
	Success bool                                      `json:"success" api:"required"`
	Result  CachePurgeEnvironmentResponse             `json:"result" api:"nullable"`
	JSON    cachePurgeEnvironmentResponseEnvelopeJSON `json:"-"`
}

// cachePurgeEnvironmentResponseEnvelopeJSON contains the JSON metadata for the
// struct [CachePurgeEnvironmentResponseEnvelope]
type cachePurgeEnvironmentResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeEnvironmentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeEnvironmentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
