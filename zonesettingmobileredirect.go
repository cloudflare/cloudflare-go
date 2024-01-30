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

// ZoneSettingMobileRedirectService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingMobileRedirectService] method instead.
type ZoneSettingMobileRedirectService struct {
	Options []option.RequestOption
}

// NewZoneSettingMobileRedirectService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingMobileRedirectService(opts ...option.RequestOption) (r *ZoneSettingMobileRedirectService) {
	r = &ZoneSettingMobileRedirectService{}
	r.Options = opts
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *ZoneSettingMobileRedirectService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingMobileRedirectUpdateParams, opts ...option.RequestOption) (res *ZoneSettingMobileRedirectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *ZoneSettingMobileRedirectService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingMobileRedirectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingMobileRedirectUpdateResponse struct {
	Errors   []ZoneSettingMobileRedirectUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingMobileRedirectUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                          `json:"success,required"`
	Result  ZoneSettingMobileRedirectUpdateResponseResult `json:"result"`
	JSON    zoneSettingMobileRedirectUpdateResponseJSON   `json:"-"`
}

// zoneSettingMobileRedirectUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMobileRedirectUpdateResponse]
type zoneSettingMobileRedirectUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectUpdateResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingMobileRedirectUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingMobileRedirectUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectUpdateResponseError]
type zoneSettingMobileRedirectUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectUpdateResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingMobileRedirectUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingMobileRedirectUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingMobileRedirectUpdateResponseMessage]
type zoneSettingMobileRedirectUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectUpdateResponseResult struct {
	// Identifier of the zone setting.
	ID ZoneSettingMobileRedirectUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingMobileRedirectUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMobileRedirectUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMobileRedirectUpdateResponseResultJSON `json:"-"`
}

// zoneSettingMobileRedirectUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectUpdateResponseResult]
type zoneSettingMobileRedirectUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type ZoneSettingMobileRedirectUpdateResponseResultID string

const (
	ZoneSettingMobileRedirectUpdateResponseResultIDMobileRedirect ZoneSettingMobileRedirectUpdateResponseResultID = "mobile_redirect"
)

// Value of the zone setting.
type ZoneSettingMobileRedirectUpdateResponseResultValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingMobileRedirectUpdateResponseResultValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                                   `json:"strip_uri"`
	JSON     zoneSettingMobileRedirectUpdateResponseResultValueJSON `json:"-"`
}

// zoneSettingMobileRedirectUpdateResponseResultValueJSON contains the JSON
// metadata for the struct [ZoneSettingMobileRedirectUpdateResponseResultValue]
type zoneSettingMobileRedirectUpdateResponseResultValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectUpdateResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingMobileRedirectUpdateResponseResultValueStatus string

const (
	ZoneSettingMobileRedirectUpdateResponseResultValueStatusOn  ZoneSettingMobileRedirectUpdateResponseResultValueStatus = "on"
	ZoneSettingMobileRedirectUpdateResponseResultValueStatusOff ZoneSettingMobileRedirectUpdateResponseResultValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMobileRedirectUpdateResponseResultEditable bool

const (
	ZoneSettingMobileRedirectUpdateResponseResultEditableTrue  ZoneSettingMobileRedirectUpdateResponseResultEditable = true
	ZoneSettingMobileRedirectUpdateResponseResultEditableFalse ZoneSettingMobileRedirectUpdateResponseResultEditable = false
)

type ZoneSettingMobileRedirectListResponse struct {
	Errors   []ZoneSettingMobileRedirectListResponseError   `json:"errors,required"`
	Messages []ZoneSettingMobileRedirectListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                        `json:"success,required"`
	Result  ZoneSettingMobileRedirectListResponseResult `json:"result"`
	JSON    zoneSettingMobileRedirectListResponseJSON   `json:"-"`
}

// zoneSettingMobileRedirectListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMobileRedirectListResponse]
type zoneSettingMobileRedirectListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingMobileRedirectListResponseErrorJSON `json:"-"`
}

// zoneSettingMobileRedirectListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectListResponseError]
type zoneSettingMobileRedirectListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingMobileRedirectListResponseMessageJSON `json:"-"`
}

// zoneSettingMobileRedirectListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectListResponseMessage]
type zoneSettingMobileRedirectListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMobileRedirectListResponseResult struct {
	// Identifier of the zone setting.
	ID ZoneSettingMobileRedirectListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingMobileRedirectListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMobileRedirectListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMobileRedirectListResponseResultJSON `json:"-"`
}

// zoneSettingMobileRedirectListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingMobileRedirectListResponseResult]
type zoneSettingMobileRedirectListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type ZoneSettingMobileRedirectListResponseResultID string

const (
	ZoneSettingMobileRedirectListResponseResultIDMobileRedirect ZoneSettingMobileRedirectListResponseResultID = "mobile_redirect"
)

// Value of the zone setting.
type ZoneSettingMobileRedirectListResponseResultValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingMobileRedirectListResponseResultValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                                 `json:"strip_uri"`
	JSON     zoneSettingMobileRedirectListResponseResultValueJSON `json:"-"`
}

// zoneSettingMobileRedirectListResponseResultValueJSON contains the JSON metadata
// for the struct [ZoneSettingMobileRedirectListResponseResultValue]
type zoneSettingMobileRedirectListResponseResultValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingMobileRedirectListResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingMobileRedirectListResponseResultValueStatus string

const (
	ZoneSettingMobileRedirectListResponseResultValueStatusOn  ZoneSettingMobileRedirectListResponseResultValueStatus = "on"
	ZoneSettingMobileRedirectListResponseResultValueStatusOff ZoneSettingMobileRedirectListResponseResultValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMobileRedirectListResponseResultEditable bool

const (
	ZoneSettingMobileRedirectListResponseResultEditableTrue  ZoneSettingMobileRedirectListResponseResultEditable = true
	ZoneSettingMobileRedirectListResponseResultEditableFalse ZoneSettingMobileRedirectListResponseResultEditable = false
)

type ZoneSettingMobileRedirectUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingMobileRedirectUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingMobileRedirectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMobileRedirectUpdateParamsValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[ZoneSettingMobileRedirectUpdateParamsValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri param.Field[bool] `json:"strip_uri"`
}

func (r ZoneSettingMobileRedirectUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingMobileRedirectUpdateParamsValueStatus string

const (
	ZoneSettingMobileRedirectUpdateParamsValueStatusOn  ZoneSettingMobileRedirectUpdateParamsValueStatus = "on"
	ZoneSettingMobileRedirectUpdateParamsValueStatusOff ZoneSettingMobileRedirectUpdateParamsValueStatus = "off"
)
