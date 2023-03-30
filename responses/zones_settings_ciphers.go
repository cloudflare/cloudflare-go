package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type CipherListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result Ciphers `json:"result"`
	JSON   CipherListResponseJSON
}

type CipherListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CipherListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CipherListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CipherUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result Ciphers `json:"result"`
	JSON   CipherUpdateResponseJSON
}

type CipherUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CipherUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CipherUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
