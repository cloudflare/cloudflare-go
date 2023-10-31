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

// ZoneSettingIPGeolocationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingIPGeolocationService] method instead.
type ZoneSettingIPGeolocationService struct {
	Options []option.RequestOption
}

// NewZoneSettingIPGeolocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingIPGeolocationService(opts ...option.RequestOption) (r *ZoneSettingIPGeolocationService) {
	r = &ZoneSettingIPGeolocationService{}
	r.Options = opts
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *ZoneSettingIPGeolocationService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingIPGeolocationUpdateParams, opts ...option.RequestOption) (res *ZoneSettingIPGeolocationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *ZoneSettingIPGeolocationService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingIPGeolocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingIPGeolocationUpdateResponse struct {
	Errors   []ZoneSettingIPGeolocationUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingIPGeolocationUpdateResponseMessage `json:"messages"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result ZoneSettingIPGeolocationUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingIPGeolocationUpdateResponseJSON
}

// zoneSettingIPGeolocationUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingIPGeolocationUpdateResponse]
type zoneSettingIPGeolocationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingIPGeolocationUpdateResponseErrorJSON
}

// zoneSettingIPGeolocationUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingIPGeolocationUpdateResponseError]
type zoneSettingIPGeolocationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingIPGeolocationUpdateResponseMessageJSON
}

// zoneSettingIPGeolocationUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingIPGeolocationUpdateResponseMessage]
type zoneSettingIPGeolocationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingIPGeolocationUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingIPGeolocationUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIPGeolocationUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingIPGeolocationUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingIPGeolocationUpdateResponseResultJSON
}

// zoneSettingIPGeolocationUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingIPGeolocationUpdateResponseResult]
type zoneSettingIPGeolocationUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIPGeolocationUpdateResponseResultID string

const (
	ZoneSettingIPGeolocationUpdateResponseResultIDIPGeolocation ZoneSettingIPGeolocationUpdateResponseResultID = "ip_geolocation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIPGeolocationUpdateResponseResultEditable bool

const (
	ZoneSettingIPGeolocationUpdateResponseResultEditableTrue  ZoneSettingIPGeolocationUpdateResponseResultEditable = true
	ZoneSettingIPGeolocationUpdateResponseResultEditableFalse ZoneSettingIPGeolocationUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingIPGeolocationUpdateResponseResultValue string

const (
	ZoneSettingIPGeolocationUpdateResponseResultValueOn  ZoneSettingIPGeolocationUpdateResponseResultValue = "on"
	ZoneSettingIPGeolocationUpdateResponseResultValueOff ZoneSettingIPGeolocationUpdateResponseResultValue = "off"
)

type ZoneSettingIPGeolocationListResponse struct {
	Errors   []ZoneSettingIPGeolocationListResponseError   `json:"errors"`
	Messages []ZoneSettingIPGeolocationListResponseMessage `json:"messages"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result ZoneSettingIPGeolocationListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingIPGeolocationListResponseJSON
}

// zoneSettingIPGeolocationListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingIPGeolocationListResponse]
type zoneSettingIPGeolocationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingIPGeolocationListResponseErrorJSON
}

// zoneSettingIPGeolocationListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingIPGeolocationListResponseError]
type zoneSettingIPGeolocationListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingIPGeolocationListResponseMessageJSON
}

// zoneSettingIPGeolocationListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingIPGeolocationListResponseMessage]
type zoneSettingIPGeolocationListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingIPGeolocationListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingIPGeolocationListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIPGeolocationListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingIPGeolocationListResponseResultValue `json:"value"`
	JSON  zoneSettingIPGeolocationListResponseResultJSON
}

// zoneSettingIPGeolocationListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingIPGeolocationListResponseResult]
type zoneSettingIPGeolocationListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIPGeolocationListResponseResultID string

const (
	ZoneSettingIPGeolocationListResponseResultIDIPGeolocation ZoneSettingIPGeolocationListResponseResultID = "ip_geolocation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIPGeolocationListResponseResultEditable bool

const (
	ZoneSettingIPGeolocationListResponseResultEditableTrue  ZoneSettingIPGeolocationListResponseResultEditable = true
	ZoneSettingIPGeolocationListResponseResultEditableFalse ZoneSettingIPGeolocationListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingIPGeolocationListResponseResultValue string

const (
	ZoneSettingIPGeolocationListResponseResultValueOn  ZoneSettingIPGeolocationListResponseResultValue = "on"
	ZoneSettingIPGeolocationListResponseResultValueOff ZoneSettingIPGeolocationListResponseResultValue = "off"
)

type ZoneSettingIPGeolocationUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingIPGeolocationUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingIPGeolocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingIPGeolocationUpdateParamsValue string

const (
	ZoneSettingIPGeolocationUpdateParamsValueOn  ZoneSettingIPGeolocationUpdateParamsValue = "on"
	ZoneSettingIPGeolocationUpdateParamsValueOff ZoneSettingIPGeolocationUpdateParamsValue = "off"
)
