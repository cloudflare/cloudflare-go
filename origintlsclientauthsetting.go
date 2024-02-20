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

// Enable or disable zone-level authenticated origin pulls. 'enabled' should be set
// true either before/after the certificate is uploaded to see the certificate in
// use.
func (r *OriginTLSClientAuthSettingService) Update(ctx context.Context, zoneID string, body OriginTLSClientAuthSettingUpdateParams, opts ...option.RequestOption) (res *OriginTLSClientAuthSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthSettingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get whether zone-level authenticated origin pulls is enabled or not. It is false
// by default.
func (r *OriginTLSClientAuthSettingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *OriginTLSClientAuthSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginTLSClientAuthSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginTLSClientAuthSettingUpdateResponse struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                                         `json:"enabled"`
	JSON    originTLSClientAuthSettingUpdateResponseJSON `json:"-"`
}

// originTLSClientAuthSettingUpdateResponseJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthSettingUpdateResponse]
type originTLSClientAuthSettingUpdateResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingGetResponse struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool                                      `json:"enabled"`
	JSON    originTLSClientAuthSettingGetResponseJSON `json:"-"`
}

// originTLSClientAuthSettingGetResponseJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthSettingGetResponse]
type originTLSClientAuthSettingGetResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingUpdateParams struct {
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r OriginTLSClientAuthSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthSettingUpdateResponseEnvelope struct {
	Errors   []OriginTLSClientAuthSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthSettingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthSettingUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthSettingUpdateResponseEnvelope]
type originTLSClientAuthSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    originTLSClientAuthSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthSettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthSettingUpdateResponseEnvelopeErrors]
type originTLSClientAuthSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    originTLSClientAuthSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [OriginTLSClientAuthSettingUpdateResponseEnvelopeMessages]
type originTLSClientAuthSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthSettingUpdateResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthSettingUpdateResponseEnvelopeSuccessTrue OriginTLSClientAuthSettingUpdateResponseEnvelopeSuccess = true
)

type OriginTLSClientAuthSettingGetResponseEnvelope struct {
	Errors   []OriginTLSClientAuthSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginTLSClientAuthSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originTLSClientAuthSettingGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthSettingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthSettingGetResponseEnvelope]
type originTLSClientAuthSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    originTLSClientAuthSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// originTLSClientAuthSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthSettingGetResponseEnvelopeErrors]
type originTLSClientAuthSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginTLSClientAuthSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    originTLSClientAuthSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// originTLSClientAuthSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthSettingGetResponseEnvelopeMessages]
type originTLSClientAuthSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginTLSClientAuthSettingGetResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthSettingGetResponseEnvelopeSuccessTrue OriginTLSClientAuthSettingGetResponseEnvelopeSuccess = true
)
