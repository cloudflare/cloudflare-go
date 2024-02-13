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
func (r *SettingHTTP2Service) Update(ctx context.Context, zoneID string, body SettingHTTP2UpdateParams, opts ...option.RequestOption) (res *SettingHTTP2UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP2UpdateResponseEnvelope
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
type SettingHTTP2UpdateResponse struct {
	// ID of the zone setting.
	ID SettingHTTP2UpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHTTP2UpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHTTP2UpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHTTP2UpdateResponseJSON `json:"-"`
}

// settingHTTP2UpdateResponseJSON contains the JSON metadata for the struct
// [SettingHTTP2UpdateResponse]
type settingHTTP2UpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHTTP2UpdateResponseID string

const (
	SettingHTTP2UpdateResponseIDHTTP2 SettingHTTP2UpdateResponseID = "http2"
)

// Current value of the zone setting.
type SettingHTTP2UpdateResponseValue string

const (
	SettingHTTP2UpdateResponseValueOn  SettingHTTP2UpdateResponseValue = "on"
	SettingHTTP2UpdateResponseValueOff SettingHTTP2UpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHTTP2UpdateResponseEditable bool

const (
	SettingHTTP2UpdateResponseEditableTrue  SettingHTTP2UpdateResponseEditable = true
	SettingHTTP2UpdateResponseEditableFalse SettingHTTP2UpdateResponseEditable = false
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

type SettingHTTP2UpdateParams struct {
	// Value of the HTTP2 setting.
	Value param.Field[SettingHTTP2UpdateParamsValue] `json:"value,required"`
}

func (r SettingHTTP2UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP2 setting.
type SettingHTTP2UpdateParamsValue string

const (
	SettingHTTP2UpdateParamsValueOn  SettingHTTP2UpdateParamsValue = "on"
	SettingHTTP2UpdateParamsValueOff SettingHTTP2UpdateParamsValue = "off"
)

type SettingHTTP2UpdateResponseEnvelope struct {
	Errors   []SettingHTTP2UpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP2UpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result SettingHTTP2UpdateResponse             `json:"result"`
	JSON   settingHTTP2UpdateResponseEnvelopeJSON `json:"-"`
}

// settingHTTP2UpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP2UpdateResponseEnvelope]
type settingHTTP2UpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2UpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP2UpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingHTTP2UpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP2UpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP2UpdateResponseEnvelopeErrors]
type settingHTTP2UpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2UpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP2UpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingHTTP2UpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP2UpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingHTTP2UpdateResponseEnvelopeMessages]
type settingHTTP2UpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP2UpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
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
