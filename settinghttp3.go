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

// SettingHTTP3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingHTTP3Service] method
// instead.
type SettingHTTP3Service struct {
	Options []option.RequestOption
}

// NewSettingHTTP3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingHTTP3Service(opts ...option.RequestOption) (r *SettingHTTP3Service) {
	r = &SettingHTTP3Service{}
	r.Options = opts
	return
}

// Value of the HTTP3 setting.
func (r *SettingHTTP3Service) Edit(ctx context.Context, zoneID string, body SettingHTTP3EditParams, opts ...option.RequestOption) (res *SettingHTTP3EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP3EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the HTTP3 setting.
func (r *SettingHTTP3Service) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingHTTP3GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP3GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP3 enabled for this zone.
type SettingHTTP3EditResponse struct {
	// ID of the zone setting.
	ID SettingHTTP3EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHTTP3EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHTTP3EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHTTP3EditResponseJSON `json:"-"`
}

// settingHTTP3EditResponseJSON contains the JSON metadata for the struct
// [SettingHTTP3EditResponse]
type settingHTTP3EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHTTP3EditResponseID string

const (
	SettingHTTP3EditResponseIDHTTP3 SettingHTTP3EditResponseID = "http3"
)

// Current value of the zone setting.
type SettingHTTP3EditResponseValue string

const (
	SettingHTTP3EditResponseValueOn  SettingHTTP3EditResponseValue = "on"
	SettingHTTP3EditResponseValueOff SettingHTTP3EditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHTTP3EditResponseEditable bool

const (
	SettingHTTP3EditResponseEditableTrue  SettingHTTP3EditResponseEditable = true
	SettingHTTP3EditResponseEditableFalse SettingHTTP3EditResponseEditable = false
)

// HTTP3 enabled for this zone.
type SettingHTTP3GetResponse struct {
	// ID of the zone setting.
	ID SettingHTTP3GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHTTP3GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHTTP3GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHTTP3GetResponseJSON `json:"-"`
}

// settingHTTP3GetResponseJSON contains the JSON metadata for the struct
// [SettingHTTP3GetResponse]
type settingHTTP3GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHTTP3GetResponseID string

const (
	SettingHTTP3GetResponseIDHTTP3 SettingHTTP3GetResponseID = "http3"
)

// Current value of the zone setting.
type SettingHTTP3GetResponseValue string

const (
	SettingHTTP3GetResponseValueOn  SettingHTTP3GetResponseValue = "on"
	SettingHTTP3GetResponseValueOff SettingHTTP3GetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHTTP3GetResponseEditable bool

const (
	SettingHTTP3GetResponseEditableTrue  SettingHTTP3GetResponseEditable = true
	SettingHTTP3GetResponseEditableFalse SettingHTTP3GetResponseEditable = false
)

type SettingHTTP3EditParams struct {
	// Value of the HTTP3 setting.
	Value param.Field[SettingHTTP3EditParamsValue] `json:"value,required"`
}

func (r SettingHTTP3EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP3 setting.
type SettingHTTP3EditParamsValue string

const (
	SettingHTTP3EditParamsValueOn  SettingHTTP3EditParamsValue = "on"
	SettingHTTP3EditParamsValueOff SettingHTTP3EditParamsValue = "off"
)

type SettingHTTP3EditResponseEnvelope struct {
	Errors   []SettingHTTP3EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP3EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result SettingHTTP3EditResponse             `json:"result"`
	JSON   settingHTTP3EditResponseEnvelopeJSON `json:"-"`
}

// settingHTTP3EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP3EditResponseEnvelope]
type settingHTTP3EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP3EditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingHTTP3EditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP3EditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP3EditResponseEnvelopeErrors]
type settingHTTP3EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP3EditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingHTTP3EditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP3EditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingHTTP3EditResponseEnvelopeMessages]
type settingHTTP3EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP3GetResponseEnvelope struct {
	Errors   []SettingHTTP3GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP3GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result SettingHTTP3GetResponse             `json:"result"`
	JSON   settingHTTP3GetResponseEnvelopeJSON `json:"-"`
}

// settingHTTP3GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP3GetResponseEnvelope]
type settingHTTP3GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP3GetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingHTTP3GetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP3GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP3GetResponseEnvelopeErrors]
type settingHTTP3GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP3GetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingHTTP3GetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP3GetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingHTTP3GetResponseEnvelopeMessages]
type settingHTTP3GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
