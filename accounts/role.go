// File generated from our OpenAPI spec by Stainless.

package accounts

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// RoleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRoleService] method instead.
type RoleService struct {
	Options []option.RequestOption
}

// NewRoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoleService(opts ...option.RequestOption) (r *RoleService) {
	r = &RoleService{}
	r.Options = opts
	return
}

// Get all available roles for an account.
func (r *RoleService) List(ctx context.Context, query RoleListParams, opts ...option.RequestOption) (res *[]IamSchemasRole, err error) {
	opts = append(r.Options[:], opts...)
	var env RoleListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/roles", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific role for an account.
func (r *RoleService) Get(ctx context.Context, roleID interface{}, query RoleGetParams, opts ...option.RequestOption) (res *RoleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/roles/%v", query.AccountID, roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IamSchemasRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string           `json:"permissions,required"`
	JSON        iamSchemasRoleJSON `json:"-"`
}

// iamSchemasRoleJSON contains the JSON metadata for the struct [IamSchemasRole]
type iamSchemasRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSchemasRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSchemasRoleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [accounts.RoleGetResponseUnknown] or [shared.UnionString].
type RoleGetResponse interface {
	ImplementsAccountsRoleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RoleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RoleListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type RoleListResponseEnvelope struct {
	Errors   []RoleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IamSchemasRole                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RoleListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RoleListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       roleListResponseEnvelopeJSON       `json:"-"`
}

// roleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoleListResponseEnvelope]
type roleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoleListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    roleListResponseEnvelopeErrorsJSON `json:"-"`
}

// roleListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RoleListResponseEnvelopeErrors]
type roleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoleListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    roleListResponseEnvelopeMessagesJSON `json:"-"`
}

// roleListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RoleListResponseEnvelopeMessages]
type roleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoleListResponseEnvelopeSuccess bool

const (
	RoleListResponseEnvelopeSuccessTrue RoleListResponseEnvelopeSuccess = true
)

type RoleListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       roleListResponseEnvelopeResultInfoJSON `json:"-"`
}

// roleListResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RoleListResponseEnvelopeResultInfo]
type roleListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RoleGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type RoleGetResponseEnvelope struct {
	Errors   []RoleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RoleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    roleGetResponseEnvelopeJSON    `json:"-"`
}

// roleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoleGetResponseEnvelope]
type roleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoleGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    roleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// roleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RoleGetResponseEnvelopeErrors]
type roleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoleGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    roleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// roleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RoleGetResponseEnvelopeMessages]
type roleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoleGetResponseEnvelopeSuccess bool

const (
	RoleGetResponseEnvelopeSuccessTrue RoleGetResponseEnvelopeSuccess = true
)
