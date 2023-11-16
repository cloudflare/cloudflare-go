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

// ZoneSettingOrangeToOrangeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingOrangeToOrangeService] method instead.
type ZoneSettingOrangeToOrangeService struct {
	Options []option.RequestOption
}

// NewZoneSettingOrangeToOrangeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingOrangeToOrangeService(opts ...option.RequestOption) (r *ZoneSettingOrangeToOrangeService) {
	r = &ZoneSettingOrangeToOrangeService{}
	r.Options = opts
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *ZoneSettingOrangeToOrangeService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingOrangeToOrangeUpdateParams, opts ...option.RequestOption) (res *ZoneSettingOrangeToOrangeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *ZoneSettingOrangeToOrangeService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingOrangeToOrangeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingOrangeToOrangeUpdateResponse struct {
	Errors   []ZoneSettingOrangeToOrangeUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingOrangeToOrangeUpdateResponseMessage `json:"messages"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZoneSettingOrangeToOrangeUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                        `json:"success"`
	JSON    zoneSettingOrangeToOrangeUpdateResponseJSON `json:"-"`
}

// zoneSettingOrangeToOrangeUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingOrangeToOrangeUpdateResponse]
type zoneSettingOrangeToOrangeUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeUpdateResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingOrangeToOrangeUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingOrangeToOrangeUpdateResponseError]
type zoneSettingOrangeToOrangeUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeUpdateResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingOrangeToOrangeUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingOrangeToOrangeUpdateResponseMessage]
type zoneSettingOrangeToOrangeUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingOrangeToOrangeUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOrangeToOrangeUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOrangeToOrangeUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingOrangeToOrangeUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingOrangeToOrangeUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingOrangeToOrangeUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingOrangeToOrangeUpdateResponseResult]
type zoneSettingOrangeToOrangeUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOrangeToOrangeUpdateResponseResultID string

const (
	ZoneSettingOrangeToOrangeUpdateResponseResultIDOrangeToOrange ZoneSettingOrangeToOrangeUpdateResponseResultID = "orange_to_orange"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOrangeToOrangeUpdateResponseResultEditable bool

const (
	ZoneSettingOrangeToOrangeUpdateResponseResultEditableTrue  ZoneSettingOrangeToOrangeUpdateResponseResultEditable = true
	ZoneSettingOrangeToOrangeUpdateResponseResultEditableFalse ZoneSettingOrangeToOrangeUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingOrangeToOrangeUpdateResponseResultValue string

const (
	ZoneSettingOrangeToOrangeUpdateResponseResultValueOn  ZoneSettingOrangeToOrangeUpdateResponseResultValue = "on"
	ZoneSettingOrangeToOrangeUpdateResponseResultValueOff ZoneSettingOrangeToOrangeUpdateResponseResultValue = "off"
)

type ZoneSettingOrangeToOrangeListResponse struct {
	Errors   []ZoneSettingOrangeToOrangeListResponseError   `json:"errors"`
	Messages []ZoneSettingOrangeToOrangeListResponseMessage `json:"messages"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZoneSettingOrangeToOrangeListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                      `json:"success"`
	JSON    zoneSettingOrangeToOrangeListResponseJSON `json:"-"`
}

// zoneSettingOrangeToOrangeListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingOrangeToOrangeListResponse]
type zoneSettingOrangeToOrangeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeListResponseErrorJSON `json:"-"`
}

// zoneSettingOrangeToOrangeListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingOrangeToOrangeListResponseError]
type zoneSettingOrangeToOrangeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOrangeToOrangeListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingOrangeToOrangeListResponseMessageJSON `json:"-"`
}

// zoneSettingOrangeToOrangeListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingOrangeToOrangeListResponseMessage]
type zoneSettingOrangeToOrangeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingOrangeToOrangeListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOrangeToOrangeListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOrangeToOrangeListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingOrangeToOrangeListResponseResultValue `json:"value"`
	JSON  zoneSettingOrangeToOrangeListResponseResultJSON  `json:"-"`
}

// zoneSettingOrangeToOrangeListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingOrangeToOrangeListResponseResult]
type zoneSettingOrangeToOrangeListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrangeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOrangeToOrangeListResponseResultID string

const (
	ZoneSettingOrangeToOrangeListResponseResultIDOrangeToOrange ZoneSettingOrangeToOrangeListResponseResultID = "orange_to_orange"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOrangeToOrangeListResponseResultEditable bool

const (
	ZoneSettingOrangeToOrangeListResponseResultEditableTrue  ZoneSettingOrangeToOrangeListResponseResultEditable = true
	ZoneSettingOrangeToOrangeListResponseResultEditableFalse ZoneSettingOrangeToOrangeListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingOrangeToOrangeListResponseResultValue string

const (
	ZoneSettingOrangeToOrangeListResponseResultValueOn  ZoneSettingOrangeToOrangeListResponseResultValue = "on"
	ZoneSettingOrangeToOrangeListResponseResultValueOff ZoneSettingOrangeToOrangeListResponseResultValue = "off"
)

type ZoneSettingOrangeToOrangeUpdateParams struct {
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Value param.Field[ZoneSettingOrangeToOrangeUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingOrangeToOrangeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingOrangeToOrangeUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingOrangeToOrangeUpdateParamsValueID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingOrangeToOrangeUpdateParamsValueValue] `json:"value"`
}

func (r ZoneSettingOrangeToOrangeUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingOrangeToOrangeUpdateParamsValueID string

const (
	ZoneSettingOrangeToOrangeUpdateParamsValueIDOrangeToOrange ZoneSettingOrangeToOrangeUpdateParamsValueID = "orange_to_orange"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOrangeToOrangeUpdateParamsValueEditable bool

const (
	ZoneSettingOrangeToOrangeUpdateParamsValueEditableTrue  ZoneSettingOrangeToOrangeUpdateParamsValueEditable = true
	ZoneSettingOrangeToOrangeUpdateParamsValueEditableFalse ZoneSettingOrangeToOrangeUpdateParamsValueEditable = false
)

// Value of the zone setting.
type ZoneSettingOrangeToOrangeUpdateParamsValueValue string

const (
	ZoneSettingOrangeToOrangeUpdateParamsValueValueOn  ZoneSettingOrangeToOrangeUpdateParamsValueValue = "on"
	ZoneSettingOrangeToOrangeUpdateParamsValueValueOff ZoneSettingOrangeToOrangeUpdateParamsValueValue = "off"
)
