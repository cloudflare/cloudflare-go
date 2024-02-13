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
func (r *SettingSecurityLevelService) Update(ctx context.Context, zoneID string, body SettingSecurityLevelUpdateParams, opts ...option.RequestOption) (res *SettingSecurityLevelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityLevelUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_level", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingSecurityLevelService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingSecurityLevelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingSecurityLevelGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/security_level", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type SettingSecurityLevelUpdateResponse struct {
	// ID of the zone setting.
	ID SettingSecurityLevelUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingSecurityLevelUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingSecurityLevelUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingSecurityLevelUpdateResponseJSON `json:"-"`
}

// settingSecurityLevelUpdateResponseJSON contains the JSON metadata for the struct
// [SettingSecurityLevelUpdateResponse]
type settingSecurityLevelUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingSecurityLevelUpdateResponseID string

const (
	SettingSecurityLevelUpdateResponseIDSecurityLevel SettingSecurityLevelUpdateResponseID = "security_level"
)

// Current value of the zone setting.
type SettingSecurityLevelUpdateResponseValue string

const (
	SettingSecurityLevelUpdateResponseValueOff            SettingSecurityLevelUpdateResponseValue = "off"
	SettingSecurityLevelUpdateResponseValueEssentiallyOff SettingSecurityLevelUpdateResponseValue = "essentially_off"
	SettingSecurityLevelUpdateResponseValueLow            SettingSecurityLevelUpdateResponseValue = "low"
	SettingSecurityLevelUpdateResponseValueMedium         SettingSecurityLevelUpdateResponseValue = "medium"
	SettingSecurityLevelUpdateResponseValueHigh           SettingSecurityLevelUpdateResponseValue = "high"
	SettingSecurityLevelUpdateResponseValueUnderAttack    SettingSecurityLevelUpdateResponseValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingSecurityLevelUpdateResponseEditable bool

const (
	SettingSecurityLevelUpdateResponseEditableTrue  SettingSecurityLevelUpdateResponseEditable = true
	SettingSecurityLevelUpdateResponseEditableFalse SettingSecurityLevelUpdateResponseEditable = false
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

type SettingSecurityLevelUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingSecurityLevelUpdateParamsValue] `json:"value,required"`
}

func (r SettingSecurityLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingSecurityLevelUpdateParamsValue string

const (
	SettingSecurityLevelUpdateParamsValueOff            SettingSecurityLevelUpdateParamsValue = "off"
	SettingSecurityLevelUpdateParamsValueEssentiallyOff SettingSecurityLevelUpdateParamsValue = "essentially_off"
	SettingSecurityLevelUpdateParamsValueLow            SettingSecurityLevelUpdateParamsValue = "low"
	SettingSecurityLevelUpdateParamsValueMedium         SettingSecurityLevelUpdateParamsValue = "medium"
	SettingSecurityLevelUpdateParamsValueHigh           SettingSecurityLevelUpdateParamsValue = "high"
	SettingSecurityLevelUpdateParamsValueUnderAttack    SettingSecurityLevelUpdateParamsValue = "under_attack"
)

type SettingSecurityLevelUpdateResponseEnvelope struct {
	Errors   []SettingSecurityLevelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingSecurityLevelUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Choose the appropriate security profile for your website, which will
	// automatically adjust each of the security settings. If you choose to customize
	// an individual security setting, the profile will become Custom.
	// (https://support.cloudflare.com/hc/en-us/articles/200170056).
	Result SettingSecurityLevelUpdateResponse             `json:"result"`
	JSON   settingSecurityLevelUpdateResponseEnvelopeJSON `json:"-"`
}

// settingSecurityLevelUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingSecurityLevelUpdateResponseEnvelope]
type settingSecurityLevelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSecurityLevelUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingSecurityLevelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingSecurityLevelUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingSecurityLevelUpdateResponseEnvelopeErrors]
type settingSecurityLevelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingSecurityLevelUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingSecurityLevelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingSecurityLevelUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingSecurityLevelUpdateResponseEnvelopeMessages]
type settingSecurityLevelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSecurityLevelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
