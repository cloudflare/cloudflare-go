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
func (r *ZoneOriginTlsClientAuthSettingService) ZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZone(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EnabledResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable zone-level authenticated origin pulls. 'enabled' should be set
// true either before/after the certificate is uploaded to see the certificate in
// use.
func (r *ZoneOriginTlsClientAuthSettingService) ZoneLevelAuthenticatedOriginPullsSetEnablementForZone(ctx context.Context, zoneIdentifier string, body ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams, opts ...option.RequestOption) (res *EnabledResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type EnabledResponse struct {
	Errors   []EnabledResponseError   `json:"errors"`
	Messages []EnabledResponseMessage `json:"messages"`
	Result   EnabledResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EnabledResponseSuccess `json:"success"`
	JSON    enabledResponseJSON    `json:"-"`
}

// enabledResponseJSON contains the JSON metadata for the struct [EnabledResponse]
type enabledResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnabledResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EnabledResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    enabledResponseErrorJSON `json:"-"`
}

// enabledResponseErrorJSON contains the JSON metadata for the struct
// [EnabledResponseError]
type enabledResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnabledResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EnabledResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    enabledResponseMessageJSON `json:"-"`
}

// enabledResponseMessageJSON contains the JSON metadata for the struct
// [EnabledResponseMessage]
type enabledResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnabledResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EnabledResponseResult struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                      `json:"enabled"`
	JSON    enabledResponseResultJSON `json:"-"`
}

// enabledResponseResultJSON contains the JSON metadata for the struct
// [EnabledResponseResult]
type enabledResponseResultJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnabledResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EnabledResponseSuccess bool

const (
	EnabledResponseSuccessTrue EnabledResponseSuccess = true
)

type ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneOriginTlsClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
