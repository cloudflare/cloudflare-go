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

// ZoneSettingAlwaysUseHTTPService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingAlwaysUseHTTPService] method instead.
type ZoneSettingAlwaysUseHTTPService struct {
	Options []option.RequestOption
}

// NewZoneSettingAlwaysUseHTTPService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingAlwaysUseHTTPService(opts ...option.RequestOption) (r *ZoneSettingAlwaysUseHTTPService) {
	r = &ZoneSettingAlwaysUseHTTPService{}
	r.Options = opts
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *ZoneSettingAlwaysUseHTTPService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingAlwaysUseHTTPUpdateParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysUseHTTPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_use_https", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *ZoneSettingAlwaysUseHTTPService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingAlwaysUseHTTPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_use_https", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingAlwaysUseHTTPUpdateResponse struct {
	Errors   []ZoneSettingAlwaysUseHTTPUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingAlwaysUseHTTPUpdateResponseMessage `json:"messages"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result ZoneSettingAlwaysUseHTTPUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                       `json:"success"`
	JSON    zoneSettingAlwaysUseHTTPUpdateResponseJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysUseHTTPUpdateResponse]
type zoneSettingAlwaysUseHTTPUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysUseHTTPUpdateResponseError]
type zoneSettingAlwaysUseHTTPUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysUseHTTPUpdateResponseMessage]
type zoneSettingAlwaysUseHTTPUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingAlwaysUseHTTPUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysUseHTTPUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysUseHTTPUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingAlwaysUseHTTPUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingAlwaysUseHTTPUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingAlwaysUseHTTPUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysUseHTTPUpdateResponseResult]
type zoneSettingAlwaysUseHTTPUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAlwaysUseHTTPUpdateResponseResultID string

const (
	ZoneSettingAlwaysUseHTTPUpdateResponseResultIDAlwaysUseHTTPs ZoneSettingAlwaysUseHTTPUpdateResponseResultID = "always_use_https"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysUseHTTPUpdateResponseResultEditable bool

const (
	ZoneSettingAlwaysUseHTTPUpdateResponseResultEditableTrue  ZoneSettingAlwaysUseHTTPUpdateResponseResultEditable = true
	ZoneSettingAlwaysUseHTTPUpdateResponseResultEditableFalse ZoneSettingAlwaysUseHTTPUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingAlwaysUseHTTPUpdateResponseResultValue string

const (
	ZoneSettingAlwaysUseHTTPUpdateResponseResultValueOn  ZoneSettingAlwaysUseHTTPUpdateResponseResultValue = "on"
	ZoneSettingAlwaysUseHTTPUpdateResponseResultValueOff ZoneSettingAlwaysUseHTTPUpdateResponseResultValue = "off"
)

type ZoneSettingAlwaysUseHTTPListResponse struct {
	Errors   []ZoneSettingAlwaysUseHTTPListResponseError   `json:"errors"`
	Messages []ZoneSettingAlwaysUseHTTPListResponseMessage `json:"messages"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result ZoneSettingAlwaysUseHTTPListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                     `json:"success"`
	JSON    zoneSettingAlwaysUseHTTPListResponseJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysUseHTTPListResponse]
type zoneSettingAlwaysUseHTTPListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPListResponseErrorJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysUseHTTPListResponseError]
type zoneSettingAlwaysUseHTTPListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPListResponseMessageJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysUseHTTPListResponseMessage]
type zoneSettingAlwaysUseHTTPListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingAlwaysUseHTTPListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysUseHTTPListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysUseHTTPListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingAlwaysUseHTTPListResponseResultValue `json:"value"`
	JSON  zoneSettingAlwaysUseHTTPListResponseResultJSON  `json:"-"`
}

// zoneSettingAlwaysUseHTTPListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysUseHTTPListResponseResult]
type zoneSettingAlwaysUseHTTPListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAlwaysUseHTTPListResponseResultID string

const (
	ZoneSettingAlwaysUseHTTPListResponseResultIDAlwaysUseHTTPs ZoneSettingAlwaysUseHTTPListResponseResultID = "always_use_https"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysUseHTTPListResponseResultEditable bool

const (
	ZoneSettingAlwaysUseHTTPListResponseResultEditableTrue  ZoneSettingAlwaysUseHTTPListResponseResultEditable = true
	ZoneSettingAlwaysUseHTTPListResponseResultEditableFalse ZoneSettingAlwaysUseHTTPListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingAlwaysUseHTTPListResponseResultValue string

const (
	ZoneSettingAlwaysUseHTTPListResponseResultValueOn  ZoneSettingAlwaysUseHTTPListResponseResultValue = "on"
	ZoneSettingAlwaysUseHTTPListResponseResultValueOff ZoneSettingAlwaysUseHTTPListResponseResultValue = "off"
)

type ZoneSettingAlwaysUseHTTPUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingAlwaysUseHTTPUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingAlwaysUseHTTPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingAlwaysUseHTTPUpdateParamsValue string

const (
	ZoneSettingAlwaysUseHTTPUpdateParamsValueOn  ZoneSettingAlwaysUseHTTPUpdateParamsValue = "on"
	ZoneSettingAlwaysUseHTTPUpdateParamsValueOff ZoneSettingAlwaysUseHTTPUpdateParamsValue = "off"
)
