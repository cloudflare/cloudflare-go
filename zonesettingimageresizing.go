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

// ZoneSettingImageResizingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingImageResizingService] method instead.
type ZoneSettingImageResizingService struct {
	Options []option.RequestOption
}

// NewZoneSettingImageResizingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingImageResizingService(opts ...option.RequestOption) (r *ZoneSettingImageResizingService) {
	r = &ZoneSettingImageResizingService{}
	r.Options = opts
	return
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
func (r *ZoneSettingImageResizingService) Edit(ctx context.Context, params ZoneSettingImageResizingEditParams, opts ...option.RequestOption) (res *ZonesImageResizing, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingImageResizingEditResponseEnvelope
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
func (r *ZoneSettingImageResizingService) Get(ctx context.Context, query ZoneSettingImageResizingGetParams, opts ...option.RequestOption) (res *ZonesImageResizing, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingImageResizingGetResponseEnvelope
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

func (r ZonesImageResizing) implementsZoneSettingEditResponse() {}

func (r ZonesImageResizing) implementsZoneSettingGetResponse() {}

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

func (r ZonesImageResizingParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingImageResizingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[ZonesImageResizingParam] `json:"value,required"`
}

func (r ZoneSettingImageResizingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingImageResizingEditResponseEnvelope struct {
	Errors   []ZoneSettingImageResizingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingImageResizingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ZonesImageResizing                               `json:"result"`
	JSON   zoneSettingImageResizingEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingImageResizingEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingImageResizingEditResponseEnvelope]
type zoneSettingImageResizingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingImageResizingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingImageResizingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingImageResizingEditResponseEnvelopeErrors]
type zoneSettingImageResizingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingImageResizingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingImageResizingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingImageResizingEditResponseEnvelopeMessages]
type zoneSettingImageResizingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingImageResizingGetResponseEnvelope struct {
	Errors   []ZoneSettingImageResizingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingImageResizingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ZonesImageResizing                              `json:"result"`
	JSON   zoneSettingImageResizingGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingImageResizingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingImageResizingGetResponseEnvelope]
type zoneSettingImageResizingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingImageResizingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingImageResizingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingImageResizingGetResponseEnvelopeErrors]
type zoneSettingImageResizingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingImageResizingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingImageResizingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingImageResizingGetResponseEnvelopeMessages]
type zoneSettingImageResizingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
