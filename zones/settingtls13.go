// File generated from our OpenAPI spec by Stainless.

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
func (r *SettingTLS1_3Service) Edit(ctx context.Context, params SettingTLS1_3EditParams, opts ...option.RequestOption) (res *ZonesTLS1_3, err error) {
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
func (r *SettingTLS1_3Service) Get(ctx context.Context, query SettingTLS1_3GetParams, opts ...option.RequestOption) (res *ZonesTLS1_3, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTls1_3GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables Crypto TLS 1.3 feature for a zone.
type ZonesTLS1_3 struct {
	// ID of the zone setting.
	ID ZonesTLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesTLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time       `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesTls1_3JSON `json:"-"`
}

// zonesTls1_3JSON contains the JSON metadata for the struct [ZonesTLS1_3]
type zonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesTls1_3JSON) RawJSON() string {
	return r.raw
}

func (r ZonesTLS1_3) implementsZonesSettingEditResponse() {}

func (r ZonesTLS1_3) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesTLS1_3ID string

const (
	ZonesTLS1_3IDTLS1_3 ZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type ZonesTLS1_3Value string

const (
	ZonesTLS1_3ValueOn  ZonesTLS1_3Value = "on"
	ZonesTLS1_3ValueOff ZonesTLS1_3Value = "off"
	ZonesTLS1_3ValueZrt ZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesTLS1_3Editable bool

const (
	ZonesTLS1_3EditableTrue  ZonesTLS1_3Editable = true
	ZonesTLS1_3EditableFalse ZonesTLS1_3Editable = false
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZonesTLS1_3Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesTLS1_3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesTLS1_3Value] `json:"value,required"`
}

func (r ZonesTLS1_3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesTLS1_3Param) implementsZonesSettingEditParamsItem() {}

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
	Result ZonesTLS1_3                           `json:"result"`
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
	Result ZonesTLS1_3                          `json:"result"`
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
