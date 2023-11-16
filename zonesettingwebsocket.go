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
func (r *ZoneSettingWebsocketService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingWebsocketUpdateParams, opts ...option.RequestOption) (res *ZoneSettingWebsocketUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/websockets", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets Websockets setting. For more information about Websockets, please refer to
// [Using Cloudflare with WebSockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Using-Cloudflare-with-WebSockets).
func (r *ZoneSettingWebsocketService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingWebsocketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/websockets", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingWebsocketUpdateResponse struct {
	Errors   []ZoneSettingWebsocketUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingWebsocketUpdateResponseMessage `json:"messages"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result ZoneSettingWebsocketUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                   `json:"success"`
	JSON    zoneSettingWebsocketUpdateResponseJSON `json:"-"`
}

// zoneSettingWebsocketUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebsocketUpdateResponse]
type zoneSettingWebsocketUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingWebsocketUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingWebsocketUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketUpdateResponseError]
type zoneSettingWebsocketUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingWebsocketUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingWebsocketUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketUpdateResponseMessage]
type zoneSettingWebsocketUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingWebsocketUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWebsocketUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebsocketUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingWebsocketUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingWebsocketUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingWebsocketUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketUpdateResponseResult]
type zoneSettingWebsocketUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWebsocketUpdateResponseResultID string

const (
	ZoneSettingWebsocketUpdateResponseResultIDWebsockets ZoneSettingWebsocketUpdateResponseResultID = "websockets"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebsocketUpdateResponseResultEditable bool

const (
	ZoneSettingWebsocketUpdateResponseResultEditableTrue  ZoneSettingWebsocketUpdateResponseResultEditable = true
	ZoneSettingWebsocketUpdateResponseResultEditableFalse ZoneSettingWebsocketUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingWebsocketUpdateResponseResultValue string

const (
	ZoneSettingWebsocketUpdateResponseResultValueOff ZoneSettingWebsocketUpdateResponseResultValue = "off"
	ZoneSettingWebsocketUpdateResponseResultValueOn  ZoneSettingWebsocketUpdateResponseResultValue = "on"
)

type ZoneSettingWebsocketListResponse struct {
	Errors   []ZoneSettingWebsocketListResponseError   `json:"errors"`
	Messages []ZoneSettingWebsocketListResponseMessage `json:"messages"`
	// WebSockets are open connections sustained between the client and the origin
	// server. Inside a WebSockets connection, the client and the origin can pass data
	// back and forth without having to reestablish sessions. This makes exchanging
	// data within a WebSockets connection fast. WebSockets are often used for
	// real-time applications such as live chat and gaming. For more information refer
	// to
	// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
	Result ZoneSettingWebsocketListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                 `json:"success"`
	JSON    zoneSettingWebsocketListResponseJSON `json:"-"`
}

// zoneSettingWebsocketListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebsocketListResponse]
type zoneSettingWebsocketListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketListResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingWebsocketListResponseErrorJSON `json:"-"`
}

// zoneSettingWebsocketListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketListResponseError]
type zoneSettingWebsocketListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebsocketListResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingWebsocketListResponseMessageJSON `json:"-"`
}

// zoneSettingWebsocketListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketListResponseMessage]
type zoneSettingWebsocketListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingWebsocketListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWebsocketListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebsocketListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingWebsocketListResponseResultValue `json:"value"`
	JSON  zoneSettingWebsocketListResponseResultJSON  `json:"-"`
}

// zoneSettingWebsocketListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketListResponseResult]
type zoneSettingWebsocketListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebsocketListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWebsocketListResponseResultID string

const (
	ZoneSettingWebsocketListResponseResultIDWebsockets ZoneSettingWebsocketListResponseResultID = "websockets"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebsocketListResponseResultEditable bool

const (
	ZoneSettingWebsocketListResponseResultEditableTrue  ZoneSettingWebsocketListResponseResultEditable = true
	ZoneSettingWebsocketListResponseResultEditableFalse ZoneSettingWebsocketListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingWebsocketListResponseResultValue string

const (
	ZoneSettingWebsocketListResponseResultValueOff ZoneSettingWebsocketListResponseResultValue = "off"
	ZoneSettingWebsocketListResponseResultValueOn  ZoneSettingWebsocketListResponseResultValue = "on"
)

type ZoneSettingWebsocketUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingWebsocketUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingWebsocketUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingWebsocketUpdateParamsValue string

const (
	ZoneSettingWebsocketUpdateParamsValueOff ZoneSettingWebsocketUpdateParamsValue = "off"
	ZoneSettingWebsocketUpdateParamsValueOn  ZoneSettingWebsocketUpdateParamsValue = "on"
)
