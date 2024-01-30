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

// ZoneSettingMinTlsVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingMinTlsVersionService] method instead.
type ZoneSettingMinTlsVersionService struct {
	Options []option.RequestOption
}

// NewZoneSettingMinTlsVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingMinTlsVersionService(opts ...option.RequestOption) (r *ZoneSettingMinTlsVersionService) {
	r = &ZoneSettingMinTlsVersionService{}
	r.Options = opts
	return
}

// Changes Minimum TLS Version setting.
func (r *ZoneSettingMinTlsVersionService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingMinTlsVersionUpdateParams, opts ...option.RequestOption) (res *ZoneSettingMinTlsVersionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets Minimum TLS Version setting.
func (r *ZoneSettingMinTlsVersionService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingMinTlsVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingMinTlsVersionUpdateResponse struct {
	Errors   []ZoneSettingMinTlsVersionUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingMinTlsVersionUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                         `json:"success,required"`
	Result  ZoneSettingMinTlsVersionUpdateResponseResult `json:"result"`
	JSON    zoneSettingMinTlsVersionUpdateResponseJSON   `json:"-"`
}

// zoneSettingMinTlsVersionUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMinTlsVersionUpdateResponse]
type zoneSettingMinTlsVersionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTlsVersionUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingMinTlsVersionUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingMinTlsVersionUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingMinTlsVersionUpdateResponseError]
type zoneSettingMinTlsVersionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTlsVersionUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingMinTlsVersionUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingMinTlsVersionUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingMinTlsVersionUpdateResponseMessage]
type zoneSettingMinTlsVersionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTlsVersionUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingMinTlsVersionUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingMinTlsVersionUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinTlsVersionUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinTlsVersionUpdateResponseResultJSON `json:"-"`
}

// zoneSettingMinTlsVersionUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingMinTlsVersionUpdateResponseResult]
type zoneSettingMinTlsVersionUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingMinTlsVersionUpdateResponseResultID string

const (
	ZoneSettingMinTlsVersionUpdateResponseResultIDMinTlsVersion ZoneSettingMinTlsVersionUpdateResponseResultID = "min_tls_version"
)

// Value of the zone setting.
type ZoneSettingMinTlsVersionUpdateResponseResultValue string

const (
	ZoneSettingMinTlsVersionUpdateResponseResultValue1_0 ZoneSettingMinTlsVersionUpdateResponseResultValue = "1.0"
	ZoneSettingMinTlsVersionUpdateResponseResultValue1_1 ZoneSettingMinTlsVersionUpdateResponseResultValue = "1.1"
	ZoneSettingMinTlsVersionUpdateResponseResultValue1_2 ZoneSettingMinTlsVersionUpdateResponseResultValue = "1.2"
	ZoneSettingMinTlsVersionUpdateResponseResultValue1_3 ZoneSettingMinTlsVersionUpdateResponseResultValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinTlsVersionUpdateResponseResultEditable bool

const (
	ZoneSettingMinTlsVersionUpdateResponseResultEditableTrue  ZoneSettingMinTlsVersionUpdateResponseResultEditable = true
	ZoneSettingMinTlsVersionUpdateResponseResultEditableFalse ZoneSettingMinTlsVersionUpdateResponseResultEditable = false
)

type ZoneSettingMinTlsVersionListResponse struct {
	Errors   []ZoneSettingMinTlsVersionListResponseError   `json:"errors,required"`
	Messages []ZoneSettingMinTlsVersionListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                       `json:"success,required"`
	Result  ZoneSettingMinTlsVersionListResponseResult `json:"result"`
	JSON    zoneSettingMinTlsVersionListResponseJSON   `json:"-"`
}

// zoneSettingMinTlsVersionListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMinTlsVersionListResponse]
type zoneSettingMinTlsVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTlsVersionListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingMinTlsVersionListResponseErrorJSON `json:"-"`
}

// zoneSettingMinTlsVersionListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingMinTlsVersionListResponseError]
type zoneSettingMinTlsVersionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTlsVersionListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingMinTlsVersionListResponseMessageJSON `json:"-"`
}

// zoneSettingMinTlsVersionListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingMinTlsVersionListResponseMessage]
type zoneSettingMinTlsVersionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTlsVersionListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingMinTlsVersionListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingMinTlsVersionListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinTlsVersionListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinTlsVersionListResponseResultJSON `json:"-"`
}

// zoneSettingMinTlsVersionListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingMinTlsVersionListResponseResult]
type zoneSettingMinTlsVersionListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTlsVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingMinTlsVersionListResponseResultID string

const (
	ZoneSettingMinTlsVersionListResponseResultIDMinTlsVersion ZoneSettingMinTlsVersionListResponseResultID = "min_tls_version"
)

// Value of the zone setting.
type ZoneSettingMinTlsVersionListResponseResultValue string

const (
	ZoneSettingMinTlsVersionListResponseResultValue1_0 ZoneSettingMinTlsVersionListResponseResultValue = "1.0"
	ZoneSettingMinTlsVersionListResponseResultValue1_1 ZoneSettingMinTlsVersionListResponseResultValue = "1.1"
	ZoneSettingMinTlsVersionListResponseResultValue1_2 ZoneSettingMinTlsVersionListResponseResultValue = "1.2"
	ZoneSettingMinTlsVersionListResponseResultValue1_3 ZoneSettingMinTlsVersionListResponseResultValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinTlsVersionListResponseResultEditable bool

const (
	ZoneSettingMinTlsVersionListResponseResultEditableTrue  ZoneSettingMinTlsVersionListResponseResultEditable = true
	ZoneSettingMinTlsVersionListResponseResultEditableFalse ZoneSettingMinTlsVersionListResponseResultEditable = false
)

type ZoneSettingMinTlsVersionUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingMinTlsVersionUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingMinTlsVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMinTlsVersionUpdateParamsValue string

const (
	ZoneSettingMinTlsVersionUpdateParamsValue1_0 ZoneSettingMinTlsVersionUpdateParamsValue = "1.0"
	ZoneSettingMinTlsVersionUpdateParamsValue1_1 ZoneSettingMinTlsVersionUpdateParamsValue = "1.1"
	ZoneSettingMinTlsVersionUpdateParamsValue1_2 ZoneSettingMinTlsVersionUpdateParamsValue = "1.2"
	ZoneSettingMinTlsVersionUpdateParamsValue1_3 ZoneSettingMinTlsVersionUpdateParamsValue = "1.3"
)
