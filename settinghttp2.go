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

// SettingHTTP2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingHTTP2Service] method
// instead.
type SettingHTTP2Service struct {
	Options []option.RequestOption
}

// NewSettingHTTP2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingHTTP2Service(opts ...option.RequestOption) (r *SettingHTTP2Service) {
	r = &SettingHTTP2Service{}
	r.Options = opts
	return
}

// Value of the HTTP2 setting.
func (r *SettingHTTP2Service) Edit(ctx context.Context, zoneID string, body SettingHTTP2EditParams, opts ...option.RequestOption) (res *SettingHTTP2EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP2EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http2", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the HTTP2 setting.
func (r *SettingHTTP2Service) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingHTTP2GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP2GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http2", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP2 enabled for this zone.
type SettingHTTP2EditResponse struct {
	// ID of the zone setting.
	ID SettingHTTP2EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHTTP2EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHTTP2EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHTTP2EditResponseJSON `json:"-"`
}

// settingHTTP2EditResponseJSON contains the JSON metadata for the struct
// [SettingHTTP2EditResponse]
type settingHTTP2EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHTTP2EditResponseID string

const (
	SettingHTTP2EditResponseIDHTTP2 SettingHTTP2EditResponseID = "http2"
)

// Current value of the zone setting.
type SettingHTTP2EditResponseValue string

const (
	SettingHTTP2EditResponseValueOn  SettingHTTP2EditResponseValue = "on"
	SettingHTTP2EditResponseValueOff SettingHTTP2EditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHTTP2EditResponseEditable bool

const (
	SettingHTTP2EditResponseEditableTrue  SettingHTTP2EditResponseEditable = true
	SettingHTTP2EditResponseEditableFalse SettingHTTP2EditResponseEditable = false
)

// HTTP2 enabled for this zone.
type SettingHTTP2GetResponse struct {
	// ID of the zone setting.
	ID SettingHTTP2GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHTTP2GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHTTP2GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHTTP2GetResponseJSON `json:"-"`
}

// settingHTTP2GetResponseJSON contains the JSON metadata for the struct
// [SettingHTTP2GetResponse]
type settingHTTP2GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHTTP2GetResponseID string

const (
	SettingHTTP2GetResponseIDHTTP2 SettingHTTP2GetResponseID = "http2"
)

// Current value of the zone setting.
type SettingHTTP2GetResponseValue string

const (
	SettingHTTP2GetResponseValueOn  SettingHTTP2GetResponseValue = "on"
	SettingHTTP2GetResponseValueOff SettingHTTP2GetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHTTP2GetResponseEditable bool

const (
	SettingHTTP2GetResponseEditableTrue  SettingHTTP2GetResponseEditable = true
	SettingHTTP2GetResponseEditableFalse SettingHTTP2GetResponseEditable = false
)

type SettingHTTP2EditParams struct {
	// Value of the HTTP2 setting.
	Value param.Field[SettingHTTP2EditParamsValue] `json:"value,required"`
}

func (r SettingHTTP2EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP2 setting.
type SettingHTTP2EditParamsValue string

const (
	SettingHTTP2EditParamsValueOn  SettingHTTP2EditParamsValue = "on"
	SettingHTTP2EditParamsValueOff SettingHTTP2EditParamsValue = "off"
)

type SettingHTTP2EditResponseEnvelope struct {
	Errors   []SettingHTTP2EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP2EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result SettingHTTP2EditResponse             `json:"result"`
	JSON   settingHTTP2EditResponseEnvelopeJSON `json:"-"`
}

// settingHTTP2EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP2EditResponseEnvelope]
type settingHTTP2EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP2EditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingHTTP2EditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP2EditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP2EditResponseEnvelopeErrors]
type settingHTTP2EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP2EditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingHTTP2EditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP2EditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingHTTP2EditResponseEnvelopeMessages]
type settingHTTP2EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP2GetResponseEnvelope struct {
	Errors   []SettingHTTP2GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP2GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result SettingHTTP2GetResponse             `json:"result"`
	JSON   settingHTTP2GetResponseEnvelopeJSON `json:"-"`
}

// settingHTTP2GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP2GetResponseEnvelope]
type settingHTTP2GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP2GetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingHTTP2GetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP2GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP2GetResponseEnvelopeErrors]
type settingHTTP2GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP2GetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingHTTP2GetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP2GetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingHTTP2GetResponseEnvelopeMessages]
type settingHTTP2GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
