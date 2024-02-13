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
func (r *SettingMirageService) Update(ctx context.Context, zoneID string, body SettingMirageUpdateParams, opts ...option.RequestOption) (res *SettingMirageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMirageUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mirage", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingMirageService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingMirageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMirageGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mirage", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type SettingMirageUpdateResponse struct {
	// ID of the zone setting.
	ID SettingMirageUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMirageUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMirageUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMirageUpdateResponseJSON `json:"-"`
}

// settingMirageUpdateResponseJSON contains the JSON metadata for the struct
// [SettingMirageUpdateResponse]
type settingMirageUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingMirageUpdateResponseID string

const (
	SettingMirageUpdateResponseIDMirage SettingMirageUpdateResponseID = "mirage"
)

// Current value of the zone setting.
type SettingMirageUpdateResponseValue string

const (
	SettingMirageUpdateResponseValueOn  SettingMirageUpdateResponseValue = "on"
	SettingMirageUpdateResponseValueOff SettingMirageUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMirageUpdateResponseEditable bool

const (
	SettingMirageUpdateResponseEditableTrue  SettingMirageUpdateResponseEditable = true
	SettingMirageUpdateResponseEditableFalse SettingMirageUpdateResponseEditable = false
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

type SettingMirageUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingMirageUpdateParamsValue] `json:"value,required"`
}

func (r SettingMirageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMirageUpdateParamsValue string

const (
	SettingMirageUpdateParamsValueOn  SettingMirageUpdateParamsValue = "on"
	SettingMirageUpdateParamsValueOff SettingMirageUpdateParamsValue = "off"
)

type SettingMirageUpdateResponseEnvelope struct {
	Errors   []SettingMirageUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMirageUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result SettingMirageUpdateResponse             `json:"result"`
	JSON   settingMirageUpdateResponseEnvelopeJSON `json:"-"`
}

// settingMirageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMirageUpdateResponseEnvelope]
type settingMirageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMirageUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingMirageUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMirageUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMirageUpdateResponseEnvelopeErrors]
type settingMirageUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMirageUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingMirageUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMirageUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingMirageUpdateResponseEnvelopeMessages]
type settingMirageUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMirageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
