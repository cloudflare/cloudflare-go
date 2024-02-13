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

// SettingOrangeToOrangeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingOrangeToOrangeService]
// method instead.
type SettingOrangeToOrangeService struct {
	Options []option.RequestOption
}

// NewSettingOrangeToOrangeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingOrangeToOrangeService(opts ...option.RequestOption) (r *SettingOrangeToOrangeService) {
	r = &SettingOrangeToOrangeService{}
	r.Options = opts
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *SettingOrangeToOrangeService) Update(ctx context.Context, zoneID string, body SettingOrangeToOrangeUpdateParams, opts ...option.RequestOption) (res *SettingOrangeToOrangeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOrangeToOrangeUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *SettingOrangeToOrangeService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingOrangeToOrangeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingOrangeToOrangeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingOrangeToOrangeUpdateResponse struct {
	// ID of the zone setting.
	ID SettingOrangeToOrangeUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOrangeToOrangeUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOrangeToOrangeUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOrangeToOrangeUpdateResponseJSON `json:"-"`
}

// settingOrangeToOrangeUpdateResponseJSON contains the JSON metadata for the
// struct [SettingOrangeToOrangeUpdateResponse]
type settingOrangeToOrangeUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOrangeToOrangeUpdateResponseID string

const (
	SettingOrangeToOrangeUpdateResponseIDOrangeToOrange SettingOrangeToOrangeUpdateResponseID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingOrangeToOrangeUpdateResponseValue string

const (
	SettingOrangeToOrangeUpdateResponseValueOn  SettingOrangeToOrangeUpdateResponseValue = "on"
	SettingOrangeToOrangeUpdateResponseValueOff SettingOrangeToOrangeUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOrangeToOrangeUpdateResponseEditable bool

const (
	SettingOrangeToOrangeUpdateResponseEditableTrue  SettingOrangeToOrangeUpdateResponseEditable = true
	SettingOrangeToOrangeUpdateResponseEditableFalse SettingOrangeToOrangeUpdateResponseEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingOrangeToOrangeGetResponse struct {
	// ID of the zone setting.
	ID SettingOrangeToOrangeGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingOrangeToOrangeGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingOrangeToOrangeGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingOrangeToOrangeGetResponseJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseJSON contains the JSON metadata for the struct
// [SettingOrangeToOrangeGetResponse]
type settingOrangeToOrangeGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingOrangeToOrangeGetResponseID string

const (
	SettingOrangeToOrangeGetResponseIDOrangeToOrange SettingOrangeToOrangeGetResponseID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingOrangeToOrangeGetResponseValue string

const (
	SettingOrangeToOrangeGetResponseValueOn  SettingOrangeToOrangeGetResponseValue = "on"
	SettingOrangeToOrangeGetResponseValueOff SettingOrangeToOrangeGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOrangeToOrangeGetResponseEditable bool

const (
	SettingOrangeToOrangeGetResponseEditableTrue  SettingOrangeToOrangeGetResponseEditable = true
	SettingOrangeToOrangeGetResponseEditableFalse SettingOrangeToOrangeGetResponseEditable = false
)

type SettingOrangeToOrangeUpdateParams struct {
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Value param.Field[SettingOrangeToOrangeUpdateParamsValue] `json:"value,required"`
}

func (r SettingOrangeToOrangeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingOrangeToOrangeUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingOrangeToOrangeUpdateParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingOrangeToOrangeUpdateParamsValueValue] `json:"value,required"`
}

func (r SettingOrangeToOrangeUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingOrangeToOrangeUpdateParamsValueID string

const (
	SettingOrangeToOrangeUpdateParamsValueIDOrangeToOrange SettingOrangeToOrangeUpdateParamsValueID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingOrangeToOrangeUpdateParamsValueValue string

const (
	SettingOrangeToOrangeUpdateParamsValueValueOn  SettingOrangeToOrangeUpdateParamsValueValue = "on"
	SettingOrangeToOrangeUpdateParamsValueValueOff SettingOrangeToOrangeUpdateParamsValueValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingOrangeToOrangeUpdateParamsValueEditable bool

const (
	SettingOrangeToOrangeUpdateParamsValueEditableTrue  SettingOrangeToOrangeUpdateParamsValueEditable = true
	SettingOrangeToOrangeUpdateParamsValueEditableFalse SettingOrangeToOrangeUpdateParamsValueEditable = false
)

type SettingOrangeToOrangeUpdateResponseEnvelope struct {
	Errors   []SettingOrangeToOrangeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOrangeToOrangeUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result SettingOrangeToOrangeUpdateResponse             `json:"result"`
	JSON   settingOrangeToOrangeUpdateResponseEnvelopeJSON `json:"-"`
}

// settingOrangeToOrangeUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingOrangeToOrangeUpdateResponseEnvelope]
type settingOrangeToOrangeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeUpdateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingOrangeToOrangeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOrangeToOrangeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeUpdateResponseEnvelopeErrors]
type settingOrangeToOrangeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingOrangeToOrangeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOrangeToOrangeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingOrangeToOrangeUpdateResponseEnvelopeMessages]
type settingOrangeToOrangeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeGetResponseEnvelope struct {
	Errors   []SettingOrangeToOrangeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingOrangeToOrangeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result SettingOrangeToOrangeGetResponse             `json:"result"`
	JSON   settingOrangeToOrangeGetResponseEnvelopeJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOrangeToOrangeGetResponseEnvelope]
type settingOrangeToOrangeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingOrangeToOrangeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeGetResponseEnvelopeErrors]
type settingOrangeToOrangeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingOrangeToOrangeGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingOrangeToOrangeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingOrangeToOrangeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingOrangeToOrangeGetResponseEnvelopeMessages]
type settingOrangeToOrangeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOrangeToOrangeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
