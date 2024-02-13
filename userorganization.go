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

// UserOrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserOrganizationService] method
// instead.
type UserOrganizationService struct {
	Options []option.RequestOption
}

// NewUserOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserOrganizationService(opts ...option.RequestOption) (r *UserOrganizationService) {
	r = &UserOrganizationService{}
	r.Options = opts
	return
}

// Removes association to an organization.
func (r *UserOrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *UserOrganizationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets a specific organization the user is associated with.
func (r *UserOrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *UserOrganizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserOrganizationGetResponseEnvelope
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists organizations the user is associated with.
func (r *UserOrganizationService) UserSOrganizationsListOrganizations(ctx context.Context, query UserOrganizationUserSOrganizationsListOrganizationsParams, opts ...option.RequestOption) (res *[]UserOrganizationUserSOrganizationsListOrganizationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelope
	path := "user/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserOrganizationDeleteResponse struct {
	// Identifier
	ID   string                             `json:"id"`
	JSON userOrganizationDeleteResponseJSON `json:"-"`
}

// userOrganizationDeleteResponseJSON contains the JSON metadata for the struct
// [UserOrganizationDeleteResponse]
type userOrganizationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [UserOrganizationGetResponseUnknown] or [shared.UnionString].
type UserOrganizationGetResponse interface {
	ImplementsUserOrganizationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserOrganizationGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserOrganizationUserSOrganizationsListOrganizationsResponse struct {
	// Identifier
	ID string `json:"id"`
	// Organization name.
	Name string `json:"name"`
	// Access permissions for this User.
	Permissions []string `json:"permissions"`
	// List of roles that a user has within an organization.
	Roles []string `json:"roles"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status UserOrganizationUserSOrganizationsListOrganizationsResponseStatus `json:"status"`
	JSON   userOrganizationUserSOrganizationsListOrganizationsResponseJSON   `json:"-"`
}

// userOrganizationUserSOrganizationsListOrganizationsResponseJSON contains the
// JSON metadata for the struct
// [UserOrganizationUserSOrganizationsListOrganizationsResponse]
type userOrganizationUserSOrganizationsListOrganizationsResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationUserSOrganizationsListOrganizationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the user is a member of the organization or has an inivitation pending.
type UserOrganizationUserSOrganizationsListOrganizationsResponseStatus string

const (
	UserOrganizationUserSOrganizationsListOrganizationsResponseStatusMember  UserOrganizationUserSOrganizationsListOrganizationsResponseStatus = "member"
	UserOrganizationUserSOrganizationsListOrganizationsResponseStatusInvited UserOrganizationUserSOrganizationsListOrganizationsResponseStatus = "invited"
)

type UserOrganizationGetResponseEnvelope struct {
	Errors   []UserOrganizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserOrganizationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   UserOrganizationGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserOrganizationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userOrganizationGetResponseEnvelopeJSON    `json:"-"`
}

// userOrganizationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserOrganizationGetResponseEnvelope]
type userOrganizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserOrganizationGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    userOrganizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userOrganizationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserOrganizationGetResponseEnvelopeErrors]
type userOrganizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserOrganizationGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    userOrganizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userOrganizationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserOrganizationGetResponseEnvelopeMessages]
type userOrganizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserOrganizationGetResponseEnvelopeSuccess bool

const (
	UserOrganizationGetResponseEnvelopeSuccessTrue UserOrganizationGetResponseEnvelopeSuccess = true
)

type UserOrganizationUserSOrganizationsListOrganizationsParams struct {
	// Direction to order organizations.
	Direction param.Field[UserOrganizationUserSOrganizationsListOrganizationsParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[UserOrganizationUserSOrganizationsListOrganizationsParamsMatch] `query:"match"`
	// Organization name.
	Name param.Field[string] `query:"name"`
	// Field to order organizations by.
	Order param.Field[UserOrganizationUserSOrganizationsListOrganizationsParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of organizations per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status param.Field[UserOrganizationUserSOrganizationsListOrganizationsParamsStatus] `query:"status"`
}

// URLQuery serializes
// [UserOrganizationUserSOrganizationsListOrganizationsParams]'s query parameters
// as `url.Values`.
func (r UserOrganizationUserSOrganizationsListOrganizationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order organizations.
type UserOrganizationUserSOrganizationsListOrganizationsParamsDirection string

const (
	UserOrganizationUserSOrganizationsListOrganizationsParamsDirectionAsc  UserOrganizationUserSOrganizationsListOrganizationsParamsDirection = "asc"
	UserOrganizationUserSOrganizationsListOrganizationsParamsDirectionDesc UserOrganizationUserSOrganizationsListOrganizationsParamsDirection = "desc"
)

// Whether to match all search requirements or at least one (any).
type UserOrganizationUserSOrganizationsListOrganizationsParamsMatch string

const (
	UserOrganizationUserSOrganizationsListOrganizationsParamsMatchAny UserOrganizationUserSOrganizationsListOrganizationsParamsMatch = "any"
	UserOrganizationUserSOrganizationsListOrganizationsParamsMatchAll UserOrganizationUserSOrganizationsListOrganizationsParamsMatch = "all"
)

// Field to order organizations by.
type UserOrganizationUserSOrganizationsListOrganizationsParamsOrder string

const (
	UserOrganizationUserSOrganizationsListOrganizationsParamsOrderID     UserOrganizationUserSOrganizationsListOrganizationsParamsOrder = "id"
	UserOrganizationUserSOrganizationsListOrganizationsParamsOrderName   UserOrganizationUserSOrganizationsListOrganizationsParamsOrder = "name"
	UserOrganizationUserSOrganizationsListOrganizationsParamsOrderStatus UserOrganizationUserSOrganizationsListOrganizationsParamsOrder = "status"
)

// Whether the user is a member of the organization or has an inivitation pending.
type UserOrganizationUserSOrganizationsListOrganizationsParamsStatus string

const (
	UserOrganizationUserSOrganizationsListOrganizationsParamsStatusMember  UserOrganizationUserSOrganizationsListOrganizationsParamsStatus = "member"
	UserOrganizationUserSOrganizationsListOrganizationsParamsStatusInvited UserOrganizationUserSOrganizationsListOrganizationsParamsStatus = "invited"
)

type UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelope struct {
	Errors   []UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserOrganizationUserSOrganizationsListOrganizationsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeJSON       `json:"-"`
}

// userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelope]
type userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeErrorsJSON `json:"-"`
}

// userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeErrors]
type userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeMessagesJSON `json:"-"`
}

// userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeMessages]
type userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeSuccess bool

const (
	UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeSuccessTrue UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeSuccess = true
)

type UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeResultInfoJSON `json:"-"`
}

// userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeResultInfo]
type userOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationUserSOrganizationsListOrganizationsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
