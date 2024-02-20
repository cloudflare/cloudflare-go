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
func (r *SettingWebsocketService) Edit(ctx context.Context, zoneID string, body SettingWebsocketEditParams, opts ...option.RequestOption) (res *SettingWebsocketEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebsocketEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/websockets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Websockets setting. For more information about Websockets, please refer to
// [Using Cloudflare with WebSockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Using-Cloudflare-with-WebSockets).
func (r *SettingWebsocketService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingWebsocketGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebsocketGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/websockets", zoneID)
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
type SettingWebsocketEditResponse struct {
	// ID of the zone setting.
	ID SettingWebsocketEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingWebsocketEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingWebsocketEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingWebsocketEditResponseJSON `json:"-"`
}

// settingWebsocketEditResponseJSON contains the JSON metadata for the struct
// [SettingWebsocketEditResponse]
type settingWebsocketEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingWebsocketEditResponseID string

const (
	SettingWebsocketEditResponseIDWebsockets SettingWebsocketEditResponseID = "websockets"
)

// Current value of the zone setting.
type SettingWebsocketEditResponseValue string

const (
	SettingWebsocketEditResponseValueOff SettingWebsocketEditResponseValue = "off"
	SettingWebsocketEditResponseValueOn  SettingWebsocketEditResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingWebsocketEditResponseEditable bool

const (
	SettingWebsocketEditResponseEditableTrue  SettingWebsocketEditResponseEditable = true
	SettingWebsocketEditResponseEditableFalse SettingWebsocketEditResponseEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type SettingWebsocketGetResponse struct {
	// ID of the zone setting.
	ID SettingWebsocketGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingWebsocketGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingWebsocketGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingWebsocketGetResponseJSON `json:"-"`
}

// settingWebsocketGetResponseJSON contains the JSON metadata for the struct
// [SettingWebsocketGetResponse]
type settingWebsocketGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebsocketGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingWebsocketGetResponseID string

const (
	SettingWebsocketGetResponseIDWebsockets SettingWebsocketGetResponseID = "websockets"
)

// Current value of the zone setting.
type SettingWebsocketGetResponseValue string

const (
	SettingWebsocketGetResponseValueOff SettingWebsocketGetResponseValue = "off"
	SettingWebsocketGetResponseValueOn  SettingWebsocketGetResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingWebsocketGetResponseEditable bool

const (
	SettingWebsocketGetResponseEditableTrue  SettingWebsocketGetResponseEditable = true
	SettingWebsocketGetResponseEditableFalse SettingWebsocketGetResponseEditable = false
)

type SettingWebsocketEditParams struct {
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
	Result SettingWebsocketEditResponse             `json:"result"`
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
	Result SettingWebsocketGetResponse             `json:"result"`
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
