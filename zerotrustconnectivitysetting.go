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
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Zero Trust Connectivity Settings for the given account.
func (r *ZerotrustConnectivitySettingService) Update(ctx context.Context, accountID string, body ZerotrustConnectivitySettingUpdateParams, opts ...option.RequestOption) (res *ZerotrustConnectivitySettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZerotrustConnectivitySettingGetResponse struct {
	Errors   []ZerotrustConnectivitySettingGetResponseError   `json:"errors"`
	Messages []ZerotrustConnectivitySettingGetResponseMessage `json:"messages"`
	Result   ZerotrustConnectivitySettingGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZerotrustConnectivitySettingGetResponseSuccess `json:"success"`
	JSON    zerotrustConnectivitySettingGetResponseJSON    `json:"-"`
}

// zerotrustConnectivitySettingGetResponseJSON contains the JSON metadata for the
// struct [ZerotrustConnectivitySettingGetResponse]
type zerotrustConnectivitySettingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingGetResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zerotrustConnectivitySettingGetResponseErrorJSON `json:"-"`
}

// zerotrustConnectivitySettingGetResponseErrorJSON contains the JSON metadata for
// the struct [ZerotrustConnectivitySettingGetResponseError]
type zerotrustConnectivitySettingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingGetResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zerotrustConnectivitySettingGetResponseMessageJSON `json:"-"`
}

// zerotrustConnectivitySettingGetResponseMessageJSON contains the JSON metadata
// for the struct [ZerotrustConnectivitySettingGetResponseMessage]
type zerotrustConnectivitySettingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingGetResponseResult struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled bool                                              `json:"offramp_warp_enabled"`
	JSON               zerotrustConnectivitySettingGetResponseResultJSON `json:"-"`
}

// zerotrustConnectivitySettingGetResponseResultJSON contains the JSON metadata for
// the struct [ZerotrustConnectivitySettingGetResponseResult]
type zerotrustConnectivitySettingGetResponseResultJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWarpEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZerotrustConnectivitySettingGetResponseSuccess bool

const (
	ZerotrustConnectivitySettingGetResponseSuccessTrue ZerotrustConnectivitySettingGetResponseSuccess = true
)

type ZerotrustConnectivitySettingUpdateResponse struct {
	Errors   []ZerotrustConnectivitySettingUpdateResponseError   `json:"errors"`
	Messages []ZerotrustConnectivitySettingUpdateResponseMessage `json:"messages"`
	Result   ZerotrustConnectivitySettingUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZerotrustConnectivitySettingUpdateResponseSuccess `json:"success"`
	JSON    zerotrustConnectivitySettingUpdateResponseJSON    `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseJSON contains the JSON metadata for
// the struct [ZerotrustConnectivitySettingUpdateResponse]
type zerotrustConnectivitySettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingUpdateResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zerotrustConnectivitySettingUpdateResponseErrorJSON `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZerotrustConnectivitySettingUpdateResponseError]
type zerotrustConnectivitySettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingUpdateResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zerotrustConnectivitySettingUpdateResponseMessageJSON `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZerotrustConnectivitySettingUpdateResponseMessage]
type zerotrustConnectivitySettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZerotrustConnectivitySettingUpdateResponseResult struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled bool                                                 `json:"offramp_warp_enabled"`
	JSON               zerotrustConnectivitySettingUpdateResponseResultJSON `json:"-"`
}

// zerotrustConnectivitySettingUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZerotrustConnectivitySettingUpdateResponseResult]
type zerotrustConnectivitySettingUpdateResponseResultJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWarpEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZerotrustConnectivitySettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZerotrustConnectivitySettingUpdateResponseSuccess bool

const (
	ZerotrustConnectivitySettingUpdateResponseSuccessTrue ZerotrustConnectivitySettingUpdateResponseSuccess = true
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
