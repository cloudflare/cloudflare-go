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

// ZoneSettingMirageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingMirageService] method
// instead.
type ZoneSettingMirageService struct {
	Options []option.RequestOption
}

// NewZoneSettingMirageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingMirageService(opts ...option.RequestOption) (r *ZoneSettingMirageService) {
	r = &ZoneSettingMirageService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *ZoneSettingMirageService) Edit(ctx context.Context, params ZoneSettingMirageEditParams, opts ...option.RequestOption) (res *ZoneSettingMirageEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMirageEditResponseEnvelope
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
func (r *ZoneSettingMirageService) Get(ctx context.Context, query ZoneSettingMirageGetParams, opts ...option.RequestOption) (res *ZoneSettingMirageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMirageGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mirage", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type ZoneSettingMirageEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingMirageEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMirageEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMirageEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMirageEditResponseJSON `json:"-"`
}

// zoneSettingMirageEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingMirageEditResponse]
type zoneSettingMirageEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingMirageEditResponseID string

const (
	ZoneSettingMirageEditResponseIDMirage ZoneSettingMirageEditResponseID = "mirage"
)

// Current value of the zone setting.
type ZoneSettingMirageEditResponseValue string

const (
	ZoneSettingMirageEditResponseValueOn  ZoneSettingMirageEditResponseValue = "on"
	ZoneSettingMirageEditResponseValueOff ZoneSettingMirageEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMirageEditResponseEditable bool

const (
	ZoneSettingMirageEditResponseEditableTrue  ZoneSettingMirageEditResponseEditable = true
	ZoneSettingMirageEditResponseEditableFalse ZoneSettingMirageEditResponseEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingMirageGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingMirageGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMirageGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMirageGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMirageGetResponseJSON `json:"-"`
}

// zoneSettingMirageGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingMirageGetResponse]
type zoneSettingMirageGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingMirageGetResponseID string

const (
	ZoneSettingMirageGetResponseIDMirage ZoneSettingMirageGetResponseID = "mirage"
)

// Current value of the zone setting.
type ZoneSettingMirageGetResponseValue string

const (
	ZoneSettingMirageGetResponseValueOn  ZoneSettingMirageGetResponseValue = "on"
	ZoneSettingMirageGetResponseValueOff ZoneSettingMirageGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMirageGetResponseEditable bool

const (
	ZoneSettingMirageGetResponseEditableTrue  ZoneSettingMirageGetResponseEditable = true
	ZoneSettingMirageGetResponseEditableFalse ZoneSettingMirageGetResponseEditable = false
)

type ZoneSettingMirageEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingMirageEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingMirageEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMirageEditParamsValue string

const (
	ZoneSettingMirageEditParamsValueOn  ZoneSettingMirageEditParamsValue = "on"
	ZoneSettingMirageEditParamsValueOff ZoneSettingMirageEditParamsValue = "off"
)

type ZoneSettingMirageEditResponseEnvelope struct {
	Errors   []ZoneSettingMirageEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMirageEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result ZoneSettingMirageEditResponse             `json:"result"`
	JSON   zoneSettingMirageEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMirageEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingMirageEditResponseEnvelope]
type zoneSettingMirageEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMirageEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingMirageEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMirageEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingMirageEditResponseEnvelopeErrors]
type zoneSettingMirageEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMirageEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingMirageEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMirageEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingMirageEditResponseEnvelopeMessages]
type zoneSettingMirageEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMirageGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingMirageGetResponseEnvelope struct {
	Errors   []ZoneSettingMirageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMirageGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result ZoneSettingMirageGetResponse             `json:"result"`
	JSON   zoneSettingMirageGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMirageGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingMirageGetResponseEnvelope]
type zoneSettingMirageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMirageGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingMirageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMirageGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingMirageGetResponseEnvelopeErrors]
type zoneSettingMirageGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingMirageGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingMirageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMirageGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingMirageGetResponseEnvelopeMessages]
type zoneSettingMirageGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMirageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMirageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
