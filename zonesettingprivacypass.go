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

// ZoneSettingPrivacyPassService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingPrivacyPassService]
// method instead.
type ZoneSettingPrivacyPassService struct {
	Options []option.RequestOption
}

// NewZoneSettingPrivacyPassService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingPrivacyPassService(opts ...option.RequestOption) (r *ZoneSettingPrivacyPassService) {
	r = &ZoneSettingPrivacyPassService{}
	r.Options = opts
	return
}

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
func (r *ZoneSettingPrivacyPassService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingPrivacyPassUpdateParams, opts ...option.RequestOption) (res *ZoneSettingPrivacyPassUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/privacy_pass", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
func (r *ZoneSettingPrivacyPassService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingPrivacyPassListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/privacy_pass", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingPrivacyPassUpdateResponse struct {
	Errors   []ZoneSettingPrivacyPassUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingPrivacyPassUpdateResponseMessage `json:"messages"`
	// Privacy Pass is a browser extension developed by the Privacy Pass Team to
	// improve the browsing experience for your visitors. Enabling Privacy Pass will
	// reduce the number of CAPTCHAs shown to your visitors.
	// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
	Result ZoneSettingPrivacyPassUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingPrivacyPassUpdateResponseJSON
}

// zoneSettingPrivacyPassUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingPrivacyPassUpdateResponse]
type zoneSettingPrivacyPassUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrivacyPassUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrivacyPassUpdateResponseErrorJSON
}

// zoneSettingPrivacyPassUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingPrivacyPassUpdateResponseError]
type zoneSettingPrivacyPassUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrivacyPassUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrivacyPassUpdateResponseMessageJSON
}

// zoneSettingPrivacyPassUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingPrivacyPassUpdateResponseMessage]
type zoneSettingPrivacyPassUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
type ZoneSettingPrivacyPassUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingPrivacyPassUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPrivacyPassUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingPrivacyPassUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingPrivacyPassUpdateResponseResultJSON
}

// zoneSettingPrivacyPassUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingPrivacyPassUpdateResponseResult]
type zoneSettingPrivacyPassUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPrivacyPassUpdateResponseResultID string

const (
	ZoneSettingPrivacyPassUpdateResponseResultIDPrivacyPass ZoneSettingPrivacyPassUpdateResponseResultID = "privacy_pass"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPrivacyPassUpdateResponseResultEditable bool

const (
	ZoneSettingPrivacyPassUpdateResponseResultEditableTrue  ZoneSettingPrivacyPassUpdateResponseResultEditable = true
	ZoneSettingPrivacyPassUpdateResponseResultEditableFalse ZoneSettingPrivacyPassUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingPrivacyPassUpdateResponseResultValue string

const (
	ZoneSettingPrivacyPassUpdateResponseResultValueOn  ZoneSettingPrivacyPassUpdateResponseResultValue = "on"
	ZoneSettingPrivacyPassUpdateResponseResultValueOff ZoneSettingPrivacyPassUpdateResponseResultValue = "off"
)

type ZoneSettingPrivacyPassListResponse struct {
	Errors   []ZoneSettingPrivacyPassListResponseError   `json:"errors"`
	Messages []ZoneSettingPrivacyPassListResponseMessage `json:"messages"`
	// Privacy Pass is a browser extension developed by the Privacy Pass Team to
	// improve the browsing experience for your visitors. Enabling Privacy Pass will
	// reduce the number of CAPTCHAs shown to your visitors.
	// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
	Result ZoneSettingPrivacyPassListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingPrivacyPassListResponseJSON
}

// zoneSettingPrivacyPassListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingPrivacyPassListResponse]
type zoneSettingPrivacyPassListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrivacyPassListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrivacyPassListResponseErrorJSON
}

// zoneSettingPrivacyPassListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingPrivacyPassListResponseError]
type zoneSettingPrivacyPassListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPrivacyPassListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingPrivacyPassListResponseMessageJSON
}

// zoneSettingPrivacyPassListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingPrivacyPassListResponseMessage]
type zoneSettingPrivacyPassListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
type ZoneSettingPrivacyPassListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingPrivacyPassListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPrivacyPassListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingPrivacyPassListResponseResultValue `json:"value"`
	JSON  zoneSettingPrivacyPassListResponseResultJSON
}

// zoneSettingPrivacyPassListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingPrivacyPassListResponseResult]
type zoneSettingPrivacyPassListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPrivacyPassListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPrivacyPassListResponseResultID string

const (
	ZoneSettingPrivacyPassListResponseResultIDPrivacyPass ZoneSettingPrivacyPassListResponseResultID = "privacy_pass"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPrivacyPassListResponseResultEditable bool

const (
	ZoneSettingPrivacyPassListResponseResultEditableTrue  ZoneSettingPrivacyPassListResponseResultEditable = true
	ZoneSettingPrivacyPassListResponseResultEditableFalse ZoneSettingPrivacyPassListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingPrivacyPassListResponseResultValue string

const (
	ZoneSettingPrivacyPassListResponseResultValueOn  ZoneSettingPrivacyPassListResponseResultValue = "on"
	ZoneSettingPrivacyPassListResponseResultValueOff ZoneSettingPrivacyPassListResponseResultValue = "off"
)

type ZoneSettingPrivacyPassUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingPrivacyPassUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingPrivacyPassUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingPrivacyPassUpdateParamsValue string

const (
	ZoneSettingPrivacyPassUpdateParamsValueOn  ZoneSettingPrivacyPassUpdateParamsValue = "on"
	ZoneSettingPrivacyPassUpdateParamsValueOff ZoneSettingPrivacyPassUpdateParamsValue = "off"
)
