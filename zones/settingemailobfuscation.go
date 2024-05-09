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
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingEmailObfuscationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingEmailObfuscationService] method instead.
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
func (r *SettingEmailObfuscationService) Edit(ctx context.Context, params SettingEmailObfuscationEditParams, opts ...option.RequestOption) (res *EmailObfuscation, err error) {
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
func (r *SettingEmailObfuscationService) Get(ctx context.Context, query SettingEmailObfuscationGetParams, opts ...option.RequestOption) (res *EmailObfuscation, err error) {
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
type EmailObfuscation struct {
	// ID of the zone setting.
	ID EmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value EmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       emailObfuscationJSON `json:"-"`
}

// emailObfuscationJSON contains the JSON metadata for the struct
// [EmailObfuscation]
type emailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailObfuscationJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type EmailObfuscationID string

const (
	EmailObfuscationIDEmailObfuscation EmailObfuscationID = "email_obfuscation"
)

func (r EmailObfuscationID) IsKnown() bool {
	switch r {
	case EmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type EmailObfuscationValue string

const (
	EmailObfuscationValueOn  EmailObfuscationValue = "on"
	EmailObfuscationValueOff EmailObfuscationValue = "off"
)

func (r EmailObfuscationValue) IsKnown() bool {
	switch r {
	case EmailObfuscationValueOn, EmailObfuscationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type EmailObfuscationEditable bool

const (
	EmailObfuscationEditableTrue  EmailObfuscationEditable = true
	EmailObfuscationEditableFalse EmailObfuscationEditable = false
)

func (r EmailObfuscationEditable) IsKnown() bool {
	switch r {
	case EmailObfuscationEditableTrue, EmailObfuscationEditableFalse:
		return true
	}
	return false
}

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

func (r SettingEmailObfuscationEditParamsValue) IsKnown() bool {
	switch r {
	case SettingEmailObfuscationEditParamsValueOn, SettingEmailObfuscationEditParamsValueOff:
		return true
	}
	return false
}

type SettingEmailObfuscationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result EmailObfuscation                                `json:"result"`
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

type SettingEmailObfuscationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingEmailObfuscationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Encrypt email adresses on your web page from bots, while keeping them visible to
	// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
	Result EmailObfuscation                               `json:"result"`
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
