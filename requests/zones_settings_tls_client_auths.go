package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type TlsClientAuthListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result fields.Field[TlsClientAuth] `json:"result"`
}

// MarshalJSON serializes TlsClientAuthListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TlsClientAuthListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TlsClientAuthListResponse) String() (result string) {
	return fmt.Sprintf("&TlsClientAuthListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type TlsClientAuthUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result fields.Field[TlsClientAuth] `json:"result"`
}

// MarshalJSON serializes TlsClientAuthUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TlsClientAuthUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TlsClientAuthUpdateResponse) String() (result string) {
	return fmt.Sprintf("&TlsClientAuthUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type TlsClientAuthUpdateParams struct {
	// value of the zone setting.
	Value fields.Field[TlsClientAuthValue] `json:"value,required"`
}

// MarshalJSON serializes TlsClientAuthUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *TlsClientAuthUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TlsClientAuthUpdateParams) String() (result string) {
	return fmt.Sprintf("&TlsClientAuthUpdateParams{Value:%s}", r.Value)
}
