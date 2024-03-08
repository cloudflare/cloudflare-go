// File generated from our OpenAPI spec by Stainless.

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

// SettingNELService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingNELService] method instead.
type SettingNELService struct {
	Options []option.RequestOption
}

// NewSettingNELService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingNELService(opts ...option.RequestOption) (r *SettingNELService) {
	r = &SettingNELService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/nel-solving-mobile-speed)
// for more information.
func (r *SettingNELService) Edit(ctx context.Context, params SettingNELEditParams, opts ...option.RequestOption) (res *ZonesNEL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingNELEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
func (r *SettingNELService) Get(ctx context.Context, query SettingNELGetParams, opts ...option.RequestOption) (res *ZonesNEL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingNELGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZonesNEL struct {
	// Zone setting identifier.
	ID ZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time    `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesNELJSON `json:"-"`
}

// zonesNELJSON contains the JSON metadata for the struct [ZonesNEL]
type zonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesNELJSON) RawJSON() string {
	return r.raw
}

func (r ZonesNEL) implementsZonesSettingEditResponse() {}

func (r ZonesNEL) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type ZonesNELID string

const (
	ZonesNELIDNEL ZonesNELID = "nel"
)

// Current value of the zone setting.
type ZonesNELValue struct {
	Enabled bool              `json:"enabled"`
	JSON    zonesNELValueJSON `json:"-"`
}

// zonesNELValueJSON contains the JSON metadata for the struct [ZonesNELValue]
type zonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesNELValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesNELEditable bool

const (
	ZonesNELEditableTrue  ZonesNELEditable = true
	ZonesNELEditableFalse ZonesNELEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZonesNELParam struct {
	// Zone setting identifier.
	ID param.Field[ZonesNELID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesNELValueParam] `json:"value,required"`
}

func (r ZonesNELParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesNELParam) implementsZonesSettingEditParamsItem() {}

// Current value of the zone setting.
type ZonesNELValueParam struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZonesNELValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingNELEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[ZonesNELParam] `json:"value,required"`
}

func (r SettingNELEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingNELEditResponseEnvelope struct {
	Errors   []SettingNELEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingNELEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZonesNEL                           `json:"result"`
	JSON   settingNELEditResponseEnvelopeJSON `json:"-"`
}

// settingNELEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingNELEditResponseEnvelope]
type settingNELEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingNELEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingNELEditResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingNELEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingNELEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingNELEditResponseEnvelopeErrors]
type settingNELEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingNELEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingNELEditResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingNELEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingNELEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingNELEditResponseEnvelopeMessages]
type settingNELEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingNELEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingNELGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingNELGetResponseEnvelope struct {
	Errors   []SettingNELGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingNELGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZonesNEL                          `json:"result"`
	JSON   settingNELGetResponseEnvelopeJSON `json:"-"`
}

// settingNELGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingNELGetResponseEnvelope]
type settingNELGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingNELGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingNELGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingNELGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingNELGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingNELGetResponseEnvelopeErrors]
type settingNELGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingNELGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingNELGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingNELGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingNELGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingNELGetResponseEnvelopeMessages]
type settingNELGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingNELGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
