package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type ProxyReadTimeoutListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Maximum time between two read operations from origin.
	Result fields.Field[ProxyReadTimeout] `json:"result"`
}

// MarshalJSON serializes ProxyReadTimeoutListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ProxyReadTimeoutListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ProxyReadTimeoutListResponse) String() (result string) {
	return fmt.Sprintf("&ProxyReadTimeoutListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ProxyReadTimeoutUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Maximum time between two read operations from origin.
	Result fields.Field[ProxyReadTimeout] `json:"result"`
}

// MarshalJSON serializes ProxyReadTimeoutUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ProxyReadTimeoutUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ProxyReadTimeoutUpdateResponse) String() (result string) {
	return fmt.Sprintf("&ProxyReadTimeoutUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ProxyReadTimeoutUpdateParams struct {
	// Maximum time between two read operations from origin.
	Value fields.Field[ProxyReadTimeout] `json:"value,required"`
}

// MarshalJSON serializes ProxyReadTimeoutUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ProxyReadTimeoutUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ProxyReadTimeoutUpdateParams) String() (result string) {
	return fmt.Sprintf("&ProxyReadTimeoutUpdateParams{Value:%s}", r.Value)
}
