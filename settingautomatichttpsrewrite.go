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

// SettingAutomaticHTTPsRewriteService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingAutomaticHTTPsRewriteService] method instead.
type SettingAutomaticHTTPsRewriteService struct {
	Options []option.RequestOption
}

// NewSettingAutomaticHTTPsRewriteService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingAutomaticHTTPsRewriteService(opts ...option.RequestOption) (r *SettingAutomaticHTTPsRewriteService) {
	r = &SettingAutomaticHTTPsRewriteService{}
	r.Options = opts
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *SettingAutomaticHTTPsRewriteService) Update(ctx context.Context, zoneID string, body SettingAutomaticHTTPsRewriteUpdateParams, opts ...option.RequestOption) (res *SettingAutomaticHTTPsRewriteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPsRewriteUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *SettingAutomaticHTTPsRewriteService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingAutomaticHTTPsRewriteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPsRewriteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingAutomaticHTTPsRewriteUpdateResponse struct {
	// ID of the zone setting.
	ID SettingAutomaticHTTPsRewriteUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAutomaticHTTPsRewriteUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAutomaticHTTPsRewriteUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAutomaticHTTPsRewriteUpdateResponseJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteUpdateResponseJSON contains the JSON metadata for
// the struct [SettingAutomaticHTTPsRewriteUpdateResponse]
type settingAutomaticHTTPsRewriteUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAutomaticHTTPsRewriteUpdateResponseID string

const (
	SettingAutomaticHTTPsRewriteUpdateResponseIDAutomaticHTTPsRewrites SettingAutomaticHTTPsRewriteUpdateResponseID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingAutomaticHTTPsRewriteUpdateResponseValue string

const (
	SettingAutomaticHTTPsRewriteUpdateResponseValueOn  SettingAutomaticHTTPsRewriteUpdateResponseValue = "on"
	SettingAutomaticHTTPsRewriteUpdateResponseValueOff SettingAutomaticHTTPsRewriteUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAutomaticHTTPsRewriteUpdateResponseEditable bool

const (
	SettingAutomaticHTTPsRewriteUpdateResponseEditableTrue  SettingAutomaticHTTPsRewriteUpdateResponseEditable = true
	SettingAutomaticHTTPsRewriteUpdateResponseEditableFalse SettingAutomaticHTTPsRewriteUpdateResponseEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingAutomaticHTTPsRewriteGetResponse struct {
	// ID of the zone setting.
	ID SettingAutomaticHTTPsRewriteGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAutomaticHTTPsRewriteGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAutomaticHTTPsRewriteGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAutomaticHTTPsRewriteGetResponseJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteGetResponseJSON contains the JSON metadata for the
// struct [SettingAutomaticHTTPsRewriteGetResponse]
type settingAutomaticHTTPsRewriteGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAutomaticHTTPsRewriteGetResponseID string

const (
	SettingAutomaticHTTPsRewriteGetResponseIDAutomaticHTTPsRewrites SettingAutomaticHTTPsRewriteGetResponseID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingAutomaticHTTPsRewriteGetResponseValue string

const (
	SettingAutomaticHTTPsRewriteGetResponseValueOn  SettingAutomaticHTTPsRewriteGetResponseValue = "on"
	SettingAutomaticHTTPsRewriteGetResponseValueOff SettingAutomaticHTTPsRewriteGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAutomaticHTTPsRewriteGetResponseEditable bool

const (
	SettingAutomaticHTTPsRewriteGetResponseEditableTrue  SettingAutomaticHTTPsRewriteGetResponseEditable = true
	SettingAutomaticHTTPsRewriteGetResponseEditableFalse SettingAutomaticHTTPsRewriteGetResponseEditable = false
)

type SettingAutomaticHTTPsRewriteUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingAutomaticHTTPsRewriteUpdateParamsValue] `json:"value,required"`
}

func (r SettingAutomaticHTTPsRewriteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingAutomaticHTTPsRewriteUpdateParamsValue string

const (
	SettingAutomaticHTTPsRewriteUpdateParamsValueOn  SettingAutomaticHTTPsRewriteUpdateParamsValue = "on"
	SettingAutomaticHTTPsRewriteUpdateParamsValueOff SettingAutomaticHTTPsRewriteUpdateParamsValue = "off"
)

type SettingAutomaticHTTPsRewriteUpdateResponseEnvelope struct {
	Errors   []SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result SettingAutomaticHTTPsRewriteUpdateResponse             `json:"result"`
	JSON   settingAutomaticHTTPsRewriteUpdateResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingAutomaticHTTPsRewriteUpdateResponseEnvelope]
type settingAutomaticHTTPsRewriteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    settingAutomaticHTTPsRewriteUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeErrors]
type settingAutomaticHTTPsRewriteUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    settingAutomaticHTTPsRewriteUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeMessages]
type settingAutomaticHTTPsRewriteUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPsRewriteGetResponseEnvelope struct {
	Errors   []SettingAutomaticHTTPsRewriteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAutomaticHTTPsRewriteGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result SettingAutomaticHTTPsRewriteGetResponse             `json:"result"`
	JSON   settingAutomaticHTTPsRewriteGetResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAutomaticHTTPsRewriteGetResponseEnvelope]
type settingAutomaticHTTPsRewriteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPsRewriteGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingAutomaticHTTPsRewriteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingAutomaticHTTPsRewriteGetResponseEnvelopeErrors]
type settingAutomaticHTTPsRewriteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPsRewriteGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingAutomaticHTTPsRewriteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAutomaticHTTPsRewriteGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingAutomaticHTTPsRewriteGetResponseEnvelopeMessages]
type settingAutomaticHTTPsRewriteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPsRewriteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
