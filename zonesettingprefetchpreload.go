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

// ZoneSettingPrefetchPreloadService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingPrefetchPreloadService] method instead.
type ZoneSettingPrefetchPreloadService struct {
	Options []option.RequestOption
}

// NewZoneSettingPrefetchPreloadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingPrefetchPreloadService(opts ...option.RequestOption) (r *ZoneSettingPrefetchPreloadService) {
	r = &ZoneSettingPrefetchPreloadService{}
	r.Options = opts
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *ZoneSettingPrefetchPreloadService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingPrefetchPreloadUpdateParams, opts ...option.RequestOption) (res *ZoneSettingPrefetchPreloadUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *ZoneSettingPrefetchPreloadService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingPrefetchPreloadListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingPrefetchPreloadUpdateResponse struct {
	Errors   []ZoneSettingPrefetchPreloadUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingPrefetchPreloadUpdateResponseMessage `json:"messages"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZoneSettingPrefetchPreloadUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingPrefetchPreloadUpdateResponseJSON
}

// zoneSettingPrefetchPreloadUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingPrefetchPreloadUpdateResponse]
type zoneSettingPrefetchPreloadUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadUpdateResponseErrorJSON
}

// zoneSettingPrefetchPreloadUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingPrefetchPreloadUpdateResponseError]
type zoneSettingPrefetchPreloadUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadUpdateResponseMessageJSON
}

// zoneSettingPrefetchPreloadUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingPrefetchPreloadUpdateResponseMessage]
type zoneSettingPrefetchPreloadUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingPrefetchPreloadUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingPrefetchPreloadUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPrefetchPreloadUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingPrefetchPreloadUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingPrefetchPreloadUpdateResponseResultJSON
}

// zoneSettingPrefetchPreloadUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingPrefetchPreloadUpdateResponseResult]
type zoneSettingPrefetchPreloadUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPrefetchPreloadUpdateResponseResultID string

const (
	ZoneSettingPrefetchPreloadUpdateResponseResultIDPrefetchPreload ZoneSettingPrefetchPreloadUpdateResponseResultID = "prefetch_preload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPrefetchPreloadUpdateResponseResultEditable bool

const (
	ZoneSettingPrefetchPreloadUpdateResponseResultEditableTrue  ZoneSettingPrefetchPreloadUpdateResponseResultEditable = true
	ZoneSettingPrefetchPreloadUpdateResponseResultEditableFalse ZoneSettingPrefetchPreloadUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingPrefetchPreloadUpdateResponseResultValue string

const (
	ZoneSettingPrefetchPreloadUpdateResponseResultValueOn  ZoneSettingPrefetchPreloadUpdateResponseResultValue = "on"
	ZoneSettingPrefetchPreloadUpdateResponseResultValueOff ZoneSettingPrefetchPreloadUpdateResponseResultValue = "off"
)

type ZoneSettingPrefetchPreloadListResponse struct {
	Errors   []ZoneSettingPrefetchPreloadListResponseError   `json:"errors"`
	Messages []ZoneSettingPrefetchPreloadListResponseMessage `json:"messages"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result ZoneSettingPrefetchPreloadListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingPrefetchPreloadListResponseJSON
}

// zoneSettingPrefetchPreloadListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingPrefetchPreloadListResponse]
type zoneSettingPrefetchPreloadListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadListResponseErrorJSON
}

// zoneSettingPrefetchPreloadListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingPrefetchPreloadListResponseError]
type zoneSettingPrefetchPreloadListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrefetchPreloadListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrefetchPreloadListResponseMessageJSON
}

// zoneSettingPrefetchPreloadListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingPrefetchPreloadListResponseMessage]
type zoneSettingPrefetchPreloadListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingPrefetchPreloadListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingPrefetchPreloadListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPrefetchPreloadListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingPrefetchPreloadListResponseResultValue `json:"value"`
	JSON  zoneSettingPrefetchPreloadListResponseResultJSON
}

// zoneSettingPrefetchPreloadListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingPrefetchPreloadListResponseResult]
type zoneSettingPrefetchPreloadListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrefetchPreloadListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPrefetchPreloadListResponseResultID string

const (
	ZoneSettingPrefetchPreloadListResponseResultIDPrefetchPreload ZoneSettingPrefetchPreloadListResponseResultID = "prefetch_preload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPrefetchPreloadListResponseResultEditable bool

const (
	ZoneSettingPrefetchPreloadListResponseResultEditableTrue  ZoneSettingPrefetchPreloadListResponseResultEditable = true
	ZoneSettingPrefetchPreloadListResponseResultEditableFalse ZoneSettingPrefetchPreloadListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingPrefetchPreloadListResponseResultValue string

const (
	ZoneSettingPrefetchPreloadListResponseResultValueOn  ZoneSettingPrefetchPreloadListResponseResultValue = "on"
	ZoneSettingPrefetchPreloadListResponseResultValueOff ZoneSettingPrefetchPreloadListResponseResultValue = "off"
)

type ZoneSettingPrefetchPreloadUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingPrefetchPreloadUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingPrefetchPreloadUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingPrefetchPreloadUpdateParamsValue string

const (
	ZoneSettingPrefetchPreloadUpdateParamsValueOn  ZoneSettingPrefetchPreloadUpdateParamsValue = "on"
	ZoneSettingPrefetchPreloadUpdateParamsValueOff ZoneSettingPrefetchPreloadUpdateParamsValue = "off"
)
