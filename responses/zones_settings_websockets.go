package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type WebsocketListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result Websockets `json:"result"`
	JSON   WebsocketListResponseJSON
}

type WebsocketListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WebsocketListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WebsocketListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WebsocketUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result Websockets `json:"result"`
	JSON   WebsocketUpdateResponseJSON
}

type WebsocketUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WebsocketUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WebsocketUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
