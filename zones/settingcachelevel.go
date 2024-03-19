// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingCacheLevelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingCacheLevelService] method
// instead.
type SettingCacheLevelService struct {
	Options []option.RequestOption
}

// NewSettingCacheLevelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingCacheLevelService(opts ...option.RequestOption) (r *SettingCacheLevelService) {
	r = &SettingCacheLevelService{}
	r.Options = opts
	return
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
func (r *SettingCacheLevelService) Edit(ctx context.Context, params SettingCacheLevelEditParams, opts ...option.RequestOption) (res *ZonesCacheLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCacheLevelEditResponseEnvelope
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
func (r *SettingCacheLevelService) Get(ctx context.Context, query SettingCacheLevelGetParams, opts ...option.RequestOption) (res *ZonesCacheLevel, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCacheLevelGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/cache_level", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r zonesCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZonesCacheLevel) implementsZonesSettingEditResponse() {}

func (r ZonesCacheLevel) implementsZonesSettingGetResponse() {}

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

func (r ZonesCacheLevelParam) implementsZonesSettingEditParamsItem() {}

type SettingCacheLevelEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingCacheLevelEditParamsValue] `json:"value,required"`
}

func (r SettingCacheLevelEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingCacheLevelEditParamsValue string

const (
	SettingCacheLevelEditParamsValueAggressive SettingCacheLevelEditParamsValue = "aggressive"
	SettingCacheLevelEditParamsValueBasic      SettingCacheLevelEditParamsValue = "basic"
	SettingCacheLevelEditParamsValueSimplified SettingCacheLevelEditParamsValue = "simplified"
)

type SettingCacheLevelEditResponseEnvelope struct {
	Errors   []SettingCacheLevelEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingCacheLevelEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result ZonesCacheLevel                           `json:"result"`
	JSON   settingCacheLevelEditResponseEnvelopeJSON `json:"-"`
}

// settingCacheLevelEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingCacheLevelEditResponseEnvelope]
type settingCacheLevelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingCacheLevelEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingCacheLevelEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingCacheLevelEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingCacheLevelEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingCacheLevelEditResponseEnvelopeErrors]
type settingCacheLevelEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingCacheLevelEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingCacheLevelEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingCacheLevelEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingCacheLevelEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingCacheLevelEditResponseEnvelopeMessages]
type settingCacheLevelEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingCacheLevelEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingCacheLevelGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingCacheLevelGetResponseEnvelope struct {
	Errors   []SettingCacheLevelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingCacheLevelGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result ZonesCacheLevel                          `json:"result"`
	JSON   settingCacheLevelGetResponseEnvelopeJSON `json:"-"`
}

// settingCacheLevelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingCacheLevelGetResponseEnvelope]
type settingCacheLevelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingCacheLevelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingCacheLevelGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingCacheLevelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingCacheLevelGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingCacheLevelGetResponseEnvelopeErrors]
type settingCacheLevelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingCacheLevelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingCacheLevelGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingCacheLevelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingCacheLevelGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingCacheLevelGetResponseEnvelopeMessages]
type settingCacheLevelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingCacheLevelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
