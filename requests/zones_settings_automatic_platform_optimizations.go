package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type AutomaticPlatformOptimizationListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool]                          `json:"success"`
	Result  fields.Field[AutomaticPlatformOptimization] `json:"result"`
}

// MarshalJSON serializes AutomaticPlatformOptimizationListResponse into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *AutomaticPlatformOptimizationListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticPlatformOptimizationListResponse) String() (result string) {
	return fmt.Sprintf("&AutomaticPlatformOptimizationListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AutomaticPlatformOptimizationUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool]                          `json:"success"`
	Result  fields.Field[AutomaticPlatformOptimization] `json:"result"`
}

// MarshalJSON serializes AutomaticPlatformOptimizationUpdateResponse into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *AutomaticPlatformOptimizationUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticPlatformOptimizationUpdateResponse) String() (result string) {
	return fmt.Sprintf("&AutomaticPlatformOptimizationUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AutomaticPlatformOptimizationUpdateParams struct {
	Value fields.Field[AutomaticPlatformOptimization] `json:"value,required"`
}

// MarshalJSON serializes AutomaticPlatformOptimizationUpdateParams into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *AutomaticPlatformOptimizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticPlatformOptimizationUpdateParams) String() (result string) {
	return fmt.Sprintf("&AutomaticPlatformOptimizationUpdateParams{Value:%s}", r.Value)
}
