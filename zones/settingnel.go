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

// SettingNELService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingNELService] method instead.
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
func (r *SettingNELService) Edit(ctx context.Context, params SettingNELEditParams, opts ...option.RequestOption) (res *NEL, err error) {
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
func (r *SettingNELService) Get(ctx context.Context, query SettingNELGetParams, opts ...option.RequestOption) (res *NEL, err error) {
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
type NEL struct {
	// Zone setting identifier.
	ID NELID `json:"id,required"`
	// Current value of the zone setting.
	Value NELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable NELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       nelJSON   `json:"-"`
}

// nelJSON contains the JSON metadata for the struct [NEL]
type nelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelJSON) RawJSON() string {
	return r.raw
}

// Zone setting identifier.
type NELID string

const (
	NELIDNEL NELID = "nel"
)

func (r NELID) IsKnown() bool {
	switch r {
	case NELIDNEL:
		return true
	}
	return false
}

// Current value of the zone setting.
type NELValue struct {
	Enabled bool         `json:"enabled"`
	JSON    nelValueJSON `json:"-"`
}

// nelValueJSON contains the JSON metadata for the struct [NELValue]
type nelValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type NELEditable bool

const (
	NELEditableTrue  NELEditable = true
	NELEditableFalse NELEditable = false
)

func (r NELEditable) IsKnown() bool {
	switch r {
	case NELEditableTrue, NELEditableFalse:
		return true
	}
	return false
}

// Enable Network Error Logging reporting on your zone. (Beta)
type NELParam struct {
	// Zone setting identifier.
	ID param.Field[NELID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[NELValueParam] `json:"value,required"`
}

func (r NELParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Current value of the zone setting.
type NELValueParam struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r NELValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingNELEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[NELParam] `json:"value,required"`
}

func (r SettingNELEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingNELEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result NEL                                `json:"result"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result NEL                               `json:"result"`
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
