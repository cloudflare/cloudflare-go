// File generated from our OpenAPI spec by Stainless.

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
func (r *SettingImageResizingService) Edit(ctx context.Context, params SettingImageResizingEditParams, opts ...option.RequestOption) (res *ZonesImageResizing, err error) {
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
func (r *SettingImageResizingService) Get(ctx context.Context, query SettingImageResizingGetParams, opts ...option.RequestOption) (res *ZonesImageResizing, err error) {
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
type ZonesImageResizing struct {
	// ID of the zone setting.
	ID ZonesImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesImageResizingJSON `json:"-"`
}

// zonesImageResizingJSON contains the JSON metadata for the struct
// [ZonesImageResizing]
type zonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesImageResizingJSON) RawJSON() string {
	return r.raw
}

func (r ZonesImageResizing) implementsZonesSettingEditResponse() {}

func (r ZonesImageResizing) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesImageResizingID string

const (
	ZonesImageResizingIDImageResizing ZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type ZonesImageResizingValue string

const (
	ZonesImageResizingValueOn   ZonesImageResizingValue = "on"
	ZonesImageResizingValueOff  ZonesImageResizingValue = "off"
	ZonesImageResizingValueOpen ZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesImageResizingEditable bool

const (
	ZonesImageResizingEditableTrue  ZonesImageResizingEditable = true
	ZonesImageResizingEditableFalse ZonesImageResizingEditable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZonesImageResizingParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesImageResizingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesImageResizingValue] `json:"value,required"`
}

func (r ZonesImageResizingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesImageResizingParam) implementsZonesSettingEditParamsItem() {}

type SettingImageResizingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[ZonesImageResizingParam] `json:"value,required"`
}

func (r SettingImageResizingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingImageResizingEditResponseEnvelope struct {
	Errors   []SettingImageResizingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingImageResizingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ZonesImageResizing                           `json:"result"`
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

func (r settingImageResizingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingImageResizingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZonesImageResizing                          `json:"result"`
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

func (r settingImageResizingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingImageResizingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
