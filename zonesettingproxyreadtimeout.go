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

// ZoneSettingProxyReadTimeoutService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingProxyReadTimeoutService] method instead.
type ZoneSettingProxyReadTimeoutService struct {
	Options []option.RequestOption
}

// NewZoneSettingProxyReadTimeoutService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingProxyReadTimeoutService(opts ...option.RequestOption) (r *ZoneSettingProxyReadTimeoutService) {
	r = &ZoneSettingProxyReadTimeoutService{}
	r.Options = opts
	return
}

// Maximum time between two read operations from origin.
func (r *ZoneSettingProxyReadTimeoutService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingProxyReadTimeoutUpdateParams, opts ...option.RequestOption) (res *ZoneSettingProxyReadTimeoutUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Maximum time between two read operations from origin.
func (r *ZoneSettingProxyReadTimeoutService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingProxyReadTimeoutListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingProxyReadTimeoutUpdateResponse struct {
	Errors   []ZoneSettingProxyReadTimeoutUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingProxyReadTimeoutUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                            `json:"success,required"`
	Result  ZoneSettingProxyReadTimeoutUpdateResponseResult `json:"result"`
	JSON    zoneSettingProxyReadTimeoutUpdateResponseJSON   `json:"-"`
}

// zoneSettingProxyReadTimeoutUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingProxyReadTimeoutUpdateResponse]
type zoneSettingProxyReadTimeoutUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingProxyReadTimeoutUpdateResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingProxyReadTimeoutUpdateResponseError]
type zoneSettingProxyReadTimeoutUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingProxyReadTimeoutUpdateResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingProxyReadTimeoutUpdateResponseMessage]
type zoneSettingProxyReadTimeoutUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingProxyReadTimeoutUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingProxyReadTimeoutUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingProxyReadTimeoutUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingProxyReadTimeoutUpdateResponseResultJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingProxyReadTimeoutUpdateResponseResult]
type zoneSettingProxyReadTimeoutUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingProxyReadTimeoutUpdateResponseResultID string

const (
	ZoneSettingProxyReadTimeoutUpdateResponseResultIDProxyReadTimeout ZoneSettingProxyReadTimeoutUpdateResponseResultID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingProxyReadTimeoutUpdateResponseResultEditable bool

const (
	ZoneSettingProxyReadTimeoutUpdateResponseResultEditableTrue  ZoneSettingProxyReadTimeoutUpdateResponseResultEditable = true
	ZoneSettingProxyReadTimeoutUpdateResponseResultEditableFalse ZoneSettingProxyReadTimeoutUpdateResponseResultEditable = false
)

type ZoneSettingProxyReadTimeoutListResponse struct {
	Errors   []ZoneSettingProxyReadTimeoutListResponseError   `json:"errors,required"`
	Messages []ZoneSettingProxyReadTimeoutListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                          `json:"success,required"`
	Result  ZoneSettingProxyReadTimeoutListResponseResult `json:"result"`
	JSON    zoneSettingProxyReadTimeoutListResponseJSON   `json:"-"`
}

// zoneSettingProxyReadTimeoutListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingProxyReadTimeoutListResponse]
type zoneSettingProxyReadTimeoutListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingProxyReadTimeoutListResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutListResponseErrorJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingProxyReadTimeoutListResponseError]
type zoneSettingProxyReadTimeoutListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingProxyReadTimeoutListResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutListResponseMessageJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingProxyReadTimeoutListResponseMessage]
type zoneSettingProxyReadTimeoutListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingProxyReadTimeoutListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingProxyReadTimeoutListResponseResultID `json:"id,required"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingProxyReadTimeoutListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingProxyReadTimeoutListResponseResultJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingProxyReadTimeoutListResponseResult]
type zoneSettingProxyReadTimeoutListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingProxyReadTimeoutListResponseResultID string

const (
	ZoneSettingProxyReadTimeoutListResponseResultIDProxyReadTimeout ZoneSettingProxyReadTimeoutListResponseResultID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingProxyReadTimeoutListResponseResultEditable bool

const (
	ZoneSettingProxyReadTimeoutListResponseResultEditableTrue  ZoneSettingProxyReadTimeoutListResponseResultEditable = true
	ZoneSettingProxyReadTimeoutListResponseResultEditableFalse ZoneSettingProxyReadTimeoutListResponseResultEditable = false
)

type ZoneSettingProxyReadTimeoutUpdateParams struct {
	// Maximum time between two read operations from origin.
	Value param.Field[ZoneSettingProxyReadTimeoutUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingProxyReadTimeoutUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Maximum time between two read operations from origin.
type ZoneSettingProxyReadTimeoutUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingProxyReadTimeoutUpdateParamsValueID] `json:"id,required"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value param.Field[float64] `json:"value,required"`
}

func (r ZoneSettingProxyReadTimeoutUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingProxyReadTimeoutUpdateParamsValueID string

const (
	ZoneSettingProxyReadTimeoutUpdateParamsValueIDProxyReadTimeout ZoneSettingProxyReadTimeoutUpdateParamsValueID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingProxyReadTimeoutUpdateParamsValueEditable bool

const (
	ZoneSettingProxyReadTimeoutUpdateParamsValueEditableTrue  ZoneSettingProxyReadTimeoutUpdateParamsValueEditable = true
	ZoneSettingProxyReadTimeoutUpdateParamsValueEditableFalse ZoneSettingProxyReadTimeoutUpdateParamsValueEditable = false
)
