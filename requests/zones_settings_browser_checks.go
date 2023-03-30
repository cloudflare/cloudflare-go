package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type BrowserCheckListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result fields.Field[BrowserCheck] `json:"result"`
}

// MarshalJSON serializes BrowserCheckListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *BrowserCheckListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCheckListResponse) String() (result string) {
	return fmt.Sprintf("&BrowserCheckListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type BrowserCheckUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result fields.Field[BrowserCheck] `json:"result"`
}

// MarshalJSON serializes BrowserCheckUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *BrowserCheckUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCheckUpdateResponse) String() (result string) {
	return fmt.Sprintf("&BrowserCheckUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type BrowserCheckUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[BrowserCheckValue] `json:"value,required"`
}

// MarshalJSON serializes BrowserCheckUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *BrowserCheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCheckUpdateParams) String() (result string) {
	return fmt.Sprintf("&BrowserCheckUpdateParams{Value:%s}", r.Value)
}
