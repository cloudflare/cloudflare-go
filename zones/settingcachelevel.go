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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *SettingCacheLevelService) Edit(ctx context.Context, params SettingCacheLevelEditParams, opts ...option.RequestOption) (res *CacheLevel, err error) {
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
func (r *SettingCacheLevelService) Get(ctx context.Context, query SettingCacheLevelGetParams, opts ...option.RequestOption) (res *CacheLevel, err error) {
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
type CacheLevel struct {
	// ID of the zone setting.
	ID CacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value CacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       cacheLevelJSON `json:"-"`
}

// cacheLevelJSON contains the JSON metadata for the struct [CacheLevel]
type cacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheLevelJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type CacheLevelID string

const (
	CacheLevelIDCacheLevel CacheLevelID = "cache_level"
)

func (r CacheLevelID) IsKnown() bool {
	switch r {
	case CacheLevelIDCacheLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type CacheLevelValue string

const (
	CacheLevelValueAggressive CacheLevelValue = "aggressive"
	CacheLevelValueBasic      CacheLevelValue = "basic"
	CacheLevelValueSimplified CacheLevelValue = "simplified"
)

func (r CacheLevelValue) IsKnown() bool {
	switch r {
	case CacheLevelValueAggressive, CacheLevelValueBasic, CacheLevelValueSimplified:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type CacheLevelEditable bool

const (
	CacheLevelEditableTrue  CacheLevelEditable = true
	CacheLevelEditableFalse CacheLevelEditable = false
)

func (r CacheLevelEditable) IsKnown() bool {
	switch r {
	case CacheLevelEditableTrue, CacheLevelEditableFalse:
		return true
	}
	return false
}

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

func (r SettingCacheLevelEditParamsValue) IsKnown() bool {
	switch r {
	case SettingCacheLevelEditParamsValueAggressive, SettingCacheLevelEditParamsValueBasic, SettingCacheLevelEditParamsValueSimplified:
		return true
	}
	return false
}

type SettingCacheLevelEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result CacheLevel                                `json:"result"`
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

type SettingCacheLevelGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingCacheLevelGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result CacheLevel                               `json:"result"`
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
