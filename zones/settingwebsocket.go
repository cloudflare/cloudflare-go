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
func (r *SettingWebsocketService) Edit(ctx context.Context, params SettingWebsocketEditParams, opts ...option.RequestOption) (res *ZonesWebsockets, err error) {
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
func (r *SettingWebsocketService) Get(ctx context.Context, query SettingWebsocketGetParams, opts ...option.RequestOption) (res *ZonesWebsockets, err error) {
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
type ZonesWebsockets struct {
	// ID of the zone setting.
	ID ZonesWebsocketsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesWebsocketsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesWebsocketsJSON `json:"-"`
}

// zonesWebsocketsJSON contains the JSON metadata for the struct [ZonesWebsockets]
type zonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesWebsocketsJSON) RawJSON() string {
	return r.raw
}

func (r ZonesWebsockets) implementsZonesSettingEditResponse() {}

func (r ZonesWebsockets) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesWebsocketsID string

const (
	ZonesWebsocketsIDWebsockets ZonesWebsocketsID = "websockets"
)

func (r ZonesWebsocketsID) IsKnown() bool {
	switch r {
	case ZonesWebsocketsIDWebsockets:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesWebsocketsValue string

const (
	ZonesWebsocketsValueOff ZonesWebsocketsValue = "off"
	ZonesWebsocketsValueOn  ZonesWebsocketsValue = "on"
)

func (r ZonesWebsocketsValue) IsKnown() bool {
	switch r {
	case ZonesWebsocketsValueOff, ZonesWebsocketsValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesWebsocketsEditable bool

const (
	ZonesWebsocketsEditableTrue  ZonesWebsocketsEditable = true
	ZonesWebsocketsEditableFalse ZonesWebsocketsEditable = false
)

func (r ZonesWebsocketsEditable) IsKnown() bool {
	switch r {
	case ZonesWebsocketsEditableTrue, ZonesWebsocketsEditableFalse:
		return true
	}
	return false
}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZonesWebsocketsParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesWebsocketsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesWebsocketsValue] `json:"value,required"`
}

func (r ZonesWebsocketsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesWebsocketsParam) implementsZonesSettingEditParamsItem() {}

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
	Errors   []SettingWebsocketEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebsocketEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result ZonesWebsockets                          `json:"result"`
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

type SettingWebsocketEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingWebsocketEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebsocketEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingWebsocketEditResponseEnvelopeErrors]
type settingWebsocketEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebsocketEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWebsocketEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingWebsocketEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebsocketEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingWebsocketEditResponseEnvelopeMessages]
type settingWebsocketEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebsocketEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingWebsocketGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingWebsocketGetResponseEnvelope struct {
	Errors   []SettingWebsocketGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebsocketGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result ZonesWebsockets                         `json:"result"`
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

type SettingWebsocketGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingWebsocketGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebsocketGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWebsocketGetResponseEnvelopeErrors]
type settingWebsocketGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebsocketGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWebsocketGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingWebsocketGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebsocketGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingWebsocketGetResponseEnvelopeMessages]
type settingWebsocketGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebsocketGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
