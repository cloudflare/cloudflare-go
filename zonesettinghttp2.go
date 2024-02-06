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

// ZoneSettingHTTP2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingHTTP2Service] method
// instead.
type ZoneSettingHTTP2Service struct {
	Options []option.RequestOption
}

// NewZoneSettingHTTP2Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingHTTP2Service(opts ...option.RequestOption) (r *ZoneSettingHTTP2Service) {
	r = &ZoneSettingHTTP2Service{}
	r.Options = opts
	return
}

// Value of the HTTP2 setting.
func (r *ZoneSettingHTTP2Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingHTTP2UpdateParams, opts ...option.RequestOption) (res *ZoneSettingHTTP2UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http2", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the HTTP2 setting.
func (r *ZoneSettingHTTP2Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingHTTP2ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http2", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingHTTP2UpdateResponse struct {
	Errors   []ZoneSettingHTTP2UpdateResponseError   `json:"errors"`
	Messages []ZoneSettingHTTP2UpdateResponseMessage `json:"messages"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHTTP2UpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                               `json:"success"`
	JSON    zoneSettingHTTP2UpdateResponseJSON `json:"-"`
}

// zoneSettingHTTP2UpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP2UpdateResponse]
type zoneSettingHTTP2UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP2UpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHTTP2UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingHTTP2UpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP2UpdateResponseError]
type zoneSettingHTTP2UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP2UpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingHTTP2UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingHTTP2UpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP2UpdateResponseMessage]
type zoneSettingHTTP2UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP2 enabled for this zone.
type ZoneSettingHTTP2UpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP2UpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP2UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingHTTP2UpdateResponseResultValue `json:"value"`
	JSON  zoneSettingHTTP2UpdateResponseResultJSON  `json:"-"`
}

// zoneSettingHTTP2UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP2UpdateResponseResult]
type zoneSettingHTTP2UpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHTTP2UpdateResponseResultID string

const (
	ZoneSettingHTTP2UpdateResponseResultIDHTTP2 ZoneSettingHTTP2UpdateResponseResultID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP2UpdateResponseResultEditable bool

const (
	ZoneSettingHTTP2UpdateResponseResultEditableTrue  ZoneSettingHTTP2UpdateResponseResultEditable = true
	ZoneSettingHTTP2UpdateResponseResultEditableFalse ZoneSettingHTTP2UpdateResponseResultEditable = false
)

// Value of the HTTP2 setting.
type ZoneSettingHTTP2UpdateResponseResultValue string

const (
	ZoneSettingHTTP2UpdateResponseResultValueOn  ZoneSettingHTTP2UpdateResponseResultValue = "on"
	ZoneSettingHTTP2UpdateResponseResultValueOff ZoneSettingHTTP2UpdateResponseResultValue = "off"
)

type ZoneSettingHTTP2ListResponse struct {
	Errors   []ZoneSettingHTTP2ListResponseError   `json:"errors"`
	Messages []ZoneSettingHTTP2ListResponseMessage `json:"messages"`
	// HTTP2 enabled for this zone.
	Result ZoneSettingHTTP2ListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                             `json:"success"`
	JSON    zoneSettingHTTP2ListResponseJSON `json:"-"`
}

// zoneSettingHTTP2ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP2ListResponse]
type zoneSettingHTTP2ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP2ListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingHTTP2ListResponseErrorJSON `json:"-"`
}

// zoneSettingHTTP2ListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP2ListResponseError]
type zoneSettingHTTP2ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP2ListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHTTP2ListResponseMessageJSON `json:"-"`
}

// zoneSettingHTTP2ListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP2ListResponseMessage]
type zoneSettingHTTP2ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP2 enabled for this zone.
type ZoneSettingHTTP2ListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP2ListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP2ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingHTTP2ListResponseResultValue `json:"value"`
	JSON  zoneSettingHTTP2ListResponseResultJSON  `json:"-"`
}

// zoneSettingHTTP2ListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP2ListResponseResult]
type zoneSettingHTTP2ListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP2ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHTTP2ListResponseResultID string

const (
	ZoneSettingHTTP2ListResponseResultIDHTTP2 ZoneSettingHTTP2ListResponseResultID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP2ListResponseResultEditable bool

const (
	ZoneSettingHTTP2ListResponseResultEditableTrue  ZoneSettingHTTP2ListResponseResultEditable = true
	ZoneSettingHTTP2ListResponseResultEditableFalse ZoneSettingHTTP2ListResponseResultEditable = false
)

// Value of the HTTP2 setting.
type ZoneSettingHTTP2ListResponseResultValue string

const (
	ZoneSettingHTTP2ListResponseResultValueOn  ZoneSettingHTTP2ListResponseResultValue = "on"
	ZoneSettingHTTP2ListResponseResultValueOff ZoneSettingHTTP2ListResponseResultValue = "off"
)

type ZoneSettingHTTP2UpdateParams struct {
	// Value of the HTTP2 setting.
	Value param.Field[ZoneSettingHTTP2UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingHTTP2UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP2 setting.
type ZoneSettingHTTP2UpdateParamsValue string

const (
	ZoneSettingHTTP2UpdateParamsValueOn  ZoneSettingHTTP2UpdateParamsValue = "on"
	ZoneSettingHTTP2UpdateParamsValueOff ZoneSettingHTTP2UpdateParamsValue = "off"
)
