package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type TlsClientAuthListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result TlsClientAuth `json:"result"`
	JSON   TlsClientAuthListResponseJSON
}

type TlsClientAuthListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TlsClientAuthListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TlsClientAuthListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TlsClientAuthUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result TlsClientAuth `json:"result"`
	JSON   TlsClientAuthUpdateResponseJSON
}

type TlsClientAuthUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TlsClientAuthUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *TlsClientAuthUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
