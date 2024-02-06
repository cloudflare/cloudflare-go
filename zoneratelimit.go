// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRateLimitService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRateLimitService] method
// instead.
type ZoneRateLimitService struct {
	Options []option.RequestOption
}

// NewZoneRateLimitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRateLimitService(opts ...option.RequestOption) (r *ZoneRateLimitService) {
	r = &ZoneRateLimitService{}
	r.Options = opts
	return
}

// Fetches the details of a rate limit.
func (r *ZoneRateLimitService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneRateLimitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the rate limits for a zone.
func (r *ZoneRateLimitService) List(ctx context.Context, zoneIdentifier string, query ZoneRateLimitListParams, opts ...option.RequestOption) (res *ZoneRateLimitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneRateLimitGetResponse struct {
	Errors   []ZoneRateLimitGetResponseError   `json:"errors"`
	Messages []ZoneRateLimitGetResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success ZoneRateLimitGetResponseSuccess `json:"success"`
	JSON    zoneRateLimitGetResponseJSON    `json:"-"`
}

// zoneRateLimitGetResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitGetResponse]
type zoneRateLimitGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneRateLimitGetResponseErrorJSON `json:"-"`
}

// zoneRateLimitGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneRateLimitGetResponseError]
type zoneRateLimitGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneRateLimitGetResponseMessageJSON `json:"-"`
}

// zoneRateLimitGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRateLimitGetResponseMessage]
type zoneRateLimitGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRateLimitGetResponseSuccess bool

const (
	ZoneRateLimitGetResponseSuccessTrue ZoneRateLimitGetResponseSuccess = true
)

type ZoneRateLimitListResponse struct {
	Errors     []ZoneRateLimitListResponseError    `json:"errors"`
	Messages   []ZoneRateLimitListResponseMessage  `json:"messages"`
	Result     []ZoneRateLimitListResponseResult   `json:"result"`
	ResultInfo ZoneRateLimitListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneRateLimitListResponseSuccess `json:"success"`
	JSON    zoneRateLimitListResponseJSON    `json:"-"`
}

// zoneRateLimitListResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponse]
type zoneRateLimitListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneRateLimitListResponseErrorJSON `json:"-"`
}

// zoneRateLimitListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponseError]
type zoneRateLimitListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneRateLimitListResponseMessageJSON `json:"-"`
}

// zoneRateLimitListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponseMessage]
type zoneRateLimitListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseResult struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action ZoneRateLimitListResponseResultAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []ZoneRateLimitListResponseResultBypass `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match ZoneRateLimitListResponseResultMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                             `json:"threshold"`
	JSON      zoneRateLimitListResponseResultJSON `json:"-"`
}

// zoneRateLimitListResponseResultJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponseResult]
type zoneRateLimitListResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Bypass      apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	Match       apijson.Field
	Period      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type ZoneRateLimitListResponseResultAction struct {
	// The action to perform.
	Mode ZoneRateLimitListResponseResultActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response ZoneRateLimitListResponseResultActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64                                   `json:"timeout"`
	JSON    zoneRateLimitListResponseResultActionJSON `json:"-"`
}

// zoneRateLimitListResponseResultActionJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseResultAction]
type zoneRateLimitListResponseResultActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform.
type ZoneRateLimitListResponseResultActionMode string

const (
	ZoneRateLimitListResponseResultActionModeSimulate         ZoneRateLimitListResponseResultActionMode = "simulate"
	ZoneRateLimitListResponseResultActionModeBan              ZoneRateLimitListResponseResultActionMode = "ban"
	ZoneRateLimitListResponseResultActionModeChallenge        ZoneRateLimitListResponseResultActionMode = "challenge"
	ZoneRateLimitListResponseResultActionModeJsChallenge      ZoneRateLimitListResponseResultActionMode = "js_challenge"
	ZoneRateLimitListResponseResultActionModeManagedChallenge ZoneRateLimitListResponseResultActionMode = "managed_challenge"
)

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type ZoneRateLimitListResponseResultActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                                            `json:"content_type"`
	JSON        zoneRateLimitListResponseResultActionResponseJSON `json:"-"`
}

// zoneRateLimitListResponseResultActionResponseJSON contains the JSON metadata for
// the struct [ZoneRateLimitListResponseResultActionResponse]
type zoneRateLimitListResponseResultActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseResultBypass struct {
	Name ZoneRateLimitListResponseResultBypassName `json:"name"`
	// The URL to bypass.
	Value string                                    `json:"value"`
	JSON  zoneRateLimitListResponseResultBypassJSON `json:"-"`
}

// zoneRateLimitListResponseResultBypassJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseResultBypass]
type zoneRateLimitListResponseResultBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseResultBypassName string

const (
	ZoneRateLimitListResponseResultBypassNameURL ZoneRateLimitListResponseResultBypassName = "url"
)

// Determines which traffic the rate limit counts towards the threshold.
type ZoneRateLimitListResponseResultMatch struct {
	Headers  []ZoneRateLimitListResponseResultMatchHeader `json:"headers"`
	Request  ZoneRateLimitListResponseResultMatchRequest  `json:"request"`
	Response ZoneRateLimitListResponseResultMatchResponse `json:"response"`
	JSON     zoneRateLimitListResponseResultMatchJSON     `json:"-"`
}

// zoneRateLimitListResponseResultMatchJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseResultMatch]
type zoneRateLimitListResponseResultMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseResultMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op ZoneRateLimitListResponseResultMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                                         `json:"value"`
	JSON  zoneRateLimitListResponseResultMatchHeaderJSON `json:"-"`
}

// zoneRateLimitListResponseResultMatchHeaderJSON contains the JSON metadata for
// the struct [ZoneRateLimitListResponseResultMatchHeader]
type zoneRateLimitListResponseResultMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type ZoneRateLimitListResponseResultMatchHeadersOp string

const (
	ZoneRateLimitListResponseResultMatchHeadersOpEq ZoneRateLimitListResponseResultMatchHeadersOp = "eq"
	ZoneRateLimitListResponseResultMatchHeadersOpNe ZoneRateLimitListResponseResultMatchHeadersOp = "ne"
)

type ZoneRateLimitListResponseResultMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []ZoneRateLimitListResponseResultMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                                          `json:"url"`
	JSON zoneRateLimitListResponseResultMatchRequestJSON `json:"-"`
}

// zoneRateLimitListResponseResultMatchRequestJSON contains the JSON metadata for
// the struct [ZoneRateLimitListResponseResultMatchRequest]
type zoneRateLimitListResponseResultMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type ZoneRateLimitListResponseResultMatchRequestMethod string

const (
	ZoneRateLimitListResponseResultMatchRequestMethodGet    ZoneRateLimitListResponseResultMatchRequestMethod = "GET"
	ZoneRateLimitListResponseResultMatchRequestMethodPost   ZoneRateLimitListResponseResultMatchRequestMethod = "POST"
	ZoneRateLimitListResponseResultMatchRequestMethodPut    ZoneRateLimitListResponseResultMatchRequestMethod = "PUT"
	ZoneRateLimitListResponseResultMatchRequestMethodDelete ZoneRateLimitListResponseResultMatchRequestMethod = "DELETE"
	ZoneRateLimitListResponseResultMatchRequestMethodPatch  ZoneRateLimitListResponseResultMatchRequestMethod = "PATCH"
	ZoneRateLimitListResponseResultMatchRequestMethodHead   ZoneRateLimitListResponseResultMatchRequestMethod = "HEAD"
	ZoneRateLimitListResponseResultMatchRequestMethod_All   ZoneRateLimitListResponseResultMatchRequestMethod = "_ALL_"
)

type ZoneRateLimitListResponseResultMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                                             `json:"origin_traffic"`
	JSON          zoneRateLimitListResponseResultMatchResponseJSON `json:"-"`
}

// zoneRateLimitListResponseResultMatchResponseJSON contains the JSON metadata for
// the struct [ZoneRateLimitListResponseResultMatchResponse]
type zoneRateLimitListResponseResultMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       zoneRateLimitListResponseResultInfoJSON `json:"-"`
}

// zoneRateLimitListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseResultInfo]
type zoneRateLimitListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRateLimitListResponseSuccess bool

const (
	ZoneRateLimitListResponseSuccessTrue ZoneRateLimitListResponseSuccess = true
)

type ZoneRateLimitListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneRateLimitListParams]'s query parameters as
// `url.Values`.
func (r ZoneRateLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
