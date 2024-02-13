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

// SettingEarlyHintService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingEarlyHintService] method
// instead.
type SettingEarlyHintService struct {
	Options []option.RequestOption
}

// NewSettingEarlyHintService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingEarlyHintService(opts ...option.RequestOption) (r *SettingEarlyHintService) {
	r = &SettingEarlyHintService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *SettingEarlyHintService) Update(ctx context.Context, zoneID string, body SettingEarlyHintUpdateParams, opts ...option.RequestOption) (res *SettingEarlyHintUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEarlyHintUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *SettingEarlyHintService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingEarlyHintGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEarlyHintGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type SettingEarlyHintUpdateResponse struct {
	// ID of the zone setting.
	ID SettingEarlyHintUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEarlyHintUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEarlyHintUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEarlyHintUpdateResponseJSON `json:"-"`
}

// settingEarlyHintUpdateResponseJSON contains the JSON metadata for the struct
// [SettingEarlyHintUpdateResponse]
type settingEarlyHintUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingEarlyHintUpdateResponseID string

const (
	SettingEarlyHintUpdateResponseIDEarlyHints SettingEarlyHintUpdateResponseID = "early_hints"
)

// Current value of the zone setting.
type SettingEarlyHintUpdateResponseValue string

const (
	SettingEarlyHintUpdateResponseValueOn  SettingEarlyHintUpdateResponseValue = "on"
	SettingEarlyHintUpdateResponseValueOff SettingEarlyHintUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEarlyHintUpdateResponseEditable bool

const (
	SettingEarlyHintUpdateResponseEditableTrue  SettingEarlyHintUpdateResponseEditable = true
	SettingEarlyHintUpdateResponseEditableFalse SettingEarlyHintUpdateResponseEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type SettingEarlyHintGetResponse struct {
	// ID of the zone setting.
	ID SettingEarlyHintGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEarlyHintGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEarlyHintGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEarlyHintGetResponseJSON `json:"-"`
}

// settingEarlyHintGetResponseJSON contains the JSON metadata for the struct
// [SettingEarlyHintGetResponse]
type settingEarlyHintGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingEarlyHintGetResponseID string

const (
	SettingEarlyHintGetResponseIDEarlyHints SettingEarlyHintGetResponseID = "early_hints"
)

// Current value of the zone setting.
type SettingEarlyHintGetResponseValue string

const (
	SettingEarlyHintGetResponseValueOn  SettingEarlyHintGetResponseValue = "on"
	SettingEarlyHintGetResponseValueOff SettingEarlyHintGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEarlyHintGetResponseEditable bool

const (
	SettingEarlyHintGetResponseEditableTrue  SettingEarlyHintGetResponseEditable = true
	SettingEarlyHintGetResponseEditableFalse SettingEarlyHintGetResponseEditable = false
)

type SettingEarlyHintUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingEarlyHintUpdateParamsValue] `json:"value,required"`
}

func (r SettingEarlyHintUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingEarlyHintUpdateParamsValue string

const (
	SettingEarlyHintUpdateParamsValueOn  SettingEarlyHintUpdateParamsValue = "on"
	SettingEarlyHintUpdateParamsValueOff SettingEarlyHintUpdateParamsValue = "off"
)

type SettingEarlyHintUpdateResponseEnvelope struct {
	Errors   []SettingEarlyHintUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEarlyHintUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result SettingEarlyHintUpdateResponse             `json:"result"`
	JSON   settingEarlyHintUpdateResponseEnvelopeJSON `json:"-"`
}

// settingEarlyHintUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingEarlyHintUpdateResponseEnvelope]
type settingEarlyHintUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEarlyHintUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingEarlyHintUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEarlyHintUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingEarlyHintUpdateResponseEnvelopeErrors]
type settingEarlyHintUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEarlyHintUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingEarlyHintUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEarlyHintUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingEarlyHintUpdateResponseEnvelopeMessages]
type settingEarlyHintUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEarlyHintGetResponseEnvelope struct {
	Errors   []SettingEarlyHintGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEarlyHintGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result SettingEarlyHintGetResponse             `json:"result"`
	JSON   settingEarlyHintGetResponseEnvelopeJSON `json:"-"`
}

// settingEarlyHintGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingEarlyHintGetResponseEnvelope]
type settingEarlyHintGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEarlyHintGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingEarlyHintGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEarlyHintGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingEarlyHintGetResponseEnvelopeErrors]
type settingEarlyHintGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEarlyHintGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingEarlyHintGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEarlyHintGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingEarlyHintGetResponseEnvelopeMessages]
type settingEarlyHintGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
