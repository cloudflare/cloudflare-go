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

// SettingZeroRTTService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingZeroRTTService] method
// instead.
type SettingZeroRTTService struct {
	Options []option.RequestOption
}

// NewSettingZeroRTTService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingZeroRTTService(opts ...option.RequestOption) (r *SettingZeroRTTService) {
	r = &SettingZeroRTTService{}
	r.Options = opts
	return
}

// Changes the 0-RTT session resumption setting.
func (r *SettingZeroRTTService) Edit(ctx context.Context, params SettingZeroRTTEditParams, opts ...option.RequestOption) (res *ZeroRTT, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRTTEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets 0-RTT session resumption setting.
func (r *SettingZeroRTTService) Get(ctx context.Context, query SettingZeroRTTGetParams, opts ...option.RequestOption) (res *ZeroRTT, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRTTGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// 0-RTT session resumption enabled for this zone.
type ZeroRTT struct {
	// ID of the zone setting.
	ID ZeroRTTID `json:"id,required"`
	// Current value of the zone setting.
	Value ZeroRTTValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZeroRTTEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on,nullable" format:"date-time"`
	JSON       zeroRTTJSON `json:"-"`
}

// zeroRTTJSON contains the JSON metadata for the struct [ZeroRTT]
type zeroRTTJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroRTT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroRTTJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZeroRTTID string

const (
	ZeroRTTID0rtt ZeroRTTID = "0rtt"
)

func (r ZeroRTTID) IsKnown() bool {
	switch r {
	case ZeroRTTID0rtt:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZeroRTTValue string

const (
	ZeroRTTValueOn  ZeroRTTValue = "on"
	ZeroRTTValueOff ZeroRTTValue = "off"
)

func (r ZeroRTTValue) IsKnown() bool {
	switch r {
	case ZeroRTTValueOn, ZeroRTTValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZeroRTTEditable bool

const (
	ZeroRTTEditableTrue  ZeroRTTEditable = true
	ZeroRTTEditableFalse ZeroRTTEditable = false
)

func (r ZeroRTTEditable) IsKnown() bool {
	switch r {
	case ZeroRTTEditableTrue, ZeroRTTEditableFalse:
		return true
	}
	return false
}

type SettingZeroRTTEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the 0-RTT setting.
	Value param.Field[SettingZeroRTTEditParamsValue] `json:"value,required"`
}

func (r SettingZeroRTTEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the 0-RTT setting.
type SettingZeroRTTEditParamsValue string

const (
	SettingZeroRTTEditParamsValueOn  SettingZeroRTTEditParamsValue = "on"
	SettingZeroRTTEditParamsValueOff SettingZeroRTTEditParamsValue = "off"
)

func (r SettingZeroRTTEditParamsValue) IsKnown() bool {
	switch r {
	case SettingZeroRTTEditParamsValueOn, SettingZeroRTTEditParamsValueOff:
		return true
	}
	return false
}

type SettingZeroRTTEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result ZeroRTT                                `json:"result"`
	JSON   settingZeroRTTEditResponseEnvelopeJSON `json:"-"`
}

// settingZeroRTTEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZeroRTTEditResponseEnvelope]
type settingZeroRTTEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZeroRTTGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingZeroRTTGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result ZeroRTT                               `json:"result"`
	JSON   settingZeroRTTGetResponseEnvelopeJSON `json:"-"`
}

// settingZeroRTTGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZeroRTTGetResponseEnvelope]
type settingZeroRTTGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
