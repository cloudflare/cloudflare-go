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

// ZoneSettingTlsClientAuthService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingTlsClientAuthService] method instead.
type ZoneSettingTlsClientAuthService struct {
	Options []option.RequestOption
}

// NewZoneSettingTlsClientAuthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingTlsClientAuthService(opts ...option.RequestOption) (r *ZoneSettingTlsClientAuthService) {
	r = &ZoneSettingTlsClientAuthService{}
	r.Options = opts
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZoneSettingTlsClientAuthService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingTlsClientAuthUpdateParams, opts ...option.RequestOption) (res *ZoneSettingTlsClientAuthUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZoneSettingTlsClientAuthService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingTlsClientAuthListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingTlsClientAuthUpdateResponse struct {
	Errors   []ZoneSettingTlsClientAuthUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingTlsClientAuthUpdateResponseMessage `json:"messages"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTlsClientAuthUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingTlsClientAuthUpdateResponseJSON
}

// zoneSettingTlsClientAuthUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingTlsClientAuthUpdateResponse]
type zoneSettingTlsClientAuthUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTlsClientAuthUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingTlsClientAuthUpdateResponseErrorJSON
}

// zoneSettingTlsClientAuthUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingTlsClientAuthUpdateResponseError]
type zoneSettingTlsClientAuthUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTlsClientAuthUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingTlsClientAuthUpdateResponseMessageJSON
}

// zoneSettingTlsClientAuthUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingTlsClientAuthUpdateResponseMessage]
type zoneSettingTlsClientAuthUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingTlsClientAuthUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTlsClientAuthUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTlsClientAuthUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingTlsClientAuthUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingTlsClientAuthUpdateResponseResultJSON
}

// zoneSettingTlsClientAuthUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingTlsClientAuthUpdateResponseResult]
type zoneSettingTlsClientAuthUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTlsClientAuthUpdateResponseResultID string

const (
	ZoneSettingTlsClientAuthUpdateResponseResultIDTlsClientAuth ZoneSettingTlsClientAuthUpdateResponseResultID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTlsClientAuthUpdateResponseResultEditable bool

const (
	ZoneSettingTlsClientAuthUpdateResponseResultEditableTrue  ZoneSettingTlsClientAuthUpdateResponseResultEditable = true
	ZoneSettingTlsClientAuthUpdateResponseResultEditableFalse ZoneSettingTlsClientAuthUpdateResponseResultEditable = false
)

// value of the zone setting.
type ZoneSettingTlsClientAuthUpdateResponseResultValue string

const (
	ZoneSettingTlsClientAuthUpdateResponseResultValueOn  ZoneSettingTlsClientAuthUpdateResponseResultValue = "on"
	ZoneSettingTlsClientAuthUpdateResponseResultValueOff ZoneSettingTlsClientAuthUpdateResponseResultValue = "off"
)

type ZoneSettingTlsClientAuthListResponse struct {
	Errors   []ZoneSettingTlsClientAuthListResponseError   `json:"errors"`
	Messages []ZoneSettingTlsClientAuthListResponseMessage `json:"messages"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTlsClientAuthListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingTlsClientAuthListResponseJSON
}

// zoneSettingTlsClientAuthListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingTlsClientAuthListResponse]
type zoneSettingTlsClientAuthListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTlsClientAuthListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingTlsClientAuthListResponseErrorJSON
}

// zoneSettingTlsClientAuthListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingTlsClientAuthListResponseError]
type zoneSettingTlsClientAuthListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTlsClientAuthListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingTlsClientAuthListResponseMessageJSON
}

// zoneSettingTlsClientAuthListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingTlsClientAuthListResponseMessage]
type zoneSettingTlsClientAuthListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingTlsClientAuthListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTlsClientAuthListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTlsClientAuthListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingTlsClientAuthListResponseResultValue `json:"value"`
	JSON  zoneSettingTlsClientAuthListResponseResultJSON
}

// zoneSettingTlsClientAuthListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingTlsClientAuthListResponseResult]
type zoneSettingTlsClientAuthListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTlsClientAuthListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTlsClientAuthListResponseResultID string

const (
	ZoneSettingTlsClientAuthListResponseResultIDTlsClientAuth ZoneSettingTlsClientAuthListResponseResultID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTlsClientAuthListResponseResultEditable bool

const (
	ZoneSettingTlsClientAuthListResponseResultEditableTrue  ZoneSettingTlsClientAuthListResponseResultEditable = true
	ZoneSettingTlsClientAuthListResponseResultEditableFalse ZoneSettingTlsClientAuthListResponseResultEditable = false
)

// value of the zone setting.
type ZoneSettingTlsClientAuthListResponseResultValue string

const (
	ZoneSettingTlsClientAuthListResponseResultValueOn  ZoneSettingTlsClientAuthListResponseResultValue = "on"
	ZoneSettingTlsClientAuthListResponseResultValueOff ZoneSettingTlsClientAuthListResponseResultValue = "off"
)

type ZoneSettingTlsClientAuthUpdateParams struct {
	// value of the zone setting.
	Value param.Field[ZoneSettingTlsClientAuthUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingTlsClientAuthUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// value of the zone setting.
type ZoneSettingTlsClientAuthUpdateParamsValue string

const (
	ZoneSettingTlsClientAuthUpdateParamsValueOn  ZoneSettingTlsClientAuthUpdateParamsValue = "on"
	ZoneSettingTlsClientAuthUpdateParamsValueOff ZoneSettingTlsClientAuthUpdateParamsValue = "off"
)
