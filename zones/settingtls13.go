// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingTLS1_3Service contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingTLS1_3Service] method instead.
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
func (r *SettingTLS1_3Service) Edit(ctx context.Context, params SettingTLS1_3EditParams, opts ...option.RequestOption) (res *TLS1_3, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTls1_3EditResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets TLS 1.3 setting enabled for a zone.
func (r *SettingTLS1_3Service) Get(ctx context.Context, query SettingTLS1_3GetParams, opts ...option.RequestOption) (res *TLS1_3, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTls1_3GetResponseEnvelope
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables Crypto TLS 1.3 feature for a zone.
type TLS1_3 struct {
	// ID of the zone setting.
	ID TLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value TLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       tls1_3JSON `json:"-"`
}

// tls1_3JSON contains the JSON metadata for the struct [TLS1_3]
type tls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tls1_3JSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type TLS1_3ID string

const (
	TLS1_3IDTLS1_3 TLS1_3ID = "tls_1_3"
)

func (r TLS1_3ID) IsKnown() bool {
	switch r {
	case TLS1_3IDTLS1_3:
		return true
	}
	return false
}

// Current value of the zone setting.
type TLS1_3Value string

const (
	TLS1_3ValueOn  TLS1_3Value = "on"
	TLS1_3ValueOff TLS1_3Value = "off"
	TLS1_3ValueZrt TLS1_3Value = "zrt"
)

func (r TLS1_3Value) IsKnown() bool {
	switch r {
	case TLS1_3ValueOn, TLS1_3ValueOff, TLS1_3ValueZrt:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type TLS1_3Editable bool

const (
	TLS1_3EditableTrue  TLS1_3Editable = true
	TLS1_3EditableFalse TLS1_3Editable = false
)

func (r TLS1_3Editable) IsKnown() bool {
	switch r {
	case TLS1_3EditableTrue, TLS1_3EditableFalse:
		return true
	}
	return false
}

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

func (r SettingTls1_3EditParamsValue) IsKnown() bool {
	switch r {
	case SettingTls1_3EditParamsValueOn, SettingTls1_3EditParamsValueOff, SettingTls1_3EditParamsValueZrt:
		return true
	}
	return false
}

type SettingTls1_3EditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result TLS1_3                                `json:"result"`
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

type SettingTLS1_3GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingTls1_3GetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result TLS1_3                               `json:"result"`
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
