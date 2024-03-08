// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZoneSettingWebpService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingWebpService] method
// instead.
type ZoneSettingWebpService struct {
	Options []option.RequestOption
}

// NewZoneSettingWebpService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingWebpService(opts ...option.RequestOption) (r *ZoneSettingWebpService) {
	r = &ZoneSettingWebpService{}
	r.Options = opts
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *ZoneSettingWebpService) Edit(ctx context.Context, params ZoneSettingWebpEditParams, opts ...option.RequestOption) (res *ZoneSettingWebpEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingWebpEditResponseEnvelope
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
func (r *ZoneSettingWebpService) Get(ctx context.Context, query ZoneSettingWebpGetParams, opts ...option.RequestOption) (res *ZoneSettingWebpGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingWebpGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/webp", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingWebpEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingWebpEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingWebpEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebpEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWebpEditResponseJSON `json:"-"`
}

// zoneSettingWebpEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebpEditResponse]
type zoneSettingWebpEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingWebpEditResponseID string

const (
	ZoneSettingWebpEditResponseIDWebp ZoneSettingWebpEditResponseID = "webp"
)

// Current value of the zone setting.
type ZoneSettingWebpEditResponseValue string

const (
	ZoneSettingWebpEditResponseValueOff ZoneSettingWebpEditResponseValue = "off"
	ZoneSettingWebpEditResponseValueOn  ZoneSettingWebpEditResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebpEditResponseEditable bool

const (
	ZoneSettingWebpEditResponseEditableTrue  ZoneSettingWebpEditResponseEditable = true
	ZoneSettingWebpEditResponseEditableFalse ZoneSettingWebpEditResponseEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingWebpGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingWebpGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingWebpGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebpGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWebpGetResponseJSON `json:"-"`
}

// zoneSettingWebpGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebpGetResponse]
type zoneSettingWebpGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingWebpGetResponseID string

const (
	ZoneSettingWebpGetResponseIDWebp ZoneSettingWebpGetResponseID = "webp"
)

// Current value of the zone setting.
type ZoneSettingWebpGetResponseValue string

const (
	ZoneSettingWebpGetResponseValueOff ZoneSettingWebpGetResponseValue = "off"
	ZoneSettingWebpGetResponseValueOn  ZoneSettingWebpGetResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebpGetResponseEditable bool

const (
	ZoneSettingWebpGetResponseEditableTrue  ZoneSettingWebpGetResponseEditable = true
	ZoneSettingWebpGetResponseEditableFalse ZoneSettingWebpGetResponseEditable = false
)

type ZoneSettingWebpEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingWebpEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingWebpEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingWebpEditParamsValue string

const (
	ZoneSettingWebpEditParamsValueOff ZoneSettingWebpEditParamsValue = "off"
	ZoneSettingWebpEditParamsValueOn  ZoneSettingWebpEditParamsValue = "on"
)

type ZoneSettingWebpEditResponseEnvelope struct {
	Errors   []ZoneSettingWebpEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingWebpEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZoneSettingWebpEditResponse             `json:"result"`
	JSON   zoneSettingWebpEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingWebpEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingWebpEditResponseEnvelope]
type zoneSettingWebpEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWebpEditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingWebpEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingWebpEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingWebpEditResponseEnvelopeErrors]
type zoneSettingWebpEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWebpEditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingWebpEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingWebpEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingWebpEditResponseEnvelopeMessages]
type zoneSettingWebpEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWebpGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingWebpGetResponseEnvelope struct {
	Errors   []ZoneSettingWebpGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingWebpGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZoneSettingWebpGetResponse             `json:"result"`
	JSON   zoneSettingWebpGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingWebpGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingWebpGetResponseEnvelope]
type zoneSettingWebpGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWebpGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingWebpGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingWebpGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingWebpGetResponseEnvelopeErrors]
type zoneSettingWebpGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingWebpGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingWebpGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingWebpGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingWebpGetResponseEnvelopeMessages]
type zoneSettingWebpGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingWebpGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
