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

// ZoneSettingChallengeTtlService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingChallengeTtlService] method instead.
type ZoneSettingChallengeTtlService struct {
	Options []option.RequestOption
}

// NewZoneSettingChallengeTtlService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingChallengeTtlService(opts ...option.RequestOption) (r *ZoneSettingChallengeTtlService) {
	r = &ZoneSettingChallengeTtlService{}
	r.Options = opts
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *ZoneSettingChallengeTtlService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingChallengeTtlUpdateParams, opts ...option.RequestOption) (res *ZoneSettingChallengeTtlUpdateResponse, err error) {
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
func (r *ZoneSettingChallengeTtlService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingChallengeTtlListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingChallengeTtlUpdateResponse struct {
	Errors   []ZoneSettingChallengeTtlUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingChallengeTtlUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                        `json:"success,required"`
	Result  ZoneSettingChallengeTtlUpdateResponseResult `json:"result"`
	JSON    zoneSettingChallengeTtlUpdateResponseJSON   `json:"-"`
}

// zoneSettingChallengeTtlUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTtlUpdateResponse]
type zoneSettingChallengeTtlUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTtlUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingChallengeTtlUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingChallengeTtlUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTtlUpdateResponseError]
type zoneSettingChallengeTtlUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTtlUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingChallengeTtlUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingChallengeTtlUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTtlUpdateResponseMessage]
type zoneSettingChallengeTtlUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTtlUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTtlUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingChallengeTtlUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTtlUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingChallengeTtlUpdateResponseResultJSON `json:"-"`
}

// zoneSettingChallengeTtlUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTtlUpdateResponseResult]
type zoneSettingChallengeTtlUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingChallengeTtlUpdateResponseResultID string

const (
	ZoneSettingChallengeTtlUpdateResponseResultIDChallengeTtl ZoneSettingChallengeTtlUpdateResponseResultID = "challenge_ttl"
)

// Value of the zone setting.
type ZoneSettingChallengeTtlUpdateResponseResultValue float64

const (
	ZoneSettingChallengeTtlUpdateResponseResultValue300      ZoneSettingChallengeTtlUpdateResponseResultValue = 300
	ZoneSettingChallengeTtlUpdateResponseResultValue900      ZoneSettingChallengeTtlUpdateResponseResultValue = 900
	ZoneSettingChallengeTtlUpdateResponseResultValue1800     ZoneSettingChallengeTtlUpdateResponseResultValue = 1800
	ZoneSettingChallengeTtlUpdateResponseResultValue2700     ZoneSettingChallengeTtlUpdateResponseResultValue = 2700
	ZoneSettingChallengeTtlUpdateResponseResultValue3600     ZoneSettingChallengeTtlUpdateResponseResultValue = 3600
	ZoneSettingChallengeTtlUpdateResponseResultValue7200     ZoneSettingChallengeTtlUpdateResponseResultValue = 7200
	ZoneSettingChallengeTtlUpdateResponseResultValue10800    ZoneSettingChallengeTtlUpdateResponseResultValue = 10800
	ZoneSettingChallengeTtlUpdateResponseResultValue14400    ZoneSettingChallengeTtlUpdateResponseResultValue = 14400
	ZoneSettingChallengeTtlUpdateResponseResultValue28800    ZoneSettingChallengeTtlUpdateResponseResultValue = 28800
	ZoneSettingChallengeTtlUpdateResponseResultValue57600    ZoneSettingChallengeTtlUpdateResponseResultValue = 57600
	ZoneSettingChallengeTtlUpdateResponseResultValue86400    ZoneSettingChallengeTtlUpdateResponseResultValue = 86400
	ZoneSettingChallengeTtlUpdateResponseResultValue604800   ZoneSettingChallengeTtlUpdateResponseResultValue = 604800
	ZoneSettingChallengeTtlUpdateResponseResultValue2592000  ZoneSettingChallengeTtlUpdateResponseResultValue = 2592000
	ZoneSettingChallengeTtlUpdateResponseResultValue31536000 ZoneSettingChallengeTtlUpdateResponseResultValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTtlUpdateResponseResultEditable bool

const (
	ZoneSettingChallengeTtlUpdateResponseResultEditableTrue  ZoneSettingChallengeTtlUpdateResponseResultEditable = true
	ZoneSettingChallengeTtlUpdateResponseResultEditableFalse ZoneSettingChallengeTtlUpdateResponseResultEditable = false
)

type ZoneSettingChallengeTtlListResponse struct {
	Errors   []ZoneSettingChallengeTtlListResponseError   `json:"errors,required"`
	Messages []ZoneSettingChallengeTtlListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                      `json:"success,required"`
	Result  ZoneSettingChallengeTtlListResponseResult `json:"result"`
	JSON    zoneSettingChallengeTtlListResponseJSON   `json:"-"`
}

// zoneSettingChallengeTtlListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTtlListResponse]
type zoneSettingChallengeTtlListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTtlListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingChallengeTtlListResponseErrorJSON `json:"-"`
}

// zoneSettingChallengeTtlListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTtlListResponseError]
type zoneSettingChallengeTtlListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTtlListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingChallengeTtlListResponseMessageJSON `json:"-"`
}

// zoneSettingChallengeTtlListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingChallengeTtlListResponseMessage]
type zoneSettingChallengeTtlListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingChallengeTtlListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingChallengeTtlListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingChallengeTtlListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingChallengeTtlListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingChallengeTtlListResponseResultJSON `json:"-"`
}

// zoneSettingChallengeTtlListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingChallengeTtlListResponseResult]
type zoneSettingChallengeTtlListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingChallengeTtlListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingChallengeTtlListResponseResultID string

const (
	ZoneSettingChallengeTtlListResponseResultIDChallengeTtl ZoneSettingChallengeTtlListResponseResultID = "challenge_ttl"
)

// Value of the zone setting.
type ZoneSettingChallengeTtlListResponseResultValue float64

const (
	ZoneSettingChallengeTtlListResponseResultValue300      ZoneSettingChallengeTtlListResponseResultValue = 300
	ZoneSettingChallengeTtlListResponseResultValue900      ZoneSettingChallengeTtlListResponseResultValue = 900
	ZoneSettingChallengeTtlListResponseResultValue1800     ZoneSettingChallengeTtlListResponseResultValue = 1800
	ZoneSettingChallengeTtlListResponseResultValue2700     ZoneSettingChallengeTtlListResponseResultValue = 2700
	ZoneSettingChallengeTtlListResponseResultValue3600     ZoneSettingChallengeTtlListResponseResultValue = 3600
	ZoneSettingChallengeTtlListResponseResultValue7200     ZoneSettingChallengeTtlListResponseResultValue = 7200
	ZoneSettingChallengeTtlListResponseResultValue10800    ZoneSettingChallengeTtlListResponseResultValue = 10800
	ZoneSettingChallengeTtlListResponseResultValue14400    ZoneSettingChallengeTtlListResponseResultValue = 14400
	ZoneSettingChallengeTtlListResponseResultValue28800    ZoneSettingChallengeTtlListResponseResultValue = 28800
	ZoneSettingChallengeTtlListResponseResultValue57600    ZoneSettingChallengeTtlListResponseResultValue = 57600
	ZoneSettingChallengeTtlListResponseResultValue86400    ZoneSettingChallengeTtlListResponseResultValue = 86400
	ZoneSettingChallengeTtlListResponseResultValue604800   ZoneSettingChallengeTtlListResponseResultValue = 604800
	ZoneSettingChallengeTtlListResponseResultValue2592000  ZoneSettingChallengeTtlListResponseResultValue = 2592000
	ZoneSettingChallengeTtlListResponseResultValue31536000 ZoneSettingChallengeTtlListResponseResultValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingChallengeTtlListResponseResultEditable bool

const (
	ZoneSettingChallengeTtlListResponseResultEditableTrue  ZoneSettingChallengeTtlListResponseResultEditable = true
	ZoneSettingChallengeTtlListResponseResultEditableFalse ZoneSettingChallengeTtlListResponseResultEditable = false
)

type ZoneSettingChallengeTtlUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingChallengeTtlUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingChallengeTtlUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingChallengeTtlUpdateParamsValue float64

const (
	ZoneSettingChallengeTtlUpdateParamsValue300      ZoneSettingChallengeTtlUpdateParamsValue = 300
	ZoneSettingChallengeTtlUpdateParamsValue900      ZoneSettingChallengeTtlUpdateParamsValue = 900
	ZoneSettingChallengeTtlUpdateParamsValue1800     ZoneSettingChallengeTtlUpdateParamsValue = 1800
	ZoneSettingChallengeTtlUpdateParamsValue2700     ZoneSettingChallengeTtlUpdateParamsValue = 2700
	ZoneSettingChallengeTtlUpdateParamsValue3600     ZoneSettingChallengeTtlUpdateParamsValue = 3600
	ZoneSettingChallengeTtlUpdateParamsValue7200     ZoneSettingChallengeTtlUpdateParamsValue = 7200
	ZoneSettingChallengeTtlUpdateParamsValue10800    ZoneSettingChallengeTtlUpdateParamsValue = 10800
	ZoneSettingChallengeTtlUpdateParamsValue14400    ZoneSettingChallengeTtlUpdateParamsValue = 14400
	ZoneSettingChallengeTtlUpdateParamsValue28800    ZoneSettingChallengeTtlUpdateParamsValue = 28800
	ZoneSettingChallengeTtlUpdateParamsValue57600    ZoneSettingChallengeTtlUpdateParamsValue = 57600
	ZoneSettingChallengeTtlUpdateParamsValue86400    ZoneSettingChallengeTtlUpdateParamsValue = 86400
	ZoneSettingChallengeTtlUpdateParamsValue604800   ZoneSettingChallengeTtlUpdateParamsValue = 604800
	ZoneSettingChallengeTtlUpdateParamsValue2592000  ZoneSettingChallengeTtlUpdateParamsValue = 2592000
	ZoneSettingChallengeTtlUpdateParamsValue31536000 ZoneSettingChallengeTtlUpdateParamsValue = 31536000
)
