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

// SettingH2PrioritizationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingH2PrioritizationService] method instead.
type SettingH2PrioritizationService struct {
	Options []option.RequestOption
}

// NewSettingH2PrioritizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingH2PrioritizationService(opts ...option.RequestOption) (r *SettingH2PrioritizationService) {
	r = &SettingH2PrioritizationService{}
	r.Options = opts
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *SettingH2PrioritizationService) Update(ctx context.Context, zoneID string, body SettingH2PrioritizationUpdateParams, opts ...option.RequestOption) (res *SettingH2PrioritizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingH2PrioritizationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *SettingH2PrioritizationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingH2PrioritizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingH2PrioritizationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingH2PrioritizationUpdateResponse struct {
	// ID of the zone setting.
	ID SettingH2PrioritizationUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingH2PrioritizationUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingH2PrioritizationUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingH2PrioritizationUpdateResponseJSON `json:"-"`
}

// settingH2PrioritizationUpdateResponseJSON contains the JSON metadata for the
// struct [SettingH2PrioritizationUpdateResponse]
type settingH2PrioritizationUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingH2PrioritizationUpdateResponseID string

const (
	SettingH2PrioritizationUpdateResponseIDH2Prioritization SettingH2PrioritizationUpdateResponseID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingH2PrioritizationUpdateResponseValue string

const (
	SettingH2PrioritizationUpdateResponseValueOn     SettingH2PrioritizationUpdateResponseValue = "on"
	SettingH2PrioritizationUpdateResponseValueOff    SettingH2PrioritizationUpdateResponseValue = "off"
	SettingH2PrioritizationUpdateResponseValueCustom SettingH2PrioritizationUpdateResponseValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingH2PrioritizationUpdateResponseEditable bool

const (
	SettingH2PrioritizationUpdateResponseEditableTrue  SettingH2PrioritizationUpdateResponseEditable = true
	SettingH2PrioritizationUpdateResponseEditableFalse SettingH2PrioritizationUpdateResponseEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingH2PrioritizationGetResponse struct {
	// ID of the zone setting.
	ID SettingH2PrioritizationGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingH2PrioritizationGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingH2PrioritizationGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingH2PrioritizationGetResponseJSON `json:"-"`
}

// settingH2PrioritizationGetResponseJSON contains the JSON metadata for the struct
// [SettingH2PrioritizationGetResponse]
type settingH2PrioritizationGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingH2PrioritizationGetResponseID string

const (
	SettingH2PrioritizationGetResponseIDH2Prioritization SettingH2PrioritizationGetResponseID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingH2PrioritizationGetResponseValue string

const (
	SettingH2PrioritizationGetResponseValueOn     SettingH2PrioritizationGetResponseValue = "on"
	SettingH2PrioritizationGetResponseValueOff    SettingH2PrioritizationGetResponseValue = "off"
	SettingH2PrioritizationGetResponseValueCustom SettingH2PrioritizationGetResponseValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingH2PrioritizationGetResponseEditable bool

const (
	SettingH2PrioritizationGetResponseEditableTrue  SettingH2PrioritizationGetResponseEditable = true
	SettingH2PrioritizationGetResponseEditableFalse SettingH2PrioritizationGetResponseEditable = false
)

type SettingH2PrioritizationUpdateParams struct {
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value param.Field[SettingH2PrioritizationUpdateParamsValue] `json:"value,required"`
}

func (r SettingH2PrioritizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingH2PrioritizationUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingH2PrioritizationUpdateParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingH2PrioritizationUpdateParamsValueValue] `json:"value,required"`
}

func (r SettingH2PrioritizationUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingH2PrioritizationUpdateParamsValueID string

const (
	SettingH2PrioritizationUpdateParamsValueIDH2Prioritization SettingH2PrioritizationUpdateParamsValueID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingH2PrioritizationUpdateParamsValueValue string

const (
	SettingH2PrioritizationUpdateParamsValueValueOn     SettingH2PrioritizationUpdateParamsValueValue = "on"
	SettingH2PrioritizationUpdateParamsValueValueOff    SettingH2PrioritizationUpdateParamsValueValue = "off"
	SettingH2PrioritizationUpdateParamsValueValueCustom SettingH2PrioritizationUpdateParamsValueValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingH2PrioritizationUpdateParamsValueEditable bool

const (
	SettingH2PrioritizationUpdateParamsValueEditableTrue  SettingH2PrioritizationUpdateParamsValueEditable = true
	SettingH2PrioritizationUpdateParamsValueEditableFalse SettingH2PrioritizationUpdateParamsValueEditable = false
)

type SettingH2PrioritizationUpdateResponseEnvelope struct {
	Errors   []SettingH2PrioritizationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingH2PrioritizationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result SettingH2PrioritizationUpdateResponse             `json:"result"`
	JSON   settingH2PrioritizationUpdateResponseEnvelopeJSON `json:"-"`
}

// settingH2PrioritizationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingH2PrioritizationUpdateResponseEnvelope]
type settingH2PrioritizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingH2PrioritizationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingH2PrioritizationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingH2PrioritizationUpdateResponseEnvelopeErrors]
type settingH2PrioritizationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    settingH2PrioritizationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingH2PrioritizationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingH2PrioritizationUpdateResponseEnvelopeMessages]
type settingH2PrioritizationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationGetResponseEnvelope struct {
	Errors   []SettingH2PrioritizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingH2PrioritizationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result SettingH2PrioritizationGetResponse             `json:"result"`
	JSON   settingH2PrioritizationGetResponseEnvelopeJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingH2PrioritizationGetResponseEnvelope]
type settingH2PrioritizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingH2PrioritizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingH2PrioritizationGetResponseEnvelopeErrors]
type settingH2PrioritizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingH2PrioritizationGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingH2PrioritizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingH2PrioritizationGetResponseEnvelopeMessages]
type settingH2PrioritizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
