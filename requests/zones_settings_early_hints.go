package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type EarlyHintListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result fields.Field[EarlyHints] `json:"result"`
}

// MarshalJSON serializes EarlyHintListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EarlyHintListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EarlyHintListResponse) String() (result string) {
	return fmt.Sprintf("&EarlyHintListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type EarlyHintUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result fields.Field[EarlyHints] `json:"result"`
}

// MarshalJSON serializes EarlyHintUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EarlyHintUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EarlyHintUpdateResponse) String() (result string) {
	return fmt.Sprintf("&EarlyHintUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type EarlyHintUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[EarlyHintsValue] `json:"value,required"`
}

// MarshalJSON serializes EarlyHintUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EarlyHintUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EarlyHintUpdateParams) String() (result string) {
	return fmt.Sprintf("&EarlyHintUpdateParams{Value:%s}", r.Value)
}
