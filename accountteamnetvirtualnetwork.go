// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountTeamnetVirtualNetworkService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountTeamnetVirtualNetworkService] method instead.
type AccountTeamnetVirtualNetworkService struct {
	Options []option.RequestOption
}

// NewAccountTeamnetVirtualNetworkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountTeamnetVirtualNetworkService(opts ...option.RequestOption) (r *AccountTeamnetVirtualNetworkService) {
	r = &AccountTeamnetVirtualNetworkService{}
	r.Options = opts
	return
}

// Updates an existing virtual network.
func (r *AccountTeamnetVirtualNetworkService) Update(ctx context.Context, accountIdentifier string, vnetID string, body AccountTeamnetVirtualNetworkUpdateParams, opts ...option.RequestOption) (res *AccountTeamnetVirtualNetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountIdentifier, vnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing virtual network.
func (r *AccountTeamnetVirtualNetworkService) Delete(ctx context.Context, accountIdentifier string, vnetID string, opts ...option.RequestOption) (res *AccountTeamnetVirtualNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountIdentifier, vnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new virtual network to an account.
func (r *AccountTeamnetVirtualNetworkService) TunnelVirtualNetworkNewAVirtualNetwork(ctx context.Context, accountIdentifier string, body AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkParams, opts ...option.RequestOption) (res *AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters virtual networks in an account.
func (r *AccountTeamnetVirtualNetworkService) TunnelVirtualNetworkListVirtualNetworks(ctx context.Context, accountIdentifier string, query AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams, opts ...option.RequestOption) (res *AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountTeamnetVirtualNetworkUpdateResponse struct {
	Errors   []AccountTeamnetVirtualNetworkUpdateResponseError   `json:"errors"`
	Messages []AccountTeamnetVirtualNetworkUpdateResponseMessage `json:"messages"`
	Result   interface{}                                         `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetVirtualNetworkUpdateResponseSuccess `json:"success"`
	JSON    accountTeamnetVirtualNetworkUpdateResponseJSON    `json:"-"`
}

// accountTeamnetVirtualNetworkUpdateResponseJSON contains the JSON metadata for
// the struct [AccountTeamnetVirtualNetworkUpdateResponse]
type accountTeamnetVirtualNetworkUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkUpdateResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkUpdateResponseErrorJSON `json:"-"`
}

// accountTeamnetVirtualNetworkUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountTeamnetVirtualNetworkUpdateResponseError]
type accountTeamnetVirtualNetworkUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkUpdateResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkUpdateResponseMessageJSON `json:"-"`
}

// accountTeamnetVirtualNetworkUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountTeamnetVirtualNetworkUpdateResponseMessage]
type accountTeamnetVirtualNetworkUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetVirtualNetworkUpdateResponseSuccess bool

const (
	AccountTeamnetVirtualNetworkUpdateResponseSuccessTrue AccountTeamnetVirtualNetworkUpdateResponseSuccess = true
)

type AccountTeamnetVirtualNetworkDeleteResponse struct {
	Errors   []AccountTeamnetVirtualNetworkDeleteResponseError   `json:"errors"`
	Messages []AccountTeamnetVirtualNetworkDeleteResponseMessage `json:"messages"`
	Result   interface{}                                         `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetVirtualNetworkDeleteResponseSuccess `json:"success"`
	JSON    accountTeamnetVirtualNetworkDeleteResponseJSON    `json:"-"`
}

// accountTeamnetVirtualNetworkDeleteResponseJSON contains the JSON metadata for
// the struct [AccountTeamnetVirtualNetworkDeleteResponse]
type accountTeamnetVirtualNetworkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkDeleteResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkDeleteResponseErrorJSON `json:"-"`
}

// accountTeamnetVirtualNetworkDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountTeamnetVirtualNetworkDeleteResponseError]
type accountTeamnetVirtualNetworkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkDeleteResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkDeleteResponseMessageJSON `json:"-"`
}

// accountTeamnetVirtualNetworkDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountTeamnetVirtualNetworkDeleteResponseMessage]
type accountTeamnetVirtualNetworkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetVirtualNetworkDeleteResponseSuccess bool

const (
	AccountTeamnetVirtualNetworkDeleteResponseSuccessTrue AccountTeamnetVirtualNetworkDeleteResponseSuccess = true
)

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse struct {
	Errors   []AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseError   `json:"errors"`
	Messages []AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseMessage `json:"messages"`
	Result   interface{}                                                                         `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseSuccess `json:"success"`
	JSON    accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseJSON    `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseErrorJSON `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseError]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseMessageJSON `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseMessage]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseSuccess bool

const (
	AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseSuccessTrue AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseSuccess = true
)

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse struct {
	Errors     []AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseError    `json:"errors"`
	Messages   []AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseMessage  `json:"messages"`
	Result     []AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResult   `json:"result"`
	ResultInfo AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseSuccess `json:"success"`
	JSON    accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseJSON    `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseErrorJSON `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseError]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseMessageJSON `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseMessage]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResult struct {
	// UUID of the virtual network.
	ID string `json:"id,required"`
	// Optional remark describing the virtual network.
	Comment string `json:"comment,required"`
	// Timestamp of when the virtual network was created.
	CreatedAt interface{} `json:"created_at,required"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork bool `json:"is_default_network,required"`
	// A user-friendly name for the virtual network.
	Name string `json:"name,required"`
	// Timestamp of when the virtual network was deleted. If `null`, the virtual
	// network has not been deleted.
	DeletedAt interface{}                                                                           `json:"deleted_at"`
	JSON      accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultJSON `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResult]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	DeletedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                   `json:"total_count"`
	JSON       accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultInfoJSON `json:"-"`
}

// accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultInfo]
type accountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseSuccess bool

const (
	AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseSuccessTrue AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseSuccess = true
)

type AccountTeamnetVirtualNetworkUpdateParams struct {
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name"`
}

func (r AccountTeamnetVirtualNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkParams struct {
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefault param.Field[bool] `json:"is_default"`
}

func (r AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams struct {
	// If `true`, only include the default virtual network. If `false`, exclude the
	// default virtual network. If empty, all virtual networks will be included.
	IsDefault param.Field[interface{}] `query:"is_default"`
	// If `true`, only include deleted virtual networks. If `false`, exclude deleted
	// virtual networks. If empty, all virtual networks will be included.
	IsDeleted param.Field[interface{}] `query:"is_deleted"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `query:"name"`
	// A user-friendly name for the virtual network.
	VnetName param.Field[string] `query:"vnet_name"`
}

// URLQuery serializes
// [AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams]'s
// query parameters as `url.Values`.
func (r AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
