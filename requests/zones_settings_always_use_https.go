package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type AlwaysUseHTTPListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result fields.Field[AlwaysUseHTTPs] `json:"result"`
}

// MarshalJSON serializes AlwaysUseHTTPListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AlwaysUseHTTPListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysUseHTTPListResponse) String() (result string) {
	return fmt.Sprintf("&AlwaysUseHTTPListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AlwaysUseHTTPUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result fields.Field[AlwaysUseHTTPs] `json:"result"`
}

// MarshalJSON serializes AlwaysUseHTTPUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AlwaysUseHTTPUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysUseHTTPUpdateResponse) String() (result string) {
	return fmt.Sprintf("&AlwaysUseHTTPUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AlwaysUseHTTPUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[AlwaysUseHTTPsValue] `json:"value,required"`
}

// MarshalJSON serializes AlwaysUseHTTPUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AlwaysUseHTTPUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysUseHTTPUpdateParams) String() (result string) {
	return fmt.Sprintf("&AlwaysUseHTTPUpdateParams{Value:%s}", r.Value)
}
