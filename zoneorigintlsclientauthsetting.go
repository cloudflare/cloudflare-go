// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneOriginTlsClientAuthSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneOriginTlsClientAuthSettingService] method instead.
type ZoneOriginTlsClientAuthSettingService struct {
	Options []option.RequestOption
}

// NewZoneOriginTlsClientAuthSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneOriginTlsClientAuthSettingService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthSettingService) {
	r = &ZoneOriginTlsClientAuthSettingService{}
	r.Options = opts
	return
}

// Get whether zone-level authenticated origin pulls is enabled or not. It is false
// by default.
func (r *ZoneOriginTlsClientAuthSettingService) ZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZone(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable zone-level authenticated origin pulls. 'enabled' should be set
// true either before/after the certificate is uploaded to see the certificate in
// use.
func (r *ZoneOriginTlsClientAuthSettingService) ZoneLevelAuthenticatedOriginPullsSetEnablementForZone(ctx context.Context, zoneIdentifier string, body ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse struct {
	Errors   []ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseMessage `json:"messages"`
	Result   ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseError struct {
	Code    int64                                                                                                       `json:"code,required"`
	Message string                                                                                                      `json:"message,required"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseError]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseMessage struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseMessage]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseResult struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                                                                                                         `json:"enabled"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseResult]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseSuccess bool

const (
	ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseSuccessTrue ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseSuccess = true
)

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse struct {
	Errors   []ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseError   `json:"errors"`
	Messages []ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseMessage `json:"messages"`
	Result   ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseSuccess `json:"success"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseError struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseErrorJSON `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseError]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseMessage struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseMessageJSON `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseMessage]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseResult struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                                                                                                  `json:"enabled"`
	JSON    zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseResultJSON `json:"-"`
}

// zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseResult]
type zoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseSuccess bool

const (
	ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseSuccessTrue ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseSuccess = true
)

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
