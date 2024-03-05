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

// ZoneSettingNELService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingNELService] method
// instead.
type ZoneSettingNELService struct {
	Options []option.RequestOption
}

// NewZoneSettingNELService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingNELService(opts ...option.RequestOption) (r *ZoneSettingNELService) {
	r = &ZoneSettingNELService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/nel-solving-mobile-speed)
// for more information.
func (r *ZoneSettingNELService) Edit(ctx context.Context, params ZoneSettingNELEditParams, opts ...option.RequestOption) (res *ZonesNEL, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingNELEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
func (r *ZoneSettingNELService) Get(ctx context.Context, query ZoneSettingNELGetParams, opts ...option.RequestOption) (res *ZonesNEL, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingNELGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZonesNEL struct {
	// Zone setting identifier.
	ID ZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time    `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesNELJSON `json:"-"`
}

// zonesNELJSON contains the JSON metadata for the struct [ZonesNEL]
type zonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesNEL) implementsZoneSettingEditResponse() {}

func (r ZonesNEL) implementsZoneSettingGetResponse() {}

// Zone setting identifier.
type ZonesNELID string

const (
	ZonesNELIDNEL ZonesNELID = "nel"
)

// Current value of the zone setting.
type ZonesNELValue struct {
	Enabled bool              `json:"enabled"`
	JSON    zonesNELValueJSON `json:"-"`
}

// zonesNELValueJSON contains the JSON metadata for the struct [ZonesNELValue]
type zonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesNELEditable bool

const (
	ZonesNELEditableTrue  ZonesNELEditable = true
	ZonesNELEditableFalse ZonesNELEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZonesNELParam struct {
	// Zone setting identifier.
	ID param.Field[ZonesNELID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesNELValueParam] `json:"value,required"`
}

func (r ZonesNELParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesNELParam) implementsZoneSettingEditParamsItem() {}

// Current value of the zone setting.
type ZonesNELValueParam struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZonesNELValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingNELEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[ZonesNELParam] `json:"value,required"`
}

func (r ZoneSettingNELEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingNELEditResponseEnvelope struct {
	Errors   []ZoneSettingNELEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingNELEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZonesNEL                               `json:"result"`
	JSON   zoneSettingNELEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingNELEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingNELEditResponseEnvelope]
type zoneSettingNELEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELEditResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingNELEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingNELEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingNELEditResponseEnvelopeErrors]
type zoneSettingNELEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELEditResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingNELEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingNELEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingNELEditResponseEnvelopeMessages]
type zoneSettingNELEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingNELGetResponseEnvelope struct {
	Errors   []ZoneSettingNELGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingNELGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZonesNEL                              `json:"result"`
	JSON   zoneSettingNELGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingNELGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingNELGetResponseEnvelope]
type zoneSettingNELGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingNELGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingNELGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingNELGetResponseEnvelopeErrors]
type zoneSettingNELGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingNELGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingNELGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingNELGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneSettingNELGetResponseEnvelopeMessages]
type zoneSettingNELGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
