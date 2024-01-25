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

// AccountZerotrustConnectivitySettingService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountZerotrustConnectivitySettingService] method instead.
type AccountZerotrustConnectivitySettingService struct {
	Options []option.RequestOption
}

// NewAccountZerotrustConnectivitySettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountZerotrustConnectivitySettingService(opts ...option.RequestOption) (r *AccountZerotrustConnectivitySettingService) {
	r = &AccountZerotrustConnectivitySettingService{}
	r.Options = opts
	return
}

// Updates the Zero Trust Connectivity Settings for the given account.
func (r *AccountZerotrustConnectivitySettingService) Update(ctx context.Context, accountIdentifier string, body AccountZerotrustConnectivitySettingUpdateParams, opts ...option.RequestOption) (res *AccountZerotrustConnectivitySettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets the Zero Trust Connectivity Settings for the given account.
func (r *AccountZerotrustConnectivitySettingService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountZerotrustConnectivitySettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountZerotrustConnectivitySettingUpdateResponse struct {
	Errors   []AccountZerotrustConnectivitySettingUpdateResponseError   `json:"errors"`
	Messages []AccountZerotrustConnectivitySettingUpdateResponseMessage `json:"messages"`
	Result   AccountZerotrustConnectivitySettingUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountZerotrustConnectivitySettingUpdateResponseSuccess `json:"success"`
	JSON    accountZerotrustConnectivitySettingUpdateResponseJSON    `json:"-"`
}

// accountZerotrustConnectivitySettingUpdateResponseJSON contains the JSON metadata
// for the struct [AccountZerotrustConnectivitySettingUpdateResponse]
type accountZerotrustConnectivitySettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountZerotrustConnectivitySettingUpdateResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountZerotrustConnectivitySettingUpdateResponseErrorJSON `json:"-"`
}

// accountZerotrustConnectivitySettingUpdateResponseErrorJSON contains the JSON
// metadata for the struct [AccountZerotrustConnectivitySettingUpdateResponseError]
type accountZerotrustConnectivitySettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountZerotrustConnectivitySettingUpdateResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountZerotrustConnectivitySettingUpdateResponseMessageJSON `json:"-"`
}

// accountZerotrustConnectivitySettingUpdateResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountZerotrustConnectivitySettingUpdateResponseMessage]
type accountZerotrustConnectivitySettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountZerotrustConnectivitySettingUpdateResponseResult struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled bool                                                        `json:"offramp_warp_enabled"`
	JSON               accountZerotrustConnectivitySettingUpdateResponseResultJSON `json:"-"`
}

// accountZerotrustConnectivitySettingUpdateResponseResultJSON contains the JSON
// metadata for the struct
// [AccountZerotrustConnectivitySettingUpdateResponseResult]
type accountZerotrustConnectivitySettingUpdateResponseResultJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWarpEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountZerotrustConnectivitySettingUpdateResponseSuccess bool

const (
	AccountZerotrustConnectivitySettingUpdateResponseSuccessTrue AccountZerotrustConnectivitySettingUpdateResponseSuccess = true
)

type AccountZerotrustConnectivitySettingListResponse struct {
	Errors   []AccountZerotrustConnectivitySettingListResponseError   `json:"errors"`
	Messages []AccountZerotrustConnectivitySettingListResponseMessage `json:"messages"`
	Result   AccountZerotrustConnectivitySettingListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountZerotrustConnectivitySettingListResponseSuccess `json:"success"`
	JSON    accountZerotrustConnectivitySettingListResponseJSON    `json:"-"`
}

// accountZerotrustConnectivitySettingListResponseJSON contains the JSON metadata
// for the struct [AccountZerotrustConnectivitySettingListResponse]
type accountZerotrustConnectivitySettingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountZerotrustConnectivitySettingListResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountZerotrustConnectivitySettingListResponseErrorJSON `json:"-"`
}

// accountZerotrustConnectivitySettingListResponseErrorJSON contains the JSON
// metadata for the struct [AccountZerotrustConnectivitySettingListResponseError]
type accountZerotrustConnectivitySettingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountZerotrustConnectivitySettingListResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountZerotrustConnectivitySettingListResponseMessageJSON `json:"-"`
}

// accountZerotrustConnectivitySettingListResponseMessageJSON contains the JSON
// metadata for the struct [AccountZerotrustConnectivitySettingListResponseMessage]
type accountZerotrustConnectivitySettingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountZerotrustConnectivitySettingListResponseResult struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled bool                                                      `json:"offramp_warp_enabled"`
	JSON               accountZerotrustConnectivitySettingListResponseResultJSON `json:"-"`
}

// accountZerotrustConnectivitySettingListResponseResultJSON contains the JSON
// metadata for the struct [AccountZerotrustConnectivitySettingListResponseResult]
type accountZerotrustConnectivitySettingListResponseResultJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWarpEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountZerotrustConnectivitySettingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountZerotrustConnectivitySettingListResponseSuccess bool

const (
	AccountZerotrustConnectivitySettingListResponseSuccessTrue AccountZerotrustConnectivitySettingListResponseSuccess = true
)

type AccountZerotrustConnectivitySettingUpdateParams struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled param.Field[bool] `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled param.Field[bool] `json:"offramp_warp_enabled"`
}

func (r AccountZerotrustConnectivitySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
