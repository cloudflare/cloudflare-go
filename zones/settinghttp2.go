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
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
func (r *SettingHTTP2Service) Edit(ctx context.Context, params SettingHTTP2EditParams, opts ...option.RequestOption) (res *HTTP2, err error) {
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
func (r *SettingHTTP2Service) Get(ctx context.Context, query SettingHTTP2GetParams, opts ...option.RequestOption) (res *HTTP2, err error) {
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
type HTTP2 struct {
	// ID of the zone setting.
	ID HTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value HTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       http2JSON `json:"-"`
}

// http2JSON contains the JSON metadata for the struct [HTTP2]
type http2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r http2JSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type HTTP2ID string

const (
	HTTP2IDHTTP2 HTTP2ID = "http2"
)

func (r HTTP2ID) IsKnown() bool {
	switch r {
	case HTTP2IDHTTP2:
		return true
	}
	return false
}

// Current value of the zone setting.
type HTTP2Value string

const (
	HTTP2ValueOn  HTTP2Value = "on"
	HTTP2ValueOff HTTP2Value = "off"
)

func (r HTTP2Value) IsKnown() bool {
	switch r {
	case HTTP2ValueOn, HTTP2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type HTTP2Editable bool

const (
	HTTP2EditableTrue  HTTP2Editable = true
	HTTP2EditableFalse HTTP2Editable = false
)

func (r HTTP2Editable) IsKnown() bool {
	switch r {
	case HTTP2EditableTrue, HTTP2EditableFalse:
		return true
	}
	return false
}

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
	Result HTTP2                                `json:"result"`
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
	Result HTTP2                               `json:"result"`
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
