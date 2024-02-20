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

// SettingPseudoIPV4Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingPseudoIPV4Service] method
// instead.
type SettingPseudoIPV4Service struct {
	Options []option.RequestOption
}

// NewSettingPseudoIPV4Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingPseudoIPV4Service(opts ...option.RequestOption) (r *SettingPseudoIPV4Service) {
	r = &SettingPseudoIPV4Service{}
	r.Options = opts
	return
}

// Value of the Pseudo IPv4 setting.
func (r *SettingPseudoIPV4Service) Edit(ctx context.Context, zoneID string, body SettingPseudoIPV4EditParams, opts ...option.RequestOption) (res *SettingPseudoIPV4EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPseudoIPV4EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the Pseudo IPv4 setting.
func (r *SettingPseudoIPV4Service) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingPseudoIPV4GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPseudoIPV4GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The value set for the Pseudo IPv4 setting.
type SettingPseudoIPV4EditResponse struct {
	// Value of the Pseudo IPv4 setting.
	ID SettingPseudoIPV4EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingPseudoIPV4EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingPseudoIPV4EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingPseudoIPV4EditResponseJSON `json:"-"`
}

// settingPseudoIPV4EditResponseJSON contains the JSON metadata for the struct
// [SettingPseudoIPV4EditResponse]
type settingPseudoIPV4EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Value of the Pseudo IPv4 setting.
type SettingPseudoIPV4EditResponseID string

const (
	SettingPseudoIPV4EditResponseIDPseudoIPV4 SettingPseudoIPV4EditResponseID = "pseudo_ipv4"
)

// Current value of the zone setting.
type SettingPseudoIPV4EditResponseValue string

const (
	SettingPseudoIPV4EditResponseValueOff             SettingPseudoIPV4EditResponseValue = "off"
	SettingPseudoIPV4EditResponseValueAddHeader       SettingPseudoIPV4EditResponseValue = "add_header"
	SettingPseudoIPV4EditResponseValueOverwriteHeader SettingPseudoIPV4EditResponseValue = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingPseudoIPV4EditResponseEditable bool

const (
	SettingPseudoIPV4EditResponseEditableTrue  SettingPseudoIPV4EditResponseEditable = true
	SettingPseudoIPV4EditResponseEditableFalse SettingPseudoIPV4EditResponseEditable = false
)

// The value set for the Pseudo IPv4 setting.
type SettingPseudoIPV4GetResponse struct {
	// Value of the Pseudo IPv4 setting.
	ID SettingPseudoIPV4GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingPseudoIPV4GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingPseudoIPV4GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingPseudoIPV4GetResponseJSON `json:"-"`
}

// settingPseudoIPV4GetResponseJSON contains the JSON metadata for the struct
// [SettingPseudoIPV4GetResponse]
type settingPseudoIPV4GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Value of the Pseudo IPv4 setting.
type SettingPseudoIPV4GetResponseID string

const (
	SettingPseudoIPV4GetResponseIDPseudoIPV4 SettingPseudoIPV4GetResponseID = "pseudo_ipv4"
)

// Current value of the zone setting.
type SettingPseudoIPV4GetResponseValue string

const (
	SettingPseudoIPV4GetResponseValueOff             SettingPseudoIPV4GetResponseValue = "off"
	SettingPseudoIPV4GetResponseValueAddHeader       SettingPseudoIPV4GetResponseValue = "add_header"
	SettingPseudoIPV4GetResponseValueOverwriteHeader SettingPseudoIPV4GetResponseValue = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingPseudoIPV4GetResponseEditable bool

const (
	SettingPseudoIPV4GetResponseEditableTrue  SettingPseudoIPV4GetResponseEditable = true
	SettingPseudoIPV4GetResponseEditableFalse SettingPseudoIPV4GetResponseEditable = false
)

type SettingPseudoIPV4EditParams struct {
	// Value of the Pseudo IPv4 setting.
	Value param.Field[SettingPseudoIPV4EditParamsValue] `json:"value,required"`
}

func (r SettingPseudoIPV4EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Pseudo IPv4 setting.
type SettingPseudoIPV4EditParamsValue string

const (
	SettingPseudoIPV4EditParamsValueOff             SettingPseudoIPV4EditParamsValue = "off"
	SettingPseudoIPV4EditParamsValueAddHeader       SettingPseudoIPV4EditParamsValue = "add_header"
	SettingPseudoIPV4EditParamsValueOverwriteHeader SettingPseudoIPV4EditParamsValue = "overwrite_header"
)

type SettingPseudoIPV4EditResponseEnvelope struct {
	Errors   []SettingPseudoIPV4EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPseudoIPV4EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result SettingPseudoIPV4EditResponse             `json:"result"`
	JSON   settingPseudoIPV4EditResponseEnvelopeJSON `json:"-"`
}

// settingPseudoIPV4EditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingPseudoIPV4EditResponseEnvelope]
type settingPseudoIPV4EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPseudoIPV4EditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingPseudoIPV4EditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPseudoIPV4EditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingPseudoIPV4EditResponseEnvelopeErrors]
type settingPseudoIPV4EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPseudoIPV4EditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingPseudoIPV4EditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPseudoIPV4EditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingPseudoIPV4EditResponseEnvelopeMessages]
type settingPseudoIPV4EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPseudoIPV4GetResponseEnvelope struct {
	Errors   []SettingPseudoIPV4GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPseudoIPV4GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result SettingPseudoIPV4GetResponse             `json:"result"`
	JSON   settingPseudoIPV4GetResponseEnvelopeJSON `json:"-"`
}

// settingPseudoIPV4GetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingPseudoIPV4GetResponseEnvelope]
type settingPseudoIPV4GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPseudoIPV4GetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingPseudoIPV4GetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPseudoIPV4GetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingPseudoIPV4GetResponseEnvelopeErrors]
type settingPseudoIPV4GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPseudoIPV4GetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingPseudoIPV4GetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPseudoIPV4GetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingPseudoIPV4GetResponseEnvelopeMessages]
type settingPseudoIPV4GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPseudoIPV4GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
