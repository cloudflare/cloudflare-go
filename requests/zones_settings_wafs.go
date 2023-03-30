package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type WafListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// The WAF examines HTTP requests to your website. It inspects both GET and POST
	// requests and applies rules to help filter out illegitimate traffic from
	// legitimate website visitors. The Cloudflare WAF inspects website addresses or
	// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
	// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
	// page that asks them to submit a CAPTCHA successfully to continue their action.
	// If the challenge is failed, the action will be stopped. What this means is that
	// Cloudflare's WAF will block any traffic identified as illegitimate before it
	// reaches your origin web server.
	// (https://support.cloudflare.com/hc/en-us/articles/200172016).
	Result fields.Field[Waf] `json:"result"`
}

// MarshalJSON serializes WafListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WafListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WafListResponse) String() (result string) {
	return fmt.Sprintf("&WafListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type WafUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// The WAF examines HTTP requests to your website. It inspects both GET and POST
	// requests and applies rules to help filter out illegitimate traffic from
	// legitimate website visitors. The Cloudflare WAF inspects website addresses or
	// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
	// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
	// page that asks them to submit a CAPTCHA successfully to continue their action.
	// If the challenge is failed, the action will be stopped. What this means is that
	// Cloudflare's WAF will block any traffic identified as illegitimate before it
	// reaches your origin web server.
	// (https://support.cloudflare.com/hc/en-us/articles/200172016).
	Result fields.Field[Waf] `json:"result"`
}

// MarshalJSON serializes WafUpdateResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WafUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WafUpdateResponse) String() (result string) {
	return fmt.Sprintf("&WafUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type WafUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[WafValue] `json:"value,required"`
}

// MarshalJSON serializes WafUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WafUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WafUpdateParams) String() (result string) {
	return fmt.Sprintf("&WafUpdateParams{Value:%s}", r.Value)
}
