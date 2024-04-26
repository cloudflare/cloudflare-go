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
func (r *SettingMirageService) Edit(ctx context.Context, params SettingMirageEditParams, opts ...option.RequestOption) (res *Mirage, err error) {
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
func (r *SettingMirageService) Get(ctx context.Context, query SettingMirageGetParams, opts ...option.RequestOption) (res *Mirage, err error) {
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
type Mirage struct {
	// ID of the zone setting.
	ID MirageID `json:"id,required"`
	// Current value of the zone setting.
	Value MirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       mirageJSON `json:"-"`
}

// mirageJSON contains the JSON metadata for the struct [Mirage]
type mirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Mirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mirageJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type MirageID string

const (
	MirageIDMirage MirageID = "mirage"
)

func (r MirageID) IsKnown() bool {
	switch r {
	case MirageIDMirage:
		return true
	}
	return false
}

// Current value of the zone setting.
type MirageValue string

const (
	MirageValueOn  MirageValue = "on"
	MirageValueOff MirageValue = "off"
)

func (r MirageValue) IsKnown() bool {
	switch r {
	case MirageValueOn, MirageValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type MirageEditable bool

const (
	MirageEditableTrue  MirageEditable = true
	MirageEditableFalse MirageEditable = false
)

func (r MirageEditable) IsKnown() bool {
	switch r {
	case MirageEditableTrue, MirageEditableFalse:
		return true
	}
	return false
}

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

func (r SettingMirageEditParamsValue) IsKnown() bool {
	switch r {
	case SettingMirageEditParamsValueOn, SettingMirageEditParamsValueOff:
		return true
	}
	return false
}

type SettingMirageEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result Mirage                                `json:"result"`
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

type SettingMirageGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMirageGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result Mirage                               `json:"result"`
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
