// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limits

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// RateLimitService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRateLimitService] method instead.
type RateLimitService struct {
	Options []option.RequestOption
}

// NewRateLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRateLimitService(opts ...option.RequestOption) (r *RateLimitService) {
	r = &RateLimitService{}
	r.Options = opts
	return
}

// Creates a new rate limit for a zone. Refer to the object definition for a list
// of required attributes.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) New(ctx context.Context, zoneIdentifier string, body RateLimitNewParams, opts ...option.RequestOption) (res *RateLimitNewResponseUnion, err error) {
	var env RateLimitNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the rate limits for a zone.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) List(ctx context.Context, zoneIdentifier string, query RateLimitListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RateLimit], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
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

// Fetches the rate limits for a zone.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query RateLimitListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RateLimit] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing rate limit.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *RateLimitDeleteResponse, err error) {
	var env RateLimitDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rate limit.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Edit(ctx context.Context, zoneIdentifier string, id string, body RateLimitEditParams, opts ...option.RequestOption) (res *RateLimitEditResponseUnion, err error) {
	var env RateLimitEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a rate limit.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *RateLimitGetResponseUnion, err error) {
	var env RateLimitGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type Action string

const (
	ActionBlock            Action = "block"
	ActionChallenge        Action = "challenge"
	ActionJSChallenge      Action = "js_challenge"
	ActionManagedChallenge Action = "managed_challenge"
	ActionAllow            Action = "allow"
	ActionLog              Action = "log"
	ActionBypass           Action = "bypass"
)

func (r Action) IsKnown() bool {
	switch r {
	case ActionBlock, ActionChallenge, ActionJSChallenge, ActionManagedChallenge, ActionAllow, ActionLog, ActionBypass:
		return true
	}
	return false
}

// An HTTP method or `_ALL_` to indicate all methods.
type Methods string

const (
	MethodsGet    Methods = "GET"
	MethodsPost   Methods = "POST"
	MethodsPut    Methods = "PUT"
	MethodsDelete Methods = "DELETE"
	MethodsPatch  Methods = "PATCH"
	MethodsHead   Methods = "HEAD"
	Methods_All   Methods = "_ALL_"
)

func (r Methods) IsKnown() bool {
	switch r {
	case MethodsGet, MethodsPost, MethodsPut, MethodsDelete, MethodsPatch, MethodsHead, Methods_All:
		return true
	}
	return false
}

type RateLimit struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action RateLimitAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []RateLimitBypass `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match RateLimitMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64       `json:"threshold"`
	JSON      rateLimitJSON `json:"-"`
}

// rateLimitJSON contains the JSON metadata for the struct [RateLimit]
type rateLimitJSON struct {
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

func (r *RateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitAction struct {
	// The action to perform.
	Mode RateLimitActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response RateLimitActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64             `json:"timeout"`
	JSON    rateLimitActionJSON `json:"-"`
}

// rateLimitActionJSON contains the JSON metadata for the struct [RateLimitAction]
type rateLimitActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitActionJSON) RawJSON() string {
	return r.raw
}

// The action to perform.
type RateLimitActionMode string

const (
	RateLimitActionModeSimulate         RateLimitActionMode = "simulate"
	RateLimitActionModeBan              RateLimitActionMode = "ban"
	RateLimitActionModeChallenge        RateLimitActionMode = "challenge"
	RateLimitActionModeJSChallenge      RateLimitActionMode = "js_challenge"
	RateLimitActionModeManagedChallenge RateLimitActionMode = "managed_challenge"
)

func (r RateLimitActionMode) IsKnown() bool {
	switch r {
	case RateLimitActionModeSimulate, RateLimitActionModeBan, RateLimitActionModeChallenge, RateLimitActionModeJSChallenge, RateLimitActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RateLimitActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                      `json:"content_type"`
	JSON        rateLimitActionResponseJSON `json:"-"`
}

// rateLimitActionResponseJSON contains the JSON metadata for the struct
// [RateLimitActionResponse]
type rateLimitActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitActionResponseJSON) RawJSON() string {
	return r.raw
}

type RateLimitBypass struct {
	Name RateLimitBypassName `json:"name"`
	// The URL to bypass.
	Value string              `json:"value"`
	JSON  rateLimitBypassJSON `json:"-"`
}

// rateLimitBypassJSON contains the JSON metadata for the struct [RateLimitBypass]
type rateLimitBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitBypassJSON) RawJSON() string {
	return r.raw
}

type RateLimitBypassName string

const (
	RateLimitBypassNameURL RateLimitBypassName = "url"
)

func (r RateLimitBypassName) IsKnown() bool {
	switch r {
	case RateLimitBypassNameURL:
		return true
	}
	return false
}

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitMatch struct {
	Headers  []RateLimitMatchHeader `json:"headers"`
	Request  RateLimitMatchRequest  `json:"request"`
	Response RateLimitMatchResponse `json:"response"`
	JSON     rateLimitMatchJSON     `json:"-"`
}

// rateLimitMatchJSON contains the JSON metadata for the struct [RateLimitMatch]
type rateLimitMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchJSON) RawJSON() string {
	return r.raw
}

type RateLimitMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op RateLimitMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                   `json:"value"`
	JSON  rateLimitMatchHeaderJSON `json:"-"`
}

// rateLimitMatchHeaderJSON contains the JSON metadata for the struct
// [RateLimitMatchHeader]
type rateLimitMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchHeaderJSON) RawJSON() string {
	return r.raw
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitMatchHeadersOp string

const (
	RateLimitMatchHeadersOpEq RateLimitMatchHeadersOp = "eq"
	RateLimitMatchHeadersOpNe RateLimitMatchHeadersOp = "ne"
)

func (r RateLimitMatchHeadersOp) IsKnown() bool {
	switch r {
	case RateLimitMatchHeadersOpEq, RateLimitMatchHeadersOpNe:
		return true
	}
	return false
}

type RateLimitMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []Methods `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                    `json:"url"`
	JSON rateLimitMatchRequestJSON `json:"-"`
}

// rateLimitMatchRequestJSON contains the JSON metadata for the struct
// [RateLimitMatchRequest]
type rateLimitMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchRequestJSON) RawJSON() string {
	return r.raw
}

type RateLimitMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                       `json:"origin_traffic"`
	JSON          rateLimitMatchResponseJSON `json:"-"`
}

// rateLimitMatchResponseJSON contains the JSON metadata for the struct
// [RateLimitMatchResponse]
type rateLimitMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RateLimitMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [rate_limits.RateLimitNewResponseUnknown] or
// [shared.UnionString].
type RateLimitNewResponseUnion interface {
	ImplementsRateLimitsRateLimitNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RateLimitDeleteResponse struct {
	// The unique identifier of the rate limit.
	ID   string                      `json:"id"`
	JSON rateLimitDeleteResponseJSON `json:"-"`
}

// rateLimitDeleteResponseJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponse]
type rateLimitDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [rate_limits.RateLimitEditResponseUnknown] or
// [shared.UnionString].
type RateLimitEditResponseUnion interface {
	ImplementsRateLimitsRateLimitEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [rate_limits.RateLimitGetResponseUnknown] or
// [shared.UnionString].
type RateLimitGetResponseUnion interface {
	ImplementsRateLimitsRateLimitGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RateLimitNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r RateLimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RateLimitNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   RateLimitNewResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success RateLimitNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitNewResponseEnvelopeJSON    `json:"-"`
}

// rateLimitNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitNewResponseEnvelope]
type rateLimitNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitNewResponseEnvelopeSuccess bool

const (
	RateLimitNewResponseEnvelopeSuccessTrue RateLimitNewResponseEnvelopeSuccess = true
)

func (r RateLimitNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RateLimitListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RateLimitListParams]'s query parameters as `url.Values`.
func (r RateLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RateLimitDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   RateLimitDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success RateLimitDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitDeleteResponseEnvelopeJSON    `json:"-"`
}

// rateLimitDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponseEnvelope]
type rateLimitDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitDeleteResponseEnvelopeSuccess bool

const (
	RateLimitDeleteResponseEnvelopeSuccessTrue RateLimitDeleteResponseEnvelopeSuccess = true
)

func (r RateLimitDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RateLimitEditParams struct {
	Body interface{} `json:"body,required"`
}

func (r RateLimitEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RateLimitEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   RateLimitEditResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success RateLimitEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitEditResponseEnvelopeJSON    `json:"-"`
}

// rateLimitEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitEditResponseEnvelope]
type rateLimitEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitEditResponseEnvelopeSuccess bool

const (
	RateLimitEditResponseEnvelopeSuccessTrue RateLimitEditResponseEnvelopeSuccess = true
)

func (r RateLimitEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RateLimitGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   RateLimitGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success RateLimitGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitGetResponseEnvelopeJSON    `json:"-"`
}

// rateLimitGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitGetResponseEnvelope]
type rateLimitGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RateLimitGetResponseEnvelopeSuccess bool

const (
	RateLimitGetResponseEnvelopeSuccessTrue RateLimitGetResponseEnvelopeSuccess = true
)

func (r RateLimitGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
