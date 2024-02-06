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

// ZoneSettingTLSClientAuthService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingTLSClientAuthService] method instead.
type ZoneSettingTLSClientAuthService struct {
	Options []option.RequestOption
}

// NewZoneSettingTLSClientAuthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingTLSClientAuthService(opts ...option.RequestOption) (r *ZoneSettingTLSClientAuthService) {
	r = &ZoneSettingTLSClientAuthService{}
	r.Options = opts
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZoneSettingTLSClientAuthService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingTLSClientAuthUpdateParams, opts ...option.RequestOption) (res *ZoneSettingTLSClientAuthUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZoneSettingTLSClientAuthService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingTLSClientAuthListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingTLSClientAuthUpdateResponse struct {
	Errors   []ZoneSettingTLSClientAuthUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingTLSClientAuthUpdateResponseMessage `json:"messages"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTLSClientAuthUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                       `json:"success"`
	JSON    zoneSettingTLSClientAuthUpdateResponseJSON `json:"-"`
}

// zoneSettingTLSClientAuthUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingTLSClientAuthUpdateResponse]
type zoneSettingTLSClientAuthUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTLSClientAuthUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingTLSClientAuthUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingTLSClientAuthUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingTLSClientAuthUpdateResponseError]
type zoneSettingTLSClientAuthUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTLSClientAuthUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingTLSClientAuthUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingTLSClientAuthUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingTLSClientAuthUpdateResponseMessage]
type zoneSettingTLSClientAuthUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingTLSClientAuthUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTLSClientAuthUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTLSClientAuthUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingTLSClientAuthUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingTLSClientAuthUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingTLSClientAuthUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingTLSClientAuthUpdateResponseResult]
type zoneSettingTLSClientAuthUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTLSClientAuthUpdateResponseResultID string

const (
	ZoneSettingTLSClientAuthUpdateResponseResultIDTLSClientAuth ZoneSettingTLSClientAuthUpdateResponseResultID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTLSClientAuthUpdateResponseResultEditable bool

const (
	ZoneSettingTLSClientAuthUpdateResponseResultEditableTrue  ZoneSettingTLSClientAuthUpdateResponseResultEditable = true
	ZoneSettingTLSClientAuthUpdateResponseResultEditableFalse ZoneSettingTLSClientAuthUpdateResponseResultEditable = false
)

// value of the zone setting.
type ZoneSettingTLSClientAuthUpdateResponseResultValue string

const (
	ZoneSettingTLSClientAuthUpdateResponseResultValueOn  ZoneSettingTLSClientAuthUpdateResponseResultValue = "on"
	ZoneSettingTLSClientAuthUpdateResponseResultValueOff ZoneSettingTLSClientAuthUpdateResponseResultValue = "off"
)

type ZoneSettingTLSClientAuthListResponse struct {
	Errors   []ZoneSettingTLSClientAuthListResponseError   `json:"errors"`
	Messages []ZoneSettingTLSClientAuthListResponseMessage `json:"messages"`
	// TLS Client Auth requires Cloudflare to connect to your origin server using a
	// client certificate (Enterprise Only).
	Result ZoneSettingTLSClientAuthListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                     `json:"success"`
	JSON    zoneSettingTLSClientAuthListResponseJSON `json:"-"`
}

// zoneSettingTLSClientAuthListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingTLSClientAuthListResponse]
type zoneSettingTLSClientAuthListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTLSClientAuthListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingTLSClientAuthListResponseErrorJSON `json:"-"`
}

// zoneSettingTLSClientAuthListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingTLSClientAuthListResponseError]
type zoneSettingTLSClientAuthListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingTLSClientAuthListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingTLSClientAuthListResponseMessageJSON `json:"-"`
}

// zoneSettingTLSClientAuthListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingTLSClientAuthListResponseMessage]
type zoneSettingTLSClientAuthListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingTLSClientAuthListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingTLSClientAuthListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingTLSClientAuthListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingTLSClientAuthListResponseResultValue `json:"value"`
	JSON  zoneSettingTLSClientAuthListResponseResultJSON  `json:"-"`
}

// zoneSettingTLSClientAuthListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingTLSClientAuthListResponseResult]
type zoneSettingTLSClientAuthListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingTLSClientAuthListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingTLSClientAuthListResponseResultID string

const (
	ZoneSettingTLSClientAuthListResponseResultIDTLSClientAuth ZoneSettingTLSClientAuthListResponseResultID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingTLSClientAuthListResponseResultEditable bool

const (
	ZoneSettingTLSClientAuthListResponseResultEditableTrue  ZoneSettingTLSClientAuthListResponseResultEditable = true
	ZoneSettingTLSClientAuthListResponseResultEditableFalse ZoneSettingTLSClientAuthListResponseResultEditable = false
)

// value of the zone setting.
type ZoneSettingTLSClientAuthListResponseResultValue string

const (
	ZoneSettingTLSClientAuthListResponseResultValueOn  ZoneSettingTLSClientAuthListResponseResultValue = "on"
	ZoneSettingTLSClientAuthListResponseResultValueOff ZoneSettingTLSClientAuthListResponseResultValue = "off"
)

type ZoneSettingTLSClientAuthUpdateParams struct {
	// value of the zone setting.
	Value param.Field[ZoneSettingTLSClientAuthUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingTLSClientAuthUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// value of the zone setting.
type ZoneSettingTLSClientAuthUpdateParamsValue string

const (
	ZoneSettingTLSClientAuthUpdateParamsValueOn  ZoneSettingTLSClientAuthUpdateParamsValue = "on"
	ZoneSettingTLSClientAuthUpdateParamsValueOff ZoneSettingTLSClientAuthUpdateParamsValue = "off"
)
