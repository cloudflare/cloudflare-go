package requests

import (
	"fmt"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/core/query"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type RatelimitCollection struct {
	Errors   fields.Field[[]Messages]           `json:"errors"`
	Messages fields.Field[[]Messages]           `json:"messages"`
	Result   fields.Field[[]Ratelimit_877kvwJo] `json:"result"`
	// Whether the API call was successful
	Success    fields.Field[RatelimitCollectionSuccess] `json:"success"`
	ResultInfo fields.Field[ResultInfo]                 `json:"result_info"`
}

// MarshalJSON serializes RatelimitCollection into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *RatelimitCollection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RatelimitCollection) String() (result string) {
	return fmt.Sprintf("&RatelimitCollection{Errors:%s Messages:%s Result:%s Success:%s ResultInfo:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), core.Fmt(r.Result), r.Success, r.ResultInfo)
}

type Ratelimit_877kvwJo struct {
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action fields.Field[SchemasAction] `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass fields.Field[[]Bypass] `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description fields.Field[string] `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled fields.Field[bool] `json:"disabled"`
	// The unique identifier of the rate limit.
	ID fields.Field[string] `json:"id"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match fields.Field[Match] `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period fields.Field[float64] `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold fields.Field[float64] `json:"threshold"`
}

// MarshalJSON serializes Ratelimit_877kvwJo into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Ratelimit_877kvwJo) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Ratelimit_877kvwJo) String() (result string) {
	return fmt.Sprintf("&Ratelimit_877kvwJo{Action:%s Bypass:%s Description:%s Disabled:%s ID:%s Match:%s Period:%s Threshold:%s}", r.Action, core.Fmt(r.Bypass), r.Description, r.Disabled, r.ID, r.Match, r.Period, r.Threshold)
}

type SchemasAction struct {
	// The action to perform.
	Mode fields.Field[Mode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response fields.Field[CustomResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout fields.Field[float64] `json:"timeout"`
}

// MarshalJSON serializes SchemasAction into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SchemasAction) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SchemasAction) String() (result string) {
	return fmt.Sprintf("&SchemasAction{Mode:%s Response:%s Timeout:%s}", r.Mode, r.Response, r.Timeout)
}

type Mode string

const (
	ModeSimulate         Mode = "simulate"
	ModeBan              Mode = "ban"
	ModeChallenge        Mode = "challenge"
	ModeJsChallenge      Mode = "js_challenge"
	ModeManagedChallenge Mode = "managed_challenge"
)

type CustomResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body fields.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType fields.Field[string] `json:"content_type"`
}

// MarshalJSON serializes CustomResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CustomResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CustomResponse) String() (result string) {
	return fmt.Sprintf("&CustomResponse{Body:%s ContentType:%s}", r.Body, r.ContentType)
}

type Bypass struct {
	Name fields.Field[BypassName] `json:"name"`
	// The URL to bypass.
	Value fields.Field[string] `json:"value"`
}

// MarshalJSON serializes Bypass into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Bypass) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Bypass) String() (result string) {
	return fmt.Sprintf("&Bypass{Name:%s Value:%s}", r.Name, r.Value)
}

type BypassName string

const (
	BypassNameURL BypassName = "url"
)

type Match struct {
	Headers  fields.Field[[]MatchHeaders] `json:"headers"`
	Request  fields.Field[MatchRequest]   `json:"request"`
	Response fields.Field[MatchResponse]  `json:"response"`
}

// MarshalJSON serializes Match into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Match) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Match) String() (result string) {
	return fmt.Sprintf("&Match{Headers:%s Request:%s Response:%s}", core.Fmt(r.Headers), r.Request, r.Response)
}

type MatchHeaders struct {
	// The name of the response header to match.
	Name fields.Field[string] `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op fields.Field[HeaderOp] `json:"op"`
	// The value of the response header, which must match exactly.
	Value fields.Field[string] `json:"value"`
}

// MarshalJSON serializes MatchHeaders into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MatchHeaders) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MatchHeaders) String() (result string) {
	return fmt.Sprintf("&MatchHeaders{Name:%s Op:%s Value:%s}", r.Name, r.Op, r.Value)
}

type HeaderOp string

const (
	HeaderOpEq HeaderOp = "eq"
	HeaderOpNe HeaderOp = "ne"
)

type MatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods fields.Field[[]Methods] `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes fields.Field[[]string] `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL fields.Field[string] `json:"url"`
}

// MarshalJSON serializes MatchRequest into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MatchRequest) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MatchRequest) String() (result string) {
	return fmt.Sprintf("&MatchRequest{Methods:%s Schemes:%s URL:%s}", core.Fmt(r.Methods), core.Fmt(r.Schemes), r.URL)
}

type Methods string

const (
	MethodsGet    Methods = "GET"
	MethodsPost   Methods = "POST"
	MethodsPut    Methods = "PUT"
	MethodsDelete Methods = "DELETE"
	MethodsPatch  Methods = "PATCH"
	MethodsHead   Methods = "HEAD"
	MethodsAll    Methods = "_ALL_"
)

type MatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic fields.Field[bool] `json:"origin_traffic"`
}

// MarshalJSON serializes MatchResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MatchResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MatchResponse) String() (result string) {
	return fmt.Sprintf("&MatchResponse{OriginTraffic:%s}", r.OriginTraffic)
}

type RatelimitCollectionSuccess bool

const (
	RatelimitCollectionSuccessTrue RatelimitCollectionSuccess = true
)

type RatelimitSingle struct {
	Errors   fields.Field[[]Messages]             `json:"errors"`
	Messages fields.Field[[]Messages]             `json:"messages"`
	Result   fields.Field[map[string]interface{}] `json:"result"`
	// Whether the API call was successful
	Success fields.Field[RatelimitSingleSuccess] `json:"success"`
}

// MarshalJSON serializes RatelimitSingle into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *RatelimitSingle) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RatelimitSingle) String() (result string) {
	return fmt.Sprintf("&RatelimitSingle{Errors:%s Messages:%s Result:%s Success:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Result, r.Success)
}

type RatelimitSingleSuccess bool

const (
	RatelimitSingleSuccessTrue RatelimitSingleSuccess = true
)

type RateLimitDeleteResponse struct {
	Errors   fields.Field[[]Messages]                    `json:"errors"`
	Messages fields.Field[[]Messages]                    `json:"messages"`
	Result   fields.Field[RateLimitDeleteResponseResult] `json:"result"`
	// Whether the API call was successful
	Success fields.Field[RateLimitDeleteResponseSuccess] `json:"success"`
}

// MarshalJSON serializes RateLimitDeleteResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *RateLimitDeleteResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RateLimitDeleteResponse) String() (result string) {
	return fmt.Sprintf("&RateLimitDeleteResponse{Errors:%s Messages:%s Result:%s Success:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Result, r.Success)
}

type RateLimitDeleteResponseResult struct {
	// The unique identifier of the rate limit.
	ID fields.Field[string] `json:"id"`
}

// MarshalJSON serializes RateLimitDeleteResponseResult into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *RateLimitDeleteResponseResult) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RateLimitDeleteResponseResult) String() (result string) {
	return fmt.Sprintf("&RateLimitDeleteResponseResult{ID:%s}", r.ID)
}

type RateLimitDeleteResponseSuccess bool

const (
	RateLimitDeleteResponseSuccessTrue RateLimitDeleteResponseSuccess = true
)

type RateLimitsGetParams struct {
	// The {zone_identifier} parameter of /zones/{zone_identifier}/rate_limits/{id}
	ZoneIdentifier string `path:"zone_identifier" json:"-"`
}

type RateLimitListParams struct {
	// The page number of paginated results.
	Page fields.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage fields.Field[float64] `query:"per_page"`
}

// URLQuery serializes RateLimitListParams into a url.Values of the query
// parameters associated with this value
func (r *RateLimitListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r RateLimitListParams) String() (result string) {
	return fmt.Sprintf("&RateLimitListParams{Page:%s PerPage:%s}", r.Page, r.PerPage)
}

type RateLimitsDeleteParams struct {
	// The {zone_identifier} parameter of /zones/{zone_identifier}/rate_limits/{id}
	ZoneIdentifier string `path:"zone_identifier" json:"-"`
}
