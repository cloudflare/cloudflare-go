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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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

// Gets a specific organization the user is associated with.
func (r *UserOrganizationService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserOrganizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/organizations/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes association to an organization.
func (r *UserOrganizationService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserOrganizationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/organizations/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists organizations the user is associated with.
func (r *UserOrganizationService) UserSOrganizationsListOrganizations(ctx context.Context, query UserOrganizationUserSOrganizationsListOrganizationsParams, opts ...option.RequestOption) (res *shared.Page[UserOrganizationUserSOrganizationsListOrganizationsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/organizations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

type UserOrganizationGetResponse struct {
	Errors   []UserOrganizationGetResponseError   `json:"errors"`
	Messages []UserOrganizationGetResponseMessage `json:"messages"`
	Result   interface{}                          `json:"result"`
	// Whether the API call was successful
	Success UserOrganizationGetResponseSuccess `json:"success"`
	JSON    userOrganizationGetResponseJSON    `json:"-"`
}

// userOrganizationGetResponseJSON contains the JSON metadata for the struct
// [UserOrganizationGetResponse]
type userOrganizationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserOrganizationGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    userOrganizationGetResponseErrorJSON `json:"-"`
}

// userOrganizationGetResponseErrorJSON contains the JSON metadata for the struct
// [UserOrganizationGetResponseError]
type userOrganizationGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserOrganizationGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    userOrganizationGetResponseMessageJSON `json:"-"`
}

// userOrganizationGetResponseMessageJSON contains the JSON metadata for the struct
// [UserOrganizationGetResponseMessage]
type userOrganizationGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserOrganizationGetResponseSuccess bool

const (
	UserOrganizationGetResponseSuccessTrue UserOrganizationGetResponseSuccess = true
)

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
