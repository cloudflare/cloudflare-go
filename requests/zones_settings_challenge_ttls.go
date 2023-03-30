package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type ChallengeTtlListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result fields.Field[ChallengeTtl] `json:"result"`
}

// MarshalJSON serializes ChallengeTtlListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ChallengeTtlListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ChallengeTtlListResponse) String() (result string) {
	return fmt.Sprintf("&ChallengeTtlListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ChallengeTtlUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result fields.Field[ChallengeTtl] `json:"result"`
}

// MarshalJSON serializes ChallengeTtlUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ChallengeTtlUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ChallengeTtlUpdateResponse) String() (result string) {
	return fmt.Sprintf("&ChallengeTtlUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ChallengeTtlUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[ChallengeTtlValue] `json:"value,required"`
}

// MarshalJSON serializes ChallengeTtlUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ChallengeTtlUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ChallengeTtlUpdateParams) String() (result string) {
	return fmt.Sprintf("&ChallengeTtlUpdateParams{Value:%s}", r.Value)
}
