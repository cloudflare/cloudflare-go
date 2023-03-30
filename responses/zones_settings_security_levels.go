package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type SecurityLevelListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result SecurityLevel `json:"result"`
	JSON   SecurityLevelListResponseJSON
}

type SecurityLevelListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SecurityLevelListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *SecurityLevelListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SecurityLevelUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result SecurityLevel `json:"result"`
	JSON   SecurityLevelUpdateResponseJSON
}

type SecurityLevelUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SecurityLevelUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *SecurityLevelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
