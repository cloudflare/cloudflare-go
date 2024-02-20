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

// SettingOrangeToOrangeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingOrangeToOrangeService]
// method instead.
type SettingOrangeToOrangeService struct {
	Options []option.RequestOption
}

// NewSettingOrangeToOrangeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingOrangeToOrangeService(opts ...option.RequestOption) (r *SettingOrangeToOrangeService) {
	r = &SettingOrangeToOrangeService{}
	r.Options = opts
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *SettingOrangeToOrangeService) Edit(ctx context.Context, zoneID string, body SettingOrangeToOrangeEditParams, opts ...option.RequestOption) (res *SettingOrangeToOrangeEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOrangeToOrangeEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *SettingOrangeToOrangeService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingOrangeToOrangeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOrangeToOrangeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingOrangeToOrangeEditResponse struct {
	// ID of the zone setting.
	ID SettingOrangeToOrangeEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOrangeToOrangeEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOrangeToOrangeEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOrangeToOrangeEditResponseJSON `json:"-"`
}

// settingOrangeToOrangeEditResponseJSON contains the JSON metadata for the struct
// [SettingOrangeToOrangeEditResponse]
type settingOrangeToOrangeEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOrangeToOrangeEditResponseID string

const (
	SettingOrangeToOrangeEditResponseIDOrangeToOrange SettingOrangeToOrangeEditResponseID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingOrangeToOrangeEditResponseValue string

const (
	SettingOrangeToOrangeEditResponseValueOn  SettingOrangeToOrangeEditResponseValue = "on"
	SettingOrangeToOrangeEditResponseValueOff SettingOrangeToOrangeEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOrangeToOrangeEditResponseEditable bool

const (
	SettingOrangeToOrangeEditResponseEditableTrue  SettingOrangeToOrangeEditResponseEditable = true
	SettingOrangeToOrangeEditResponseEditableFalse SettingOrangeToOrangeEditResponseEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingOrangeToOrangeGetResponse struct {
	// ID of the zone setting.
	ID SettingOrangeToOrangeGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOrangeToOrangeGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOrangeToOrangeGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOrangeToOrangeGetResponseJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseJSON contains the JSON metadata for the struct
// [SettingOrangeToOrangeGetResponse]
type settingOrangeToOrangeGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOrangeToOrangeGetResponseID string

const (
	SettingOrangeToOrangeGetResponseIDOrangeToOrange SettingOrangeToOrangeGetResponseID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingOrangeToOrangeGetResponseValue string

const (
	SettingOrangeToOrangeGetResponseValueOn  SettingOrangeToOrangeGetResponseValue = "on"
	SettingOrangeToOrangeGetResponseValueOff SettingOrangeToOrangeGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOrangeToOrangeGetResponseEditable bool

const (
	SettingOrangeToOrangeGetResponseEditableTrue  SettingOrangeToOrangeGetResponseEditable = true
	SettingOrangeToOrangeGetResponseEditableFalse SettingOrangeToOrangeGetResponseEditable = false
)

type SettingOrangeToOrangeEditParams struct {
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Value param.Field[SettingOrangeToOrangeEditParamsValue] `json:"value,required"`
}

func (r SettingOrangeToOrangeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingOrangeToOrangeEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingOrangeToOrangeEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingOrangeToOrangeEditParamsValueValue] `json:"value,required"`
}

func (r SettingOrangeToOrangeEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingOrangeToOrangeEditParamsValueID string

const (
	SettingOrangeToOrangeEditParamsValueIDOrangeToOrange SettingOrangeToOrangeEditParamsValueID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingOrangeToOrangeEditParamsValueValue string

const (
	SettingOrangeToOrangeEditParamsValueValueOn  SettingOrangeToOrangeEditParamsValueValue = "on"
	SettingOrangeToOrangeEditParamsValueValueOff SettingOrangeToOrangeEditParamsValueValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOrangeToOrangeEditParamsValueEditable bool

const (
	SettingOrangeToOrangeEditParamsValueEditableTrue  SettingOrangeToOrangeEditParamsValueEditable = true
	SettingOrangeToOrangeEditParamsValueEditableFalse SettingOrangeToOrangeEditParamsValueEditable = false
)

type SettingOrangeToOrangeEditResponseEnvelope struct {
	Errors   []SettingOrangeToOrangeEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOrangeToOrangeEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result SettingOrangeToOrangeEditResponse             `json:"result"`
	JSON   settingOrangeToOrangeEditResponseEnvelopeJSON `json:"-"`
}

// settingOrangeToOrangeEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOrangeToOrangeEditResponseEnvelope]
type settingOrangeToOrangeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingOrangeToOrangeEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOrangeToOrangeEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeEditResponseEnvelopeErrors]
type settingOrangeToOrangeEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingOrangeToOrangeEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOrangeToOrangeEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeEditResponseEnvelopeMessages]
type settingOrangeToOrangeEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeGetResponseEnvelope struct {
	Errors   []SettingOrangeToOrangeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOrangeToOrangeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result SettingOrangeToOrangeGetResponse             `json:"result"`
	JSON   settingOrangeToOrangeGetResponseEnvelopeJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOrangeToOrangeGetResponseEnvelope]
type settingOrangeToOrangeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingOrangeToOrangeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeGetResponseEnvelopeErrors]
type settingOrangeToOrangeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingOrangeToOrangeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeGetResponseEnvelopeMessages]
type settingOrangeToOrangeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
