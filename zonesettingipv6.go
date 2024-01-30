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

// ZoneSettingIpv6Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingIpv6Service] method
// instead.
type ZoneSettingIpv6Service struct {
	Options []option.RequestOption
}

// NewZoneSettingIpv6Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingIpv6Service(opts ...option.RequestOption) (r *ZoneSettingIpv6Service) {
	r = &ZoneSettingIpv6Service{}
	r.Options = opts
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZoneSettingIpv6Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingIpv6UpdateParams, opts ...option.RequestOption) (res *ZoneSettingIpv6UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ipv6", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZoneSettingIpv6Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingIpv6ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ipv6", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingIpv6UpdateResponse struct {
	Errors   []ZoneSettingIpv6UpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingIpv6UpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	Result  ZoneSettingIpv6UpdateResponseResult `json:"result"`
	JSON    zoneSettingIpv6UpdateResponseJSON   `json:"-"`
}

// zoneSettingIpv6UpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingIpv6UpdateResponse]
type zoneSettingIpv6UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIpv6UpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingIpv6UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingIpv6UpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingIpv6UpdateResponseError]
type zoneSettingIpv6UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIpv6UpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingIpv6UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingIpv6UpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingIpv6UpdateResponseMessage]
type zoneSettingIpv6UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIpv6UpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingIpv6UpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingIpv6UpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIpv6UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingIpv6UpdateResponseResultJSON `json:"-"`
}

// zoneSettingIpv6UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingIpv6UpdateResponseResult]
type zoneSettingIpv6UpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIpv6UpdateResponseResultID string

const (
	ZoneSettingIpv6UpdateResponseResultIDIpv6 ZoneSettingIpv6UpdateResponseResultID = "ipv6"
)

// Value of the zone setting.
type ZoneSettingIpv6UpdateResponseResultValue string

const (
	ZoneSettingIpv6UpdateResponseResultValueOff ZoneSettingIpv6UpdateResponseResultValue = "off"
	ZoneSettingIpv6UpdateResponseResultValueOn  ZoneSettingIpv6UpdateResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIpv6UpdateResponseResultEditable bool

const (
	ZoneSettingIpv6UpdateResponseResultEditableTrue  ZoneSettingIpv6UpdateResponseResultEditable = true
	ZoneSettingIpv6UpdateResponseResultEditableFalse ZoneSettingIpv6UpdateResponseResultEditable = false
)

type ZoneSettingIpv6ListResponse struct {
	Errors   []ZoneSettingIpv6ListResponseError   `json:"errors,required"`
	Messages []ZoneSettingIpv6ListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                              `json:"success,required"`
	Result  ZoneSettingIpv6ListResponseResult `json:"result"`
	JSON    zoneSettingIpv6ListResponseJSON   `json:"-"`
}

// zoneSettingIpv6ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingIpv6ListResponse]
type zoneSettingIpv6ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIpv6ListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneSettingIpv6ListResponseErrorJSON `json:"-"`
}

// zoneSettingIpv6ListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingIpv6ListResponseError]
type zoneSettingIpv6ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIpv6ListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingIpv6ListResponseMessageJSON `json:"-"`
}

// zoneSettingIpv6ListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingIpv6ListResponseMessage]
type zoneSettingIpv6ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIpv6ListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingIpv6ListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingIpv6ListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingIpv6ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingIpv6ListResponseResultJSON `json:"-"`
}

// zoneSettingIpv6ListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingIpv6ListResponseResult]
type zoneSettingIpv6ListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIpv6ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingIpv6ListResponseResultID string

const (
	ZoneSettingIpv6ListResponseResultIDIpv6 ZoneSettingIpv6ListResponseResultID = "ipv6"
)

// Value of the zone setting.
type ZoneSettingIpv6ListResponseResultValue string

const (
	ZoneSettingIpv6ListResponseResultValueOff ZoneSettingIpv6ListResponseResultValue = "off"
	ZoneSettingIpv6ListResponseResultValueOn  ZoneSettingIpv6ListResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingIpv6ListResponseResultEditable bool

const (
	ZoneSettingIpv6ListResponseResultEditableTrue  ZoneSettingIpv6ListResponseResultEditable = true
	ZoneSettingIpv6ListResponseResultEditableFalse ZoneSettingIpv6ListResponseResultEditable = false
)

type ZoneSettingIpv6UpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingIpv6UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingIpv6UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingIpv6UpdateParamsValue string

const (
	ZoneSettingIpv6UpdateParamsValueOff ZoneSettingIpv6UpdateParamsValue = "off"
	ZoneSettingIpv6UpdateParamsValueOn  ZoneSettingIpv6UpdateParamsValue = "on"
)
