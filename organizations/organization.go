// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/accounts"
	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// User API tokens can list and create Organizations, but reading, updating, or
// deleting an Organization requires a Global API key and your registered email
// address. The `cloudflare_organization` Terraform resource therefore requires a
// Global API key for its full lifecycle.
//
// See
// [Organizations authentication limitations](https://developers.cloudflare.com/fundamentals/organizations/limitations/#api-authentication)
// for credential compatibility by operation. Each endpoint’s **Security** section
// lists its supported authentication methods, with **Accepted Permissions** where
// applicable.
//
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
	Logs                 *LogService
	Billing              *BillingService
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
	r.Logs = NewLogService(opts...)
	r.Billing = NewBillingService(opts...)
	return
}

// Create a new organization for a user. Sub-organization creation availability
// depends on the organization's capabilities. The optional `profile` object
// accepts partial profile details; its individual fields are not required.
// Organization names are limited to 65 UTF-8 bytes. Each profile field is limited
// to 1,000 UTF-8 bytes, and a non-empty business email must be valid. At least one
// of business name, address, phone, or email must be non-empty to persist a
// profile.
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update an organization's name. Supplying `parent_id` verifies the current
// parent; this endpoint cannot move an organization. Organization names are
// limited to 65 UTF-8 bytes.
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a list of organizations a particular user has access to. At most one of
// `containing.account`, `containing.user`, or `containing.organization` may be
// provided.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Organization], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "organizations"
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

// Retrieve a list of organizations a particular user has access to. At most one of
// `containing.account`, `containing.user`, or `containing.organization` may be
// provided.
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Organization] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete an organization. It must not contain active sub-organizations or
// accounts. A root organization must also have no active managed users.
// Sub-organization deletion availability depends on the organization's
// capabilities.
//
// **Access Control:** Root deletion is restricted to enterprise organizations.
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationDeleteResponse, err error) {
	var env OrganizationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the details of a certain organization.
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// References an Organization in the Cloudflare data model.
type Organization struct {
	ID         string                  `json:"id" api:"required"`
	CreateTime time.Time               `json:"create_time" api:"required" format:"date-time"`
	Meta       OrganizationMeta        `json:"meta" api:"required"`
	Name       string                  `json:"name" api:"required"`
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
	AreChildAccountsAccessibleByOrgMembers bool `json:"are_child_accounts_accessible_by_org_members"`
	// Ordered chain of organization tags from the root organization down to (and
	// including) this organization itself. Root organizations return a single-element
	// array containing their own tag; sub-organizations return
	// `[rootTag, ...intermediateTags, parentTag, selfTag]`. Useful for constructing
	// authorization scopes that need to cover every ancestor in the hierarchy.
	HierarchyTags []string `json:"hierarchy_tags"`
	ManagedBy     string   `json:"managed_by"`
	// Enable features for Organizations.
	TenantFlags OrganizationMetaTenantFlags `json:"tenant_flags"`
	TenantType  string                      `json:"tenant_type"`
	ExtraFields map[string]interface{}      `json:"-" api:"extrafields"`
	JSON        organizationMetaJSON        `json:"-"`
}

// organizationMetaJSON contains the JSON metadata for the struct
// [OrganizationMeta]
type organizationMetaJSON struct {
	AreChildAccountsAccessibleByOrgMembers apijson.Field
	HierarchyTags                          apijson.Field
	ManagedBy                              apijson.Field
	TenantFlags                            apijson.Field
	TenantType                             apijson.Field
	raw                                    string
	ExtraFields                            map[string]apijson.Field
}

func (r *OrganizationMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMetaJSON) RawJSON() string {
	return r.raw
}

// Enable features for Organizations.
type OrganizationMetaTenantFlags struct {
	AccountCreation                      string                          `json:"account_creation" api:"required"`
	AccountCreationAppliesTenantDefaults string                          `json:"account_creation_applies_tenant_defaults" api:"required"`
	AccountDeletion                      string                          `json:"account_deletion" api:"required"`
	AccountMigration                     string                          `json:"account_migration" api:"required"`
	AccountMobility                      string                          `json:"account_mobility" api:"required"`
	EnterpriseCapability                 string                          `json:"enterprise_capability" api:"required"`
	MemberManagement                     string                          `json:"member_management" api:"required"`
	SubOrgCreation                       string                          `json:"sub_org_creation" api:"required"`
	JSON                                 organizationMetaTenantFlagsJSON `json:"-"`
}

// organizationMetaTenantFlagsJSON contains the JSON metadata for the struct
// [OrganizationMetaTenantFlags]
type organizationMetaTenantFlagsJSON struct {
	AccountCreation                      apijson.Field
	AccountCreationAppliesTenantDefaults apijson.Field
	AccountDeletion                      apijson.Field
	AccountMigration                     apijson.Field
	AccountMobility                      apijson.Field
	EnterpriseCapability                 apijson.Field
	MemberManagement                     apijson.Field
	SubOrgCreation                       apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *OrganizationMetaTenantFlags) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMetaTenantFlagsJSON) RawJSON() string {
	return r.raw
}

type OrganizationParent struct {
	ID   string                 `json:"id" api:"required"`
	Name string                 `json:"name" api:"required"`
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

type OrganizationDeleteResponse struct {
	ID   string                         `json:"id" api:"required"`
	JSON organizationDeleteResponseJSON `json:"-"`
}

// organizationDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationDeleteResponse]
type organizationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewParams struct {
	// The Organization name, limited to 65 UTF-8 bytes.
	Name   param.Field[string]                      `json:"name" api:"required"`
	Parent param.Field[OrganizationNewParamsParent] `json:"parent"`
	// Optional profile fields, each limited to 1,000 UTF-8 bytes. An empty profile is
	// ignored.
	Profile param.Field[OrganizationNewParamsProfile] `json:"profile"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationNewParamsParent struct {
	ID param.Field[string] `json:"id" api:"required"`
}

func (r OrganizationNewParamsParent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional profile fields, each limited to 1,000 UTF-8 bytes. An empty profile is
// ignored.
type OrganizationNewParamsProfile struct {
	BusinessAddress  param.Field[string] `json:"business_address"`
	BusinessEmail    param.Field[string] `json:"business_email"`
	BusinessName     param.Field[string] `json:"business_name"`
	BusinessPhone    param.Field[string] `json:"business_phone"`
	ExternalMetadata param.Field[string] `json:"external_metadata"`
}

func (r OrganizationNewParamsProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationNewResponseEnvelope struct {
	Errors   []interface{}         `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// References an Organization in the Cloudflare data model.
	Result  Organization                           `json:"result" api:"required"`
	Success OrganizationNewResponseEnvelopeSuccess `json:"success" api:"required"`
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
	// The Organization name, limited to 65 UTF-8 bytes. An omitted, empty, or
	// unchanged value is a no-op.
	Name param.Field[string] `json:"name"`
	// The current parent Organization ID. A different value is rejected; this endpoint
	// cannot move an Organization.
	ParentID param.Field[string] `json:"parent_id"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateResponseEnvelope struct {
	Errors   []interface{}         `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// References an Organization in the Cloudflare data model.
	Result  Organization                              `json:"result" api:"required"`
	Success OrganizationUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
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

type OrganizationListParams struct {
	// Only return organizations with the specified IDs (ex. id=foo&id=bar). Send
	// multiple elements by repeating the query value.
	ID         param.Field[[]string]                         `query:"id"`
	Containing param.Field[OrganizationListParamsContaining] `query:"containing"`
	// Sort direction for the `order_by` field. Defaults to `asc` when `order_by` is
	// specified.
	Direction param.Field[OrganizationListParamsDirection] `query:"direction"`
	Member    param.Field[OrganizationListParamsMember]    `query:"member"`
	Name      param.Field[OrganizationListParamsName]      `query:"name"`
	// Field to order results by. Currently supported values: `name`. When not
	// specified, results are ordered by hierarchy.
	OrderBy param.Field[OrganizationListParamsOrderBy] `query:"order_by"`
	// The amount of items to return. Defaults to 10.
	PageSize param.Field[int64] `query:"page_size"`
	// An opaque token returned from the last list response that when provided will
	// retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent requests
	// with a page token.
	PageToken param.Field[string]                       `query:"page_token"`
	Parent    param.Field[OrganizationListParamsParent] `query:"parent"`
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OrganizationListParamsContaining struct {
	// Filter the list of organizations to the ones that contain this particular
	// account.
	Account param.Field[string] `query:"account"`
	// Filter the list of organizations to the ones that contain this particular
	// organization.
	Organization param.Field[string] `query:"organization"`
	// Filter the list of organizations to the ones that contain this particular user.
	//
	// IMPORTANT: Just because an organization "contains" a user is not a
	// representation of any authorization or privilege to manage any resources
	// therein. An organization "containing" a user simply means the user is managed by
	// that organization.
	User param.Field[string] `query:"user"`
}

// URLQuery serializes [OrganizationListParamsContaining]'s query parameters as
// `url.Values`.
func (r OrganizationListParamsContaining) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction for the `order_by` field. Defaults to `asc` when `order_by` is
// specified.
type OrganizationListParamsDirection string

const (
	OrganizationListParamsDirectionAsc  OrganizationListParamsDirection = "asc"
	OrganizationListParamsDirectionDesc OrganizationListParamsDirection = "desc"
)

func (r OrganizationListParamsDirection) IsKnown() bool {
	switch r {
	case OrganizationListParamsDirectionAsc, OrganizationListParamsDirectionDesc:
		return true
	}
	return false
}

type OrganizationListParamsMember struct {
	// Filter organizations to those where the given user, by email, is an active
	// member of that exact organization. This does not include hierarchy-derived
	// access.
	Email param.Field[string] `query:"email"`
}

// URLQuery serializes [OrganizationListParamsMember]'s query parameters as
// `url.Values`.
func (r OrganizationListParamsMember) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OrganizationListParamsName struct {
	// (case-insensitive) Filter the list of organizations to where the name contains a
	// particular string.
	Contains param.Field[string] `query:"contains"`
	// (case-insensitive) Filter the list of organizations to where the name ends with
	// a particular string.
	EndsWith param.Field[string] `query:"endsWith"`
	// (case-insensitive) Filter the list of organizations to where the name starts
	// with a particular string.
	StartsWith param.Field[string] `query:"startsWith"`
}

// URLQuery serializes [OrganizationListParamsName]'s query parameters as
// `url.Values`.
func (r OrganizationListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order results by. Currently supported values: `name`. When not
// specified, results are ordered by hierarchy.
type OrganizationListParamsOrderBy string

const (
	OrganizationListParamsOrderByName OrganizationListParamsOrderBy = "name"
)

func (r OrganizationListParamsOrderBy) IsKnown() bool {
	switch r {
	case OrganizationListParamsOrderByName:
		return true
	}
	return false
}

type OrganizationListParamsParent struct {
	// Filter the list of organizations to the ones that are a sub-organization of the
	// specified organization.
	//
	// "null" is a valid value to provide for this parameter. It means "where an
	// organization has no parent (i.e. it is a 'root' organization)."
	ID param.Field[OrganizationListParamsParentID] `query:"id"`
}

// URLQuery serializes [OrganizationListParamsParent]'s query parameters as
// `url.Values`.
func (r OrganizationListParamsParent) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter the list of organizations to the ones that are a sub-organization of the
// specified organization.
//
// "null" is a valid value to provide for this parameter. It means "where an
// organization has no parent (i.e. it is a 'root' organization)."
type OrganizationListParamsParentID string

const (
	OrganizationListParamsParentIDNull OrganizationListParamsParentID = "null"
)

func (r OrganizationListParamsParentID) IsKnown() bool {
	switch r {
	case OrganizationListParamsParentIDNull:
		return true
	}
	return false
}

type OrganizationDeleteResponseEnvelope struct {
	Errors   []interface{}                             `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                     `json:"messages" api:"required"`
	Result   OrganizationDeleteResponse                `json:"result" api:"required"`
	Success  OrganizationDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON     organizationDeleteResponseEnvelopeJSON    `json:"-"`
}

// organizationDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationDeleteResponseEnvelope]
type organizationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteResponseEnvelopeSuccess bool

const (
	OrganizationDeleteResponseEnvelopeSuccessTrue OrganizationDeleteResponseEnvelopeSuccess = true
)

func (r OrganizationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationGetResponseEnvelope struct {
	Errors   []interface{}         `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// References an Organization in the Cloudflare data model.
	Result  Organization                           `json:"result" api:"required"`
	Success OrganizationGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
