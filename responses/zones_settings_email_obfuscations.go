package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type EmailObfuscationListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result EmailObfuscation `json:"result"`
	JSON   EmailObfuscationListResponseJSON
}

type EmailObfuscationListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EmailObfuscationListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *EmailObfuscationListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EmailObfuscationUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result EmailObfuscation `json:"result"`
	JSON   EmailObfuscationUpdateResponseJSON
}

type EmailObfuscationUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// EmailObfuscationUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *EmailObfuscationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
