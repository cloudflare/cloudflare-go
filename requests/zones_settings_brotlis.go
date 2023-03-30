package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type BrotliListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result fields.Field[Brotli] `json:"result"`
}

// MarshalJSON serializes BrotliListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *BrotliListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrotliListResponse) String() (result string) {
	return fmt.Sprintf("&BrotliListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type BrotliUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result fields.Field[Brotli] `json:"result"`
}

// MarshalJSON serializes BrotliUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *BrotliUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrotliUpdateResponse) String() (result string) {
	return fmt.Sprintf("&BrotliUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type BrotliUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[BrotliValue] `json:"value,required"`
}

// MarshalJSON serializes BrotliUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *BrotliUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrotliUpdateParams) String() (result string) {
	return fmt.Sprintf("&BrotliUpdateParams{Value:%s}", r.Value)
}
