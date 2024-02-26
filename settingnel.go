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

// SettingNELService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingNELService] method instead.
type SettingNELService struct {
	Options []option.RequestOption
}

// NewSettingNELService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingNELService(opts ...option.RequestOption) (r *SettingNELService) {
	r = &SettingNELService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/nel-solving-mobile-speed)
// for more information.
func (r *SettingNELService) Edit(ctx context.Context, params SettingNELEditParams, opts ...option.RequestOption) (res *SettingNELEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingNELEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
func (r *SettingNELService) Get(ctx context.Context, query SettingNELGetParams, opts ...option.RequestOption) (res *SettingNELGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingNELGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingNELEditResponse struct {
	// Zone setting identifier.
	ID SettingNELEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingNELEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingNELEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingNELEditResponseJSON `json:"-"`
}

// settingNELEditResponseJSON contains the JSON metadata for the struct
// [SettingNELEditResponse]
type settingNELEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type SettingNELEditResponseID string

const (
	SettingNELEditResponseIDNEL SettingNELEditResponseID = "nel"
)

// Current value of the zone setting.
type SettingNELEditResponseValue struct {
	Enabled bool                            `json:"enabled"`
	JSON    settingNELEditResponseValueJSON `json:"-"`
}

// settingNELEditResponseValueJSON contains the JSON metadata for the struct
// [SettingNELEditResponseValue]
type settingNELEditResponseValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingNELEditResponseEditable bool

const (
	SettingNELEditResponseEditableTrue  SettingNELEditResponseEditable = true
	SettingNELEditResponseEditableFalse SettingNELEditResponseEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingNELGetResponse struct {
	// Zone setting identifier.
	ID SettingNELGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingNELGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingNELGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingNELGetResponseJSON `json:"-"`
}

// settingNELGetResponseJSON contains the JSON metadata for the struct
// [SettingNELGetResponse]
type settingNELGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type SettingNELGetResponseID string

const (
	SettingNELGetResponseIDNEL SettingNELGetResponseID = "nel"
)

// Current value of the zone setting.
type SettingNELGetResponseValue struct {
	Enabled bool                           `json:"enabled"`
	JSON    settingNELGetResponseValueJSON `json:"-"`
}

// settingNELGetResponseValueJSON contains the JSON metadata for the struct
// [SettingNELGetResponseValue]
type settingNELGetResponseValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingNELGetResponseEditable bool

const (
	SettingNELGetResponseEditableTrue  SettingNELGetResponseEditable = true
	SettingNELGetResponseEditableFalse SettingNELGetResponseEditable = false
)

type SettingNELEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[SettingNELEditParamsValue] `json:"value,required"`
}

func (r SettingNELEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingNELEditParamsValue struct {
	// Zone setting identifier.
	ID param.Field[SettingNELEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingNELEditParamsValueValue] `json:"value,required"`
}

func (r SettingNELEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Zone setting identifier.
type SettingNELEditParamsValueID string

const (
	SettingNELEditParamsValueIDNEL SettingNELEditParamsValueID = "nel"
)

// Current value of the zone setting.
type SettingNELEditParamsValueValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingNELEditParamsValueValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingNELEditParamsValueEditable bool

const (
	SettingNELEditParamsValueEditableTrue  SettingNELEditParamsValueEditable = true
	SettingNELEditParamsValueEditableFalse SettingNELEditParamsValueEditable = false
)

type SettingNELEditResponseEnvelope struct {
	Errors   []SettingNELEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingNELEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result SettingNELEditResponse             `json:"result"`
	JSON   settingNELEditResponseEnvelopeJSON `json:"-"`
}

// settingNELEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingNELEditResponseEnvelope]
type settingNELEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingNELEditResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingNELEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingNELEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingNELEditResponseEnvelopeErrors]
type settingNELEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingNELEditResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingNELEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingNELEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingNELEditResponseEnvelopeMessages]
type settingNELEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingNELGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingNELGetResponseEnvelope struct {
	Errors   []SettingNELGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingNELGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result SettingNELGetResponse             `json:"result"`
	JSON   settingNELGetResponseEnvelopeJSON `json:"-"`
}

// settingNELGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingNELGetResponseEnvelope]
type settingNELGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingNELGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingNELGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingNELGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingNELGetResponseEnvelopeErrors]
type settingNELGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingNELGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingNELGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingNELGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingNELGetResponseEnvelopeMessages]
type settingNELGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
