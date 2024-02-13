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

// SettingTrueClientIPHeaderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingTrueClientIPHeaderService] method instead.
type SettingTrueClientIPHeaderService struct {
	Options []option.RequestOption
}

// NewSettingTrueClientIPHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingTrueClientIPHeaderService(opts ...option.RequestOption) (r *SettingTrueClientIPHeaderService) {
	r = &SettingTrueClientIPHeaderService{}
	r.Options = opts
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *SettingTrueClientIPHeaderService) Update(ctx context.Context, zoneID string, body SettingTrueClientIPHeaderUpdateParams, opts ...option.RequestOption) (res *SettingTrueClientIPHeaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTrueClientIPHeaderUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *SettingTrueClientIPHeaderService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingTrueClientIPHeaderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingTrueClientIPHeaderGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingTrueClientIPHeaderUpdateResponse struct {
	// ID of the zone setting.
	ID SettingTrueClientIPHeaderUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTrueClientIPHeaderUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTrueClientIPHeaderUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTrueClientIPHeaderUpdateResponseJSON `json:"-"`
}

// settingTrueClientIPHeaderUpdateResponseJSON contains the JSON metadata for the
// struct [SettingTrueClientIPHeaderUpdateResponse]
type settingTrueClientIPHeaderUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingTrueClientIPHeaderUpdateResponseID string

const (
	SettingTrueClientIPHeaderUpdateResponseIDTrueClientIPHeader SettingTrueClientIPHeaderUpdateResponseID = "true_client_ip_header"
)

// Current value of the zone setting.
type SettingTrueClientIPHeaderUpdateResponseValue string

const (
	SettingTrueClientIPHeaderUpdateResponseValueOn  SettingTrueClientIPHeaderUpdateResponseValue = "on"
	SettingTrueClientIPHeaderUpdateResponseValueOff SettingTrueClientIPHeaderUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTrueClientIPHeaderUpdateResponseEditable bool

const (
	SettingTrueClientIPHeaderUpdateResponseEditableTrue  SettingTrueClientIPHeaderUpdateResponseEditable = true
	SettingTrueClientIPHeaderUpdateResponseEditableFalse SettingTrueClientIPHeaderUpdateResponseEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingTrueClientIPHeaderGetResponse struct {
	// ID of the zone setting.
	ID SettingTrueClientIPHeaderGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingTrueClientIPHeaderGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingTrueClientIPHeaderGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingTrueClientIPHeaderGetResponseJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseJSON contains the JSON metadata for the
// struct [SettingTrueClientIPHeaderGetResponse]
type settingTrueClientIPHeaderGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingTrueClientIPHeaderGetResponseID string

const (
	SettingTrueClientIPHeaderGetResponseIDTrueClientIPHeader SettingTrueClientIPHeaderGetResponseID = "true_client_ip_header"
)

// Current value of the zone setting.
type SettingTrueClientIPHeaderGetResponseValue string

const (
	SettingTrueClientIPHeaderGetResponseValueOn  SettingTrueClientIPHeaderGetResponseValue = "on"
	SettingTrueClientIPHeaderGetResponseValueOff SettingTrueClientIPHeaderGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingTrueClientIPHeaderGetResponseEditable bool

const (
	SettingTrueClientIPHeaderGetResponseEditableTrue  SettingTrueClientIPHeaderGetResponseEditable = true
	SettingTrueClientIPHeaderGetResponseEditableFalse SettingTrueClientIPHeaderGetResponseEditable = false
)

type SettingTrueClientIPHeaderUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingTrueClientIPHeaderUpdateParamsValue] `json:"value,required"`
}

func (r SettingTrueClientIPHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingTrueClientIPHeaderUpdateParamsValue string

const (
	SettingTrueClientIPHeaderUpdateParamsValueOn  SettingTrueClientIPHeaderUpdateParamsValue = "on"
	SettingTrueClientIPHeaderUpdateParamsValueOff SettingTrueClientIPHeaderUpdateParamsValue = "off"
)

type SettingTrueClientIPHeaderUpdateResponseEnvelope struct {
	Errors   []SettingTrueClientIPHeaderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTrueClientIPHeaderUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result SettingTrueClientIPHeaderUpdateResponse             `json:"result"`
	JSON   settingTrueClientIPHeaderUpdateResponseEnvelopeJSON `json:"-"`
}

// settingTrueClientIPHeaderUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingTrueClientIPHeaderUpdateResponseEnvelope]
type settingTrueClientIPHeaderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTrueClientIPHeaderUpdateResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingTrueClientIPHeaderUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTrueClientIPHeaderUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingTrueClientIPHeaderUpdateResponseEnvelopeErrors]
type settingTrueClientIPHeaderUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTrueClientIPHeaderUpdateResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingTrueClientIPHeaderUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTrueClientIPHeaderUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingTrueClientIPHeaderUpdateResponseEnvelopeMessages]
type settingTrueClientIPHeaderUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTrueClientIPHeaderGetResponseEnvelope struct {
	Errors   []SettingTrueClientIPHeaderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTrueClientIPHeaderGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result SettingTrueClientIPHeaderGetResponse             `json:"result"`
	JSON   settingTrueClientIPHeaderGetResponseEnvelopeJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTrueClientIPHeaderGetResponseEnvelope]
type settingTrueClientIPHeaderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTrueClientIPHeaderGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingTrueClientIPHeaderGetResponseEnvelopeErrors]
type settingTrueClientIPHeaderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingTrueClientIPHeaderGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingTrueClientIPHeaderGetResponseEnvelopeMessages]
type settingTrueClientIPHeaderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrueClientIPHeaderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
