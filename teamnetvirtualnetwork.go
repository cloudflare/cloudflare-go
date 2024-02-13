// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// TeamnetVirtualNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTeamnetVirtualNetworkService]
// method instead.
type TeamnetVirtualNetworkService struct {
	Options []option.RequestOption
}

// NewTeamnetVirtualNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTeamnetVirtualNetworkService(opts ...option.RequestOption) (r *TeamnetVirtualNetworkService) {
	r = &TeamnetVirtualNetworkService{}
	r.Options = opts
	return
}

// Updates an existing virtual network.
func (r *TeamnetVirtualNetworkService) Update(ctx context.Context, accountID string, virtualNetworkID string, body TeamnetVirtualNetworkUpdateParams, opts ...option.RequestOption) (res *TeamnetVirtualNetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetVirtualNetworkUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing virtual network.
func (r *TeamnetVirtualNetworkService) Delete(ctx context.Context, accountID string, virtualNetworkID string, opts ...option.RequestOption) (res *TeamnetVirtualNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetVirtualNetworkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds a new virtual network to an account.
func (r *TeamnetVirtualNetworkService) TunnelVirtualNetworkNewAVirtualNetwork(ctx context.Context, accountID string, body TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkParams, opts ...option.RequestOption) (res *TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters virtual networks in an account.
func (r *TeamnetVirtualNetworkService) TunnelVirtualNetworkListVirtualNetworks(ctx context.Context, accountID string, query TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams, opts ...option.RequestOption) (res *[]TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [TeamnetVirtualNetworkUpdateResponseUnknown],
// [TeamnetVirtualNetworkUpdateResponseArray] or [shared.UnionString].
type TeamnetVirtualNetworkUpdateResponse interface {
	ImplementsTeamnetVirtualNetworkUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamnetVirtualNetworkUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TeamnetVirtualNetworkUpdateResponseArray []interface{}

func (r TeamnetVirtualNetworkUpdateResponseArray) ImplementsTeamnetVirtualNetworkUpdateResponse() {}

// Union satisfied by [TeamnetVirtualNetworkDeleteResponseUnknown],
// [TeamnetVirtualNetworkDeleteResponseArray] or [shared.UnionString].
type TeamnetVirtualNetworkDeleteResponse interface {
	ImplementsTeamnetVirtualNetworkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamnetVirtualNetworkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TeamnetVirtualNetworkDeleteResponseArray []interface{}

func (r TeamnetVirtualNetworkDeleteResponseArray) ImplementsTeamnetVirtualNetworkDeleteResponse() {}

// Union satisfied by
// [TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseUnknown],
// [TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseArray] or
// [shared.UnionString].
type TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse interface {
	ImplementsTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseArray []interface{}

func (r TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseArray) ImplementsTeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse() {
}

type TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse struct {
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
	DeletedAt interface{}                                                              `json:"deleted_at"`
	JSON      teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseJSON `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse]
type teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	DeletedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkUpdateParams struct {
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name"`
}

func (r TeamnetVirtualNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamnetVirtualNetworkUpdateResponseEnvelope struct {
	Errors   []TeamnetVirtualNetworkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetVirtualNetworkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetVirtualNetworkUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetVirtualNetworkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetVirtualNetworkUpdateResponseEnvelopeJSON    `json:"-"`
}

// teamnetVirtualNetworkUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [TeamnetVirtualNetworkUpdateResponseEnvelope]
type teamnetVirtualNetworkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkUpdateResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    teamnetVirtualNetworkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetVirtualNetworkUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [TeamnetVirtualNetworkUpdateResponseEnvelopeErrors]
type teamnetVirtualNetworkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkUpdateResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    teamnetVirtualNetworkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetVirtualNetworkUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [TeamnetVirtualNetworkUpdateResponseEnvelopeMessages]
type teamnetVirtualNetworkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetVirtualNetworkUpdateResponseEnvelopeSuccess bool

const (
	TeamnetVirtualNetworkUpdateResponseEnvelopeSuccessTrue TeamnetVirtualNetworkUpdateResponseEnvelopeSuccess = true
)

type TeamnetVirtualNetworkDeleteResponseEnvelope struct {
	Errors   []TeamnetVirtualNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetVirtualNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetVirtualNetworkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetVirtualNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetVirtualNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// teamnetVirtualNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [TeamnetVirtualNetworkDeleteResponseEnvelope]
type teamnetVirtualNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    teamnetVirtualNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetVirtualNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [TeamnetVirtualNetworkDeleteResponseEnvelopeErrors]
type teamnetVirtualNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    teamnetVirtualNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetVirtualNetworkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [TeamnetVirtualNetworkDeleteResponseEnvelopeMessages]
type teamnetVirtualNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetVirtualNetworkDeleteResponseEnvelopeSuccess bool

const (
	TeamnetVirtualNetworkDeleteResponseEnvelopeSuccessTrue TeamnetVirtualNetworkDeleteResponseEnvelopeSuccess = true
)

type TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkParams struct {
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefault param.Field[bool] `json:"is_default"`
}

func (r TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelope struct {
	Errors   []TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeMessages `json:"messages,required"`
	Result   TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeSuccess `json:"success,required"`
	JSON    teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeJSON    `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelope]
type teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeErrors]
type teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeMessages]
type teamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeSuccess bool

const (
	TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeSuccessTrue TeamnetVirtualNetworkTunnelVirtualNetworkNewAVirtualNetworkResponseEnvelopeSuccess = true
)

type TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams struct {
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
// [TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams]'s query
// parameters as `url.Values`.
func (r TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelope struct {
	Errors   []TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeMessages `json:"messages,required"`
	Result   []TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeResultInfo `json:"result_info"`
	JSON       teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeJSON       `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelope]
type teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeErrors struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeErrorsJSON `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeErrors]
type teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeMessages struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeMessagesJSON `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeMessages]
type teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeSuccess bool

const (
	TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeSuccessTrue TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeSuccess = true
)

type TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                    `json:"total_count"`
	JSON       teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeResultInfoJSON `json:"-"`
}

// teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeResultInfo]
type teamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamnetVirtualNetworkTunnelVirtualNetworkListVirtualNetworksResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
