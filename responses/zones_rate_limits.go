package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type RatelimitCollection struct {
	Errors   []Messages           `json:"errors"`
	Messages []Messages           `json:"messages"`
	Result   []Ratelimit_877kvwJo `json:"result"`
	// Whether the API call was successful
	Success    RatelimitCollectionSuccess `json:"success"`
	ResultInfo ResultInfo                 `json:"result_info"`
	JSON       RatelimitCollectionJSON
}

type RatelimitCollectionJSON struct {
	Errors     pjson.Metadata
	Messages   pjson.Metadata
	Result     pjson.Metadata
	Success    pjson.Metadata
	ResultInfo pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RatelimitCollection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RatelimitCollection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Ratelimit_877kvwJo struct {
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action SchemasAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []Bypass `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match Match `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64 `json:"threshold"`
	JSON      Ratelimit_877kvwJoJSON
}

type Ratelimit_877kvwJoJSON struct {
	Action      pjson.Metadata
	Bypass      pjson.Metadata
	Description pjson.Metadata
	Disabled    pjson.Metadata
	ID          pjson.Metadata
	Match       pjson.Metadata
	Period      pjson.Metadata
	Threshold   pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ratelimit_877kvwJo using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Ratelimit_877kvwJo) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SchemasAction struct {
	// The action to perform.
	Mode Mode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response CustomResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64 `json:"timeout"`
	JSON    SchemasActionJSON
}

type SchemasActionJSON struct {
	Mode     pjson.Metadata
	Response pjson.Metadata
	Timeout  pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SchemasAction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *SchemasAction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string `json:"content_type"`
	JSON        CustomResponseJSON
}

type CustomResponseJSON struct {
	Body        pjson.Metadata
	ContentType pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CustomResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CustomResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Bypass struct {
	Name BypassName `json:"name"`
	// The URL to bypass.
	Value string `json:"value"`
	JSON  BypassJSON
}

type BypassJSON struct {
	Name   pjson.Metadata
	Value  pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Bypass using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Bypass) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BypassName string

const (
	BypassNameURL BypassName = "url"
)

type Match struct {
	Headers  []MatchHeaders `json:"headers"`
	Request  MatchRequest   `json:"request"`
	Response MatchResponse  `json:"response"`
	JSON     MatchJSON
}

type MatchJSON struct {
	Headers  pjson.Metadata
	Request  pjson.Metadata
	Response pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Match using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Match) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MatchHeaders struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op HeaderOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string `json:"value"`
	JSON  MatchHeadersJSON
}

type MatchHeadersJSON struct {
	Name   pjson.Metadata
	Op     pjson.Metadata
	Value  pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MatchHeaders using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MatchHeaders) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	Methods []Methods `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string `json:"url"`
	JSON MatchRequestJSON
}

type MatchRequestJSON struct {
	Methods pjson.Metadata
	Schemes pjson.Metadata
	URL     pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MatchRequest using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MatchRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	OriginTraffic bool `json:"origin_traffic"`
	JSON          MatchResponseJSON
}

type MatchResponseJSON struct {
	OriginTraffic pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MatchResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MatchResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RatelimitCollectionSuccess bool

const (
	RatelimitCollectionSuccessTrue RatelimitCollectionSuccess = true
)

type RatelimitSingle struct {
	Errors   []Messages             `json:"errors"`
	Messages []Messages             `json:"messages"`
	Result   map[string]interface{} `json:"result"`
	// Whether the API call was successful
	Success RatelimitSingleSuccess `json:"success"`
	JSON    RatelimitSingleJSON
}

type RatelimitSingleJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Result   pjson.Metadata
	Success  pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RatelimitSingle using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RatelimitSingle) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RatelimitSingleSuccess bool

const (
	RatelimitSingleSuccessTrue RatelimitSingleSuccess = true
)

type RateLimitDeleteResponse struct {
	Errors   []Messages                    `json:"errors"`
	Messages []Messages                    `json:"messages"`
	Result   RateLimitDeleteResponseResult `json:"result"`
	// Whether the API call was successful
	Success RateLimitDeleteResponseSuccess `json:"success"`
	JSON    RateLimitDeleteResponseJSON
}

type RateLimitDeleteResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Result   pjson.Metadata
	Success  pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RateLimitDeleteResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RateLimitDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RateLimitDeleteResponseResult struct {
	// The unique identifier of the rate limit.
	ID   string `json:"id"`
	JSON RateLimitDeleteResponseResultJSON
}

type RateLimitDeleteResponseResultJSON struct {
	ID     pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RateLimitDeleteResponseResult
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RateLimitDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RateLimitDeleteResponseSuccess bool

const (
	RateLimitDeleteResponseSuccessTrue RateLimitDeleteResponseSuccess = true
)

type RateLimitsGetParams struct {
	// The {zone_identifier} parameter of /zones/{zone_identifier}/rate_limits/{id}
	ZoneIdentifier string `path:"zone_identifier" json:"-"`
}

type RateLimitsDeleteParams struct {
	// The {zone_identifier} parameter of /zones/{zone_identifier}/rate_limits/{id}
	ZoneIdentifier string `path:"zone_identifier" json:"-"`
}
