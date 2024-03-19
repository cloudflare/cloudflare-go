// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
func (r *SettingHTTP2Service) Edit(ctx context.Context, params SettingHTTP2EditParams, opts ...option.RequestOption) (res *ZonesHTTP2, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP2EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http2", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the HTTP2 setting.
func (r *SettingHTTP2Service) Get(ctx context.Context, query SettingHTTP2GetParams, opts ...option.RequestOption) (res *ZonesHTTP2, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP2GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http2", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP2 enabled for this zone.
type ZonesHTTP2 struct {
	// ID of the zone setting.
	ID ZonesHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesHTTP2JSON `json:"-"`
}

// zonesHTTP2JSON contains the JSON metadata for the struct [ZonesHTTP2]
type zonesHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesHTTP2JSON) RawJSON() string {
	return r.raw
}

func (r ZonesHTTP2) implementsZonesSettingEditResponse() {}

func (r ZonesHTTP2) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesHTTP2ID string

const (
	ZonesHTTP2IDHTTP2 ZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type ZonesHTTP2Value string

const (
	ZonesHTTP2ValueOn  ZonesHTTP2Value = "on"
	ZonesHTTP2ValueOff ZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesHTTP2Editable bool

const (
	ZonesHTTP2EditableTrue  ZonesHTTP2Editable = true
	ZonesHTTP2EditableFalse ZonesHTTP2Editable = false
)

// HTTP2 enabled for this zone.
type ZonesHTTP2Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesHTTP2ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesHTTP2Value] `json:"value,required"`
}

func (r ZonesHTTP2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesHTTP2Param) implementsZonesSettingEditParamsItem() {}

type SettingHTTP2EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	Result ZonesHTTP2                           `json:"result"`
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

func (r settingHTTP2EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r settingHTTP2EditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingHTTP2EditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingHTTP2GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingHTTP2GetResponseEnvelope struct {
	Errors   []SettingHTTP2GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP2GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result ZonesHTTP2                          `json:"result"`
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

func (r settingHTTP2GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r settingHTTP2GetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingHTTP2GetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
