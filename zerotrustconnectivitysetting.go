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

// Updates the Zero Trust Connectivity Settings for the given account.
func (r *ZerotrustConnectivitySettingService) Edit(ctx context.Context, accountID string, body ZerotrustConnectivitySettingEditParams, opts ...option.RequestOption) (res *ZerotrustConnectivitySettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZerotrustConnectivitySettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type ZerotrustConnectivitySettingEditResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled bool                                         `json:"offramp_warp_enabled"`
	JSON               zerotrustConnectivitySettingEditResponseJSON `json:"-"`
}

// zerotrustConnectivitySettingEditResponseJSON contains the JSON metadata for the
// struct [ZerotrustConnectivitySettingEditResponse]
type zerotrustConnectivitySettingEditResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWARPEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingGetResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled bool                                        `json:"offramp_warp_enabled"`
	JSON               zerotrustConnectivitySettingGetResponseJSON `json:"-"`
}

// zerotrustConnectivitySettingGetResponseJSON contains the JSON metadata for the
// struct [ZerotrustConnectivitySettingGetResponse]
type zerotrustConnectivitySettingGetResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWARPEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingEditParams struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled param.Field[bool] `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled param.Field[bool] `json:"offramp_warp_enabled"`
}

func (r ZerotrustConnectivitySettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZerotrustConnectivitySettingEditResponseEnvelope struct {
	Errors   []ZerotrustConnectivitySettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZerotrustConnectivitySettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZerotrustConnectivitySettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZerotrustConnectivitySettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zerotrustConnectivitySettingEditResponseEnvelopeJSON    `json:"-"`
}

// zerotrustConnectivitySettingEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZerotrustConnectivitySettingEditResponseEnvelope]
type zerotrustConnectivitySettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zerotrustConnectivitySettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zerotrustConnectivitySettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZerotrustConnectivitySettingEditResponseEnvelopeErrors]
type zerotrustConnectivitySettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zerotrustConnectivitySettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zerotrustConnectivitySettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZerotrustConnectivitySettingEditResponseEnvelopeMessages]
type zerotrustConnectivitySettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZerotrustConnectivitySettingEditResponseEnvelopeSuccess bool

const (
	ZerotrustConnectivitySettingEditResponseEnvelopeSuccessTrue ZerotrustConnectivitySettingEditResponseEnvelopeSuccess = true
)

type ZerotrustConnectivitySettingGetResponseEnvelope struct {
	Errors   []ZerotrustConnectivitySettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZerotrustConnectivitySettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZerotrustConnectivitySettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZerotrustConnectivitySettingGetResponseEnvelopeSuccess `json:"success,required"`
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
