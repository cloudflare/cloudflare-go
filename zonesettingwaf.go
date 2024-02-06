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

// ZoneSettingWAFService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingWAFService] method
// instead.
type ZoneSettingWAFService struct {
	Options []option.RequestOption
}

// NewZoneSettingWAFService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingWAFService(opts ...option.RequestOption) (r *ZoneSettingWAFService) {
	r = &ZoneSettingWAFService{}
	r.Options = opts
	return
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
func (r *ZoneSettingWAFService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingWAFUpdateParams, opts ...option.RequestOption) (res *ZoneSettingWAFUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/waf", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
func (r *ZoneSettingWAFService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingWAFListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/waf", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingWAFUpdateResponse struct {
	Errors   []ZoneSettingWAFUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingWAFUpdateResponseMessage `json:"messages"`
	// The WAF examines HTTP requests to your website. It inspects both GET and POST
	// requests and applies rules to help filter out illegitimate traffic from
	// legitimate website visitors. The Cloudflare WAF inspects website addresses or
	// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
	// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
	// page that asks them to submit a CAPTCHA successfully to continue their action.
	// If the challenge is failed, the action will be stopped. What this means is that
	// Cloudflare's WAF will block any traffic identified as illegitimate before it
	// reaches your origin web server.
	// (https://support.cloudflare.com/hc/en-us/articles/200172016).
	Result ZoneSettingWAFUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                             `json:"success"`
	JSON    zoneSettingWAFUpdateResponseJSON `json:"-"`
}

// zoneSettingWAFUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWAFUpdateResponse]
type zoneSettingWAFUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWAFUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingWAFUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingWAFUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingWAFUpdateResponseError]
type zoneSettingWAFUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWAFUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSettingWAFUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingWAFUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingWAFUpdateResponseMessage]
type zoneSettingWAFUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type ZoneSettingWAFUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWAFUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWAFUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingWAFUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingWAFUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingWAFUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingWAFUpdateResponseResult]
type zoneSettingWAFUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWAFUpdateResponseResultID string

const (
	ZoneSettingWAFUpdateResponseResultIDWAF ZoneSettingWAFUpdateResponseResultID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWAFUpdateResponseResultEditable bool

const (
	ZoneSettingWAFUpdateResponseResultEditableTrue  ZoneSettingWAFUpdateResponseResultEditable = true
	ZoneSettingWAFUpdateResponseResultEditableFalse ZoneSettingWAFUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingWAFUpdateResponseResultValue string

const (
	ZoneSettingWAFUpdateResponseResultValueOn  ZoneSettingWAFUpdateResponseResultValue = "on"
	ZoneSettingWAFUpdateResponseResultValueOff ZoneSettingWAFUpdateResponseResultValue = "off"
)

type ZoneSettingWAFListResponse struct {
	Errors   []ZoneSettingWAFListResponseError   `json:"errors"`
	Messages []ZoneSettingWAFListResponseMessage `json:"messages"`
	// The WAF examines HTTP requests to your website. It inspects both GET and POST
	// requests and applies rules to help filter out illegitimate traffic from
	// legitimate website visitors. The Cloudflare WAF inspects website addresses or
	// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
	// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
	// page that asks them to submit a CAPTCHA successfully to continue their action.
	// If the challenge is failed, the action will be stopped. What this means is that
	// Cloudflare's WAF will block any traffic identified as illegitimate before it
	// reaches your origin web server.
	// (https://support.cloudflare.com/hc/en-us/articles/200172016).
	Result ZoneSettingWAFListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	JSON    zoneSettingWAFListResponseJSON `json:"-"`
}

// zoneSettingWAFListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWAFListResponse]
type zoneSettingWAFListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWAFListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneSettingWAFListResponseErrorJSON `json:"-"`
}

// zoneSettingWAFListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingWAFListResponseError]
type zoneSettingWAFListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWAFListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingWAFListResponseMessageJSON `json:"-"`
}

// zoneSettingWAFListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingWAFListResponseMessage]
type zoneSettingWAFListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type ZoneSettingWAFListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWAFListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWAFListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingWAFListResponseResultValue `json:"value"`
	JSON  zoneSettingWAFListResponseResultJSON  `json:"-"`
}

// zoneSettingWAFListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingWAFListResponseResult]
type zoneSettingWAFListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWAFListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWAFListResponseResultID string

const (
	ZoneSettingWAFListResponseResultIDWAF ZoneSettingWAFListResponseResultID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWAFListResponseResultEditable bool

const (
	ZoneSettingWAFListResponseResultEditableTrue  ZoneSettingWAFListResponseResultEditable = true
	ZoneSettingWAFListResponseResultEditableFalse ZoneSettingWAFListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingWAFListResponseResultValue string

const (
	ZoneSettingWAFListResponseResultValueOn  ZoneSettingWAFListResponseResultValue = "on"
	ZoneSettingWAFListResponseResultValueOff ZoneSettingWAFListResponseResultValue = "off"
)

type ZoneSettingWAFUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingWAFUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingWAFUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingWAFUpdateParamsValue string

const (
	ZoneSettingWAFUpdateParamsValueOn  ZoneSettingWAFUpdateParamsValue = "on"
	ZoneSettingWAFUpdateParamsValueOff ZoneSettingWAFUpdateParamsValue = "off"
)
