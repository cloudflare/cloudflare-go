// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
func (r *SettingImageResizingService) Edit(ctx context.Context, params SettingImageResizingEditParams, opts ...option.RequestOption) (res *ImageResizing, err error) {
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
func (r *SettingImageResizingService) Get(ctx context.Context, query SettingImageResizingGetParams, opts ...option.RequestOption) (res *ImageResizing, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingImageResizingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/image_resizing", query.ZoneID)
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
type ImageResizing struct {
	// ID of the zone setting.
	ID ImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value ImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       imageResizingJSON `json:"-"`
}

// imageResizingJSON contains the JSON metadata for the struct [ImageResizing]
type imageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResizingJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ImageResizingID string

const (
	ImageResizingIDImageResizing ImageResizingID = "image_resizing"
)

func (r ImageResizingID) IsKnown() bool {
	switch r {
	case ImageResizingIDImageResizing:
		return true
	}
	return false
}

// Current value of the zone setting.
type ImageResizingValue string

const (
	ImageResizingValueOn   ImageResizingValue = "on"
	ImageResizingValueOff  ImageResizingValue = "off"
	ImageResizingValueOpen ImageResizingValue = "open"
)

func (r ImageResizingValue) IsKnown() bool {
	switch r {
	case ImageResizingValueOn, ImageResizingValueOff, ImageResizingValueOpen:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ImageResizingEditable bool

const (
	ImageResizingEditableTrue  ImageResizingEditable = true
	ImageResizingEditableFalse ImageResizingEditable = false
)

func (r ImageResizingEditable) IsKnown() bool {
	switch r {
	case ImageResizingEditableTrue, ImageResizingEditableFalse:
		return true
	}
	return false
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ImageResizingParam struct {
	// ID of the zone setting.
	ID param.Field[ImageResizingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ImageResizingValue] `json:"value,required"`
}

func (r ImageResizingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingImageResizingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[ImageResizingParam] `json:"value,required"`
}

func (r SettingImageResizingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingImageResizingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ImageResizing                                `json:"result"`
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

func (r settingImageResizingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingImageResizingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingImageResizingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ImageResizing                               `json:"result"`
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

func (r settingImageResizingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
