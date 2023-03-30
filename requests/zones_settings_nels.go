package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type NelListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result fields.Field[Nel] `json:"result"`
}

// MarshalJSON serializes NelListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *NelListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r NelListResponse) String() (result string) {
	return fmt.Sprintf("&NelListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type NelUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result fields.Field[Nel] `json:"result"`
}

// MarshalJSON serializes NelUpdateResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *NelUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r NelUpdateResponse) String() (result string) {
	return fmt.Sprintf("&NelUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type NelUpdateParams struct {
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value fields.Field[Nel] `json:"value,required"`
}

// MarshalJSON serializes NelUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *NelUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r NelUpdateParams) String() (result string) {
	return fmt.Sprintf("&NelUpdateParams{Value:%s}", r.Value)
}
