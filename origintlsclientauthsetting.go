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

// OriginTLSClientAuthSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewOriginTLSClientAuthSettingService] method instead.
type OriginTLSClientAuthSettingService struct {
	Options []option.RequestOption
}

// NewOriginTLSClientAuthSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOriginTLSClientAuthSettingService(opts ...option.RequestOption) (r *OriginTLSClientAuthSettingService) {
	r = &OriginTLSClientAuthSettingService{}
	r.Options = opts
	return
}

// Get whether zone-level authenticated origin pulls is enabled or not. It is false
// by default.
func (r *OriginTLSClientAuthSettingService) ZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZone(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable or disable zone-level authenticated origin pulls. 'enabled' should be set
// true either before/after the certificate is uploaded to see the certificate in
// use.
func (r *OriginTLSClientAuthSettingService) ZoneLevelAuthenticatedOriginPullsSetEnablementForZone(ctx context.Context, zoneID string, body OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams, opts ...option.RequestOption) (res *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                                                                                               `json:"enabled"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseJSON `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                                                                                        `json:"enabled"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseJSON `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelope struct {
	Errors   []OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelope]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeErrors struct {
	Code    int64                                                                                                            `json:"code,required"`
	Message string                                                                                                           `json:"message,required"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeErrors]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeMessages struct {
	Code    int64                                                                                                              `json:"code,required"`
	Message string                                                                                                             `json:"message,required"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeMessages]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeSuccessTrue OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsGetEnablementSettingForZoneResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelope struct {
	Errors   []OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelope]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeErrors struct {
	Code    int64                                                                                                     `json:"code,required"`
	Message string                                                                                                    `json:"message,required"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeErrors]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeMessages struct {
	Code    int64                                                                                                       `json:"code,required"`
	Message string                                                                                                      `json:"message,required"`
	JSON    originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeMessages]
type originTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeSuccessTrue OriginTLSClientAuthSettingZoneLevelAuthenticatedOriginPullsSetEnablementForZoneResponseEnvelopeSuccess = true
)
