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
func (r *SettingNELService) Update(ctx context.Context, zoneID string, body SettingNELUpdateParams, opts ...option.RequestOption) (res *SettingNELUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingNELUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
func (r *SettingNELService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingNELGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingNELGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/nel", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingNELUpdateResponse struct {
	// Zone setting identifier.
	ID SettingNELUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingNELUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingNELUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingNELUpdateResponseJSON `json:"-"`
}

// settingNELUpdateResponseJSON contains the JSON metadata for the struct
// [SettingNELUpdateResponse]
type settingNELUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type SettingNELUpdateResponseID string

const (
	SettingNELUpdateResponseIDNEL SettingNELUpdateResponseID = "nel"
)

// Current value of the zone setting.
type SettingNELUpdateResponseValue struct {
	Enabled bool                              `json:"enabled"`
	JSON    settingNELUpdateResponseValueJSON `json:"-"`
}

// settingNELUpdateResponseValueJSON contains the JSON metadata for the struct
// [SettingNELUpdateResponseValue]
type settingNELUpdateResponseValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELUpdateResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingNELUpdateResponseEditable bool

const (
	SettingNELUpdateResponseEditableTrue  SettingNELUpdateResponseEditable = true
	SettingNELUpdateResponseEditableFalse SettingNELUpdateResponseEditable = false
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

type SettingNELUpdateParams struct {
	// Enable Network Error Logging reporting on your zone. (Beta)
	Value param.Field[SettingNELUpdateParamsValue] `json:"value,required"`
}

func (r SettingNELUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingNELUpdateParamsValue struct {
	// Zone setting identifier.
	ID param.Field[SettingNELUpdateParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingNELUpdateParamsValueValue] `json:"value,required"`
}

func (r SettingNELUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Zone setting identifier.
type SettingNELUpdateParamsValueID string

const (
	SettingNELUpdateParamsValueIDNEL SettingNELUpdateParamsValueID = "nel"
)

// Current value of the zone setting.
type SettingNELUpdateParamsValueValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingNELUpdateParamsValueValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingNELUpdateParamsValueEditable bool

const (
	SettingNELUpdateParamsValueEditableTrue  SettingNELUpdateParamsValueEditable = true
	SettingNELUpdateParamsValueEditableFalse SettingNELUpdateParamsValueEditable = false
)

type SettingNELUpdateResponseEnvelope struct {
	Errors   []SettingNELUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingNELUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result SettingNELUpdateResponse             `json:"result"`
	JSON   settingNELUpdateResponseEnvelopeJSON `json:"-"`
}

// settingNELUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingNELUpdateResponseEnvelope]
type settingNELUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingNELUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingNELUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingNELUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingNELUpdateResponseEnvelopeErrors]
type settingNELUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingNELUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingNELUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingNELUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingNELUpdateResponseEnvelopeMessages]
type settingNELUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingNELUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
