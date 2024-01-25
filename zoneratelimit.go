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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *ZoneRateLimitService) New(ctx context.Context, zoneIdentifier string, body ZoneRateLimitNewParams, opts ...option.RequestOption) (res *ZoneRateLimitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a rate limit.
func (r *ZoneRateLimitService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneRateLimitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing rate limit.
func (r *ZoneRateLimitService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneRateLimitUpdateParams, opts ...option.RequestOption) (res *ZoneRateLimitUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches the rate limits for a zone.
func (r *ZoneRateLimitService) List(ctx context.Context, zoneIdentifier string, query ZoneRateLimitListParams, opts ...option.RequestOption) (res *shared.Page[ZoneRateLimitListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Deletes an existing rate limit.
func (r *ZoneRateLimitService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneRateLimitDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneRateLimitNewResponse struct {
	Errors   []ZoneRateLimitNewResponseError   `json:"errors"`
	Messages []ZoneRateLimitNewResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success ZoneRateLimitNewResponseSuccess `json:"success"`
	JSON    zoneRateLimitNewResponseJSON    `json:"-"`
}

// zoneRateLimitNewResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitNewResponse]
type zoneRateLimitNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitNewResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneRateLimitNewResponseErrorJSON `json:"-"`
}

// zoneRateLimitNewResponseErrorJSON contains the JSON metadata for the struct
// [ZoneRateLimitNewResponseError]
type zoneRateLimitNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitNewResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneRateLimitNewResponseMessageJSON `json:"-"`
}

// zoneRateLimitNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRateLimitNewResponseMessage]
type zoneRateLimitNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRateLimitNewResponseSuccess bool

const (
	ZoneRateLimitNewResponseSuccessTrue ZoneRateLimitNewResponseSuccess = true
)

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

type ZoneRateLimitUpdateResponse struct {
	Errors   []ZoneRateLimitUpdateResponseError   `json:"errors"`
	Messages []ZoneRateLimitUpdateResponseMessage `json:"messages"`
	Result   interface{}                          `json:"result"`
	// Whether the API call was successful
	Success ZoneRateLimitUpdateResponseSuccess `json:"success"`
	JSON    zoneRateLimitUpdateResponseJSON    `json:"-"`
}

// zoneRateLimitUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitUpdateResponse]
type zoneRateLimitUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitUpdateResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneRateLimitUpdateResponseErrorJSON `json:"-"`
}

// zoneRateLimitUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneRateLimitUpdateResponseError]
type zoneRateLimitUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitUpdateResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneRateLimitUpdateResponseMessageJSON `json:"-"`
}

// zoneRateLimitUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRateLimitUpdateResponseMessage]
type zoneRateLimitUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRateLimitUpdateResponseSuccess bool

const (
	ZoneRateLimitUpdateResponseSuccessTrue ZoneRateLimitUpdateResponseSuccess = true
)

type ZoneRateLimitListResponse struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action ZoneRateLimitListResponseAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []ZoneRateLimitListResponseBypass `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match ZoneRateLimitListResponseMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                       `json:"threshold"`
	JSON      zoneRateLimitListResponseJSON `json:"-"`
}

// zoneRateLimitListResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponse]
type zoneRateLimitListResponseJSON struct {
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

func (r *ZoneRateLimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type ZoneRateLimitListResponseAction struct {
	// The action to perform.
	Mode ZoneRateLimitListResponseActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response ZoneRateLimitListResponseActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64                             `json:"timeout"`
	JSON    zoneRateLimitListResponseActionJSON `json:"-"`
}

// zoneRateLimitListResponseActionJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponseAction]
type zoneRateLimitListResponseActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform.
type ZoneRateLimitListResponseActionMode string

const (
	ZoneRateLimitListResponseActionModeSimulate         ZoneRateLimitListResponseActionMode = "simulate"
	ZoneRateLimitListResponseActionModeBan              ZoneRateLimitListResponseActionMode = "ban"
	ZoneRateLimitListResponseActionModeChallenge        ZoneRateLimitListResponseActionMode = "challenge"
	ZoneRateLimitListResponseActionModeJsChallenge      ZoneRateLimitListResponseActionMode = "js_challenge"
	ZoneRateLimitListResponseActionModeManagedChallenge ZoneRateLimitListResponseActionMode = "managed_challenge"
)

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type ZoneRateLimitListResponseActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                                      `json:"content_type"`
	JSON        zoneRateLimitListResponseActionResponseJSON `json:"-"`
}

// zoneRateLimitListResponseActionResponseJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseActionResponse]
type zoneRateLimitListResponseActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseBypass struct {
	Name ZoneRateLimitListResponseBypassName `json:"name"`
	// The URL to bypass.
	Value string                              `json:"value"`
	JSON  zoneRateLimitListResponseBypassJSON `json:"-"`
}

// zoneRateLimitListResponseBypassJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponseBypass]
type zoneRateLimitListResponseBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseBypassName string

const (
	ZoneRateLimitListResponseBypassNameURL ZoneRateLimitListResponseBypassName = "url"
)

// Determines which traffic the rate limit counts towards the threshold.
type ZoneRateLimitListResponseMatch struct {
	Headers  []ZoneRateLimitListResponseMatchHeader `json:"headers"`
	Request  ZoneRateLimitListResponseMatchRequest  `json:"request"`
	Response ZoneRateLimitListResponseMatchResponse `json:"response"`
	JSON     zoneRateLimitListResponseMatchJSON     `json:"-"`
}

// zoneRateLimitListResponseMatchJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponseMatch]
type zoneRateLimitListResponseMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRateLimitListResponseMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op ZoneRateLimitListResponseMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                                   `json:"value"`
	JSON  zoneRateLimitListResponseMatchHeaderJSON `json:"-"`
}

// zoneRateLimitListResponseMatchHeaderJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseMatchHeader]
type zoneRateLimitListResponseMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type ZoneRateLimitListResponseMatchHeadersOp string

const (
	ZoneRateLimitListResponseMatchHeadersOpEq ZoneRateLimitListResponseMatchHeadersOp = "eq"
	ZoneRateLimitListResponseMatchHeadersOpNe ZoneRateLimitListResponseMatchHeadersOp = "ne"
)

type ZoneRateLimitListResponseMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []ZoneRateLimitListResponseMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                                    `json:"url"`
	JSON zoneRateLimitListResponseMatchRequestJSON `json:"-"`
}

// zoneRateLimitListResponseMatchRequestJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseMatchRequest]
type zoneRateLimitListResponseMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type ZoneRateLimitListResponseMatchRequestMethod string

const (
	ZoneRateLimitListResponseMatchRequestMethodGet    ZoneRateLimitListResponseMatchRequestMethod = "GET"
	ZoneRateLimitListResponseMatchRequestMethodPost   ZoneRateLimitListResponseMatchRequestMethod = "POST"
	ZoneRateLimitListResponseMatchRequestMethodPut    ZoneRateLimitListResponseMatchRequestMethod = "PUT"
	ZoneRateLimitListResponseMatchRequestMethodDelete ZoneRateLimitListResponseMatchRequestMethod = "DELETE"
	ZoneRateLimitListResponseMatchRequestMethodPatch  ZoneRateLimitListResponseMatchRequestMethod = "PATCH"
	ZoneRateLimitListResponseMatchRequestMethodHead   ZoneRateLimitListResponseMatchRequestMethod = "HEAD"
	ZoneRateLimitListResponseMatchRequestMethod_All   ZoneRateLimitListResponseMatchRequestMethod = "_ALL_"
)

type ZoneRateLimitListResponseMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                                       `json:"origin_traffic"`
	JSON          zoneRateLimitListResponseMatchResponseJSON `json:"-"`
}

// zoneRateLimitListResponseMatchResponseJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseMatchResponse]
type zoneRateLimitListResponseMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
