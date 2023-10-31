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

// ZoneSettingSortQueryStringForCachService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingSortQueryStringForCachService] method instead.
type ZoneSettingSortQueryStringForCachService struct {
	Options []option.RequestOption
}

// NewZoneSettingSortQueryStringForCachService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSortQueryStringForCachService(opts ...option.RequestOption) (r *ZoneSettingSortQueryStringForCachService) {
	r = &ZoneSettingSortQueryStringForCachService{}
	r.Options = opts
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *ZoneSettingSortQueryStringForCachService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingSortQueryStringForCachUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSortQueryStringForCachUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *ZoneSettingSortQueryStringForCachService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingSortQueryStringForCachListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingSortQueryStringForCachUpdateResponse struct {
	Errors   []ZoneSettingSortQueryStringForCachUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingSortQueryStringForCachUpdateResponseMessage `json:"messages"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result ZoneSettingSortQueryStringForCachUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingSortQueryStringForCachUpdateResponseJSON
}

// zoneSettingSortQueryStringForCachUpdateResponseJSON contains the JSON metadata
// for the struct [ZoneSettingSortQueryStringForCachUpdateResponse]
type zoneSettingSortQueryStringForCachUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSortQueryStringForCachUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCachUpdateResponseErrorJSON
}

// zoneSettingSortQueryStringForCachUpdateResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCachUpdateResponseError]
type zoneSettingSortQueryStringForCachUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSortQueryStringForCachUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCachUpdateResponseMessageJSON
}

// zoneSettingSortQueryStringForCachUpdateResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCachUpdateResponseMessage]
type zoneSettingSortQueryStringForCachUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingSortQueryStringForCachUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSortQueryStringForCachUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSortQueryStringForCachUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingSortQueryStringForCachUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingSortQueryStringForCachUpdateResponseResultJSON
}

// zoneSettingSortQueryStringForCachUpdateResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCachUpdateResponseResult]
type zoneSettingSortQueryStringForCachUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSortQueryStringForCachUpdateResponseResultID string

const (
	ZoneSettingSortQueryStringForCachUpdateResponseResultIDSortQueryStringForCache ZoneSettingSortQueryStringForCachUpdateResponseResultID = "sort_query_string_for_cache"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSortQueryStringForCachUpdateResponseResultEditable bool

const (
	ZoneSettingSortQueryStringForCachUpdateResponseResultEditableTrue  ZoneSettingSortQueryStringForCachUpdateResponseResultEditable = true
	ZoneSettingSortQueryStringForCachUpdateResponseResultEditableFalse ZoneSettingSortQueryStringForCachUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingSortQueryStringForCachUpdateResponseResultValue string

const (
	ZoneSettingSortQueryStringForCachUpdateResponseResultValueOn  ZoneSettingSortQueryStringForCachUpdateResponseResultValue = "on"
	ZoneSettingSortQueryStringForCachUpdateResponseResultValueOff ZoneSettingSortQueryStringForCachUpdateResponseResultValue = "off"
)

type ZoneSettingSortQueryStringForCachListResponse struct {
	Errors   []ZoneSettingSortQueryStringForCachListResponseError   `json:"errors"`
	Messages []ZoneSettingSortQueryStringForCachListResponseMessage `json:"messages"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result ZoneSettingSortQueryStringForCachListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingSortQueryStringForCachListResponseJSON
}

// zoneSettingSortQueryStringForCachListResponseJSON contains the JSON metadata for
// the struct [ZoneSettingSortQueryStringForCachListResponse]
type zoneSettingSortQueryStringForCachListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSortQueryStringForCachListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCachListResponseErrorJSON
}

// zoneSettingSortQueryStringForCachListResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCachListResponseError]
type zoneSettingSortQueryStringForCachListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSortQueryStringForCachListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSortQueryStringForCachListResponseMessageJSON
}

// zoneSettingSortQueryStringForCachListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCachListResponseMessage]
type zoneSettingSortQueryStringForCachListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingSortQueryStringForCachListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingSortQueryStringForCachListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSortQueryStringForCachListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingSortQueryStringForCachListResponseResultValue `json:"value"`
	JSON  zoneSettingSortQueryStringForCachListResponseResultJSON
}

// zoneSettingSortQueryStringForCachListResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingSortQueryStringForCachListResponseResult]
type zoneSettingSortQueryStringForCachListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSortQueryStringForCachListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingSortQueryStringForCachListResponseResultID string

const (
	ZoneSettingSortQueryStringForCachListResponseResultIDSortQueryStringForCache ZoneSettingSortQueryStringForCachListResponseResultID = "sort_query_string_for_cache"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSortQueryStringForCachListResponseResultEditable bool

const (
	ZoneSettingSortQueryStringForCachListResponseResultEditableTrue  ZoneSettingSortQueryStringForCachListResponseResultEditable = true
	ZoneSettingSortQueryStringForCachListResponseResultEditableFalse ZoneSettingSortQueryStringForCachListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingSortQueryStringForCachListResponseResultValue string

const (
	ZoneSettingSortQueryStringForCachListResponseResultValueOn  ZoneSettingSortQueryStringForCachListResponseResultValue = "on"
	ZoneSettingSortQueryStringForCachListResponseResultValueOff ZoneSettingSortQueryStringForCachListResponseResultValue = "off"
)

type ZoneSettingSortQueryStringForCachUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingSortQueryStringForCachUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSortQueryStringForCachUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingSortQueryStringForCachUpdateParamsValue string

const (
	ZoneSettingSortQueryStringForCachUpdateParamsValueOn  ZoneSettingSortQueryStringForCachUpdateParamsValue = "on"
	ZoneSettingSortQueryStringForCachUpdateParamsValueOff ZoneSettingSortQueryStringForCachUpdateParamsValue = "off"
)
