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

// Lists organizations the user is associated with.
func (r *UserOrganizationService) List(ctx context.Context, query UserOrganizationListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[IamOrganization], err error) {
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

// Lists organizations the user is associated with.
func (r *UserOrganizationService) ListAutoPaging(ctx context.Context, query UserOrganizationListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[IamOrganization] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
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

type IamOrganization struct {
	// Identifier
	ID string `json:"id"`
	// Organization name.
	Name string `json:"name"`
	// Access permissions for this User.
	Permissions []string `json:"permissions"`
	// List of roles that a user has within an organization.
	Roles []string `json:"roles"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status IamOrganizationStatus `json:"status"`
	JSON   iamOrganizationJSON   `json:"-"`
}

// iamOrganizationJSON contains the JSON metadata for the struct [IamOrganization]
type iamOrganizationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the user is a member of the organization or has an inivitation pending.
type IamOrganizationStatus string

const (
	IamOrganizationStatusMember  IamOrganizationStatus = "member"
	IamOrganizationStatusInvited IamOrganizationStatus = "invited"
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

type UserOrganizationListParams struct {
	// Direction to order organizations.
	Direction param.Field[UserOrganizationListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[UserOrganizationListParamsMatch] `query:"match"`
	// Organization name.
	Name param.Field[string] `query:"name"`
	// Field to order organizations by.
	Order param.Field[UserOrganizationListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of organizations per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status param.Field[UserOrganizationListParamsStatus] `query:"status"`
}

// URLQuery serializes [UserOrganizationListParams]'s query parameters as
// `url.Values`.
func (r UserOrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order organizations.
type UserOrganizationListParamsDirection string

const (
	UserOrganizationListParamsDirectionAsc  UserOrganizationListParamsDirection = "asc"
	UserOrganizationListParamsDirectionDesc UserOrganizationListParamsDirection = "desc"
)

// Whether to match all search requirements or at least one (any).
type UserOrganizationListParamsMatch string

const (
	UserOrganizationListParamsMatchAny UserOrganizationListParamsMatch = "any"
	UserOrganizationListParamsMatchAll UserOrganizationListParamsMatch = "all"
)

// Field to order organizations by.
type UserOrganizationListParamsOrder string

const (
	UserOrganizationListParamsOrderID     UserOrganizationListParamsOrder = "id"
	UserOrganizationListParamsOrderName   UserOrganizationListParamsOrder = "name"
	UserOrganizationListParamsOrderStatus UserOrganizationListParamsOrder = "status"
)

// Whether the user is a member of the organization or has an inivitation pending.
type UserOrganizationListParamsStatus string

const (
	UserOrganizationListParamsStatusMember  UserOrganizationListParamsStatus = "member"
	UserOrganizationListParamsStatusInvited UserOrganizationListParamsStatus = "invited"
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
