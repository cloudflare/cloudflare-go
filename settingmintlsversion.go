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

// SettingMinTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMinTLSVersionService]
// method instead.
type SettingMinTLSVersionService struct {
	Options []option.RequestOption
}

// NewSettingMinTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingMinTLSVersionService(opts ...option.RequestOption) (r *SettingMinTLSVersionService) {
	r = &SettingMinTLSVersionService{}
	r.Options = opts
	return
}

// Changes Minimum TLS Version setting.
func (r *SettingMinTLSVersionService) Edit(ctx context.Context, params SettingMinTLSVersionEditParams, opts ...option.RequestOption) (res *SettingMinTLSVersionEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinTLSVersionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Minimum TLS Version setting.
func (r *SettingMinTLSVersionService) Get(ctx context.Context, query SettingMinTLSVersionGetParams, opts ...option.RequestOption) (res *SettingMinTLSVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinTLSVersionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type SettingMinTLSVersionEditResponse struct {
	// ID of the zone setting.
	ID SettingMinTLSVersionEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMinTLSVersionEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMinTLSVersionEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMinTLSVersionEditResponseJSON `json:"-"`
}

// settingMinTLSVersionEditResponseJSON contains the JSON metadata for the struct
// [SettingMinTLSVersionEditResponse]
type settingMinTLSVersionEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingMinTLSVersionEditResponseID string

const (
	SettingMinTLSVersionEditResponseIDMinTLSVersion SettingMinTLSVersionEditResponseID = "min_tls_version"
)

// Current value of the zone setting.
type SettingMinTLSVersionEditResponseValue string

const (
	SettingMinTLSVersionEditResponseValue1_0 SettingMinTLSVersionEditResponseValue = "1.0"
	SettingMinTLSVersionEditResponseValue1_1 SettingMinTLSVersionEditResponseValue = "1.1"
	SettingMinTLSVersionEditResponseValue1_2 SettingMinTLSVersionEditResponseValue = "1.2"
	SettingMinTLSVersionEditResponseValue1_3 SettingMinTLSVersionEditResponseValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMinTLSVersionEditResponseEditable bool

const (
	SettingMinTLSVersionEditResponseEditableTrue  SettingMinTLSVersionEditResponseEditable = true
	SettingMinTLSVersionEditResponseEditableFalse SettingMinTLSVersionEditResponseEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type SettingMinTLSVersionGetResponse struct {
	// ID of the zone setting.
	ID SettingMinTLSVersionGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMinTLSVersionGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMinTLSVersionGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMinTLSVersionGetResponseJSON `json:"-"`
}

// settingMinTLSVersionGetResponseJSON contains the JSON metadata for the struct
// [SettingMinTLSVersionGetResponse]
type settingMinTLSVersionGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingMinTLSVersionGetResponseID string

const (
	SettingMinTLSVersionGetResponseIDMinTLSVersion SettingMinTLSVersionGetResponseID = "min_tls_version"
)

// Current value of the zone setting.
type SettingMinTLSVersionGetResponseValue string

const (
	SettingMinTLSVersionGetResponseValue1_0 SettingMinTLSVersionGetResponseValue = "1.0"
	SettingMinTLSVersionGetResponseValue1_1 SettingMinTLSVersionGetResponseValue = "1.1"
	SettingMinTLSVersionGetResponseValue1_2 SettingMinTLSVersionGetResponseValue = "1.2"
	SettingMinTLSVersionGetResponseValue1_3 SettingMinTLSVersionGetResponseValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMinTLSVersionGetResponseEditable bool

const (
	SettingMinTLSVersionGetResponseEditableTrue  SettingMinTLSVersionGetResponseEditable = true
	SettingMinTLSVersionGetResponseEditableFalse SettingMinTLSVersionGetResponseEditable = false
)

type SettingMinTLSVersionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMinTLSVersionEditParamsValue] `json:"value,required"`
}

func (r SettingMinTLSVersionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMinTLSVersionEditParamsValue string

const (
	SettingMinTLSVersionEditParamsValue1_0 SettingMinTLSVersionEditParamsValue = "1.0"
	SettingMinTLSVersionEditParamsValue1_1 SettingMinTLSVersionEditParamsValue = "1.1"
	SettingMinTLSVersionEditParamsValue1_2 SettingMinTLSVersionEditParamsValue = "1.2"
	SettingMinTLSVersionEditParamsValue1_3 SettingMinTLSVersionEditParamsValue = "1.3"
)

type SettingMinTLSVersionEditResponseEnvelope struct {
	Errors   []SettingMinTLSVersionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinTLSVersionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result SettingMinTLSVersionEditResponse             `json:"result"`
	JSON   settingMinTLSVersionEditResponseEnvelopeJSON `json:"-"`
}

// settingMinTLSVersionEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMinTLSVersionEditResponseEnvelope]
type settingMinTLSVersionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinTLSVersionEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingMinTLSVersionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinTLSVersionEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMinTLSVersionEditResponseEnvelopeErrors]
type settingMinTLSVersionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinTLSVersionEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingMinTLSVersionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinTLSVersionEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMinTLSVersionEditResponseEnvelopeMessages]
type settingMinTLSVersionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinTLSVersionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMinTLSVersionGetResponseEnvelope struct {
	Errors   []SettingMinTLSVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinTLSVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result SettingMinTLSVersionGetResponse             `json:"result"`
	JSON   settingMinTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// settingMinTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMinTLSVersionGetResponseEnvelope]
type settingMinTLSVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinTLSVersionGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingMinTLSVersionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinTLSVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingMinTLSVersionGetResponseEnvelopeErrors]
type settingMinTLSVersionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinTLSVersionGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingMinTLSVersionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinTLSVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMinTLSVersionGetResponseEnvelopeMessages]
type settingMinTLSVersionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
