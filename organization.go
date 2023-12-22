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
func (r *OrganizationService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SingleOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing Organization.
func (r *OrganizationService) Update(ctx context.Context, identifier string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *SingleOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SingleOrganizationResponse struct {
	Errors   []SingleOrganizationResponseError   `json:"errors"`
	Messages []SingleOrganizationResponseMessage `json:"messages"`
	Result   interface{}                         `json:"result"`
	// Whether the API call was successful
	Success SingleOrganizationResponseSuccess `json:"success"`
	JSON    singleOrganizationResponseJSON    `json:"-"`
}

// singleOrganizationResponseJSON contains the JSON metadata for the struct
// [SingleOrganizationResponse]
type singleOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleOrganizationResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    singleOrganizationResponseErrorJSON `json:"-"`
}

// singleOrganizationResponseErrorJSON contains the JSON metadata for the struct
// [SingleOrganizationResponseError]
type singleOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleOrganizationResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    singleOrganizationResponseMessageJSON `json:"-"`
}

// singleOrganizationResponseMessageJSON contains the JSON metadata for the struct
// [SingleOrganizationResponseMessage]
type singleOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleOrganizationResponseSuccess bool

const (
	SingleOrganizationResponseSuccessTrue SingleOrganizationResponseSuccess = true
)

type OrganizationUpdateParams struct {
	// Organization name.
	Name param.Field[string] `json:"name"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
