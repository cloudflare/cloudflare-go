// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationService] method
// instead.
type OrganizationService struct {
	Options  []option.RequestOption
	Invites  *OrganizationInviteService
	Members  *OrganizationMemberService
	Railguns *OrganizationRailgunService
	Roles    *OrganizationRoleService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Invites = NewOrganizationInviteService(opts...)
	r.Members = NewOrganizationMemberService(opts...)
	r.Railguns = NewOrganizationRailgunService(opts...)
	r.Roles = NewOrganizationRoleService(opts...)
	return
}

// Get information about a specific organization that you are a member of.
func (r *OrganizationService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing Organization.
func (r *OrganizationService) Update(ctx context.Context, identifier string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *OrganizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type OrganizationGetResponse struct {
	Errors   []OrganizationGetResponseError   `json:"errors"`
	Messages []OrganizationGetResponseMessage `json:"messages"`
	Result   interface{}                      `json:"result"`
	// Whether the API call was successful
	Success OrganizationGetResponseSuccess `json:"success"`
	JSON    organizationGetResponseJSON    `json:"-"`
}

// organizationGetResponseJSON contains the JSON metadata for the struct
// [OrganizationGetResponse]
type organizationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationGetResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    organizationGetResponseErrorJSON `json:"-"`
}

// organizationGetResponseErrorJSON contains the JSON metadata for the struct
// [OrganizationGetResponseError]
type organizationGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationGetResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    organizationGetResponseMessageJSON `json:"-"`
}

// organizationGetResponseMessageJSON contains the JSON metadata for the struct
// [OrganizationGetResponseMessage]
type organizationGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationGetResponseSuccess bool

const (
	OrganizationGetResponseSuccessTrue OrganizationGetResponseSuccess = true
)

type OrganizationUpdateResponse struct {
	Errors   []OrganizationUpdateResponseError   `json:"errors"`
	Messages []OrganizationUpdateResponseMessage `json:"messages"`
	Result   interface{}                         `json:"result"`
	// Whether the API call was successful
	Success OrganizationUpdateResponseSuccess `json:"success"`
	JSON    organizationUpdateResponseJSON    `json:"-"`
}

// organizationUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponse]
type organizationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUpdateResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    organizationUpdateResponseErrorJSON `json:"-"`
}

// organizationUpdateResponseErrorJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponseError]
type organizationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUpdateResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    organizationUpdateResponseMessageJSON `json:"-"`
}

// organizationUpdateResponseMessageJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponseMessage]
type organizationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationUpdateResponseSuccess bool

const (
	OrganizationUpdateResponseSuccessTrue OrganizationUpdateResponseSuccess = true
)

type OrganizationUpdateParams struct {
	// Organization name.
	Name param.Field[string] `json:"name"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
