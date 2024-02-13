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

// SettingTLS1_3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingTLS1_3Service] method
// instead.
type SettingTLS1_3Service struct {
	Options []option.RequestOption
}

// NewSettingTLS1_3Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingTLS1_3Service(opts ...option.RequestOption) (r *SettingTLS1_3Service) {
	r = &SettingTLS1_3Service{}
	r.Options = opts
	return
}

// Changes TLS 1.3 setting.
func (r *SettingTLS1_3Service) Update(ctx context.Context, zoneID string, body SettingTLS1_3UpdateParams, opts ...option.RequestOption) (res *SettingTls1_3UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTls1_3UpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets TLS 1.3 setting enabled for a zone.
func (r *SettingTLS1_3Service) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingTls1_3GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTls1_3GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables Crypto TLS 1.3 feature for a zone.
type SettingTls1_3UpdateResponse struct {
	// ID of the zone setting.
	ID SettingTls1_3UpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTls1_3UpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTls1_3UpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTls1_3UpdateResponseJSON `json:"-"`
}

// settingTls1_3UpdateResponseJSON contains the JSON metadata for the struct
// [SettingTls1_3UpdateResponse]
type settingTls1_3UpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingTls1_3UpdateResponseID string

const (
	SettingTls1_3UpdateResponseIDTLS1_3 SettingTls1_3UpdateResponseID = "tls_1_3"
)

// Current value of the zone setting.
type SettingTls1_3UpdateResponseValue string

const (
	SettingTls1_3UpdateResponseValueOn  SettingTls1_3UpdateResponseValue = "on"
	SettingTls1_3UpdateResponseValueOff SettingTls1_3UpdateResponseValue = "off"
	SettingTls1_3UpdateResponseValueZrt SettingTls1_3UpdateResponseValue = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTls1_3UpdateResponseEditable bool

const (
	SettingTls1_3UpdateResponseEditableTrue  SettingTls1_3UpdateResponseEditable = true
	SettingTls1_3UpdateResponseEditableFalse SettingTls1_3UpdateResponseEditable = false
)

// Enables Crypto TLS 1.3 feature for a zone.
type SettingTls1_3GetResponse struct {
	// ID of the zone setting.
	ID SettingTls1_3GetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTls1_3GetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTls1_3GetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTls1_3GetResponseJSON `json:"-"`
}

// settingTls1_3GetResponseJSON contains the JSON metadata for the struct
// [SettingTls1_3GetResponse]
type settingTls1_3GetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingTls1_3GetResponseID string

const (
	SettingTls1_3GetResponseIDTLS1_3 SettingTls1_3GetResponseID = "tls_1_3"
)

// Current value of the zone setting.
type SettingTls1_3GetResponseValue string

const (
	SettingTls1_3GetResponseValueOn  SettingTls1_3GetResponseValue = "on"
	SettingTls1_3GetResponseValueOff SettingTls1_3GetResponseValue = "off"
	SettingTls1_3GetResponseValueZrt SettingTls1_3GetResponseValue = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTls1_3GetResponseEditable bool

const (
	SettingTls1_3GetResponseEditableTrue  SettingTls1_3GetResponseEditable = true
	SettingTls1_3GetResponseEditableFalse SettingTls1_3GetResponseEditable = false
)

type SettingTLS1_3UpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingTls1_3UpdateParamsValue] `json:"value,required"`
}

func (r SettingTLS1_3UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingTls1_3UpdateParamsValue string

const (
	SettingTls1_3UpdateParamsValueOn  SettingTls1_3UpdateParamsValue = "on"
	SettingTls1_3UpdateParamsValueOff SettingTls1_3UpdateParamsValue = "off"
	SettingTls1_3UpdateParamsValueZrt SettingTls1_3UpdateParamsValue = "zrt"
)

type SettingTls1_3UpdateResponseEnvelope struct {
	Errors   []SettingTls1_3UpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTls1_3UpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result SettingTls1_3UpdateResponse             `json:"result"`
	JSON   settingTls1_3UpdateResponseEnvelopeJSON `json:"-"`
}

// settingTls1_3UpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTls1_3UpdateResponseEnvelope]
type settingTls1_3UpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3UpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTls1_3UpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingTls1_3UpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTls1_3UpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTls1_3UpdateResponseEnvelopeErrors]
type settingTls1_3UpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3UpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTls1_3UpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingTls1_3UpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTls1_3UpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingTls1_3UpdateResponseEnvelopeMessages]
type settingTls1_3UpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3UpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTls1_3GetResponseEnvelope struct {
	Errors   []SettingTls1_3GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTls1_3GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result SettingTls1_3GetResponse             `json:"result"`
	JSON   settingTls1_3GetResponseEnvelopeJSON `json:"-"`
}

// settingTls1_3GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingTls1_3GetResponseEnvelope]
type settingTls1_3GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTls1_3GetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingTls1_3GetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTls1_3GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTls1_3GetResponseEnvelopeErrors]
type settingTls1_3GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTls1_3GetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingTls1_3GetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTls1_3GetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingTls1_3GetResponseEnvelopeMessages]
type settingTls1_3GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
