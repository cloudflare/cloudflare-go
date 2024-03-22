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
func (r *SettingWebPService) Edit(ctx context.Context, params SettingWebPEditParams, opts ...option.RequestOption) (res *ZonesWebP, err error) {
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
func (r *SettingWebPService) Get(ctx context.Context, query SettingWebPGetParams, opts ...option.RequestOption) (res *ZonesWebP, err error) {
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
type ZonesWebP struct {
	// ID of the zone setting.
	ID ZonesWebPID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesWebPValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesWebPEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesWebPJSON `json:"-"`
}

// zonesWebPJSON contains the JSON metadata for the struct [ZonesWebP]
type zonesWebPJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesWebP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesWebPJSON) RawJSON() string {
	return r.raw
}

func (r ZonesWebP) implementsZonesSettingEditResponse() {}

func (r ZonesWebP) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesWebPID string

const (
	ZonesWebPIDWebP ZonesWebPID = "webp"
)

func (r ZonesWebPID) IsKnown() bool {
	switch r {
	case ZonesWebPIDWebP:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesWebPValue string

const (
	ZonesWebPValueOff ZonesWebPValue = "off"
	ZonesWebPValueOn  ZonesWebPValue = "on"
)

func (r ZonesWebPValue) IsKnown() bool {
	switch r {
	case ZonesWebPValueOff, ZonesWebPValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesWebPEditable bool

const (
	ZonesWebPEditableTrue  ZonesWebPEditable = true
	ZonesWebPEditableFalse ZonesWebPEditable = false
)

func (r ZonesWebPEditable) IsKnown() bool {
	switch r {
	case ZonesWebPEditableTrue, ZonesWebPEditableFalse:
		return true
	}
	return false
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZonesWebPParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesWebPID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesWebPValue] `json:"value,required"`
}

func (r ZonesWebPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesWebPParam) implementsZonesSettingEditParamsItem() {}

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
	Errors   []SettingWebPEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebPEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZonesWebP                           `json:"result"`
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

type SettingWebPEditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingWebPEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebPEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWebPEditResponseEnvelopeErrors]
type settingWebPEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebPEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebPEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWebPEditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingWebPEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebPEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWebPEditResponseEnvelopeMessages]
type settingWebPEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebPEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebPEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingWebPGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingWebPGetResponseEnvelope struct {
	Errors   []SettingWebPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebPGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZonesWebP                          `json:"result"`
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

type SettingWebPGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingWebPGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebPGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWebPGetResponseEnvelopeErrors]
type settingWebPGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebPGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWebPGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingWebPGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebPGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWebPGetResponseEnvelopeMessages]
type settingWebPGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebPGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
