package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type PrivacyPassListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Privacy Pass is a browser extension developed by the Privacy Pass Team to
	// improve the browsing experience for your visitors. Enabling Privacy Pass will
	// reduce the number of CAPTCHAs shown to your visitors.
	// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
	Result fields.Field[PrivacyPass] `json:"result"`
}

// MarshalJSON serializes PrivacyPassListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PrivacyPassListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrivacyPassListResponse) String() (result string) {
	return fmt.Sprintf("&PrivacyPassListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PrivacyPassUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Privacy Pass is a browser extension developed by the Privacy Pass Team to
	// improve the browsing experience for your visitors. Enabling Privacy Pass will
	// reduce the number of CAPTCHAs shown to your visitors.
	// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
	Result fields.Field[PrivacyPass] `json:"result"`
}

// MarshalJSON serializes PrivacyPassUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PrivacyPassUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrivacyPassUpdateResponse) String() (result string) {
	return fmt.Sprintf("&PrivacyPassUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PrivacyPassUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[PrivacyPassValue] `json:"value,required"`
}

// MarshalJSON serializes PrivacyPassUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PrivacyPassUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrivacyPassUpdateParams) String() (result string) {
	return fmt.Sprintf("&PrivacyPassUpdateParams{Value:%s}", r.Value)
}
