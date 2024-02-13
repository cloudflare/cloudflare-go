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

// SettingWebpService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingWebpService] method
// instead.
type SettingWebpService struct {
	Options []option.RequestOption
}

// NewSettingWebpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingWebpService(opts ...option.RequestOption) (r *SettingWebpService) {
	r = &SettingWebpService{}
	r.Options = opts
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *SettingWebpService) Update(ctx context.Context, zoneID string, body SettingWebpUpdateParams, opts ...option.RequestOption) (res *SettingWebpUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebpUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/webp", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *SettingWebpService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingWebpGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebpGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/webp", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type SettingWebpUpdateResponse struct {
	// ID of the zone setting.
	ID SettingWebpUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingWebpUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingWebpUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingWebpUpdateResponseJSON `json:"-"`
}

// settingWebpUpdateResponseJSON contains the JSON metadata for the struct
// [SettingWebpUpdateResponse]
type settingWebpUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingWebpUpdateResponseID string

const (
	SettingWebpUpdateResponseIDWebp SettingWebpUpdateResponseID = "webp"
)

// Current value of the zone setting.
type SettingWebpUpdateResponseValue string

const (
	SettingWebpUpdateResponseValueOff SettingWebpUpdateResponseValue = "off"
	SettingWebpUpdateResponseValueOn  SettingWebpUpdateResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingWebpUpdateResponseEditable bool

const (
	SettingWebpUpdateResponseEditableTrue  SettingWebpUpdateResponseEditable = true
	SettingWebpUpdateResponseEditableFalse SettingWebpUpdateResponseEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type SettingWebpGetResponse struct {
	// ID of the zone setting.
	ID SettingWebpGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingWebpGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingWebpGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingWebpGetResponseJSON `json:"-"`
}

// settingWebpGetResponseJSON contains the JSON metadata for the struct
// [SettingWebpGetResponse]
type settingWebpGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingWebpGetResponseID string

const (
	SettingWebpGetResponseIDWebp SettingWebpGetResponseID = "webp"
)

// Current value of the zone setting.
type SettingWebpGetResponseValue string

const (
	SettingWebpGetResponseValueOff SettingWebpGetResponseValue = "off"
	SettingWebpGetResponseValueOn  SettingWebpGetResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingWebpGetResponseEditable bool

const (
	SettingWebpGetResponseEditableTrue  SettingWebpGetResponseEditable = true
	SettingWebpGetResponseEditableFalse SettingWebpGetResponseEditable = false
)

type SettingWebpUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingWebpUpdateParamsValue] `json:"value,required"`
}

func (r SettingWebpUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingWebpUpdateParamsValue string

const (
	SettingWebpUpdateParamsValueOff SettingWebpUpdateParamsValue = "off"
	SettingWebpUpdateParamsValueOn  SettingWebpUpdateParamsValue = "on"
)

type SettingWebpUpdateResponseEnvelope struct {
	Errors   []SettingWebpUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebpUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result SettingWebpUpdateResponse             `json:"result"`
	JSON   settingWebpUpdateResponseEnvelopeJSON `json:"-"`
}

// settingWebpUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWebpUpdateResponseEnvelope]
type settingWebpUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingWebpUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingWebpUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebpUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWebpUpdateResponseEnvelopeErrors]
type settingWebpUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingWebpUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingWebpUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebpUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWebpUpdateResponseEnvelopeMessages]
type settingWebpUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingWebpGetResponseEnvelope struct {
	Errors   []SettingWebpGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebpGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result SettingWebpGetResponse             `json:"result"`
	JSON   settingWebpGetResponseEnvelopeJSON `json:"-"`
}

// settingWebpGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWebpGetResponseEnvelope]
type settingWebpGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingWebpGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingWebpGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebpGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWebpGetResponseEnvelopeErrors]
type settingWebpGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingWebpGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingWebpGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebpGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWebpGetResponseEnvelopeMessages]
type settingWebpGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
