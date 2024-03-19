// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingEmailObfuscationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingEmailObfuscationService] method instead.
type SettingEmailObfuscationService struct {
	Options []option.RequestOption
}

// NewSettingEmailObfuscationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingEmailObfuscationService(opts ...option.RequestOption) (r *SettingEmailObfuscationService) {
	r = &SettingEmailObfuscationService{}
	r.Options = opts
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *SettingEmailObfuscationService) Edit(ctx context.Context, params SettingEmailObfuscationEditParams, opts ...option.RequestOption) (res *ZonesEmailObfuscation, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEmailObfuscationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *SettingEmailObfuscationService) Get(ctx context.Context, query SettingEmailObfuscationGetParams, opts ...option.RequestOption) (res *ZonesEmailObfuscation, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEmailObfuscationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID ZonesEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesEmailObfuscationJSON `json:"-"`
}

// zonesEmailObfuscationJSON contains the JSON metadata for the struct
// [ZonesEmailObfuscation]
type zonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r ZonesEmailObfuscation) implementsZonesSettingEditResponse() {}

func (r ZonesEmailObfuscation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesEmailObfuscationID string

const (
	ZonesEmailObfuscationIDEmailObfuscation ZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type ZonesEmailObfuscationValue string

const (
	ZonesEmailObfuscationValueOn  ZonesEmailObfuscationValue = "on"
	ZonesEmailObfuscationValueOff ZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesEmailObfuscationEditable bool

const (
	ZonesEmailObfuscationEditableTrue  ZonesEmailObfuscationEditable = true
	ZonesEmailObfuscationEditableFalse ZonesEmailObfuscationEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZonesEmailObfuscationParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesEmailObfuscationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesEmailObfuscationValue] `json:"value,required"`
}

func (r ZonesEmailObfuscationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesEmailObfuscationParam) implementsZonesSettingEditParamsItem() {}

type SettingEmailObfuscationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingEmailObfuscationEditParamsValue] `json:"value,required"`
}

func (r SettingEmailObfuscationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingEmailObfuscationEditParamsValue string

const (
	SettingEmailObfuscationEditParamsValueOn  SettingEmailObfuscationEditParamsValue = "on"
	SettingEmailObfuscationEditParamsValueOff SettingEmailObfuscationEditParamsValue = "off"
)

type SettingEmailObfuscationEditResponseEnvelope struct {
	Errors   []SettingEmailObfuscationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEmailObfuscationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result ZonesEmailObfuscation                           `json:"result"`
	JSON   settingEmailObfuscationEditResponseEnvelopeJSON `json:"-"`
}

// settingEmailObfuscationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingEmailObfuscationEditResponseEnvelope]
type settingEmailObfuscationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingEmailObfuscationEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingEmailObfuscationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEmailObfuscationEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingEmailObfuscationEditResponseEnvelopeErrors]
type settingEmailObfuscationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingEmailObfuscationEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingEmailObfuscationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEmailObfuscationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingEmailObfuscationEditResponseEnvelopeMessages]
type settingEmailObfuscationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingEmailObfuscationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingEmailObfuscationGetResponseEnvelope struct {
	Errors   []SettingEmailObfuscationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEmailObfuscationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result ZonesEmailObfuscation                          `json:"result"`
	JSON   settingEmailObfuscationGetResponseEnvelopeJSON `json:"-"`
}

// settingEmailObfuscationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingEmailObfuscationGetResponseEnvelope]
type settingEmailObfuscationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingEmailObfuscationGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingEmailObfuscationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEmailObfuscationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingEmailObfuscationGetResponseEnvelopeErrors]
type settingEmailObfuscationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingEmailObfuscationGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingEmailObfuscationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEmailObfuscationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingEmailObfuscationGetResponseEnvelopeMessages]
type settingEmailObfuscationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEmailObfuscationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEmailObfuscationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
