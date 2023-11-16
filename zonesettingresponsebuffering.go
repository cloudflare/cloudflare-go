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

// ZoneSettingResponseBufferingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingResponseBufferingService] method instead.
type ZoneSettingResponseBufferingService struct {
	Options []option.RequestOption
}

// NewZoneSettingResponseBufferingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingResponseBufferingService(opts ...option.RequestOption) (r *ZoneSettingResponseBufferingService) {
	r = &ZoneSettingResponseBufferingService{}
	r.Options = opts
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *ZoneSettingResponseBufferingService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingResponseBufferingUpdateParams, opts ...option.RequestOption) (res *ZoneSettingResponseBufferingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/response_buffering", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *ZoneSettingResponseBufferingService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingResponseBufferingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/response_buffering", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingResponseBufferingUpdateResponse struct {
	Errors   []ZoneSettingResponseBufferingUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingResponseBufferingUpdateResponseMessage `json:"messages"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ZoneSettingResponseBufferingUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                           `json:"success"`
	JSON    zoneSettingResponseBufferingUpdateResponseJSON `json:"-"`
}

// zoneSettingResponseBufferingUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingResponseBufferingUpdateResponse]
type zoneSettingResponseBufferingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingResponseBufferingUpdateResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingResponseBufferingUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingResponseBufferingUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingResponseBufferingUpdateResponseError]
type zoneSettingResponseBufferingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingResponseBufferingUpdateResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingResponseBufferingUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingResponseBufferingUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingResponseBufferingUpdateResponseMessage]
type zoneSettingResponseBufferingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingResponseBufferingUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingResponseBufferingUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingResponseBufferingUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingResponseBufferingUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingResponseBufferingUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingResponseBufferingUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingResponseBufferingUpdateResponseResult]
type zoneSettingResponseBufferingUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingResponseBufferingUpdateResponseResultID string

const (
	ZoneSettingResponseBufferingUpdateResponseResultIDResponseBuffering ZoneSettingResponseBufferingUpdateResponseResultID = "response_buffering"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingResponseBufferingUpdateResponseResultEditable bool

const (
	ZoneSettingResponseBufferingUpdateResponseResultEditableTrue  ZoneSettingResponseBufferingUpdateResponseResultEditable = true
	ZoneSettingResponseBufferingUpdateResponseResultEditableFalse ZoneSettingResponseBufferingUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingResponseBufferingUpdateResponseResultValue string

const (
	ZoneSettingResponseBufferingUpdateResponseResultValueOn  ZoneSettingResponseBufferingUpdateResponseResultValue = "on"
	ZoneSettingResponseBufferingUpdateResponseResultValueOff ZoneSettingResponseBufferingUpdateResponseResultValue = "off"
)

type ZoneSettingResponseBufferingListResponse struct {
	Errors   []ZoneSettingResponseBufferingListResponseError   `json:"errors"`
	Messages []ZoneSettingResponseBufferingListResponseMessage `json:"messages"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ZoneSettingResponseBufferingListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                         `json:"success"`
	JSON    zoneSettingResponseBufferingListResponseJSON `json:"-"`
}

// zoneSettingResponseBufferingListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingResponseBufferingListResponse]
type zoneSettingResponseBufferingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingResponseBufferingListResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingResponseBufferingListResponseErrorJSON `json:"-"`
}

// zoneSettingResponseBufferingListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingResponseBufferingListResponseError]
type zoneSettingResponseBufferingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingResponseBufferingListResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingResponseBufferingListResponseMessageJSON `json:"-"`
}

// zoneSettingResponseBufferingListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingResponseBufferingListResponseMessage]
type zoneSettingResponseBufferingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingResponseBufferingListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingResponseBufferingListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingResponseBufferingListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingResponseBufferingListResponseResultValue `json:"value"`
	JSON  zoneSettingResponseBufferingListResponseResultJSON  `json:"-"`
}

// zoneSettingResponseBufferingListResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingResponseBufferingListResponseResult]
type zoneSettingResponseBufferingListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingResponseBufferingListResponseResultID string

const (
	ZoneSettingResponseBufferingListResponseResultIDResponseBuffering ZoneSettingResponseBufferingListResponseResultID = "response_buffering"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingResponseBufferingListResponseResultEditable bool

const (
	ZoneSettingResponseBufferingListResponseResultEditableTrue  ZoneSettingResponseBufferingListResponseResultEditable = true
	ZoneSettingResponseBufferingListResponseResultEditableFalse ZoneSettingResponseBufferingListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingResponseBufferingListResponseResultValue string

const (
	ZoneSettingResponseBufferingListResponseResultValueOn  ZoneSettingResponseBufferingListResponseResultValue = "on"
	ZoneSettingResponseBufferingListResponseResultValueOff ZoneSettingResponseBufferingListResponseResultValue = "off"
)

type ZoneSettingResponseBufferingUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingResponseBufferingUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingResponseBufferingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingResponseBufferingUpdateParamsValue string

const (
	ZoneSettingResponseBufferingUpdateParamsValueOn  ZoneSettingResponseBufferingUpdateParamsValue = "on"
	ZoneSettingResponseBufferingUpdateParamsValueOff ZoneSettingResponseBufferingUpdateParamsValue = "off"
)
