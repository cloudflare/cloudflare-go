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

// ZoneSettingWafService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingWafService] method
// instead.
type ZoneSettingWafService struct {
	Options []option.RequestOption
}

// NewZoneSettingWafService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingWafService(opts ...option.RequestOption) (r *ZoneSettingWafService) {
	r = &ZoneSettingWafService{}
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
func (r *ZoneSettingWafService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingWafUpdateParams, opts ...option.RequestOption) (res *ZoneSettingWafUpdateResponse, err error) {
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
func (r *ZoneSettingWafService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingWafListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/waf", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingWafUpdateResponse struct {
	Errors   []ZoneSettingWafUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingWafUpdateResponseMessage `json:"messages"`
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
	Result ZoneSettingWafUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingWafUpdateResponseJSON
}

// zoneSettingWafUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWafUpdateResponse]
type zoneSettingWafUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWafUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingWafUpdateResponseErrorJSON
}

// zoneSettingWafUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingWafUpdateResponseError]
type zoneSettingWafUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWafUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingWafUpdateResponseMessageJSON
}

// zoneSettingWafUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingWafUpdateResponseMessage]
type zoneSettingWafUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
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
type ZoneSettingWafUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWafUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWafUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingWafUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingWafUpdateResponseResultJSON
}

// zoneSettingWafUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingWafUpdateResponseResult]
type zoneSettingWafUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWafUpdateResponseResultID string

const (
	ZoneSettingWafUpdateResponseResultIDWaf ZoneSettingWafUpdateResponseResultID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWafUpdateResponseResultEditable bool

const (
	ZoneSettingWafUpdateResponseResultEditableTrue  ZoneSettingWafUpdateResponseResultEditable = true
	ZoneSettingWafUpdateResponseResultEditableFalse ZoneSettingWafUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingWafUpdateResponseResultValue string

const (
	ZoneSettingWafUpdateResponseResultValueOn  ZoneSettingWafUpdateResponseResultValue = "on"
	ZoneSettingWafUpdateResponseResultValueOff ZoneSettingWafUpdateResponseResultValue = "off"
)

type ZoneSettingWafListResponse struct {
	Errors   []ZoneSettingWafListResponseError   `json:"errors"`
	Messages []ZoneSettingWafListResponseMessage `json:"messages"`
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
	Result ZoneSettingWafListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingWafListResponseJSON
}

// zoneSettingWafListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWafListResponse]
type zoneSettingWafListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWafListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingWafListResponseErrorJSON
}

// zoneSettingWafListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingWafListResponseError]
type zoneSettingWafListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWafListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingWafListResponseMessageJSON
}

// zoneSettingWafListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingWafListResponseMessage]
type zoneSettingWafListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafListResponseMessage) UnmarshalJSON(data []byte) (err error) {
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
type ZoneSettingWafListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWafListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWafListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingWafListResponseResultValue `json:"value"`
	JSON  zoneSettingWafListResponseResultJSON
}

// zoneSettingWafListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingWafListResponseResult]
type zoneSettingWafListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWafListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWafListResponseResultID string

const (
	ZoneSettingWafListResponseResultIDWaf ZoneSettingWafListResponseResultID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWafListResponseResultEditable bool

const (
	ZoneSettingWafListResponseResultEditableTrue  ZoneSettingWafListResponseResultEditable = true
	ZoneSettingWafListResponseResultEditableFalse ZoneSettingWafListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingWafListResponseResultValue string

const (
	ZoneSettingWafListResponseResultValueOn  ZoneSettingWafListResponseResultValue = "on"
	ZoneSettingWafListResponseResultValueOff ZoneSettingWafListResponseResultValue = "off"
)

type ZoneSettingWafUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingWafUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingWafUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingWafUpdateParamsValue string

const (
	ZoneSettingWafUpdateParamsValueOn  ZoneSettingWafUpdateParamsValue = "on"
	ZoneSettingWafUpdateParamsValueOff ZoneSettingWafUpdateParamsValue = "off"
)
