package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type MinTlsVersionListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result fields.Field[MinTlsVersion] `json:"result"`
}

// MarshalJSON serializes MinTlsVersionListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MinTlsVersionListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinTlsVersionListResponse) String() (result string) {
	return fmt.Sprintf("&MinTlsVersionListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MinTlsVersionUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result fields.Field[MinTlsVersion] `json:"result"`
}

// MarshalJSON serializes MinTlsVersionUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MinTlsVersionUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinTlsVersionUpdateResponse) String() (result string) {
	return fmt.Sprintf("&MinTlsVersionUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MinTlsVersionUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[MinTlsVersionValue] `json:"value,required"`
}

// MarshalJSON serializes MinTlsVersionUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MinTlsVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinTlsVersionUpdateParams) String() (result string) {
	return fmt.Sprintf("&MinTlsVersionUpdateParams{Value:%s}", r.Value)
}
