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

// SettingSortQueryStringForCacheService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingSortQueryStringForCacheService] method instead.
type SettingSortQueryStringForCacheService struct {
	Options []option.RequestOption
}

// NewSettingSortQueryStringForCacheService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingSortQueryStringForCacheService(opts ...option.RequestOption) (r *SettingSortQueryStringForCacheService) {
	r = &SettingSortQueryStringForCacheService{}
	r.Options = opts
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *SettingSortQueryStringForCacheService) Edit(ctx context.Context, params SettingSortQueryStringForCacheEditParams, opts ...option.RequestOption) (res *ZoneSettingSortQueryStringForCache, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSortQueryStringForCacheEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *SettingSortQueryStringForCacheService) Get(ctx context.Context, query SettingSortQueryStringForCacheGetParams, opts ...option.RequestOption) (res *ZoneSettingSortQueryStringForCache, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSortQueryStringForCacheGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZoneSettingSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSortQueryStringForCacheJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheJSON contains the JSON metadata for the struct
// [ZoneSettingSortQueryStringForCache]
type zoneSettingSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingSortQueryStringForCache) implementsZonesSettingEditResponse() {}

func (r ZoneSettingSortQueryStringForCache) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingSortQueryStringForCacheID string

const (
	ZoneSettingSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingSortQueryStringForCacheID = "sort_query_string_for_cache"
)

func (r ZoneSettingSortQueryStringForCacheID) IsKnown() bool {
	switch r {
	case ZoneSettingSortQueryStringForCacheIDSortQueryStringForCache:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingSortQueryStringForCacheValue string

const (
	ZoneSettingSortQueryStringForCacheValueOn  ZoneSettingSortQueryStringForCacheValue = "on"
	ZoneSettingSortQueryStringForCacheValueOff ZoneSettingSortQueryStringForCacheValue = "off"
)

func (r ZoneSettingSortQueryStringForCacheValue) IsKnown() bool {
	switch r {
	case ZoneSettingSortQueryStringForCacheValueOn, ZoneSettingSortQueryStringForCacheValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSortQueryStringForCacheEditable bool

const (
	ZoneSettingSortQueryStringForCacheEditableTrue  ZoneSettingSortQueryStringForCacheEditable = true
	ZoneSettingSortQueryStringForCacheEditableFalse ZoneSettingSortQueryStringForCacheEditable = false
)

func (r ZoneSettingSortQueryStringForCacheEditable) IsKnown() bool {
	switch r {
	case ZoneSettingSortQueryStringForCacheEditableTrue, ZoneSettingSortQueryStringForCacheEditableFalse:
		return true
	}
	return false
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingSortQueryStringForCacheParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingSortQueryStringForCacheID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingSortQueryStringForCacheValue] `json:"value,required"`
}

func (r ZoneSettingSortQueryStringForCacheParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingSortQueryStringForCacheParam) implementsZonesSettingEditParamsItem() {}

type SettingSortQueryStringForCacheEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingSortQueryStringForCacheEditParamsValue] `json:"value,required"`
}

func (r SettingSortQueryStringForCacheEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingSortQueryStringForCacheEditParamsValue string

const (
	SettingSortQueryStringForCacheEditParamsValueOn  SettingSortQueryStringForCacheEditParamsValue = "on"
	SettingSortQueryStringForCacheEditParamsValueOff SettingSortQueryStringForCacheEditParamsValue = "off"
)

func (r SettingSortQueryStringForCacheEditParamsValue) IsKnown() bool {
	switch r {
	case SettingSortQueryStringForCacheEditParamsValueOn, SettingSortQueryStringForCacheEditParamsValueOff:
		return true
	}
	return false
}

type SettingSortQueryStringForCacheEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result ZoneSettingSortQueryStringForCache                     `json:"result"`
	JSON   settingSortQueryStringForCacheEditResponseEnvelopeJSON `json:"-"`
}

// settingSortQueryStringForCacheEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingSortQueryStringForCacheEditResponseEnvelope]
type settingSortQueryStringForCacheEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSortQueryStringForCacheEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSortQueryStringForCacheGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSortQueryStringForCacheGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result ZoneSettingSortQueryStringForCache                    `json:"result"`
	JSON   settingSortQueryStringForCacheGetResponseEnvelopeJSON `json:"-"`
}

// settingSortQueryStringForCacheGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingSortQueryStringForCacheGetResponseEnvelope]
type settingSortQueryStringForCacheGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSortQueryStringForCacheGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
