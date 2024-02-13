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
func (r *SettingSortQueryStringForCacheService) Update(ctx context.Context, zoneID string, body SettingSortQueryStringForCacheUpdateParams, opts ...option.RequestOption) (res *SettingSortQueryStringForCacheUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSortQueryStringForCacheUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *SettingSortQueryStringForCacheService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingSortQueryStringForCacheGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSortQueryStringForCacheGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", zoneID)
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
type SettingSortQueryStringForCacheUpdateResponse struct {
	// ID of the zone setting.
	ID SettingSortQueryStringForCacheUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSortQueryStringForCacheUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSortQueryStringForCacheUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSortQueryStringForCacheUpdateResponseJSON `json:"-"`
}

// settingSortQueryStringForCacheUpdateResponseJSON contains the JSON metadata for
// the struct [SettingSortQueryStringForCacheUpdateResponse]
type settingSortQueryStringForCacheUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingSortQueryStringForCacheUpdateResponseID string

const (
	SettingSortQueryStringForCacheUpdateResponseIDSortQueryStringForCache SettingSortQueryStringForCacheUpdateResponseID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type SettingSortQueryStringForCacheUpdateResponseValue string

const (
	SettingSortQueryStringForCacheUpdateResponseValueOn  SettingSortQueryStringForCacheUpdateResponseValue = "on"
	SettingSortQueryStringForCacheUpdateResponseValueOff SettingSortQueryStringForCacheUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSortQueryStringForCacheUpdateResponseEditable bool

const (
	SettingSortQueryStringForCacheUpdateResponseEditableTrue  SettingSortQueryStringForCacheUpdateResponseEditable = true
	SettingSortQueryStringForCacheUpdateResponseEditableFalse SettingSortQueryStringForCacheUpdateResponseEditable = false
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingSortQueryStringForCacheGetResponse struct {
	// ID of the zone setting.
	ID SettingSortQueryStringForCacheGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSortQueryStringForCacheGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSortQueryStringForCacheGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSortQueryStringForCacheGetResponseJSON `json:"-"`
}

// settingSortQueryStringForCacheGetResponseJSON contains the JSON metadata for the
// struct [SettingSortQueryStringForCacheGetResponse]
type settingSortQueryStringForCacheGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingSortQueryStringForCacheGetResponseID string

const (
	SettingSortQueryStringForCacheGetResponseIDSortQueryStringForCache SettingSortQueryStringForCacheGetResponseID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type SettingSortQueryStringForCacheGetResponseValue string

const (
	SettingSortQueryStringForCacheGetResponseValueOn  SettingSortQueryStringForCacheGetResponseValue = "on"
	SettingSortQueryStringForCacheGetResponseValueOff SettingSortQueryStringForCacheGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSortQueryStringForCacheGetResponseEditable bool

const (
	SettingSortQueryStringForCacheGetResponseEditableTrue  SettingSortQueryStringForCacheGetResponseEditable = true
	SettingSortQueryStringForCacheGetResponseEditableFalse SettingSortQueryStringForCacheGetResponseEditable = false
)

type SettingSortQueryStringForCacheUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingSortQueryStringForCacheUpdateParamsValue] `json:"value,required"`
}

func (r SettingSortQueryStringForCacheUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingSortQueryStringForCacheUpdateParamsValue string

const (
	SettingSortQueryStringForCacheUpdateParamsValueOn  SettingSortQueryStringForCacheUpdateParamsValue = "on"
	SettingSortQueryStringForCacheUpdateParamsValueOff SettingSortQueryStringForCacheUpdateParamsValue = "off"
)

type SettingSortQueryStringForCacheUpdateResponseEnvelope struct {
	Errors   []SettingSortQueryStringForCacheUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSortQueryStringForCacheUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result SettingSortQueryStringForCacheUpdateResponse             `json:"result"`
	JSON   settingSortQueryStringForCacheUpdateResponseEnvelopeJSON `json:"-"`
}

// settingSortQueryStringForCacheUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingSortQueryStringForCacheUpdateResponseEnvelope]
type settingSortQueryStringForCacheUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSortQueryStringForCacheUpdateResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    settingSortQueryStringForCacheUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSortQueryStringForCacheUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingSortQueryStringForCacheUpdateResponseEnvelopeErrors]
type settingSortQueryStringForCacheUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSortQueryStringForCacheUpdateResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    settingSortQueryStringForCacheUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSortQueryStringForCacheUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SettingSortQueryStringForCacheUpdateResponseEnvelopeMessages]
type settingSortQueryStringForCacheUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSortQueryStringForCacheGetResponseEnvelope struct {
	Errors   []SettingSortQueryStringForCacheGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSortQueryStringForCacheGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result SettingSortQueryStringForCacheGetResponse             `json:"result"`
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

type SettingSortQueryStringForCacheGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingSortQueryStringForCacheGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSortQueryStringForCacheGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingSortQueryStringForCacheGetResponseEnvelopeErrors]
type settingSortQueryStringForCacheGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSortQueryStringForCacheGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    settingSortQueryStringForCacheGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSortQueryStringForCacheGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingSortQueryStringForCacheGetResponseEnvelopeMessages]
type settingSortQueryStringForCacheGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSortQueryStringForCacheGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
