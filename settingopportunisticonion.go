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

// SettingOpportunisticOnionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingOpportunisticOnionService] method instead.
type SettingOpportunisticOnionService struct {
	Options []option.RequestOption
}

// NewSettingOpportunisticOnionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingOpportunisticOnionService(opts ...option.RequestOption) (r *SettingOpportunisticOnionService) {
	r = &SettingOpportunisticOnionService{}
	r.Options = opts
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *SettingOpportunisticOnionService) Update(ctx context.Context, zoneID string, body SettingOpportunisticOnionUpdateParams, opts ...option.RequestOption) (res *SettingOpportunisticOnionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticOnionUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *SettingOpportunisticOnionService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingOpportunisticOnionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOpportunisticOnionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type SettingOpportunisticOnionUpdateResponse struct {
	// ID of the zone setting.
	ID SettingOpportunisticOnionUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOpportunisticOnionUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOpportunisticOnionUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOpportunisticOnionUpdateResponseJSON `json:"-"`
}

// settingOpportunisticOnionUpdateResponseJSON contains the JSON metadata for the
// struct [SettingOpportunisticOnionUpdateResponse]
type settingOpportunisticOnionUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOpportunisticOnionUpdateResponseID string

const (
	SettingOpportunisticOnionUpdateResponseIDOpportunisticOnion SettingOpportunisticOnionUpdateResponseID = "opportunistic_onion"
)

// Current value of the zone setting.
type SettingOpportunisticOnionUpdateResponseValue string

const (
	SettingOpportunisticOnionUpdateResponseValueOn  SettingOpportunisticOnionUpdateResponseValue = "on"
	SettingOpportunisticOnionUpdateResponseValueOff SettingOpportunisticOnionUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOpportunisticOnionUpdateResponseEditable bool

const (
	SettingOpportunisticOnionUpdateResponseEditableTrue  SettingOpportunisticOnionUpdateResponseEditable = true
	SettingOpportunisticOnionUpdateResponseEditableFalse SettingOpportunisticOnionUpdateResponseEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type SettingOpportunisticOnionGetResponse struct {
	// ID of the zone setting.
	ID SettingOpportunisticOnionGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOpportunisticOnionGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOpportunisticOnionGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOpportunisticOnionGetResponseJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseJSON contains the JSON metadata for the
// struct [SettingOpportunisticOnionGetResponse]
type settingOpportunisticOnionGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOpportunisticOnionGetResponseID string

const (
	SettingOpportunisticOnionGetResponseIDOpportunisticOnion SettingOpportunisticOnionGetResponseID = "opportunistic_onion"
)

// Current value of the zone setting.
type SettingOpportunisticOnionGetResponseValue string

const (
	SettingOpportunisticOnionGetResponseValueOn  SettingOpportunisticOnionGetResponseValue = "on"
	SettingOpportunisticOnionGetResponseValueOff SettingOpportunisticOnionGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOpportunisticOnionGetResponseEditable bool

const (
	SettingOpportunisticOnionGetResponseEditableTrue  SettingOpportunisticOnionGetResponseEditable = true
	SettingOpportunisticOnionGetResponseEditableFalse SettingOpportunisticOnionGetResponseEditable = false
)

type SettingOpportunisticOnionUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[SettingOpportunisticOnionUpdateParamsValue] `json:"value,required"`
}

func (r SettingOpportunisticOnionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type SettingOpportunisticOnionUpdateParamsValue string

const (
	SettingOpportunisticOnionUpdateParamsValueOn  SettingOpportunisticOnionUpdateParamsValue = "on"
	SettingOpportunisticOnionUpdateParamsValueOff SettingOpportunisticOnionUpdateParamsValue = "off"
)

type SettingOpportunisticOnionUpdateResponseEnvelope struct {
	Errors   []SettingOpportunisticOnionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticOnionUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result SettingOpportunisticOnionUpdateResponse             `json:"result"`
	JSON   settingOpportunisticOnionUpdateResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticOnionUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingOpportunisticOnionUpdateResponseEnvelope]
type settingOpportunisticOnionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticOnionUpdateResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingOpportunisticOnionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticOnionUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingOpportunisticOnionUpdateResponseEnvelopeErrors]
type settingOpportunisticOnionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticOnionUpdateResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    settingOpportunisticOnionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticOnionUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingOpportunisticOnionUpdateResponseEnvelopeMessages]
type settingOpportunisticOnionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticOnionGetResponseEnvelope struct {
	Errors   []SettingOpportunisticOnionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOpportunisticOnionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result SettingOpportunisticOnionGetResponse             `json:"result"`
	JSON   settingOpportunisticOnionGetResponseEnvelopeJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingOpportunisticOnionGetResponseEnvelope]
type settingOpportunisticOnionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticOnionGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingOpportunisticOnionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingOpportunisticOnionGetResponseEnvelopeErrors]
type settingOpportunisticOnionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOpportunisticOnionGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingOpportunisticOnionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOpportunisticOnionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingOpportunisticOnionGetResponseEnvelopeMessages]
type settingOpportunisticOnionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOpportunisticOnionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
