// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// OrganizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options              []option.RequestOption
	OrganizationAccounts *OrganizationAccountService
	OrganizationProfile  *OrganizationProfileService
	Members              *MemberService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.OrganizationAccounts = NewOrganizationAccountService(opts...)
	r.OrganizationProfile = NewOrganizationProfileService(opts...)
	r.Members = NewMemberService(opts...)
	return
}

// Create a new organization for a user.
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify organization
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an organization. The organization MUST be empty before deleting. It must
// not contain any sub-organizations, accounts, members or users.
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve the details of a certain organization.
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Represents an Organization in the Cloudflare data model
type Organization struct {
	ID         string                  `json:"id,required"`
	CreateTime time.Time               `json:"create_time,required" format:"date-time"`
	Meta       OrganizationMeta        `json:"meta,required"`
	Name       string                  `json:"name,required"`
	Parent     OrganizationParent      `json:"parent"`
	Profile    accounts.AccountProfile `json:"profile"`
	JSON       organizationJSON        `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	ID          apijson.Field
	CreateTime  apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	Parent      apijson.Field
	Profile     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationMeta struct {
	// Organization flags for feature enablement
	Flags       OrganizationMetaFlags  `json:"flags"`
	ManagedBy   string                 `json:"managed_by"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        organizationMetaJSON   `json:"-"`
}

// organizationMetaJSON contains the JSON metadata for the struct
// [OrganizationMeta]
type organizationMetaJSON struct {
	Flags       apijson.Field
	ManagedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMetaJSON) RawJSON() string {
	return r.raw
}

// Organization flags for feature enablement
type OrganizationMetaFlags struct {
	AccountCreation  string                    `json:"account_creation,required"`
	AccountDeletion  string                    `json:"account_deletion,required"`
	AccountMigration string                    `json:"account_migration,required"`
	AccountMobility  string                    `json:"account_mobility,required"`
	SubOrgCreation   string                    `json:"sub_org_creation,required"`
	JSON             organizationMetaFlagsJSON `json:"-"`
}

// organizationMetaFlagsJSON contains the JSON metadata for the struct
// [OrganizationMetaFlags]
type organizationMetaFlagsJSON struct {
	AccountCreation  apijson.Field
	AccountDeletion  apijson.Field
	AccountMigration apijson.Field
	AccountMobility  apijson.Field
	SubOrgCreation   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationMetaFlags) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMetaFlagsJSON) RawJSON() string {
	return r.raw
}

type OrganizationParent struct {
	ID   string                 `json:"id,required"`
	Name string                 `json:"name,required"`
	JSON organizationParentJSON `json:"-"`
}

// organizationParentJSON contains the JSON metadata for the struct
// [OrganizationParent]
type organizationParentJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationParentJSON) RawJSON() string {
	return r.raw
}

// Represents an Organization in the Cloudflare data model
type OrganizationParam struct {
	Name    param.Field[string]                       `json:"name,required"`
	Parent  param.Field[OrganizationParentParam]      `json:"parent"`
	Profile param.Field[accounts.AccountProfileParam] `json:"profile"`
}

func (r OrganizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationMetaParam struct {
	// Organization flags for feature enablement
	Flags       param.Field[OrganizationMetaFlagsParam] `json:"flags"`
	ManagedBy   param.Field[string]                     `json:"managed_by"`
	ExtraFields map[string]interface{}                  `json:"-,extras"`
}

func (r OrganizationMetaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Organization flags for feature enablement
type OrganizationMetaFlagsParam struct {
	AccountCreation  param.Field[string] `json:"account_creation,required"`
	AccountDeletion  param.Field[string] `json:"account_deletion,required"`
	AccountMigration param.Field[string] `json:"account_migration,required"`
	AccountMobility  param.Field[string] `json:"account_mobility,required"`
	SubOrgCreation   param.Field[string] `json:"sub_org_creation,required"`
}

func (r OrganizationMetaFlagsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationParentParam struct {
	ID param.Field[string] `json:"id,required"`
}

func (r OrganizationParentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationNewParams struct {
	// Represents an Organization in the Cloudflare data model
	Organization OrganizationParam `json:"organization,required"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Organization)
}

type OrganizationNewResponseEnvelope struct {
	Errors   []interface{}         `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Represents an Organization in the Cloudflare data model
	Result  Organization                           `json:"result,required"`
	Success OrganizationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    organizationNewResponseEnvelopeJSON    `json:"-"`
}

// organizationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationNewResponseEnvelope]
type organizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeSuccess bool

const (
	OrganizationNewResponseEnvelopeSuccessTrue OrganizationNewResponseEnvelopeSuccess = true
)

func (r OrganizationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationUpdateParams struct {
	// Represents an Organization in the Cloudflare data model
	Organization OrganizationParam `json:"organization,required"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Organization)
}

type OrganizationUpdateResponseEnvelope struct {
	Errors   []interface{}         `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Represents an Organization in the Cloudflare data model
	Result  Organization                              `json:"result,required"`
	Success OrganizationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    organizationUpdateResponseEnvelopeJSON    `json:"-"`
}

// organizationUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponseEnvelope]
type organizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeSuccess bool

const (
	OrganizationUpdateResponseEnvelopeSuccessTrue OrganizationUpdateResponseEnvelopeSuccess = true
)

func (r OrganizationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationGetResponseEnvelope struct {
	Errors   []interface{}         `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Represents an Organization in the Cloudflare data model
	Result  Organization                           `json:"result,required"`
	Success OrganizationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    organizationGetResponseEnvelopeJSON    `json:"-"`
}

// organizationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationGetResponseEnvelope]
type organizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseEnvelopeSuccess bool

const (
	OrganizationGetResponseEnvelopeSuccessTrue OrganizationGetResponseEnvelopeSuccess = true
)

func (r OrganizationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
