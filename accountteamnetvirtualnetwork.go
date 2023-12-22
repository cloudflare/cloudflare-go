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
func (r *AccountTeamnetVirtualNetworkService) Update(ctx context.Context, accountIdentifier string, vnetID string, body AccountTeamnetVirtualNetworkUpdateParams, opts ...option.RequestOption) (res *VnetResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountIdentifier, vnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing virtual network.
func (r *AccountTeamnetVirtualNetworkService) Delete(ctx context.Context, accountIdentifier string, vnetID string, opts ...option.RequestOption) (res *VnetResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountIdentifier, vnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new virtual network to an account.
func (r *AccountTeamnetVirtualNetworkService) TunnelVirtualNetworkNewAVirtualNetwork(ctx context.Context, accountIdentifier string, body AccountTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkParams, opts ...option.RequestOption) (res *VnetResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists and filters virtual networks in an account.
func (r *AccountTeamnetVirtualNetworkService) TunnelVirtualNetworkListVirtualNetworks(ctx context.Context, accountIdentifier string, query AccountTeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams, opts ...option.RequestOption) (res *VnetResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VnetResponseCollection struct {
	Errors     []VnetResponseCollectionError    `json:"errors"`
	Messages   []VnetResponseCollectionMessage  `json:"messages"`
	Result     []VnetResponseCollectionResult   `json:"result"`
	ResultInfo VnetResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success VnetResponseCollectionSuccess `json:"success"`
	JSON    vnetResponseCollectionJSON    `json:"-"`
}

// vnetResponseCollectionJSON contains the JSON metadata for the struct
// [VnetResponseCollection]
type vnetResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VnetResponseCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    vnetResponseCollectionErrorJSON `json:"-"`
}

// vnetResponseCollectionErrorJSON contains the JSON metadata for the struct
// [VnetResponseCollectionError]
type vnetResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VnetResponseCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    vnetResponseCollectionMessageJSON `json:"-"`
}

// vnetResponseCollectionMessageJSON contains the JSON metadata for the struct
// [VnetResponseCollectionMessage]
type vnetResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VnetResponseCollectionResult struct {
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
	DeletedAt interface{}                      `json:"deleted_at"`
	JSON      vnetResponseCollectionResultJSON `json:"-"`
}

// vnetResponseCollectionResultJSON contains the JSON metadata for the struct
// [VnetResponseCollectionResult]
type vnetResponseCollectionResultJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	DeletedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VnetResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VnetResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       vnetResponseCollectionResultInfoJSON `json:"-"`
}

// vnetResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [VnetResponseCollectionResultInfo]
type vnetResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VnetResponseCollectionSuccess bool

const (
	VnetResponseCollectionSuccessTrue VnetResponseCollectionSuccess = true
)

type VnetResponseSingle struct {
	Errors   []VnetResponseSingleError   `json:"errors"`
	Messages []VnetResponseSingleMessage `json:"messages"`
	Result   interface{}                 `json:"result"`
	// Whether the API call was successful
	Success VnetResponseSingleSuccess `json:"success"`
	JSON    vnetResponseSingleJSON    `json:"-"`
}

// vnetResponseSingleJSON contains the JSON metadata for the struct
// [VnetResponseSingle]
type vnetResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VnetResponseSingleError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    vnetResponseSingleErrorJSON `json:"-"`
}

// vnetResponseSingleErrorJSON contains the JSON metadata for the struct
// [VnetResponseSingleError]
type vnetResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VnetResponseSingleMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    vnetResponseSingleMessageJSON `json:"-"`
}

// vnetResponseSingleMessageJSON contains the JSON metadata for the struct
// [VnetResponseSingleMessage]
type vnetResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VnetResponseSingleSuccess bool

const (
	VnetResponseSingleSuccessTrue VnetResponseSingleSuccess = true
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
