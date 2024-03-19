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
func (r *SettingOrangeToOrangeService) Edit(ctx context.Context, params SettingOrangeToOrangeEditParams, opts ...option.RequestOption) (res *ZonesOrangeToOrange, err error) {
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
func (r *SettingOrangeToOrangeService) Get(ctx context.Context, query SettingOrangeToOrangeGetParams, opts ...option.RequestOption) (res *ZonesOrangeToOrange, err error) {
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
type ZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time               `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesOrangeToOrangeJSON `json:"-"`
}

// zonesOrangeToOrangeJSON contains the JSON metadata for the struct
// [ZonesOrangeToOrange]
type zonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesOrangeToOrangeJSON) RawJSON() string {
	return r.raw
}

func (r ZonesOrangeToOrange) implementsZonesSettingEditResponse() {}

func (r ZonesOrangeToOrange) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesOrangeToOrangeID string

const (
	ZonesOrangeToOrangeIDOrangeToOrange ZonesOrangeToOrangeID = "orange_to_orange"
)

func (r ZonesOrangeToOrangeID) IsKnown() bool {
	switch r {
	case ZonesOrangeToOrangeIDOrangeToOrange:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesOrangeToOrangeValue string

const (
	ZonesOrangeToOrangeValueOn  ZonesOrangeToOrangeValue = "on"
	ZonesOrangeToOrangeValueOff ZonesOrangeToOrangeValue = "off"
)

func (r ZonesOrangeToOrangeValue) IsKnown() bool {
	switch r {
	case ZonesOrangeToOrangeValueOn, ZonesOrangeToOrangeValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesOrangeToOrangeEditable bool

const (
	ZonesOrangeToOrangeEditableTrue  ZonesOrangeToOrangeEditable = true
	ZonesOrangeToOrangeEditableFalse ZonesOrangeToOrangeEditable = false
)

func (r ZonesOrangeToOrangeEditable) IsKnown() bool {
	switch r {
	case ZonesOrangeToOrangeEditableTrue, ZonesOrangeToOrangeEditableFalse:
		return true
	}
	return false
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZonesOrangeToOrangeParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOrangeToOrangeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesOrangeToOrangeValue] `json:"value,required"`
}

func (r ZonesOrangeToOrangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOrangeToOrangeParam) implementsZonesSettingEditParamsItem() {}

type SettingOrangeToOrangeEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Value param.Field[ZonesOrangeToOrangeParam] `json:"value,required"`
}

func (r SettingOrangeToOrangeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingOrangeToOrangeEditResponseEnvelope struct {
	Errors   []SettingOrangeToOrangeEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOrangeToOrangeEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZonesOrangeToOrange                           `json:"result"`
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

type SettingOrangeToOrangeEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingOrangeToOrangeEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOrangeToOrangeEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeEditResponseEnvelopeErrors]
type settingOrangeToOrangeEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOrangeToOrangeEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOrangeToOrangeEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingOrangeToOrangeEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOrangeToOrangeEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeEditResponseEnvelopeMessages]
type settingOrangeToOrangeEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOrangeToOrangeEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingOrangeToOrangeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOrangeToOrangeGetResponseEnvelope struct {
	Errors   []SettingOrangeToOrangeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOrangeToOrangeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result ZonesOrangeToOrange                          `json:"result"`
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

type SettingOrangeToOrangeGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingOrangeToOrangeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeGetResponseEnvelopeErrors]
type settingOrangeToOrangeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOrangeToOrangeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingOrangeToOrangeGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingOrangeToOrangeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeGetResponseEnvelopeMessages]
type settingOrangeToOrangeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOrangeToOrangeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
