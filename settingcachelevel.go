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
func (r *SettingCacheLevelService) Update(ctx context.Context, zoneID string, body SettingCacheLevelUpdateParams, opts ...option.RequestOption) (res *SettingCacheLevelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCacheLevelUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/cache_level", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingCacheLevelService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingCacheLevelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingCacheLevelGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/cache_level", zoneID)
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
type SettingCacheLevelUpdateResponse struct {
	// ID of the zone setting.
	ID SettingCacheLevelUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingCacheLevelUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingCacheLevelUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingCacheLevelUpdateResponseJSON `json:"-"`
}

// settingCacheLevelUpdateResponseJSON contains the JSON metadata for the struct
// [SettingCacheLevelUpdateResponse]
type settingCacheLevelUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingCacheLevelUpdateResponseID string

const (
	SettingCacheLevelUpdateResponseIDCacheLevel SettingCacheLevelUpdateResponseID = "cache_level"
)

// Current value of the zone setting.
type SettingCacheLevelUpdateResponseValue string

const (
	SettingCacheLevelUpdateResponseValueAggressive SettingCacheLevelUpdateResponseValue = "aggressive"
	SettingCacheLevelUpdateResponseValueBasic      SettingCacheLevelUpdateResponseValue = "basic"
	SettingCacheLevelUpdateResponseValueSimplified SettingCacheLevelUpdateResponseValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingCacheLevelUpdateResponseEditable bool

const (
	SettingCacheLevelUpdateResponseEditableTrue  SettingCacheLevelUpdateResponseEditable = true
	SettingCacheLevelUpdateResponseEditableFalse SettingCacheLevelUpdateResponseEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingCacheLevelGetResponse struct {
	// ID of the zone setting.
	ID SettingCacheLevelGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingCacheLevelGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingCacheLevelGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingCacheLevelGetResponseJSON `json:"-"`
}

// settingCacheLevelGetResponseJSON contains the JSON metadata for the struct
// [SettingCacheLevelGetResponse]
type settingCacheLevelGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingCacheLevelGetResponseID string

const (
	SettingCacheLevelGetResponseIDCacheLevel SettingCacheLevelGetResponseID = "cache_level"
)

// Current value of the zone setting.
type SettingCacheLevelGetResponseValue string

const (
	SettingCacheLevelGetResponseValueAggressive SettingCacheLevelGetResponseValue = "aggressive"
	SettingCacheLevelGetResponseValueBasic      SettingCacheLevelGetResponseValue = "basic"
	SettingCacheLevelGetResponseValueSimplified SettingCacheLevelGetResponseValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingCacheLevelGetResponseEditable bool

const (
	SettingCacheLevelGetResponseEditableTrue  SettingCacheLevelGetResponseEditable = true
	SettingCacheLevelGetResponseEditableFalse SettingCacheLevelGetResponseEditable = false
)

type SettingCacheLevelUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingCacheLevelUpdateParamsValue] `json:"value,required"`
}

func (r SettingCacheLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingCacheLevelUpdateParamsValue string

const (
	SettingCacheLevelUpdateParamsValueAggressive SettingCacheLevelUpdateParamsValue = "aggressive"
	SettingCacheLevelUpdateParamsValueBasic      SettingCacheLevelUpdateParamsValue = "basic"
	SettingCacheLevelUpdateParamsValueSimplified SettingCacheLevelUpdateParamsValue = "simplified"
)

type SettingCacheLevelUpdateResponseEnvelope struct {
	Errors   []SettingCacheLevelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingCacheLevelUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result SettingCacheLevelUpdateResponse             `json:"result"`
	JSON   settingCacheLevelUpdateResponseEnvelopeJSON `json:"-"`
}

// settingCacheLevelUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingCacheLevelUpdateResponseEnvelope]
type settingCacheLevelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingCacheLevelUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingCacheLevelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingCacheLevelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingCacheLevelUpdateResponseEnvelopeErrors]
type settingCacheLevelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingCacheLevelUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingCacheLevelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingCacheLevelUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingCacheLevelUpdateResponseEnvelopeMessages]
type settingCacheLevelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingCacheLevelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Result SettingCacheLevelGetResponse             `json:"result"`
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
