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
func (r *SettingMinTLSVersionService) Update(ctx context.Context, zoneID string, body SettingMinTLSVersionUpdateParams, opts ...option.RequestOption) (res *SettingMinTLSVersionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinTLSVersionUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Minimum TLS Version setting.
func (r *SettingMinTLSVersionService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingMinTLSVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinTLSVersionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type SettingMinTLSVersionUpdateResponse struct {
	// ID of the zone setting.
	ID SettingMinTLSVersionUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMinTLSVersionUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMinTLSVersionUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMinTLSVersionUpdateResponseJSON `json:"-"`
}

// settingMinTLSVersionUpdateResponseJSON contains the JSON metadata for the struct
// [SettingMinTLSVersionUpdateResponse]
type settingMinTLSVersionUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingMinTLSVersionUpdateResponseID string

const (
	SettingMinTLSVersionUpdateResponseIDMinTLSVersion SettingMinTLSVersionUpdateResponseID = "min_tls_version"
)

// Current value of the zone setting.
type SettingMinTLSVersionUpdateResponseValue string

const (
	SettingMinTLSVersionUpdateResponseValue1_0 SettingMinTLSVersionUpdateResponseValue = "1.0"
	SettingMinTLSVersionUpdateResponseValue1_1 SettingMinTLSVersionUpdateResponseValue = "1.1"
	SettingMinTLSVersionUpdateResponseValue1_2 SettingMinTLSVersionUpdateResponseValue = "1.2"
	SettingMinTLSVersionUpdateResponseValue1_3 SettingMinTLSVersionUpdateResponseValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMinTLSVersionUpdateResponseEditable bool

const (
	SettingMinTLSVersionUpdateResponseEditableTrue  SettingMinTLSVersionUpdateResponseEditable = true
	SettingMinTLSVersionUpdateResponseEditableFalse SettingMinTLSVersionUpdateResponseEditable = false
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

type SettingMinTLSVersionUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingMinTLSVersionUpdateParamsValue] `json:"value,required"`
}

func (r SettingMinTLSVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMinTLSVersionUpdateParamsValue string

const (
	SettingMinTLSVersionUpdateParamsValue1_0 SettingMinTLSVersionUpdateParamsValue = "1.0"
	SettingMinTLSVersionUpdateParamsValue1_1 SettingMinTLSVersionUpdateParamsValue = "1.1"
	SettingMinTLSVersionUpdateParamsValue1_2 SettingMinTLSVersionUpdateParamsValue = "1.2"
	SettingMinTLSVersionUpdateParamsValue1_3 SettingMinTLSVersionUpdateParamsValue = "1.3"
)

type SettingMinTLSVersionUpdateResponseEnvelope struct {
	Errors   []SettingMinTLSVersionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinTLSVersionUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result SettingMinTLSVersionUpdateResponse             `json:"result"`
	JSON   settingMinTLSVersionUpdateResponseEnvelopeJSON `json:"-"`
}

// settingMinTLSVersionUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingMinTLSVersionUpdateResponseEnvelope]
type settingMinTLSVersionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinTLSVersionUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingMinTLSVersionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinTLSVersionUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMinTLSVersionUpdateResponseEnvelopeErrors]
type settingMinTLSVersionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinTLSVersionUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingMinTLSVersionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinTLSVersionUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingMinTLSVersionUpdateResponseEnvelopeMessages]
type settingMinTLSVersionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinTLSVersionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
