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

// ZoneSettingOriginErrorPagePassThrusService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingOriginErrorPagePassThrusService] method instead.
type ZoneSettingOriginErrorPagePassThrusService struct {
	Options []option.RequestOption
}

// NewZoneSettingOriginErrorPagePassThrusService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSettingOriginErrorPagePassThrusService(opts ...option.RequestOption) (r *ZoneSettingOriginErrorPagePassThrusService) {
	r = &ZoneSettingOriginErrorPagePassThrusService{}
	r.Options = opts
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *ZoneSettingOriginErrorPagePassThrusService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingOriginErrorPagePassThrusUpdateParams, opts ...option.RequestOption) (res *ZoneSettingOriginErrorPagePassThrusUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *ZoneSettingOriginErrorPagePassThrusService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingOriginErrorPagePassThrusListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingOriginErrorPagePassThrusUpdateResponse struct {
	Errors   []ZoneSettingOriginErrorPagePassThrusUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingOriginErrorPagePassThrusUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                    `json:"success,required"`
	Result  ZoneSettingOriginErrorPagePassThrusUpdateResponseResult `json:"result"`
	JSON    zoneSettingOriginErrorPagePassThrusUpdateResponseJSON   `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusUpdateResponseJSON contains the JSON metadata
// for the struct [ZoneSettingOriginErrorPagePassThrusUpdateResponse]
type zoneSettingOriginErrorPagePassThrusUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThrusUpdateResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThrusUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusUpdateResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingOriginErrorPagePassThrusUpdateResponseError]
type zoneSettingOriginErrorPagePassThrusUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThrusUpdateResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThrusUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusUpdateResponseMessageJSON contains the JSON
// metadata for the struct
// [ZoneSettingOriginErrorPagePassThrusUpdateResponseMessage]
type zoneSettingOriginErrorPagePassThrusUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThrusUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOriginErrorPagePassThrusUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingOriginErrorPagePassThrusUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOriginErrorPagePassThrusUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginErrorPagePassThrusUpdateResponseResultJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusUpdateResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneSettingOriginErrorPagePassThrusUpdateResponseResult]
type zoneSettingOriginErrorPagePassThrusUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOriginErrorPagePassThrusUpdateResponseResultID string

const (
	ZoneSettingOriginErrorPagePassThrusUpdateResponseResultIDOriginErrorPagePassThru ZoneSettingOriginErrorPagePassThrusUpdateResponseResultID = "origin_error_page_pass_thru"
)

// Value of the zone setting.
type ZoneSettingOriginErrorPagePassThrusUpdateResponseResultValue string

const (
	ZoneSettingOriginErrorPagePassThrusUpdateResponseResultValueOn  ZoneSettingOriginErrorPagePassThrusUpdateResponseResultValue = "on"
	ZoneSettingOriginErrorPagePassThrusUpdateResponseResultValueOff ZoneSettingOriginErrorPagePassThrusUpdateResponseResultValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOriginErrorPagePassThrusUpdateResponseResultEditable bool

const (
	ZoneSettingOriginErrorPagePassThrusUpdateResponseResultEditableTrue  ZoneSettingOriginErrorPagePassThrusUpdateResponseResultEditable = true
	ZoneSettingOriginErrorPagePassThrusUpdateResponseResultEditableFalse ZoneSettingOriginErrorPagePassThrusUpdateResponseResultEditable = false
)

type ZoneSettingOriginErrorPagePassThrusListResponse struct {
	Errors   []ZoneSettingOriginErrorPagePassThrusListResponseError   `json:"errors,required"`
	Messages []ZoneSettingOriginErrorPagePassThrusListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                                  `json:"success,required"`
	Result  ZoneSettingOriginErrorPagePassThrusListResponseResult `json:"result"`
	JSON    zoneSettingOriginErrorPagePassThrusListResponseJSON   `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusListResponseJSON contains the JSON metadata
// for the struct [ZoneSettingOriginErrorPagePassThrusListResponse]
type zoneSettingOriginErrorPagePassThrusListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThrusListResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThrusListResponseErrorJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusListResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingOriginErrorPagePassThrusListResponseError]
type zoneSettingOriginErrorPagePassThrusListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThrusListResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingOriginErrorPagePassThrusListResponseMessageJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingOriginErrorPagePassThrusListResponseMessage]
type zoneSettingOriginErrorPagePassThrusListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginErrorPagePassThrusListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOriginErrorPagePassThrusListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingOriginErrorPagePassThrusListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOriginErrorPagePassThrusListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginErrorPagePassThrusListResponseResultJSON `json:"-"`
}

// zoneSettingOriginErrorPagePassThrusListResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingOriginErrorPagePassThrusListResponseResult]
type zoneSettingOriginErrorPagePassThrusListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginErrorPagePassThrusListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOriginErrorPagePassThrusListResponseResultID string

const (
	ZoneSettingOriginErrorPagePassThrusListResponseResultIDOriginErrorPagePassThru ZoneSettingOriginErrorPagePassThrusListResponseResultID = "origin_error_page_pass_thru"
)

// Value of the zone setting.
type ZoneSettingOriginErrorPagePassThrusListResponseResultValue string

const (
	ZoneSettingOriginErrorPagePassThrusListResponseResultValueOn  ZoneSettingOriginErrorPagePassThrusListResponseResultValue = "on"
	ZoneSettingOriginErrorPagePassThrusListResponseResultValueOff ZoneSettingOriginErrorPagePassThrusListResponseResultValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOriginErrorPagePassThrusListResponseResultEditable bool

const (
	ZoneSettingOriginErrorPagePassThrusListResponseResultEditableTrue  ZoneSettingOriginErrorPagePassThrusListResponseResultEditable = true
	ZoneSettingOriginErrorPagePassThrusListResponseResultEditableFalse ZoneSettingOriginErrorPagePassThrusListResponseResultEditable = false
)

type ZoneSettingOriginErrorPagePassThrusUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingOriginErrorPagePassThrusUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingOriginErrorPagePassThrusUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingOriginErrorPagePassThrusUpdateParamsValue string

const (
	ZoneSettingOriginErrorPagePassThrusUpdateParamsValueOn  ZoneSettingOriginErrorPagePassThrusUpdateParamsValue = "on"
	ZoneSettingOriginErrorPagePassThrusUpdateParamsValueOff ZoneSettingOriginErrorPagePassThrusUpdateParamsValue = "off"
)
