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

// ZoneSettingNelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingNelService] method
// instead.
type ZoneSettingNelService struct {
	Options []option.RequestOption
}

// NewZoneSettingNelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingNelService(opts ...option.RequestOption) (r *ZoneSettingNelService) {
	r = &ZoneSettingNelService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/nel-solving-mobile-speed)
// for more information.
func (r *ZoneSettingNelService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingNelUpdateParams, opts ...option.RequestOption) (res *ZoneSettingNelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/nel", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
func (r *ZoneSettingNelService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingNelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/nel", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingNelUpdateResponse struct {
	Errors   []ZoneSettingNelUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingNelUpdateResponseMessage `json:"messages"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZoneSettingNelUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                             `json:"success"`
	JSON    zoneSettingNelUpdateResponseJSON `json:"-"`
}

// zoneSettingNelUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingNelUpdateResponse]
type zoneSettingNelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNelUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingNelUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingNelUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingNelUpdateResponseError]
type zoneSettingNelUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNelUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingNelUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingNelUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingNelUpdateResponseMessage]
type zoneSettingNelUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNelUpdateResponseResult struct {
	// Zone setting identifier.
	ID ZoneSettingNelUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingNelUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingNelUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingNelUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingNelUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingNelUpdateResponseResult]
type zoneSettingNelUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type ZoneSettingNelUpdateResponseResultID string

const (
	ZoneSettingNelUpdateResponseResultIDNel ZoneSettingNelUpdateResponseResultID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNelUpdateResponseResultEditable bool

const (
	ZoneSettingNelUpdateResponseResultEditableTrue  ZoneSettingNelUpdateResponseResultEditable = true
	ZoneSettingNelUpdateResponseResultEditableFalse ZoneSettingNelUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingNelUpdateResponseResultValue struct {
	Enabled bool                                        `json:"enabled"`
	JSON    zoneSettingNelUpdateResponseResultValueJSON `json:"-"`
}

// zoneSettingNelUpdateResponseResultValueJSON contains the JSON metadata for the
// struct [ZoneSettingNelUpdateResponseResultValue]
type zoneSettingNelUpdateResponseResultValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelUpdateResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNelListResponse struct {
	Errors   []ZoneSettingNelListResponseError   `json:"errors"`
	Messages []ZoneSettingNelListResponseMessage `json:"messages"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZoneSettingNelListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	JSON    zoneSettingNelListResponseJSON `json:"-"`
}

// zoneSettingNelListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingNelListResponse]
type zoneSettingNelListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNelListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneSettingNelListResponseErrorJSON `json:"-"`
}

// zoneSettingNelListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingNelListResponseError]
type zoneSettingNelListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNelListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingNelListResponseMessageJSON `json:"-"`
}

// zoneSettingNelListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingNelListResponseMessage]
type zoneSettingNelListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNelListResponseResult struct {
	// Zone setting identifier.
	ID ZoneSettingNelListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingNelListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingNelListResponseResultValue `json:"value"`
	JSON  zoneSettingNelListResponseResultJSON  `json:"-"`
}

// zoneSettingNelListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingNelListResponseResult]
type zoneSettingNelListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type ZoneSettingNelListResponseResultID string

const (
	ZoneSettingNelListResponseResultIDNel ZoneSettingNelListResponseResultID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNelListResponseResultEditable bool

const (
	ZoneSettingNelListResponseResultEditableTrue  ZoneSettingNelListResponseResultEditable = true
	ZoneSettingNelListResponseResultEditableFalse ZoneSettingNelListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingNelListResponseResultValue struct {
	Enabled bool                                      `json:"enabled"`
	JSON    zoneSettingNelListResponseResultValueJSON `json:"-"`
}

// zoneSettingNelListResponseResultValueJSON contains the JSON metadata for the
// struct [ZoneSettingNelListResponseResultValue]
type zoneSettingNelListResponseResultValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNelListResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNelUpdateParams struct {
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[ZoneSettingNelUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingNelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNelUpdateParamsValue struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingNelUpdateParamsValueID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingNelUpdateParamsValueValue] `json:"value"`
}

func (r ZoneSettingNelUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Zone setting identifier.
type ZoneSettingNelUpdateParamsValueID string

const (
	ZoneSettingNelUpdateParamsValueIDNel ZoneSettingNelUpdateParamsValueID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNelUpdateParamsValueEditable bool

const (
	ZoneSettingNelUpdateParamsValueEditableTrue  ZoneSettingNelUpdateParamsValueEditable = true
	ZoneSettingNelUpdateParamsValueEditableFalse ZoneSettingNelUpdateParamsValueEditable = false
)

// Value of the zone setting.
type ZoneSettingNelUpdateParamsValueValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingNelUpdateParamsValueValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
