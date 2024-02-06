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

// ZoneSettingChallengeTTLService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingChallengeTTLService] method instead.
type ZoneSettingChallengeTTLService struct {
	Options []option.RequestOption
}

// NewZoneSettingChallengeTTLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingChallengeTTLService(opts ...option.RequestOption) (r *ZoneSettingChallengeTTLService) {
	r = &ZoneSettingChallengeTTLService{}
	r.Options = opts
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *ZoneSettingChallengeTTLService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingChallengeTTLUpdateParams, opts ...option.RequestOption) (res *ZoneSettingChallengeTTLUpdateResponse, err error) {
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
func (r *ZoneSettingChallengeTTLService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingChallengeTTLListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingChallengeTTLUpdateResponse struct {
	Errors   []ZoneSettingChallengeTTLUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingChallengeTTLUpdateResponseMessage `json:"messages"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZoneSettingChallengeTTLUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                      `json:"success"`
	JSON    zoneSettingChallengeTTLUpdateResponseJSON `json:"-"`
}

// zoneSettingChallengeTTLUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLUpdateResponse]
type zoneSettingChallengeTTLUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingChallengeTTLUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingChallengeTTLUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLUpdateResponseError]
type zoneSettingChallengeTTLUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingChallengeTTLUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingChallengeTTLUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLUpdateResponseMessage]
type zoneSettingChallengeTTLUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingChallengeTTLUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTTLUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTTLUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingChallengeTTLUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingChallengeTTLUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingChallengeTTLUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLUpdateResponseResult]
type zoneSettingChallengeTTLUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingChallengeTTLUpdateResponseResultID string

const (
	ZoneSettingChallengeTTLUpdateResponseResultIDChallengeTTL ZoneSettingChallengeTTLUpdateResponseResultID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTTLUpdateResponseResultEditable bool

const (
	ZoneSettingChallengeTTLUpdateResponseResultEditableTrue  ZoneSettingChallengeTTLUpdateResponseResultEditable = true
	ZoneSettingChallengeTTLUpdateResponseResultEditableFalse ZoneSettingChallengeTTLUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingChallengeTTLUpdateResponseResultValue float64

const (
	ZoneSettingChallengeTTLUpdateResponseResultValue300      ZoneSettingChallengeTTLUpdateResponseResultValue = 300
	ZoneSettingChallengeTTLUpdateResponseResultValue900      ZoneSettingChallengeTTLUpdateResponseResultValue = 900
	ZoneSettingChallengeTTLUpdateResponseResultValue1800     ZoneSettingChallengeTTLUpdateResponseResultValue = 1800
	ZoneSettingChallengeTTLUpdateResponseResultValue2700     ZoneSettingChallengeTTLUpdateResponseResultValue = 2700
	ZoneSettingChallengeTTLUpdateResponseResultValue3600     ZoneSettingChallengeTTLUpdateResponseResultValue = 3600
	ZoneSettingChallengeTTLUpdateResponseResultValue7200     ZoneSettingChallengeTTLUpdateResponseResultValue = 7200
	ZoneSettingChallengeTTLUpdateResponseResultValue10800    ZoneSettingChallengeTTLUpdateResponseResultValue = 10800
	ZoneSettingChallengeTTLUpdateResponseResultValue14400    ZoneSettingChallengeTTLUpdateResponseResultValue = 14400
	ZoneSettingChallengeTTLUpdateResponseResultValue28800    ZoneSettingChallengeTTLUpdateResponseResultValue = 28800
	ZoneSettingChallengeTTLUpdateResponseResultValue57600    ZoneSettingChallengeTTLUpdateResponseResultValue = 57600
	ZoneSettingChallengeTTLUpdateResponseResultValue86400    ZoneSettingChallengeTTLUpdateResponseResultValue = 86400
	ZoneSettingChallengeTTLUpdateResponseResultValue604800   ZoneSettingChallengeTTLUpdateResponseResultValue = 604800
	ZoneSettingChallengeTTLUpdateResponseResultValue2592000  ZoneSettingChallengeTTLUpdateResponseResultValue = 2592000
	ZoneSettingChallengeTTLUpdateResponseResultValue31536000 ZoneSettingChallengeTTLUpdateResponseResultValue = 31536000
)

type ZoneSettingChallengeTTLListResponse struct {
	Errors   []ZoneSettingChallengeTTLListResponseError   `json:"errors"`
	Messages []ZoneSettingChallengeTTLListResponseMessage `json:"messages"`
	// Specify how long a visitor is allowed access to your site after successfully
	// completing a challenge (such as a CAPTCHA). After the TTL has expired the
	// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
	// setting and will attempt to honor any setting above 45 minutes.
	// (https://support.cloudflare.com/hc/en-us/articles/200170136).
	Result ZoneSettingChallengeTTLListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                    `json:"success"`
	JSON    zoneSettingChallengeTTLListResponseJSON `json:"-"`
}

// zoneSettingChallengeTTLListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLListResponse]
type zoneSettingChallengeTTLListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingChallengeTTLListResponseErrorJSON `json:"-"`
}

// zoneSettingChallengeTTLListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLListResponseError]
type zoneSettingChallengeTTLListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTTLListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingChallengeTTLListResponseMessageJSON `json:"-"`
}

// zoneSettingChallengeTTLListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTTLListResponseMessage]
type zoneSettingChallengeTTLListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingChallengeTTLListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTTLListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTTLListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingChallengeTTLListResponseResultValue `json:"value"`
	JSON  zoneSettingChallengeTTLListResponseResultJSON  `json:"-"`
}

// zoneSettingChallengeTTLListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTTLListResponseResult]
type zoneSettingChallengeTTLListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTTLListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingChallengeTTLListResponseResultID string

const (
	ZoneSettingChallengeTTLListResponseResultIDChallengeTTL ZoneSettingChallengeTTLListResponseResultID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTTLListResponseResultEditable bool

const (
	ZoneSettingChallengeTTLListResponseResultEditableTrue  ZoneSettingChallengeTTLListResponseResultEditable = true
	ZoneSettingChallengeTTLListResponseResultEditableFalse ZoneSettingChallengeTTLListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingChallengeTTLListResponseResultValue float64

const (
	ZoneSettingChallengeTTLListResponseResultValue300      ZoneSettingChallengeTTLListResponseResultValue = 300
	ZoneSettingChallengeTTLListResponseResultValue900      ZoneSettingChallengeTTLListResponseResultValue = 900
	ZoneSettingChallengeTTLListResponseResultValue1800     ZoneSettingChallengeTTLListResponseResultValue = 1800
	ZoneSettingChallengeTTLListResponseResultValue2700     ZoneSettingChallengeTTLListResponseResultValue = 2700
	ZoneSettingChallengeTTLListResponseResultValue3600     ZoneSettingChallengeTTLListResponseResultValue = 3600
	ZoneSettingChallengeTTLListResponseResultValue7200     ZoneSettingChallengeTTLListResponseResultValue = 7200
	ZoneSettingChallengeTTLListResponseResultValue10800    ZoneSettingChallengeTTLListResponseResultValue = 10800
	ZoneSettingChallengeTTLListResponseResultValue14400    ZoneSettingChallengeTTLListResponseResultValue = 14400
	ZoneSettingChallengeTTLListResponseResultValue28800    ZoneSettingChallengeTTLListResponseResultValue = 28800
	ZoneSettingChallengeTTLListResponseResultValue57600    ZoneSettingChallengeTTLListResponseResultValue = 57600
	ZoneSettingChallengeTTLListResponseResultValue86400    ZoneSettingChallengeTTLListResponseResultValue = 86400
	ZoneSettingChallengeTTLListResponseResultValue604800   ZoneSettingChallengeTTLListResponseResultValue = 604800
	ZoneSettingChallengeTTLListResponseResultValue2592000  ZoneSettingChallengeTTLListResponseResultValue = 2592000
	ZoneSettingChallengeTTLListResponseResultValue31536000 ZoneSettingChallengeTTLListResponseResultValue = 31536000
)

type ZoneSettingChallengeTTLUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingChallengeTTLUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingChallengeTTLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingChallengeTTLUpdateParamsValue float64

const (
	ZoneSettingChallengeTTLUpdateParamsValue300      ZoneSettingChallengeTTLUpdateParamsValue = 300
	ZoneSettingChallengeTTLUpdateParamsValue900      ZoneSettingChallengeTTLUpdateParamsValue = 900
	ZoneSettingChallengeTTLUpdateParamsValue1800     ZoneSettingChallengeTTLUpdateParamsValue = 1800
	ZoneSettingChallengeTTLUpdateParamsValue2700     ZoneSettingChallengeTTLUpdateParamsValue = 2700
	ZoneSettingChallengeTTLUpdateParamsValue3600     ZoneSettingChallengeTTLUpdateParamsValue = 3600
	ZoneSettingChallengeTTLUpdateParamsValue7200     ZoneSettingChallengeTTLUpdateParamsValue = 7200
	ZoneSettingChallengeTTLUpdateParamsValue10800    ZoneSettingChallengeTTLUpdateParamsValue = 10800
	ZoneSettingChallengeTTLUpdateParamsValue14400    ZoneSettingChallengeTTLUpdateParamsValue = 14400
	ZoneSettingChallengeTTLUpdateParamsValue28800    ZoneSettingChallengeTTLUpdateParamsValue = 28800
	ZoneSettingChallengeTTLUpdateParamsValue57600    ZoneSettingChallengeTTLUpdateParamsValue = 57600
	ZoneSettingChallengeTTLUpdateParamsValue86400    ZoneSettingChallengeTTLUpdateParamsValue = 86400
	ZoneSettingChallengeTTLUpdateParamsValue604800   ZoneSettingChallengeTTLUpdateParamsValue = 604800
	ZoneSettingChallengeTTLUpdateParamsValue2592000  ZoneSettingChallengeTTLUpdateParamsValue = 2592000
	ZoneSettingChallengeTTLUpdateParamsValue31536000 ZoneSettingChallengeTTLUpdateParamsValue = 31536000
)
