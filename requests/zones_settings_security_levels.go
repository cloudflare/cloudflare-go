package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type SecurityLevelListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result fields.Field[SecurityLevel] `json:"result"`
}

// MarshalJSON serializes SecurityLevelListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SecurityLevelListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityLevelListResponse) String() (result string) {
	return fmt.Sprintf("&SecurityLevelListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SecurityLevelUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result fields.Field[SecurityLevel] `json:"result"`
}

// MarshalJSON serializes SecurityLevelUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SecurityLevelUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityLevelUpdateResponse) String() (result string) {
	return fmt.Sprintf("&SecurityLevelUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SecurityLevelUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[SecurityLevelValue] `json:"value,required"`
}

// MarshalJSON serializes SecurityLevelUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SecurityLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityLevelUpdateParams) String() (result string) {
	return fmt.Sprintf("&SecurityLevelUpdateParams{Value:%s}", r.Value)
}
