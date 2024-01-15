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

// ZoneSettingZeroRttService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingZeroRttService] method
// instead.
type ZoneSettingZeroRttService struct {
	Options []option.RequestOption
}

// NewZoneSettingZeroRttService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZeroRttService(opts ...option.RequestOption) (r *ZoneSettingZeroRttService) {
	r = &ZoneSettingZeroRttService{}
	r.Options = opts
	return
}

// Gets 0-RTT session resumption setting.
func (r *ZoneSettingZeroRttService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingZeroRttListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/0rtt", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Changes the 0-RTT session resumption setting.
func (r *ZoneSettingZeroRttService) ZoneSettingsChange0RttSessionResumptionSetting(ctx context.Context, zoneIdentifier string, body ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParams, opts ...option.RequestOption) (res *ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/0rtt", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneSettingZeroRttListResponse struct {
	Errors   []ZoneSettingZeroRttListResponseError   `json:"errors"`
	Messages []ZoneSettingZeroRttListResponseMessage `json:"messages"`
	// 0-RTT session resumption enabled for this zone.
	Result ZoneSettingZeroRttListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                               `json:"success"`
	JSON    zoneSettingZeroRttListResponseJSON `json:"-"`
}

// zoneSettingZeroRttListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingZeroRttListResponse]
type zoneSettingZeroRttListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRttListResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingZeroRttListResponseErrorJSON `json:"-"`
}

// zoneSettingZeroRttListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingZeroRttListResponseError]
type zoneSettingZeroRttListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRttListResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingZeroRttListResponseMessageJSON `json:"-"`
}

// zoneSettingZeroRttListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingZeroRttListResponseMessage]
type zoneSettingZeroRttListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingZeroRttListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingZeroRttListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingZeroRttListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the 0-RTT setting.
	Value ZoneSettingZeroRttListResponseResultValue `json:"value"`
	JSON  zoneSettingZeroRttListResponseResultJSON  `json:"-"`
}

// zoneSettingZeroRttListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingZeroRttListResponseResult]
type zoneSettingZeroRttListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingZeroRttListResponseResultID string

const (
	ZoneSettingZeroRttListResponseResultID0rtt ZoneSettingZeroRttListResponseResultID = "0rtt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingZeroRttListResponseResultEditable bool

const (
	ZoneSettingZeroRttListResponseResultEditableTrue  ZoneSettingZeroRttListResponseResultEditable = true
	ZoneSettingZeroRttListResponseResultEditableFalse ZoneSettingZeroRttListResponseResultEditable = false
)

// Value of the 0-RTT setting.
type ZoneSettingZeroRttListResponseResultValue string

const (
	ZoneSettingZeroRttListResponseResultValueOn  ZoneSettingZeroRttListResponseResultValue = "on"
	ZoneSettingZeroRttListResponseResultValueOff ZoneSettingZeroRttListResponseResultValue = "off"
)

type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse struct {
	Errors   []ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseError   `json:"errors"`
	Messages []ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseMessage `json:"messages"`
	// 0-RTT session resumption enabled for this zone.
	Result ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                                                         `json:"success"`
	JSON    zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseJSON `json:"-"`
}

// zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseJSON
// contains the JSON metadata for the struct
// [ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse]
type zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseErrorJSON `json:"-"`
}

// zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseError]
type zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseMessageJSON `json:"-"`
}

// zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseMessage]
type zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the 0-RTT setting.
	Value ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultValue `json:"value"`
	JSON  zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultJSON  `json:"-"`
}

// zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResult]
type zoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultID string

const (
	ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultID0rtt ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultID = "0rtt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultEditable bool

const (
	ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultEditableTrue  ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultEditable = true
	ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultEditableFalse ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultEditable = false
)

// Value of the 0-RTT setting.
type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultValue string

const (
	ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultValueOn  ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultValue = "on"
	ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultValueOff ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseResultValue = "off"
)

type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParams struct {
	// Value of the 0-RTT setting.
	Value param.Field[ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue] `json:"value,required"`
}

func (r ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the 0-RTT setting.
type ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue string

const (
	ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValueOn  ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue = "on"
	ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValueOff ZoneSettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue = "off"
)
