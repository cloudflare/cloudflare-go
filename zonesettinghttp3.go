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

// ZoneSettingHttp3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingHttp3Service] method
// instead.
type ZoneSettingHttp3Service struct {
	Options []option.RequestOption
}

// NewZoneSettingHttp3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingHttp3Service(opts ...option.RequestOption) (r *ZoneSettingHttp3Service) {
	r = &ZoneSettingHttp3Service{}
	r.Options = opts
	return
}

// Value of the HTTP3 setting.
func (r *ZoneSettingHttp3Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingHttp3UpdateParams, opts ...option.RequestOption) (res *ZoneSettingHttp3UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http3", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the HTTP3 setting.
func (r *ZoneSettingHttp3Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingHttp3ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http3", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingHttp3UpdateResponse struct {
	Errors   []ZoneSettingHttp3UpdateResponseError   `json:"errors"`
	Messages []ZoneSettingHttp3UpdateResponseMessage `json:"messages"`
	// HTTP3 enabled for this zone.
	Result ZoneSettingHttp3UpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                               `json:"success"`
	JSON    zoneSettingHttp3UpdateResponseJSON `json:"-"`
}

// zoneSettingHttp3UpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHttp3UpdateResponse]
type zoneSettingHttp3UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp3UpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHttp3UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingHttp3UpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingHttp3UpdateResponseError]
type zoneSettingHttp3UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp3UpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingHttp3UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingHttp3UpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHttp3UpdateResponseMessage]
type zoneSettingHttp3UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP3 enabled for this zone.
type ZoneSettingHttp3UpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHttp3UpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHttp3UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value ZoneSettingHttp3UpdateResponseResultValue `json:"value"`
	JSON  zoneSettingHttp3UpdateResponseResultJSON  `json:"-"`
}

// zoneSettingHttp3UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingHttp3UpdateResponseResult]
type zoneSettingHttp3UpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHttp3UpdateResponseResultID string

const (
	ZoneSettingHttp3UpdateResponseResultIDHttp3 ZoneSettingHttp3UpdateResponseResultID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHttp3UpdateResponseResultEditable bool

const (
	ZoneSettingHttp3UpdateResponseResultEditableTrue  ZoneSettingHttp3UpdateResponseResultEditable = true
	ZoneSettingHttp3UpdateResponseResultEditableFalse ZoneSettingHttp3UpdateResponseResultEditable = false
)

// Value of the HTTP3 setting.
type ZoneSettingHttp3UpdateResponseResultValue string

const (
	ZoneSettingHttp3UpdateResponseResultValueOn  ZoneSettingHttp3UpdateResponseResultValue = "on"
	ZoneSettingHttp3UpdateResponseResultValueOff ZoneSettingHttp3UpdateResponseResultValue = "off"
)

type ZoneSettingHttp3ListResponse struct {
	Errors   []ZoneSettingHttp3ListResponseError   `json:"errors"`
	Messages []ZoneSettingHttp3ListResponseMessage `json:"messages"`
	// HTTP3 enabled for this zone.
	Result ZoneSettingHttp3ListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                             `json:"success"`
	JSON    zoneSettingHttp3ListResponseJSON `json:"-"`
}

// zoneSettingHttp3ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHttp3ListResponse]
type zoneSettingHttp3ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp3ListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingHttp3ListResponseErrorJSON `json:"-"`
}

// zoneSettingHttp3ListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingHttp3ListResponseError]
type zoneSettingHttp3ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp3ListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHttp3ListResponseMessageJSON `json:"-"`
}

// zoneSettingHttp3ListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHttp3ListResponseMessage]
type zoneSettingHttp3ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP3 enabled for this zone.
type ZoneSettingHttp3ListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHttp3ListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHttp3ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value ZoneSettingHttp3ListResponseResultValue `json:"value"`
	JSON  zoneSettingHttp3ListResponseResultJSON  `json:"-"`
}

// zoneSettingHttp3ListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingHttp3ListResponseResult]
type zoneSettingHttp3ListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp3ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHttp3ListResponseResultID string

const (
	ZoneSettingHttp3ListResponseResultIDHttp3 ZoneSettingHttp3ListResponseResultID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHttp3ListResponseResultEditable bool

const (
	ZoneSettingHttp3ListResponseResultEditableTrue  ZoneSettingHttp3ListResponseResultEditable = true
	ZoneSettingHttp3ListResponseResultEditableFalse ZoneSettingHttp3ListResponseResultEditable = false
)

// Value of the HTTP3 setting.
type ZoneSettingHttp3ListResponseResultValue string

const (
	ZoneSettingHttp3ListResponseResultValueOn  ZoneSettingHttp3ListResponseResultValue = "on"
	ZoneSettingHttp3ListResponseResultValueOff ZoneSettingHttp3ListResponseResultValue = "off"
)

type ZoneSettingHttp3UpdateParams struct {
	// Value of the HTTP3 setting.
	Value param.Field[ZoneSettingHttp3UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingHttp3UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP3 setting.
type ZoneSettingHttp3UpdateParamsValue string

const (
	ZoneSettingHttp3UpdateParamsValueOn  ZoneSettingHttp3UpdateParamsValue = "on"
	ZoneSettingHttp3UpdateParamsValueOff ZoneSettingHttp3UpdateParamsValue = "off"
)
