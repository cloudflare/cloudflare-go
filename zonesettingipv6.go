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

// ZoneSettingIPV6Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingIPV6Service] method
// instead.
type ZoneSettingIPV6Service struct {
	Options []option.RequestOption
}

// NewZoneSettingIPV6Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingIPV6Service(opts ...option.RequestOption) (r *ZoneSettingIPV6Service) {
	r = &ZoneSettingIPV6Service{}
	r.Options = opts
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZoneSettingIPV6Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingIPV6UpdateParams, opts ...option.RequestOption) (res *ZoneSettingIPV6UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ipv6", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZoneSettingIPV6Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingIPV6ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ipv6", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingIPV6UpdateResponse struct {
	Errors   []ZoneSettingIPV6UpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingIPV6UpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	Result  ZoneSettingIPV6UpdateResponseResult `json:"result"`
	JSON    zoneSettingIPV6UpdateResponseJSON   `json:"-"`
}

// zoneSettingIPV6UpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6UpdateResponse]
type zoneSettingIPV6UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6UpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingIPV6UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingIPV6UpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6UpdateResponseError]
type zoneSettingIPV6UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6UpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingIPV6UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingIPV6UpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingIPV6UpdateResponseMessage]
type zoneSettingIPV6UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6UpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingIPV6UpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingIPV6UpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIPV6UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingIPV6UpdateResponseResultJSON `json:"-"`
}

// zoneSettingIPV6UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingIPV6UpdateResponseResult]
type zoneSettingIPV6UpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIPV6UpdateResponseResultID string

const (
	ZoneSettingIPV6UpdateResponseResultIDIPV6 ZoneSettingIPV6UpdateResponseResultID = "ipv6"
)

// Value of the zone setting.
type ZoneSettingIPV6UpdateResponseResultValue string

const (
	ZoneSettingIPV6UpdateResponseResultValueOff ZoneSettingIPV6UpdateResponseResultValue = "off"
	ZoneSettingIPV6UpdateResponseResultValueOn  ZoneSettingIPV6UpdateResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIPV6UpdateResponseResultEditable bool

const (
	ZoneSettingIPV6UpdateResponseResultEditableTrue  ZoneSettingIPV6UpdateResponseResultEditable = true
	ZoneSettingIPV6UpdateResponseResultEditableFalse ZoneSettingIPV6UpdateResponseResultEditable = false
)

type ZoneSettingIPV6ListResponse struct {
	Errors   []ZoneSettingIPV6ListResponseError   `json:"errors,required"`
	Messages []ZoneSettingIPV6ListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                              `json:"success,required"`
	Result  ZoneSettingIPV6ListResponseResult `json:"result"`
	JSON    zoneSettingIPV6ListResponseJSON   `json:"-"`
}

// zoneSettingIPV6ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6ListResponse]
type zoneSettingIPV6ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6ListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneSettingIPV6ListResponseErrorJSON `json:"-"`
}

// zoneSettingIPV6ListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6ListResponseError]
type zoneSettingIPV6ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6ListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingIPV6ListResponseMessageJSON `json:"-"`
}

// zoneSettingIPV6ListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6ListResponseMessage]
type zoneSettingIPV6ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPV6ListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingIPV6ListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingIPV6ListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIPV6ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingIPV6ListResponseResultJSON `json:"-"`
}

// zoneSettingIPV6ListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingIPV6ListResponseResult]
type zoneSettingIPV6ListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPV6ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIPV6ListResponseResultID string

const (
	ZoneSettingIPV6ListResponseResultIDIPV6 ZoneSettingIPV6ListResponseResultID = "ipv6"
)

// Value of the zone setting.
type ZoneSettingIPV6ListResponseResultValue string

const (
	ZoneSettingIPV6ListResponseResultValueOff ZoneSettingIPV6ListResponseResultValue = "off"
	ZoneSettingIPV6ListResponseResultValueOn  ZoneSettingIPV6ListResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIPV6ListResponseResultEditable bool

const (
	ZoneSettingIPV6ListResponseResultEditableTrue  ZoneSettingIPV6ListResponseResultEditable = true
	ZoneSettingIPV6ListResponseResultEditableFalse ZoneSettingIPV6ListResponseResultEditable = false
)

type ZoneSettingIPV6UpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingIPV6UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingIPV6UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingIPV6UpdateParamsValue string

const (
	ZoneSettingIPV6UpdateParamsValueOff ZoneSettingIPV6UpdateParamsValue = "off"
	ZoneSettingIPV6UpdateParamsValueOn  ZoneSettingIPV6UpdateParamsValue = "on"
)
