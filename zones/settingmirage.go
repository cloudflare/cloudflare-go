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

// SettingMirageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMirageService] method
// instead.
type SettingMirageService struct {
	Options []option.RequestOption
}

// NewSettingMirageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingMirageService(opts ...option.RequestOption) (r *SettingMirageService) {
	r = &SettingMirageService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *SettingMirageService) Edit(ctx context.Context, params SettingMirageEditParams, opts ...option.RequestOption) (res *ZonesMirage, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMirageEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mirage", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *SettingMirageService) Get(ctx context.Context, query SettingMirageGetParams, opts ...option.RequestOption) (res *ZonesMirage, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMirageGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mirage", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZonesMirage struct {
	// ID of the zone setting.
	ID ZonesMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time       `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesMirageJSON `json:"-"`
}

// zonesMirageJSON contains the JSON metadata for the struct [ZonesMirage]
type zonesMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMirageJSON) RawJSON() string {
	return r.raw
}

func (r ZonesMirage) implementsZonesSettingEditResponse() {}

func (r ZonesMirage) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesMirageID string

const (
	ZonesMirageIDMirage ZonesMirageID = "mirage"
)

// Current value of the zone setting.
type ZonesMirageValue string

const (
	ZonesMirageValueOn  ZonesMirageValue = "on"
	ZonesMirageValueOff ZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesMirageEditable bool

const (
	ZonesMirageEditableTrue  ZonesMirageEditable = true
	ZonesMirageEditableFalse ZonesMirageEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZonesMirageParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesMirageID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesMirageValue] `json:"value,required"`
}

func (r ZonesMirageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesMirageParam) implementsZonesSettingEditParamsItem() {}

type SettingMirageEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMirageEditParamsValue] `json:"value,required"`
}

func (r SettingMirageEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMirageEditParamsValue string

const (
	SettingMirageEditParamsValueOn  SettingMirageEditParamsValue = "on"
	SettingMirageEditParamsValueOff SettingMirageEditParamsValue = "off"
)

type SettingMirageEditResponseEnvelope struct {
	Errors   []SettingMirageEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMirageEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result ZonesMirage                           `json:"result"`
	JSON   settingMirageEditResponseEnvelopeJSON `json:"-"`
}

// settingMirageEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingMirageEditResponseEnvelope]
type settingMirageEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMirageEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingMirageEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMirageEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMirageEditResponseEnvelopeErrors]
type settingMirageEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMirageEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingMirageEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMirageEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingMirageEditResponseEnvelopeMessages]
type settingMirageEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingMirageGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMirageGetResponseEnvelope struct {
	Errors   []SettingMirageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMirageGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result ZonesMirage                          `json:"result"`
	JSON   settingMirageGetResponseEnvelopeJSON `json:"-"`
}

// settingMirageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingMirageGetResponseEnvelope]
type settingMirageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMirageGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingMirageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMirageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMirageGetResponseEnvelopeErrors]
type settingMirageGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMirageGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingMirageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMirageGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingMirageGetResponseEnvelopeMessages]
type settingMirageGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
