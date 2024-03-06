// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingFontSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingFontSettingService]
// method instead.
type ZoneSettingFontSettingService struct {
	Options []option.RequestOption
}

// NewZoneSettingFontSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingFontSettingService(opts ...option.RequestOption) (r *ZoneSettingFontSettingService) {
	r = &ZoneSettingFontSettingService{}
	r.Options = opts
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *ZoneSettingFontSettingService) Edit(ctx context.Context, params ZoneSettingFontSettingEditParams, opts ...option.RequestOption) (res *ZoneSettingFontSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingFontSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *ZoneSettingFontSettingService) Get(ctx context.Context, query ZoneSettingFontSettingGetParams, opts ...option.RequestOption) (res *ZoneSettingFontSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingFontSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type ZoneSettingFontSettingEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingFontSettingEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingFontSettingEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingFontSettingEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingFontSettingEditResponseJSON `json:"-"`
}

// zoneSettingFontSettingEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingFontSettingEditResponse]
type zoneSettingFontSettingEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingFontSettingEditResponseID string

const (
	ZoneSettingFontSettingEditResponseIDFonts ZoneSettingFontSettingEditResponseID = "fonts"
)

// Current value of the zone setting.
type ZoneSettingFontSettingEditResponseValue string

const (
	ZoneSettingFontSettingEditResponseValueOn  ZoneSettingFontSettingEditResponseValue = "on"
	ZoneSettingFontSettingEditResponseValueOff ZoneSettingFontSettingEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingFontSettingEditResponseEditable bool

const (
	ZoneSettingFontSettingEditResponseEditableTrue  ZoneSettingFontSettingEditResponseEditable = true
	ZoneSettingFontSettingEditResponseEditableFalse ZoneSettingFontSettingEditResponseEditable = false
)

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type ZoneSettingFontSettingGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingFontSettingGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingFontSettingGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingFontSettingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingFontSettingGetResponseJSON `json:"-"`
}

// zoneSettingFontSettingGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingFontSettingGetResponse]
type zoneSettingFontSettingGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingFontSettingGetResponseID string

const (
	ZoneSettingFontSettingGetResponseIDFonts ZoneSettingFontSettingGetResponseID = "fonts"
)

// Current value of the zone setting.
type ZoneSettingFontSettingGetResponseValue string

const (
	ZoneSettingFontSettingGetResponseValueOn  ZoneSettingFontSettingGetResponseValue = "on"
	ZoneSettingFontSettingGetResponseValueOff ZoneSettingFontSettingGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingFontSettingGetResponseEditable bool

const (
	ZoneSettingFontSettingGetResponseEditableTrue  ZoneSettingFontSettingGetResponseEditable = true
	ZoneSettingFontSettingGetResponseEditableFalse ZoneSettingFontSettingGetResponseEditable = false
)

type ZoneSettingFontSettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether the feature is enabled or disabled.
	Value param.Field[ZoneSettingFontSettingEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingFontSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the feature is enabled or disabled.
type ZoneSettingFontSettingEditParamsValue string

const (
	ZoneSettingFontSettingEditParamsValueOn  ZoneSettingFontSettingEditParamsValue = "on"
	ZoneSettingFontSettingEditParamsValueOff ZoneSettingFontSettingEditParamsValue = "off"
)

type ZoneSettingFontSettingEditResponseEnvelope struct {
	Errors   []ZoneSettingFontSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingFontSettingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result ZoneSettingFontSettingEditResponse             `json:"result"`
	JSON   zoneSettingFontSettingEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingFontSettingEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingFontSettingEditResponseEnvelope]
type zoneSettingFontSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontSettingEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingFontSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingFontSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingFontSettingEditResponseEnvelopeErrors]
type zoneSettingFontSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontSettingEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingFontSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingFontSettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingFontSettingEditResponseEnvelopeMessages]
type zoneSettingFontSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingFontSettingGetResponseEnvelope struct {
	Errors   []ZoneSettingFontSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingFontSettingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result ZoneSettingFontSettingGetResponse             `json:"result"`
	JSON   zoneSettingFontSettingGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingFontSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingFontSettingGetResponseEnvelope]
type zoneSettingFontSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontSettingGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingFontSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingFontSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingFontSettingGetResponseEnvelopeErrors]
type zoneSettingFontSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingFontSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingFontSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingFontSettingGetResponseEnvelopeMessages]
type zoneSettingFontSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
