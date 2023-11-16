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

// ZoneSettingTrueClientIPHeaderService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingTrueClientIPHeaderService] method instead.
type ZoneSettingTrueClientIPHeaderService struct {
	Options []option.RequestOption
}

// NewZoneSettingTrueClientIPHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingTrueClientIPHeaderService(opts ...option.RequestOption) (r *ZoneSettingTrueClientIPHeaderService) {
	r = &ZoneSettingTrueClientIPHeaderService{}
	r.Options = opts
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *ZoneSettingTrueClientIPHeaderService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingTrueClientIPHeaderUpdateParams, opts ...option.RequestOption) (res *ZoneSettingTrueClientIPHeaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *ZoneSettingTrueClientIPHeaderService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingTrueClientIPHeaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingTrueClientIPHeaderUpdateResponse struct {
	Errors   []ZoneSettingTrueClientIPHeaderUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingTrueClientIPHeaderUpdateResponseMessage `json:"messages"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result ZoneSettingTrueClientIPHeaderUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                            `json:"success"`
	JSON    zoneSettingTrueClientIPHeaderUpdateResponseJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingTrueClientIPHeaderUpdateResponse]
type zoneSettingTrueClientIPHeaderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderUpdateResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingTrueClientIPHeaderUpdateResponseError]
type zoneSettingTrueClientIPHeaderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderUpdateResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderUpdateResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingTrueClientIPHeaderUpdateResponseMessage]
type zoneSettingTrueClientIPHeaderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingTrueClientIPHeaderUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTrueClientIPHeaderUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTrueClientIPHeaderUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingTrueClientIPHeaderUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingTrueClientIPHeaderUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingTrueClientIPHeaderUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingTrueClientIPHeaderUpdateResponseResult]
type zoneSettingTrueClientIPHeaderUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTrueClientIPHeaderUpdateResponseResultID string

const (
	ZoneSettingTrueClientIPHeaderUpdateResponseResultIDTrueClientIPHeader ZoneSettingTrueClientIPHeaderUpdateResponseResultID = "true_client_ip_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTrueClientIPHeaderUpdateResponseResultEditable bool

const (
	ZoneSettingTrueClientIPHeaderUpdateResponseResultEditableTrue  ZoneSettingTrueClientIPHeaderUpdateResponseResultEditable = true
	ZoneSettingTrueClientIPHeaderUpdateResponseResultEditableFalse ZoneSettingTrueClientIPHeaderUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingTrueClientIPHeaderUpdateResponseResultValue string

const (
	ZoneSettingTrueClientIPHeaderUpdateResponseResultValueOn  ZoneSettingTrueClientIPHeaderUpdateResponseResultValue = "on"
	ZoneSettingTrueClientIPHeaderUpdateResponseResultValueOff ZoneSettingTrueClientIPHeaderUpdateResponseResultValue = "off"
)

type ZoneSettingTrueClientIPHeaderListResponse struct {
	Errors   []ZoneSettingTrueClientIPHeaderListResponseError   `json:"errors"`
	Messages []ZoneSettingTrueClientIPHeaderListResponseMessage `json:"messages"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result ZoneSettingTrueClientIPHeaderListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                          `json:"success"`
	JSON    zoneSettingTrueClientIPHeaderListResponseJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingTrueClientIPHeaderListResponse]
type zoneSettingTrueClientIPHeaderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderListResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderListResponseErrorJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderListResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingTrueClientIPHeaderListResponseError]
type zoneSettingTrueClientIPHeaderListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTrueClientIPHeaderListResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingTrueClientIPHeaderListResponseMessageJSON `json:"-"`
}

// zoneSettingTrueClientIPHeaderListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingTrueClientIPHeaderListResponseMessage]
type zoneSettingTrueClientIPHeaderListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingTrueClientIPHeaderListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTrueClientIPHeaderListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTrueClientIPHeaderListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingTrueClientIPHeaderListResponseResultValue `json:"value"`
	JSON  zoneSettingTrueClientIPHeaderListResponseResultJSON  `json:"-"`
}

// zoneSettingTrueClientIPHeaderListResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingTrueClientIPHeaderListResponseResult]
type zoneSettingTrueClientIPHeaderListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTrueClientIPHeaderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTrueClientIPHeaderListResponseResultID string

const (
	ZoneSettingTrueClientIPHeaderListResponseResultIDTrueClientIPHeader ZoneSettingTrueClientIPHeaderListResponseResultID = "true_client_ip_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTrueClientIPHeaderListResponseResultEditable bool

const (
	ZoneSettingTrueClientIPHeaderListResponseResultEditableTrue  ZoneSettingTrueClientIPHeaderListResponseResultEditable = true
	ZoneSettingTrueClientIPHeaderListResponseResultEditableFalse ZoneSettingTrueClientIPHeaderListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingTrueClientIPHeaderListResponseResultValue string

const (
	ZoneSettingTrueClientIPHeaderListResponseResultValueOn  ZoneSettingTrueClientIPHeaderListResponseResultValue = "on"
	ZoneSettingTrueClientIPHeaderListResponseResultValueOff ZoneSettingTrueClientIPHeaderListResponseResultValue = "off"
)

type ZoneSettingTrueClientIPHeaderUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingTrueClientIPHeaderUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingTrueClientIPHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingTrueClientIPHeaderUpdateParamsValue string

const (
	ZoneSettingTrueClientIPHeaderUpdateParamsValueOn  ZoneSettingTrueClientIPHeaderUpdateParamsValue = "on"
	ZoneSettingTrueClientIPHeaderUpdateParamsValueOff ZoneSettingTrueClientIPHeaderUpdateParamsValue = "off"
)
