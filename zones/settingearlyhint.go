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
func (r *SettingEarlyHintService) Edit(ctx context.Context, params SettingEarlyHintEditParams, opts ...option.RequestOption) (res *SettingEarlyHintEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEarlyHintEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
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
func (r *SettingEarlyHintService) Get(ctx context.Context, query SettingEarlyHintGetParams, opts ...option.RequestOption) (res *SettingEarlyHintGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEarlyHintGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/early_hints", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type SettingEarlyHintEditResponse struct {
	// ID of the zone setting.
	ID SettingEarlyHintEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEarlyHintEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEarlyHintEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEarlyHintEditResponseJSON `json:"-"`
}

// settingEarlyHintEditResponseJSON contains the JSON metadata for the struct
// [SettingEarlyHintEditResponse]
type settingEarlyHintEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SettingEarlyHintEditResponseID string

const (
	SettingEarlyHintEditResponseIDEarlyHints SettingEarlyHintEditResponseID = "early_hints"
)

// Current value of the zone setting.
type SettingEarlyHintEditResponseValue string

const (
	SettingEarlyHintEditResponseValueOn  SettingEarlyHintEditResponseValue = "on"
	SettingEarlyHintEditResponseValueOff SettingEarlyHintEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEarlyHintEditResponseEditable bool

const (
	SettingEarlyHintEditResponseEditableTrue  SettingEarlyHintEditResponseEditable = true
	SettingEarlyHintEditResponseEditableFalse SettingEarlyHintEditResponseEditable = false
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

func (r settingEarlyHintGetResponseJSON) RawJSON() string {
	return r.raw
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

type SettingEarlyHintEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingEarlyHintEditParamsValue] `json:"value,required"`
}

func (r SettingEarlyHintEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingEarlyHintEditParamsValue string

const (
	SettingEarlyHintEditParamsValueOn  SettingEarlyHintEditParamsValue = "on"
	SettingEarlyHintEditParamsValueOff SettingEarlyHintEditParamsValue = "off"
)

type SettingEarlyHintEditResponseEnvelope struct {
	Errors   []SettingEarlyHintEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEarlyHintEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result SettingEarlyHintEditResponse             `json:"result"`
	JSON   settingEarlyHintEditResponseEnvelopeJSON `json:"-"`
}

// settingEarlyHintEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingEarlyHintEditResponseEnvelope]
type settingEarlyHintEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingEarlyHintEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingEarlyHintEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEarlyHintEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingEarlyHintEditResponseEnvelopeErrors]
type settingEarlyHintEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingEarlyHintEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingEarlyHintEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEarlyHintEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingEarlyHintEditResponseEnvelopeMessages]
type settingEarlyHintEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEarlyHintEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEarlyHintEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingEarlyHintGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

func (r settingEarlyHintGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r settingEarlyHintGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingEarlyHintGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
