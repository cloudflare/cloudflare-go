package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type Ipv6ListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result Ipv6 `json:"result"`
	JSON   Ipv6ListResponseJSON
}

type Ipv6ListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ipv6ListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Ipv6ListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Ipv6UpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result Ipv6 `json:"result"`
	JSON   Ipv6UpdateResponseJSON
}

type Ipv6UpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ipv6UpdateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Ipv6UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
