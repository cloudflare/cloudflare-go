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
func (r *ZoneSettingWebpService) Edit(ctx context.Context, params ZoneSettingWebpEditParams, opts ...option.RequestOption) (res *ZonesWebp, err error) {
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
func (r *ZoneSettingWebpService) Get(ctx context.Context, query ZoneSettingWebpGetParams, opts ...option.RequestOption) (res *ZonesWebp, err error) {
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
type ZonesWebp struct {
	// ID of the zone setting.
	ID ZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesWebpJSON `json:"-"`
}

// zonesWebpJSON contains the JSON metadata for the struct [ZonesWebp]
type zonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesWebp) implementsZoneSettingEditResponse() {}

func (r ZonesWebp) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesWebpID string

const (
	ZonesWebpIDWebp ZonesWebpID = "webp"
)

// Current value of the zone setting.
type ZonesWebpValue string

const (
	ZonesWebpValueOff ZonesWebpValue = "off"
	ZonesWebpValueOn  ZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesWebpEditable bool

const (
	ZonesWebpEditableTrue  ZonesWebpEditable = true
	ZonesWebpEditableFalse ZonesWebpEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZonesWebpParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesWebpID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesWebpValue] `json:"value,required"`
}

func (r ZonesWebpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesWebpParam) implementsZoneSettingEditParamsItem() {}

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
	Result ZonesWebp                               `json:"result"`
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
	Result ZonesWebp                              `json:"result"`
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
