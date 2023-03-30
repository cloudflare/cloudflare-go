package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type Ipv6ListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result fields.Field[Ipv6] `json:"result"`
}

// MarshalJSON serializes Ipv6ListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Ipv6ListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Ipv6ListResponse) String() (result string) {
	return fmt.Sprintf("&Ipv6ListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type Ipv6UpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable IPv6 on all subdomains that are Cloudflare enabled.
	// (https://support.cloudflare.com/hc/en-us/articles/200168586).
	Result fields.Field[Ipv6] `json:"result"`
}

// MarshalJSON serializes Ipv6UpdateResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Ipv6UpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Ipv6UpdateResponse) String() (result string) {
	return fmt.Sprintf("&Ipv6UpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type Ipv6UpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[Ipv6Value] `json:"value,required"`
}

// MarshalJSON serializes Ipv6UpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Ipv6UpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Ipv6UpdateParams) String() (result string) {
	return fmt.Sprintf("&Ipv6UpdateParams{Value:%s}", r.Value)
}
