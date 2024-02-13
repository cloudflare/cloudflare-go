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
func (r *SettingHTTP3Service) Update(ctx context.Context, zoneID string, body SettingHTTP3UpdateParams, opts ...option.RequestOption) (res *SettingHTTP3UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP3UpdateResponseEnvelope
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
type SettingHTTP3UpdateResponse struct {
	// ID of the zone setting.
	ID SettingHTTP3UpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingHTTP3UpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingHTTP3UpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingHTTP3UpdateResponseJSON `json:"-"`
}

// settingHTTP3UpdateResponseJSON contains the JSON metadata for the struct
// [SettingHTTP3UpdateResponse]
type settingHTTP3UpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingHTTP3UpdateResponseID string

const (
	SettingHTTP3UpdateResponseIDHTTP3 SettingHTTP3UpdateResponseID = "http3"
)

// Current value of the zone setting.
type SettingHTTP3UpdateResponseValue string

const (
	SettingHTTP3UpdateResponseValueOn  SettingHTTP3UpdateResponseValue = "on"
	SettingHTTP3UpdateResponseValueOff SettingHTTP3UpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingHTTP3UpdateResponseEditable bool

const (
	SettingHTTP3UpdateResponseEditableTrue  SettingHTTP3UpdateResponseEditable = true
	SettingHTTP3UpdateResponseEditableFalse SettingHTTP3UpdateResponseEditable = false
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

type SettingHTTP3UpdateParams struct {
	// Value of the HTTP3 setting.
	Value param.Field[SettingHTTP3UpdateParamsValue] `json:"value,required"`
}

func (r SettingHTTP3UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP3 setting.
type SettingHTTP3UpdateParamsValue string

const (
	SettingHTTP3UpdateParamsValueOn  SettingHTTP3UpdateParamsValue = "on"
	SettingHTTP3UpdateParamsValueOff SettingHTTP3UpdateParamsValue = "off"
)

type SettingHTTP3UpdateResponseEnvelope struct {
	Errors   []SettingHTTP3UpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP3UpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result SettingHTTP3UpdateResponse             `json:"result"`
	JSON   settingHTTP3UpdateResponseEnvelopeJSON `json:"-"`
}

// settingHTTP3UpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP3UpdateResponseEnvelope]
type settingHTTP3UpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3UpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP3UpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingHTTP3UpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP3UpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP3UpdateResponseEnvelopeErrors]
type settingHTTP3UpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3UpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingHTTP3UpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingHTTP3UpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP3UpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingHTTP3UpdateResponseEnvelopeMessages]
type settingHTTP3UpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3UpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
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
