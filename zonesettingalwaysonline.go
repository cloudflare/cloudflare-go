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

// ZoneSettingAlwaysOnlineService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingAlwaysOnlineService] method instead.
type ZoneSettingAlwaysOnlineService struct {
	Options []option.RequestOption
}

// NewZoneSettingAlwaysOnlineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingAlwaysOnlineService(opts ...option.RequestOption) (r *ZoneSettingAlwaysOnlineService) {
	r = &ZoneSettingAlwaysOnlineService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *ZoneSettingAlwaysOnlineService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingAlwaysOnlineUpdateParams, opts ...option.RequestOption) (res *ZoneSettingAlwaysOnlineUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_online", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *ZoneSettingAlwaysOnlineService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingAlwaysOnlineListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_online", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingAlwaysOnlineUpdateResponse struct {
	Errors   []ZoneSettingAlwaysOnlineUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingAlwaysOnlineUpdateResponseMessage `json:"messages"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZoneSettingAlwaysOnlineUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingAlwaysOnlineUpdateResponseJSON
}

// zoneSettingAlwaysOnlineUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineUpdateResponse]
type zoneSettingAlwaysOnlineUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineUpdateResponseErrorJSON
}

// zoneSettingAlwaysOnlineUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysOnlineUpdateResponseError]
type zoneSettingAlwaysOnlineUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineUpdateResponseMessageJSON
}

// zoneSettingAlwaysOnlineUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysOnlineUpdateResponseMessage]
type zoneSettingAlwaysOnlineUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingAlwaysOnlineUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysOnlineUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysOnlineUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingAlwaysOnlineUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingAlwaysOnlineUpdateResponseResultJSON
}

// zoneSettingAlwaysOnlineUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysOnlineUpdateResponseResult]
type zoneSettingAlwaysOnlineUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAlwaysOnlineUpdateResponseResultID string

const (
	ZoneSettingAlwaysOnlineUpdateResponseResultIDAlwaysOnline ZoneSettingAlwaysOnlineUpdateResponseResultID = "always_online"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysOnlineUpdateResponseResultEditable bool

const (
	ZoneSettingAlwaysOnlineUpdateResponseResultEditableTrue  ZoneSettingAlwaysOnlineUpdateResponseResultEditable = true
	ZoneSettingAlwaysOnlineUpdateResponseResultEditableFalse ZoneSettingAlwaysOnlineUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingAlwaysOnlineUpdateResponseResultValue string

const (
	ZoneSettingAlwaysOnlineUpdateResponseResultValueOn  ZoneSettingAlwaysOnlineUpdateResponseResultValue = "on"
	ZoneSettingAlwaysOnlineUpdateResponseResultValueOff ZoneSettingAlwaysOnlineUpdateResponseResultValue = "off"
)

type ZoneSettingAlwaysOnlineListResponse struct {
	Errors   []ZoneSettingAlwaysOnlineListResponseError   `json:"errors"`
	Messages []ZoneSettingAlwaysOnlineListResponseMessage `json:"messages"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZoneSettingAlwaysOnlineListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingAlwaysOnlineListResponseJSON
}

// zoneSettingAlwaysOnlineListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineListResponse]
type zoneSettingAlwaysOnlineListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineListResponseErrorJSON
}

// zoneSettingAlwaysOnlineListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineListResponseError]
type zoneSettingAlwaysOnlineListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineListResponseMessageJSON
}

// zoneSettingAlwaysOnlineListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysOnlineListResponseMessage]
type zoneSettingAlwaysOnlineListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingAlwaysOnlineListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysOnlineListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysOnlineListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingAlwaysOnlineListResponseResultValue `json:"value"`
	JSON  zoneSettingAlwaysOnlineListResponseResultJSON
}

// zoneSettingAlwaysOnlineListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineListResponseResult]
type zoneSettingAlwaysOnlineListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAlwaysOnlineListResponseResultID string

const (
	ZoneSettingAlwaysOnlineListResponseResultIDAlwaysOnline ZoneSettingAlwaysOnlineListResponseResultID = "always_online"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysOnlineListResponseResultEditable bool

const (
	ZoneSettingAlwaysOnlineListResponseResultEditableTrue  ZoneSettingAlwaysOnlineListResponseResultEditable = true
	ZoneSettingAlwaysOnlineListResponseResultEditableFalse ZoneSettingAlwaysOnlineListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingAlwaysOnlineListResponseResultValue string

const (
	ZoneSettingAlwaysOnlineListResponseResultValueOn  ZoneSettingAlwaysOnlineListResponseResultValue = "on"
	ZoneSettingAlwaysOnlineListResponseResultValueOff ZoneSettingAlwaysOnlineListResponseResultValue = "off"
)

type ZoneSettingAlwaysOnlineUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingAlwaysOnlineUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingAlwaysOnlineUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingAlwaysOnlineUpdateParamsValue string

const (
	ZoneSettingAlwaysOnlineUpdateParamsValueOn  ZoneSettingAlwaysOnlineUpdateParamsValue = "on"
	ZoneSettingAlwaysOnlineUpdateParamsValueOff ZoneSettingAlwaysOnlineUpdateParamsValue = "off"
)
