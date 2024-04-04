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

// SettingWebPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingWebPService] method
// instead.
type SettingWebPService struct {
	Options []option.RequestOption
}

// NewSettingWebPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingWebPService(opts ...option.RequestOption) (r *SettingWebPService) {
	r = &SettingWebPService{}
	r.Options = opts
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *SettingWebPService) Edit(ctx context.Context, params SettingWebPEditParams, opts ...option.RequestOption) (res *ZoneSettingWebP, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebPEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/webp", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *SettingWebPService) Get(ctx context.Context, query SettingWebPGetParams, opts ...option.RequestOption) (res *ZoneSettingWebP, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebPGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/webp", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingWebP struct {
	// ID of the zone setting.
	ID ZoneSettingWebPID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingWebPValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebPEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWebPJSON `json:"-"`
}

// zoneSettingWebPJSON contains the JSON metadata for the struct [ZoneSettingWebP]
type zoneSettingWebPJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebPJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingWebP) implementsZonesSettingEditResponse() {}

func (r ZoneSettingWebP) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingWebPID string

const (
	ZoneSettingWebPIDWebP ZoneSettingWebPID = "webp"
)

func (r ZoneSettingWebPID) IsKnown() bool {
	switch r {
	case ZoneSettingWebPIDWebP:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingWebPValue string

const (
	ZoneSettingWebPValueOff ZoneSettingWebPValue = "off"
	ZoneSettingWebPValueOn  ZoneSettingWebPValue = "on"
)

func (r ZoneSettingWebPValue) IsKnown() bool {
	switch r {
	case ZoneSettingWebPValueOff, ZoneSettingWebPValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebPEditable bool

const (
	ZoneSettingWebPEditableTrue  ZoneSettingWebPEditable = true
	ZoneSettingWebPEditableFalse ZoneSettingWebPEditable = false
)

func (r ZoneSettingWebPEditable) IsKnown() bool {
	switch r {
	case ZoneSettingWebPEditableTrue, ZoneSettingWebPEditableFalse:
		return true
	}
	return false
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingWebPParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingWebPID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingWebPValue] `json:"value,required"`
}

func (r ZoneSettingWebPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingWebPParam) implementsZonesSettingEditParamsItemUnion() {}

type SettingWebPEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingWebPEditParamsValue] `json:"value,required"`
}

func (r SettingWebPEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingWebPEditParamsValue string

const (
	SettingWebPEditParamsValueOff SettingWebPEditParamsValue = "off"
	SettingWebPEditParamsValueOn  SettingWebPEditParamsValue = "on"
)

func (r SettingWebPEditParamsValue) IsKnown() bool {
	switch r {
	case SettingWebPEditParamsValueOff, SettingWebPEditParamsValueOn:
		return true
	}
	return false
}

type SettingWebPEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZoneSettingWebP                     `json:"result"`
	JSON   settingWebPEditResponseEnvelopeJSON `json:"-"`
}

// settingWebPEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWebPEditResponseEnvelope]
type settingWebPEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebPEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebPEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingWebPGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingWebPGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZoneSettingWebP                    `json:"result"`
	JSON   settingWebPGetResponseEnvelopeJSON `json:"-"`
}

// settingWebPGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWebPGetResponseEnvelope]
type settingWebPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebPGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
