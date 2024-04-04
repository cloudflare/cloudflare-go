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

// SettingOrangeToOrangeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingOrangeToOrangeService]
// method instead.
type SettingOrangeToOrangeService struct {
	Options []option.RequestOption
}

// NewSettingOrangeToOrangeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingOrangeToOrangeService(opts ...option.RequestOption) (r *SettingOrangeToOrangeService) {
	r = &SettingOrangeToOrangeService{}
	r.Options = opts
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *SettingOrangeToOrangeService) Edit(ctx context.Context, params SettingOrangeToOrangeEditParams, opts ...option.RequestOption) (res *ZoneSettingOrangeToOrange, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOrangeToOrangeEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *SettingOrangeToOrangeService) Get(ctx context.Context, query SettingOrangeToOrangeGetParams, opts ...option.RequestOption) (res *ZoneSettingOrangeToOrange, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOrangeToOrangeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingOrangeToOrange struct {
	// ID of the zone setting.
	ID ZoneSettingOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOrangeToOrangeJSON `json:"-"`
}

// zoneSettingOrangeToOrangeJSON contains the JSON metadata for the struct
// [ZoneSettingOrangeToOrange]
type zoneSettingOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOrangeToOrangeJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingOrangeToOrange) implementsZonesSettingEditResponse() {}

func (r ZoneSettingOrangeToOrange) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingOrangeToOrangeID string

const (
	ZoneSettingOrangeToOrangeIDOrangeToOrange ZoneSettingOrangeToOrangeID = "orange_to_orange"
)

func (r ZoneSettingOrangeToOrangeID) IsKnown() bool {
	switch r {
	case ZoneSettingOrangeToOrangeIDOrangeToOrange:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingOrangeToOrangeValue string

const (
	ZoneSettingOrangeToOrangeValueOn  ZoneSettingOrangeToOrangeValue = "on"
	ZoneSettingOrangeToOrangeValueOff ZoneSettingOrangeToOrangeValue = "off"
)

func (r ZoneSettingOrangeToOrangeValue) IsKnown() bool {
	switch r {
	case ZoneSettingOrangeToOrangeValueOn, ZoneSettingOrangeToOrangeValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOrangeToOrangeEditable bool

const (
	ZoneSettingOrangeToOrangeEditableTrue  ZoneSettingOrangeToOrangeEditable = true
	ZoneSettingOrangeToOrangeEditableFalse ZoneSettingOrangeToOrangeEditable = false
)

func (r ZoneSettingOrangeToOrangeEditable) IsKnown() bool {
	switch r {
	case ZoneSettingOrangeToOrangeEditableTrue, ZoneSettingOrangeToOrangeEditableFalse:
		return true
	}
	return false
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingOrangeToOrangeParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingOrangeToOrangeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingOrangeToOrangeValue] `json:"value,required"`
}

func (r ZoneSettingOrangeToOrangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingOrangeToOrangeParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingOrangeToOrangeEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Value param.Field[ZoneSettingOrangeToOrangeParam] `json:"value,required"`
}

func (r SettingOrangeToOrangeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingOrangeToOrangeEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZoneSettingOrangeToOrange                     `json:"result"`
	JSON   settingOrangeToOrangeEditResponseEnvelopeJSON `json:"-"`
}

// settingOrangeToOrangeEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOrangeToOrangeEditResponseEnvelope]
type settingOrangeToOrangeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOrangeToOrangeEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingOrangeToOrangeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOrangeToOrangeGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZoneSettingOrangeToOrange                    `json:"result"`
	JSON   settingOrangeToOrangeGetResponseEnvelopeJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOrangeToOrangeGetResponseEnvelope]
type settingOrangeToOrangeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOrangeToOrangeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
