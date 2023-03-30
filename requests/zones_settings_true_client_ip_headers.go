package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type TrueClientIpHeaderListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result fields.Field[TrueClientIpHeader] `json:"result"`
}

// MarshalJSON serializes TrueClientIpHeaderListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TrueClientIpHeaderListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TrueClientIpHeaderListResponse) String() (result string) {
	return fmt.Sprintf("&TrueClientIpHeaderListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type TrueClientIpHeaderUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result fields.Field[TrueClientIpHeader] `json:"result"`
}

// MarshalJSON serializes TrueClientIpHeaderUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TrueClientIpHeaderUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TrueClientIpHeaderUpdateResponse) String() (result string) {
	return fmt.Sprintf("&TrueClientIpHeaderUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type TrueClientIpHeaderUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[TrueClientIpHeaderValue] `json:"value,required"`
}

// MarshalJSON serializes TrueClientIpHeaderUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *TrueClientIpHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TrueClientIpHeaderUpdateParams) String() (result string) {
	return fmt.Sprintf("&TrueClientIpHeaderUpdateParams{Value:%s}", r.Value)
}
