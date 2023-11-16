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

// ZoneSettingEarlyHintService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingEarlyHintService]
// method instead.
type ZoneSettingEarlyHintService struct {
	Options []option.RequestOption
}

// NewZoneSettingEarlyHintService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingEarlyHintService(opts ...option.RequestOption) (r *ZoneSettingEarlyHintService) {
	r = &ZoneSettingEarlyHintService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *ZoneSettingEarlyHintService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingEarlyHintUpdateParams, opts ...option.RequestOption) (res *ZoneSettingEarlyHintUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/early_hints", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *ZoneSettingEarlyHintService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingEarlyHintListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/early_hints", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingEarlyHintUpdateResponse struct {
	Errors   []ZoneSettingEarlyHintUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingEarlyHintUpdateResponseMessage `json:"messages"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZoneSettingEarlyHintUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                   `json:"success"`
	JSON    zoneSettingEarlyHintUpdateResponseJSON `json:"-"`
}

// zoneSettingEarlyHintUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingEarlyHintUpdateResponse]
type zoneSettingEarlyHintUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingEarlyHintUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingEarlyHintUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintUpdateResponseError]
type zoneSettingEarlyHintUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingEarlyHintUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingEarlyHintUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintUpdateResponseMessage]
type zoneSettingEarlyHintUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEarlyHintUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingEarlyHintUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEarlyHintUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEarlyHintUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingEarlyHintUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingEarlyHintUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintUpdateResponseResult]
type zoneSettingEarlyHintUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingEarlyHintUpdateResponseResultID string

const (
	ZoneSettingEarlyHintUpdateResponseResultIDEarlyHints ZoneSettingEarlyHintUpdateResponseResultID = "early_hints"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEarlyHintUpdateResponseResultEditable bool

const (
	ZoneSettingEarlyHintUpdateResponseResultEditableTrue  ZoneSettingEarlyHintUpdateResponseResultEditable = true
	ZoneSettingEarlyHintUpdateResponseResultEditableFalse ZoneSettingEarlyHintUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingEarlyHintUpdateResponseResultValue string

const (
	ZoneSettingEarlyHintUpdateResponseResultValueOn  ZoneSettingEarlyHintUpdateResponseResultValue = "on"
	ZoneSettingEarlyHintUpdateResponseResultValueOff ZoneSettingEarlyHintUpdateResponseResultValue = "off"
)

type ZoneSettingEarlyHintListResponse struct {
	Errors   []ZoneSettingEarlyHintListResponseError   `json:"errors"`
	Messages []ZoneSettingEarlyHintListResponseMessage `json:"messages"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result ZoneSettingEarlyHintListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                 `json:"success"`
	JSON    zoneSettingEarlyHintListResponseJSON `json:"-"`
}

// zoneSettingEarlyHintListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingEarlyHintListResponse]
type zoneSettingEarlyHintListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintListResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingEarlyHintListResponseErrorJSON `json:"-"`
}

// zoneSettingEarlyHintListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintListResponseError]
type zoneSettingEarlyHintListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEarlyHintListResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingEarlyHintListResponseMessageJSON `json:"-"`
}

// zoneSettingEarlyHintListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintListResponseMessage]
type zoneSettingEarlyHintListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEarlyHintListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingEarlyHintListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEarlyHintListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEarlyHintListResponseResultValue `json:"value"`
	JSON  zoneSettingEarlyHintListResponseResultJSON  `json:"-"`
}

// zoneSettingEarlyHintListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingEarlyHintListResponseResult]
type zoneSettingEarlyHintListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEarlyHintListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingEarlyHintListResponseResultID string

const (
	ZoneSettingEarlyHintListResponseResultIDEarlyHints ZoneSettingEarlyHintListResponseResultID = "early_hints"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEarlyHintListResponseResultEditable bool

const (
	ZoneSettingEarlyHintListResponseResultEditableTrue  ZoneSettingEarlyHintListResponseResultEditable = true
	ZoneSettingEarlyHintListResponseResultEditableFalse ZoneSettingEarlyHintListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingEarlyHintListResponseResultValue string

const (
	ZoneSettingEarlyHintListResponseResultValueOn  ZoneSettingEarlyHintListResponseResultValue = "on"
	ZoneSettingEarlyHintListResponseResultValueOff ZoneSettingEarlyHintListResponseResultValue = "off"
)

type ZoneSettingEarlyHintUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingEarlyHintUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingEarlyHintUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingEarlyHintUpdateParamsValue string

const (
	ZoneSettingEarlyHintUpdateParamsValueOn  ZoneSettingEarlyHintUpdateParamsValue = "on"
	ZoneSettingEarlyHintUpdateParamsValueOff ZoneSettingEarlyHintUpdateParamsValue = "off"
)
