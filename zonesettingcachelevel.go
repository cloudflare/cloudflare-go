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
func (r *ZoneSettingCacheLevelService) Edit(ctx context.Context, params ZoneSettingCacheLevelEditParams, opts ...option.RequestOption) (res *ZonesCacheLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingCacheLevelEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/cache_level", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
func (r *ZoneSettingCacheLevelService) Get(ctx context.Context, query ZoneSettingCacheLevelGetParams, opts ...option.RequestOption) (res *ZonesCacheLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingCacheLevelGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/cache_level", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZonesCacheLevel struct {
	// ID of the zone setting.
	ID ZonesCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesCacheLevelJSON `json:"-"`
}

// zonesCacheLevelJSON contains the JSON metadata for the struct [ZonesCacheLevel]
type zonesCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesCacheLevel) implementsZoneSettingEditResponse() {}

func (r ZonesCacheLevel) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesCacheLevelID string

const (
	ZonesCacheLevelIDCacheLevel ZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type ZonesCacheLevelValue string

const (
	ZonesCacheLevelValueAggressive ZonesCacheLevelValue = "aggressive"
	ZonesCacheLevelValueBasic      ZonesCacheLevelValue = "basic"
	ZonesCacheLevelValueSimplified ZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesCacheLevelEditable bool

const (
	ZonesCacheLevelEditableTrue  ZonesCacheLevelEditable = true
	ZonesCacheLevelEditableFalse ZonesCacheLevelEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZonesCacheLevelParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesCacheLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesCacheLevelValue] `json:"value,required"`
}

func (r ZonesCacheLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesCacheLevelParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingCacheLevelEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingCacheLevelEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingCacheLevelEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingCacheLevelEditParamsValue string

const (
	ZoneSettingCacheLevelEditParamsValueAggressive ZoneSettingCacheLevelEditParamsValue = "aggressive"
	ZoneSettingCacheLevelEditParamsValueBasic      ZoneSettingCacheLevelEditParamsValue = "basic"
	ZoneSettingCacheLevelEditParamsValueSimplified ZoneSettingCacheLevelEditParamsValue = "simplified"
)

type ZoneSettingCacheLevelEditResponseEnvelope struct {
	Errors   []ZoneSettingCacheLevelEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingCacheLevelEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result ZonesCacheLevel                               `json:"result"`
	JSON   zoneSettingCacheLevelEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingCacheLevelEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelEditResponseEnvelope]
type zoneSettingCacheLevelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingCacheLevelEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingCacheLevelEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingCacheLevelEditResponseEnvelopeErrors]
type zoneSettingCacheLevelEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingCacheLevelEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingCacheLevelEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingCacheLevelEditResponseEnvelopeMessages]
type zoneSettingCacheLevelEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingCacheLevelGetResponseEnvelope struct {
	Errors   []ZoneSettingCacheLevelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingCacheLevelGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result ZonesCacheLevel                              `json:"result"`
	JSON   zoneSettingCacheLevelGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingCacheLevelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingCacheLevelGetResponseEnvelope]
type zoneSettingCacheLevelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingCacheLevelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingCacheLevelGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingCacheLevelGetResponseEnvelopeErrors]
type zoneSettingCacheLevelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCacheLevelGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingCacheLevelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingCacheLevelGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingCacheLevelGetResponseEnvelopeMessages]
type zoneSettingCacheLevelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCacheLevelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
