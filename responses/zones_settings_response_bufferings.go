package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type ResponseBufferingListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ResponseBuffering `json:"result"`
	JSON   ResponseBufferingListResponseJSON
}

type ResponseBufferingListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ResponseBufferingListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ResponseBufferingListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ResponseBufferingUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ResponseBuffering `json:"result"`
	JSON   ResponseBufferingUpdateResponseJSON
}

type ResponseBufferingUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ResponseBufferingUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ResponseBufferingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
