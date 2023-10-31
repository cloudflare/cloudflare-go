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

// ZoneSettingHotlinkProtectionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingHotlinkProtectionService] method instead.
type ZoneSettingHotlinkProtectionService struct {
	Options []option.RequestOption
}

// NewZoneSettingHotlinkProtectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingHotlinkProtectionService(opts ...option.RequestOption) (r *ZoneSettingHotlinkProtectionService) {
	r = &ZoneSettingHotlinkProtectionService{}
	r.Options = opts
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *ZoneSettingHotlinkProtectionService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingHotlinkProtectionUpdateParams, opts ...option.RequestOption) (res *ZoneSettingHotlinkProtectionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *ZoneSettingHotlinkProtectionService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingHotlinkProtectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingHotlinkProtectionUpdateResponse struct {
	Errors   []ZoneSettingHotlinkProtectionUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingHotlinkProtectionUpdateResponseMessage `json:"messages"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result ZoneSettingHotlinkProtectionUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingHotlinkProtectionUpdateResponseJSON
}

// zoneSettingHotlinkProtectionUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingHotlinkProtectionUpdateResponse]
type zoneSettingHotlinkProtectionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionUpdateResponseErrorJSON
}

// zoneSettingHotlinkProtectionUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingHotlinkProtectionUpdateResponseError]
type zoneSettingHotlinkProtectionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionUpdateResponseMessageJSON
}

// zoneSettingHotlinkProtectionUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingHotlinkProtectionUpdateResponseMessage]
type zoneSettingHotlinkProtectionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingHotlinkProtectionUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHotlinkProtectionUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHotlinkProtectionUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingHotlinkProtectionUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingHotlinkProtectionUpdateResponseResultJSON
}

// zoneSettingHotlinkProtectionUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingHotlinkProtectionUpdateResponseResult]
type zoneSettingHotlinkProtectionUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHotlinkProtectionUpdateResponseResultID string

const (
	ZoneSettingHotlinkProtectionUpdateResponseResultIDHotlinkProtection ZoneSettingHotlinkProtectionUpdateResponseResultID = "hotlink_protection"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHotlinkProtectionUpdateResponseResultEditable bool

const (
	ZoneSettingHotlinkProtectionUpdateResponseResultEditableTrue  ZoneSettingHotlinkProtectionUpdateResponseResultEditable = true
	ZoneSettingHotlinkProtectionUpdateResponseResultEditableFalse ZoneSettingHotlinkProtectionUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingHotlinkProtectionUpdateResponseResultValue string

const (
	ZoneSettingHotlinkProtectionUpdateResponseResultValueOn  ZoneSettingHotlinkProtectionUpdateResponseResultValue = "on"
	ZoneSettingHotlinkProtectionUpdateResponseResultValueOff ZoneSettingHotlinkProtectionUpdateResponseResultValue = "off"
)

type ZoneSettingHotlinkProtectionListResponse struct {
	Errors   []ZoneSettingHotlinkProtectionListResponseError   `json:"errors"`
	Messages []ZoneSettingHotlinkProtectionListResponseMessage `json:"messages"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result ZoneSettingHotlinkProtectionListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingHotlinkProtectionListResponseJSON
}

// zoneSettingHotlinkProtectionListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingHotlinkProtectionListResponse]
type zoneSettingHotlinkProtectionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionListResponseErrorJSON
}

// zoneSettingHotlinkProtectionListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingHotlinkProtectionListResponseError]
type zoneSettingHotlinkProtectionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionListResponseMessageJSON
}

// zoneSettingHotlinkProtectionListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingHotlinkProtectionListResponseMessage]
type zoneSettingHotlinkProtectionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingHotlinkProtectionListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHotlinkProtectionListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHotlinkProtectionListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingHotlinkProtectionListResponseResultValue `json:"value"`
	JSON  zoneSettingHotlinkProtectionListResponseResultJSON
}

// zoneSettingHotlinkProtectionListResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingHotlinkProtectionListResponseResult]
type zoneSettingHotlinkProtectionListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHotlinkProtectionListResponseResultID string

const (
	ZoneSettingHotlinkProtectionListResponseResultIDHotlinkProtection ZoneSettingHotlinkProtectionListResponseResultID = "hotlink_protection"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHotlinkProtectionListResponseResultEditable bool

const (
	ZoneSettingHotlinkProtectionListResponseResultEditableTrue  ZoneSettingHotlinkProtectionListResponseResultEditable = true
	ZoneSettingHotlinkProtectionListResponseResultEditableFalse ZoneSettingHotlinkProtectionListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingHotlinkProtectionListResponseResultValue string

const (
	ZoneSettingHotlinkProtectionListResponseResultValueOn  ZoneSettingHotlinkProtectionListResponseResultValue = "on"
	ZoneSettingHotlinkProtectionListResponseResultValueOff ZoneSettingHotlinkProtectionListResponseResultValue = "off"
)

type ZoneSettingHotlinkProtectionUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingHotlinkProtectionUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingHotlinkProtectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingHotlinkProtectionUpdateParamsValue string

const (
	ZoneSettingHotlinkProtectionUpdateParamsValueOn  ZoneSettingHotlinkProtectionUpdateParamsValue = "on"
	ZoneSettingHotlinkProtectionUpdateParamsValueOff ZoneSettingHotlinkProtectionUpdateParamsValue = "off"
)
