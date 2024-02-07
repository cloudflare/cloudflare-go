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

// ZerotrustConnectivitySettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZerotrustConnectivitySettingService] method instead.
type ZerotrustConnectivitySettingService struct {
	Options []option.RequestOption
}

// NewZerotrustConnectivitySettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZerotrustConnectivitySettingService(opts ...option.RequestOption) (r *ZerotrustConnectivitySettingService) {
	r = &ZerotrustConnectivitySettingService{}
	r.Options = opts
	return
}

// Gets the Zero Trust Connectivity Settings for the given account.
func (r *ZerotrustConnectivitySettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ZerotrustConnectivitySettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZerotrustConnectivitySettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the Zero Trust Connectivity Settings for the given account.
func (r *ZerotrustConnectivitySettingService) Update(ctx context.Context, accountID string, body ZerotrustConnectivitySettingUpdateParams, opts ...option.RequestOption) (res *ZerotrustConnectivitySettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZerotrustConnectivitySettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZerotrustConnectivitySettingGetResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled bool                                        `json:"offramp_warp_enabled"`
	JSON               zerotrustConnectivitySettingGetResponseJSON `json:"-"`
}

// zerotrustConnectivitySettingGetResponseJSON contains the JSON metadata for the
// struct [ZerotrustConnectivitySettingGetResponse]
type zerotrustConnectivitySettingGetResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWarpEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingUpdateResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled bool                                           `json:"offramp_warp_enabled"`
	JSON               zerotrustConnectivitySettingUpdateResponseJSON `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseJSON contains the JSON metadata for
// the struct [ZerotrustConnectivitySettingUpdateResponse]
type zerotrustConnectivitySettingUpdateResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWarpEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingGetResponseEnvelope struct {
	Errors   []ZerotrustConnectivitySettingGetResponseEnvelopeErrors   `json:"errors"`
	Messages []ZerotrustConnectivitySettingGetResponseEnvelopeMessages `json:"messages"`
	Result   ZerotrustConnectivitySettingGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success ZerotrustConnectivitySettingGetResponseEnvelopeSuccess `json:"success"`
	JSON    zerotrustConnectivitySettingGetResponseEnvelopeJSON    `json:"-"`
}

// zerotrustConnectivitySettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZerotrustConnectivitySettingGetResponseEnvelope]
type zerotrustConnectivitySettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zerotrustConnectivitySettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zerotrustConnectivitySettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZerotrustConnectivitySettingGetResponseEnvelopeErrors]
type zerotrustConnectivitySettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zerotrustConnectivitySettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zerotrustConnectivitySettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZerotrustConnectivitySettingGetResponseEnvelopeMessages]
type zerotrustConnectivitySettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZerotrustConnectivitySettingGetResponseEnvelopeSuccess bool

const (
	ZerotrustConnectivitySettingGetResponseEnvelopeSuccessTrue ZerotrustConnectivitySettingGetResponseEnvelopeSuccess = true
)

type ZerotrustConnectivitySettingUpdateParams struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled param.Field[bool] `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled param.Field[bool] `json:"offramp_warp_enabled"`
}

func (r ZerotrustConnectivitySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZerotrustConnectivitySettingUpdateResponseEnvelope struct {
	Errors   []ZerotrustConnectivitySettingUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []ZerotrustConnectivitySettingUpdateResponseEnvelopeMessages `json:"messages"`
	Result   ZerotrustConnectivitySettingUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success ZerotrustConnectivitySettingUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    zerotrustConnectivitySettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZerotrustConnectivitySettingUpdateResponseEnvelope]
type zerotrustConnectivitySettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zerotrustConnectivitySettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZerotrustConnectivitySettingUpdateResponseEnvelopeErrors]
type zerotrustConnectivitySettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zerotrustConnectivitySettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZerotrustConnectivitySettingUpdateResponseEnvelopeMessages]
type zerotrustConnectivitySettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZerotrustConnectivitySettingUpdateResponseEnvelopeSuccess bool

const (
	ZerotrustConnectivitySettingUpdateResponseEnvelopeSuccessTrue ZerotrustConnectivitySettingUpdateResponseEnvelopeSuccess = true
)
