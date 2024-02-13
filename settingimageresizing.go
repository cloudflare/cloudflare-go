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

// SettingImageResizingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingImageResizingService]
// method instead.
type SettingImageResizingService struct {
	Options []option.RequestOption
}

// NewSettingImageResizingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingImageResizingService(opts ...option.RequestOption) (r *SettingImageResizingService) {
	r = &SettingImageResizingService{}
	r.Options = opts
	return
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
func (r *SettingImageResizingService) Update(ctx context.Context, zoneID string, body SettingImageResizingUpdateParams, opts ...option.RequestOption) (res *SettingImageResizingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingImageResizingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/image_resizing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
func (r *SettingImageResizingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingImageResizingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingImageResizingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/image_resizing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingImageResizingUpdateResponse struct {
	// ID of the zone setting.
	ID SettingImageResizingUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingImageResizingUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingImageResizingUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingImageResizingUpdateResponseJSON `json:"-"`
}

// settingImageResizingUpdateResponseJSON contains the JSON metadata for the struct
// [SettingImageResizingUpdateResponse]
type settingImageResizingUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingImageResizingUpdateResponseID string

const (
	SettingImageResizingUpdateResponseIDImageResizing SettingImageResizingUpdateResponseID = "image_resizing"
)

// Current value of the zone setting.
type SettingImageResizingUpdateResponseValue string

const (
	SettingImageResizingUpdateResponseValueOn   SettingImageResizingUpdateResponseValue = "on"
	SettingImageResizingUpdateResponseValueOff  SettingImageResizingUpdateResponseValue = "off"
	SettingImageResizingUpdateResponseValueOpen SettingImageResizingUpdateResponseValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingImageResizingUpdateResponseEditable bool

const (
	SettingImageResizingUpdateResponseEditableTrue  SettingImageResizingUpdateResponseEditable = true
	SettingImageResizingUpdateResponseEditableFalse SettingImageResizingUpdateResponseEditable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingImageResizingGetResponse struct {
	// ID of the zone setting.
	ID SettingImageResizingGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingImageResizingGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingImageResizingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingImageResizingGetResponseJSON `json:"-"`
}

// settingImageResizingGetResponseJSON contains the JSON metadata for the struct
// [SettingImageResizingGetResponse]
type settingImageResizingGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingImageResizingGetResponseID string

const (
	SettingImageResizingGetResponseIDImageResizing SettingImageResizingGetResponseID = "image_resizing"
)

// Current value of the zone setting.
type SettingImageResizingGetResponseValue string

const (
	SettingImageResizingGetResponseValueOn   SettingImageResizingGetResponseValue = "on"
	SettingImageResizingGetResponseValueOff  SettingImageResizingGetResponseValue = "off"
	SettingImageResizingGetResponseValueOpen SettingImageResizingGetResponseValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingImageResizingGetResponseEditable bool

const (
	SettingImageResizingGetResponseEditableTrue  SettingImageResizingGetResponseEditable = true
	SettingImageResizingGetResponseEditableFalse SettingImageResizingGetResponseEditable = false
)

type SettingImageResizingUpdateParams struct {
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[SettingImageResizingUpdateParamsValue] `json:"value,required"`
}

func (r SettingImageResizingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingImageResizingUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingImageResizingUpdateParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingImageResizingUpdateParamsValueValue] `json:"value,required"`
}

func (r SettingImageResizingUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingImageResizingUpdateParamsValueID string

const (
	SettingImageResizingUpdateParamsValueIDImageResizing SettingImageResizingUpdateParamsValueID = "image_resizing"
)

// Current value of the zone setting.
type SettingImageResizingUpdateParamsValueValue string

const (
	SettingImageResizingUpdateParamsValueValueOn   SettingImageResizingUpdateParamsValueValue = "on"
	SettingImageResizingUpdateParamsValueValueOff  SettingImageResizingUpdateParamsValueValue = "off"
	SettingImageResizingUpdateParamsValueValueOpen SettingImageResizingUpdateParamsValueValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingImageResizingUpdateParamsValueEditable bool

const (
	SettingImageResizingUpdateParamsValueEditableTrue  SettingImageResizingUpdateParamsValueEditable = true
	SettingImageResizingUpdateParamsValueEditableFalse SettingImageResizingUpdateParamsValueEditable = false
)

type SettingImageResizingUpdateResponseEnvelope struct {
	Errors   []SettingImageResizingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingImageResizingUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result SettingImageResizingUpdateResponse             `json:"result"`
	JSON   settingImageResizingUpdateResponseEnvelopeJSON `json:"-"`
}

// settingImageResizingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingImageResizingUpdateResponseEnvelope]
type settingImageResizingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingImageResizingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingImageResizingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingImageResizingUpdateResponseEnvelopeErrors]
type settingImageResizingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingImageResizingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingImageResizingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingImageResizingUpdateResponseEnvelopeMessages]
type settingImageResizingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingGetResponseEnvelope struct {
	Errors   []SettingImageResizingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingImageResizingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result SettingImageResizingGetResponse             `json:"result"`
	JSON   settingImageResizingGetResponseEnvelopeJSON `json:"-"`
}

// settingImageResizingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingImageResizingGetResponseEnvelope]
type settingImageResizingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingImageResizingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingImageResizingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingImageResizingGetResponseEnvelopeErrors]
type settingImageResizingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingImageResizingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingImageResizingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingImageResizingGetResponseEnvelopeMessages]
type settingImageResizingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
