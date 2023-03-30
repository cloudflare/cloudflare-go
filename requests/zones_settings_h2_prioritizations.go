package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type H2PrioritizationListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result fields.Field[H2Prioritization] `json:"result"`
}

// MarshalJSON serializes H2PrioritizationListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *H2PrioritizationListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r H2PrioritizationListResponse) String() (result string) {
	return fmt.Sprintf("&H2PrioritizationListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type H2PrioritizationUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result fields.Field[H2Prioritization] `json:"result"`
}

// MarshalJSON serializes H2PrioritizationUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *H2PrioritizationUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r H2PrioritizationUpdateResponse) String() (result string) {
	return fmt.Sprintf("&H2PrioritizationUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type H2PrioritizationUpdateParams struct {
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value fields.Field[H2Prioritization] `json:"value,required"`
}

// MarshalJSON serializes H2PrioritizationUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *H2PrioritizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r H2PrioritizationUpdateParams) String() (result string) {
	return fmt.Sprintf("&H2PrioritizationUpdateParams{Value:%s}", r.Value)
}
