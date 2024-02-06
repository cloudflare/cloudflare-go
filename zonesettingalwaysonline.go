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
func (r *ZoneSettingAlwaysOnlineService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingAlwaysOnlineGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_online", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

type ZoneSettingAlwaysOnlineGetResponse struct {
	Errors   []ZoneSettingAlwaysOnlineGetResponseError   `json:"errors"`
	Messages []ZoneSettingAlwaysOnlineGetResponseMessage `json:"messages"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZoneSettingAlwaysOnlineGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                   `json:"success"`
	JSON    zoneSettingAlwaysOnlineGetResponseJSON `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAlwaysOnlineGetResponse]
type zoneSettingAlwaysOnlineGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineGetResponseErrorJSON `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineGetResponseError]
type zoneSettingAlwaysOnlineGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysOnlineGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineGetResponseMessageJSON `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineGetResponseMessage]
type zoneSettingAlwaysOnlineGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingAlwaysOnlineGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAlwaysOnlineGetResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAlwaysOnlineGetResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingAlwaysOnlineGetResponseResultValue `json:"value"`
	JSON  zoneSettingAlwaysOnlineGetResponseResultJSON  `json:"-"`
}

// zoneSettingAlwaysOnlineGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingAlwaysOnlineGetResponseResult]
type zoneSettingAlwaysOnlineGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysOnlineGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAlwaysOnlineGetResponseResultID string

const (
	ZoneSettingAlwaysOnlineGetResponseResultIDAlwaysOnline ZoneSettingAlwaysOnlineGetResponseResultID = "always_online"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAlwaysOnlineGetResponseResultEditable bool

const (
	ZoneSettingAlwaysOnlineGetResponseResultEditableTrue  ZoneSettingAlwaysOnlineGetResponseResultEditable = true
	ZoneSettingAlwaysOnlineGetResponseResultEditableFalse ZoneSettingAlwaysOnlineGetResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingAlwaysOnlineGetResponseResultValue string

const (
	ZoneSettingAlwaysOnlineGetResponseResultValueOn  ZoneSettingAlwaysOnlineGetResponseResultValue = "on"
	ZoneSettingAlwaysOnlineGetResponseResultValueOff ZoneSettingAlwaysOnlineGetResponseResultValue = "off"
)

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
	Success bool                                      `json:"success"`
	JSON    zoneSettingAlwaysOnlineUpdateResponseJSON `json:"-"`
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
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineUpdateResponseErrorJSON `json:"-"`
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
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingAlwaysOnlineUpdateResponseMessageJSON `json:"-"`
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
	JSON  zoneSettingAlwaysOnlineUpdateResponseResultJSON  `json:"-"`
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
