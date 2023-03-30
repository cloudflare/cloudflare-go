package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type EmailObfuscationListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result fields.Field[EmailObfuscation] `json:"result"`
}

// MarshalJSON serializes EmailObfuscationListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EmailObfuscationListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EmailObfuscationListResponse) String() (result string) {
	return fmt.Sprintf("&EmailObfuscationListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type EmailObfuscationUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result fields.Field[EmailObfuscation] `json:"result"`
}

// MarshalJSON serializes EmailObfuscationUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *EmailObfuscationUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EmailObfuscationUpdateResponse) String() (result string) {
	return fmt.Sprintf("&EmailObfuscationUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type EmailObfuscationUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[EmailObfuscationValue] `json:"value,required"`
}

// MarshalJSON serializes EmailObfuscationUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *EmailObfuscationUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EmailObfuscationUpdateParams) String() (result string) {
	return fmt.Sprintf("&EmailObfuscationUpdateParams{Value:%s}", r.Value)
}
