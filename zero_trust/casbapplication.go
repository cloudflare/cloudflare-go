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

// Application item in list response.
type CasbApplicationListResponse struct {
	// Vendor identifier (e.g. microsoft_internal, google_workspace).
	//
	// - `BITBUCKET` - BITBUCKET
	// - `BOX` - BOX
	// - `CONFLUENCE` - CONFLUENCE
	// - `DROPBOX` - DROPBOX
	// - `GITHUB` - GITHUB
	// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
	// - `JIRA` - JIRA
	// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
	// - `SALESFORCE` - SALESFORCE
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
// - `BITBUCKET` - BITBUCKET
// - `BOX` - BOX
// - `CONFLUENCE` - CONFLUENCE
// - `DROPBOX` - DROPBOX
// - `GITHUB` - GITHUB
// - `GOOGLE_WORKSPACE` - GOOGLE_WORKSPACE
// - `JIRA` - JIRA
// - `MICROSOFT_INTERNAL` - MICROSOFT_INTERNAL
// - `SALESFORCE` - SALESFORCE
// - `SLACK` - SLACK
type CasbApplicationListResponseID string

const (
	CasbApplicationListResponseIDBitbucket         CasbApplicationListResponseID = "BITBUCKET"
	CasbApplicationListResponseIDBox               CasbApplicationListResponseID = "BOX"
	CasbApplicationListResponseIDConfluence        CasbApplicationListResponseID = "CONFLUENCE"
	CasbApplicationListResponseIDDropbox           CasbApplicationListResponseID = "DROPBOX"
	CasbApplicationListResponseIDGitHub            CasbApplicationListResponseID = "GITHUB"
	CasbApplicationListResponseIDGoogleWorkspace   CasbApplicationListResponseID = "GOOGLE_WORKSPACE"
	CasbApplicationListResponseIDJira              CasbApplicationListResponseID = "JIRA"
	CasbApplicationListResponseIDMicrosoftInternal CasbApplicationListResponseID = "MICROSOFT_INTERNAL"
	CasbApplicationListResponseIDSalesforce        CasbApplicationListResponseID = "SALESFORCE"
	CasbApplicationListResponseIDSlack             CasbApplicationListResponseID = "SLACK"
)

func (r CasbApplicationListResponseID) IsKnown() bool {
	switch r {
	case CasbApplicationListResponseIDBitbucket, CasbApplicationListResponseIDBox, CasbApplicationListResponseIDConfluence, CasbApplicationListResponseIDDropbox, CasbApplicationListResponseIDGitHub, CasbApplicationListResponseIDGoogleWorkspace, CasbApplicationListResponseIDJira, CasbApplicationListResponseIDMicrosoftInternal, CasbApplicationListResponseIDSalesforce, CasbApplicationListResponseIDSlack:
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
