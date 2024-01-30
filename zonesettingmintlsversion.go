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

// ZoneSettingMinTLSVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingMinTLSVersionService] method instead.
type ZoneSettingMinTLSVersionService struct {
	Options []option.RequestOption
}

// NewZoneSettingMinTLSVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingMinTLSVersionService(opts ...option.RequestOption) (r *ZoneSettingMinTLSVersionService) {
	r = &ZoneSettingMinTLSVersionService{}
	r.Options = opts
	return
}

// Changes Minimum TLS Version setting.
func (r *ZoneSettingMinTLSVersionService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingMinTLSVersionUpdateParams, opts ...option.RequestOption) (res *ZoneSettingMinTLSVersionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets Minimum TLS Version setting.
func (r *ZoneSettingMinTLSVersionService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingMinTLSVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingMinTLSVersionUpdateResponse struct {
	Errors   []ZoneSettingMinTLSVersionUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingMinTLSVersionUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                         `json:"success,required"`
	Result  ZoneSettingMinTLSVersionUpdateResponseResult `json:"result"`
	JSON    zoneSettingMinTLSVersionUpdateResponseJSON   `json:"-"`
}

// zoneSettingMinTLSVersionUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMinTLSVersionUpdateResponse]
type zoneSettingMinTLSVersionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTLSVersionUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingMinTLSVersionUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingMinTLSVersionUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingMinTLSVersionUpdateResponseError]
type zoneSettingMinTLSVersionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTLSVersionUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingMinTLSVersionUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingMinTLSVersionUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingMinTLSVersionUpdateResponseMessage]
type zoneSettingMinTLSVersionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTLSVersionUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingMinTLSVersionUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingMinTLSVersionUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinTLSVersionUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinTLSVersionUpdateResponseResultJSON `json:"-"`
}

// zoneSettingMinTLSVersionUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingMinTLSVersionUpdateResponseResult]
type zoneSettingMinTLSVersionUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingMinTLSVersionUpdateResponseResultID string

const (
	ZoneSettingMinTLSVersionUpdateResponseResultIDMinTLSVersion ZoneSettingMinTLSVersionUpdateResponseResultID = "min_tls_version"
)

// Value of the zone setting.
type ZoneSettingMinTLSVersionUpdateResponseResultValue string

const (
	ZoneSettingMinTLSVersionUpdateResponseResultValue1_0 ZoneSettingMinTLSVersionUpdateResponseResultValue = "1.0"
	ZoneSettingMinTLSVersionUpdateResponseResultValue1_1 ZoneSettingMinTLSVersionUpdateResponseResultValue = "1.1"
	ZoneSettingMinTLSVersionUpdateResponseResultValue1_2 ZoneSettingMinTLSVersionUpdateResponseResultValue = "1.2"
	ZoneSettingMinTLSVersionUpdateResponseResultValue1_3 ZoneSettingMinTLSVersionUpdateResponseResultValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinTLSVersionUpdateResponseResultEditable bool

const (
	ZoneSettingMinTLSVersionUpdateResponseResultEditableTrue  ZoneSettingMinTLSVersionUpdateResponseResultEditable = true
	ZoneSettingMinTLSVersionUpdateResponseResultEditableFalse ZoneSettingMinTLSVersionUpdateResponseResultEditable = false
)

type ZoneSettingMinTLSVersionListResponse struct {
	Errors   []ZoneSettingMinTLSVersionListResponseError   `json:"errors,required"`
	Messages []ZoneSettingMinTLSVersionListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                       `json:"success,required"`
	Result  ZoneSettingMinTLSVersionListResponseResult `json:"result"`
	JSON    zoneSettingMinTLSVersionListResponseJSON   `json:"-"`
}

// zoneSettingMinTLSVersionListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingMinTLSVersionListResponse]
type zoneSettingMinTLSVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTLSVersionListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingMinTLSVersionListResponseErrorJSON `json:"-"`
}

// zoneSettingMinTLSVersionListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingMinTLSVersionListResponseError]
type zoneSettingMinTLSVersionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTLSVersionListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingMinTLSVersionListResponseMessageJSON `json:"-"`
}

// zoneSettingMinTLSVersionListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingMinTLSVersionListResponseMessage]
type zoneSettingMinTLSVersionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinTLSVersionListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingMinTLSVersionListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingMinTLSVersionListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinTLSVersionListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinTLSVersionListResponseResultJSON `json:"-"`
}

// zoneSettingMinTLSVersionListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingMinTLSVersionListResponseResult]
type zoneSettingMinTLSVersionListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinTLSVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingMinTLSVersionListResponseResultID string

const (
	ZoneSettingMinTLSVersionListResponseResultIDMinTLSVersion ZoneSettingMinTLSVersionListResponseResultID = "min_tls_version"
)

// Value of the zone setting.
type ZoneSettingMinTLSVersionListResponseResultValue string

const (
	ZoneSettingMinTLSVersionListResponseResultValue1_0 ZoneSettingMinTLSVersionListResponseResultValue = "1.0"
	ZoneSettingMinTLSVersionListResponseResultValue1_1 ZoneSettingMinTLSVersionListResponseResultValue = "1.1"
	ZoneSettingMinTLSVersionListResponseResultValue1_2 ZoneSettingMinTLSVersionListResponseResultValue = "1.2"
	ZoneSettingMinTLSVersionListResponseResultValue1_3 ZoneSettingMinTLSVersionListResponseResultValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinTLSVersionListResponseResultEditable bool

const (
	ZoneSettingMinTLSVersionListResponseResultEditableTrue  ZoneSettingMinTLSVersionListResponseResultEditable = true
	ZoneSettingMinTLSVersionListResponseResultEditableFalse ZoneSettingMinTLSVersionListResponseResultEditable = false
)

type ZoneSettingMinTLSVersionUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingMinTLSVersionUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingMinTLSVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMinTLSVersionUpdateParamsValue string

const (
	ZoneSettingMinTLSVersionUpdateParamsValue1_0 ZoneSettingMinTLSVersionUpdateParamsValue = "1.0"
	ZoneSettingMinTLSVersionUpdateParamsValue1_1 ZoneSettingMinTLSVersionUpdateParamsValue = "1.1"
	ZoneSettingMinTLSVersionUpdateParamsValue1_2 ZoneSettingMinTLSVersionUpdateParamsValue = "1.2"
	ZoneSettingMinTLSVersionUpdateParamsValue1_3 ZoneSettingMinTLSVersionUpdateParamsValue = "1.3"
)
