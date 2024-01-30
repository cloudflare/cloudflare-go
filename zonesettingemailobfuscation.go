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

// ZoneSettingEmailObfuscationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingEmailObfuscationService] method instead.
type ZoneSettingEmailObfuscationService struct {
	Options []option.RequestOption
}

// NewZoneSettingEmailObfuscationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingEmailObfuscationService(opts ...option.RequestOption) (r *ZoneSettingEmailObfuscationService) {
	r = &ZoneSettingEmailObfuscationService{}
	r.Options = opts
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *ZoneSettingEmailObfuscationService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingEmailObfuscationUpdateParams, opts ...option.RequestOption) (res *ZoneSettingEmailObfuscationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *ZoneSettingEmailObfuscationService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingEmailObfuscationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingEmailObfuscationUpdateResponse struct {
	Errors   []ZoneSettingEmailObfuscationUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingEmailObfuscationUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                            `json:"success,required"`
	Result  ZoneSettingEmailObfuscationUpdateResponseResult `json:"result"`
	JSON    zoneSettingEmailObfuscationUpdateResponseJSON   `json:"-"`
}

// zoneSettingEmailObfuscationUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingEmailObfuscationUpdateResponse]
type zoneSettingEmailObfuscationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationUpdateResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingEmailObfuscationUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingEmailObfuscationUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingEmailObfuscationUpdateResponseError]
type zoneSettingEmailObfuscationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationUpdateResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingEmailObfuscationUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingEmailObfuscationUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingEmailObfuscationUpdateResponseMessage]
type zoneSettingEmailObfuscationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingEmailObfuscationUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingEmailObfuscationUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEmailObfuscationUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEmailObfuscationUpdateResponseResultJSON `json:"-"`
}

// zoneSettingEmailObfuscationUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingEmailObfuscationUpdateResponseResult]
type zoneSettingEmailObfuscationUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingEmailObfuscationUpdateResponseResultID string

const (
	ZoneSettingEmailObfuscationUpdateResponseResultIDEmailObfuscation ZoneSettingEmailObfuscationUpdateResponseResultID = "email_obfuscation"
)

// Value of the zone setting.
type ZoneSettingEmailObfuscationUpdateResponseResultValue string

const (
	ZoneSettingEmailObfuscationUpdateResponseResultValueOn  ZoneSettingEmailObfuscationUpdateResponseResultValue = "on"
	ZoneSettingEmailObfuscationUpdateResponseResultValueOff ZoneSettingEmailObfuscationUpdateResponseResultValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEmailObfuscationUpdateResponseResultEditable bool

const (
	ZoneSettingEmailObfuscationUpdateResponseResultEditableTrue  ZoneSettingEmailObfuscationUpdateResponseResultEditable = true
	ZoneSettingEmailObfuscationUpdateResponseResultEditableFalse ZoneSettingEmailObfuscationUpdateResponseResultEditable = false
)

type ZoneSettingEmailObfuscationListResponse struct {
	Errors   []ZoneSettingEmailObfuscationListResponseError   `json:"errors,required"`
	Messages []ZoneSettingEmailObfuscationListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                          `json:"success,required"`
	Result  ZoneSettingEmailObfuscationListResponseResult `json:"result"`
	JSON    zoneSettingEmailObfuscationListResponseJSON   `json:"-"`
}

// zoneSettingEmailObfuscationListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingEmailObfuscationListResponse]
type zoneSettingEmailObfuscationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationListResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingEmailObfuscationListResponseErrorJSON `json:"-"`
}

// zoneSettingEmailObfuscationListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingEmailObfuscationListResponseError]
type zoneSettingEmailObfuscationListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationListResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingEmailObfuscationListResponseMessageJSON `json:"-"`
}

// zoneSettingEmailObfuscationListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingEmailObfuscationListResponseMessage]
type zoneSettingEmailObfuscationListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEmailObfuscationListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingEmailObfuscationListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingEmailObfuscationListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEmailObfuscationListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEmailObfuscationListResponseResultJSON `json:"-"`
}

// zoneSettingEmailObfuscationListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingEmailObfuscationListResponseResult]
type zoneSettingEmailObfuscationListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEmailObfuscationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingEmailObfuscationListResponseResultID string

const (
	ZoneSettingEmailObfuscationListResponseResultIDEmailObfuscation ZoneSettingEmailObfuscationListResponseResultID = "email_obfuscation"
)

// Value of the zone setting.
type ZoneSettingEmailObfuscationListResponseResultValue string

const (
	ZoneSettingEmailObfuscationListResponseResultValueOn  ZoneSettingEmailObfuscationListResponseResultValue = "on"
	ZoneSettingEmailObfuscationListResponseResultValueOff ZoneSettingEmailObfuscationListResponseResultValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEmailObfuscationListResponseResultEditable bool

const (
	ZoneSettingEmailObfuscationListResponseResultEditableTrue  ZoneSettingEmailObfuscationListResponseResultEditable = true
	ZoneSettingEmailObfuscationListResponseResultEditableFalse ZoneSettingEmailObfuscationListResponseResultEditable = false
)

type ZoneSettingEmailObfuscationUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingEmailObfuscationUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingEmailObfuscationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingEmailObfuscationUpdateParamsValue string

const (
	ZoneSettingEmailObfuscationUpdateParamsValueOn  ZoneSettingEmailObfuscationUpdateParamsValue = "on"
	ZoneSettingEmailObfuscationUpdateParamsValueOff ZoneSettingEmailObfuscationUpdateParamsValue = "off"
)
