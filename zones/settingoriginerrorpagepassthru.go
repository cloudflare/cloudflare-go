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

// SettingOriginErrorPagePassThruService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingOriginErrorPagePassThruService] method instead.
type SettingOriginErrorPagePassThruService struct {
	Options []option.RequestOption
}

// NewSettingOriginErrorPagePassThruService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOriginErrorPagePassThruService(opts ...option.RequestOption) (r *SettingOriginErrorPagePassThruService) {
	r = &SettingOriginErrorPagePassThruService{}
	r.Options = opts
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *SettingOriginErrorPagePassThruService) Edit(ctx context.Context, params SettingOriginErrorPagePassThruEditParams, opts ...option.RequestOption) (res *OriginErrorPagePassThru, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginErrorPagePassThruEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *SettingOriginErrorPagePassThruService) Get(ctx context.Context, query SettingOriginErrorPagePassThruGetParams, opts ...option.RequestOption) (res *OriginErrorPagePassThru, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginErrorPagePassThruGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type OriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID OriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value OriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       originErrorPagePassThruJSON `json:"-"`
}

// originErrorPagePassThruJSON contains the JSON metadata for the struct
// [OriginErrorPagePassThru]
type originErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type OriginErrorPagePassThruID string

const (
	OriginErrorPagePassThruIDOriginErrorPagePassThru OriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

func (r OriginErrorPagePassThruID) IsKnown() bool {
	switch r {
	case OriginErrorPagePassThruIDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Current value of the zone setting.
type OriginErrorPagePassThruValue string

const (
	OriginErrorPagePassThruValueOn  OriginErrorPagePassThruValue = "on"
	OriginErrorPagePassThruValueOff OriginErrorPagePassThruValue = "off"
)

func (r OriginErrorPagePassThruValue) IsKnown() bool {
	switch r {
	case OriginErrorPagePassThruValueOn, OriginErrorPagePassThruValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type OriginErrorPagePassThruEditable bool

const (
	OriginErrorPagePassThruEditableTrue  OriginErrorPagePassThruEditable = true
	OriginErrorPagePassThruEditableFalse OriginErrorPagePassThruEditable = false
)

func (r OriginErrorPagePassThruEditable) IsKnown() bool {
	switch r {
	case OriginErrorPagePassThruEditableTrue, OriginErrorPagePassThruEditableFalse:
		return true
	}
	return false
}

type SettingOriginErrorPagePassThruEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingOriginErrorPagePassThruEditParamsValue] `json:"value,required"`
}

func (r SettingOriginErrorPagePassThruEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingOriginErrorPagePassThruEditParamsValue string

const (
	SettingOriginErrorPagePassThruEditParamsValueOn  SettingOriginErrorPagePassThruEditParamsValue = "on"
	SettingOriginErrorPagePassThruEditParamsValueOff SettingOriginErrorPagePassThruEditParamsValue = "off"
)

func (r SettingOriginErrorPagePassThruEditParamsValue) IsKnown() bool {
	switch r {
	case SettingOriginErrorPagePassThruEditParamsValueOn, SettingOriginErrorPagePassThruEditParamsValueOff:
		return true
	}
	return false
}

type SettingOriginErrorPagePassThruEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result OriginErrorPagePassThru                                `json:"result"`
	JSON   settingOriginErrorPagePassThruEditResponseEnvelopeJSON `json:"-"`
}

// settingOriginErrorPagePassThruEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingOriginErrorPagePassThruEditResponseEnvelope]
type settingOriginErrorPagePassThruEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOriginErrorPagePassThruGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOriginErrorPagePassThruGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result OriginErrorPagePassThru                               `json:"result"`
	JSON   settingOriginErrorPagePassThruGetResponseEnvelopeJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOriginErrorPagePassThruGetResponseEnvelope]
type settingOriginErrorPagePassThruGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOriginErrorPagePassThruGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
