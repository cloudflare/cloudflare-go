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

// ZoneSettingMirageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingMirageService] method
// instead.
type ZoneSettingMirageService struct {
	Options []option.RequestOption
}

// NewZoneSettingMirageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingMirageService(opts ...option.RequestOption) (r *ZoneSettingMirageService) {
	r = &ZoneSettingMirageService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *ZoneSettingMirageService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingMirageUpdateParams, opts ...option.RequestOption) (res *ZoneSettingMirageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mirage", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *ZoneSettingMirageService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingMirageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mirage", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingMirageUpdateResponse struct {
	Errors   []ZoneSettingMirageUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingMirageUpdateResponseMessage `json:"messages"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result ZoneSettingMirageUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                `json:"success"`
	JSON    zoneSettingMirageUpdateResponseJSON `json:"-"`
}

// zoneSettingMirageUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingMirageUpdateResponse]
type zoneSettingMirageUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMirageUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingMirageUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingMirageUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingMirageUpdateResponseError]
type zoneSettingMirageUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMirageUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSettingMirageUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingMirageUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingMirageUpdateResponseMessage]
type zoneSettingMirageUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingMirageUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingMirageUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMirageUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingMirageUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingMirageUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingMirageUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingMirageUpdateResponseResult]
type zoneSettingMirageUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingMirageUpdateResponseResultID string

const (
	ZoneSettingMirageUpdateResponseResultIDMirage ZoneSettingMirageUpdateResponseResultID = "mirage"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMirageUpdateResponseResultEditable bool

const (
	ZoneSettingMirageUpdateResponseResultEditableTrue  ZoneSettingMirageUpdateResponseResultEditable = true
	ZoneSettingMirageUpdateResponseResultEditableFalse ZoneSettingMirageUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingMirageUpdateResponseResultValue string

const (
	ZoneSettingMirageUpdateResponseResultValueOn  ZoneSettingMirageUpdateResponseResultValue = "on"
	ZoneSettingMirageUpdateResponseResultValueOff ZoneSettingMirageUpdateResponseResultValue = "off"
)

type ZoneSettingMirageListResponse struct {
	Errors   []ZoneSettingMirageListResponseError   `json:"errors"`
	Messages []ZoneSettingMirageListResponseMessage `json:"messages"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result ZoneSettingMirageListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    zoneSettingMirageListResponseJSON `json:"-"`
}

// zoneSettingMirageListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingMirageListResponse]
type zoneSettingMirageListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMirageListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingMirageListResponseErrorJSON `json:"-"`
}

// zoneSettingMirageListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingMirageListResponseError]
type zoneSettingMirageListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMirageListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingMirageListResponseMessageJSON `json:"-"`
}

// zoneSettingMirageListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingMirageListResponseMessage]
type zoneSettingMirageListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingMirageListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingMirageListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMirageListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingMirageListResponseResultValue `json:"value"`
	JSON  zoneSettingMirageListResponseResultJSON  `json:"-"`
}

// zoneSettingMirageListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingMirageListResponseResult]
type zoneSettingMirageListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingMirageListResponseResultID string

const (
	ZoneSettingMirageListResponseResultIDMirage ZoneSettingMirageListResponseResultID = "mirage"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMirageListResponseResultEditable bool

const (
	ZoneSettingMirageListResponseResultEditableTrue  ZoneSettingMirageListResponseResultEditable = true
	ZoneSettingMirageListResponseResultEditableFalse ZoneSettingMirageListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingMirageListResponseResultValue string

const (
	ZoneSettingMirageListResponseResultValueOn  ZoneSettingMirageListResponseResultValue = "on"
	ZoneSettingMirageListResponseResultValueOff ZoneSettingMirageListResponseResultValue = "off"
)

type ZoneSettingMirageUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingMirageUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingMirageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMirageUpdateParamsValue string

const (
	ZoneSettingMirageUpdateParamsValueOn  ZoneSettingMirageUpdateParamsValue = "on"
	ZoneSettingMirageUpdateParamsValueOff ZoneSettingMirageUpdateParamsValue = "off"
)
