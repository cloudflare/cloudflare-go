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
func (r *SettingHTTP3Service) Edit(ctx context.Context, params SettingHTTP3EditParams, opts ...option.RequestOption) (res *ZoneSettingHTTP3, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP3EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the HTTP3 setting.
func (r *SettingHTTP3Service) Get(ctx context.Context, query SettingHTTP3GetParams, opts ...option.RequestOption) (res *ZoneSettingHTTP3, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP3GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP3 enabled for this zone.
type ZoneSettingHTTP3 struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP3JSON `json:"-"`
}

// zoneSettingHTTP3JSON contains the JSON metadata for the struct
// [ZoneSettingHTTP3]
type zoneSettingHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingHTTP3JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingHTTP3) implementsZonesSettingEditResponse() {}

func (r ZoneSettingHTTP3) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingHTTP3ID string

const (
	ZoneSettingHTTP3IDHTTP3 ZoneSettingHTTP3ID = "http3"
)

func (r ZoneSettingHTTP3ID) IsKnown() bool {
	switch r {
	case ZoneSettingHTTP3IDHTTP3:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingHTTP3Value string

const (
	ZoneSettingHTTP3ValueOn  ZoneSettingHTTP3Value = "on"
	ZoneSettingHTTP3ValueOff ZoneSettingHTTP3Value = "off"
)

func (r ZoneSettingHTTP3Value) IsKnown() bool {
	switch r {
	case ZoneSettingHTTP3ValueOn, ZoneSettingHTTP3ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP3Editable bool

const (
	ZoneSettingHTTP3EditableTrue  ZoneSettingHTTP3Editable = true
	ZoneSettingHTTP3EditableFalse ZoneSettingHTTP3Editable = false
)

func (r ZoneSettingHTTP3Editable) IsKnown() bool {
	switch r {
	case ZoneSettingHTTP3EditableTrue, ZoneSettingHTTP3EditableFalse:
		return true
	}
	return false
}

// HTTP3 enabled for this zone.
type ZoneSettingHTTP3Param struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingHTTP3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingHTTP3Value] `json:"value,required"`
}

func (r ZoneSettingHTTP3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingHTTP3Param) implementsZonesSettingEditParamsItemUnion() {}

type SettingHTTP3EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the HTTP3 setting.
	Value param.Field[SettingHTTP3EditParamsValue] `json:"value,required"`
}

func (r SettingHTTP3EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP3 setting.
type SettingHTTP3EditParamsValue string

const (
	SettingHTTP3EditParamsValueOn  SettingHTTP3EditParamsValue = "on"
	SettingHTTP3EditParamsValueOff SettingHTTP3EditParamsValue = "off"
)

func (r SettingHTTP3EditParamsValue) IsKnown() bool {
	switch r {
	case SettingHTTP3EditParamsValueOn, SettingHTTP3EditParamsValueOff:
		return true
	}
	return false
}

type SettingHTTP3EditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result ZoneSettingHTTP3                     `json:"result"`
	JSON   settingHTTP3EditResponseEnvelopeJSON `json:"-"`
}

// settingHTTP3EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP3EditResponseEnvelope]
type settingHTTP3EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingHTTP3EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingHTTP3GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingHTTP3GetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result ZoneSettingHTTP3                    `json:"result"`
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

func (r settingHTTP3GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
