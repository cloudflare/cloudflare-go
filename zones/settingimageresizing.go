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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *SettingImageResizingService) Edit(ctx context.Context, params SettingImageResizingEditParams, opts ...option.RequestOption) (res *ZoneSettingImageResizing, err error) {
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
func (r *SettingImageResizingService) Get(ctx context.Context, query SettingImageResizingGetParams, opts ...option.RequestOption) (res *ZoneSettingImageResizing, err error) {
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
type ZoneSettingImageResizing struct {
	// ID of the zone setting.
	ID ZoneSettingImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingImageResizingJSON `json:"-"`
}

// zoneSettingImageResizingJSON contains the JSON metadata for the struct
// [ZoneSettingImageResizing]
type zoneSettingImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingImageResizingJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingImageResizing) implementsZonesSettingEditResponse() {}

func (r ZoneSettingImageResizing) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingImageResizingID string

const (
	ZoneSettingImageResizingIDImageResizing ZoneSettingImageResizingID = "image_resizing"
)

func (r ZoneSettingImageResizingID) IsKnown() bool {
	switch r {
	case ZoneSettingImageResizingIDImageResizing:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingImageResizingValue string

const (
	ZoneSettingImageResizingValueOn   ZoneSettingImageResizingValue = "on"
	ZoneSettingImageResizingValueOff  ZoneSettingImageResizingValue = "off"
	ZoneSettingImageResizingValueOpen ZoneSettingImageResizingValue = "open"
)

func (r ZoneSettingImageResizingValue) IsKnown() bool {
	switch r {
	case ZoneSettingImageResizingValueOn, ZoneSettingImageResizingValueOff, ZoneSettingImageResizingValueOpen:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingImageResizingEditable bool

const (
	ZoneSettingImageResizingEditableTrue  ZoneSettingImageResizingEditable = true
	ZoneSettingImageResizingEditableFalse ZoneSettingImageResizingEditable = false
)

func (r ZoneSettingImageResizingEditable) IsKnown() bool {
	switch r {
	case ZoneSettingImageResizingEditableTrue, ZoneSettingImageResizingEditableFalse:
		return true
	}
	return false
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingImageResizingParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingImageResizingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingImageResizingValue] `json:"value,required"`
}

func (r ZoneSettingImageResizingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingImageResizingParam) implementsZonesSettingEditParamsItem() {}

type SettingImageResizingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[ZoneSettingImageResizingParam] `json:"value,required"`
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
	Result ZoneSettingImageResizing                     `json:"result"`
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
	Result ZoneSettingImageResizing                    `json:"result"`
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
