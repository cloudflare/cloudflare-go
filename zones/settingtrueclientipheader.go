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

// SettingTrueClientIPHeaderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingTrueClientIPHeaderService] method instead.
type SettingTrueClientIPHeaderService struct {
	Options []option.RequestOption
}

// NewSettingTrueClientIPHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingTrueClientIPHeaderService(opts ...option.RequestOption) (r *SettingTrueClientIPHeaderService) {
	r = &SettingTrueClientIPHeaderService{}
	r.Options = opts
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *SettingTrueClientIPHeaderService) Edit(ctx context.Context, params SettingTrueClientIPHeaderEditParams, opts ...option.RequestOption) (res *TrueClientIPHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTrueClientIPHeaderEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *SettingTrueClientIPHeaderService) Get(ctx context.Context, query SettingTrueClientIPHeaderGetParams, opts ...option.RequestOption) (res *TrueClientIPHeader, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTrueClientIPHeaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type TrueClientIPHeader struct {
	// ID of the zone setting.
	ID TrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value TrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       trueClientIPHeaderJSON `json:"-"`
}

// trueClientIPHeaderJSON contains the JSON metadata for the struct
// [TrueClientIPHeader]
type trueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type TrueClientIPHeaderID string

const (
	TrueClientIPHeaderIDTrueClientIPHeader TrueClientIPHeaderID = "true_client_ip_header"
)

func (r TrueClientIPHeaderID) IsKnown() bool {
	switch r {
	case TrueClientIPHeaderIDTrueClientIPHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type TrueClientIPHeaderValue string

const (
	TrueClientIPHeaderValueOn  TrueClientIPHeaderValue = "on"
	TrueClientIPHeaderValueOff TrueClientIPHeaderValue = "off"
)

func (r TrueClientIPHeaderValue) IsKnown() bool {
	switch r {
	case TrueClientIPHeaderValueOn, TrueClientIPHeaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type TrueClientIPHeaderEditable bool

const (
	TrueClientIPHeaderEditableTrue  TrueClientIPHeaderEditable = true
	TrueClientIPHeaderEditableFalse TrueClientIPHeaderEditable = false
)

func (r TrueClientIPHeaderEditable) IsKnown() bool {
	switch r {
	case TrueClientIPHeaderEditableTrue, TrueClientIPHeaderEditableFalse:
		return true
	}
	return false
}

type SettingTrueClientIPHeaderEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingTrueClientIPHeaderEditParamsValue] `json:"value,required"`
}

func (r SettingTrueClientIPHeaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingTrueClientIPHeaderEditParamsValue string

const (
	SettingTrueClientIPHeaderEditParamsValueOn  SettingTrueClientIPHeaderEditParamsValue = "on"
	SettingTrueClientIPHeaderEditParamsValueOff SettingTrueClientIPHeaderEditParamsValue = "off"
)

func (r SettingTrueClientIPHeaderEditParamsValue) IsKnown() bool {
	switch r {
	case SettingTrueClientIPHeaderEditParamsValueOn, SettingTrueClientIPHeaderEditParamsValueOff:
		return true
	}
	return false
}

type SettingTrueClientIPHeaderEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result TrueClientIPHeader                                `json:"result"`
	JSON   settingTrueClientIPHeaderEditResponseEnvelopeJSON `json:"-"`
}

// settingTrueClientIPHeaderEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTrueClientIPHeaderEditResponseEnvelope]
type settingTrueClientIPHeaderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrueClientIPHeaderGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingTrueClientIPHeaderGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result TrueClientIPHeader                               `json:"result"`
	JSON   settingTrueClientIPHeaderGetResponseEnvelopeJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTrueClientIPHeaderGetResponseEnvelope]
type settingTrueClientIPHeaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrueClientIPHeaderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
