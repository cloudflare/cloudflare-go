package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type BrowserCheckListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result BrowserCheck `json:"result"`
	JSON   BrowserCheckListResponseJSON
}

type BrowserCheckListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrowserCheckListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BrowserCheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BrowserCheckUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result BrowserCheck `json:"result"`
	JSON   BrowserCheckUpdateResponseJSON
}

type BrowserCheckUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrowserCheckUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BrowserCheckUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
