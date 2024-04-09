// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingWebsocketService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingWebsocketService] method
// instead.
type SettingWebsocketService struct {
	Options []option.RequestOption
}

// NewSettingWebsocketService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingWebsocketService(opts ...option.RequestOption) (r *SettingWebsocketService) {
	r = &SettingWebsocketService{}
	r.Options = opts
	return
}

// Changes Websockets setting. For more information about Websockets, please refer
// to
// [Using Cloudflare with WebSockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Using-Cloudflare-with-WebSockets).
func (r *SettingWebsocketService) Edit(ctx context.Context, params SettingWebsocketEditParams, opts ...option.RequestOption) (res *Websocket, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebsocketEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/websockets", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Websockets setting. For more information about Websockets, please refer to
// [Using Cloudflare with WebSockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Using-Cloudflare-with-WebSockets).
func (r *SettingWebsocketService) Get(ctx context.Context, query SettingWebsocketGetParams, opts ...option.RequestOption) (res *Websocket, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebsocketGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/websockets", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type Websocket struct {
	// ID of the zone setting.
	ID WebsocketID `json:"id,required"`
	// Current value of the zone setting.
	Value WebsocketValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebsocketEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       websocketJSON `json:"-"`
}

// websocketJSON contains the JSON metadata for the struct [Websocket]
type websocketJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Websocket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r websocketJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type WebsocketID string

const (
	WebsocketIDWebsockets WebsocketID = "websockets"
)

func (r WebsocketID) IsKnown() bool {
	switch r {
	case WebsocketIDWebsockets:
		return true
	}
	return false
}

// Current value of the zone setting.
type WebsocketValue string

const (
	WebsocketValueOff WebsocketValue = "off"
	WebsocketValueOn  WebsocketValue = "on"
)

func (r WebsocketValue) IsKnown() bool {
	switch r {
	case WebsocketValueOff, WebsocketValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type WebsocketEditable bool

const (
	WebsocketEditableTrue  WebsocketEditable = true
	WebsocketEditableFalse WebsocketEditable = false
)

func (r WebsocketEditable) IsKnown() bool {
	switch r {
	case WebsocketEditableTrue, WebsocketEditableFalse:
		return true
	}
	return false
}

type SettingWebsocketEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingWebsocketEditParamsValue] `json:"value,required"`
}

func (r SettingWebsocketEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingWebsocketEditParamsValue string

const (
	SettingWebsocketEditParamsValueOff SettingWebsocketEditParamsValue = "off"
	SettingWebsocketEditParamsValueOn  SettingWebsocketEditParamsValue = "on"
)

func (r SettingWebsocketEditParamsValue) IsKnown() bool {
	switch r {
	case SettingWebsocketEditParamsValueOff, SettingWebsocketEditParamsValueOn:
		return true
	}
	return false
}

type SettingWebsocketEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result Websocket                                `json:"result"`
	JSON   settingWebsocketEditResponseEnvelopeJSON `json:"-"`
}

// settingWebsocketEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingWebsocketEditResponseEnvelope]
type settingWebsocketEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebsocketEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingWebsocketGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingWebsocketGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result Websocket                               `json:"result"`
	JSON   settingWebsocketGetResponseEnvelopeJSON `json:"-"`
}

// settingWebsocketGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingWebsocketGetResponseEnvelope]
type settingWebsocketGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebsocketGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
