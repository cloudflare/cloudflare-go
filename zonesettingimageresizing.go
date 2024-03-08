// File generated from our OpenAPI spec by Stainless.

package cloudflare

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
func (r *ZoneSettingImageResizingService) Edit(ctx context.Context, params ZoneSettingImageResizingEditParams, opts ...option.RequestOption) (res *ZoneSettingImageResizingEditResponse, err error) {
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
func (r *ZoneSettingImageResizingService) Get(ctx context.Context, query ZoneSettingImageResizingGetParams, opts ...option.RequestOption) (res *ZoneSettingImageResizingGetResponse, err error) {
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
type ZoneSettingImageResizingEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingImageResizingEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingImageResizingEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingImageResizingEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingImageResizingEditResponseJSON `json:"-"`
}

// zoneSettingImageResizingEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingImageResizingEditResponse]
type zoneSettingImageResizingEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingImageResizingEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingImageResizingEditResponseID string

const (
	ZoneSettingImageResizingEditResponseIDImageResizing ZoneSettingImageResizingEditResponseID = "image_resizing"
)

// Current value of the zone setting.
type ZoneSettingImageResizingEditResponseValue string

const (
	ZoneSettingImageResizingEditResponseValueOn   ZoneSettingImageResizingEditResponseValue = "on"
	ZoneSettingImageResizingEditResponseValueOff  ZoneSettingImageResizingEditResponseValue = "off"
	ZoneSettingImageResizingEditResponseValueOpen ZoneSettingImageResizingEditResponseValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingImageResizingEditResponseEditable bool

const (
	ZoneSettingImageResizingEditResponseEditableTrue  ZoneSettingImageResizingEditResponseEditable = true
	ZoneSettingImageResizingEditResponseEditableFalse ZoneSettingImageResizingEditResponseEditable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingImageResizingGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingImageResizingGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingImageResizingGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingImageResizingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingImageResizingGetResponseJSON `json:"-"`
}

// zoneSettingImageResizingGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingImageResizingGetResponse]
type zoneSettingImageResizingGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingImageResizingGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingImageResizingGetResponseID string

const (
	ZoneSettingImageResizingGetResponseIDImageResizing ZoneSettingImageResizingGetResponseID = "image_resizing"
)

// Current value of the zone setting.
type ZoneSettingImageResizingGetResponseValue string

const (
	ZoneSettingImageResizingGetResponseValueOn   ZoneSettingImageResizingGetResponseValue = "on"
	ZoneSettingImageResizingGetResponseValueOff  ZoneSettingImageResizingGetResponseValue = "off"
	ZoneSettingImageResizingGetResponseValueOpen ZoneSettingImageResizingGetResponseValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingImageResizingGetResponseEditable bool

const (
	ZoneSettingImageResizingGetResponseEditableTrue  ZoneSettingImageResizingGetResponseEditable = true
	ZoneSettingImageResizingGetResponseEditableFalse ZoneSettingImageResizingGetResponseEditable = false
)

type ZoneSettingImageResizingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[ZoneSettingImageResizingEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingImageResizingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingImageResizingEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingImageResizingEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingImageResizingEditParamsValueValue] `json:"value,required"`
}

func (r ZoneSettingImageResizingEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingImageResizingEditParamsValueID string

const (
	ZoneSettingImageResizingEditParamsValueIDImageResizing ZoneSettingImageResizingEditParamsValueID = "image_resizing"
)

// Current value of the zone setting.
type ZoneSettingImageResizingEditParamsValueValue string

const (
	ZoneSettingImageResizingEditParamsValueValueOn   ZoneSettingImageResizingEditParamsValueValue = "on"
	ZoneSettingImageResizingEditParamsValueValueOff  ZoneSettingImageResizingEditParamsValueValue = "off"
	ZoneSettingImageResizingEditParamsValueValueOpen ZoneSettingImageResizingEditParamsValueValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingImageResizingEditParamsValueEditable bool

const (
	ZoneSettingImageResizingEditParamsValueEditableTrue  ZoneSettingImageResizingEditParamsValueEditable = true
	ZoneSettingImageResizingEditParamsValueEditableFalse ZoneSettingImageResizingEditParamsValueEditable = false
)

type ZoneSettingImageResizingEditResponseEnvelope struct {
	Errors   []ZoneSettingImageResizingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingImageResizingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ZoneSettingImageResizingEditResponse             `json:"result"`
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

func (r zoneSettingImageResizingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingImageResizingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingImageResizingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingImageResizingGetResponse             `json:"result"`
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

func (r zoneSettingImageResizingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingImageResizingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingImageResizingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
