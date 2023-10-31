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

// ZoneSettingBrowserCheckService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingBrowserCheckService] method instead.
type ZoneSettingBrowserCheckService struct {
	Options []option.RequestOption
}

// NewZoneSettingBrowserCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingBrowserCheckService(opts ...option.RequestOption) (r *ZoneSettingBrowserCheckService) {
	r = &ZoneSettingBrowserCheckService{}
	r.Options = opts
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
func (r *ZoneSettingBrowserCheckService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingBrowserCheckUpdateParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCheckUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_check", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
func (r *ZoneSettingBrowserCheckService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingBrowserCheckListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_check", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingBrowserCheckUpdateResponse struct {
	Errors   []ZoneSettingBrowserCheckUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingBrowserCheckUpdateResponseMessage `json:"messages"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZoneSettingBrowserCheckUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingBrowserCheckUpdateResponseJSON
}

// zoneSettingBrowserCheckUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCheckUpdateResponse]
type zoneSettingBrowserCheckUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingBrowserCheckUpdateResponseErrorJSON
}

// zoneSettingBrowserCheckUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCheckUpdateResponseError]
type zoneSettingBrowserCheckUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingBrowserCheckUpdateResponseMessageJSON
}

// zoneSettingBrowserCheckUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCheckUpdateResponseMessage]
type zoneSettingBrowserCheckUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingBrowserCheckUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCheckUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCheckUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingBrowserCheckUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingBrowserCheckUpdateResponseResultJSON
}

// zoneSettingBrowserCheckUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCheckUpdateResponseResult]
type zoneSettingBrowserCheckUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCheckUpdateResponseResultID string

const (
	ZoneSettingBrowserCheckUpdateResponseResultIDBrowserCheck ZoneSettingBrowserCheckUpdateResponseResultID = "browser_check"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCheckUpdateResponseResultEditable bool

const (
	ZoneSettingBrowserCheckUpdateResponseResultEditableTrue  ZoneSettingBrowserCheckUpdateResponseResultEditable = true
	ZoneSettingBrowserCheckUpdateResponseResultEditableFalse ZoneSettingBrowserCheckUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingBrowserCheckUpdateResponseResultValue string

const (
	ZoneSettingBrowserCheckUpdateResponseResultValueOn  ZoneSettingBrowserCheckUpdateResponseResultValue = "on"
	ZoneSettingBrowserCheckUpdateResponseResultValueOff ZoneSettingBrowserCheckUpdateResponseResultValue = "off"
)

type ZoneSettingBrowserCheckListResponse struct {
	Errors   []ZoneSettingBrowserCheckListResponseError   `json:"errors"`
	Messages []ZoneSettingBrowserCheckListResponseMessage `json:"messages"`
	// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
	// headers abused most commonly by spammers and denies access to your page. It will
	// also challenge visitors that do not have a user agent or a non standard user
	// agent (also commonly used by abuse bots, crawlers or visitors).
	// (https://support.cloudflare.com/hc/en-us/articles/200170086).
	Result ZoneSettingBrowserCheckListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingBrowserCheckListResponseJSON
}

// zoneSettingBrowserCheckListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCheckListResponse]
type zoneSettingBrowserCheckListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingBrowserCheckListResponseErrorJSON
}

// zoneSettingBrowserCheckListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCheckListResponseError]
type zoneSettingBrowserCheckListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCheckListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingBrowserCheckListResponseMessageJSON
}

// zoneSettingBrowserCheckListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCheckListResponseMessage]
type zoneSettingBrowserCheckListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingBrowserCheckListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCheckListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCheckListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingBrowserCheckListResponseResultValue `json:"value"`
	JSON  zoneSettingBrowserCheckListResponseResultJSON
}

// zoneSettingBrowserCheckListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCheckListResponseResult]
type zoneSettingBrowserCheckListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCheckListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCheckListResponseResultID string

const (
	ZoneSettingBrowserCheckListResponseResultIDBrowserCheck ZoneSettingBrowserCheckListResponseResultID = "browser_check"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCheckListResponseResultEditable bool

const (
	ZoneSettingBrowserCheckListResponseResultEditableTrue  ZoneSettingBrowserCheckListResponseResultEditable = true
	ZoneSettingBrowserCheckListResponseResultEditableFalse ZoneSettingBrowserCheckListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingBrowserCheckListResponseResultValue string

const (
	ZoneSettingBrowserCheckListResponseResultValueOn  ZoneSettingBrowserCheckListResponseResultValue = "on"
	ZoneSettingBrowserCheckListResponseResultValueOff ZoneSettingBrowserCheckListResponseResultValue = "off"
)

type ZoneSettingBrowserCheckUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingBrowserCheckUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingBrowserCheckUpdateParamsValue string

const (
	ZoneSettingBrowserCheckUpdateParamsValueOn  ZoneSettingBrowserCheckUpdateParamsValue = "on"
	ZoneSettingBrowserCheckUpdateParamsValueOff ZoneSettingBrowserCheckUpdateParamsValue = "off"
)
