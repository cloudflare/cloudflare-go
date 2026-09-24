// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limits

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// RateLimitService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRateLimitService] method instead.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
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

// **Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API
// instead.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) New(ctx context.Context, params RateLimitNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/rate_limits", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// **Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API
// instead.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) List(ctx context.Context, params RateLimitListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/rate_limits", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return err
}

// **Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API
// instead.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Delete(ctx context.Context, rateLimitID string, body RateLimitDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", body.ZoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API
// instead.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Edit(ctx context.Context, rateLimitID string, params RateLimitEditParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", params.ZoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// **Deprecated**: This endpoint returns 410 Gone. Please use the Rulesets API
// instead.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Get(ctx context.Context, rateLimitID string, query RateLimitGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", query.ZoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
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

type RateLimitNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RateLimitNewParamsAction] `json:"action" api:"required"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match param.Field[RateLimitNewParamsMatch] `json:"match" api:"required"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period param.Field[float64] `json:"period" api:"required"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold param.Field[float64] `json:"threshold" api:"required"`
}

func (r RateLimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitNewParamsAction struct {
	// The action to perform.
	Mode param.Field[RateLimitNewParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RateLimitNewParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RateLimitNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RateLimitNewParamsActionMode string

const (
	RateLimitNewParamsActionModeSimulate         RateLimitNewParamsActionMode = "simulate"
	RateLimitNewParamsActionModeBan              RateLimitNewParamsActionMode = "ban"
	RateLimitNewParamsActionModeChallenge        RateLimitNewParamsActionMode = "challenge"
	RateLimitNewParamsActionModeJSChallenge      RateLimitNewParamsActionMode = "js_challenge"
	RateLimitNewParamsActionModeManagedChallenge RateLimitNewParamsActionMode = "managed_challenge"
)

func (r RateLimitNewParamsActionMode) IsKnown() bool {
	switch r {
	case RateLimitNewParamsActionModeSimulate, RateLimitNewParamsActionModeBan, RateLimitNewParamsActionModeChallenge, RateLimitNewParamsActionModeJSChallenge, RateLimitNewParamsActionModeManagedChallenge:
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
type RateLimitNewParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RateLimitNewParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitNewParamsMatch struct {
	Headers  param.Field[[]RateLimitNewParamsMatchHeader] `json:"headers"`
	Request  param.Field[RateLimitNewParamsMatchRequest]  `json:"request"`
	Response param.Field[RateLimitNewParamsMatchResponse] `json:"response"`
}

func (r RateLimitNewParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitNewParamsMatchHeader struct {
	// The name of the response header to match.
	Name param.Field[string] `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op param.Field[RateLimitNewParamsMatchHeadersOp] `json:"op"`
	// The value of the response header, which must match exactly.
	Value param.Field[string] `json:"value"`
}

func (r RateLimitNewParamsMatchHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitNewParamsMatchHeadersOp string

const (
	RateLimitNewParamsMatchHeadersOpEq RateLimitNewParamsMatchHeadersOp = "eq"
	RateLimitNewParamsMatchHeadersOpNe RateLimitNewParamsMatchHeadersOp = "ne"
)

func (r RateLimitNewParamsMatchHeadersOp) IsKnown() bool {
	switch r {
	case RateLimitNewParamsMatchHeadersOpEq, RateLimitNewParamsMatchHeadersOpNe:
		return true
	}
	return false
}

type RateLimitNewParamsMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods param.Field[[]RateLimitNewParamsMatchRequestMethod] `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes param.Field[[]string] `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL param.Field[string] `json:"url"`
}

func (r RateLimitNewParamsMatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitNewParamsMatchRequestMethod string

const (
	RateLimitNewParamsMatchRequestMethodGet    RateLimitNewParamsMatchRequestMethod = "GET"
	RateLimitNewParamsMatchRequestMethodPost   RateLimitNewParamsMatchRequestMethod = "POST"
	RateLimitNewParamsMatchRequestMethodPut    RateLimitNewParamsMatchRequestMethod = "PUT"
	RateLimitNewParamsMatchRequestMethodDelete RateLimitNewParamsMatchRequestMethod = "DELETE"
	RateLimitNewParamsMatchRequestMethodPatch  RateLimitNewParamsMatchRequestMethod = "PATCH"
	RateLimitNewParamsMatchRequestMethodHead   RateLimitNewParamsMatchRequestMethod = "HEAD"
	RateLimitNewParamsMatchRequestMethod_All   RateLimitNewParamsMatchRequestMethod = "_ALL_"
)

func (r RateLimitNewParamsMatchRequestMethod) IsKnown() bool {
	switch r {
	case RateLimitNewParamsMatchRequestMethodGet, RateLimitNewParamsMatchRequestMethodPost, RateLimitNewParamsMatchRequestMethodPut, RateLimitNewParamsMatchRequestMethodDelete, RateLimitNewParamsMatchRequestMethodPatch, RateLimitNewParamsMatchRequestMethodHead, RateLimitNewParamsMatchRequestMethod_All:
		return true
	}
	return false
}

type RateLimitNewParamsMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic param.Field[bool] `json:"origin_traffic"`
}

func (r RateLimitNewParamsMatchResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Defines the page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Defines the maximum number of results per page. You can only set the value to
	// `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RateLimitListParams]'s query parameters as `url.Values`.
func (r RateLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RateLimitDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type RateLimitEditParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RateLimitEditParamsAction] `json:"action" api:"required"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match param.Field[RateLimitEditParamsMatch] `json:"match" api:"required"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period param.Field[float64] `json:"period" api:"required"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold param.Field[float64] `json:"threshold" api:"required"`
}

func (r RateLimitEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitEditParamsAction struct {
	// The action to perform.
	Mode param.Field[RateLimitEditParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RateLimitEditParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RateLimitEditParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RateLimitEditParamsActionMode string

const (
	RateLimitEditParamsActionModeSimulate         RateLimitEditParamsActionMode = "simulate"
	RateLimitEditParamsActionModeBan              RateLimitEditParamsActionMode = "ban"
	RateLimitEditParamsActionModeChallenge        RateLimitEditParamsActionMode = "challenge"
	RateLimitEditParamsActionModeJSChallenge      RateLimitEditParamsActionMode = "js_challenge"
	RateLimitEditParamsActionModeManagedChallenge RateLimitEditParamsActionMode = "managed_challenge"
)

func (r RateLimitEditParamsActionMode) IsKnown() bool {
	switch r {
	case RateLimitEditParamsActionModeSimulate, RateLimitEditParamsActionModeBan, RateLimitEditParamsActionModeChallenge, RateLimitEditParamsActionModeJSChallenge, RateLimitEditParamsActionModeManagedChallenge:
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
type RateLimitEditParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RateLimitEditParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitEditParamsMatch struct {
	Headers  param.Field[[]RateLimitEditParamsMatchHeader] `json:"headers"`
	Request  param.Field[RateLimitEditParamsMatchRequest]  `json:"request"`
	Response param.Field[RateLimitEditParamsMatchResponse] `json:"response"`
}

func (r RateLimitEditParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitEditParamsMatchHeader struct {
	// The name of the response header to match.
	Name param.Field[string] `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op param.Field[RateLimitEditParamsMatchHeadersOp] `json:"op"`
	// The value of the response header, which must match exactly.
	Value param.Field[string] `json:"value"`
}

func (r RateLimitEditParamsMatchHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitEditParamsMatchHeadersOp string

const (
	RateLimitEditParamsMatchHeadersOpEq RateLimitEditParamsMatchHeadersOp = "eq"
	RateLimitEditParamsMatchHeadersOpNe RateLimitEditParamsMatchHeadersOp = "ne"
)

func (r RateLimitEditParamsMatchHeadersOp) IsKnown() bool {
	switch r {
	case RateLimitEditParamsMatchHeadersOpEq, RateLimitEditParamsMatchHeadersOpNe:
		return true
	}
	return false
}

type RateLimitEditParamsMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods param.Field[[]RateLimitEditParamsMatchRequestMethod] `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes param.Field[[]string] `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL param.Field[string] `json:"url"`
}

func (r RateLimitEditParamsMatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitEditParamsMatchRequestMethod string

const (
	RateLimitEditParamsMatchRequestMethodGet    RateLimitEditParamsMatchRequestMethod = "GET"
	RateLimitEditParamsMatchRequestMethodPost   RateLimitEditParamsMatchRequestMethod = "POST"
	RateLimitEditParamsMatchRequestMethodPut    RateLimitEditParamsMatchRequestMethod = "PUT"
	RateLimitEditParamsMatchRequestMethodDelete RateLimitEditParamsMatchRequestMethod = "DELETE"
	RateLimitEditParamsMatchRequestMethodPatch  RateLimitEditParamsMatchRequestMethod = "PATCH"
	RateLimitEditParamsMatchRequestMethodHead   RateLimitEditParamsMatchRequestMethod = "HEAD"
	RateLimitEditParamsMatchRequestMethod_All   RateLimitEditParamsMatchRequestMethod = "_ALL_"
)

func (r RateLimitEditParamsMatchRequestMethod) IsKnown() bool {
	switch r {
	case RateLimitEditParamsMatchRequestMethodGet, RateLimitEditParamsMatchRequestMethodPost, RateLimitEditParamsMatchRequestMethodPut, RateLimitEditParamsMatchRequestMethodDelete, RateLimitEditParamsMatchRequestMethodPatch, RateLimitEditParamsMatchRequestMethodHead, RateLimitEditParamsMatchRequestMethod_All:
		return true
	}
	return false
}

type RateLimitEditParamsMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic param.Field[bool] `json:"origin_traffic"`
}

func (r RateLimitEditParamsMatchResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}
