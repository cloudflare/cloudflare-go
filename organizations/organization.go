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

	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
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

// Create a new organization for a user. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
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

// Modify organization. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
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

// Retrieve a list of organizations a particular user has access to. (Currently in
// Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
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

// Retrieve a list of organizations a particular user has access to. (Currently in
// Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Organization] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete an organization. The organization MUST be empty before deleting. It must
// not contain any sub-organizations, accounts, members or users. (Currently in
// Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationDeleteResponse, err error) {
	var env OrganizationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve the details of a certain organization. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
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

// References an Organization in the Cloudflare data model.
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
	// Enable features for Organizations.
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

// Enable features for Organizations.
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

// References an Organization in the Cloudflare data model.
type OrganizationParam struct {
	Name    param.Field[string]                       `json:"name,required"`
	Parent  param.Field[OrganizationParentParam]      `json:"parent"`
	Profile param.Field[accounts.AccountProfileParam] `json:"profile"`
}

func (r OrganizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationMetaParam struct {
	// Enable features for Organizations.
	Flags       param.Field[OrganizationMetaFlagsParam] `json:"flags"`
	ManagedBy   param.Field[string]                     `json:"managed_by"`
	ExtraFields map[string]interface{}                  `json:"-,extras"`
}

func (r OrganizationMetaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable features for Organizations.
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

type OrganizationDeleteResponse struct {
	ID   string                         `json:"id,required"`
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
	// References an Organization in the Cloudflare data model.
	Organization OrganizationParam `json:"organization,required"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Organization)
}

type OrganizationNewResponseEnvelope struct {
	Errors   []interface{}         `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// References an Organization in the Cloudflare data model.
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
	// References an Organization in the Cloudflare data model.
	Organization OrganizationParam `json:"organization,required"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Organization)
}

type OrganizationUpdateResponseEnvelope struct {
	Errors   []interface{}         `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// References an Organization in the Cloudflare data model.
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

type OrganizationListParams struct {
	// Only return organizations with the specified IDs (ex. id=foo&id=bar). Send
	// multiple elements by repeating the query value.
	ID         param.Field[[]string]                         `query:"id"`
	Containing param.Field[OrganizationListParamsContaining] `query:"containing"`
	Name       param.Field[OrganizationListParamsName]       `query:"name"`
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
	Errors   []interface{}                             `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   OrganizationDeleteResponse                `json:"result,required"`
	Success  OrganizationDeleteResponseEnvelopeSuccess `json:"success,required"`
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
	Errors   []interface{}         `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// References an Organization in the Cloudflare data model.
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
