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
func (r *SettingImageResizingService) Edit(ctx context.Context, params SettingImageResizingEditParams, opts ...option.RequestOption) (res *SettingImageResizingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingImageResizingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/image_resizing", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
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
func (r *SettingImageResizingService) Get(ctx context.Context, query SettingImageResizingGetParams, opts ...option.RequestOption) (res *SettingImageResizingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingImageResizingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/image_resizing", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type SettingImageResizingEditResponse struct {
	// ID of the zone setting.
	ID SettingImageResizingEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingImageResizingEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingImageResizingEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingImageResizingEditResponseJSON `json:"-"`
}

// settingImageResizingEditResponseJSON contains the JSON metadata for the struct
// [SettingImageResizingEditResponse]
type settingImageResizingEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingImageResizingEditResponseID string

const (
	SettingImageResizingEditResponseIDImageResizing SettingImageResizingEditResponseID = "image_resizing"
)

// Current value of the zone setting.
type SettingImageResizingEditResponseValue string

const (
	SettingImageResizingEditResponseValueOn   SettingImageResizingEditResponseValue = "on"
	SettingImageResizingEditResponseValueOff  SettingImageResizingEditResponseValue = "off"
	SettingImageResizingEditResponseValueOpen SettingImageResizingEditResponseValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingImageResizingEditResponseEditable bool

const (
	SettingImageResizingEditResponseEditableTrue  SettingImageResizingEditResponseEditable = true
	SettingImageResizingEditResponseEditableFalse SettingImageResizingEditResponseEditable = false
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

type SettingImageResizingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[SettingImageResizingEditParamsValue] `json:"value,required"`
}

func (r SettingImageResizingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingImageResizingEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingImageResizingEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingImageResizingEditParamsValueValue] `json:"value,required"`
}

func (r SettingImageResizingEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingImageResizingEditParamsValueID string

const (
	SettingImageResizingEditParamsValueIDImageResizing SettingImageResizingEditParamsValueID = "image_resizing"
)

// Current value of the zone setting.
type SettingImageResizingEditParamsValueValue string

const (
	SettingImageResizingEditParamsValueValueOn   SettingImageResizingEditParamsValueValue = "on"
	SettingImageResizingEditParamsValueValueOff  SettingImageResizingEditParamsValueValue = "off"
	SettingImageResizingEditParamsValueValueOpen SettingImageResizingEditParamsValueValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingImageResizingEditParamsValueEditable bool

const (
	SettingImageResizingEditParamsValueEditableTrue  SettingImageResizingEditParamsValueEditable = true
	SettingImageResizingEditParamsValueEditableFalse SettingImageResizingEditParamsValueEditable = false
)

type SettingImageResizingEditResponseEnvelope struct {
	Errors   []SettingImageResizingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingImageResizingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result SettingImageResizingEditResponse             `json:"result"`
	JSON   settingImageResizingEditResponseEnvelopeJSON `json:"-"`
}

// settingImageResizingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingImageResizingEditResponseEnvelope]
type settingImageResizingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingImageResizingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingImageResizingEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingImageResizingEditResponseEnvelopeErrors]
type settingImageResizingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingImageResizingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingImageResizingEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingImageResizingEditResponseEnvelopeMessages]
type settingImageResizingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingImageResizingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingImageResizingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
