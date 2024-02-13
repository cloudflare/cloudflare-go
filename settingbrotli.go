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

// SettingBrotliService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingBrotliService] method
// instead.
type SettingBrotliService struct {
	Options []option.RequestOption
}

// NewSettingBrotliService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingBrotliService(opts ...option.RequestOption) (r *SettingBrotliService) {
	r = &SettingBrotliService{}
	r.Options = opts
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *SettingBrotliService) Update(ctx context.Context, zoneID string, body SettingBrotliUpdateParams, opts ...option.RequestOption) (res *SettingBrotliUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrotliUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/brotli", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *SettingBrotliService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingBrotliGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrotliGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/brotli", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type SettingBrotliUpdateResponse struct {
	// ID of the zone setting.
	ID SettingBrotliUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingBrotliUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingBrotliUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingBrotliUpdateResponseJSON `json:"-"`
}

// settingBrotliUpdateResponseJSON contains the JSON metadata for the struct
// [SettingBrotliUpdateResponse]
type settingBrotliUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingBrotliUpdateResponseID string

const (
	SettingBrotliUpdateResponseIDBrotli SettingBrotliUpdateResponseID = "brotli"
)

// Current value of the zone setting.
type SettingBrotliUpdateResponseValue string

const (
	SettingBrotliUpdateResponseValueOff SettingBrotliUpdateResponseValue = "off"
	SettingBrotliUpdateResponseValueOn  SettingBrotliUpdateResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingBrotliUpdateResponseEditable bool

const (
	SettingBrotliUpdateResponseEditableTrue  SettingBrotliUpdateResponseEditable = true
	SettingBrotliUpdateResponseEditableFalse SettingBrotliUpdateResponseEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type SettingBrotliGetResponse struct {
	// ID of the zone setting.
	ID SettingBrotliGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingBrotliGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingBrotliGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingBrotliGetResponseJSON `json:"-"`
}

// settingBrotliGetResponseJSON contains the JSON metadata for the struct
// [SettingBrotliGetResponse]
type settingBrotliGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingBrotliGetResponseID string

const (
	SettingBrotliGetResponseIDBrotli SettingBrotliGetResponseID = "brotli"
)

// Current value of the zone setting.
type SettingBrotliGetResponseValue string

const (
	SettingBrotliGetResponseValueOff SettingBrotliGetResponseValue = "off"
	SettingBrotliGetResponseValueOn  SettingBrotliGetResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingBrotliGetResponseEditable bool

const (
	SettingBrotliGetResponseEditableTrue  SettingBrotliGetResponseEditable = true
	SettingBrotliGetResponseEditableFalse SettingBrotliGetResponseEditable = false
)

type SettingBrotliUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingBrotliUpdateParamsValue] `json:"value,required"`
}

func (r SettingBrotliUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingBrotliUpdateParamsValue string

const (
	SettingBrotliUpdateParamsValueOff SettingBrotliUpdateParamsValue = "off"
	SettingBrotliUpdateParamsValueOn  SettingBrotliUpdateParamsValue = "on"
)

type SettingBrotliUpdateResponseEnvelope struct {
	Errors   []SettingBrotliUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrotliUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result SettingBrotliUpdateResponse             `json:"result"`
	JSON   settingBrotliUpdateResponseEnvelopeJSON `json:"-"`
}

// settingBrotliUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBrotliUpdateResponseEnvelope]
type settingBrotliUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingBrotliUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingBrotliUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrotliUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingBrotliUpdateResponseEnvelopeErrors]
type settingBrotliUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingBrotliUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingBrotliUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrotliUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingBrotliUpdateResponseEnvelopeMessages]
type settingBrotliUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingBrotliGetResponseEnvelope struct {
	Errors   []SettingBrotliGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrotliGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result SettingBrotliGetResponse             `json:"result"`
	JSON   settingBrotliGetResponseEnvelopeJSON `json:"-"`
}

// settingBrotliGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingBrotliGetResponseEnvelope]
type settingBrotliGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingBrotliGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingBrotliGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrotliGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingBrotliGetResponseEnvelopeErrors]
type settingBrotliGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingBrotliGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingBrotliGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrotliGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingBrotliGetResponseEnvelopeMessages]
type settingBrotliGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrotliGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
