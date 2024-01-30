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
	Errors   []ZoneSettingWebsocketUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingWebsocketUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                     `json:"success,required"`
	Result  ZoneSettingWebsocketUpdateResponseResult `json:"result"`
	JSON    zoneSettingWebsocketUpdateResponseJSON   `json:"-"`
}

// zoneSettingWebsocketUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebsocketUpdateResponse]
type zoneSettingWebsocketUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

type ZoneSettingWebsocketUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWebsocketUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingWebsocketUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebsocketUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWebsocketUpdateResponseResultJSON `json:"-"`
}

// zoneSettingWebsocketUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketUpdateResponseResult]
type zoneSettingWebsocketUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
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

// Value of the zone setting.
type ZoneSettingWebsocketUpdateResponseResultValue string

const (
	ZoneSettingWebsocketUpdateResponseResultValueOff ZoneSettingWebsocketUpdateResponseResultValue = "off"
	ZoneSettingWebsocketUpdateResponseResultValueOn  ZoneSettingWebsocketUpdateResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebsocketUpdateResponseResultEditable bool

const (
	ZoneSettingWebsocketUpdateResponseResultEditableTrue  ZoneSettingWebsocketUpdateResponseResultEditable = true
	ZoneSettingWebsocketUpdateResponseResultEditableFalse ZoneSettingWebsocketUpdateResponseResultEditable = false
)

type ZoneSettingWebsocketListResponse struct {
	Errors   []ZoneSettingWebsocketListResponseError   `json:"errors,required"`
	Messages []ZoneSettingWebsocketListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                   `json:"success,required"`
	Result  ZoneSettingWebsocketListResponseResult `json:"result"`
	JSON    zoneSettingWebsocketListResponseJSON   `json:"-"`
}

// zoneSettingWebsocketListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebsocketListResponse]
type zoneSettingWebsocketListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

type ZoneSettingWebsocketListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWebsocketListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingWebsocketListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebsocketListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWebsocketListResponseResultJSON `json:"-"`
}

// zoneSettingWebsocketListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingWebsocketListResponseResult]
type zoneSettingWebsocketListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
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

// Value of the zone setting.
type ZoneSettingWebsocketListResponseResultValue string

const (
	ZoneSettingWebsocketListResponseResultValueOff ZoneSettingWebsocketListResponseResultValue = "off"
	ZoneSettingWebsocketListResponseResultValueOn  ZoneSettingWebsocketListResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebsocketListResponseResultEditable bool

const (
	ZoneSettingWebsocketListResponseResultEditableTrue  ZoneSettingWebsocketListResponseResultEditable = true
	ZoneSettingWebsocketListResponseResultEditableFalse ZoneSettingWebsocketListResponseResultEditable = false
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
