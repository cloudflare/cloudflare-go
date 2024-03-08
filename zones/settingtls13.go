// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *SettingTLS1_3Service) Edit(ctx context.Context, params SettingTLS1_3EditParams, opts ...option.RequestOption) (res *SettingTls1_3EditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTls1_3EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets TLS 1.3 setting enabled for a zone.
func (r *SettingTLS1_3Service) Get(ctx context.Context, query SettingTLS1_3GetParams, opts ...option.RequestOption) (res *SettingTls1_3GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTls1_3GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables Crypto TLS 1.3 feature for a zone.
type SettingTls1_3EditResponse struct {
	// ID of the zone setting.
	ID SettingTls1_3EditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTls1_3EditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTls1_3EditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTls1_3EditResponseJSON `json:"-"`
}

// settingTls1_3EditResponseJSON contains the JSON metadata for the struct
// [SettingTls1_3EditResponse]
type settingTls1_3EditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTls1_3EditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingTls1_3EditResponseID string

const (
	SettingTls1_3EditResponseIDTLS1_3 SettingTls1_3EditResponseID = "tls_1_3"
)

// Current value of the zone setting.
type SettingTls1_3EditResponseValue string

const (
	SettingTls1_3EditResponseValueOn  SettingTls1_3EditResponseValue = "on"
	SettingTls1_3EditResponseValueOff SettingTls1_3EditResponseValue = "off"
	SettingTls1_3EditResponseValueZrt SettingTls1_3EditResponseValue = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTls1_3EditResponseEditable bool

const (
	SettingTls1_3EditResponseEditableTrue  SettingTls1_3EditResponseEditable = true
	SettingTls1_3EditResponseEditableFalse SettingTls1_3EditResponseEditable = false
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

func (r settingTls1_3GetResponseJSON) RawJSON() string {
	return r.raw
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

type SettingTLS1_3EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingTls1_3EditParamsValue] `json:"value,required"`
}

func (r SettingTLS1_3EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingTls1_3EditParamsValue string

const (
	SettingTls1_3EditParamsValueOn  SettingTls1_3EditParamsValue = "on"
	SettingTls1_3EditParamsValueOff SettingTls1_3EditParamsValue = "off"
	SettingTls1_3EditParamsValueZrt SettingTls1_3EditParamsValue = "zrt"
)

type SettingTls1_3EditResponseEnvelope struct {
	Errors   []SettingTls1_3EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTls1_3EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result SettingTls1_3EditResponse             `json:"result"`
	JSON   settingTls1_3EditResponseEnvelopeJSON `json:"-"`
}

// settingTls1_3EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingTls1_3EditResponseEnvelope]
type settingTls1_3EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTls1_3EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTls1_3EditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingTls1_3EditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTls1_3EditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTls1_3EditResponseEnvelopeErrors]
type settingTls1_3EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTls1_3EditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTls1_3EditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingTls1_3EditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTls1_3EditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingTls1_3EditResponseEnvelopeMessages]
type settingTls1_3EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTls1_3EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTls1_3EditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTLS1_3GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

func (r settingTls1_3GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r settingTls1_3GetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingTls1_3GetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
