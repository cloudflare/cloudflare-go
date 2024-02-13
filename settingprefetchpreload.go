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

// SettingPrefetchPreloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingPrefetchPreloadService]
// method instead.
type SettingPrefetchPreloadService struct {
	Options []option.RequestOption
}

// NewSettingPrefetchPreloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingPrefetchPreloadService(opts ...option.RequestOption) (r *SettingPrefetchPreloadService) {
	r = &SettingPrefetchPreloadService{}
	r.Options = opts
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *SettingPrefetchPreloadService) Update(ctx context.Context, zoneID string, body SettingPrefetchPreloadUpdateParams, opts ...option.RequestOption) (res *SettingPrefetchPreloadUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPrefetchPreloadUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *SettingPrefetchPreloadService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingPrefetchPreloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPrefetchPreloadGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type SettingPrefetchPreloadUpdateResponse struct {
	// ID of the zone setting.
	ID SettingPrefetchPreloadUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingPrefetchPreloadUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingPrefetchPreloadUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingPrefetchPreloadUpdateResponseJSON `json:"-"`
}

// settingPrefetchPreloadUpdateResponseJSON contains the JSON metadata for the
// struct [SettingPrefetchPreloadUpdateResponse]
type settingPrefetchPreloadUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingPrefetchPreloadUpdateResponseID string

const (
	SettingPrefetchPreloadUpdateResponseIDPrefetchPreload SettingPrefetchPreloadUpdateResponseID = "prefetch_preload"
)

// Current value of the zone setting.
type SettingPrefetchPreloadUpdateResponseValue string

const (
	SettingPrefetchPreloadUpdateResponseValueOn  SettingPrefetchPreloadUpdateResponseValue = "on"
	SettingPrefetchPreloadUpdateResponseValueOff SettingPrefetchPreloadUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingPrefetchPreloadUpdateResponseEditable bool

const (
	SettingPrefetchPreloadUpdateResponseEditableTrue  SettingPrefetchPreloadUpdateResponseEditable = true
	SettingPrefetchPreloadUpdateResponseEditableFalse SettingPrefetchPreloadUpdateResponseEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type SettingPrefetchPreloadGetResponse struct {
	// ID of the zone setting.
	ID SettingPrefetchPreloadGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingPrefetchPreloadGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingPrefetchPreloadGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingPrefetchPreloadGetResponseJSON `json:"-"`
}

// settingPrefetchPreloadGetResponseJSON contains the JSON metadata for the struct
// [SettingPrefetchPreloadGetResponse]
type settingPrefetchPreloadGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingPrefetchPreloadGetResponseID string

const (
	SettingPrefetchPreloadGetResponseIDPrefetchPreload SettingPrefetchPreloadGetResponseID = "prefetch_preload"
)

// Current value of the zone setting.
type SettingPrefetchPreloadGetResponseValue string

const (
	SettingPrefetchPreloadGetResponseValueOn  SettingPrefetchPreloadGetResponseValue = "on"
	SettingPrefetchPreloadGetResponseValueOff SettingPrefetchPreloadGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingPrefetchPreloadGetResponseEditable bool

const (
	SettingPrefetchPreloadGetResponseEditableTrue  SettingPrefetchPreloadGetResponseEditable = true
	SettingPrefetchPreloadGetResponseEditableFalse SettingPrefetchPreloadGetResponseEditable = false
)

type SettingPrefetchPreloadUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingPrefetchPreloadUpdateParamsValue] `json:"value,required"`
}

func (r SettingPrefetchPreloadUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingPrefetchPreloadUpdateParamsValue string

const (
	SettingPrefetchPreloadUpdateParamsValueOn  SettingPrefetchPreloadUpdateParamsValue = "on"
	SettingPrefetchPreloadUpdateParamsValueOff SettingPrefetchPreloadUpdateParamsValue = "off"
)

type SettingPrefetchPreloadUpdateResponseEnvelope struct {
	Errors   []SettingPrefetchPreloadUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPrefetchPreloadUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result SettingPrefetchPreloadUpdateResponse             `json:"result"`
	JSON   settingPrefetchPreloadUpdateResponseEnvelopeJSON `json:"-"`
}

// settingPrefetchPreloadUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingPrefetchPreloadUpdateResponseEnvelope]
type settingPrefetchPreloadUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPrefetchPreloadUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingPrefetchPreloadUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPrefetchPreloadUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingPrefetchPreloadUpdateResponseEnvelopeErrors]
type settingPrefetchPreloadUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPrefetchPreloadUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingPrefetchPreloadUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPrefetchPreloadUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingPrefetchPreloadUpdateResponseEnvelopeMessages]
type settingPrefetchPreloadUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPrefetchPreloadGetResponseEnvelope struct {
	Errors   []SettingPrefetchPreloadGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPrefetchPreloadGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result SettingPrefetchPreloadGetResponse             `json:"result"`
	JSON   settingPrefetchPreloadGetResponseEnvelopeJSON `json:"-"`
}

// settingPrefetchPreloadGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingPrefetchPreloadGetResponseEnvelope]
type settingPrefetchPreloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPrefetchPreloadGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingPrefetchPreloadGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPrefetchPreloadGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingPrefetchPreloadGetResponseEnvelopeErrors]
type settingPrefetchPreloadGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPrefetchPreloadGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingPrefetchPreloadGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPrefetchPreloadGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingPrefetchPreloadGetResponseEnvelopeMessages]
type settingPrefetchPreloadGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPrefetchPreloadGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
