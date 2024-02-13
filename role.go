// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *RoleService) AccountRolesListRoles(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]RoleAccountRolesListRolesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoleAccountRolesListRolesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/roles", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific role for an account.
func (r *RoleService) Get(ctx context.Context, accountID interface{}, roleID interface{}, opts ...option.RequestOption) (res *RoleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/roles/%v", accountID, roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoleAccountRolesListRolesResponse struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                              `json:"permissions,required"`
	JSON        roleAccountRolesListRolesResponseJSON `json:"-"`
}

// roleAccountRolesListRolesResponseJSON contains the JSON metadata for the struct
// [RoleAccountRolesListRolesResponse]
type roleAccountRolesListRolesResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleAccountRolesListRolesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [RoleGetResponseUnknown] or [shared.UnionString].
type RoleGetResponse interface {
	ImplementsRoleGetResponse()
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

type RoleAccountRolesListRolesResponseEnvelope struct {
	Errors   []RoleAccountRolesListRolesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoleAccountRolesListRolesResponseEnvelopeMessages `json:"messages,required"`
	Result   []RoleAccountRolesListRolesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RoleAccountRolesListRolesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RoleAccountRolesListRolesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       roleAccountRolesListRolesResponseEnvelopeJSON       `json:"-"`
}

// roleAccountRolesListRolesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RoleAccountRolesListRolesResponseEnvelope]
type roleAccountRolesListRolesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleAccountRolesListRolesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoleAccountRolesListRolesResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    roleAccountRolesListRolesResponseEnvelopeErrorsJSON `json:"-"`
}

// roleAccountRolesListRolesResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [RoleAccountRolesListRolesResponseEnvelopeErrors]
type roleAccountRolesListRolesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleAccountRolesListRolesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoleAccountRolesListRolesResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    roleAccountRolesListRolesResponseEnvelopeMessagesJSON `json:"-"`
}

// roleAccountRolesListRolesResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RoleAccountRolesListRolesResponseEnvelopeMessages]
type roleAccountRolesListRolesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleAccountRolesListRolesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RoleAccountRolesListRolesResponseEnvelopeSuccess bool

const (
	RoleAccountRolesListRolesResponseEnvelopeSuccessTrue RoleAccountRolesListRolesResponseEnvelopeSuccess = true
)

type RoleAccountRolesListRolesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       roleAccountRolesListRolesResponseEnvelopeResultInfoJSON `json:"-"`
}

// roleAccountRolesListRolesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [RoleAccountRolesListRolesResponseEnvelopeResultInfo]
type roleAccountRolesListRolesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleAccountRolesListRolesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Whether the API call was successful
type RoleGetResponseEnvelopeSuccess bool

const (
	RoleGetResponseEnvelopeSuccessTrue RoleGetResponseEnvelopeSuccess = true
)
