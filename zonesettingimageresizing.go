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
func (r *ZoneSettingImageResizingService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingImageResizingUpdateParams, opts ...option.RequestOption) (res *ZoneSettingImageResizingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/image_resizing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
func (r *ZoneSettingImageResizingService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingImageResizingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/image_resizing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingImageResizingUpdateResponse struct {
	Errors   []ZoneSettingImageResizingUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingImageResizingUpdateResponseMessage `json:"messages"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ZoneSettingImageResizingUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingImageResizingUpdateResponseJSON
}

// zoneSettingImageResizingUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingImageResizingUpdateResponse]
type zoneSettingImageResizingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingImageResizingUpdateResponseErrorJSON
}

// zoneSettingImageResizingUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingImageResizingUpdateResponseError]
type zoneSettingImageResizingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingImageResizingUpdateResponseMessageJSON
}

// zoneSettingImageResizingUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingImageResizingUpdateResponseMessage]
type zoneSettingImageResizingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingImageResizingUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingImageResizingUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingImageResizingUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value ZoneSettingImageResizingUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingImageResizingUpdateResponseResultJSON
}

// zoneSettingImageResizingUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingImageResizingUpdateResponseResult]
type zoneSettingImageResizingUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingImageResizingUpdateResponseResultID string

const (
	ZoneSettingImageResizingUpdateResponseResultIDImageResizing ZoneSettingImageResizingUpdateResponseResultID = "image_resizing"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingImageResizingUpdateResponseResultEditable bool

const (
	ZoneSettingImageResizingUpdateResponseResultEditableTrue  ZoneSettingImageResizingUpdateResponseResultEditable = true
	ZoneSettingImageResizingUpdateResponseResultEditableFalse ZoneSettingImageResizingUpdateResponseResultEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingImageResizingUpdateResponseResultValue string

const (
	ZoneSettingImageResizingUpdateResponseResultValueOn   ZoneSettingImageResizingUpdateResponseResultValue = "on"
	ZoneSettingImageResizingUpdateResponseResultValueOff  ZoneSettingImageResizingUpdateResponseResultValue = "off"
	ZoneSettingImageResizingUpdateResponseResultValueOpen ZoneSettingImageResizingUpdateResponseResultValue = "open"
)

type ZoneSettingImageResizingListResponse struct {
	Errors   []ZoneSettingImageResizingListResponseError   `json:"errors"`
	Messages []ZoneSettingImageResizingListResponseMessage `json:"messages"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ZoneSettingImageResizingListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingImageResizingListResponseJSON
}

// zoneSettingImageResizingListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingImageResizingListResponse]
type zoneSettingImageResizingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingImageResizingListResponseErrorJSON
}

// zoneSettingImageResizingListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingImageResizingListResponseError]
type zoneSettingImageResizingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingImageResizingListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingImageResizingListResponseMessageJSON
}

// zoneSettingImageResizingListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingImageResizingListResponseMessage]
type zoneSettingImageResizingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingImageResizingListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingImageResizingListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingImageResizingListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value ZoneSettingImageResizingListResponseResultValue `json:"value"`
	JSON  zoneSettingImageResizingListResponseResultJSON
}

// zoneSettingImageResizingListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingImageResizingListResponseResult]
type zoneSettingImageResizingListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingImageResizingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingImageResizingListResponseResultID string

const (
	ZoneSettingImageResizingListResponseResultIDImageResizing ZoneSettingImageResizingListResponseResultID = "image_resizing"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingImageResizingListResponseResultEditable bool

const (
	ZoneSettingImageResizingListResponseResultEditableTrue  ZoneSettingImageResizingListResponseResultEditable = true
	ZoneSettingImageResizingListResponseResultEditableFalse ZoneSettingImageResizingListResponseResultEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingImageResizingListResponseResultValue string

const (
	ZoneSettingImageResizingListResponseResultValueOn   ZoneSettingImageResizingListResponseResultValue = "on"
	ZoneSettingImageResizingListResponseResultValueOff  ZoneSettingImageResizingListResponseResultValue = "off"
	ZoneSettingImageResizingListResponseResultValueOpen ZoneSettingImageResizingListResponseResultValue = "open"
)

type ZoneSettingImageResizingUpdateParams struct {
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value param.Field[ZoneSettingImageResizingUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingImageResizingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingImageResizingUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingImageResizingUpdateParamsValueID] `json:"id"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value param.Field[ZoneSettingImageResizingUpdateParamsValueValue] `json:"value"`
}

func (r ZoneSettingImageResizingUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingImageResizingUpdateParamsValueID string

const (
	ZoneSettingImageResizingUpdateParamsValueIDImageResizing ZoneSettingImageResizingUpdateParamsValueID = "image_resizing"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingImageResizingUpdateParamsValueEditable bool

const (
	ZoneSettingImageResizingUpdateParamsValueEditableTrue  ZoneSettingImageResizingUpdateParamsValueEditable = true
	ZoneSettingImageResizingUpdateParamsValueEditableFalse ZoneSettingImageResizingUpdateParamsValueEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingImageResizingUpdateParamsValueValue string

const (
	ZoneSettingImageResizingUpdateParamsValueValueOn   ZoneSettingImageResizingUpdateParamsValueValue = "on"
	ZoneSettingImageResizingUpdateParamsValueValueOff  ZoneSettingImageResizingUpdateParamsValueValue = "off"
	ZoneSettingImageResizingUpdateParamsValueValueOpen ZoneSettingImageResizingUpdateParamsValueValue = "open"
)
