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
func (r *SettingZeroRTTService) Edit(ctx context.Context, params SettingZeroRTTEditParams, opts ...option.RequestOption) (res *ZoneSetting0rtt, err error) {
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
func (r *SettingZeroRTTService) Get(ctx context.Context, query SettingZeroRTTGetParams, opts ...option.RequestOption) (res *ZoneSetting0rtt, err error) {
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
type ZoneSetting0rtt struct {
	// ID of the zone setting.
	ID ZoneSetting0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSetting0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSetting0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSetting0rttJSON `json:"-"`
}

// zoneSetting0rttJSON contains the JSON metadata for the struct [ZoneSetting0rtt]
type zoneSetting0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSetting0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSetting0rttJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSetting0rttID string

const (
	ZoneSetting0rttID0rtt ZoneSetting0rttID = "0rtt"
)

func (r ZoneSetting0rttID) IsKnown() bool {
	switch r {
	case ZoneSetting0rttID0rtt:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSetting0rttValue string

const (
	ZoneSetting0rttValueOn  ZoneSetting0rttValue = "on"
	ZoneSetting0rttValueOff ZoneSetting0rttValue = "off"
)

func (r ZoneSetting0rttValue) IsKnown() bool {
	switch r {
	case ZoneSetting0rttValueOn, ZoneSetting0rttValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSetting0rttEditable bool

const (
	ZoneSetting0rttEditableTrue  ZoneSetting0rttEditable = true
	ZoneSetting0rttEditableFalse ZoneSetting0rttEditable = false
)

func (r ZoneSetting0rttEditable) IsKnown() bool {
	switch r {
	case ZoneSetting0rttEditableTrue, ZoneSetting0rttEditableFalse:
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result ZoneSetting0rtt                        `json:"result"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result ZoneSetting0rtt                       `json:"result"`
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
