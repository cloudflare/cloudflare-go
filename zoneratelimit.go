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

// Creates a new rate limit for a zone. Refer to the object definition for a list
// of required attributes.
func (r *ZoneRateLimitService) New(ctx context.Context, zoneIdentifier string, body ZoneRateLimitNewParams, opts ...option.RequestOption) (res *RatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a rate limit.
func (r *ZoneRateLimitService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *RatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing rate limit.
func (r *ZoneRateLimitService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneRateLimitUpdateParams, opts ...option.RequestOption) (res *RatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches the rate limits for a zone.
func (r *ZoneRateLimitService) List(ctx context.Context, zoneIdentifier string, query ZoneRateLimitListParams, opts ...option.RequestOption) (res *RatelimitCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing rate limit.
func (r *ZoneRateLimitService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneRateLimitDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RatelimitCollection struct {
	Errors     []RatelimitCollectionError    `json:"errors"`
	Messages   []RatelimitCollectionMessage  `json:"messages"`
	Result     []RatelimitCollectionResult   `json:"result"`
	ResultInfo RatelimitCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RatelimitCollectionSuccess `json:"success"`
	JSON    ratelimitCollectionJSON    `json:"-"`
}

// ratelimitCollectionJSON contains the JSON metadata for the struct
// [RatelimitCollection]
type ratelimitCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitCollectionError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    ratelimitCollectionErrorJSON `json:"-"`
}

// ratelimitCollectionErrorJSON contains the JSON metadata for the struct
// [RatelimitCollectionError]
type ratelimitCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitCollectionMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    ratelimitCollectionMessageJSON `json:"-"`
}

// ratelimitCollectionMessageJSON contains the JSON metadata for the struct
// [RatelimitCollectionMessage]
type ratelimitCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitCollectionResult struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action RatelimitCollectionResultAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []RatelimitCollectionResultBypass `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match RatelimitCollectionResultMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                       `json:"threshold"`
	JSON      ratelimitCollectionResultJSON `json:"-"`
}

// ratelimitCollectionResultJSON contains the JSON metadata for the struct
// [RatelimitCollectionResult]
type ratelimitCollectionResultJSON struct {
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

func (r *RatelimitCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RatelimitCollectionResultAction struct {
	// The action to perform.
	Mode RatelimitCollectionResultActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response RatelimitCollectionResultActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64                             `json:"timeout"`
	JSON    ratelimitCollectionResultActionJSON `json:"-"`
}

// ratelimitCollectionResultActionJSON contains the JSON metadata for the struct
// [RatelimitCollectionResultAction]
type ratelimitCollectionResultActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform.
type RatelimitCollectionResultActionMode string

const (
	RatelimitCollectionResultActionModeSimulate         RatelimitCollectionResultActionMode = "simulate"
	RatelimitCollectionResultActionModeBan              RatelimitCollectionResultActionMode = "ban"
	RatelimitCollectionResultActionModeChallenge        RatelimitCollectionResultActionMode = "challenge"
	RatelimitCollectionResultActionModeJsChallenge      RatelimitCollectionResultActionMode = "js_challenge"
	RatelimitCollectionResultActionModeManagedChallenge RatelimitCollectionResultActionMode = "managed_challenge"
)

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RatelimitCollectionResultActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                                      `json:"content_type"`
	JSON        ratelimitCollectionResultActionResponseJSON `json:"-"`
}

// ratelimitCollectionResultActionResponseJSON contains the JSON metadata for the
// struct [RatelimitCollectionResultActionResponse]
type ratelimitCollectionResultActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionResultActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitCollectionResultBypass struct {
	Name RatelimitCollectionResultBypassName `json:"name"`
	// The URL to bypass.
	Value string                              `json:"value"`
	JSON  ratelimitCollectionResultBypassJSON `json:"-"`
}

// ratelimitCollectionResultBypassJSON contains the JSON metadata for the struct
// [RatelimitCollectionResultBypass]
type ratelimitCollectionResultBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionResultBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitCollectionResultBypassName string

const (
	RatelimitCollectionResultBypassNameURL RatelimitCollectionResultBypassName = "url"
)

// Determines which traffic the rate limit counts towards the threshold.
type RatelimitCollectionResultMatch struct {
	Headers  []RatelimitCollectionResultMatchHeader `json:"headers"`
	Request  RatelimitCollectionResultMatchRequest  `json:"request"`
	Response RatelimitCollectionResultMatchResponse `json:"response"`
	JSON     ratelimitCollectionResultMatchJSON     `json:"-"`
}

// ratelimitCollectionResultMatchJSON contains the JSON metadata for the struct
// [RatelimitCollectionResultMatch]
type ratelimitCollectionResultMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitCollectionResultMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op RatelimitCollectionResultMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                                   `json:"value"`
	JSON  ratelimitCollectionResultMatchHeaderJSON `json:"-"`
}

// ratelimitCollectionResultMatchHeaderJSON contains the JSON metadata for the
// struct [RatelimitCollectionResultMatchHeader]
type ratelimitCollectionResultMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionResultMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RatelimitCollectionResultMatchHeadersOp string

const (
	RatelimitCollectionResultMatchHeadersOpEq RatelimitCollectionResultMatchHeadersOp = "eq"
	RatelimitCollectionResultMatchHeadersOpNe RatelimitCollectionResultMatchHeadersOp = "ne"
)

type RatelimitCollectionResultMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []RatelimitCollectionResultMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                                    `json:"url"`
	JSON ratelimitCollectionResultMatchRequestJSON `json:"-"`
}

// ratelimitCollectionResultMatchRequestJSON contains the JSON metadata for the
// struct [RatelimitCollectionResultMatchRequest]
type ratelimitCollectionResultMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionResultMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type RatelimitCollectionResultMatchRequestMethod string

const (
	RatelimitCollectionResultMatchRequestMethodGet    RatelimitCollectionResultMatchRequestMethod = "GET"
	RatelimitCollectionResultMatchRequestMethodPost   RatelimitCollectionResultMatchRequestMethod = "POST"
	RatelimitCollectionResultMatchRequestMethodPut    RatelimitCollectionResultMatchRequestMethod = "PUT"
	RatelimitCollectionResultMatchRequestMethodDelete RatelimitCollectionResultMatchRequestMethod = "DELETE"
	RatelimitCollectionResultMatchRequestMethodPatch  RatelimitCollectionResultMatchRequestMethod = "PATCH"
	RatelimitCollectionResultMatchRequestMethodHead   RatelimitCollectionResultMatchRequestMethod = "HEAD"
	RatelimitCollectionResultMatchRequestMethod_All   RatelimitCollectionResultMatchRequestMethod = "_ALL_"
)

type RatelimitCollectionResultMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                                       `json:"origin_traffic"`
	JSON          ratelimitCollectionResultMatchResponseJSON `json:"-"`
}

// ratelimitCollectionResultMatchResponseJSON contains the JSON metadata for the
// struct [RatelimitCollectionResultMatchResponse]
type ratelimitCollectionResultMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RatelimitCollectionResultMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                           `json:"total_count"`
	JSON       ratelimitCollectionResultInfoJSON `json:"-"`
}

// ratelimitCollectionResultInfoJSON contains the JSON metadata for the struct
// [RatelimitCollectionResultInfo]
type ratelimitCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RatelimitCollectionSuccess bool

const (
	RatelimitCollectionSuccessTrue RatelimitCollectionSuccess = true
)

type RatelimitSingle struct {
	Errors   []RatelimitSingleError   `json:"errors"`
	Messages []RatelimitSingleMessage `json:"messages"`
	Result   interface{}              `json:"result"`
	// Whether the API call was successful
	Success RatelimitSingleSuccess `json:"success"`
	JSON    ratelimitSingleJSON    `json:"-"`
}

// ratelimitSingleJSON contains the JSON metadata for the struct [RatelimitSingle]
type ratelimitSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitSingleError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    ratelimitSingleErrorJSON `json:"-"`
}

// ratelimitSingleErrorJSON contains the JSON metadata for the struct
// [RatelimitSingleError]
type ratelimitSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RatelimitSingleMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    ratelimitSingleMessageJSON `json:"-"`
}

// ratelimitSingleMessageJSON contains the JSON metadata for the struct
// [RatelimitSingleMessage]
type ratelimitSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RatelimitSingleSuccess bool

const (
	RatelimitSingleSuccessTrue RatelimitSingleSuccess = true
)

type ZoneRateLimitDeleteResponse struct {
	Errors   []ZoneRateLimitDeleteResponseError   `json:"errors"`
	Messages []ZoneRateLimitDeleteResponseMessage `json:"messages"`
	Result   ZoneRateLimitDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneRateLimitDeleteResponseSuccess `json:"success"`
	JSON    zoneRateLimitDeleteResponseJSON    `json:"-"`
}

// zoneRateLimitDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitDeleteResponse]
type zoneRateLimitDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitDeleteResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneRateLimitDeleteResponseErrorJSON `json:"-"`
}

// zoneRateLimitDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneRateLimitDeleteResponseError]
type zoneRateLimitDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitDeleteResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneRateLimitDeleteResponseMessageJSON `json:"-"`
}

// zoneRateLimitDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRateLimitDeleteResponseMessage]
type zoneRateLimitDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitDeleteResponseResult struct {
	// The unique identifier of the rate limit.
	ID   string                                `json:"id"`
	JSON zoneRateLimitDeleteResponseResultJSON `json:"-"`
}

// zoneRateLimitDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneRateLimitDeleteResponseResult]
type zoneRateLimitDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRateLimitDeleteResponseSuccess bool

const (
	ZoneRateLimitDeleteResponseSuccessTrue ZoneRateLimitDeleteResponseSuccess = true
)

type ZoneRateLimitNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneRateLimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneRateLimitUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneRateLimitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

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
