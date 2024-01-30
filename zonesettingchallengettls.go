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

// ZoneSettingChallengeTTLSService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingChallengeTTLSService] method instead.
type ZoneSettingChallengeTTLSService struct {
	Options []option.RequestOption
}

// NewZoneSettingChallengeTTLSService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingChallengeTTLSService(opts ...option.RequestOption) (r *ZoneSettingChallengeTTLSService) {
	r = &ZoneSettingChallengeTTLSService{}
	r.Options = opts
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *ZoneSettingChallengeTTLSService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingChallengeTTLSUpdateParams, opts ...option.RequestOption) (res *ZoneSettingChallengeTTLSUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *ZoneSettingChallengeTTLSService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingChallengeTTLSListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingChallengeTTLSUpdateResponse struct {
	Errors   []ZoneSettingChallengeTTLSUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingChallengeTTLSUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                         `json:"success,required"`
	Result  ZoneSettingChallengeTTLSUpdateResponseResult `json:"result"`
	JSON    zoneSettingChallengeTTLSUpdateResponseJSON   `json:"-"`
}

// zoneSettingChallengeTTLSUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLSUpdateResponse]
type zoneSettingChallengeTTLSUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLSUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingChallengeTTLSUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingChallengeTTLSUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLSUpdateResponseError]
type zoneSettingChallengeTTLSUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLSUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingChallengeTTLSUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingChallengeTTLSUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLSUpdateResponseMessage]
type zoneSettingChallengeTTLSUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLSUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTTLSUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingChallengeTTLSUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTTLSUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingChallengeTTLSUpdateResponseResultJSON `json:"-"`
}

// zoneSettingChallengeTTLSUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLSUpdateResponseResult]
type zoneSettingChallengeTTLSUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingChallengeTTLSUpdateResponseResultID string

const (
	ZoneSettingChallengeTTLSUpdateResponseResultIDChallengeTtl ZoneSettingChallengeTTLSUpdateResponseResultID = "challenge_ttl"
)

// Value of the zone setting.
type ZoneSettingChallengeTTLSUpdateResponseResultValue float64

const (
	ZoneSettingChallengeTTLSUpdateResponseResultValue300      ZoneSettingChallengeTTLSUpdateResponseResultValue = 300
	ZoneSettingChallengeTTLSUpdateResponseResultValue900      ZoneSettingChallengeTTLSUpdateResponseResultValue = 900
	ZoneSettingChallengeTTLSUpdateResponseResultValue1800     ZoneSettingChallengeTTLSUpdateResponseResultValue = 1800
	ZoneSettingChallengeTTLSUpdateResponseResultValue2700     ZoneSettingChallengeTTLSUpdateResponseResultValue = 2700
	ZoneSettingChallengeTTLSUpdateResponseResultValue3600     ZoneSettingChallengeTTLSUpdateResponseResultValue = 3600
	ZoneSettingChallengeTTLSUpdateResponseResultValue7200     ZoneSettingChallengeTTLSUpdateResponseResultValue = 7200
	ZoneSettingChallengeTTLSUpdateResponseResultValue10800    ZoneSettingChallengeTTLSUpdateResponseResultValue = 10800
	ZoneSettingChallengeTTLSUpdateResponseResultValue14400    ZoneSettingChallengeTTLSUpdateResponseResultValue = 14400
	ZoneSettingChallengeTTLSUpdateResponseResultValue28800    ZoneSettingChallengeTTLSUpdateResponseResultValue = 28800
	ZoneSettingChallengeTTLSUpdateResponseResultValue57600    ZoneSettingChallengeTTLSUpdateResponseResultValue = 57600
	ZoneSettingChallengeTTLSUpdateResponseResultValue86400    ZoneSettingChallengeTTLSUpdateResponseResultValue = 86400
	ZoneSettingChallengeTTLSUpdateResponseResultValue604800   ZoneSettingChallengeTTLSUpdateResponseResultValue = 604800
	ZoneSettingChallengeTTLSUpdateResponseResultValue2592000  ZoneSettingChallengeTTLSUpdateResponseResultValue = 2592000
	ZoneSettingChallengeTTLSUpdateResponseResultValue31536000 ZoneSettingChallengeTTLSUpdateResponseResultValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTTLSUpdateResponseResultEditable bool

const (
	ZoneSettingChallengeTTLSUpdateResponseResultEditableTrue  ZoneSettingChallengeTTLSUpdateResponseResultEditable = true
	ZoneSettingChallengeTTLSUpdateResponseResultEditableFalse ZoneSettingChallengeTTLSUpdateResponseResultEditable = false
)

type ZoneSettingChallengeTTLSListResponse struct {
	Errors   []ZoneSettingChallengeTTLSListResponseError   `json:"errors,required"`
	Messages []ZoneSettingChallengeTTLSListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                       `json:"success,required"`
	Result  ZoneSettingChallengeTTLSListResponseResult `json:"result"`
	JSON    zoneSettingChallengeTTLSListResponseJSON   `json:"-"`
}

// zoneSettingChallengeTTLSListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLSListResponse]
type zoneSettingChallengeTTLSListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLSListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingChallengeTTLSListResponseErrorJSON `json:"-"`
}

// zoneSettingChallengeTTLSListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLSListResponseError]
type zoneSettingChallengeTTLSListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLSListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingChallengeTTLSListResponseMessageJSON `json:"-"`
}

// zoneSettingChallengeTTLSListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLSListResponseMessage]
type zoneSettingChallengeTTLSListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLSListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTTLSListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingChallengeTTLSListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTTLSListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingChallengeTTLSListResponseResultJSON `json:"-"`
}

// zoneSettingChallengeTTLSListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLSListResponseResult]
type zoneSettingChallengeTTLSListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLSListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingChallengeTTLSListResponseResultID string

const (
	ZoneSettingChallengeTTLSListResponseResultIDChallengeTtl ZoneSettingChallengeTTLSListResponseResultID = "challenge_ttl"
)

// Value of the zone setting.
type ZoneSettingChallengeTTLSListResponseResultValue float64

const (
	ZoneSettingChallengeTTLSListResponseResultValue300      ZoneSettingChallengeTTLSListResponseResultValue = 300
	ZoneSettingChallengeTTLSListResponseResultValue900      ZoneSettingChallengeTTLSListResponseResultValue = 900
	ZoneSettingChallengeTTLSListResponseResultValue1800     ZoneSettingChallengeTTLSListResponseResultValue = 1800
	ZoneSettingChallengeTTLSListResponseResultValue2700     ZoneSettingChallengeTTLSListResponseResultValue = 2700
	ZoneSettingChallengeTTLSListResponseResultValue3600     ZoneSettingChallengeTTLSListResponseResultValue = 3600
	ZoneSettingChallengeTTLSListResponseResultValue7200     ZoneSettingChallengeTTLSListResponseResultValue = 7200
	ZoneSettingChallengeTTLSListResponseResultValue10800    ZoneSettingChallengeTTLSListResponseResultValue = 10800
	ZoneSettingChallengeTTLSListResponseResultValue14400    ZoneSettingChallengeTTLSListResponseResultValue = 14400
	ZoneSettingChallengeTTLSListResponseResultValue28800    ZoneSettingChallengeTTLSListResponseResultValue = 28800
	ZoneSettingChallengeTTLSListResponseResultValue57600    ZoneSettingChallengeTTLSListResponseResultValue = 57600
	ZoneSettingChallengeTTLSListResponseResultValue86400    ZoneSettingChallengeTTLSListResponseResultValue = 86400
	ZoneSettingChallengeTTLSListResponseResultValue604800   ZoneSettingChallengeTTLSListResponseResultValue = 604800
	ZoneSettingChallengeTTLSListResponseResultValue2592000  ZoneSettingChallengeTTLSListResponseResultValue = 2592000
	ZoneSettingChallengeTTLSListResponseResultValue31536000 ZoneSettingChallengeTTLSListResponseResultValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTTLSListResponseResultEditable bool

const (
	ZoneSettingChallengeTTLSListResponseResultEditableTrue  ZoneSettingChallengeTTLSListResponseResultEditable = true
	ZoneSettingChallengeTTLSListResponseResultEditableFalse ZoneSettingChallengeTTLSListResponseResultEditable = false
)

type ZoneSettingChallengeTTLSUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingChallengeTTLSUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingChallengeTTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingChallengeTTLSUpdateParamsValue float64

const (
	ZoneSettingChallengeTTLSUpdateParamsValue300      ZoneSettingChallengeTTLSUpdateParamsValue = 300
	ZoneSettingChallengeTTLSUpdateParamsValue900      ZoneSettingChallengeTTLSUpdateParamsValue = 900
	ZoneSettingChallengeTTLSUpdateParamsValue1800     ZoneSettingChallengeTTLSUpdateParamsValue = 1800
	ZoneSettingChallengeTTLSUpdateParamsValue2700     ZoneSettingChallengeTTLSUpdateParamsValue = 2700
	ZoneSettingChallengeTTLSUpdateParamsValue3600     ZoneSettingChallengeTTLSUpdateParamsValue = 3600
	ZoneSettingChallengeTTLSUpdateParamsValue7200     ZoneSettingChallengeTTLSUpdateParamsValue = 7200
	ZoneSettingChallengeTTLSUpdateParamsValue10800    ZoneSettingChallengeTTLSUpdateParamsValue = 10800
	ZoneSettingChallengeTTLSUpdateParamsValue14400    ZoneSettingChallengeTTLSUpdateParamsValue = 14400
	ZoneSettingChallengeTTLSUpdateParamsValue28800    ZoneSettingChallengeTTLSUpdateParamsValue = 28800
	ZoneSettingChallengeTTLSUpdateParamsValue57600    ZoneSettingChallengeTTLSUpdateParamsValue = 57600
	ZoneSettingChallengeTTLSUpdateParamsValue86400    ZoneSettingChallengeTTLSUpdateParamsValue = 86400
	ZoneSettingChallengeTTLSUpdateParamsValue604800   ZoneSettingChallengeTTLSUpdateParamsValue = 604800
	ZoneSettingChallengeTTLSUpdateParamsValue2592000  ZoneSettingChallengeTTLSUpdateParamsValue = 2592000
	ZoneSettingChallengeTTLSUpdateParamsValue31536000 ZoneSettingChallengeTTLSUpdateParamsValue = 31536000
)
