package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type OpportunisticOnionListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result fields.Field[OpportunisticOnion] `json:"result"`
}

// MarshalJSON serializes OpportunisticOnionListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *OpportunisticOnionListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticOnionListResponse) String() (result string) {
	return fmt.Sprintf("&OpportunisticOnionListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OpportunisticOnionUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result fields.Field[OpportunisticOnion] `json:"result"`
}

// MarshalJSON serializes OpportunisticOnionUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *OpportunisticOnionUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticOnionUpdateResponse) String() (result string) {
	return fmt.Sprintf("&OpportunisticOnionUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OpportunisticOnionUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[OpportunisticOnionValue] `json:"value,required"`
}

// MarshalJSON serializes OpportunisticOnionUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *OpportunisticOnionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticOnionUpdateParams) String() (result string) {
	return fmt.Sprintf("&OpportunisticOnionUpdateParams{Value:%s}", r.Value)
}
