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

// ZoneSettingHttp2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingHttp2Service] method
// instead.
type ZoneSettingHttp2Service struct {
	Options []option.RequestOption
}

// NewZoneSettingHttp2Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingHttp2Service(opts ...option.RequestOption) (r *ZoneSettingHttp2Service) {
	r = &ZoneSettingHttp2Service{}
	r.Options = opts
	return
}

// Value of the HTTP2 setting.
func (r *ZoneSettingHttp2Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingHttp2UpdateParams, opts ...option.RequestOption) (res *ZoneSettingHttp2UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http2", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the HTTP2 setting.
func (r *ZoneSettingHttp2Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingHttp2ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http2", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingHttp2UpdateResponse struct {
	Errors   []ZoneSettingHttp2UpdateResponseError   `json:"errors"`
	Messages []ZoneSettingHttp2UpdateResponseMessage `json:"messages"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHttp2UpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                               `json:"success"`
	JSON    zoneSettingHttp2UpdateResponseJSON `json:"-"`
}

// zoneSettingHttp2UpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHttp2UpdateResponse]
type zoneSettingHttp2UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp2UpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHttp2UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingHttp2UpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingHttp2UpdateResponseError]
type zoneSettingHttp2UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp2UpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingHttp2UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingHttp2UpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHttp2UpdateResponseMessage]
type zoneSettingHttp2UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP2 enabled for this zone.
type ZoneSettingHttp2UpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHttp2UpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHttp2UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingHttp2UpdateResponseResultValue `json:"value"`
	JSON  zoneSettingHttp2UpdateResponseResultJSON  `json:"-"`
}

// zoneSettingHttp2UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingHttp2UpdateResponseResult]
type zoneSettingHttp2UpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHttp2UpdateResponseResultID string

const (
	ZoneSettingHttp2UpdateResponseResultIDHttp2 ZoneSettingHttp2UpdateResponseResultID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHttp2UpdateResponseResultEditable bool

const (
	ZoneSettingHttp2UpdateResponseResultEditableTrue  ZoneSettingHttp2UpdateResponseResultEditable = true
	ZoneSettingHttp2UpdateResponseResultEditableFalse ZoneSettingHttp2UpdateResponseResultEditable = false
)

// Value of the HTTP2 setting.
type ZoneSettingHttp2UpdateResponseResultValue string

const (
	ZoneSettingHttp2UpdateResponseResultValueOn  ZoneSettingHttp2UpdateResponseResultValue = "on"
	ZoneSettingHttp2UpdateResponseResultValueOff ZoneSettingHttp2UpdateResponseResultValue = "off"
)

type ZoneSettingHttp2ListResponse struct {
	Errors   []ZoneSettingHttp2ListResponseError   `json:"errors"`
	Messages []ZoneSettingHttp2ListResponseMessage `json:"messages"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHttp2ListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                             `json:"success"`
	JSON    zoneSettingHttp2ListResponseJSON `json:"-"`
}

// zoneSettingHttp2ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHttp2ListResponse]
type zoneSettingHttp2ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp2ListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingHttp2ListResponseErrorJSON `json:"-"`
}

// zoneSettingHttp2ListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingHttp2ListResponseError]
type zoneSettingHttp2ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHttp2ListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHttp2ListResponseMessageJSON `json:"-"`
}

// zoneSettingHttp2ListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHttp2ListResponseMessage]
type zoneSettingHttp2ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP2 enabled for this zone.
type ZoneSettingHttp2ListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHttp2ListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHttp2ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingHttp2ListResponseResultValue `json:"value"`
	JSON  zoneSettingHttp2ListResponseResultJSON  `json:"-"`
}

// zoneSettingHttp2ListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingHttp2ListResponseResult]
type zoneSettingHttp2ListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHttp2ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHttp2ListResponseResultID string

const (
	ZoneSettingHttp2ListResponseResultIDHttp2 ZoneSettingHttp2ListResponseResultID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHttp2ListResponseResultEditable bool

const (
	ZoneSettingHttp2ListResponseResultEditableTrue  ZoneSettingHttp2ListResponseResultEditable = true
	ZoneSettingHttp2ListResponseResultEditableFalse ZoneSettingHttp2ListResponseResultEditable = false
)

// Value of the HTTP2 setting.
type ZoneSettingHttp2ListResponseResultValue string

const (
	ZoneSettingHttp2ListResponseResultValueOn  ZoneSettingHttp2ListResponseResultValue = "on"
	ZoneSettingHttp2ListResponseResultValueOff ZoneSettingHttp2ListResponseResultValue = "off"
)

type ZoneSettingHttp2UpdateParams struct {
	// Value of the HTTP2 setting.
	Value param.Field[ZoneSettingHttp2UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingHttp2UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP2 setting.
type ZoneSettingHttp2UpdateParamsValue string

const (
	ZoneSettingHttp2UpdateParamsValueOn  ZoneSettingHttp2UpdateParamsValue = "on"
	ZoneSettingHttp2UpdateParamsValueOff ZoneSettingHttp2UpdateParamsValue = "off"
)
