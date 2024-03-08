// File generated from our OpenAPI spec by Stainless.

package zones

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

// SettingMirageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMirageService] method
// instead.
type SettingMirageService struct {
	Options []option.RequestOption
}

// NewSettingMirageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingMirageService(opts ...option.RequestOption) (r *SettingMirageService) {
	r = &SettingMirageService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *SettingMirageService) Edit(ctx context.Context, params SettingMirageEditParams, opts ...option.RequestOption) (res *SettingMirageEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMirageEditResponseEnvelope
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
func (r *SettingMirageService) Get(ctx context.Context, query SettingMirageGetParams, opts ...option.RequestOption) (res *SettingMirageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMirageGetResponseEnvelope
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
type SettingMirageEditResponse struct {
	// ID of the zone setting.
	ID SettingMirageEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMirageEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMirageEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMirageEditResponseJSON `json:"-"`
}

// settingMirageEditResponseJSON contains the JSON metadata for the struct
// [SettingMirageEditResponse]
type settingMirageEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingMirageEditResponseID string

const (
	SettingMirageEditResponseIDMirage SettingMirageEditResponseID = "mirage"
)

// Current value of the zone setting.
type SettingMirageEditResponseValue string

const (
	SettingMirageEditResponseValueOn  SettingMirageEditResponseValue = "on"
	SettingMirageEditResponseValueOff SettingMirageEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMirageEditResponseEditable bool

const (
	SettingMirageEditResponseEditableTrue  SettingMirageEditResponseEditable = true
	SettingMirageEditResponseEditableFalse SettingMirageEditResponseEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingMirageGetResponse struct {
	// ID of the zone setting.
	ID SettingMirageGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMirageGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMirageGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMirageGetResponseJSON `json:"-"`
}

// settingMirageGetResponseJSON contains the JSON metadata for the struct
// [SettingMirageGetResponse]
type settingMirageGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingMirageGetResponseID string

const (
	SettingMirageGetResponseIDMirage SettingMirageGetResponseID = "mirage"
)

// Current value of the zone setting.
type SettingMirageGetResponseValue string

const (
	SettingMirageGetResponseValueOn  SettingMirageGetResponseValue = "on"
	SettingMirageGetResponseValueOff SettingMirageGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMirageGetResponseEditable bool

const (
	SettingMirageGetResponseEditableTrue  SettingMirageGetResponseEditable = true
	SettingMirageGetResponseEditableFalse SettingMirageGetResponseEditable = false
)

type SettingMirageEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMirageEditParamsValue] `json:"value,required"`
}

func (r SettingMirageEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMirageEditParamsValue string

const (
	SettingMirageEditParamsValueOn  SettingMirageEditParamsValue = "on"
	SettingMirageEditParamsValueOff SettingMirageEditParamsValue = "off"
)

type SettingMirageEditResponseEnvelope struct {
	Errors   []SettingMirageEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMirageEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result SettingMirageEditResponse             `json:"result"`
	JSON   settingMirageEditResponseEnvelopeJSON `json:"-"`
}

// settingMirageEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingMirageEditResponseEnvelope]
type settingMirageEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMirageEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingMirageEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMirageEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMirageEditResponseEnvelopeErrors]
type settingMirageEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMirageEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingMirageEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMirageEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingMirageEditResponseEnvelopeMessages]
type settingMirageEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingMirageGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMirageGetResponseEnvelope struct {
	Errors   []SettingMirageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMirageGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result SettingMirageGetResponse             `json:"result"`
	JSON   settingMirageGetResponseEnvelopeJSON `json:"-"`
}

// settingMirageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingMirageGetResponseEnvelope]
type settingMirageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMirageGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingMirageGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMirageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMirageGetResponseEnvelopeErrors]
type settingMirageGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMirageGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingMirageGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMirageGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingMirageGetResponseEnvelopeMessages]
type settingMirageGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMirageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
