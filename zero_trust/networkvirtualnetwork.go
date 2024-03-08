// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// NetworkVirtualNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNetworkVirtualNetworkService]
// method instead.
type NetworkVirtualNetworkService struct {
	Options []option.RequestOption
}

// NewNetworkVirtualNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkVirtualNetworkService(opts ...option.RequestOption) (r *NetworkVirtualNetworkService) {
	r = &NetworkVirtualNetworkService{}
	r.Options = opts
	return
}

// Adds a new virtual network to an account.
func (r *NetworkVirtualNetworkService) New(ctx context.Context, params NetworkVirtualNetworkNewParams, opts ...option.RequestOption) (res *NetworkVirtualNetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkVirtualNetworkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters virtual networks in an account.
func (r *NetworkVirtualNetworkService) List(ctx context.Context, params NetworkVirtualNetworkListParams, opts ...option.RequestOption) (res *[]TunnelVirtualNetwork, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkVirtualNetworkListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing virtual network.
func (r *NetworkVirtualNetworkService) Delete(ctx context.Context, virtualNetworkID string, body NetworkVirtualNetworkDeleteParams, opts ...option.RequestOption) (res *NetworkVirtualNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkVirtualNetworkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", body.AccountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing virtual network.
func (r *NetworkVirtualNetworkService) Edit(ctx context.Context, virtualNetworkID string, params NetworkVirtualNetworkEditParams, opts ...option.RequestOption) (res *NetworkVirtualNetworkEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetworkVirtualNetworkEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", params.AccountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelVirtualNetwork struct {
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
	DeletedAt interface{}              `json:"deleted_at"`
	JSON      tunnelVirtualNetworkJSON `json:"-"`
}

// tunnelVirtualNetworkJSON contains the JSON metadata for the struct
// [TunnelVirtualNetwork]
type tunnelVirtualNetworkJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	DeletedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelVirtualNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelVirtualNetworkJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.NetworkVirtualNetworkNewResponseUnknown],
// [zero_trust.NetworkVirtualNetworkNewResponseArray] or [shared.UnionString].
type NetworkVirtualNetworkNewResponse interface {
	ImplementsZeroTrustNetworkVirtualNetworkNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NetworkVirtualNetworkNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkVirtualNetworkNewResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NetworkVirtualNetworkNewResponseArray []interface{}

func (r NetworkVirtualNetworkNewResponseArray) ImplementsZeroTrustNetworkVirtualNetworkNewResponse() {
}

// Union satisfied by [zero_trust.NetworkVirtualNetworkDeleteResponseUnknown],
// [zero_trust.NetworkVirtualNetworkDeleteResponseArray] or [shared.UnionString].
type NetworkVirtualNetworkDeleteResponse interface {
	ImplementsZeroTrustNetworkVirtualNetworkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NetworkVirtualNetworkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkVirtualNetworkDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NetworkVirtualNetworkDeleteResponseArray []interface{}

func (r NetworkVirtualNetworkDeleteResponseArray) ImplementsZeroTrustNetworkVirtualNetworkDeleteResponse() {
}

// Union satisfied by [zero_trust.NetworkVirtualNetworkEditResponseUnknown],
// [zero_trust.NetworkVirtualNetworkEditResponseArray] or [shared.UnionString].
type NetworkVirtualNetworkEditResponse interface {
	ImplementsZeroTrustNetworkVirtualNetworkEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NetworkVirtualNetworkEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkVirtualNetworkEditResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NetworkVirtualNetworkEditResponseArray []interface{}

func (r NetworkVirtualNetworkEditResponseArray) ImplementsZeroTrustNetworkVirtualNetworkEditResponse() {
}

type NetworkVirtualNetworkNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefault param.Field[bool] `json:"is_default"`
}

func (r NetworkVirtualNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkVirtualNetworkNewResponseEnvelope struct {
	Errors   []NetworkVirtualNetworkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkVirtualNetworkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   NetworkVirtualNetworkNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NetworkVirtualNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkVirtualNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// networkVirtualNetworkNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkVirtualNetworkNewResponseEnvelope]
type networkVirtualNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    networkVirtualNetworkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// networkVirtualNetworkNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [NetworkVirtualNetworkNewResponseEnvelopeErrors]
type networkVirtualNetworkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    networkVirtualNetworkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// networkVirtualNetworkNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [NetworkVirtualNetworkNewResponseEnvelopeMessages]
type networkVirtualNetworkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkVirtualNetworkNewResponseEnvelopeSuccess bool

const (
	NetworkVirtualNetworkNewResponseEnvelopeSuccessTrue NetworkVirtualNetworkNewResponseEnvelopeSuccess = true
)

type NetworkVirtualNetworkListParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
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

// URLQuery serializes [NetworkVirtualNetworkListParams]'s query parameters as
// `url.Values`.
func (r NetworkVirtualNetworkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetworkVirtualNetworkListResponseEnvelope struct {
	Errors   []NetworkVirtualNetworkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkVirtualNetworkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []TunnelVirtualNetwork                              `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    NetworkVirtualNetworkListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo NetworkVirtualNetworkListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       networkVirtualNetworkListResponseEnvelopeJSON       `json:"-"`
}

// networkVirtualNetworkListResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkVirtualNetworkListResponseEnvelope]
type networkVirtualNetworkListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    networkVirtualNetworkListResponseEnvelopeErrorsJSON `json:"-"`
}

// networkVirtualNetworkListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [NetworkVirtualNetworkListResponseEnvelopeErrors]
type networkVirtualNetworkListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    networkVirtualNetworkListResponseEnvelopeMessagesJSON `json:"-"`
}

// networkVirtualNetworkListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [NetworkVirtualNetworkListResponseEnvelopeMessages]
type networkVirtualNetworkListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkVirtualNetworkListResponseEnvelopeSuccess bool

const (
	NetworkVirtualNetworkListResponseEnvelopeSuccessTrue NetworkVirtualNetworkListResponseEnvelopeSuccess = true
)

type NetworkVirtualNetworkListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       networkVirtualNetworkListResponseEnvelopeResultInfoJSON `json:"-"`
}

// networkVirtualNetworkListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [NetworkVirtualNetworkListResponseEnvelopeResultInfo]
type networkVirtualNetworkListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkVirtualNetworkDeleteResponseEnvelope struct {
	Errors   []NetworkVirtualNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkVirtualNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   NetworkVirtualNetworkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NetworkVirtualNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkVirtualNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// networkVirtualNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [NetworkVirtualNetworkDeleteResponseEnvelope]
type networkVirtualNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    networkVirtualNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// networkVirtualNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [NetworkVirtualNetworkDeleteResponseEnvelopeErrors]
type networkVirtualNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    networkVirtualNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// networkVirtualNetworkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [NetworkVirtualNetworkDeleteResponseEnvelopeMessages]
type networkVirtualNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkVirtualNetworkDeleteResponseEnvelopeSuccess bool

const (
	NetworkVirtualNetworkDeleteResponseEnvelopeSuccessTrue NetworkVirtualNetworkDeleteResponseEnvelopeSuccess = true
)

type NetworkVirtualNetworkEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name"`
}

func (r NetworkVirtualNetworkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkVirtualNetworkEditResponseEnvelope struct {
	Errors   []NetworkVirtualNetworkEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NetworkVirtualNetworkEditResponseEnvelopeMessages `json:"messages,required"`
	Result   NetworkVirtualNetworkEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success NetworkVirtualNetworkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkVirtualNetworkEditResponseEnvelopeJSON    `json:"-"`
}

// networkVirtualNetworkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkVirtualNetworkEditResponseEnvelope]
type networkVirtualNetworkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    networkVirtualNetworkEditResponseEnvelopeErrorsJSON `json:"-"`
}

// networkVirtualNetworkEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [NetworkVirtualNetworkEditResponseEnvelopeErrors]
type networkVirtualNetworkEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NetworkVirtualNetworkEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    networkVirtualNetworkEditResponseEnvelopeMessagesJSON `json:"-"`
}

// networkVirtualNetworkEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [NetworkVirtualNetworkEditResponseEnvelopeMessages]
type networkVirtualNetworkEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkVirtualNetworkEditResponseEnvelopeSuccess bool

const (
	NetworkVirtualNetworkEditResponseEnvelopeSuccessTrue NetworkVirtualNetworkEditResponseEnvelopeSuccess = true
)
