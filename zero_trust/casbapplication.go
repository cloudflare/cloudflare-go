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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CasbApplicationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbApplicationService] method instead.
type CasbApplicationService struct {
	Options     []option.RequestOption
	AuthMethods *CasbApplicationAuthMethodService
}

// NewCasbApplicationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCasbApplicationService(opts ...option.RequestOption) (r *CasbApplicationService) {
	r = &CasbApplicationService{}
	r.Options = opts
	r.AuthMethods = NewCasbApplicationAuthMethodService(opts...)
	return
}

// Returns a list of available applications with use cases and permissions.
func (r *CasbApplicationService) List(ctx context.Context, params CasbApplicationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CasbApplicationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/applications", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns a list of available applications with use cases and permissions.
func (r *CasbApplicationService) ListAutoPaging(ctx context.Context, params CasbApplicationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CasbApplicationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Returns full application details including auth methods, use cases, and
// permissions.
func (r *CasbApplicationService) Get(ctx context.Context, applicationID CasbApplicationGetParamsApplicationID, query CasbApplicationGetParams, opts ...option.RequestOption) (res *CasbApplicationGetResponse, err error) {
	var env CasbApplicationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/one/applications/%v", query.AccountID, applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Application item in list response.
type CasbApplicationListResponse struct {
	// Vendor identifier (e.g. microsoft_internal, google_workspace).
	//
	// - `ANTHROPIC` - ANTHROPIC
	// - `AWS` - AWS
	// - `BITBUCKET` - BITBUCKET
	// - `BOX` - BOX
	// - `CONFLUENCE` - CONFLUENCE
	// - `DROPBOX` - DROPBOX
	// - `GITHUB` - GITHUB
	// - `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM
	// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
	// - `JIRA` - JIRA
	// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
	// - `OPENAI` - OPENAI
	// - `SALESFORCE` - SALESFORCE
	// - `SERVICENOW` - SERVICENOW
	// - `SLACK` - SLACK
	ID CasbApplicationListResponseID `json:"id" api:"required"`
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
	// Environments this vendor supports (standard, fedramp).
	SupportedEnvironments []string `json:"supported_environments" api:"required"`
	// Supported use cases.
	UseCases []CasbApplicationListResponseUseCase `json:"use_cases" api:"required"`
	JSON     casbApplicationListResponseJSON      `json:"-"`
}

// casbApplicationListResponseJSON contains the JSON metadata for the struct
// [CasbApplicationListResponse]
type casbApplicationListResponseJSON struct {
	ID                    apijson.Field
	AuthMethods           apijson.Field
	Category              apijson.Field
	Description           apijson.Field
	DisplayName           apijson.Field
	DLPEnabled            apijson.Field
	Logo                  apijson.Field
	Permissions           apijson.Field
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

// Vendor identifier (e.g. microsoft_internal, google_workspace).
//
// - `ANTHROPIC` - ANTHROPIC
// - `AWS` - AWS
// - `BITBUCKET` - BITBUCKET
// - `BOX` - BOX
// - `CONFLUENCE` - CONFLUENCE
// - `DROPBOX` - DROPBOX
// - `GITHUB` - GITHUB
// - `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM
// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
// - `JIRA` - JIRA
// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
// - `OPENAI` - OPENAI
// - `SALESFORCE` - SALESFORCE
// - `SERVICENOW` - SERVICENOW
// - `SLACK` - SLACK
type CasbApplicationListResponseID string

const (
	CasbApplicationListResponseIDAnthropic           CasbApplicationListResponseID = "ANTHROPIC"
	CasbApplicationListResponseIDAws                 CasbApplicationListResponseID = "AWS"
	CasbApplicationListResponseIDBitbucket           CasbApplicationListResponseID = "BITBUCKET"
	CasbApplicationListResponseIDBox                 CasbApplicationListResponseID = "BOX"
	CasbApplicationListResponseIDConfluence          CasbApplicationListResponseID = "CONFLUENCE"
	CasbApplicationListResponseIDDropbox             CasbApplicationListResponseID = "DROPBOX"
	CasbApplicationListResponseIDGitHub              CasbApplicationListResponseID = "GITHUB"
	CasbApplicationListResponseIDGoogleCloudPlatform CasbApplicationListResponseID = "GOOGLE_CLOUD_PLATFORM"
	CasbApplicationListResponseIDGoogleWorkspace     CasbApplicationListResponseID = "GOOGLE_WORKSPACE"
	CasbApplicationListResponseIDJira                CasbApplicationListResponseID = "JIRA"
	CasbApplicationListResponseIDMicrosoftInternal   CasbApplicationListResponseID = "MICROSOFT_INTERNAL"
	CasbApplicationListResponseIDOpenAI              CasbApplicationListResponseID = "OPENAI"
	CasbApplicationListResponseIDSalesforce          CasbApplicationListResponseID = "SALESFORCE"
	CasbApplicationListResponseIDServicenow          CasbApplicationListResponseID = "SERVICENOW"
	CasbApplicationListResponseIDSlack               CasbApplicationListResponseID = "SLACK"
)

func (r CasbApplicationListResponseID) IsKnown() bool {
	switch r {
	case CasbApplicationListResponseIDAnthropic, CasbApplicationListResponseIDAws, CasbApplicationListResponseIDBitbucket, CasbApplicationListResponseIDBox, CasbApplicationListResponseIDConfluence, CasbApplicationListResponseIDDropbox, CasbApplicationListResponseIDGitHub, CasbApplicationListResponseIDGoogleCloudPlatform, CasbApplicationListResponseIDGoogleWorkspace, CasbApplicationListResponseIDJira, CasbApplicationListResponseIDMicrosoftInternal, CasbApplicationListResponseIDOpenAI, CasbApplicationListResponseIDSalesforce, CasbApplicationListResponseIDServicenow, CasbApplicationListResponseIDSlack:
		return true
	}
	return false
}

// Auth method summary for list endpoint.
type CasbApplicationListResponseAuthMethod struct {
	// Auth method identifier.
	ID string `json:"id" api:"required"`
	// Human-readable auth method name.
	DisplayName string                                    `json:"display_name" api:"required"`
	JSON        casbApplicationListResponseAuthMethodJSON `json:"-"`
}

// casbApplicationListResponseAuthMethodJSON contains the JSON metadata for the
// struct [CasbApplicationListResponseAuthMethod]
type casbApplicationListResponseAuthMethodJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
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

// Lightweight use case for list endpoint.
type CasbApplicationListResponseUseCase struct {
	// Use case identifier (e.g. casb, ces).
	ID string `json:"id" api:"required"`
	// Human-readable use case name.
	DisplayName string                                 `json:"display_name" api:"required"`
	JSON        casbApplicationListResponseUseCaseJSON `json:"-"`
}

// casbApplicationListResponseUseCaseJSON contains the JSON metadata for the struct
// [CasbApplicationListResponseUseCase]
type casbApplicationListResponseUseCaseJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationListResponseUseCase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationListResponseUseCaseJSON) RawJSON() string {
	return r.raw
}

// The requested item.
type CasbApplicationGetResponse struct {
	// Vendor identifier.
	//
	// - `ANTHROPIC` - ANTHROPIC
	// - `AWS` - AWS
	// - `BITBUCKET` - BITBUCKET
	// - `BOX` - BOX
	// - `CONFLUENCE` - CONFLUENCE
	// - `DROPBOX` - DROPBOX
	// - `GITHUB` - GITHUB
	// - `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM
	// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
	// - `JIRA` - JIRA
	// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
	// - `OPENAI` - OPENAI
	// - `SALESFORCE` - SALESFORCE
	// - `SERVICENOW` - SERVICENOW
	// - `SLACK` - SLACK
	ID CasbApplicationGetResponseID `json:"id" api:"required"`
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
	// Use cases with full scope details.
	UseCases []CasbApplicationGetResponseUseCase `json:"use_cases" api:"required"`
	JSON     casbApplicationGetResponseJSON      `json:"-"`
}

// casbApplicationGetResponseJSON contains the JSON metadata for the struct
// [CasbApplicationGetResponse]
type casbApplicationGetResponseJSON struct {
	ID           apijson.Field
	AuthMethods  apijson.Field
	Category     apijson.Field
	Description  apijson.Field
	DisplayName  apijson.Field
	DLPEnabled   apijson.Field
	Instructions apijson.Field
	Logo         apijson.Field
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

// Vendor identifier.
//
// - `ANTHROPIC` - ANTHROPIC
// - `AWS` - AWS
// - `BITBUCKET` - BITBUCKET
// - `BOX` - BOX
// - `CONFLUENCE` - CONFLUENCE
// - `DROPBOX` - DROPBOX
// - `GITHUB` - GITHUB
// - `GOOGLE_CLOUD_PLATFORM` - GOOGLE_CLOUD_PLATFORM
// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
// - `JIRA` - JIRA
// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
// - `OPENAI` - OPENAI
// - `SALESFORCE` - SALESFORCE
// - `SERVICENOW` - SERVICENOW
// - `SLACK` - SLACK
type CasbApplicationGetResponseID string

const (
	CasbApplicationGetResponseIDAnthropic           CasbApplicationGetResponseID = "ANTHROPIC"
	CasbApplicationGetResponseIDAws                 CasbApplicationGetResponseID = "AWS"
	CasbApplicationGetResponseIDBitbucket           CasbApplicationGetResponseID = "BITBUCKET"
	CasbApplicationGetResponseIDBox                 CasbApplicationGetResponseID = "BOX"
	CasbApplicationGetResponseIDConfluence          CasbApplicationGetResponseID = "CONFLUENCE"
	CasbApplicationGetResponseIDDropbox             CasbApplicationGetResponseID = "DROPBOX"
	CasbApplicationGetResponseIDGitHub              CasbApplicationGetResponseID = "GITHUB"
	CasbApplicationGetResponseIDGoogleCloudPlatform CasbApplicationGetResponseID = "GOOGLE_CLOUD_PLATFORM"
	CasbApplicationGetResponseIDGoogleWorkspace     CasbApplicationGetResponseID = "GOOGLE_WORKSPACE"
	CasbApplicationGetResponseIDJira                CasbApplicationGetResponseID = "JIRA"
	CasbApplicationGetResponseIDMicrosoftInternal   CasbApplicationGetResponseID = "MICROSOFT_INTERNAL"
	CasbApplicationGetResponseIDOpenAI              CasbApplicationGetResponseID = "OPENAI"
	CasbApplicationGetResponseIDSalesforce          CasbApplicationGetResponseID = "SALESFORCE"
	CasbApplicationGetResponseIDServicenow          CasbApplicationGetResponseID = "SERVICENOW"
	CasbApplicationGetResponseIDSlack               CasbApplicationGetResponseID = "SLACK"
)

func (r CasbApplicationGetResponseID) IsKnown() bool {
	switch r {
	case CasbApplicationGetResponseIDAnthropic, CasbApplicationGetResponseIDAws, CasbApplicationGetResponseIDBitbucket, CasbApplicationGetResponseIDBox, CasbApplicationGetResponseIDConfluence, CasbApplicationGetResponseIDDropbox, CasbApplicationGetResponseIDGitHub, CasbApplicationGetResponseIDGoogleCloudPlatform, CasbApplicationGetResponseIDGoogleWorkspace, CasbApplicationGetResponseIDJira, CasbApplicationGetResponseIDMicrosoftInternal, CasbApplicationGetResponseIDOpenAI, CasbApplicationGetResponseIDSalesforce, CasbApplicationGetResponseIDServicenow, CasbApplicationGetResponseIDSlack:
		return true
	}
	return false
}

// Authentication method available for a vendor.
type CasbApplicationGetResponseAuthMethod struct {
	// Auth method identifier.
	ID string `json:"id" api:"required"`
	// Human-readable auth method name.
	DisplayName string `json:"display_name" api:"required"`
	// Whether this is the default auth method.
	IsDefault bool `json:"is_default" api:"required"`
	// Environments this auth method supports.
	SupportedEnvironments []string                                 `json:"supported_environments" api:"required"`
	JSON                  casbApplicationGetResponseAuthMethodJSON `json:"-"`
}

// casbApplicationGetResponseAuthMethodJSON contains the JSON metadata for the
// struct [CasbApplicationGetResponseAuthMethod]
type casbApplicationGetResponseAuthMethodJSON struct {
	ID                    apijson.Field
	DisplayName           apijson.Field
	IsDefault             apijson.Field
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

// Full use case with scopes and features for detail endpoint.
type CasbApplicationGetResponseUseCase struct {
	// Use case identifier.
	ID string `json:"id" api:"required"`
	// Scopes always required for this use case.
	BaseScopes []CasbApplicationGetResponseUseCasesBaseScope `json:"base_scopes" api:"required"`
	// Use case description.
	Description string `json:"description" api:"required"`
	// Human-readable use case name.
	DisplayName string `json:"display_name" api:"required"`
	// Optional features with extra scopes.
	Features []CasbApplicationGetResponseUseCasesFeature `json:"features" api:"required"`
	JSON     casbApplicationGetResponseUseCaseJSON       `json:"-"`
}

// casbApplicationGetResponseUseCaseJSON contains the JSON metadata for the struct
// [CasbApplicationGetResponseUseCase]
type casbApplicationGetResponseUseCaseJSON struct {
	ID          apijson.Field
	BaseScopes  apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Features    apijson.Field
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
	// Feature identifier.
	ID string `json:"id" api:"required"`
	// Feature description.
	Description string `json:"description" api:"required"`
	// Human-readable feature name.
	DisplayName string `json:"display_name" api:"required"`
	// Additional scopes when feature is enabled.
	Scopes []CasbApplicationGetResponseUseCasesFeaturesScope `json:"scopes" api:"required"`
	JSON   casbApplicationGetResponseUseCasesFeatureJSON     `json:"-"`
}

// casbApplicationGetResponseUseCasesFeatureJSON contains the JSON metadata for the
// struct [CasbApplicationGetResponseUseCasesFeature]
type casbApplicationGetResponseUseCasesFeatureJSON struct {
	ID          apijson.Field
	Description apijson.Field
	DisplayName apijson.Field
	Scopes      apijson.Field
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
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PageSize param.Field[int64] `query:"page_size"`
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

type CasbApplicationGetParamsApplicationID string

const (
	CasbApplicationGetParamsApplicationIDAnthropic           CasbApplicationGetParamsApplicationID = "ANTHROPIC"
	CasbApplicationGetParamsApplicationIDAws                 CasbApplicationGetParamsApplicationID = "AWS"
	CasbApplicationGetParamsApplicationIDBitbucket           CasbApplicationGetParamsApplicationID = "BITBUCKET"
	CasbApplicationGetParamsApplicationIDBox                 CasbApplicationGetParamsApplicationID = "BOX"
	CasbApplicationGetParamsApplicationIDConfluence          CasbApplicationGetParamsApplicationID = "CONFLUENCE"
	CasbApplicationGetParamsApplicationIDDropbox             CasbApplicationGetParamsApplicationID = "DROPBOX"
	CasbApplicationGetParamsApplicationIDGitHub              CasbApplicationGetParamsApplicationID = "GITHUB"
	CasbApplicationGetParamsApplicationIDGoogleCloudPlatform CasbApplicationGetParamsApplicationID = "GOOGLE_CLOUD_PLATFORM"
	CasbApplicationGetParamsApplicationIDGoogleWorkspace     CasbApplicationGetParamsApplicationID = "GOOGLE_WORKSPACE"
	CasbApplicationGetParamsApplicationIDJira                CasbApplicationGetParamsApplicationID = "JIRA"
	CasbApplicationGetParamsApplicationIDMicrosoftInternal   CasbApplicationGetParamsApplicationID = "MICROSOFT_INTERNAL"
	CasbApplicationGetParamsApplicationIDOpenAI              CasbApplicationGetParamsApplicationID = "OPENAI"
	CasbApplicationGetParamsApplicationIDSalesforce          CasbApplicationGetParamsApplicationID = "SALESFORCE"
	CasbApplicationGetParamsApplicationIDServicenow          CasbApplicationGetParamsApplicationID = "SERVICENOW"
	CasbApplicationGetParamsApplicationIDSlack               CasbApplicationGetParamsApplicationID = "SLACK"
)

func (r CasbApplicationGetParamsApplicationID) IsKnown() bool {
	switch r {
	case CasbApplicationGetParamsApplicationIDAnthropic, CasbApplicationGetParamsApplicationIDAws, CasbApplicationGetParamsApplicationIDBitbucket, CasbApplicationGetParamsApplicationIDBox, CasbApplicationGetParamsApplicationIDConfluence, CasbApplicationGetParamsApplicationIDDropbox, CasbApplicationGetParamsApplicationIDGitHub, CasbApplicationGetParamsApplicationIDGoogleCloudPlatform, CasbApplicationGetParamsApplicationIDGoogleWorkspace, CasbApplicationGetParamsApplicationIDJira, CasbApplicationGetParamsApplicationIDMicrosoftInternal, CasbApplicationGetParamsApplicationIDOpenAI, CasbApplicationGetParamsApplicationIDSalesforce, CasbApplicationGetParamsApplicationIDServicenow, CasbApplicationGetParamsApplicationIDSlack:
		return true
	}
	return false
}

type CasbApplicationGetResponseEnvelope struct {
	// The requested item.
	Result CasbApplicationGetResponse `json:"result" api:"required"`
	// Whether the request succeeded.
	Success bool `json:"success" api:"required"`
	// List of errors.
	Errors []map[string]interface{} `json:"errors"`
	// List of messages.
	Messages []string                               `json:"messages"`
	JSON     casbApplicationGetResponseEnvelopeJSON `json:"-"`
}

// casbApplicationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CasbApplicationGetResponseEnvelope]
type casbApplicationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbApplicationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbApplicationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
