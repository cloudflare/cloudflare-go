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

// ZoneSettingTLS1_3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingTLS1_3Service] method
// instead.
type ZoneSettingTLS1_3Service struct {
	Options []option.RequestOption
}

// NewZoneSettingTLS1_3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingTLS1_3Service(opts ...option.RequestOption) (r *ZoneSettingTLS1_3Service) {
	r = &ZoneSettingTLS1_3Service{}
	r.Options = opts
	return
}

// Changes TLS 1.3 setting.
func (r *ZoneSettingTLS1_3Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingTLS1_3UpdateParams, opts ...option.RequestOption) (res *ZoneSettingTls1_3UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets TLS 1.3 setting enabled for a zone.
func (r *ZoneSettingTLS1_3Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingTls1_3ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingTls1_3UpdateResponse struct {
	Errors   []ZoneSettingTls1_3UpdateResponseError   `json:"errors"`
	Messages []ZoneSettingTls1_3UpdateResponseMessage `json:"messages"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result ZoneSettingTls1_3UpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                `json:"success"`
	JSON    zoneSettingTls1_3UpdateResponseJSON `json:"-"`
}

// zoneSettingTls1_3UpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingTls1_3UpdateResponse]
type zoneSettingTls1_3UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTls1_3UpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingTls1_3UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingTls1_3UpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingTls1_3UpdateResponseError]
type zoneSettingTls1_3UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTls1_3UpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSettingTls1_3UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingTls1_3UpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingTls1_3UpdateResponseMessage]
type zoneSettingTls1_3UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingTls1_3UpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTls1_3UpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTls1_3UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingTls1_3UpdateResponseResultValue `json:"value"`
	JSON  zoneSettingTls1_3UpdateResponseResultJSON  `json:"-"`
}

// zoneSettingTls1_3UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingTls1_3UpdateResponseResult]
type zoneSettingTls1_3UpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTls1_3UpdateResponseResultID string

const (
	ZoneSettingTls1_3UpdateResponseResultIDTLS1_3 ZoneSettingTls1_3UpdateResponseResultID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTls1_3UpdateResponseResultEditable bool

const (
	ZoneSettingTls1_3UpdateResponseResultEditableTrue  ZoneSettingTls1_3UpdateResponseResultEditable = true
	ZoneSettingTls1_3UpdateResponseResultEditableFalse ZoneSettingTls1_3UpdateResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingTls1_3UpdateResponseResultValue string

const (
	ZoneSettingTls1_3UpdateResponseResultValueOn  ZoneSettingTls1_3UpdateResponseResultValue = "on"
	ZoneSettingTls1_3UpdateResponseResultValueOff ZoneSettingTls1_3UpdateResponseResultValue = "off"
	ZoneSettingTls1_3UpdateResponseResultValueZrt ZoneSettingTls1_3UpdateResponseResultValue = "zrt"
)

type ZoneSettingTls1_3ListResponse struct {
	Errors   []ZoneSettingTls1_3ListResponseError   `json:"errors"`
	Messages []ZoneSettingTls1_3ListResponseMessage `json:"messages"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result ZoneSettingTls1_3ListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    zoneSettingTls1_3ListResponseJSON `json:"-"`
}

// zoneSettingTls1_3ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingTls1_3ListResponse]
type zoneSettingTls1_3ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTls1_3ListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingTls1_3ListResponseErrorJSON `json:"-"`
}

// zoneSettingTls1_3ListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingTls1_3ListResponseError]
type zoneSettingTls1_3ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTls1_3ListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingTls1_3ListResponseMessageJSON `json:"-"`
}

// zoneSettingTls1_3ListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingTls1_3ListResponseMessage]
type zoneSettingTls1_3ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingTls1_3ListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTls1_3ListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTls1_3ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingTls1_3ListResponseResultValue `json:"value"`
	JSON  zoneSettingTls1_3ListResponseResultJSON  `json:"-"`
}

// zoneSettingTls1_3ListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingTls1_3ListResponseResult]
type zoneSettingTls1_3ListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTls1_3ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTls1_3ListResponseResultID string

const (
	ZoneSettingTls1_3ListResponseResultIDTLS1_3 ZoneSettingTls1_3ListResponseResultID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTls1_3ListResponseResultEditable bool

const (
	ZoneSettingTls1_3ListResponseResultEditableTrue  ZoneSettingTls1_3ListResponseResultEditable = true
	ZoneSettingTls1_3ListResponseResultEditableFalse ZoneSettingTls1_3ListResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingTls1_3ListResponseResultValue string

const (
	ZoneSettingTls1_3ListResponseResultValueOn  ZoneSettingTls1_3ListResponseResultValue = "on"
	ZoneSettingTls1_3ListResponseResultValueOff ZoneSettingTls1_3ListResponseResultValue = "off"
	ZoneSettingTls1_3ListResponseResultValueZrt ZoneSettingTls1_3ListResponseResultValue = "zrt"
)

type ZoneSettingTLS1_3UpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingTls1_3UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingTLS1_3UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingTls1_3UpdateParamsValue string

const (
	ZoneSettingTls1_3UpdateParamsValueOn  ZoneSettingTls1_3UpdateParamsValue = "on"
	ZoneSettingTls1_3UpdateParamsValueOff ZoneSettingTls1_3UpdateParamsValue = "off"
	ZoneSettingTls1_3UpdateParamsValueZrt ZoneSettingTls1_3UpdateParamsValue = "zrt"
)
