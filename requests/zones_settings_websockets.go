package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type WebsocketListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result fields.Field[Websockets] `json:"result"`
}

// MarshalJSON serializes WebsocketListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WebsocketListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WebsocketListResponse) String() (result string) {
	return fmt.Sprintf("&WebsocketListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type WebsocketUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result fields.Field[Websockets] `json:"result"`
}

// MarshalJSON serializes WebsocketUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WebsocketUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WebsocketUpdateResponse) String() (result string) {
	return fmt.Sprintf("&WebsocketUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type WebsocketUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[WebsocketsValue] `json:"value,required"`
}

// MarshalJSON serializes WebsocketUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *WebsocketUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WebsocketUpdateParams) String() (result string) {
	return fmt.Sprintf("&WebsocketUpdateParams{Value:%s}", r.Value)
}
