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

// ZoneSettingCacheLevelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingCacheLevelService]
// method instead.
type ZoneSettingCacheLevelService struct {
	Options []option.RequestOption
}

// NewZoneSettingCacheLevelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingCacheLevelService(opts ...option.RequestOption) (r *ZoneSettingCacheLevelService) {
	r = &ZoneSettingCacheLevelService{}
	r.Options = opts
	return
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
func (r *ZoneSettingCacheLevelService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingCacheLevelUpdateParams, opts ...option.RequestOption) (res *ZoneSettingCacheLevelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/cache_level", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
func (r *ZoneSettingCacheLevelService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingCacheLevelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/cache_level", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingCacheLevelUpdateResponse struct {
	Errors   []ZoneSettingCacheLevelUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingCacheLevelUpdateResponseMessage `json:"messages"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result ZoneSettingCacheLevelUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingCacheLevelUpdateResponseJSON
}

// zoneSettingCacheLevelUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelUpdateResponse]
type zoneSettingCacheLevelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingCacheLevelUpdateResponseErrorJSON
}

// zoneSettingCacheLevelUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelUpdateResponseError]
type zoneSettingCacheLevelUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingCacheLevelUpdateResponseMessageJSON
}

// zoneSettingCacheLevelUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingCacheLevelUpdateResponseMessage]
type zoneSettingCacheLevelUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingCacheLevelUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingCacheLevelUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingCacheLevelUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingCacheLevelUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingCacheLevelUpdateResponseResultJSON
}

// zoneSettingCacheLevelUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelUpdateResponseResult]
type zoneSettingCacheLevelUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingCacheLevelUpdateResponseResultID string

const (
	ZoneSettingCacheLevelUpdateResponseResultIDCacheLevel ZoneSettingCacheLevelUpdateResponseResultID = "cache_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingCacheLevelUpdateResponseResultEditable bool

const (
	ZoneSettingCacheLevelUpdateResponseResultEditableTrue  ZoneSettingCacheLevelUpdateResponseResultEditable = true
	ZoneSettingCacheLevelUpdateResponseResultEditableFalse ZoneSettingCacheLevelUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingCacheLevelUpdateResponseResultValue string

const (
	ZoneSettingCacheLevelUpdateResponseResultValueAggressive ZoneSettingCacheLevelUpdateResponseResultValue = "aggressive"
	ZoneSettingCacheLevelUpdateResponseResultValueBasic      ZoneSettingCacheLevelUpdateResponseResultValue = "basic"
	ZoneSettingCacheLevelUpdateResponseResultValueSimplified ZoneSettingCacheLevelUpdateResponseResultValue = "simplified"
)

type ZoneSettingCacheLevelListResponse struct {
	Errors   []ZoneSettingCacheLevelListResponseError   `json:"errors"`
	Messages []ZoneSettingCacheLevelListResponseMessage `json:"messages"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result ZoneSettingCacheLevelListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingCacheLevelListResponseJSON
}

// zoneSettingCacheLevelListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingCacheLevelListResponse]
type zoneSettingCacheLevelListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingCacheLevelListResponseErrorJSON
}

// zoneSettingCacheLevelListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelListResponseError]
type zoneSettingCacheLevelListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingCacheLevelListResponseMessageJSON
}

// zoneSettingCacheLevelListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelListResponseMessage]
type zoneSettingCacheLevelListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingCacheLevelListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingCacheLevelListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingCacheLevelListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingCacheLevelListResponseResultValue `json:"value"`
	JSON  zoneSettingCacheLevelListResponseResultJSON
}

// zoneSettingCacheLevelListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelListResponseResult]
type zoneSettingCacheLevelListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingCacheLevelListResponseResultID string

const (
	ZoneSettingCacheLevelListResponseResultIDCacheLevel ZoneSettingCacheLevelListResponseResultID = "cache_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingCacheLevelListResponseResultEditable bool

const (
	ZoneSettingCacheLevelListResponseResultEditableTrue  ZoneSettingCacheLevelListResponseResultEditable = true
	ZoneSettingCacheLevelListResponseResultEditableFalse ZoneSettingCacheLevelListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingCacheLevelListResponseResultValue string

const (
	ZoneSettingCacheLevelListResponseResultValueAggressive ZoneSettingCacheLevelListResponseResultValue = "aggressive"
	ZoneSettingCacheLevelListResponseResultValueBasic      ZoneSettingCacheLevelListResponseResultValue = "basic"
	ZoneSettingCacheLevelListResponseResultValueSimplified ZoneSettingCacheLevelListResponseResultValue = "simplified"
)

type ZoneSettingCacheLevelUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingCacheLevelUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingCacheLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingCacheLevelUpdateParamsValue string

const (
	ZoneSettingCacheLevelUpdateParamsValueAggressive ZoneSettingCacheLevelUpdateParamsValue = "aggressive"
	ZoneSettingCacheLevelUpdateParamsValueBasic      ZoneSettingCacheLevelUpdateParamsValue = "basic"
	ZoneSettingCacheLevelUpdateParamsValueSimplified ZoneSettingCacheLevelUpdateParamsValue = "simplified"
)
