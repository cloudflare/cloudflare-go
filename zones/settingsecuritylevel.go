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

// SettingSecurityLevelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingSecurityLevelService]
// method instead.
type SettingSecurityLevelService struct {
	Options []option.RequestOption
}

// NewSettingSecurityLevelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingSecurityLevelService(opts ...option.RequestOption) (r *SettingSecurityLevelService) {
	r = &SettingSecurityLevelService{}
	r.Options = opts
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *SettingSecurityLevelService) Edit(ctx context.Context, params SettingSecurityLevelEditParams, opts ...option.RequestOption) (res *SettingSecurityLevelEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityLevelEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_level", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *SettingSecurityLevelService) Get(ctx context.Context, query SettingSecurityLevelGetParams, opts ...option.RequestOption) (res *SettingSecurityLevelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityLevelGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_level", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingSecurityLevelEditResponse struct {
	// ID of the zone setting.
	ID SettingSecurityLevelEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSecurityLevelEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSecurityLevelEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSecurityLevelEditResponseJSON `json:"-"`
}

// settingSecurityLevelEditResponseJSON contains the JSON metadata for the struct
// [SettingSecurityLevelEditResponse]
type settingSecurityLevelEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingSecurityLevelEditResponseID string

const (
	SettingSecurityLevelEditResponseIDSecurityLevel SettingSecurityLevelEditResponseID = "security_level"
)

// Current value of the zone setting.
type SettingSecurityLevelEditResponseValue string

const (
	SettingSecurityLevelEditResponseValueOff            SettingSecurityLevelEditResponseValue = "off"
	SettingSecurityLevelEditResponseValueEssentiallyOff SettingSecurityLevelEditResponseValue = "essentially_off"
	SettingSecurityLevelEditResponseValueLow            SettingSecurityLevelEditResponseValue = "low"
	SettingSecurityLevelEditResponseValueMedium         SettingSecurityLevelEditResponseValue = "medium"
	SettingSecurityLevelEditResponseValueHigh           SettingSecurityLevelEditResponseValue = "high"
	SettingSecurityLevelEditResponseValueUnderAttack    SettingSecurityLevelEditResponseValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSecurityLevelEditResponseEditable bool

const (
	SettingSecurityLevelEditResponseEditableTrue  SettingSecurityLevelEditResponseEditable = true
	SettingSecurityLevelEditResponseEditableFalse SettingSecurityLevelEditResponseEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingSecurityLevelGetResponse struct {
	// ID of the zone setting.
	ID SettingSecurityLevelGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSecurityLevelGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSecurityLevelGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSecurityLevelGetResponseJSON `json:"-"`
}

// settingSecurityLevelGetResponseJSON contains the JSON metadata for the struct
// [SettingSecurityLevelGetResponse]
type settingSecurityLevelGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingSecurityLevelGetResponseID string

const (
	SettingSecurityLevelGetResponseIDSecurityLevel SettingSecurityLevelGetResponseID = "security_level"
)

// Current value of the zone setting.
type SettingSecurityLevelGetResponseValue string

const (
	SettingSecurityLevelGetResponseValueOff            SettingSecurityLevelGetResponseValue = "off"
	SettingSecurityLevelGetResponseValueEssentiallyOff SettingSecurityLevelGetResponseValue = "essentially_off"
	SettingSecurityLevelGetResponseValueLow            SettingSecurityLevelGetResponseValue = "low"
	SettingSecurityLevelGetResponseValueMedium         SettingSecurityLevelGetResponseValue = "medium"
	SettingSecurityLevelGetResponseValueHigh           SettingSecurityLevelGetResponseValue = "high"
	SettingSecurityLevelGetResponseValueUnderAttack    SettingSecurityLevelGetResponseValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSecurityLevelGetResponseEditable bool

const (
	SettingSecurityLevelGetResponseEditableTrue  SettingSecurityLevelGetResponseEditable = true
	SettingSecurityLevelGetResponseEditableFalse SettingSecurityLevelGetResponseEditable = false
)

type SettingSecurityLevelEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingSecurityLevelEditParamsValue] `json:"value,required"`
}

func (r SettingSecurityLevelEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingSecurityLevelEditParamsValue string

const (
	SettingSecurityLevelEditParamsValueOff            SettingSecurityLevelEditParamsValue = "off"
	SettingSecurityLevelEditParamsValueEssentiallyOff SettingSecurityLevelEditParamsValue = "essentially_off"
	SettingSecurityLevelEditParamsValueLow            SettingSecurityLevelEditParamsValue = "low"
	SettingSecurityLevelEditParamsValueMedium         SettingSecurityLevelEditParamsValue = "medium"
	SettingSecurityLevelEditParamsValueHigh           SettingSecurityLevelEditParamsValue = "high"
	SettingSecurityLevelEditParamsValueUnderAttack    SettingSecurityLevelEditParamsValue = "under_attack"
)

type SettingSecurityLevelEditResponseEnvelope struct {
	Errors   []SettingSecurityLevelEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSecurityLevelEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result SettingSecurityLevelEditResponse             `json:"result"`
	JSON   settingSecurityLevelEditResponseEnvelopeJSON `json:"-"`
}

// settingSecurityLevelEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityLevelEditResponseEnvelope]
type settingSecurityLevelEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityLevelEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingSecurityLevelEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSecurityLevelEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSecurityLevelEditResponseEnvelopeErrors]
type settingSecurityLevelEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityLevelEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingSecurityLevelEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSecurityLevelEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSecurityLevelEditResponseEnvelopeMessages]
type settingSecurityLevelEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityLevelGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingSecurityLevelGetResponseEnvelope struct {
	Errors   []SettingSecurityLevelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSecurityLevelGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result SettingSecurityLevelGetResponse             `json:"result"`
	JSON   settingSecurityLevelGetResponseEnvelopeJSON `json:"-"`
}

// settingSecurityLevelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingSecurityLevelGetResponseEnvelope]
type settingSecurityLevelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityLevelGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingSecurityLevelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSecurityLevelGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingSecurityLevelGetResponseEnvelopeErrors]
type settingSecurityLevelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSecurityLevelGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingSecurityLevelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSecurityLevelGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingSecurityLevelGetResponseEnvelopeMessages]
type settingSecurityLevelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSecurityLevelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
