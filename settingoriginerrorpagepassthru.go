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

// SettingOriginErrorPagePassThruService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSettingOriginErrorPagePassThruService] method instead.
type SettingOriginErrorPagePassThruService struct {
	Options []option.RequestOption
}

// NewSettingOriginErrorPagePassThruService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOriginErrorPagePassThruService(opts ...option.RequestOption) (r *SettingOriginErrorPagePassThruService) {
	r = &SettingOriginErrorPagePassThruService{}
	r.Options = opts
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *SettingOriginErrorPagePassThruService) Update(ctx context.Context, zoneID string, body SettingOriginErrorPagePassThruUpdateParams, opts ...option.RequestOption) (res *SettingOriginErrorPagePassThruUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginErrorPagePassThruUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *SettingOriginErrorPagePassThruService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingOriginErrorPagePassThruGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOriginErrorPagePassThruGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingOriginErrorPagePassThruUpdateResponse struct {
	// ID of the zone setting.
	ID SettingOriginErrorPagePassThruUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOriginErrorPagePassThruUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOriginErrorPagePassThruUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOriginErrorPagePassThruUpdateResponseJSON `json:"-"`
}

// settingOriginErrorPagePassThruUpdateResponseJSON contains the JSON metadata for
// the struct [SettingOriginErrorPagePassThruUpdateResponse]
type settingOriginErrorPagePassThruUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOriginErrorPagePassThruUpdateResponseID string

const (
	SettingOriginErrorPagePassThruUpdateResponseIDOriginErrorPagePassThru SettingOriginErrorPagePassThruUpdateResponseID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingOriginErrorPagePassThruUpdateResponseValue string

const (
	SettingOriginErrorPagePassThruUpdateResponseValueOn  SettingOriginErrorPagePassThruUpdateResponseValue = "on"
	SettingOriginErrorPagePassThruUpdateResponseValueOff SettingOriginErrorPagePassThruUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOriginErrorPagePassThruUpdateResponseEditable bool

const (
	SettingOriginErrorPagePassThruUpdateResponseEditableTrue  SettingOriginErrorPagePassThruUpdateResponseEditable = true
	SettingOriginErrorPagePassThruUpdateResponseEditableFalse SettingOriginErrorPagePassThruUpdateResponseEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingOriginErrorPagePassThruGetResponse struct {
	// ID of the zone setting.
	ID SettingOriginErrorPagePassThruGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOriginErrorPagePassThruGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOriginErrorPagePassThruGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOriginErrorPagePassThruGetResponseJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseJSON contains the JSON metadata for the
// struct [SettingOriginErrorPagePassThruGetResponse]
type settingOriginErrorPagePassThruGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOriginErrorPagePassThruGetResponseID string

const (
	SettingOriginErrorPagePassThruGetResponseIDOriginErrorPagePassThru SettingOriginErrorPagePassThruGetResponseID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingOriginErrorPagePassThruGetResponseValue string

const (
	SettingOriginErrorPagePassThruGetResponseValueOn  SettingOriginErrorPagePassThruGetResponseValue = "on"
	SettingOriginErrorPagePassThruGetResponseValueOff SettingOriginErrorPagePassThruGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOriginErrorPagePassThruGetResponseEditable bool

const (
	SettingOriginErrorPagePassThruGetResponseEditableTrue  SettingOriginErrorPagePassThruGetResponseEditable = true
	SettingOriginErrorPagePassThruGetResponseEditableFalse SettingOriginErrorPagePassThruGetResponseEditable = false
)

type SettingOriginErrorPagePassThruUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingOriginErrorPagePassThruUpdateParamsValue] `json:"value,required"`
}

func (r SettingOriginErrorPagePassThruUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingOriginErrorPagePassThruUpdateParamsValue string

const (
	SettingOriginErrorPagePassThruUpdateParamsValueOn  SettingOriginErrorPagePassThruUpdateParamsValue = "on"
	SettingOriginErrorPagePassThruUpdateParamsValueOff SettingOriginErrorPagePassThruUpdateParamsValue = "off"
)

type SettingOriginErrorPagePassThruUpdateResponseEnvelope struct {
	Errors   []SettingOriginErrorPagePassThruUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOriginErrorPagePassThruUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result SettingOriginErrorPagePassThruUpdateResponse             `json:"result"`
	JSON   settingOriginErrorPagePassThruUpdateResponseEnvelopeJSON `json:"-"`
}

// settingOriginErrorPagePassThruUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingOriginErrorPagePassThruUpdateResponseEnvelope]
type settingOriginErrorPagePassThruUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOriginErrorPagePassThruUpdateResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    settingOriginErrorPagePassThruUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOriginErrorPagePassThruUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOriginErrorPagePassThruUpdateResponseEnvelopeErrors]
type settingOriginErrorPagePassThruUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOriginErrorPagePassThruUpdateResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    settingOriginErrorPagePassThruUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOriginErrorPagePassThruUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SettingOriginErrorPagePassThruUpdateResponseEnvelopeMessages]
type settingOriginErrorPagePassThruUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOriginErrorPagePassThruGetResponseEnvelope struct {
	Errors   []SettingOriginErrorPagePassThruGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOriginErrorPagePassThruGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result SettingOriginErrorPagePassThruGetResponse             `json:"result"`
	JSON   settingOriginErrorPagePassThruGetResponseEnvelopeJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOriginErrorPagePassThruGetResponseEnvelope]
type settingOriginErrorPagePassThruGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOriginErrorPagePassThruGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingOriginErrorPagePassThruGetResponseEnvelopeErrors]
type settingOriginErrorPagePassThruGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOriginErrorPagePassThruGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    settingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingOriginErrorPagePassThruGetResponseEnvelopeMessages]
type settingOriginErrorPagePassThruGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOriginErrorPagePassThruGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
