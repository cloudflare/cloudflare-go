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

// SettingAutomaticHTTPSRewriteService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingAutomaticHTTPSRewriteService] method instead.
type SettingAutomaticHTTPSRewriteService struct {
	Options []option.RequestOption
}

// NewSettingAutomaticHTTPSRewriteService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingAutomaticHTTPSRewriteService(opts ...option.RequestOption) (r *SettingAutomaticHTTPSRewriteService) {
	r = &SettingAutomaticHTTPSRewriteService{}
	r.Options = opts
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *SettingAutomaticHTTPSRewriteService) Update(ctx context.Context, zoneID string, body SettingAutomaticHTTPSRewriteUpdateParams, opts ...option.RequestOption) (res *SettingAutomaticHTTPSRewriteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPSRewriteUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *SettingAutomaticHTTPSRewriteService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingAutomaticHTTPSRewriteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPSRewriteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingAutomaticHTTPSRewriteUpdateResponse struct {
	// ID of the zone setting.
	ID SettingAutomaticHTTPSRewriteUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAutomaticHTTPSRewriteUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAutomaticHTTPSRewriteUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAutomaticHTTPSRewriteUpdateResponseJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteUpdateResponseJSON contains the JSON metadata for
// the struct [SettingAutomaticHTTPSRewriteUpdateResponse]
type settingAutomaticHTTPSRewriteUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAutomaticHTTPSRewriteUpdateResponseID string

const (
	SettingAutomaticHTTPSRewriteUpdateResponseIDAutomaticHTTPSRewrites SettingAutomaticHTTPSRewriteUpdateResponseID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingAutomaticHTTPSRewriteUpdateResponseValue string

const (
	SettingAutomaticHTTPSRewriteUpdateResponseValueOn  SettingAutomaticHTTPSRewriteUpdateResponseValue = "on"
	SettingAutomaticHTTPSRewriteUpdateResponseValueOff SettingAutomaticHTTPSRewriteUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAutomaticHTTPSRewriteUpdateResponseEditable bool

const (
	SettingAutomaticHTTPSRewriteUpdateResponseEditableTrue  SettingAutomaticHTTPSRewriteUpdateResponseEditable = true
	SettingAutomaticHTTPSRewriteUpdateResponseEditableFalse SettingAutomaticHTTPSRewriteUpdateResponseEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingAutomaticHTTPSRewriteGetResponse struct {
	// ID of the zone setting.
	ID SettingAutomaticHTTPSRewriteGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAutomaticHTTPSRewriteGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAutomaticHTTPSRewriteGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAutomaticHTTPSRewriteGetResponseJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteGetResponseJSON contains the JSON metadata for the
// struct [SettingAutomaticHTTPSRewriteGetResponse]
type settingAutomaticHTTPSRewriteGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAutomaticHTTPSRewriteGetResponseID string

const (
	SettingAutomaticHTTPSRewriteGetResponseIDAutomaticHTTPSRewrites SettingAutomaticHTTPSRewriteGetResponseID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingAutomaticHTTPSRewriteGetResponseValue string

const (
	SettingAutomaticHTTPSRewriteGetResponseValueOn  SettingAutomaticHTTPSRewriteGetResponseValue = "on"
	SettingAutomaticHTTPSRewriteGetResponseValueOff SettingAutomaticHTTPSRewriteGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAutomaticHTTPSRewriteGetResponseEditable bool

const (
	SettingAutomaticHTTPSRewriteGetResponseEditableTrue  SettingAutomaticHTTPSRewriteGetResponseEditable = true
	SettingAutomaticHTTPSRewriteGetResponseEditableFalse SettingAutomaticHTTPSRewriteGetResponseEditable = false
)

type SettingAutomaticHTTPSRewriteUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingAutomaticHTTPSRewriteUpdateParamsValue] `json:"value,required"`
}

func (r SettingAutomaticHTTPSRewriteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingAutomaticHTTPSRewriteUpdateParamsValue string

const (
	SettingAutomaticHTTPSRewriteUpdateParamsValueOn  SettingAutomaticHTTPSRewriteUpdateParamsValue = "on"
	SettingAutomaticHTTPSRewriteUpdateParamsValueOff SettingAutomaticHTTPSRewriteUpdateParamsValue = "off"
)

type SettingAutomaticHTTPSRewriteUpdateResponseEnvelope struct {
	Errors   []SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result SettingAutomaticHTTPSRewriteUpdateResponse             `json:"result"`
	JSON   settingAutomaticHTTPSRewriteUpdateResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingAutomaticHTTPSRewriteUpdateResponseEnvelope]
type settingAutomaticHTTPSRewriteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    settingAutomaticHTTPSRewriteUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeErrors]
type settingAutomaticHTTPSRewriteUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    settingAutomaticHTTPSRewriteUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeMessages]
type settingAutomaticHTTPSRewriteUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPSRewriteGetResponseEnvelope struct {
	Errors   []SettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result SettingAutomaticHTTPSRewriteGetResponse             `json:"result"`
	JSON   settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAutomaticHTTPSRewriteGetResponseEnvelope]
type settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingAutomaticHTTPSRewriteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors]
type settingAutomaticHTTPSRewriteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingAutomaticHTTPSRewriteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages]
type settingAutomaticHTTPSRewriteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
