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

// ZeroTrustNetworkVirtualNetworkService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustNetworkVirtualNetworkService] method instead.
type ZeroTrustNetworkVirtualNetworkService struct {
	Options []option.RequestOption
}

// NewZeroTrustNetworkVirtualNetworkService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustNetworkVirtualNetworkService(opts ...option.RequestOption) (r *ZeroTrustNetworkVirtualNetworkService) {
	r = &ZeroTrustNetworkVirtualNetworkService{}
	r.Options = opts
	return
}

// Adds a new virtual network to an account.
func (r *ZeroTrustNetworkVirtualNetworkService) New(ctx context.Context, params ZeroTrustNetworkVirtualNetworkNewParams, opts ...option.RequestOption) (res *ZeroTrustNetworkVirtualNetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkVirtualNetworkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters virtual networks in an account.
func (r *ZeroTrustNetworkVirtualNetworkService) List(ctx context.Context, params ZeroTrustNetworkVirtualNetworkListParams, opts ...option.RequestOption) (res *[]ZeroTrustNetworkVirtualNetworkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkVirtualNetworkListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing virtual network.
func (r *ZeroTrustNetworkVirtualNetworkService) Delete(ctx context.Context, virtualNetworkID string, body ZeroTrustNetworkVirtualNetworkDeleteParams, opts ...option.RequestOption) (res *ZeroTrustNetworkVirtualNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", body.AccountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing virtual network.
func (r *ZeroTrustNetworkVirtualNetworkService) Edit(ctx context.Context, virtualNetworkID string, params ZeroTrustNetworkVirtualNetworkEditParams, opts ...option.RequestOption) (res *ZeroTrustNetworkVirtualNetworkEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustNetworkVirtualNetworkEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", params.AccountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustNetworkVirtualNetworkNewResponseUnknown],
// [ZeroTrustNetworkVirtualNetworkNewResponseArray] or [shared.UnionString].
type ZeroTrustNetworkVirtualNetworkNewResponse interface {
	ImplementsZeroTrustNetworkVirtualNetworkNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustNetworkVirtualNetworkNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustNetworkVirtualNetworkNewResponseArray []interface{}

func (r ZeroTrustNetworkVirtualNetworkNewResponseArray) ImplementsZeroTrustNetworkVirtualNetworkNewResponse() {
}

type ZeroTrustNetworkVirtualNetworkListResponse struct {
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
	DeletedAt interface{}                                    `json:"deleted_at"`
	JSON      zeroTrustNetworkVirtualNetworkListResponseJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkListResponseJSON contains the JSON metadata for
// the struct [ZeroTrustNetworkVirtualNetworkListResponse]
type zeroTrustNetworkVirtualNetworkListResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	DeletedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZeroTrustNetworkVirtualNetworkDeleteResponseUnknown],
// [ZeroTrustNetworkVirtualNetworkDeleteResponseArray] or [shared.UnionString].
type ZeroTrustNetworkVirtualNetworkDeleteResponse interface {
	ImplementsZeroTrustNetworkVirtualNetworkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustNetworkVirtualNetworkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustNetworkVirtualNetworkDeleteResponseArray []interface{}

func (r ZeroTrustNetworkVirtualNetworkDeleteResponseArray) ImplementsZeroTrustNetworkVirtualNetworkDeleteResponse() {
}

// Union satisfied by [ZeroTrustNetworkVirtualNetworkEditResponseUnknown],
// [ZeroTrustNetworkVirtualNetworkEditResponseArray] or [shared.UnionString].
type ZeroTrustNetworkVirtualNetworkEditResponse interface {
	ImplementsZeroTrustNetworkVirtualNetworkEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustNetworkVirtualNetworkEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustNetworkVirtualNetworkEditResponseArray []interface{}

func (r ZeroTrustNetworkVirtualNetworkEditResponseArray) ImplementsZeroTrustNetworkVirtualNetworkEditResponse() {
}

type ZeroTrustNetworkVirtualNetworkNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefault param.Field[bool] `json:"is_default"`
}

func (r ZeroTrustNetworkVirtualNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustNetworkVirtualNetworkNewResponseEnvelope struct {
	Errors   []ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustNetworkVirtualNetworkNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkVirtualNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkVirtualNetworkNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustNetworkVirtualNetworkNewResponseEnvelope]
type zeroTrustNetworkVirtualNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeErrors]
type zeroTrustNetworkVirtualNetworkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeMessages]
type zeroTrustNetworkVirtualNetworkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeSuccessTrue ZeroTrustNetworkVirtualNetworkNewResponseEnvelopeSuccess = true
)

type ZeroTrustNetworkVirtualNetworkListParams struct {
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

// URLQuery serializes [ZeroTrustNetworkVirtualNetworkListParams]'s query
// parameters as `url.Values`.
func (r ZeroTrustNetworkVirtualNetworkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZeroTrustNetworkVirtualNetworkListResponseEnvelope struct {
	Errors   []ZeroTrustNetworkVirtualNetworkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkVirtualNetworkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustNetworkVirtualNetworkListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustNetworkVirtualNetworkListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustNetworkVirtualNetworkListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustNetworkVirtualNetworkListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustNetworkVirtualNetworkListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkVirtualNetworkListResponseEnvelope]
type zeroTrustNetworkVirtualNetworkListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkListResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkVirtualNetworkListResponseEnvelopeErrors]
type zeroTrustNetworkVirtualNetworkListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkListResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkVirtualNetworkListResponseEnvelopeMessages]
type zeroTrustNetworkVirtualNetworkListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkVirtualNetworkListResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkVirtualNetworkListResponseEnvelopeSuccessTrue ZeroTrustNetworkVirtualNetworkListResponseEnvelopeSuccess = true
)

type ZeroTrustNetworkVirtualNetworkListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       zeroTrustNetworkVirtualNetworkListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustNetworkVirtualNetworkListResponseEnvelopeResultInfo]
type zeroTrustNetworkVirtualNetworkListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelope struct {
	Errors   []ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustNetworkVirtualNetworkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelope]
type zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeErrors]
type zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeMessages]
type zeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeSuccessTrue ZeroTrustNetworkVirtualNetworkDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustNetworkVirtualNetworkEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name"`
}

func (r ZeroTrustNetworkVirtualNetworkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustNetworkVirtualNetworkEditResponseEnvelope struct {
	Errors   []ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustNetworkVirtualNetworkEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustNetworkVirtualNetworkEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustNetworkVirtualNetworkEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustNetworkVirtualNetworkEditResponseEnvelope]
type zeroTrustNetworkVirtualNetworkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeErrors]
type zeroTrustNetworkVirtualNetworkEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustNetworkVirtualNetworkEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustNetworkVirtualNetworkEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeMessages]
type zeroTrustNetworkVirtualNetworkEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeSuccess bool

const (
	ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeSuccessTrue ZeroTrustNetworkVirtualNetworkEditResponseEnvelopeSuccess = true
)
