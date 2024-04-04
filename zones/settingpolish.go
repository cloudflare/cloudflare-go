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

// SettingPolishService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingPolishService] method
// instead.
type SettingPolishService struct {
	Options []option.RequestOption
}

// NewSettingPolishService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingPolishService(opts ...option.RequestOption) (r *SettingPolishService) {
	r = &SettingPolishService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *SettingPolishService) Edit(ctx context.Context, params SettingPolishEditParams, opts ...option.RequestOption) (res *ZoneSettingPolish, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPolishEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/polish", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *SettingPolishService) Get(ctx context.Context, query SettingPolishGetParams, opts ...option.RequestOption) (res *ZoneSettingPolish, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPolishGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/polish", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingPolish struct {
	// ID of the zone setting.
	ID ZoneSettingPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPolishJSON `json:"-"`
}

// zoneSettingPolishJSON contains the JSON metadata for the struct
// [ZoneSettingPolish]
type zoneSettingPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingPolishID string

const (
	ZoneSettingPolishIDPolish ZoneSettingPolishID = "polish"
)

func (r ZoneSettingPolishID) IsKnown() bool {
	switch r {
	case ZoneSettingPolishIDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingPolishValue string

const (
	ZoneSettingPolishValueOff      ZoneSettingPolishValue = "off"
	ZoneSettingPolishValueLossless ZoneSettingPolishValue = "lossless"
	ZoneSettingPolishValueLossy    ZoneSettingPolishValue = "lossy"
)

func (r ZoneSettingPolishValue) IsKnown() bool {
	switch r {
	case ZoneSettingPolishValueOff, ZoneSettingPolishValueLossless, ZoneSettingPolishValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPolishEditable bool

const (
	ZoneSettingPolishEditableTrue  ZoneSettingPolishEditable = true
	ZoneSettingPolishEditableFalse ZoneSettingPolishEditable = false
)

func (r ZoneSettingPolishEditable) IsKnown() bool {
	switch r {
	case ZoneSettingPolishEditableTrue, ZoneSettingPolishEditableFalse:
		return true
	}
	return false
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingPolishParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingPolishID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingPolishValue] `json:"value,required"`
}

func (r ZoneSettingPolishParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingPolishEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Value param.Field[ZoneSettingPolishParam] `json:"value,required"`
}

func (r SettingPolishEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingPolishEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result ZoneSettingPolish                     `json:"result"`
	JSON   settingPolishEditResponseEnvelopeJSON `json:"-"`
}

// settingPolishEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingPolishEditResponseEnvelope]
type settingPolishEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPolishEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingPolishGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingPolishGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result ZoneSettingPolish                    `json:"result"`
	JSON   settingPolishGetResponseEnvelopeJSON `json:"-"`
}

// settingPolishGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingPolishGetResponseEnvelope]
type settingPolishGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPolishGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
