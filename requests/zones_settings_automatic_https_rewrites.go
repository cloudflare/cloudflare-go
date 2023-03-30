package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type AutomaticHTTPsRewriteListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result fields.Field[AutomaticHTTPsRewrites] `json:"result"`
}

// MarshalJSON serializes AutomaticHTTPsRewriteListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AutomaticHTTPsRewriteListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticHTTPsRewriteListResponse) String() (result string) {
	return fmt.Sprintf("&AutomaticHTTPsRewriteListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AutomaticHTTPsRewriteUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result fields.Field[AutomaticHTTPsRewrites] `json:"result"`
}

// MarshalJSON serializes AutomaticHTTPsRewriteUpdateResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *AutomaticHTTPsRewriteUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticHTTPsRewriteUpdateResponse) String() (result string) {
	return fmt.Sprintf("&AutomaticHTTPsRewriteUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AutomaticHTTPsRewriteUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[AutomaticHTTPsRewritesValue] `json:"value,required"`
}

// MarshalJSON serializes AutomaticHTTPsRewriteUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AutomaticHTTPsRewriteUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticHTTPsRewriteUpdateParams) String() (result string) {
	return fmt.Sprintf("&AutomaticHTTPsRewriteUpdateParams{Value:%s}", r.Value)
}
