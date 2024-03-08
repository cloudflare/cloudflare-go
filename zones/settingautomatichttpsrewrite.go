// File generated from our OpenAPI spec by Stainless.

package zones

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
func (r *SettingAutomaticHTTPSRewriteService) Edit(ctx context.Context, params SettingAutomaticHTTPSRewriteEditParams, opts ...option.RequestOption) (res *SettingAutomaticHTTPSRewriteEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPSRewriteEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *SettingAutomaticHTTPSRewriteService) Get(ctx context.Context, query SettingAutomaticHTTPSRewriteGetParams, opts ...option.RequestOption) (res *SettingAutomaticHTTPSRewriteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAutomaticHTTPSRewriteGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingAutomaticHTTPSRewriteEditResponse struct {
	// ID of the zone setting.
	ID SettingAutomaticHTTPSRewriteEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAutomaticHTTPSRewriteEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAutomaticHTTPSRewriteEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAutomaticHTTPSRewriteEditResponseJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteEditResponseJSON contains the JSON metadata for the
// struct [SettingAutomaticHTTPSRewriteEditResponse]
type settingAutomaticHTTPSRewriteEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAutomaticHTTPSRewriteEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingAutomaticHTTPSRewriteEditResponseID string

const (
	SettingAutomaticHTTPSRewriteEditResponseIDAutomaticHTTPSRewrites SettingAutomaticHTTPSRewriteEditResponseID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingAutomaticHTTPSRewriteEditResponseValue string

const (
	SettingAutomaticHTTPSRewriteEditResponseValueOn  SettingAutomaticHTTPSRewriteEditResponseValue = "on"
	SettingAutomaticHTTPSRewriteEditResponseValueOff SettingAutomaticHTTPSRewriteEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAutomaticHTTPSRewriteEditResponseEditable bool

const (
	SettingAutomaticHTTPSRewriteEditResponseEditableTrue  SettingAutomaticHTTPSRewriteEditResponseEditable = true
	SettingAutomaticHTTPSRewriteEditResponseEditableFalse SettingAutomaticHTTPSRewriteEditResponseEditable = false
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

func (r settingAutomaticHTTPSRewriteGetResponseJSON) RawJSON() string {
	return r.raw
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

type SettingAutomaticHTTPSRewriteEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingAutomaticHTTPSRewriteEditParamsValue] `json:"value,required"`
}

func (r SettingAutomaticHTTPSRewriteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingAutomaticHTTPSRewriteEditParamsValue string

const (
	SettingAutomaticHTTPSRewriteEditParamsValueOn  SettingAutomaticHTTPSRewriteEditParamsValue = "on"
	SettingAutomaticHTTPSRewriteEditParamsValueOff SettingAutomaticHTTPSRewriteEditParamsValue = "off"
)

type SettingAutomaticHTTPSRewriteEditResponseEnvelope struct {
	Errors   []SettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result SettingAutomaticHTTPSRewriteEditResponse             `json:"result"`
	JSON   settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingAutomaticHTTPSRewriteEditResponseEnvelope]
type settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAutomaticHTTPSRewriteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    settingAutomaticHTTPSRewriteEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors]
type settingAutomaticHTTPSRewriteEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAutomaticHTTPSRewriteEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    settingAutomaticHTTPSRewriteEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAutomaticHTTPSRewriteEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages]
type settingAutomaticHTTPSRewriteEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAutomaticHTTPSRewriteEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAutomaticHTTPSRewriteEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAutomaticHTTPSRewriteGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

func (r settingAutomaticHTTPSRewriteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r settingAutomaticHTTPSRewriteGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingAutomaticHTTPSRewriteGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
