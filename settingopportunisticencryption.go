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

// SettingOpportunisticEncryptionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingOpportunisticEncryptionService] method instead.
type SettingOpportunisticEncryptionService struct {
	Options []option.RequestOption
}

// NewSettingOpportunisticEncryptionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOpportunisticEncryptionService(opts ...option.RequestOption) (r *SettingOpportunisticEncryptionService) {
	r = &SettingOpportunisticEncryptionService{}
	r.Options = opts
	return
}

// Changes Opportunistic Encryption setting.
func (r *SettingOpportunisticEncryptionService) Update(ctx context.Context, zoneID string, body SettingOpportunisticEncryptionUpdateParams, opts ...option.RequestOption) (res *SettingOpportunisticEncryptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticEncryptionUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Opportunistic Encryption setting.
func (r *SettingOpportunisticEncryptionService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingOpportunisticEncryptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticEncryptionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables the Opportunistic Encryption feature for a zone.
type SettingOpportunisticEncryptionUpdateResponse struct {
	// ID of the zone setting.
	ID SettingOpportunisticEncryptionUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOpportunisticEncryptionUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOpportunisticEncryptionUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOpportunisticEncryptionUpdateResponseJSON `json:"-"`
}

// settingOpportunisticEncryptionUpdateResponseJSON contains the JSON metadata for
// the struct [SettingOpportunisticEncryptionUpdateResponse]
type settingOpportunisticEncryptionUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOpportunisticEncryptionUpdateResponseID string

const (
	SettingOpportunisticEncryptionUpdateResponseIDOpportunisticEncryption SettingOpportunisticEncryptionUpdateResponseID = "opportunistic_encryption"
)

// Current value of the zone setting.
type SettingOpportunisticEncryptionUpdateResponseValue string

const (
	SettingOpportunisticEncryptionUpdateResponseValueOn  SettingOpportunisticEncryptionUpdateResponseValue = "on"
	SettingOpportunisticEncryptionUpdateResponseValueOff SettingOpportunisticEncryptionUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOpportunisticEncryptionUpdateResponseEditable bool

const (
	SettingOpportunisticEncryptionUpdateResponseEditableTrue  SettingOpportunisticEncryptionUpdateResponseEditable = true
	SettingOpportunisticEncryptionUpdateResponseEditableFalse SettingOpportunisticEncryptionUpdateResponseEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type SettingOpportunisticEncryptionGetResponse struct {
	// ID of the zone setting.
	ID SettingOpportunisticEncryptionGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOpportunisticEncryptionGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOpportunisticEncryptionGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOpportunisticEncryptionGetResponseJSON `json:"-"`
}

// settingOpportunisticEncryptionGetResponseJSON contains the JSON metadata for the
// struct [SettingOpportunisticEncryptionGetResponse]
type settingOpportunisticEncryptionGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOpportunisticEncryptionGetResponseID string

const (
	SettingOpportunisticEncryptionGetResponseIDOpportunisticEncryption SettingOpportunisticEncryptionGetResponseID = "opportunistic_encryption"
)

// Current value of the zone setting.
type SettingOpportunisticEncryptionGetResponseValue string

const (
	SettingOpportunisticEncryptionGetResponseValueOn  SettingOpportunisticEncryptionGetResponseValue = "on"
	SettingOpportunisticEncryptionGetResponseValueOff SettingOpportunisticEncryptionGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOpportunisticEncryptionGetResponseEditable bool

const (
	SettingOpportunisticEncryptionGetResponseEditableTrue  SettingOpportunisticEncryptionGetResponseEditable = true
	SettingOpportunisticEncryptionGetResponseEditableFalse SettingOpportunisticEncryptionGetResponseEditable = false
)

type SettingOpportunisticEncryptionUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingOpportunisticEncryptionUpdateParamsValue] `json:"value,required"`
}

func (r SettingOpportunisticEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingOpportunisticEncryptionUpdateParamsValue string

const (
	SettingOpportunisticEncryptionUpdateParamsValueOn  SettingOpportunisticEncryptionUpdateParamsValue = "on"
	SettingOpportunisticEncryptionUpdateParamsValueOff SettingOpportunisticEncryptionUpdateParamsValue = "off"
)

type SettingOpportunisticEncryptionUpdateResponseEnvelope struct {
	Errors   []SettingOpportunisticEncryptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticEncryptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result SettingOpportunisticEncryptionUpdateResponse             `json:"result"`
	JSON   settingOpportunisticEncryptionUpdateResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticEncryptionUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingOpportunisticEncryptionUpdateResponseEnvelope]
type settingOpportunisticEncryptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticEncryptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    settingOpportunisticEncryptionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticEncryptionUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticEncryptionUpdateResponseEnvelopeErrors]
type settingOpportunisticEncryptionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticEncryptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    settingOpportunisticEncryptionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticEncryptionUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SettingOpportunisticEncryptionUpdateResponseEnvelopeMessages]
type settingOpportunisticEncryptionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticEncryptionGetResponseEnvelope struct {
	Errors   []SettingOpportunisticEncryptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticEncryptionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result SettingOpportunisticEncryptionGetResponse             `json:"result"`
	JSON   settingOpportunisticEncryptionGetResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticEncryptionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOpportunisticEncryptionGetResponseEnvelope]
type settingOpportunisticEncryptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticEncryptionGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticEncryptionGetResponseEnvelopeErrors]
type settingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticEncryptionGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    settingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticEncryptionGetResponseEnvelopeMessages]
type settingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticEncryptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
