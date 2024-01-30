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

// ZoneSettingHTTP3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingHTTP3Service] method
// instead.
type ZoneSettingHTTP3Service struct {
	Options []option.RequestOption
}

// NewZoneSettingHTTP3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingHTTP3Service(opts ...option.RequestOption) (r *ZoneSettingHTTP3Service) {
	r = &ZoneSettingHTTP3Service{}
	r.Options = opts
	return
}

// Value of the HTTP3 setting.
func (r *ZoneSettingHTTP3Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingHTTP3UpdateParams, opts ...option.RequestOption) (res *ZoneSettingHTTP3UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http3", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the HTTP3 setting.
func (r *ZoneSettingHTTP3Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingHTTP3ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http3", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingHTTP3UpdateResponse struct {
	Errors   []ZoneSettingHTTP3UpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingHTTP3UpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                 `json:"success,required"`
	Result  ZoneSettingHTTP3UpdateResponseResult `json:"result"`
	JSON    zoneSettingHTTP3UpdateResponseJSON   `json:"-"`
}

// zoneSettingHTTP3UpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP3UpdateResponse]
type zoneSettingHTTP3UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3UpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHTTP3UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingHTTP3UpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP3UpdateResponseError]
type zoneSettingHTTP3UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3UpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingHTTP3UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingHTTP3UpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP3UpdateResponseMessage]
type zoneSettingHTTP3UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3UpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP3UpdateResponseResultID `json:"id,required"`
	// Value of the HTTP3 setting.
	Value ZoneSettingHTTP3UpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP3UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP3UpdateResponseResultJSON `json:"-"`
}

// zoneSettingHTTP3UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP3UpdateResponseResult]
type zoneSettingHTTP3UpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHTTP3UpdateResponseResultID string

const (
	ZoneSettingHTTP3UpdateResponseResultIDHTTP3 ZoneSettingHTTP3UpdateResponseResultID = "http3"
)

// Value of the HTTP3 setting.
type ZoneSettingHTTP3UpdateResponseResultValue string

const (
	ZoneSettingHTTP3UpdateResponseResultValueOn  ZoneSettingHTTP3UpdateResponseResultValue = "on"
	ZoneSettingHTTP3UpdateResponseResultValueOff ZoneSettingHTTP3UpdateResponseResultValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP3UpdateResponseResultEditable bool

const (
	ZoneSettingHTTP3UpdateResponseResultEditableTrue  ZoneSettingHTTP3UpdateResponseResultEditable = true
	ZoneSettingHTTP3UpdateResponseResultEditableFalse ZoneSettingHTTP3UpdateResponseResultEditable = false
)

type ZoneSettingHTTP3ListResponse struct {
	Errors   []ZoneSettingHTTP3ListResponseError   `json:"errors,required"`
	Messages []ZoneSettingHTTP3ListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                               `json:"success,required"`
	Result  ZoneSettingHTTP3ListResponseResult `json:"result"`
	JSON    zoneSettingHTTP3ListResponseJSON   `json:"-"`
}

// zoneSettingHTTP3ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP3ListResponse]
type zoneSettingHTTP3ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3ListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingHTTP3ListResponseErrorJSON `json:"-"`
}

// zoneSettingHTTP3ListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP3ListResponseError]
type zoneSettingHTTP3ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3ListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingHTTP3ListResponseMessageJSON `json:"-"`
}

// zoneSettingHTTP3ListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingHTTP3ListResponseMessage]
type zoneSettingHTTP3ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHTTP3ListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingHTTP3ListResponseResultID `json:"id,required"`
	// Value of the HTTP3 setting.
	Value ZoneSettingHTTP3ListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingHTTP3ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingHTTP3ListResponseResultJSON `json:"-"`
}

// zoneSettingHTTP3ListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingHTTP3ListResponseResult]
type zoneSettingHTTP3ListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHTTP3ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingHTTP3ListResponseResultID string

const (
	ZoneSettingHTTP3ListResponseResultIDHTTP3 ZoneSettingHTTP3ListResponseResultID = "http3"
)

// Value of the HTTP3 setting.
type ZoneSettingHTTP3ListResponseResultValue string

const (
	ZoneSettingHTTP3ListResponseResultValueOn  ZoneSettingHTTP3ListResponseResultValue = "on"
	ZoneSettingHTTP3ListResponseResultValueOff ZoneSettingHTTP3ListResponseResultValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingHTTP3ListResponseResultEditable bool

const (
	ZoneSettingHTTP3ListResponseResultEditableTrue  ZoneSettingHTTP3ListResponseResultEditable = true
	ZoneSettingHTTP3ListResponseResultEditableFalse ZoneSettingHTTP3ListResponseResultEditable = false
)

type ZoneSettingHTTP3UpdateParams struct {
	// Value of the HTTP3 setting.
	Value param.Field[ZoneSettingHTTP3UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingHTTP3UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP3 setting.
type ZoneSettingHTTP3UpdateParamsValue string

const (
	ZoneSettingHTTP3UpdateParamsValueOn  ZoneSettingHTTP3UpdateParamsValue = "on"
	ZoneSettingHTTP3UpdateParamsValueOff ZoneSettingHTTP3UpdateParamsValue = "off"
)
