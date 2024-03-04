// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccountRoleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRoleService] method
// instead.
type AccountRoleService struct {
	Options []option.RequestOption
}

// NewAccountRoleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRoleService(opts ...option.RequestOption) (r *AccountRoleService) {
	r = &AccountRoleService{}
	r.Options = opts
	return
}

// Get all available roles for an account.
func (r *AccountRoleService) List(ctx context.Context, query AccountRoleListParams, opts ...option.RequestOption) (res *[]AccountRoleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountRoleListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/roles", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific role for an account.
func (r *AccountRoleService) Get(ctx context.Context, roleID interface{}, query AccountRoleGetParams, opts ...option.RequestOption) (res *AccountRoleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountRoleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/roles/%v", query.AccountID, roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountRoleListResponse struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                    `json:"permissions,required"`
	JSON        accountRoleListResponseJSON `json:"-"`
}

// accountRoleListResponseJSON contains the JSON metadata for the struct
// [AccountRoleListResponse]
type accountRoleListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountRoleGetResponseUnknown] or [shared.UnionString].
type AccountRoleGetResponse interface {
	ImplementsAccountRoleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountRoleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountRoleListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type AccountRoleListResponseEnvelope struct {
	Errors   []AccountRoleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountRoleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccountRoleListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccountRoleListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccountRoleListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accountRoleListResponseEnvelopeJSON       `json:"-"`
}

// accountRoleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountRoleListResponseEnvelope]
type accountRoleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountRoleListResponseEnvelopeErrorsJSON `json:"-"`
}

// accountRoleListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountRoleListResponseEnvelopeErrors]
type accountRoleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountRoleListResponseEnvelopeMessagesJSON `json:"-"`
}

// accountRoleListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountRoleListResponseEnvelopeMessages]
type accountRoleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRoleListResponseEnvelopeSuccess bool

const (
	AccountRoleListResponseEnvelopeSuccessTrue AccountRoleListResponseEnvelopeSuccess = true
)

type AccountRoleListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       accountRoleListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accountRoleListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccountRoleListResponseEnvelopeResultInfo]
type accountRoleListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type AccountRoleGetResponseEnvelope struct {
	Errors   []AccountRoleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountRoleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountRoleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountRoleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountRoleGetResponseEnvelopeJSON    `json:"-"`
}

// accountRoleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountRoleGetResponseEnvelope]
type accountRoleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountRoleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accountRoleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountRoleGetResponseEnvelopeErrors]
type accountRoleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRoleGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountRoleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accountRoleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountRoleGetResponseEnvelopeMessages]
type accountRoleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRoleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRoleGetResponseEnvelopeSuccess bool

const (
	AccountRoleGetResponseEnvelopeSuccessTrue AccountRoleGetResponseEnvelopeSuccess = true
)
