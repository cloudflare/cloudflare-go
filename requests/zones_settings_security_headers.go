package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type SecurityHeaderListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare security header for a zone.
	Result fields.Field[SecurityHeader] `json:"result"`
}

// MarshalJSON serializes SecurityHeaderListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SecurityHeaderListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityHeaderListResponse) String() (result string) {
	return fmt.Sprintf("&SecurityHeaderListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SecurityHeaderUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare security header for a zone.
	Result fields.Field[SecurityHeader] `json:"result"`
}

// MarshalJSON serializes SecurityHeaderUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SecurityHeaderUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityHeaderUpdateResponse) String() (result string) {
	return fmt.Sprintf("&SecurityHeaderUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SecurityHeaderUpdateParams struct {
	Value fields.Field[SecurityHeaderValue] `json:"value,required"`
}

// MarshalJSON serializes SecurityHeaderUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SecurityHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityHeaderUpdateParams) String() (result string) {
	return fmt.Sprintf("&SecurityHeaderUpdateParams{Value:%s}", r.Value)
}
