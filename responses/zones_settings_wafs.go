package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type WafListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
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
	Result Waf `json:"result"`
	JSON   WafListResponseJSON
}

type WafListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WafListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WafListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WafUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
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
	Result Waf `json:"result"`
	JSON   WafUpdateResponseJSON
}

type WafUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WafUpdateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WafUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
