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

// ZoneSettingH2PrioritizationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingH2PrioritizationService] method instead.
type ZoneSettingH2PrioritizationService struct {
	Options []option.RequestOption
}

// NewZoneSettingH2PrioritizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingH2PrioritizationService(opts ...option.RequestOption) (r *ZoneSettingH2PrioritizationService) {
	r = &ZoneSettingH2PrioritizationService{}
	r.Options = opts
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *ZoneSettingH2PrioritizationService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingH2PrioritizationUpdateParams, opts ...option.RequestOption) (res *ZoneSettingH2PrioritizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *ZoneSettingH2PrioritizationService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingH2PrioritizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingH2PrioritizationUpdateResponse struct {
	Errors   []ZoneSettingH2PrioritizationUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingH2PrioritizationUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                            `json:"success,required"`
	Result  ZoneSettingH2PrioritizationUpdateResponseResult `json:"result"`
	JSON    zoneSettingH2PrioritizationUpdateResponseJSON   `json:"-"`
}

// zoneSettingH2PrioritizationUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingH2PrioritizationUpdateResponse]
type zoneSettingH2PrioritizationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationUpdateResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingH2PrioritizationUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingH2PrioritizationUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingH2PrioritizationUpdateResponseError]
type zoneSettingH2PrioritizationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationUpdateResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingH2PrioritizationUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingH2PrioritizationUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingH2PrioritizationUpdateResponseMessage]
type zoneSettingH2PrioritizationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingH2PrioritizationUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingH2PrioritizationUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingH2PrioritizationUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingH2PrioritizationUpdateResponseResultJSON `json:"-"`
}

// zoneSettingH2PrioritizationUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingH2PrioritizationUpdateResponseResult]
type zoneSettingH2PrioritizationUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingH2PrioritizationUpdateResponseResultID string

const (
	ZoneSettingH2PrioritizationUpdateResponseResultIDH2Prioritization ZoneSettingH2PrioritizationUpdateResponseResultID = "h2_prioritization"
)

// Value of the zone setting.
type ZoneSettingH2PrioritizationUpdateResponseResultValue string

const (
	ZoneSettingH2PrioritizationUpdateResponseResultValueOn     ZoneSettingH2PrioritizationUpdateResponseResultValue = "on"
	ZoneSettingH2PrioritizationUpdateResponseResultValueOff    ZoneSettingH2PrioritizationUpdateResponseResultValue = "off"
	ZoneSettingH2PrioritizationUpdateResponseResultValueCustom ZoneSettingH2PrioritizationUpdateResponseResultValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingH2PrioritizationUpdateResponseResultEditable bool

const (
	ZoneSettingH2PrioritizationUpdateResponseResultEditableTrue  ZoneSettingH2PrioritizationUpdateResponseResultEditable = true
	ZoneSettingH2PrioritizationUpdateResponseResultEditableFalse ZoneSettingH2PrioritizationUpdateResponseResultEditable = false
)

type ZoneSettingH2PrioritizationListResponse struct {
	Errors   []ZoneSettingH2PrioritizationListResponseError   `json:"errors,required"`
	Messages []ZoneSettingH2PrioritizationListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                          `json:"success,required"`
	Result  ZoneSettingH2PrioritizationListResponseResult `json:"result"`
	JSON    zoneSettingH2PrioritizationListResponseJSON   `json:"-"`
}

// zoneSettingH2PrioritizationListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingH2PrioritizationListResponse]
type zoneSettingH2PrioritizationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationListResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingH2PrioritizationListResponseErrorJSON `json:"-"`
}

// zoneSettingH2PrioritizationListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingH2PrioritizationListResponseError]
type zoneSettingH2PrioritizationListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationListResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingH2PrioritizationListResponseMessageJSON `json:"-"`
}

// zoneSettingH2PrioritizationListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingH2PrioritizationListResponseMessage]
type zoneSettingH2PrioritizationListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingH2PrioritizationListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingH2PrioritizationListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingH2PrioritizationListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingH2PrioritizationListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingH2PrioritizationListResponseResultJSON `json:"-"`
}

// zoneSettingH2PrioritizationListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingH2PrioritizationListResponseResult]
type zoneSettingH2PrioritizationListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2PrioritizationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingH2PrioritizationListResponseResultID string

const (
	ZoneSettingH2PrioritizationListResponseResultIDH2Prioritization ZoneSettingH2PrioritizationListResponseResultID = "h2_prioritization"
)

// Value of the zone setting.
type ZoneSettingH2PrioritizationListResponseResultValue string

const (
	ZoneSettingH2PrioritizationListResponseResultValueOn     ZoneSettingH2PrioritizationListResponseResultValue = "on"
	ZoneSettingH2PrioritizationListResponseResultValueOff    ZoneSettingH2PrioritizationListResponseResultValue = "off"
	ZoneSettingH2PrioritizationListResponseResultValueCustom ZoneSettingH2PrioritizationListResponseResultValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingH2PrioritizationListResponseResultEditable bool

const (
	ZoneSettingH2PrioritizationListResponseResultEditableTrue  ZoneSettingH2PrioritizationListResponseResultEditable = true
	ZoneSettingH2PrioritizationListResponseResultEditableFalse ZoneSettingH2PrioritizationListResponseResultEditable = false
)

type ZoneSettingH2PrioritizationUpdateParams struct {
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value param.Field[ZoneSettingH2PrioritizationUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingH2PrioritizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingH2PrioritizationUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingH2PrioritizationUpdateParamsValueID] `json:"id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingH2PrioritizationUpdateParamsValueValue] `json:"value,required"`
}

func (r ZoneSettingH2PrioritizationUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingH2PrioritizationUpdateParamsValueID string

const (
	ZoneSettingH2PrioritizationUpdateParamsValueIDH2Prioritization ZoneSettingH2PrioritizationUpdateParamsValueID = "h2_prioritization"
)

// Value of the zone setting.
type ZoneSettingH2PrioritizationUpdateParamsValueValue string

const (
	ZoneSettingH2PrioritizationUpdateParamsValueValueOn     ZoneSettingH2PrioritizationUpdateParamsValueValue = "on"
	ZoneSettingH2PrioritizationUpdateParamsValueValueOff    ZoneSettingH2PrioritizationUpdateParamsValueValue = "off"
	ZoneSettingH2PrioritizationUpdateParamsValueValueCustom ZoneSettingH2PrioritizationUpdateParamsValueValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingH2PrioritizationUpdateParamsValueEditable bool

const (
	ZoneSettingH2PrioritizationUpdateParamsValueEditableTrue  ZoneSettingH2PrioritizationUpdateParamsValueEditable = true
	ZoneSettingH2PrioritizationUpdateParamsValueEditableFalse ZoneSettingH2PrioritizationUpdateParamsValueEditable = false
)
