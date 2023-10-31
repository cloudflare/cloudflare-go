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

// ZoneSettingSecurityLevelService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingSecurityLevelService] method instead.
type ZoneSettingSecurityLevelService struct {
	Options []option.RequestOption
}

// NewZoneSettingSecurityLevelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSecurityLevelService(opts ...option.RequestOption) (r *ZoneSettingSecurityLevelService) {
	r = &ZoneSettingSecurityLevelService{}
	r.Options = opts
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *ZoneSettingSecurityLevelService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingSecurityLevelUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSecurityLevelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_level", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *ZoneSettingSecurityLevelService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingSecurityLevelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_level", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingSecurityLevelUpdateResponse struct {
	Errors   []ZoneSettingSecurityLevelUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingSecurityLevelUpdateResponseMessage `json:"messages"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result ZoneSettingSecurityLevelUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingSecurityLevelUpdateResponseJSON
}

// zoneSettingSecurityLevelUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityLevelUpdateResponse]
type zoneSettingSecurityLevelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSecurityLevelUpdateResponseErrorJSON
}

// zoneSettingSecurityLevelUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityLevelUpdateResponseError]
type zoneSettingSecurityLevelUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSecurityLevelUpdateResponseMessageJSON
}

// zoneSettingSecurityLevelUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityLevelUpdateResponseMessage]
type zoneSettingSecurityLevelUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingSecurityLevelUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSecurityLevelUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityLevelUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingSecurityLevelUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingSecurityLevelUpdateResponseResultJSON
}

// zoneSettingSecurityLevelUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityLevelUpdateResponseResult]
type zoneSettingSecurityLevelUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSecurityLevelUpdateResponseResultID string

const (
	ZoneSettingSecurityLevelUpdateResponseResultIDSecurityLevel ZoneSettingSecurityLevelUpdateResponseResultID = "security_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityLevelUpdateResponseResultEditable bool

const (
	ZoneSettingSecurityLevelUpdateResponseResultEditableTrue  ZoneSettingSecurityLevelUpdateResponseResultEditable = true
	ZoneSettingSecurityLevelUpdateResponseResultEditableFalse ZoneSettingSecurityLevelUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingSecurityLevelUpdateResponseResultValue string

const (
	ZoneSettingSecurityLevelUpdateResponseResultValueOff            ZoneSettingSecurityLevelUpdateResponseResultValue = "off"
	ZoneSettingSecurityLevelUpdateResponseResultValueEssentiallyOff ZoneSettingSecurityLevelUpdateResponseResultValue = "essentially_off"
	ZoneSettingSecurityLevelUpdateResponseResultValueLow            ZoneSettingSecurityLevelUpdateResponseResultValue = "low"
	ZoneSettingSecurityLevelUpdateResponseResultValueMedium         ZoneSettingSecurityLevelUpdateResponseResultValue = "medium"
	ZoneSettingSecurityLevelUpdateResponseResultValueHigh           ZoneSettingSecurityLevelUpdateResponseResultValue = "high"
	ZoneSettingSecurityLevelUpdateResponseResultValueUnderAttack    ZoneSettingSecurityLevelUpdateResponseResultValue = "under_attack"
)

type ZoneSettingSecurityLevelListResponse struct {
	Errors   []ZoneSettingSecurityLevelListResponseError   `json:"errors"`
	Messages []ZoneSettingSecurityLevelListResponseMessage `json:"messages"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result ZoneSettingSecurityLevelListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingSecurityLevelListResponseJSON
}

// zoneSettingSecurityLevelListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityLevelListResponse]
type zoneSettingSecurityLevelListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSecurityLevelListResponseErrorJSON
}

// zoneSettingSecurityLevelListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityLevelListResponseError]
type zoneSettingSecurityLevelListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityLevelListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSecurityLevelListResponseMessageJSON
}

// zoneSettingSecurityLevelListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityLevelListResponseMessage]
type zoneSettingSecurityLevelListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingSecurityLevelListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSecurityLevelListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityLevelListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingSecurityLevelListResponseResultValue `json:"value"`
	JSON  zoneSettingSecurityLevelListResponseResultJSON
}

// zoneSettingSecurityLevelListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityLevelListResponseResult]
type zoneSettingSecurityLevelListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityLevelListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSecurityLevelListResponseResultID string

const (
	ZoneSettingSecurityLevelListResponseResultIDSecurityLevel ZoneSettingSecurityLevelListResponseResultID = "security_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityLevelListResponseResultEditable bool

const (
	ZoneSettingSecurityLevelListResponseResultEditableTrue  ZoneSettingSecurityLevelListResponseResultEditable = true
	ZoneSettingSecurityLevelListResponseResultEditableFalse ZoneSettingSecurityLevelListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingSecurityLevelListResponseResultValue string

const (
	ZoneSettingSecurityLevelListResponseResultValueOff            ZoneSettingSecurityLevelListResponseResultValue = "off"
	ZoneSettingSecurityLevelListResponseResultValueEssentiallyOff ZoneSettingSecurityLevelListResponseResultValue = "essentially_off"
	ZoneSettingSecurityLevelListResponseResultValueLow            ZoneSettingSecurityLevelListResponseResultValue = "low"
	ZoneSettingSecurityLevelListResponseResultValueMedium         ZoneSettingSecurityLevelListResponseResultValue = "medium"
	ZoneSettingSecurityLevelListResponseResultValueHigh           ZoneSettingSecurityLevelListResponseResultValue = "high"
	ZoneSettingSecurityLevelListResponseResultValueUnderAttack    ZoneSettingSecurityLevelListResponseResultValue = "under_attack"
)

type ZoneSettingSecurityLevelUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingSecurityLevelUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSecurityLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingSecurityLevelUpdateParamsValue string

const (
	ZoneSettingSecurityLevelUpdateParamsValueOff            ZoneSettingSecurityLevelUpdateParamsValue = "off"
	ZoneSettingSecurityLevelUpdateParamsValueEssentiallyOff ZoneSettingSecurityLevelUpdateParamsValue = "essentially_off"
	ZoneSettingSecurityLevelUpdateParamsValueLow            ZoneSettingSecurityLevelUpdateParamsValue = "low"
	ZoneSettingSecurityLevelUpdateParamsValueMedium         ZoneSettingSecurityLevelUpdateParamsValue = "medium"
	ZoneSettingSecurityLevelUpdateParamsValueHigh           ZoneSettingSecurityLevelUpdateParamsValue = "high"
	ZoneSettingSecurityLevelUpdateParamsValueUnderAttack    ZoneSettingSecurityLevelUpdateParamsValue = "under_attack"
)
