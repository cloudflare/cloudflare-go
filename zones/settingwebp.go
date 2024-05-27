// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingWebPService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingWebPService] method instead.
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
func (r *SettingWebPService) Edit(ctx context.Context, params SettingWebPEditParams, opts ...option.RequestOption) (res *WebP, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebPEditResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
func (r *SettingWebPService) Get(ctx context.Context, query SettingWebPGetParams, opts ...option.RequestOption) (res *WebP, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebPGetResponseEnvelope
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
type WebP struct {
	// ID of the zone setting.
	ID WebPID `json:"id,required"`
	// Current value of the zone setting.
	Value WebPValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebPEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       WebPJSON  `json:"-"`
}

// WebPJSON contains the JSON metadata for the struct [WebP]
type WebPJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WebPJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type WebPID string

const (
	WebPIDWebP WebPID = "webp"
)

func (r WebPID) IsKnown() bool {
	switch r {
	case WebPIDWebP:
		return true
	}
	return false
}

// Current value of the zone setting.
type WebPValue string

const (
	WebPValueOff WebPValue = "off"
	WebPValueOn  WebPValue = "on"
)

func (r WebPValue) IsKnown() bool {
	switch r {
	case WebPValueOff, WebPValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type WebPEditable bool

const (
	WebPEditableTrue  WebPEditable = true
	WebPEditableFalse WebPEditable = false
)

func (r WebPEditable) IsKnown() bool {
	switch r {
	case WebPEditableTrue, WebPEditableFalse:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result WebP                                `json:"result"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result WebP                               `json:"result"`
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
