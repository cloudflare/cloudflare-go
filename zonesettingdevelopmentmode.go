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

// ZoneSettingDevelopmentModeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingDevelopmentModeService] method instead.
type ZoneSettingDevelopmentModeService struct {
	Options []option.RequestOption
}

// NewZoneSettingDevelopmentModeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingDevelopmentModeService(opts ...option.RequestOption) (r *ZoneSettingDevelopmentModeService) {
	r = &ZoneSettingDevelopmentModeService{}
	r.Options = opts
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *ZoneSettingDevelopmentModeService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingDevelopmentModeUpdateParams, opts ...option.RequestOption) (res *ZoneSettingDevelopmentModeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/development_mode", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *ZoneSettingDevelopmentModeService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingDevelopmentModeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/development_mode", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingDevelopmentModeUpdateResponse struct {
	Errors   []ZoneSettingDevelopmentModeUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingDevelopmentModeUpdateResponseMessage `json:"messages"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZoneSettingDevelopmentModeUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                         `json:"success"`
	JSON    zoneSettingDevelopmentModeUpdateResponseJSON `json:"-"`
}

// zoneSettingDevelopmentModeUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingDevelopmentModeUpdateResponse]
type zoneSettingDevelopmentModeUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingDevelopmentModeUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingDevelopmentModeUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingDevelopmentModeUpdateResponseError]
type zoneSettingDevelopmentModeUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingDevelopmentModeUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingDevelopmentModeUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingDevelopmentModeUpdateResponseMessage]
type zoneSettingDevelopmentModeUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingDevelopmentModeUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingDevelopmentModeUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingDevelopmentModeUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// Value of the zone setting.
	Value ZoneSettingDevelopmentModeUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingDevelopmentModeUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingDevelopmentModeUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingDevelopmentModeUpdateResponseResult]
type zoneSettingDevelopmentModeUpdateResponseResultJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingDevelopmentModeUpdateResponseResultID string

const (
	ZoneSettingDevelopmentModeUpdateResponseResultIDDevelopmentMode ZoneSettingDevelopmentModeUpdateResponseResultID = "development_mode"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingDevelopmentModeUpdateResponseResultEditable bool

const (
	ZoneSettingDevelopmentModeUpdateResponseResultEditableTrue  ZoneSettingDevelopmentModeUpdateResponseResultEditable = true
	ZoneSettingDevelopmentModeUpdateResponseResultEditableFalse ZoneSettingDevelopmentModeUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingDevelopmentModeUpdateResponseResultValue string

const (
	ZoneSettingDevelopmentModeUpdateResponseResultValueOn  ZoneSettingDevelopmentModeUpdateResponseResultValue = "on"
	ZoneSettingDevelopmentModeUpdateResponseResultValueOff ZoneSettingDevelopmentModeUpdateResponseResultValue = "off"
)

type ZoneSettingDevelopmentModeListResponse struct {
	Errors   []ZoneSettingDevelopmentModeListResponseError   `json:"errors"`
	Messages []ZoneSettingDevelopmentModeListResponseMessage `json:"messages"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZoneSettingDevelopmentModeListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                       `json:"success"`
	JSON    zoneSettingDevelopmentModeListResponseJSON `json:"-"`
}

// zoneSettingDevelopmentModeListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingDevelopmentModeListResponse]
type zoneSettingDevelopmentModeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeListResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingDevelopmentModeListResponseErrorJSON `json:"-"`
}

// zoneSettingDevelopmentModeListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingDevelopmentModeListResponseError]
type zoneSettingDevelopmentModeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeListResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingDevelopmentModeListResponseMessageJSON `json:"-"`
}

// zoneSettingDevelopmentModeListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingDevelopmentModeListResponseMessage]
type zoneSettingDevelopmentModeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingDevelopmentModeListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingDevelopmentModeListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingDevelopmentModeListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// Value of the zone setting.
	Value ZoneSettingDevelopmentModeListResponseResultValue `json:"value"`
	JSON  zoneSettingDevelopmentModeListResponseResultJSON  `json:"-"`
}

// zoneSettingDevelopmentModeListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingDevelopmentModeListResponseResult]
type zoneSettingDevelopmentModeListResponseResultJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingDevelopmentModeListResponseResultID string

const (
	ZoneSettingDevelopmentModeListResponseResultIDDevelopmentMode ZoneSettingDevelopmentModeListResponseResultID = "development_mode"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingDevelopmentModeListResponseResultEditable bool

const (
	ZoneSettingDevelopmentModeListResponseResultEditableTrue  ZoneSettingDevelopmentModeListResponseResultEditable = true
	ZoneSettingDevelopmentModeListResponseResultEditableFalse ZoneSettingDevelopmentModeListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingDevelopmentModeListResponseResultValue string

const (
	ZoneSettingDevelopmentModeListResponseResultValueOn  ZoneSettingDevelopmentModeListResponseResultValue = "on"
	ZoneSettingDevelopmentModeListResponseResultValueOff ZoneSettingDevelopmentModeListResponseResultValue = "off"
)

type ZoneSettingDevelopmentModeUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingDevelopmentModeUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingDevelopmentModeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingDevelopmentModeUpdateParamsValue string

const (
	ZoneSettingDevelopmentModeUpdateParamsValueOn  ZoneSettingDevelopmentModeUpdateParamsValue = "on"
	ZoneSettingDevelopmentModeUpdateParamsValueOff ZoneSettingDevelopmentModeUpdateParamsValue = "off"
)
