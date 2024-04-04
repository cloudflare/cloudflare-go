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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *SettingHTTP2Service) Edit(ctx context.Context, params SettingHTTP2EditParams, opts ...option.RequestOption) (res *ZoneSettingHTTP2, err error) {
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
func (r *SettingHTTP2Service) Get(ctx context.Context, query SettingHTTP2GetParams, opts ...option.RequestOption) (res *ZoneSettingHTTP2, err error) {
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
type ZoneSettingHTTP2 struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP2JSON `json:"-"`
}

// zoneSettingHTTP2JSON contains the JSON metadata for the struct
// [ZoneSettingHTTP2]
type zoneSettingHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP2JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingHTTP2) implementsZonesSettingEditResponse() {}

func (r ZoneSettingHTTP2) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingHTTP2ID string

const (
	ZoneSettingHTTP2IDHTTP2 ZoneSettingHTTP2ID = "http2"
)

func (r ZoneSettingHTTP2ID) IsKnown() bool {
	switch r {
	case ZoneSettingHTTP2IDHTTP2:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingHTTP2Value string

const (
	ZoneSettingHTTP2ValueOn  ZoneSettingHTTP2Value = "on"
	ZoneSettingHTTP2ValueOff ZoneSettingHTTP2Value = "off"
)

func (r ZoneSettingHTTP2Value) IsKnown() bool {
	switch r {
	case ZoneSettingHTTP2ValueOn, ZoneSettingHTTP2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP2Editable bool

const (
	ZoneSettingHTTP2EditableTrue  ZoneSettingHTTP2Editable = true
	ZoneSettingHTTP2EditableFalse ZoneSettingHTTP2Editable = false
)

func (r ZoneSettingHTTP2Editable) IsKnown() bool {
	switch r {
	case ZoneSettingHTTP2EditableTrue, ZoneSettingHTTP2EditableFalse:
		return true
	}
	return false
}

// HTTP2 enabled for this zone.
type ZoneSettingHTTP2Param struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingHTTP2ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingHTTP2Value] `json:"value,required"`
}

func (r ZoneSettingHTTP2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingHTTP2Param) implementsZonesSettingEditParamsItemUnion() {}

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

func (r SettingHTTP2EditParamsValue) IsKnown() bool {
	switch r {
	case SettingHTTP2EditParamsValueOn, SettingHTTP2EditParamsValueOff:
		return true
	}
	return false
}

type SettingHTTP2EditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHTTP2                     `json:"result"`
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

type SettingHTTP2GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingHTTP2GetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHTTP2                    `json:"result"`
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
