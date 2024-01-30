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

// ZoneSettingNELService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingNELService] method
// instead.
type ZoneSettingNELService struct {
	Options []option.RequestOption
}

// NewZoneSettingNELService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingNELService(opts ...option.RequestOption) (r *ZoneSettingNELService) {
	r = &ZoneSettingNELService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/nel-solving-mobile-speed)
// for more information.
func (r *ZoneSettingNELService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingNELUpdateParams, opts ...option.RequestOption) (res *ZoneSettingNELUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/nel", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
func (r *ZoneSettingNELService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingNELListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/nel", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingNELUpdateResponse struct {
	Errors   []ZoneSettingNELUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingNELUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                               `json:"success,required"`
	Result  ZoneSettingNELUpdateResponseResult `json:"result"`
	JSON    zoneSettingNELUpdateResponseJSON   `json:"-"`
}

// zoneSettingNELUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingNELUpdateResponse]
type zoneSettingNELUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingNELUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingNELUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingNELUpdateResponseError]
type zoneSettingNELUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingNELUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingNELUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingNELUpdateResponseMessage]
type zoneSettingNELUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELUpdateResponseResult struct {
	// Zone setting identifier.
	ID ZoneSettingNELUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingNELUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingNELUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingNELUpdateResponseResultJSON `json:"-"`
}

// zoneSettingNELUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingNELUpdateResponseResult]
type zoneSettingNELUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type ZoneSettingNELUpdateResponseResultID string

const (
	ZoneSettingNELUpdateResponseResultIDNEL ZoneSettingNELUpdateResponseResultID = "nel"
)

// Value of the zone setting.
type ZoneSettingNELUpdateResponseResultValue struct {
	Enabled bool                                        `json:"enabled"`
	JSON    zoneSettingNELUpdateResponseResultValueJSON `json:"-"`
}

// zoneSettingNELUpdateResponseResultValueJSON contains the JSON metadata for the
// struct [ZoneSettingNELUpdateResponseResultValue]
type zoneSettingNELUpdateResponseResultValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELUpdateResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNELUpdateResponseResultEditable bool

const (
	ZoneSettingNELUpdateResponseResultEditableTrue  ZoneSettingNELUpdateResponseResultEditable = true
	ZoneSettingNELUpdateResponseResultEditableFalse ZoneSettingNELUpdateResponseResultEditable = false
)

type ZoneSettingNELListResponse struct {
	Errors   []ZoneSettingNELListResponseError   `json:"errors,required"`
	Messages []ZoneSettingNELListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                             `json:"success,required"`
	Result  ZoneSettingNELListResponseResult `json:"result"`
	JSON    zoneSettingNELListResponseJSON   `json:"-"`
}

// zoneSettingNELListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingNELListResponse]
type zoneSettingNELListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneSettingNELListResponseErrorJSON `json:"-"`
}

// zoneSettingNELListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingNELListResponseError]
type zoneSettingNELListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingNELListResponseMessageJSON `json:"-"`
}

// zoneSettingNELListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingNELListResponseMessage]
type zoneSettingNELListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELListResponseResult struct {
	// Zone setting identifier.
	ID ZoneSettingNELListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingNELListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingNELListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingNELListResponseResultJSON `json:"-"`
}

// zoneSettingNELListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingNELListResponseResult]
type zoneSettingNELListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type ZoneSettingNELListResponseResultID string

const (
	ZoneSettingNELListResponseResultIDNEL ZoneSettingNELListResponseResultID = "nel"
)

// Value of the zone setting.
type ZoneSettingNELListResponseResultValue struct {
	Enabled bool                                      `json:"enabled"`
	JSON    zoneSettingNELListResponseResultValueJSON `json:"-"`
}

// zoneSettingNELListResponseResultValueJSON contains the JSON metadata for the
// struct [ZoneSettingNELListResponseResultValue]
type zoneSettingNELListResponseResultValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELListResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNELListResponseResultEditable bool

const (
	ZoneSettingNELListResponseResultEditableTrue  ZoneSettingNELListResponseResultEditable = true
	ZoneSettingNELListResponseResultEditableFalse ZoneSettingNELListResponseResultEditable = false
)

type ZoneSettingNELUpdateParams struct {
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[ZoneSettingNELUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingNELUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNELUpdateParamsValue struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingNELUpdateParamsValueID] `json:"id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingNELUpdateParamsValueValue] `json:"value,required"`
}

func (r ZoneSettingNELUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Zone setting identifier.
type ZoneSettingNELUpdateParamsValueID string

const (
	ZoneSettingNELUpdateParamsValueIDNEL ZoneSettingNELUpdateParamsValueID = "nel"
)

// Value of the zone setting.
type ZoneSettingNELUpdateParamsValueValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingNELUpdateParamsValueValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNELUpdateParamsValueEditable bool

const (
	ZoneSettingNELUpdateParamsValueEditableTrue  ZoneSettingNELUpdateParamsValueEditable = true
	ZoneSettingNELUpdateParamsValueEditableFalse ZoneSettingNELUpdateParamsValueEditable = false
)
