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
func (r *ZoneSettingNELService) Edit(ctx context.Context, params ZoneSettingNELEditParams, opts ...option.RequestOption) (res *ZoneSettingNELEditResponse, err error) {
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
func (r *ZoneSettingNELService) Get(ctx context.Context, query ZoneSettingNELGetParams, opts ...option.RequestOption) (res *ZoneSettingNELGetResponse, err error) {
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
type ZoneSettingNELEditResponse struct {
	// Zone setting identifier.
	ID ZoneSettingNELEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingNELEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingNELEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingNELEditResponseJSON `json:"-"`
}

// zoneSettingNELEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingNELEditResponse]
type zoneSettingNELEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingNELEditResponseJSON) RawJSON() string {
	return r.raw
}

// Zone setting identifier.
type ZoneSettingNELEditResponseID string

const (
	ZoneSettingNELEditResponseIDNEL ZoneSettingNELEditResponseID = "nel"
)

// Current value of the zone setting.
type ZoneSettingNELEditResponseValue struct {
	Enabled bool                                `json:"enabled"`
	JSON    zoneSettingNELEditResponseValueJSON `json:"-"`
}

// zoneSettingNELEditResponseValueJSON contains the JSON metadata for the struct
// [ZoneSettingNELEditResponseValue]
type zoneSettingNELEditResponseValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingNELEditResponseValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNELEditResponseEditable bool

const (
	ZoneSettingNELEditResponseEditableTrue  ZoneSettingNELEditResponseEditable = true
	ZoneSettingNELEditResponseEditableFalse ZoneSettingNELEditResponseEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNELGetResponse struct {
	// Zone setting identifier.
	ID ZoneSettingNELGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingNELGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingNELGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingNELGetResponseJSON `json:"-"`
}

// zoneSettingNELGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingNELGetResponse]
type zoneSettingNELGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingNELGetResponseJSON) RawJSON() string {
	return r.raw
}

// Zone setting identifier.
type ZoneSettingNELGetResponseID string

const (
	ZoneSettingNELGetResponseIDNEL ZoneSettingNELGetResponseID = "nel"
)

// Current value of the zone setting.
type ZoneSettingNELGetResponseValue struct {
	Enabled bool                               `json:"enabled"`
	JSON    zoneSettingNELGetResponseValueJSON `json:"-"`
}

// zoneSettingNELGetResponseValueJSON contains the JSON metadata for the struct
// [ZoneSettingNELGetResponseValue]
type zoneSettingNELGetResponseValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingNELGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingNELGetResponseValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNELGetResponseEditable bool

const (
	ZoneSettingNELGetResponseEditableTrue  ZoneSettingNELGetResponseEditable = true
	ZoneSettingNELGetResponseEditableFalse ZoneSettingNELGetResponseEditable = false
)

type ZoneSettingNELEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[ZoneSettingNELEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingNELEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingNELEditParamsValue struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingNELEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingNELEditParamsValueValue] `json:"value,required"`
}

func (r ZoneSettingNELEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Zone setting identifier.
type ZoneSettingNELEditParamsValueID string

const (
	ZoneSettingNELEditParamsValueIDNEL ZoneSettingNELEditParamsValueID = "nel"
)

// Current value of the zone setting.
type ZoneSettingNELEditParamsValueValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingNELEditParamsValueValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingNELEditParamsValueEditable bool

const (
	ZoneSettingNELEditParamsValueEditableTrue  ZoneSettingNELEditParamsValueEditable = true
	ZoneSettingNELEditParamsValueEditableFalse ZoneSettingNELEditParamsValueEditable = false
)

type ZoneSettingNELEditResponseEnvelope struct {
	Errors   []ZoneSettingNELEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingNELEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result ZoneSettingNELEditResponse             `json:"result"`
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

func (r zoneSettingNELEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingNELEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingNELEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingNELGetResponse             `json:"result"`
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

func (r zoneSettingNELGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingNELGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingNELGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
