// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// CasbApplicationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbApplicationService] method instead.
type CasbApplicationService struct {
	Options    []option.RequestOption
	SetupFlows *CasbApplicationSetupFlowService
}

// NewCasbApplicationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCasbApplicationService(opts ...option.RequestOption) (r *CasbApplicationService) {
	r = &CasbApplicationService{}
	r.Options = opts
	r.SetupFlows = NewCasbApplicationSetupFlowService(opts...)
	return
}

// Returns a list of available applications with use cases and permissions.
func (r *CasbApplicationService) List(ctx context.Context, params CasbApplicationListParams, opts ...option.RequestOption) (res *[]CasbApplicationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/applications", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns full application details including auth methods, use cases, and
// permissions.
func (r *CasbApplicationService) Get(ctx context.Context, slug CasbApplicationGetParamsSlug, query CasbApplicationGetParams, opts ...option.RequestOption) (res *CasbApplicationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/applications/%v", query.AccountID, slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Application item in list response.
type CasbApplicationListResponse struct {
	// Available auth methods.
	AuthMethods []CasbApplicationListResponseAuthMethod `json:"auth_methods" api:"required"`
	// Vendor category (e.g. Productivity, AI).
	Category string `json:"category" api:"required"`
	// Brief description of the integration.
	Description string `json:"description" api:"required"`
	// Human-readable vendor name.
	DisplayName string `json:"display_name" api:"required"`
	// Whether DLP scanning is supported.
	DLPEnabled bool `json:"dlp_enabled" api:"required"`
	// Logo path.
	Logo string `json:"logo" api:"required,nullable"`
	// All permissions with severity.
	Permissions []CasbApplicationListResponsePermission `json:"permissions" api:"required"`
	// Vendor identifier (e.g. microsoft_internal, google_workspace).
	//
	// - `GITHUB` - GITHUB
	// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
	// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
	// - `SALESFORCE` - SALESFORCE
	// - `SLACK` - SLACK
	Slug CasbApplicationListResponseSlug `json:"slug" api:"required"`
	// Environments this vendor supports (standard, fedramp).
	SupportedEnvironments []string `json:"supported_environments" api:"required"`
	// Supported use cases.
	UseCases []CasbApplicationListResponseUseCase `json:"use_cases" api:"required"`
	JSON     casbApplicationListResponseJSON      `json:"-"`
}

// casbApplicationListResponseJSON contains the JSON metadata for the struct
// [CasbApplicationListResponse]
type casbApplicationListResponseJSON struct {
	AuthMethods           apijson.Field
	Category              apijson.Field
	Description           apijson.Field
	DisplayName           apijson.Field
	DLPEnabled            apijson.Field
	Logo                  apijson.Field
	Permissions           apijson.Field
	Slug                  apijson.Field
	SupportedEnvironments apijson.Field
	UseCases              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbApplicationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationListResponseJSON) RawJSON() string {
	return r.raw
}

// Auth method summary for list endpoint.
type CasbApplicationListResponseAuthMethod struct {
	// Human-readable auth method name.
	DisplayName string `json:"display_name" api:"required"`
	// Auth method identifier.
	Slug string                                    `json:"slug" api:"required"`
	JSON casbApplicationListResponseAuthMethodJSON `json:"-"`
}

// casbApplicationListResponseAuthMethodJSON contains the JSON metadata for the
// struct [CasbApplicationListResponseAuthMethod]
type casbApplicationListResponseAuthMethodJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationListResponseAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationListResponseAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Permission/scope with severity for display.
type CasbApplicationListResponsePermission struct {
	// Human-readable permission name.
	DisplayName string `json:"display_name" api:"required"`
	// Vendor-native scope identifier.
	Scope string `json:"scope" api:"required"`
	// Permission sensitivity level.
	//
	// - `low` - low
	// - `medium` - medium
	// - `high` - high
	// - `critical` - critical
	Severity CasbApplicationListResponsePermissionsSeverity `json:"severity" api:"required"`
	JSON     casbApplicationListResponsePermissionJSON      `json:"-"`
}

// casbApplicationListResponsePermissionJSON contains the JSON metadata for the
// struct [CasbApplicationListResponsePermission]
type casbApplicationListResponsePermissionJSON struct {
	DisplayName apijson.Field
	Scope       apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationListResponsePermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationListResponsePermissionJSON) RawJSON() string {
	return r.raw
}

// Permission sensitivity level.
//
// - `low` - low
// - `medium` - medium
// - `high` - high
// - `critical` - critical
type CasbApplicationListResponsePermissionsSeverity string

const (
	CasbApplicationListResponsePermissionsSeverityLow      CasbApplicationListResponsePermissionsSeverity = "low"
	CasbApplicationListResponsePermissionsSeverityMedium   CasbApplicationListResponsePermissionsSeverity = "medium"
	CasbApplicationListResponsePermissionsSeverityHigh     CasbApplicationListResponsePermissionsSeverity = "high"
	CasbApplicationListResponsePermissionsSeverityCritical CasbApplicationListResponsePermissionsSeverity = "critical"
)

func (r CasbApplicationListResponsePermissionsSeverity) IsKnown() bool {
	switch r {
	case CasbApplicationListResponsePermissionsSeverityLow, CasbApplicationListResponsePermissionsSeverityMedium, CasbApplicationListResponsePermissionsSeverityHigh, CasbApplicationListResponsePermissionsSeverityCritical:
		return true
	}
	return false
}

// Vendor identifier (e.g. microsoft_internal, google_workspace).
//
// - `GITHUB` - GITHUB
// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
// - `SALESFORCE` - SALESFORCE
// - `SLACK` - SLACK
type CasbApplicationListResponseSlug string

const (
	CasbApplicationListResponseSlugGitHub            CasbApplicationListResponseSlug = "GITHUB"
	CasbApplicationListResponseSlugGoogleWorkspace   CasbApplicationListResponseSlug = "GOOGLE_WORKSPACE"
	CasbApplicationListResponseSlugMicrosoftInternal CasbApplicationListResponseSlug = "MICROSOFT_INTERNAL"
	CasbApplicationListResponseSlugSalesforce        CasbApplicationListResponseSlug = "SALESFORCE"
	CasbApplicationListResponseSlugSlack             CasbApplicationListResponseSlug = "SLACK"
)

func (r CasbApplicationListResponseSlug) IsKnown() bool {
	switch r {
	case CasbApplicationListResponseSlugGitHub, CasbApplicationListResponseSlugGoogleWorkspace, CasbApplicationListResponseSlugMicrosoftInternal, CasbApplicationListResponseSlugSalesforce, CasbApplicationListResponseSlugSlack:
		return true
	}
	return false
}

// Lightweight use case for list endpoint.
type CasbApplicationListResponseUseCase struct {
	// Human-readable use case name.
	DisplayName string `json:"display_name" api:"required"`
	// Use case identifier (e.g. casb, ces).
	Slug string                                 `json:"slug" api:"required"`
	JSON casbApplicationListResponseUseCaseJSON `json:"-"`
}

// casbApplicationListResponseUseCaseJSON contains the JSON metadata for the struct
// [CasbApplicationListResponseUseCase]
type casbApplicationListResponseUseCaseJSON struct {
	DisplayName apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationListResponseUseCase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationListResponseUseCaseJSON) RawJSON() string {
	return r.raw
}

// Full application detail for onboarding UI.
type CasbApplicationGetResponse struct {
	// Available authentication methods.
	AuthMethods []CasbApplicationGetResponseAuthMethod `json:"auth_methods" api:"required"`
	// Vendor category.
	Category string `json:"category" api:"required"`
	// Brief description.
	Description string `json:"description" api:"required"`
	// Human-readable vendor name.
	DisplayName string `json:"display_name" api:"required"`
	// Whether DLP scanning is supported.
	DLPEnabled bool `json:"dlp_enabled" api:"required"`
	// Setup instructions for the user.
	Instructions string `json:"instructions" api:"required,nullable"`
	// Logo path.
	Logo string `json:"logo" api:"required,nullable"`
	// Vendor identifier.
	//
	// - `GITHUB` - GITHUB
	// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
	// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
	// - `SALESFORCE` - SALESFORCE
	// - `SLACK` - SLACK
	Slug CasbApplicationGetResponseSlug `json:"slug" api:"required"`
	// Use cases with full scope details.
	UseCases []CasbApplicationGetResponseUseCase `json:"use_cases" api:"required"`
	JSON     casbApplicationGetResponseJSON      `json:"-"`
}

// casbApplicationGetResponseJSON contains the JSON metadata for the struct
// [CasbApplicationGetResponse]
type casbApplicationGetResponseJSON struct {
	AuthMethods  apijson.Field
	Category     apijson.Field
	Description  apijson.Field
	DisplayName  apijson.Field
	DLPEnabled   apijson.Field
	Instructions apijson.Field
	Logo         apijson.Field
	Slug         apijson.Field
	UseCases     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CasbApplicationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Authentication method available for a vendor.
type CasbApplicationGetResponseAuthMethod struct {
	// Human-readable auth method name.
	DisplayName string `json:"display_name" api:"required"`
	// Whether this is the default auth method.
	IsDefault bool `json:"is_default" api:"required"`
	// Auth method identifier.
	Slug string `json:"slug" api:"required"`
	// Environments this auth method supports.
	SupportedEnvironments []string                                 `json:"supported_environments" api:"required"`
	JSON                  casbApplicationGetResponseAuthMethodJSON `json:"-"`
}

// casbApplicationGetResponseAuthMethodJSON contains the JSON metadata for the
// struct [CasbApplicationGetResponseAuthMethod]
type casbApplicationGetResponseAuthMethodJSON struct {
	DisplayName           apijson.Field
	IsDefault             apijson.Field
	Slug                  apijson.Field
	SupportedEnvironments apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbApplicationGetResponseAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationGetResponseAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Vendor identifier.
//
// - `GITHUB` - GITHUB
// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
// - `SALESFORCE` - SALESFORCE
// - `SLACK` - SLACK
type CasbApplicationGetResponseSlug string

const (
	CasbApplicationGetResponseSlugGitHub            CasbApplicationGetResponseSlug = "GITHUB"
	CasbApplicationGetResponseSlugGoogleWorkspace   CasbApplicationGetResponseSlug = "GOOGLE_WORKSPACE"
	CasbApplicationGetResponseSlugMicrosoftInternal CasbApplicationGetResponseSlug = "MICROSOFT_INTERNAL"
	CasbApplicationGetResponseSlugSalesforce        CasbApplicationGetResponseSlug = "SALESFORCE"
	CasbApplicationGetResponseSlugSlack             CasbApplicationGetResponseSlug = "SLACK"
)

func (r CasbApplicationGetResponseSlug) IsKnown() bool {
	switch r {
	case CasbApplicationGetResponseSlugGitHub, CasbApplicationGetResponseSlugGoogleWorkspace, CasbApplicationGetResponseSlugMicrosoftInternal, CasbApplicationGetResponseSlugSalesforce, CasbApplicationGetResponseSlugSlack:
		return true
	}
	return false
}

// Full use case with scopes and features for detail endpoint.
type CasbApplicationGetResponseUseCase struct {
	// Scopes always required for this use case.
	BaseScopes []CasbApplicationGetResponseUseCasesBaseScope `json:"base_scopes" api:"required"`
	// Use case description.
	Description string `json:"description" api:"required"`
	// Human-readable use case name.
	DisplayName string `json:"display_name" api:"required"`
	// Optional features with extra scopes.
	Features []CasbApplicationGetResponseUseCasesFeature `json:"features" api:"required"`
	// Use case identifier.
	Slug string                                `json:"slug" api:"required"`
	JSON casbApplicationGetResponseUseCaseJSON `json:"-"`
}

// casbApplicationGetResponseUseCaseJSON contains the JSON metadata for the struct
// [CasbApplicationGetResponseUseCase]
type casbApplicationGetResponseUseCaseJSON struct {
	BaseScopes  apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Features    apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationGetResponseUseCase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationGetResponseUseCaseJSON) RawJSON() string {
	return r.raw
}

// Permission/scope with severity for display.
type CasbApplicationGetResponseUseCasesBaseScope struct {
	// Human-readable permission name.
	DisplayName string `json:"display_name" api:"required"`
	// Vendor-native scope identifier.
	Scope string `json:"scope" api:"required"`
	// Permission sensitivity level.
	//
	// - `low` - low
	// - `medium` - medium
	// - `high` - high
	// - `critical` - critical
	Severity CasbApplicationGetResponseUseCasesBaseScopesSeverity `json:"severity" api:"required"`
	JSON     casbApplicationGetResponseUseCasesBaseScopeJSON      `json:"-"`
}

// casbApplicationGetResponseUseCasesBaseScopeJSON contains the JSON metadata for
// the struct [CasbApplicationGetResponseUseCasesBaseScope]
type casbApplicationGetResponseUseCasesBaseScopeJSON struct {
	DisplayName apijson.Field
	Scope       apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationGetResponseUseCasesBaseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationGetResponseUseCasesBaseScopeJSON) RawJSON() string {
	return r.raw
}

// Permission sensitivity level.
//
// - `low` - low
// - `medium` - medium
// - `high` - high
// - `critical` - critical
type CasbApplicationGetResponseUseCasesBaseScopesSeverity string

const (
	CasbApplicationGetResponseUseCasesBaseScopesSeverityLow      CasbApplicationGetResponseUseCasesBaseScopesSeverity = "low"
	CasbApplicationGetResponseUseCasesBaseScopesSeverityMedium   CasbApplicationGetResponseUseCasesBaseScopesSeverity = "medium"
	CasbApplicationGetResponseUseCasesBaseScopesSeverityHigh     CasbApplicationGetResponseUseCasesBaseScopesSeverity = "high"
	CasbApplicationGetResponseUseCasesBaseScopesSeverityCritical CasbApplicationGetResponseUseCasesBaseScopesSeverity = "critical"
)

func (r CasbApplicationGetResponseUseCasesBaseScopesSeverity) IsKnown() bool {
	switch r {
	case CasbApplicationGetResponseUseCasesBaseScopesSeverityLow, CasbApplicationGetResponseUseCasesBaseScopesSeverityMedium, CasbApplicationGetResponseUseCasesBaseScopesSeverityHigh, CasbApplicationGetResponseUseCasesBaseScopesSeverityCritical:
		return true
	}
	return false
}

// A feature with its additional scopes.
type CasbApplicationGetResponseUseCasesFeature struct {
	// Feature description.
	Description string `json:"description" api:"required"`
	// Human-readable feature name.
	DisplayName string `json:"display_name" api:"required"`
	// Additional scopes when feature is enabled.
	Scopes []CasbApplicationGetResponseUseCasesFeaturesScope `json:"scopes" api:"required"`
	// Feature identifier.
	Slug string                                        `json:"slug" api:"required"`
	JSON casbApplicationGetResponseUseCasesFeatureJSON `json:"-"`
}

// casbApplicationGetResponseUseCasesFeatureJSON contains the JSON metadata for the
// struct [CasbApplicationGetResponseUseCasesFeature]
type casbApplicationGetResponseUseCasesFeatureJSON struct {
	Description apijson.Field
	DisplayName apijson.Field
	Scopes      apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationGetResponseUseCasesFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationGetResponseUseCasesFeatureJSON) RawJSON() string {
	return r.raw
}

// Permission/scope with severity for display.
type CasbApplicationGetResponseUseCasesFeaturesScope struct {
	// Human-readable permission name.
	DisplayName string `json:"display_name" api:"required"`
	// Vendor-native scope identifier.
	Scope string `json:"scope" api:"required"`
	// Permission sensitivity level.
	//
	// - `low` - low
	// - `medium` - medium
	// - `high` - high
	// - `critical` - critical
	Severity CasbApplicationGetResponseUseCasesFeaturesScopesSeverity `json:"severity" api:"required"`
	JSON     casbApplicationGetResponseUseCasesFeaturesScopeJSON      `json:"-"`
}

// casbApplicationGetResponseUseCasesFeaturesScopeJSON contains the JSON metadata
// for the struct [CasbApplicationGetResponseUseCasesFeaturesScope]
type casbApplicationGetResponseUseCasesFeaturesScopeJSON struct {
	DisplayName apijson.Field
	Scope       apijson.Field
	Severity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationGetResponseUseCasesFeaturesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationGetResponseUseCasesFeaturesScopeJSON) RawJSON() string {
	return r.raw
}

// Permission sensitivity level.
//
// - `low` - low
// - `medium` - medium
// - `high` - high
// - `critical` - critical
type CasbApplicationGetResponseUseCasesFeaturesScopesSeverity string

const (
	CasbApplicationGetResponseUseCasesFeaturesScopesSeverityLow      CasbApplicationGetResponseUseCasesFeaturesScopesSeverity = "low"
	CasbApplicationGetResponseUseCasesFeaturesScopesSeverityMedium   CasbApplicationGetResponseUseCasesFeaturesScopesSeverity = "medium"
	CasbApplicationGetResponseUseCasesFeaturesScopesSeverityHigh     CasbApplicationGetResponseUseCasesFeaturesScopesSeverity = "high"
	CasbApplicationGetResponseUseCasesFeaturesScopesSeverityCritical CasbApplicationGetResponseUseCasesFeaturesScopesSeverity = "critical"
)

func (r CasbApplicationGetResponseUseCasesFeaturesScopesSeverity) IsKnown() bool {
	switch r {
	case CasbApplicationGetResponseUseCasesFeaturesScopesSeverityLow, CasbApplicationGetResponseUseCasesFeaturesScopesSeverityMedium, CasbApplicationGetResponseUseCasesFeaturesScopesSeverityHigh, CasbApplicationGetResponseUseCasesFeaturesScopesSeverityCritical:
		return true
	}
	return false
}

type CasbApplicationListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by supported environment (standard, fedramp).
	Environment param.Field[string] `query:"environment"`
}

// URLQuery serializes [CasbApplicationListParams]'s query parameters as
// `url.Values`.
func (r CasbApplicationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CasbApplicationGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CasbApplicationGetParamsSlug string

const (
	CasbApplicationGetParamsSlugGitHub            CasbApplicationGetParamsSlug = "GITHUB"
	CasbApplicationGetParamsSlugGoogleWorkspace   CasbApplicationGetParamsSlug = "GOOGLE_WORKSPACE"
	CasbApplicationGetParamsSlugMicrosoftInternal CasbApplicationGetParamsSlug = "MICROSOFT_INTERNAL"
	CasbApplicationGetParamsSlugSalesforce        CasbApplicationGetParamsSlug = "SALESFORCE"
	CasbApplicationGetParamsSlugSlack             CasbApplicationGetParamsSlug = "SLACK"
)

func (r CasbApplicationGetParamsSlug) IsKnown() bool {
	switch r {
	case CasbApplicationGetParamsSlugGitHub, CasbApplicationGetParamsSlugGoogleWorkspace, CasbApplicationGetParamsSlugMicrosoftInternal, CasbApplicationGetParamsSlugSalesforce, CasbApplicationGetParamsSlugSlack:
		return true
	}
	return false
}
