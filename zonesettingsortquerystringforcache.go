// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZoneSettingSortQueryStringForCacheService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingSortQueryStringForCacheService] method instead.
type ZoneSettingSortQueryStringForCacheService struct {
	Options []option.RequestOption
}

// NewZoneSettingSortQueryStringForCacheService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSettingSortQueryStringForCacheService(opts ...option.RequestOption) (r *ZoneSettingSortQueryStringForCacheService) {
	r = &ZoneSettingSortQueryStringForCacheService{}
	r.Options = opts
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *ZoneSettingSortQueryStringForCacheService) Edit(ctx context.Context, params ZoneSettingSortQueryStringForCacheEditParams, opts ...option.RequestOption) (res *ZoneSettingSortQueryStringForCacheEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSortQueryStringForCacheEditResponseEnvelope
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
func (r *ZoneSettingSortQueryStringForCacheService) Get(ctx context.Context, query ZoneSettingSortQueryStringForCacheGetParams, opts ...option.RequestOption) (res *ZoneSettingSortQueryStringForCacheGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingSortQueryStringForCacheGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingSortQueryStringForCacheEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingSortQueryStringForCacheEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSortQueryStringForCacheEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSortQueryStringForCacheEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSortQueryStringForCacheEditResponseJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheEditResponseJSON contains the JSON metadata
// for the struct [ZoneSettingSortQueryStringForCacheEditResponse]
type zoneSettingSortQueryStringForCacheEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingSortQueryStringForCacheEditResponseID string

const (
	ZoneSettingSortQueryStringForCacheEditResponseIDSortQueryStringForCache ZoneSettingSortQueryStringForCacheEditResponseID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type ZoneSettingSortQueryStringForCacheEditResponseValue string

const (
	ZoneSettingSortQueryStringForCacheEditResponseValueOn  ZoneSettingSortQueryStringForCacheEditResponseValue = "on"
	ZoneSettingSortQueryStringForCacheEditResponseValueOff ZoneSettingSortQueryStringForCacheEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSortQueryStringForCacheEditResponseEditable bool

const (
	ZoneSettingSortQueryStringForCacheEditResponseEditableTrue  ZoneSettingSortQueryStringForCacheEditResponseEditable = true
	ZoneSettingSortQueryStringForCacheEditResponseEditableFalse ZoneSettingSortQueryStringForCacheEditResponseEditable = false
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingSortQueryStringForCacheGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingSortQueryStringForCacheGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingSortQueryStringForCacheGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSortQueryStringForCacheGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSortQueryStringForCacheGetResponseJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheGetResponseJSON contains the JSON metadata for
// the struct [ZoneSettingSortQueryStringForCacheGetResponse]
type zoneSettingSortQueryStringForCacheGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingSortQueryStringForCacheGetResponseID string

const (
	ZoneSettingSortQueryStringForCacheGetResponseIDSortQueryStringForCache ZoneSettingSortQueryStringForCacheGetResponseID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type ZoneSettingSortQueryStringForCacheGetResponseValue string

const (
	ZoneSettingSortQueryStringForCacheGetResponseValueOn  ZoneSettingSortQueryStringForCacheGetResponseValue = "on"
	ZoneSettingSortQueryStringForCacheGetResponseValueOff ZoneSettingSortQueryStringForCacheGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSortQueryStringForCacheGetResponseEditable bool

const (
	ZoneSettingSortQueryStringForCacheGetResponseEditableTrue  ZoneSettingSortQueryStringForCacheGetResponseEditable = true
	ZoneSettingSortQueryStringForCacheGetResponseEditableFalse ZoneSettingSortQueryStringForCacheGetResponseEditable = false
)

type ZoneSettingSortQueryStringForCacheEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingSortQueryStringForCacheEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingSortQueryStringForCacheEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingSortQueryStringForCacheEditParamsValue string

const (
	ZoneSettingSortQueryStringForCacheEditParamsValueOn  ZoneSettingSortQueryStringForCacheEditParamsValue = "on"
	ZoneSettingSortQueryStringForCacheEditParamsValueOff ZoneSettingSortQueryStringForCacheEditParamsValue = "off"
)

type ZoneSettingSortQueryStringForCacheEditResponseEnvelope struct {
	Errors   []ZoneSettingSortQueryStringForCacheEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSortQueryStringForCacheEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result ZoneSettingSortQueryStringForCacheEditResponse             `json:"result"`
	JSON   zoneSettingSortQueryStringForCacheEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCacheEditResponseEnvelope]
type zoneSettingSortQueryStringForCacheEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSortQueryStringForCacheEditResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCacheEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheEditResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZoneSettingSortQueryStringForCacheEditResponseEnvelopeErrors]
type zoneSettingSortQueryStringForCacheEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSortQueryStringForCacheEditResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCacheEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingSortQueryStringForCacheEditResponseEnvelopeMessages]
type zoneSettingSortQueryStringForCacheEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSortQueryStringForCacheGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingSortQueryStringForCacheGetResponseEnvelope struct {
	Errors   []ZoneSettingSortQueryStringForCacheGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingSortQueryStringForCacheGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result ZoneSettingSortQueryStringForCacheGetResponse             `json:"result"`
	JSON   zoneSettingSortQueryStringForCacheGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCacheGetResponseEnvelope]
type zoneSettingSortQueryStringForCacheGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSortQueryStringForCacheGetResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCacheGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZoneSettingSortQueryStringForCacheGetResponseEnvelopeErrors]
type zoneSettingSortQueryStringForCacheGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSortQueryStringForCacheGetResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCacheGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingSortQueryStringForCacheGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingSortQueryStringForCacheGetResponseEnvelopeMessages]
type zoneSettingSortQueryStringForCacheGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCacheGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSortQueryStringForCacheGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
