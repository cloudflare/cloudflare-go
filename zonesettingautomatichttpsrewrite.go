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

// ZoneSettingAutomaticHTTPsRewriteService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingAutomaticHTTPsRewriteService] method instead.
type ZoneSettingAutomaticHTTPsRewriteService struct {
	Options []option.RequestOption
}

// NewZoneSettingAutomaticHTTPsRewriteService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingAutomaticHTTPsRewriteService(opts ...option.RequestOption) (r *ZoneSettingAutomaticHTTPsRewriteService) {
	r = &ZoneSettingAutomaticHTTPsRewriteService{}
	r.Options = opts
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *ZoneSettingAutomaticHTTPsRewriteService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingAutomaticHTTPsRewriteUpdateParams, opts ...option.RequestOption) (res *ZoneSettingAutomaticHTTPsRewriteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *ZoneSettingAutomaticHTTPsRewriteService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingAutomaticHTTPsRewriteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingAutomaticHTTPsRewriteUpdateResponse struct {
	Errors   []ZoneSettingAutomaticHTTPsRewriteUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingAutomaticHTTPsRewriteUpdateResponseMessage `json:"messages"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result ZoneSettingAutomaticHTTPsRewriteUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingAutomaticHTTPsRewriteUpdateResponseJSON
}

// zoneSettingAutomaticHTTPsRewriteUpdateResponseJSON contains the JSON metadata
// for the struct [ZoneSettingAutomaticHTTPsRewriteUpdateResponse]
type zoneSettingAutomaticHTTPsRewriteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPsRewriteUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPsRewriteUpdateResponseErrorJSON
}

// zoneSettingAutomaticHTTPsRewriteUpdateResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticHTTPsRewriteUpdateResponseError]
type zoneSettingAutomaticHTTPsRewriteUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPsRewriteUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPsRewriteUpdateResponseMessageJSON
}

// zoneSettingAutomaticHTTPsRewriteUpdateResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticHTTPsRewriteUpdateResponseMessage]
type zoneSettingAutomaticHTTPsRewriteUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingAutomaticHTTPsRewriteUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingAutomaticHTTPsRewriteUpdateResponseResultJSON
}

// zoneSettingAutomaticHTTPsRewriteUpdateResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticHTTPsRewriteUpdateResponseResult]
type zoneSettingAutomaticHTTPsRewriteUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultID string

const (
	ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultIDAutomaticHTTPsRewrites ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultID = "automatic_https_rewrites"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultEditable bool

const (
	ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultEditableTrue  ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultEditable = true
	ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultEditableFalse ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultValue string

const (
	ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultValueOn  ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultValue = "on"
	ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultValueOff ZoneSettingAutomaticHTTPsRewriteUpdateResponseResultValue = "off"
)

type ZoneSettingAutomaticHTTPsRewriteListResponse struct {
	Errors   []ZoneSettingAutomaticHTTPsRewriteListResponseError   `json:"errors"`
	Messages []ZoneSettingAutomaticHTTPsRewriteListResponseMessage `json:"messages"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result ZoneSettingAutomaticHTTPsRewriteListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingAutomaticHTTPsRewriteListResponseJSON
}

// zoneSettingAutomaticHTTPsRewriteListResponseJSON contains the JSON metadata for
// the struct [ZoneSettingAutomaticHTTPsRewriteListResponse]
type zoneSettingAutomaticHTTPsRewriteListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPsRewriteListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPsRewriteListResponseErrorJSON
}

// zoneSettingAutomaticHTTPsRewriteListResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingAutomaticHTTPsRewriteListResponseError]
type zoneSettingAutomaticHTTPsRewriteListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAutomaticHTTPsRewriteListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAutomaticHTTPsRewriteListResponseMessageJSON
}

// zoneSettingAutomaticHTTPsRewriteListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticHTTPsRewriteListResponseMessage]
type zoneSettingAutomaticHTTPsRewriteListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingAutomaticHTTPsRewriteListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAutomaticHTTPsRewriteListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAutomaticHTTPsRewriteListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingAutomaticHTTPsRewriteListResponseResultValue `json:"value"`
	JSON  zoneSettingAutomaticHTTPsRewriteListResponseResultJSON
}

// zoneSettingAutomaticHTTPsRewriteListResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingAutomaticHTTPsRewriteListResponseResult]
type zoneSettingAutomaticHTTPsRewriteListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAutomaticHTTPsRewriteListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAutomaticHTTPsRewriteListResponseResultID string

const (
	ZoneSettingAutomaticHTTPsRewriteListResponseResultIDAutomaticHTTPsRewrites ZoneSettingAutomaticHTTPsRewriteListResponseResultID = "automatic_https_rewrites"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAutomaticHTTPsRewriteListResponseResultEditable bool

const (
	ZoneSettingAutomaticHTTPsRewriteListResponseResultEditableTrue  ZoneSettingAutomaticHTTPsRewriteListResponseResultEditable = true
	ZoneSettingAutomaticHTTPsRewriteListResponseResultEditableFalse ZoneSettingAutomaticHTTPsRewriteListResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingAutomaticHTTPsRewriteListResponseResultValue string

const (
	ZoneSettingAutomaticHTTPsRewriteListResponseResultValueOn  ZoneSettingAutomaticHTTPsRewriteListResponseResultValue = "on"
	ZoneSettingAutomaticHTTPsRewriteListResponseResultValueOff ZoneSettingAutomaticHTTPsRewriteListResponseResultValue = "off"
)

type ZoneSettingAutomaticHTTPsRewriteUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingAutomaticHTTPsRewriteUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingAutomaticHTTPsRewriteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingAutomaticHTTPsRewriteUpdateParamsValue string

const (
	ZoneSettingAutomaticHTTPsRewriteUpdateParamsValueOn  ZoneSettingAutomaticHTTPsRewriteUpdateParamsValue = "on"
	ZoneSettingAutomaticHTTPsRewriteUpdateParamsValueOff ZoneSettingAutomaticHTTPsRewriteUpdateParamsValue = "off"
)
