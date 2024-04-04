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
func (r *SettingNELService) Edit(ctx context.Context, params SettingNELEditParams, opts ...option.RequestOption) (res *ZoneSettingNEL, err error) {
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
func (r *SettingNELService) Get(ctx context.Context, query SettingNELGetParams, opts ...option.RequestOption) (res *ZoneSettingNEL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingNELGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNEL struct {
	// Zone setting identifier.
	ID ZoneSettingNELID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingNELJSON `json:"-"`
}

// zoneSettingNELJSON contains the JSON metadata for the struct [ZoneSettingNEL]
type zoneSettingNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingNELJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingNEL) implementsZonesSettingEditResponse() {}

func (r ZoneSettingNEL) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type ZoneSettingNELID string

const (
	ZoneSettingNELIDNEL ZoneSettingNELID = "nel"
)

func (r ZoneSettingNELID) IsKnown() bool {
	switch r {
	case ZoneSettingNELIDNEL:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingNELValue struct {
	Enabled bool                    `json:"enabled"`
	JSON    zoneSettingNELValueJSON `json:"-"`
}

// zoneSettingNELValueJSON contains the JSON metadata for the struct
// [ZoneSettingNELValue]
type zoneSettingNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingNELValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNELEditable bool

const (
	ZoneSettingNELEditableTrue  ZoneSettingNELEditable = true
	ZoneSettingNELEditableFalse ZoneSettingNELEditable = false
)

func (r ZoneSettingNELEditable) IsKnown() bool {
	switch r {
	case ZoneSettingNELEditableTrue, ZoneSettingNELEditableFalse:
		return true
	}
	return false
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNELParam struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingNELID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingNELValueParam] `json:"value,required"`
}

func (r ZoneSettingNELParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingNELParam) implementsZonesSettingEditParamsItem() {}

// Current value of the zone setting.
type ZoneSettingNELValueParam struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingNELValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingNELEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[ZoneSettingNELParam] `json:"value,required"`
}

func (r SettingNELEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingNELEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZoneSettingNEL                     `json:"result"`
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

type SettingNELGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingNELGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZoneSettingNEL                    `json:"result"`
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
