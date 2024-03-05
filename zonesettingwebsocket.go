// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingWebsocketService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingWebsocketService]
// method instead.
type ZoneSettingWebsocketService struct {
	Options []option.RequestOption
}

// NewZoneSettingWebsocketService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingWebsocketService(opts ...option.RequestOption) (r *ZoneSettingWebsocketService) {
	r = &ZoneSettingWebsocketService{}
	r.Options = opts
	return
}

// Changes Websockets setting. For more information about Websockets, please refer
// to
// [Using Cloudflare with WebSockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Using-Cloudflare-with-WebSockets).
func (r *ZoneSettingWebsocketService) Edit(ctx context.Context, params ZoneSettingWebsocketEditParams, opts ...option.RequestOption) (res *ZonesWebsockets, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingWebsocketEditResponseEnvelope
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
func (r *ZoneSettingWebsocketService) Get(ctx context.Context, query ZoneSettingWebsocketGetParams, opts ...option.RequestOption) (res *ZonesWebsockets, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingWebsocketGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/websockets", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

func (r ZonesWebsockets) implementsZoneSettingEditResponse() {}

func (r ZonesWebsockets) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesWebsocketsID string

const (
	ZonesWebsocketsIDWebsockets ZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type ZonesWebsocketsValue string

const (
	ZonesWebsocketsValueOff ZonesWebsocketsValue = "off"
	ZonesWebsocketsValueOn  ZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesWebsocketsEditable bool

const (
	ZonesWebsocketsEditableTrue  ZonesWebsocketsEditable = true
	ZonesWebsocketsEditableFalse ZonesWebsocketsEditable = false
)

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

func (r ZonesWebsocketsParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingWebsocketEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingWebsocketEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingWebsocketEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingWebsocketEditParamsValue string

const (
	ZoneSettingWebsocketEditParamsValueOff ZoneSettingWebsocketEditParamsValue = "off"
	ZoneSettingWebsocketEditParamsValueOn  ZoneSettingWebsocketEditParamsValue = "on"
)

type ZoneSettingWebsocketEditResponseEnvelope struct {
	Errors   []ZoneSettingWebsocketEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingWebsocketEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result ZonesWebsockets                              `json:"result"`
	JSON   zoneSettingWebsocketEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingWebsocketEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketEditResponseEnvelope]
type zoneSettingWebsocketEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingWebsocketEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingWebsocketEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingWebsocketEditResponseEnvelopeErrors]
type zoneSettingWebsocketEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingWebsocketEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingWebsocketEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingWebsocketEditResponseEnvelopeMessages]
type zoneSettingWebsocketEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingWebsocketGetResponseEnvelope struct {
	Errors   []ZoneSettingWebsocketGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingWebsocketGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result ZonesWebsockets                             `json:"result"`
	JSON   zoneSettingWebsocketGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingWebsocketGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketGetResponseEnvelope]
type zoneSettingWebsocketGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingWebsocketGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingWebsocketGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingWebsocketGetResponseEnvelopeErrors]
type zoneSettingWebsocketGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingWebsocketGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingWebsocketGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingWebsocketGetResponseEnvelopeMessages]
type zoneSettingWebsocketGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
