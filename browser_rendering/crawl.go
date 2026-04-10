// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// CrawlService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCrawlService] method instead.
type CrawlService struct {
	Options []option.RequestOption
}

// NewCrawlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCrawlService(opts ...option.RequestOption) (r *CrawlService) {
	r = &CrawlService{}
	r.Options = opts
	return
}

// Starts a crawl job for the provided URL and its children. Check available
// options like `gotoOptions` and `waitFor*` to control page load behaviour.
func (r *CrawlService) New(ctx context.Context, params CrawlNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env CrawlNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/crawl", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Cancels an ongoing crawl job by setting its status to cancelled and stopping all
// queued URLs.
func (r *CrawlService) Delete(ctx context.Context, jobID string, body CrawlDeleteParams, opts ...option.RequestOption) (res *CrawlDeleteResponse, err error) {
	var env CrawlDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/crawl/%s", body.AccountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the result of a crawl job.
func (r *CrawlService) Get(ctx context.Context, jobID string, params CrawlGetParams, opts ...option.RequestOption) (res *CrawlGetResponse, err error) {
	var env CrawlGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/crawl/%s", params.AccountID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CrawlDeleteResponse struct {
	// The ID of the cancelled job.
	JobID string `json:"job_id" api:"required"`
	// Cancellation confirmation message.
	Message string                  `json:"message" api:"required"`
	JSON    crawlDeleteResponseJSON `json:"-"`
}

// crawlDeleteResponseJSON contains the JSON metadata for the struct
// [CrawlDeleteResponse]
type crawlDeleteResponseJSON struct {
	JobID       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CrawlGetResponse struct {
	// Crawl job ID.
	ID string `json:"id" api:"required"`
	// Total seconds spent in browser so far.
	BrowserSecondsUsed float64 `json:"browserSecondsUsed" api:"required"`
	// Total number of URLs that have been crawled so far.
	Finished float64 `json:"finished" api:"required"`
	// List of crawl job records.
	Records []CrawlGetResponseRecord `json:"records" api:"required"`
	// Total number of URLs that were skipped due to include/exclude/subdomain filters.
	// Skipped URLs are included in records but are not counted toward total/finished.
	Skipped float64 `json:"skipped" api:"required"`
	// Current crawl job status.
	Status string `json:"status" api:"required"`
	// Total current number of URLs in the crawl job.
	Total float64 `json:"total" api:"required"`
	// Cursor for pagination.
	Cursor string               `json:"cursor"`
	JSON   crawlGetResponseJSON `json:"-"`
}

// crawlGetResponseJSON contains the JSON metadata for the struct
// [CrawlGetResponse]
type crawlGetResponseJSON struct {
	ID                 apijson.Field
	BrowserSecondsUsed apijson.Field
	Finished           apijson.Field
	Records            apijson.Field
	Skipped            apijson.Field
	Status             apijson.Field
	Total              apijson.Field
	Cursor             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CrawlGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlGetResponseJSON) RawJSON() string {
	return r.raw
}

type CrawlGetResponseRecord struct {
	Metadata CrawlGetResponseRecordsMetadata `json:"metadata" api:"required"`
	// Current status of the crawled URL.
	Status CrawlGetResponseRecordsStatus `json:"status" api:"required"`
	// Crawled URL.
	URL string `json:"url" api:"required"`
	// HTML content of the crawled URL.
	HTML string `json:"html"`
	// JSON of the content of the crawled URL.
	Json map[string]interface{} `json:"json"`
	// Markdown of the content of the crawled URL.
	Markdown string                     `json:"markdown"`
	JSON     crawlGetResponseRecordJSON `json:"-"`
}

// crawlGetResponseRecordJSON contains the JSON metadata for the struct
// [CrawlGetResponseRecord]
type crawlGetResponseRecordJSON struct {
	Metadata    apijson.Field
	Status      apijson.Field
	URL         apijson.Field
	HTML        apijson.Field
	Json        apijson.Field
	Markdown    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlGetResponseRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlGetResponseRecordJSON) RawJSON() string {
	return r.raw
}

type CrawlGetResponseRecordsMetadata struct {
	// HTTP status code of the crawled page.
	Status float64 `json:"status" api:"required"`
	// Final URL of the crawled page.
	URL string `json:"url" api:"required"`
	// Title of the crawled page.
	Title string                              `json:"title"`
	JSON  crawlGetResponseRecordsMetadataJSON `json:"-"`
}

// crawlGetResponseRecordsMetadataJSON contains the JSON metadata for the struct
// [CrawlGetResponseRecordsMetadata]
type crawlGetResponseRecordsMetadataJSON struct {
	Status      apijson.Field
	URL         apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlGetResponseRecordsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlGetResponseRecordsMetadataJSON) RawJSON() string {
	return r.raw
}

// Current status of the crawled URL.
type CrawlGetResponseRecordsStatus string

const (
	CrawlGetResponseRecordsStatusQueued     CrawlGetResponseRecordsStatus = "queued"
	CrawlGetResponseRecordsStatusErrored    CrawlGetResponseRecordsStatus = "errored"
	CrawlGetResponseRecordsStatusCompleted  CrawlGetResponseRecordsStatus = "completed"
	CrawlGetResponseRecordsStatusDisallowed CrawlGetResponseRecordsStatus = "disallowed"
	CrawlGetResponseRecordsStatusSkipped    CrawlGetResponseRecordsStatus = "skipped"
	CrawlGetResponseRecordsStatusCancelled  CrawlGetResponseRecordsStatus = "cancelled"
)

func (r CrawlGetResponseRecordsStatus) IsKnown() bool {
	switch r {
	case CrawlGetResponseRecordsStatusQueued, CrawlGetResponseRecordsStatusErrored, CrawlGetResponseRecordsStatusCompleted, CrawlGetResponseRecordsStatusDisallowed, CrawlGetResponseRecordsStatusSkipped, CrawlGetResponseRecordsStatusCancelled:
		return true
	}
	return false
}

type CrawlNewParams struct {
	// Account ID.
	AccountID param.Field[string]     `path:"account_id" api:"required"`
	Body      CrawlNewParamsBodyUnion `json:"body" api:"required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
}

func (r CrawlNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [CrawlNewParams]'s query parameters as `url.Values`.
func (r CrawlNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CrawlNewParamsBody struct {
	// URL to navigate to, eg. `https://example.com`.
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout       param.Field[float64]     `json:"actionTimeout"`
	AddScriptTag        param.Field[interface{}] `json:"addScriptTag"`
	AddStyleTag         param.Field[interface{}] `json:"addStyleTag"`
	AllowRequestPattern param.Field[interface{}] `json:"allowRequestPattern"`
	AllowResourceTypes  param.Field[interface{}] `json:"allowResourceTypes"`
	Authenticate        param.Field[interface{}] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt   param.Field[bool]        `json:"bestAttempt"`
	Cookies       param.Field[interface{}] `json:"cookies"`
	CrawlPurposes param.Field[interface{}] `json:"crawlPurposes"`
	// Maximum number of levels deep the crawler will traverse from the starting URL.
	Depth            param.Field[float64]     `json:"depth"`
	EmulateMediaType param.Field[string]      `json:"emulateMediaType"`
	Formats          param.Field[interface{}] `json:"formats"`
	GotoOptions      param.Field[interface{}] `json:"gotoOptions"`
	JsonOptions      param.Field[interface{}] `json:"jsonOptions"`
	// Maximum number of URLs to crawl.
	Limit param.Field[float64] `json:"limit"`
	// Maximum age of a resource that can be returned from cache in seconds. Default is
	// 1 day.
	MaxAge param.Field[float64] `json:"maxAge"`
	// Unix timestamp (seconds since epoch) indicating to only crawl pages that were
	// modified since this time. For sitemap URLs with a lastmod field, this is
	// compared directly. For other URLs, the crawler will use If-Modified-Since header
	// when fetching. URLs without modification information (no lastmod in sitemap and
	// no Last-Modified header support) will be crawled. Note: This works in
	// conjunction with maxAge - both filters must pass for a cached resource to be
	// used. Must be within the last year and not in the future.
	ModifiedSince        param.Field[int64]       `json:"modifiedSince"`
	Options              param.Field[interface{}] `json:"options"`
	RejectRequestPattern param.Field[interface{}] `json:"rejectRequestPattern"`
	RejectResourceTypes  param.Field[interface{}] `json:"rejectResourceTypes"`
	// Whether to render the page or fetch static content. True by default.
	Render               param.Field[CrawlNewParamsBodyRender] `json:"render"`
	SetExtraHTTPHeaders  param.Field[interface{}]              `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                     `json:"setJavaScriptEnabled"`
	// Source of links to crawl. 'sitemaps' - only crawl URLs from sitemaps, 'links' -
	// only crawl URLs scraped from pages, 'all' - crawl both sitemap and scraped links
	// (default).
	Source          param.Field[CrawlNewParamsBodySource] `json:"source"`
	Viewport        param.Field[interface{}]              `json:"viewport"`
	WaitForSelector param.Field[interface{}]              `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r CrawlNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CrawlNewParamsBody) implementsCrawlNewParamsBodyUnion() {}

// Satisfied by [browser_rendering.CrawlNewParamsBodyObject],
// [browser_rendering.CrawlNewParamsBodyObject], [CrawlNewParamsBody].
type CrawlNewParamsBodyUnion interface {
	implementsCrawlNewParamsBodyUnion()
}

type CrawlNewParamsBodyObject struct {
	// URL to navigate to, eg. `https://example.com`.
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]CrawlNewParamsBodyObjectAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]CrawlNewParamsBodyObjectAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]CrawlNewParamsBodyObjectAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[CrawlNewParamsBodyObjectAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies param.Field[[]CrawlNewParamsBodyObjectCookie] `json:"cookies"`
	// List of crawl purposes to respect Content-Signal directives in robots.txt.
	// Allowed values: 'search', 'ai-input', 'ai-train'. Learn more:
	// https://contentsignals.org/. Default: ['search', 'ai-input', 'ai-train'].
	CrawlPurposes param.Field[[]CrawlNewParamsBodyObjectCrawlPurpose] `json:"crawlPurposes"`
	// Maximum number of levels deep the crawler will traverse from the starting URL.
	Depth            param.Field[float64] `json:"depth"`
	EmulateMediaType param.Field[string]  `json:"emulateMediaType"`
	// Formats to return. Default is `html`.
	Formats param.Field[[]CrawlNewParamsBodyObjectFormat] `json:"formats"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[CrawlNewParamsBodyObjectGotoOptions] `json:"gotoOptions"`
	// Options for JSON extraction.
	JsonOptions param.Field[CrawlNewParamsBodyObjectJsonOptions] `json:"jsonOptions"`
	// Maximum number of URLs to crawl.
	Limit param.Field[float64] `json:"limit"`
	// Maximum age of a resource that can be returned from cache in seconds. Default is
	// 1 day.
	MaxAge param.Field[float64] `json:"maxAge"`
	// Unix timestamp (seconds since epoch) indicating to only crawl pages that were
	// modified since this time. For sitemap URLs with a lastmod field, this is
	// compared directly. For other URLs, the crawler will use If-Modified-Since header
	// when fetching. URLs without modification information (no lastmod in sitemap and
	// no Last-Modified header support) will be crawled. Note: This works in
	// conjunction with maxAge - both filters must pass for a cached resource to be
	// used. Must be within the last year and not in the future.
	ModifiedSince param.Field[int64] `json:"modifiedSince"`
	// Additional options for the crawler.
	Options param.Field[CrawlNewParamsBodyObjectOptions] `json:"options"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes param.Field[[]CrawlNewParamsBodyObjectRejectResourceType] `json:"rejectResourceTypes"`
	// Whether to render the page or fetch static content. True by default.
	Render               param.Field[CrawlNewParamsBodyObjectRender] `json:"render"`
	SetExtraHTTPHeaders  param.Field[map[string]string]              `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                           `json:"setJavaScriptEnabled"`
	// Source of links to crawl. 'sitemaps' - only crawl URLs from sitemaps, 'links' -
	// only crawl URLs scraped from pages, 'all' - crawl both sitemap and scraped links
	// (default).
	Source param.Field[CrawlNewParamsBodyObjectSource] `json:"source"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[CrawlNewParamsBodyObjectViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[CrawlNewParamsBodyObjectWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r CrawlNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CrawlNewParamsBodyObject) implementsCrawlNewParamsBodyUnion() {}

type CrawlNewParamsBodyObjectAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r CrawlNewParamsBodyObjectAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r CrawlNewParamsBodyObjectAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectAllowResourceType string

const (
	CrawlNewParamsBodyObjectAllowResourceTypeDocument           CrawlNewParamsBodyObjectAllowResourceType = "document"
	CrawlNewParamsBodyObjectAllowResourceTypeStylesheet         CrawlNewParamsBodyObjectAllowResourceType = "stylesheet"
	CrawlNewParamsBodyObjectAllowResourceTypeImage              CrawlNewParamsBodyObjectAllowResourceType = "image"
	CrawlNewParamsBodyObjectAllowResourceTypeMedia              CrawlNewParamsBodyObjectAllowResourceType = "media"
	CrawlNewParamsBodyObjectAllowResourceTypeFont               CrawlNewParamsBodyObjectAllowResourceType = "font"
	CrawlNewParamsBodyObjectAllowResourceTypeScript             CrawlNewParamsBodyObjectAllowResourceType = "script"
	CrawlNewParamsBodyObjectAllowResourceTypeTexttrack          CrawlNewParamsBodyObjectAllowResourceType = "texttrack"
	CrawlNewParamsBodyObjectAllowResourceTypeXHR                CrawlNewParamsBodyObjectAllowResourceType = "xhr"
	CrawlNewParamsBodyObjectAllowResourceTypeFetch              CrawlNewParamsBodyObjectAllowResourceType = "fetch"
	CrawlNewParamsBodyObjectAllowResourceTypePrefetch           CrawlNewParamsBodyObjectAllowResourceType = "prefetch"
	CrawlNewParamsBodyObjectAllowResourceTypeEventsource        CrawlNewParamsBodyObjectAllowResourceType = "eventsource"
	CrawlNewParamsBodyObjectAllowResourceTypeWebsocket          CrawlNewParamsBodyObjectAllowResourceType = "websocket"
	CrawlNewParamsBodyObjectAllowResourceTypeManifest           CrawlNewParamsBodyObjectAllowResourceType = "manifest"
	CrawlNewParamsBodyObjectAllowResourceTypeSignedexchange     CrawlNewParamsBodyObjectAllowResourceType = "signedexchange"
	CrawlNewParamsBodyObjectAllowResourceTypePing               CrawlNewParamsBodyObjectAllowResourceType = "ping"
	CrawlNewParamsBodyObjectAllowResourceTypeCspviolationreport CrawlNewParamsBodyObjectAllowResourceType = "cspviolationreport"
	CrawlNewParamsBodyObjectAllowResourceTypePreflight          CrawlNewParamsBodyObjectAllowResourceType = "preflight"
	CrawlNewParamsBodyObjectAllowResourceTypeOther              CrawlNewParamsBodyObjectAllowResourceType = "other"
)

func (r CrawlNewParamsBodyObjectAllowResourceType) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectAllowResourceTypeDocument, CrawlNewParamsBodyObjectAllowResourceTypeStylesheet, CrawlNewParamsBodyObjectAllowResourceTypeImage, CrawlNewParamsBodyObjectAllowResourceTypeMedia, CrawlNewParamsBodyObjectAllowResourceTypeFont, CrawlNewParamsBodyObjectAllowResourceTypeScript, CrawlNewParamsBodyObjectAllowResourceTypeTexttrack, CrawlNewParamsBodyObjectAllowResourceTypeXHR, CrawlNewParamsBodyObjectAllowResourceTypeFetch, CrawlNewParamsBodyObjectAllowResourceTypePrefetch, CrawlNewParamsBodyObjectAllowResourceTypeEventsource, CrawlNewParamsBodyObjectAllowResourceTypeWebsocket, CrawlNewParamsBodyObjectAllowResourceTypeManifest, CrawlNewParamsBodyObjectAllowResourceTypeSignedexchange, CrawlNewParamsBodyObjectAllowResourceTypePing, CrawlNewParamsBodyObjectAllowResourceTypeCspviolationreport, CrawlNewParamsBodyObjectAllowResourceTypePreflight, CrawlNewParamsBodyObjectAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type CrawlNewParamsBodyObjectAuthenticate struct {
	Password param.Field[string] `json:"password" api:"required"`
	Username param.Field[string] `json:"username" api:"required"`
}

func (r CrawlNewParamsBodyObjectAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectCookie struct {
	// Cookie name.
	Name         param.Field[string]                                      `json:"name" api:"required"`
	Value        param.Field[string]                                      `json:"value" api:"required"`
	Domain       param.Field[string]                                      `json:"domain"`
	Expires      param.Field[float64]                                     `json:"expires"`
	HTTPOnly     param.Field[bool]                                        `json:"httpOnly"`
	PartitionKey param.Field[string]                                      `json:"partitionKey"`
	Path         param.Field[string]                                      `json:"path"`
	Priority     param.Field[CrawlNewParamsBodyObjectCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                        `json:"sameParty"`
	SameSite     param.Field[CrawlNewParamsBodyObjectCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                        `json:"secure"`
	SourcePort   param.Field[float64]                                     `json:"sourcePort"`
	SourceScheme param.Field[CrawlNewParamsBodyObjectCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                      `json:"url"`
}

func (r CrawlNewParamsBodyObjectCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectCookiesPriority string

const (
	CrawlNewParamsBodyObjectCookiesPriorityLow    CrawlNewParamsBodyObjectCookiesPriority = "Low"
	CrawlNewParamsBodyObjectCookiesPriorityMedium CrawlNewParamsBodyObjectCookiesPriority = "Medium"
	CrawlNewParamsBodyObjectCookiesPriorityHigh   CrawlNewParamsBodyObjectCookiesPriority = "High"
)

func (r CrawlNewParamsBodyObjectCookiesPriority) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectCookiesPriorityLow, CrawlNewParamsBodyObjectCookiesPriorityMedium, CrawlNewParamsBodyObjectCookiesPriorityHigh:
		return true
	}
	return false
}

type CrawlNewParamsBodyObjectCookiesSameSite string

const (
	CrawlNewParamsBodyObjectCookiesSameSiteStrict CrawlNewParamsBodyObjectCookiesSameSite = "Strict"
	CrawlNewParamsBodyObjectCookiesSameSiteLax    CrawlNewParamsBodyObjectCookiesSameSite = "Lax"
	CrawlNewParamsBodyObjectCookiesSameSiteNone   CrawlNewParamsBodyObjectCookiesSameSite = "None"
)

func (r CrawlNewParamsBodyObjectCookiesSameSite) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectCookiesSameSiteStrict, CrawlNewParamsBodyObjectCookiesSameSiteLax, CrawlNewParamsBodyObjectCookiesSameSiteNone:
		return true
	}
	return false
}

type CrawlNewParamsBodyObjectCookiesSourceScheme string

const (
	CrawlNewParamsBodyObjectCookiesSourceSchemeUnset     CrawlNewParamsBodyObjectCookiesSourceScheme = "Unset"
	CrawlNewParamsBodyObjectCookiesSourceSchemeNonSecure CrawlNewParamsBodyObjectCookiesSourceScheme = "NonSecure"
	CrawlNewParamsBodyObjectCookiesSourceSchemeSecure    CrawlNewParamsBodyObjectCookiesSourceScheme = "Secure"
)

func (r CrawlNewParamsBodyObjectCookiesSourceScheme) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectCookiesSourceSchemeUnset, CrawlNewParamsBodyObjectCookiesSourceSchemeNonSecure, CrawlNewParamsBodyObjectCookiesSourceSchemeSecure:
		return true
	}
	return false
}

type CrawlNewParamsBodyObjectCrawlPurpose string

const (
	CrawlNewParamsBodyObjectCrawlPurposeSearch  CrawlNewParamsBodyObjectCrawlPurpose = "search"
	CrawlNewParamsBodyObjectCrawlPurposeAIInput CrawlNewParamsBodyObjectCrawlPurpose = "ai-input"
	CrawlNewParamsBodyObjectCrawlPurposeAITrain CrawlNewParamsBodyObjectCrawlPurpose = "ai-train"
)

func (r CrawlNewParamsBodyObjectCrawlPurpose) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectCrawlPurposeSearch, CrawlNewParamsBodyObjectCrawlPurposeAIInput, CrawlNewParamsBodyObjectCrawlPurposeAITrain:
		return true
	}
	return false
}

type CrawlNewParamsBodyObjectFormat string

const (
	CrawlNewParamsBodyObjectFormatHTML     CrawlNewParamsBodyObjectFormat = "html"
	CrawlNewParamsBodyObjectFormatMarkdown CrawlNewParamsBodyObjectFormat = "markdown"
	CrawlNewParamsBodyObjectFormatJson     CrawlNewParamsBodyObjectFormat = "json"
)

func (r CrawlNewParamsBodyObjectFormat) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectFormatHTML, CrawlNewParamsBodyObjectFormatMarkdown, CrawlNewParamsBodyObjectFormatJson:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type CrawlNewParamsBodyObjectGotoOptions struct {
	Referer        param.Field[string]                                            `json:"referer"`
	ReferrerPolicy param.Field[string]                                            `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                           `json:"timeout"`
	WaitUntil      param.Field[CrawlNewParamsBodyObjectGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r CrawlNewParamsBodyObjectGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [browser_rendering.CrawlNewParamsBodyObjectGotoOptionsWaitUntilString],
// [browser_rendering.CrawlNewParamsBodyObjectGotoOptionsWaitUntilArray].
type CrawlNewParamsBodyObjectGotoOptionsWaitUntilUnion interface {
	implementsCrawlNewParamsBodyObjectGotoOptionsWaitUntilUnion()
}

type CrawlNewParamsBodyObjectGotoOptionsWaitUntilString string

const (
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringLoad             CrawlNewParamsBodyObjectGotoOptionsWaitUntilString = "load"
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded CrawlNewParamsBodyObjectGotoOptionsWaitUntilString = "domcontentloaded"
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0     CrawlNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle0"
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2     CrawlNewParamsBodyObjectGotoOptionsWaitUntilString = "networkidle2"
)

func (r CrawlNewParamsBodyObjectGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringLoad, CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringDomcontentloaded, CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle0, CrawlNewParamsBodyObjectGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r CrawlNewParamsBodyObjectGotoOptionsWaitUntilString) implementsCrawlNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type CrawlNewParamsBodyObjectGotoOptionsWaitUntilArray []CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItem

func (r CrawlNewParamsBodyObjectGotoOptionsWaitUntilArray) implementsCrawlNewParamsBodyObjectGotoOptionsWaitUntilUnion() {
}

type CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItem string

const (
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad             CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "load"
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0     CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle0"
	CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2     CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemLoad, CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemDomcontentloaded, CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle0, CrawlNewParamsBodyObjectGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

// Options for JSON extraction.
type CrawlNewParamsBodyObjectJsonOptions struct {
	// Optional list of custom AI models to use for the request. The models will be
	// tried in the order provided, and in case a model returns an error, the next one
	// will be used as fallback.
	CustomAI       param.Field[[]CrawlNewParamsBodyObjectJsonOptionsCustomAI]     `json:"custom_ai"`
	Prompt         param.Field[string]                                            `json:"prompt"`
	ResponseFormat param.Field[CrawlNewParamsBodyObjectJsonOptionsResponseFormat] `json:"response_format"`
}

func (r CrawlNewParamsBodyObjectJsonOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectJsonOptionsCustomAI struct {
	// AI model to use for the request. Must be formed as `<provider>/<model_name>`,
	// e.g. `workers-ai/@cf/meta/llama-3.3-70b-instruct-fp8-fast`.
	Model param.Field[string] `json:"model" api:"required"`
	// Authorization token for the AI model: `Bearer <token>`. Not needed for
	// workers-ai models.
	Authorization param.Field[string] `json:"authorization"`
}

func (r CrawlNewParamsBodyObjectJsonOptionsCustomAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectJsonOptionsResponseFormat struct {
	Type param.Field[string] `json:"type" api:"required"`
	// Schema for the response format. More information here:
	// https://developers.cloudflare.com/workers-ai/json-mode/
	JsonSchema param.Field[map[string]CrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion] `json:"json_schema"`
}

func (r CrawlNewParamsBodyObjectJsonOptionsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [browser_rendering.CrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaArray].
//
// Use [Raw()] to specify an arbitrary value for this param
type CrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion interface {
	ImplementsCrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion()
}

type CrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaArray []string

func (r CrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaArray) ImplementsCrawlNewParamsBodyObjectJsonOptionsResponseFormatJsonSchemaUnion() {
}

// Additional options for the crawler.
type CrawlNewParamsBodyObjectOptions struct {
	// Exclude links matching the provided wildcard patterns in the crawl job. Example:
	// 'https://example.com/privacy/**'.
	ExcludePatterns param.Field[[]string] `json:"excludePatterns"`
	// Include external links in the crawl job. If set to true, includeSubdomains is
	// ignored.
	IncludeExternalLinks param.Field[bool] `json:"includeExternalLinks"`
	// Include only links matching the provided wildcard patterns in the crawl job.
	// Include patterns are evaluated before exclude patterns. URLs that match any of
	// the specified include patterns will be included in the crawl job. Example:
	// 'https://example.com/blog/**'.
	IncludePatterns param.Field[[]string] `json:"includePatterns"`
	// Include links to subdomains in the crawl job. This option is ignored if
	// includeExternalLinks is true.
	IncludeSubdomains param.Field[bool] `json:"includeSubdomains"`
}

func (r CrawlNewParamsBodyObjectOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectRejectResourceType string

const (
	CrawlNewParamsBodyObjectRejectResourceTypeDocument           CrawlNewParamsBodyObjectRejectResourceType = "document"
	CrawlNewParamsBodyObjectRejectResourceTypeStylesheet         CrawlNewParamsBodyObjectRejectResourceType = "stylesheet"
	CrawlNewParamsBodyObjectRejectResourceTypeImage              CrawlNewParamsBodyObjectRejectResourceType = "image"
	CrawlNewParamsBodyObjectRejectResourceTypeMedia              CrawlNewParamsBodyObjectRejectResourceType = "media"
	CrawlNewParamsBodyObjectRejectResourceTypeFont               CrawlNewParamsBodyObjectRejectResourceType = "font"
	CrawlNewParamsBodyObjectRejectResourceTypeScript             CrawlNewParamsBodyObjectRejectResourceType = "script"
	CrawlNewParamsBodyObjectRejectResourceTypeTexttrack          CrawlNewParamsBodyObjectRejectResourceType = "texttrack"
	CrawlNewParamsBodyObjectRejectResourceTypeXHR                CrawlNewParamsBodyObjectRejectResourceType = "xhr"
	CrawlNewParamsBodyObjectRejectResourceTypeFetch              CrawlNewParamsBodyObjectRejectResourceType = "fetch"
	CrawlNewParamsBodyObjectRejectResourceTypePrefetch           CrawlNewParamsBodyObjectRejectResourceType = "prefetch"
	CrawlNewParamsBodyObjectRejectResourceTypeEventsource        CrawlNewParamsBodyObjectRejectResourceType = "eventsource"
	CrawlNewParamsBodyObjectRejectResourceTypeWebsocket          CrawlNewParamsBodyObjectRejectResourceType = "websocket"
	CrawlNewParamsBodyObjectRejectResourceTypeManifest           CrawlNewParamsBodyObjectRejectResourceType = "manifest"
	CrawlNewParamsBodyObjectRejectResourceTypeSignedexchange     CrawlNewParamsBodyObjectRejectResourceType = "signedexchange"
	CrawlNewParamsBodyObjectRejectResourceTypePing               CrawlNewParamsBodyObjectRejectResourceType = "ping"
	CrawlNewParamsBodyObjectRejectResourceTypeCspviolationreport CrawlNewParamsBodyObjectRejectResourceType = "cspviolationreport"
	CrawlNewParamsBodyObjectRejectResourceTypePreflight          CrawlNewParamsBodyObjectRejectResourceType = "preflight"
	CrawlNewParamsBodyObjectRejectResourceTypeOther              CrawlNewParamsBodyObjectRejectResourceType = "other"
)

func (r CrawlNewParamsBodyObjectRejectResourceType) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectRejectResourceTypeDocument, CrawlNewParamsBodyObjectRejectResourceTypeStylesheet, CrawlNewParamsBodyObjectRejectResourceTypeImage, CrawlNewParamsBodyObjectRejectResourceTypeMedia, CrawlNewParamsBodyObjectRejectResourceTypeFont, CrawlNewParamsBodyObjectRejectResourceTypeScript, CrawlNewParamsBodyObjectRejectResourceTypeTexttrack, CrawlNewParamsBodyObjectRejectResourceTypeXHR, CrawlNewParamsBodyObjectRejectResourceTypeFetch, CrawlNewParamsBodyObjectRejectResourceTypePrefetch, CrawlNewParamsBodyObjectRejectResourceTypeEventsource, CrawlNewParamsBodyObjectRejectResourceTypeWebsocket, CrawlNewParamsBodyObjectRejectResourceTypeManifest, CrawlNewParamsBodyObjectRejectResourceTypeSignedexchange, CrawlNewParamsBodyObjectRejectResourceTypePing, CrawlNewParamsBodyObjectRejectResourceTypeCspviolationreport, CrawlNewParamsBodyObjectRejectResourceTypePreflight, CrawlNewParamsBodyObjectRejectResourceTypeOther:
		return true
	}
	return false
}

// Whether to render the page or fetch static content. True by default.
type CrawlNewParamsBodyObjectRender bool

const (
	CrawlNewParamsBodyObjectRenderTrue CrawlNewParamsBodyObjectRender = true
)

func (r CrawlNewParamsBodyObjectRender) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectRenderTrue:
		return true
	}
	return false
}

// Source of links to crawl. 'sitemaps' - only crawl URLs from sitemaps, 'links' -
// only crawl URLs scraped from pages, 'all' - crawl both sitemap and scraped links
// (default).
type CrawlNewParamsBodyObjectSource string

const (
	CrawlNewParamsBodyObjectSourceSitemaps CrawlNewParamsBodyObjectSource = "sitemaps"
	CrawlNewParamsBodyObjectSourceLinks    CrawlNewParamsBodyObjectSource = "links"
	CrawlNewParamsBodyObjectSourceAll      CrawlNewParamsBodyObjectSource = "all"
)

func (r CrawlNewParamsBodyObjectSource) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectSourceSitemaps, CrawlNewParamsBodyObjectSourceLinks, CrawlNewParamsBodyObjectSourceAll:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type CrawlNewParamsBodyObjectViewport struct {
	Height            param.Field[float64] `json:"height" api:"required"`
	Width             param.Field[float64] `json:"width" api:"required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r CrawlNewParamsBodyObjectViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type CrawlNewParamsBodyObjectWaitForSelector struct {
	Selector param.Field[string]                                         `json:"selector" api:"required"`
	Hidden   param.Field[CrawlNewParamsBodyObjectWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                        `json:"timeout"`
	Visible  param.Field[CrawlNewParamsBodyObjectWaitForSelectorVisible] `json:"visible"`
}

func (r CrawlNewParamsBodyObjectWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CrawlNewParamsBodyObjectWaitForSelectorHidden bool

const (
	CrawlNewParamsBodyObjectWaitForSelectorHiddenTrue CrawlNewParamsBodyObjectWaitForSelectorHidden = true
)

func (r CrawlNewParamsBodyObjectWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type CrawlNewParamsBodyObjectWaitForSelectorVisible bool

const (
	CrawlNewParamsBodyObjectWaitForSelectorVisibleTrue CrawlNewParamsBodyObjectWaitForSelectorVisible = true
)

func (r CrawlNewParamsBodyObjectWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyObjectWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

// Whether to render the page or fetch static content. True by default.
type CrawlNewParamsBodyRender bool

const (
	CrawlNewParamsBodyRenderTrue  CrawlNewParamsBodyRender = true
	CrawlNewParamsBodyRenderFalse CrawlNewParamsBodyRender = false
)

func (r CrawlNewParamsBodyRender) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodyRenderTrue, CrawlNewParamsBodyRenderFalse:
		return true
	}
	return false
}

// Source of links to crawl. 'sitemaps' - only crawl URLs from sitemaps, 'links' -
// only crawl URLs scraped from pages, 'all' - crawl both sitemap and scraped links
// (default).
type CrawlNewParamsBodySource string

const (
	CrawlNewParamsBodySourceSitemaps CrawlNewParamsBodySource = "sitemaps"
	CrawlNewParamsBodySourceLinks    CrawlNewParamsBodySource = "links"
	CrawlNewParamsBodySourceAll      CrawlNewParamsBodySource = "all"
)

func (r CrawlNewParamsBodySource) IsKnown() bool {
	switch r {
	case CrawlNewParamsBodySourceSitemaps, CrawlNewParamsBodySourceLinks, CrawlNewParamsBodySourceAll:
		return true
	}
	return false
}

type CrawlNewResponseEnvelope struct {
	// Crawl job ID.
	Result string `json:"result" api:"required"`
	// Response status.
	Success bool                             `json:"success" api:"required"`
	Errors  []CrawlNewResponseEnvelopeErrors `json:"errors"`
	JSON    crawlNewResponseEnvelopeJSON     `json:"-"`
}

// crawlNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CrawlNewResponseEnvelope]
type crawlNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CrawlNewResponseEnvelopeErrors struct {
	// Error code.
	Code float64 `json:"code" api:"required"`
	// Error message.
	Message string                             `json:"message" api:"required"`
	JSON    crawlNewResponseEnvelopeErrorsJSON `json:"-"`
}

// crawlNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CrawlNewResponseEnvelopeErrors]
type crawlNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CrawlDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CrawlDeleteResponseEnvelope struct {
	Result CrawlDeleteResponse `json:"result" api:"required"`
	// Response status.
	Success bool                                `json:"success" api:"required"`
	Errors  []CrawlDeleteResponseEnvelopeErrors `json:"errors"`
	JSON    crawlDeleteResponseEnvelopeJSON     `json:"-"`
}

// crawlDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CrawlDeleteResponseEnvelope]
type crawlDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CrawlDeleteResponseEnvelopeErrors struct {
	// Error code.
	Code float64 `json:"code" api:"required"`
	// Error message.
	Message string                                `json:"message" api:"required"`
	JSON    crawlDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// crawlDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CrawlDeleteResponseEnvelopeErrors]
type crawlDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CrawlGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// Cursor for pagination.
	Cursor param.Field[float64] `query:"cursor"`
	// Limit for pagination.
	Limit param.Field[float64] `query:"limit"`
	// Filter by URL status.
	Status param.Field[CrawlGetParamsStatus] `query:"status"`
}

// URLQuery serializes [CrawlGetParams]'s query parameters as `url.Values`.
func (r CrawlGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by URL status.
type CrawlGetParamsStatus string

const (
	CrawlGetParamsStatusQueued     CrawlGetParamsStatus = "queued"
	CrawlGetParamsStatusErrored    CrawlGetParamsStatus = "errored"
	CrawlGetParamsStatusCompleted  CrawlGetParamsStatus = "completed"
	CrawlGetParamsStatusDisallowed CrawlGetParamsStatus = "disallowed"
	CrawlGetParamsStatusSkipped    CrawlGetParamsStatus = "skipped"
	CrawlGetParamsStatusCancelled  CrawlGetParamsStatus = "cancelled"
)

func (r CrawlGetParamsStatus) IsKnown() bool {
	switch r {
	case CrawlGetParamsStatusQueued, CrawlGetParamsStatusErrored, CrawlGetParamsStatusCompleted, CrawlGetParamsStatusDisallowed, CrawlGetParamsStatusSkipped, CrawlGetParamsStatusCancelled:
		return true
	}
	return false
}

type CrawlGetResponseEnvelope struct {
	Result CrawlGetResponse `json:"result" api:"required"`
	// Response status.
	Success bool                             `json:"success" api:"required"`
	Errors  []CrawlGetResponseEnvelopeErrors `json:"errors"`
	JSON    crawlGetResponseEnvelopeJSON     `json:"-"`
}

// crawlGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CrawlGetResponseEnvelope]
type crawlGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CrawlGetResponseEnvelopeErrors struct {
	// Error code.
	Code float64 `json:"code" api:"required"`
	// Error message.
	Message string                             `json:"message" api:"required"`
	JSON    crawlGetResponseEnvelopeErrorsJSON `json:"-"`
}

// crawlGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CrawlGetResponseEnvelopeErrors]
type crawlGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CrawlGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r crawlGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
