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

// SettingDevelopmentModeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingDevelopmentModeService]
// method instead.
type SettingDevelopmentModeService struct {
	Options []option.RequestOption
}

// NewSettingDevelopmentModeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingDevelopmentModeService(opts ...option.RequestOption) (r *SettingDevelopmentModeService) {
	r = &SettingDevelopmentModeService{}
	r.Options = opts
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *SettingDevelopmentModeService) Update(ctx context.Context, zoneID string, body SettingDevelopmentModeUpdateParams, opts ...option.RequestOption) (res *SettingDevelopmentModeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingDevelopmentModeUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/development_mode", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *SettingDevelopmentModeService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingDevelopmentModeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingDevelopmentModeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/development_mode", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type SettingDevelopmentModeUpdateResponse struct {
	// ID of the zone setting.
	ID SettingDevelopmentModeUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingDevelopmentModeUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingDevelopmentModeUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                  `json:"time_remaining"`
	JSON          settingDevelopmentModeUpdateResponseJSON `json:"-"`
}

// settingDevelopmentModeUpdateResponseJSON contains the JSON metadata for the
// struct [SettingDevelopmentModeUpdateResponse]
type settingDevelopmentModeUpdateResponseJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDevelopmentModeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingDevelopmentModeUpdateResponseID string

const (
	SettingDevelopmentModeUpdateResponseIDDevelopmentMode SettingDevelopmentModeUpdateResponseID = "development_mode"
)

// Current value of the zone setting.
type SettingDevelopmentModeUpdateResponseValue string

const (
	SettingDevelopmentModeUpdateResponseValueOn  SettingDevelopmentModeUpdateResponseValue = "on"
	SettingDevelopmentModeUpdateResponseValueOff SettingDevelopmentModeUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingDevelopmentModeUpdateResponseEditable bool

const (
	SettingDevelopmentModeUpdateResponseEditableTrue  SettingDevelopmentModeUpdateResponseEditable = true
	SettingDevelopmentModeUpdateResponseEditableFalse SettingDevelopmentModeUpdateResponseEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type SettingDevelopmentModeGetResponse struct {
	// ID of the zone setting.
	ID SettingDevelopmentModeGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingDevelopmentModeGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingDevelopmentModeGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                               `json:"time_remaining"`
	JSON          settingDevelopmentModeGetResponseJSON `json:"-"`
}

// settingDevelopmentModeGetResponseJSON contains the JSON metadata for the struct
// [SettingDevelopmentModeGetResponse]
type settingDevelopmentModeGetResponseJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingDevelopmentModeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingDevelopmentModeGetResponseID string

const (
	SettingDevelopmentModeGetResponseIDDevelopmentMode SettingDevelopmentModeGetResponseID = "development_mode"
)

// Current value of the zone setting.
type SettingDevelopmentModeGetResponseValue string

const (
	SettingDevelopmentModeGetResponseValueOn  SettingDevelopmentModeGetResponseValue = "on"
	SettingDevelopmentModeGetResponseValueOff SettingDevelopmentModeGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingDevelopmentModeGetResponseEditable bool

const (
	SettingDevelopmentModeGetResponseEditableTrue  SettingDevelopmentModeGetResponseEditable = true
	SettingDevelopmentModeGetResponseEditableFalse SettingDevelopmentModeGetResponseEditable = false
)

type SettingDevelopmentModeUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingDevelopmentModeUpdateParamsValue] `json:"value,required"`
}

func (r SettingDevelopmentModeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingDevelopmentModeUpdateParamsValue string

const (
	SettingDevelopmentModeUpdateParamsValueOn  SettingDevelopmentModeUpdateParamsValue = "on"
	SettingDevelopmentModeUpdateParamsValueOff SettingDevelopmentModeUpdateParamsValue = "off"
)

type SettingDevelopmentModeUpdateResponseEnvelope struct {
	Errors   []SettingDevelopmentModeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingDevelopmentModeUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result SettingDevelopmentModeUpdateResponse             `json:"result"`
	JSON   settingDevelopmentModeUpdateResponseEnvelopeJSON `json:"-"`
}

// settingDevelopmentModeUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingDevelopmentModeUpdateResponseEnvelope]
type settingDevelopmentModeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingDevelopmentModeUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingDevelopmentModeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingDevelopmentModeUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingDevelopmentModeUpdateResponseEnvelopeErrors]
type settingDevelopmentModeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingDevelopmentModeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingDevelopmentModeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingDevelopmentModeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingDevelopmentModeUpdateResponseEnvelopeMessages]
type settingDevelopmentModeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingDevelopmentModeGetResponseEnvelope struct {
	Errors   []SettingDevelopmentModeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingDevelopmentModeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result SettingDevelopmentModeGetResponse             `json:"result"`
	JSON   settingDevelopmentModeGetResponseEnvelopeJSON `json:"-"`
}

// settingDevelopmentModeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingDevelopmentModeGetResponseEnvelope]
type settingDevelopmentModeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingDevelopmentModeGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingDevelopmentModeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingDevelopmentModeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingDevelopmentModeGetResponseEnvelopeErrors]
type settingDevelopmentModeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingDevelopmentModeGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingDevelopmentModeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingDevelopmentModeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingDevelopmentModeGetResponseEnvelopeMessages]
type settingDevelopmentModeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
