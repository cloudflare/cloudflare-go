// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessAppPolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessAppPolicyService] method
// instead.
type AccessAppPolicyService struct {
	Options []option.RequestOption
}

// NewAccessAppPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessAppPolicyService(opts ...option.RequestOption) (r *AccessAppPolicyService) {
	r = &AccessAppPolicyService{}
	r.Options = opts
	return
}

// Fetches a single Access policy.
func (r *AccessAppPolicyService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, opts ...option.RequestOption) (res *AccessAppPolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppPolicyGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured Access policy.
func (r *AccessAppPolicyService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, body AccessAppPolicyUpdateParams, opts ...option.RequestOption) (res *AccessAppPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppPolicyUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an Access policy.
func (r *AccessAppPolicyService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, opts ...option.RequestOption) (res *AccessAppPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppPolicyDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new Access policy for an application.
func (r *AccessAppPolicyService) AccessPoliciesNewAnAccessPolicy(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessAppPolicyAccessPoliciesNewAnAccessPolicyParams, opts ...option.RequestOption) (res *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Access policies configured for an application.
func (r *AccessAppPolicyService) AccessPoliciesListAccessPolicies(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *[]AccessAppPolicyAccessPoliciesListAccessPoliciesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessAppPolicyGetResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyGetResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyGetResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyGetResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyGetResponseInclude `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessAppPolicyGetResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                         `json:"session_duration"`
	UpdatedAt       time.Time                      `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyGetResponseJSON `json:"-"`
}

// accessAppPolicyGetResponseJSON contains the JSON metadata for the struct
// [AccessAppPolicyGetResponse]
type accessAppPolicyGetResponseJSON struct {
	ID                           apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	Name                         apijson.Field
	Precedence                   apijson.Field
	PurposeJustificationPrompt   apijson.Field
	PurposeJustificationRequired apijson.Field
	Require                      apijson.Field
	SessionDuration              apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyGetResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                      `json:"email_list_uuid"`
	JSON          accessAppPolicyGetResponseApprovalGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseApprovalGroupJSON contains the JSON metadata for the
// struct [AccessAppPolicyGetResponseApprovalGroup]
type accessAppPolicyGetResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyGetResponseDecision string

const (
	AccessAppPolicyGetResponseDecisionAllow       AccessAppPolicyGetResponseDecision = "allow"
	AccessAppPolicyGetResponseDecisionDeny        AccessAppPolicyGetResponseDecision = "deny"
	AccessAppPolicyGetResponseDecisionNonIdentity AccessAppPolicyGetResponseDecision = "non_identity"
	AccessAppPolicyGetResponseDecisionBypass      AccessAppPolicyGetResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyGetResponseExcludeAccessEmailRule],
// [AccessAppPolicyGetResponseExcludeAccessEmailListRule],
// [AccessAppPolicyGetResponseExcludeAccessDomainRule],
// [AccessAppPolicyGetResponseExcludeAccessEveryoneRule],
// [AccessAppPolicyGetResponseExcludeAccessIPRule],
// [AccessAppPolicyGetResponseExcludeAccessIPListRule],
// [AccessAppPolicyGetResponseExcludeAccessCertificateRule],
// [AccessAppPolicyGetResponseExcludeAccessAccessGroupRule],
// [AccessAppPolicyGetResponseExcludeAccessAzureGroupRule],
// [AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyGetResponseExcludeAccessOktaGroupRule],
// [AccessAppPolicyGetResponseExcludeAccessSamlGroupRule],
// [AccessAppPolicyGetResponseExcludeAccessServiceTokenRule],
// [AccessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyGetResponseExcludeAccessCountryRule],
// [AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyGetResponseExcludeAccessDevicePostureRule].
type AccessAppPolicyGetResponseExclude interface {
	implementsAccessAppPolicyGetResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyGetResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyGetResponseExcludeAccessEmailRule struct {
	Email AccessAppPolicyGetResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyGetResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseExcludeAccessEmailRule]
type accessAppPolicyGetResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessEmailRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                    `json:"email,required" format:"email"`
	JSON  accessAppPolicyGetResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessEmailRuleEmail]
type accessAppPolicyGetResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyGetResponseExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyGetResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyGetResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessEmailListRule]
type accessAppPolicyGetResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessEmailListRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                            `json:"id,required"`
	JSON accessAppPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessEmailListRuleEmailList]
type accessAppPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyGetResponseExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyGetResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyGetResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseExcludeAccessDomainRule]
type accessAppPolicyGetResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessDomainRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                           `json:"domain,required"`
	JSON   accessAppPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyGetResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                             `json:"everyone,required"`
	JSON     accessAppPolicyGetResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessEveryoneRule]
type accessAppPolicyGetResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessEveryoneRule) implementsAccessAppPolicyGetResponseExclude() {
}

// Matches an IP address block.
type AccessAppPolicyGetResponseExcludeAccessIPRule struct {
	IP   AccessAppPolicyGetResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyGetResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessAppPolicyGetResponseExcludeAccessIPRule]
type accessAppPolicyGetResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessIPRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                              `json:"ip,required"`
	JSON accessAppPolicyGetResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseExcludeAccessIPRuleIP]
type accessAppPolicyGetResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyGetResponseExcludeAccessIPListRule struct {
	IPList AccessAppPolicyGetResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyGetResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseExcludeAccessIPListRule]
type accessAppPolicyGetResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessIPListRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                      `json:"id,required"`
	JSON accessAppPolicyGetResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessIPListRuleIPList]
type accessAppPolicyGetResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyGetResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                `json:"certificate,required"`
	JSON        accessAppPolicyGetResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessCertificateRule]
type accessAppPolicyGetResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessCertificateRule) implementsAccessAppPolicyGetResponseExclude() {
}

// Matches an Access group.
type AccessAppPolicyGetResponseExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyGetResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyGetResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessAccessGroupRule]
type accessAppPolicyGetResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessAccessGroupRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                          `json:"id,required"`
	JSON accessAppPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyGetResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyGetResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessAzureGroupRule]
type accessAppPolicyGetResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessAzureGroupRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                           `json:"connection_id,required"`
	JSON         accessAppPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRule]
type accessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                              `json:"name,required"`
	JSON accessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRule]
type accessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                           `json:"email,required"`
	JSON  accessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyGetResponseExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyGetResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyGetResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessOktaGroupRule]
type accessAppPolicyGetResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessOktaGroupRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                       `json:"email,required"`
	JSON  accessAppPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyGetResponseExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyGetResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyGetResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessSamlGroupRule]
type accessAppPolicyGetResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessSamlGroupRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                       `json:"attribute_value,required"`
	JSON           accessAppPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyGetResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyGetResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessServiceTokenRule]
type accessAppPolicyGetResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessServiceTokenRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                  `json:"token_id,required"`
	JSON    accessAppPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                         `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyGetResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRule]
type accessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                              `json:"keys_url,required"`
	JSON    accessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyGetResponseExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyGetResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyGetResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessCountryRule]
type accessAppPolicyGetResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessCountryRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                    `json:"country_code,required"`
	JSON        accessAppPolicyGetResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseExcludeAccessCountryRuleGeo]
type accessAppPolicyGetResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRule]
type accessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                        `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyGetResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyGetResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessDevicePostureRule]
type accessAppPolicyGetResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseExcludeAccessDevicePostureRule) implementsAccessAppPolicyGetResponseExclude() {
}

type AccessAppPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                    `json:"integration_uid,required"`
	JSON           accessAppPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyGetResponseIncludeAccessEmailRule],
// [AccessAppPolicyGetResponseIncludeAccessEmailListRule],
// [AccessAppPolicyGetResponseIncludeAccessDomainRule],
// [AccessAppPolicyGetResponseIncludeAccessEveryoneRule],
// [AccessAppPolicyGetResponseIncludeAccessIPRule],
// [AccessAppPolicyGetResponseIncludeAccessIPListRule],
// [AccessAppPolicyGetResponseIncludeAccessCertificateRule],
// [AccessAppPolicyGetResponseIncludeAccessAccessGroupRule],
// [AccessAppPolicyGetResponseIncludeAccessAzureGroupRule],
// [AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyGetResponseIncludeAccessOktaGroupRule],
// [AccessAppPolicyGetResponseIncludeAccessSamlGroupRule],
// [AccessAppPolicyGetResponseIncludeAccessServiceTokenRule],
// [AccessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyGetResponseIncludeAccessCountryRule],
// [AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyGetResponseIncludeAccessDevicePostureRule].
type AccessAppPolicyGetResponseInclude interface {
	implementsAccessAppPolicyGetResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyGetResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyGetResponseIncludeAccessEmailRule struct {
	Email AccessAppPolicyGetResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyGetResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseIncludeAccessEmailRule]
type accessAppPolicyGetResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessEmailRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                    `json:"email,required" format:"email"`
	JSON  accessAppPolicyGetResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessEmailRuleEmail]
type accessAppPolicyGetResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyGetResponseIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyGetResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyGetResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessEmailListRule]
type accessAppPolicyGetResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessEmailListRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                            `json:"id,required"`
	JSON accessAppPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessEmailListRuleEmailList]
type accessAppPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyGetResponseIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyGetResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyGetResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseIncludeAccessDomainRule]
type accessAppPolicyGetResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessDomainRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                           `json:"domain,required"`
	JSON   accessAppPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyGetResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                             `json:"everyone,required"`
	JSON     accessAppPolicyGetResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessEveryoneRule]
type accessAppPolicyGetResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessEveryoneRule) implementsAccessAppPolicyGetResponseInclude() {
}

// Matches an IP address block.
type AccessAppPolicyGetResponseIncludeAccessIPRule struct {
	IP   AccessAppPolicyGetResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyGetResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessAppPolicyGetResponseIncludeAccessIPRule]
type accessAppPolicyGetResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessIPRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                              `json:"ip,required"`
	JSON accessAppPolicyGetResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseIncludeAccessIPRuleIP]
type accessAppPolicyGetResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyGetResponseIncludeAccessIPListRule struct {
	IPList AccessAppPolicyGetResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyGetResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseIncludeAccessIPListRule]
type accessAppPolicyGetResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessIPListRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                      `json:"id,required"`
	JSON accessAppPolicyGetResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessIPListRuleIPList]
type accessAppPolicyGetResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyGetResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                `json:"certificate,required"`
	JSON        accessAppPolicyGetResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessCertificateRule]
type accessAppPolicyGetResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessCertificateRule) implementsAccessAppPolicyGetResponseInclude() {
}

// Matches an Access group.
type AccessAppPolicyGetResponseIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyGetResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyGetResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessAccessGroupRule]
type accessAppPolicyGetResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessAccessGroupRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                          `json:"id,required"`
	JSON accessAppPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyGetResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyGetResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessAzureGroupRule]
type accessAppPolicyGetResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessAzureGroupRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                           `json:"connection_id,required"`
	JSON         accessAppPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRule]
type accessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                              `json:"name,required"`
	JSON accessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRule]
type accessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                           `json:"email,required"`
	JSON  accessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyGetResponseIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyGetResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyGetResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessOktaGroupRule]
type accessAppPolicyGetResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessOktaGroupRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                       `json:"email,required"`
	JSON  accessAppPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyGetResponseIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyGetResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyGetResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessSamlGroupRule]
type accessAppPolicyGetResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessSamlGroupRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                       `json:"attribute_value,required"`
	JSON           accessAppPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyGetResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyGetResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessServiceTokenRule]
type accessAppPolicyGetResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessServiceTokenRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                  `json:"token_id,required"`
	JSON    accessAppPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                         `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyGetResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRule]
type accessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                              `json:"keys_url,required"`
	JSON    accessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyGetResponseIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyGetResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyGetResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessCountryRule]
type accessAppPolicyGetResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessCountryRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                    `json:"country_code,required"`
	JSON        accessAppPolicyGetResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseIncludeAccessCountryRuleGeo]
type accessAppPolicyGetResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRule]
type accessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                        `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyGetResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyGetResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessDevicePostureRule]
type accessAppPolicyGetResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseIncludeAccessDevicePostureRule) implementsAccessAppPolicyGetResponseInclude() {
}

type AccessAppPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                    `json:"integration_uid,required"`
	JSON           accessAppPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyGetResponseRequireAccessEmailRule],
// [AccessAppPolicyGetResponseRequireAccessEmailListRule],
// [AccessAppPolicyGetResponseRequireAccessDomainRule],
// [AccessAppPolicyGetResponseRequireAccessEveryoneRule],
// [AccessAppPolicyGetResponseRequireAccessIPRule],
// [AccessAppPolicyGetResponseRequireAccessIPListRule],
// [AccessAppPolicyGetResponseRequireAccessCertificateRule],
// [AccessAppPolicyGetResponseRequireAccessAccessGroupRule],
// [AccessAppPolicyGetResponseRequireAccessAzureGroupRule],
// [AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyGetResponseRequireAccessGsuiteGroupRule],
// [AccessAppPolicyGetResponseRequireAccessOktaGroupRule],
// [AccessAppPolicyGetResponseRequireAccessSamlGroupRule],
// [AccessAppPolicyGetResponseRequireAccessServiceTokenRule],
// [AccessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyGetResponseRequireAccessExternalEvaluationRule],
// [AccessAppPolicyGetResponseRequireAccessCountryRule],
// [AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRule] or
// [AccessAppPolicyGetResponseRequireAccessDevicePostureRule].
type AccessAppPolicyGetResponseRequire interface {
	implementsAccessAppPolicyGetResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyGetResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyGetResponseRequireAccessEmailRule struct {
	Email AccessAppPolicyGetResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyGetResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseRequireAccessEmailRule]
type accessAppPolicyGetResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessEmailRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                    `json:"email,required" format:"email"`
	JSON  accessAppPolicyGetResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessEmailRuleEmail]
type accessAppPolicyGetResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyGetResponseRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyGetResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyGetResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessEmailListRule]
type accessAppPolicyGetResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessEmailListRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                            `json:"id,required"`
	JSON accessAppPolicyGetResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessEmailListRuleEmailList]
type accessAppPolicyGetResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyGetResponseRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyGetResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyGetResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseRequireAccessDomainRule]
type accessAppPolicyGetResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessDomainRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                           `json:"domain,required"`
	JSON   accessAppPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessDomainRuleEmailDomain]
type accessAppPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyGetResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                             `json:"everyone,required"`
	JSON     accessAppPolicyGetResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessEveryoneRule]
type accessAppPolicyGetResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessEveryoneRule) implementsAccessAppPolicyGetResponseRequire() {
}

// Matches an IP address block.
type AccessAppPolicyGetResponseRequireAccessIPRule struct {
	IP   AccessAppPolicyGetResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyGetResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessAppPolicyGetResponseRequireAccessIPRule]
type accessAppPolicyGetResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessIPRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                              `json:"ip,required"`
	JSON accessAppPolicyGetResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseRequireAccessIPRuleIP]
type accessAppPolicyGetResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyGetResponseRequireAccessIPListRule struct {
	IPList AccessAppPolicyGetResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyGetResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyGetResponseRequireAccessIPListRule]
type accessAppPolicyGetResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessIPListRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                      `json:"id,required"`
	JSON accessAppPolicyGetResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessIPListRuleIPList]
type accessAppPolicyGetResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyGetResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                `json:"certificate,required"`
	JSON        accessAppPolicyGetResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessCertificateRule]
type accessAppPolicyGetResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessCertificateRule) implementsAccessAppPolicyGetResponseRequire() {
}

// Matches an Access group.
type AccessAppPolicyGetResponseRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyGetResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyGetResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessAccessGroupRule]
type accessAppPolicyGetResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessAccessGroupRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                          `json:"id,required"`
	JSON accessAppPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessAccessGroupRuleGroup]
type accessAppPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyGetResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyGetResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyGetResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessAzureGroupRule]
type accessAppPolicyGetResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessAzureGroupRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                           `json:"connection_id,required"`
	JSON         accessAppPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRule]
type accessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                              `json:"name,required"`
	JSON accessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyGetResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyGetResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessGsuiteGroupRule]
type accessAppPolicyGetResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessGsuiteGroupRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                           `json:"email,required"`
	JSON  accessAppPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyGetResponseRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyGetResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyGetResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessOktaGroupRule]
type accessAppPolicyGetResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessOktaGroupRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                       `json:"email,required"`
	JSON  accessAppPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessOktaGroupRuleOkta]
type accessAppPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyGetResponseRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyGetResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyGetResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessSamlGroupRule]
type accessAppPolicyGetResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessSamlGroupRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                       `json:"attribute_value,required"`
	JSON           accessAppPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessSamlGroupRuleSaml]
type accessAppPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyGetResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyGetResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyGetResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessServiceTokenRule]
type accessAppPolicyGetResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessServiceTokenRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                  `json:"token_id,required"`
	JSON    accessAppPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                         `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyGetResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyGetResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyGetResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessExternalEvaluationRule]
type accessAppPolicyGetResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessExternalEvaluationRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                              `json:"keys_url,required"`
	JSON    accessAppPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyGetResponseRequireAccessCountryRule struct {
	Geo  AccessAppPolicyGetResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyGetResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessCountryRule]
type accessAppPolicyGetResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessCountryRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                    `json:"country_code,required"`
	JSON        accessAppPolicyGetResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseRequireAccessCountryRuleGeo]
type accessAppPolicyGetResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRule]
type accessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                        `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyGetResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyGetResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessDevicePostureRule]
type accessAppPolicyGetResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseRequireAccessDevicePostureRule) implementsAccessAppPolicyGetResponseRequire() {
}

type AccessAppPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                    `json:"integration_uid,required"`
	JSON           accessAppPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyUpdateResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyUpdateResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyUpdateResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyUpdateResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyUpdateResponseInclude `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessAppPolicyUpdateResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                            `json:"session_duration"`
	UpdatedAt       time.Time                         `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyUpdateResponseJSON `json:"-"`
}

// accessAppPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponse]
type accessAppPolicyUpdateResponseJSON struct {
	ID                           apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	Name                         apijson.Field
	Precedence                   apijson.Field
	PurposeJustificationPrompt   apijson.Field
	PurposeJustificationRequired apijson.Field
	Require                      apijson.Field
	SessionDuration              apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyUpdateResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                         `json:"email_list_uuid"`
	JSON          accessAppPolicyUpdateResponseApprovalGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseApprovalGroupJSON contains the JSON metadata for
// the struct [AccessAppPolicyUpdateResponseApprovalGroup]
type accessAppPolicyUpdateResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyUpdateResponseDecision string

const (
	AccessAppPolicyUpdateResponseDecisionAllow       AccessAppPolicyUpdateResponseDecision = "allow"
	AccessAppPolicyUpdateResponseDecisionDeny        AccessAppPolicyUpdateResponseDecision = "deny"
	AccessAppPolicyUpdateResponseDecisionNonIdentity AccessAppPolicyUpdateResponseDecision = "non_identity"
	AccessAppPolicyUpdateResponseDecisionBypass      AccessAppPolicyUpdateResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyUpdateResponseExcludeAccessEmailRule],
// [AccessAppPolicyUpdateResponseExcludeAccessEmailListRule],
// [AccessAppPolicyUpdateResponseExcludeAccessDomainRule],
// [AccessAppPolicyUpdateResponseExcludeAccessEveryoneRule],
// [AccessAppPolicyUpdateResponseExcludeAccessIPRule],
// [AccessAppPolicyUpdateResponseExcludeAccessIPListRule],
// [AccessAppPolicyUpdateResponseExcludeAccessCertificateRule],
// [AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRule],
// [AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRule],
// [AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRule],
// [AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRule],
// [AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRule],
// [AccessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateResponseExcludeAccessCountryRule],
// [AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRule].
type AccessAppPolicyUpdateResponseExclude interface {
	implementsAccessAppPolicyUpdateResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyUpdateResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyUpdateResponseExcludeAccessEmailRule struct {
	Email AccessAppPolicyUpdateResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseExcludeAccessEmailRule]
type accessAppPolicyUpdateResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessEmailRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                       `json:"email,required" format:"email"`
	JSON  accessAppPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessEmailRuleEmail]
type accessAppPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateResponseExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyUpdateResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyUpdateResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessEmailListRule]
type accessAppPolicyUpdateResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessEmailListRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                               `json:"id,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessEmailListRuleEmailList]
type accessAppPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateResponseExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyUpdateResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseExcludeAccessDomainRule]
type accessAppPolicyUpdateResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessDomainRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                              `json:"domain,required"`
	JSON   accessAppPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyUpdateResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                `json:"everyone,required"`
	JSON     accessAppPolicyUpdateResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseExcludeAccessEveryoneRule]
type accessAppPolicyUpdateResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessEveryoneRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateResponseExcludeAccessIPRule struct {
	IP   AccessAppPolicyUpdateResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyUpdateResponseExcludeAccessIPRule]
type accessAppPolicyUpdateResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessIPRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                 `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseExcludeAccessIPRuleIP]
type accessAppPolicyUpdateResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateResponseExcludeAccessIPListRule struct {
	IPList AccessAppPolicyUpdateResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyUpdateResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseExcludeAccessIPListRule]
type accessAppPolicyUpdateResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessIPListRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                         `json:"id,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessIPListRuleIPList]
type accessAppPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                   `json:"certificate,required"`
	JSON        accessAppPolicyUpdateResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessCertificateRule]
type accessAppPolicyUpdateResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessCertificateRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

// Matches an Access group.
type AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRule]
type accessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                             `json:"id,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRule]
type accessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                              `json:"connection_id,required"`
	JSON         accessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRule]
type accessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                 `json:"name,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRule]
type accessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                              `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRule]
type accessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                          `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRule]
type accessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                          `json:"attribute_value,required"`
	JSON           accessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRule]
type accessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                     `json:"token_id,required"`
	JSON    accessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                            `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRule]
type accessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                 `json:"keys_url,required"`
	JSON    accessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyUpdateResponseExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyUpdateResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyUpdateResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseExcludeAccessCountryRule]
type accessAppPolicyUpdateResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessCountryRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                       `json:"country_code,required"`
	JSON        accessAppPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessCountryRuleGeo]
type accessAppPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRule]
type accessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                           `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRule]
type accessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRule) implementsAccessAppPolicyUpdateResponseExclude() {
}

type AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                       `json:"integration_uid,required"`
	JSON           accessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyUpdateResponseIncludeAccessEmailRule],
// [AccessAppPolicyUpdateResponseIncludeAccessEmailListRule],
// [AccessAppPolicyUpdateResponseIncludeAccessDomainRule],
// [AccessAppPolicyUpdateResponseIncludeAccessEveryoneRule],
// [AccessAppPolicyUpdateResponseIncludeAccessIPRule],
// [AccessAppPolicyUpdateResponseIncludeAccessIPListRule],
// [AccessAppPolicyUpdateResponseIncludeAccessCertificateRule],
// [AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRule],
// [AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRule],
// [AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRule],
// [AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRule],
// [AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRule],
// [AccessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateResponseIncludeAccessCountryRule],
// [AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRule].
type AccessAppPolicyUpdateResponseInclude interface {
	implementsAccessAppPolicyUpdateResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyUpdateResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyUpdateResponseIncludeAccessEmailRule struct {
	Email AccessAppPolicyUpdateResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseIncludeAccessEmailRule]
type accessAppPolicyUpdateResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessEmailRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                       `json:"email,required" format:"email"`
	JSON  accessAppPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessEmailRuleEmail]
type accessAppPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateResponseIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyUpdateResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyUpdateResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessEmailListRule]
type accessAppPolicyUpdateResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessEmailListRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                               `json:"id,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessEmailListRuleEmailList]
type accessAppPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateResponseIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyUpdateResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseIncludeAccessDomainRule]
type accessAppPolicyUpdateResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessDomainRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                              `json:"domain,required"`
	JSON   accessAppPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyUpdateResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                `json:"everyone,required"`
	JSON     accessAppPolicyUpdateResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseIncludeAccessEveryoneRule]
type accessAppPolicyUpdateResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessEveryoneRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateResponseIncludeAccessIPRule struct {
	IP   AccessAppPolicyUpdateResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyUpdateResponseIncludeAccessIPRule]
type accessAppPolicyUpdateResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessIPRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                 `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseIncludeAccessIPRuleIP]
type accessAppPolicyUpdateResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateResponseIncludeAccessIPListRule struct {
	IPList AccessAppPolicyUpdateResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyUpdateResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseIncludeAccessIPListRule]
type accessAppPolicyUpdateResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessIPListRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                         `json:"id,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessIPListRuleIPList]
type accessAppPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                   `json:"certificate,required"`
	JSON        accessAppPolicyUpdateResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessCertificateRule]
type accessAppPolicyUpdateResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessCertificateRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

// Matches an Access group.
type AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRule]
type accessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                             `json:"id,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRule]
type accessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                              `json:"connection_id,required"`
	JSON         accessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRule]
type accessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                 `json:"name,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRule]
type accessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                              `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRule]
type accessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                          `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRule]
type accessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                          `json:"attribute_value,required"`
	JSON           accessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRule]
type accessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                     `json:"token_id,required"`
	JSON    accessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                            `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRule]
type accessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                 `json:"keys_url,required"`
	JSON    accessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyUpdateResponseIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyUpdateResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyUpdateResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseIncludeAccessCountryRule]
type accessAppPolicyUpdateResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessCountryRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                       `json:"country_code,required"`
	JSON        accessAppPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessCountryRuleGeo]
type accessAppPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRule]
type accessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                           `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRule]
type accessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRule) implementsAccessAppPolicyUpdateResponseInclude() {
}

type AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                       `json:"integration_uid,required"`
	JSON           accessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyUpdateResponseRequireAccessEmailRule],
// [AccessAppPolicyUpdateResponseRequireAccessEmailListRule],
// [AccessAppPolicyUpdateResponseRequireAccessDomainRule],
// [AccessAppPolicyUpdateResponseRequireAccessEveryoneRule],
// [AccessAppPolicyUpdateResponseRequireAccessIPRule],
// [AccessAppPolicyUpdateResponseRequireAccessIPListRule],
// [AccessAppPolicyUpdateResponseRequireAccessCertificateRule],
// [AccessAppPolicyUpdateResponseRequireAccessAccessGroupRule],
// [AccessAppPolicyUpdateResponseRequireAccessAzureGroupRule],
// [AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateResponseRequireAccessOktaGroupRule],
// [AccessAppPolicyUpdateResponseRequireAccessSamlGroupRule],
// [AccessAppPolicyUpdateResponseRequireAccessServiceTokenRule],
// [AccessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateResponseRequireAccessCountryRule],
// [AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRule] or
// [AccessAppPolicyUpdateResponseRequireAccessDevicePostureRule].
type AccessAppPolicyUpdateResponseRequire interface {
	implementsAccessAppPolicyUpdateResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyUpdateResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyUpdateResponseRequireAccessEmailRule struct {
	Email AccessAppPolicyUpdateResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseRequireAccessEmailRule]
type accessAppPolicyUpdateResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessEmailRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                       `json:"email,required" format:"email"`
	JSON  accessAppPolicyUpdateResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessEmailRuleEmail]
type accessAppPolicyUpdateResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateResponseRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyUpdateResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyUpdateResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessEmailListRule]
type accessAppPolicyUpdateResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessEmailListRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                               `json:"id,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessEmailListRuleEmailList]
type accessAppPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateResponseRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyUpdateResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyUpdateResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseRequireAccessDomainRule]
type accessAppPolicyUpdateResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessDomainRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                              `json:"domain,required"`
	JSON   accessAppPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessDomainRuleEmailDomain]
type accessAppPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyUpdateResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                `json:"everyone,required"`
	JSON     accessAppPolicyUpdateResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseRequireAccessEveryoneRule]
type accessAppPolicyUpdateResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessEveryoneRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateResponseRequireAccessIPRule struct {
	IP   AccessAppPolicyUpdateResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessAppPolicyUpdateResponseRequireAccessIPRule]
type accessAppPolicyUpdateResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessIPRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                 `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseRequireAccessIPRuleIP]
type accessAppPolicyUpdateResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateResponseRequireAccessIPListRule struct {
	IPList AccessAppPolicyUpdateResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyUpdateResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseRequireAccessIPListRule]
type accessAppPolicyUpdateResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessIPListRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                         `json:"id,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessIPListRuleIPList]
type accessAppPolicyUpdateResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                   `json:"certificate,required"`
	JSON        accessAppPolicyUpdateResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessCertificateRule]
type accessAppPolicyUpdateResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessCertificateRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

// Matches an Access group.
type AccessAppPolicyUpdateResponseRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyUpdateResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyUpdateResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessAccessGroupRule]
type accessAppPolicyUpdateResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessAccessGroupRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                             `json:"id,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessAccessGroupRuleGroup]
type accessAppPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyUpdateResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessAzureGroupRule]
type accessAppPolicyUpdateResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessAzureGroupRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                              `json:"connection_id,required"`
	JSON         accessAppPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRule]
type accessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                 `json:"name,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRule]
type accessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                              `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateResponseRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyUpdateResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessOktaGroupRule]
type accessAppPolicyUpdateResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessOktaGroupRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                          `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessOktaGroupRuleOkta]
type accessAppPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateResponseRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyUpdateResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessSamlGroupRule]
type accessAppPolicyUpdateResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessSamlGroupRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                          `json:"attribute_value,required"`
	JSON           accessAppPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessSamlGroupRuleSaml]
type accessAppPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyUpdateResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessServiceTokenRule]
type accessAppPolicyUpdateResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessServiceTokenRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                     `json:"token_id,required"`
	JSON    accessAppPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                            `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRule]
type accessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                 `json:"keys_url,required"`
	JSON    accessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyUpdateResponseRequireAccessCountryRule struct {
	Geo  AccessAppPolicyUpdateResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyUpdateResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseRequireAccessCountryRule]
type accessAppPolicyUpdateResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessCountryRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                       `json:"country_code,required"`
	JSON        accessAppPolicyUpdateResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessCountryRuleGeo]
type accessAppPolicyUpdateResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRule]
type accessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                           `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyUpdateResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessDevicePostureRule]
type accessAppPolicyUpdateResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseRequireAccessDevicePostureRule) implementsAccessAppPolicyUpdateResponseRequire() {
}

type AccessAppPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                       `json:"integration_uid,required"`
	JSON           accessAppPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyDeleteResponse struct {
	// UUID
	ID   string                            `json:"id"`
	JSON accessAppPolicyDeleteResponseJSON `json:"-"`
}

// accessAppPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [AccessAppPolicyDeleteResponse]
type accessAppPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                                     `json:"session_duration"`
	UpdatedAt       time.Time                                                  `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseJSON contains the JSON
// metadata for the struct [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseJSON struct {
	ID                           apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	Name                         apijson.Field
	Precedence                   apijson.Field
	PurposeJustificationPrompt   apijson.Field
	PurposeJustificationRequired apijson.Field
	Require                      apijson.Field
	SessionDuration              apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                                                  `json:"email_list_uuid"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseApprovalGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseApprovalGroupJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseApprovalGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecision string

const (
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecisionAllow       AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecision = "allow"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecisionDeny        AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecision = "deny"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecisionNonIdentity AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecision = "non_identity"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecisionBypass      AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                        `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                       `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                         `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                          `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                  `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                            `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                      `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                       `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                          `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                       `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                   `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                   `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                              `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                     `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                          `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                    `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                        `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                       `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                         `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                          `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                  `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                            `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                      `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                       `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                          `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                       `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                   `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                   `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                              `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                     `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                          `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                    `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                        `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                       `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                         `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                          `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleIP]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                  `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                            `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                      `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                       `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                          `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                       `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                   `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                   `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                              `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                     `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                          `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                    `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                                      `json:"session_duration"`
	UpdatedAt       time.Time                                                   `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyAccessPoliciesListAccessPoliciesResponseJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponse]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseJSON struct {
	ID                           apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	Name                         apijson.Field
	Precedence                   apijson.Field
	PurposeJustificationPrompt   apijson.Field
	PurposeJustificationRequired apijson.Field
	Require                      apijson.Field
	SessionDuration              apijson.Field
	UpdatedAt                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                                                   `json:"email_list_uuid"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseApprovalGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseApprovalGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseApprovalGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecision string

const (
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecisionAllow       AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecision = "allow"
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecisionDeny        AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecision = "deny"
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecisionNonIdentity AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecision = "non_identity"
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecisionBypass      AccessAppPolicyAccessPoliciesListAccessPoliciesResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude interface {
	implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                 `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                        `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                          `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                           `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                   `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                             `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                       `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                        `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                           `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                        `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                    `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                    `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                               `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                      `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                           `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                 `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                     `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                 `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude interface {
	implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                 `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                        `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                          `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                           `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                   `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                             `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                       `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                        `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                           `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                        `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                    `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                    `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                               `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                      `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                           `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                 `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                     `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                 `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire interface {
	implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                 `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                        `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                          `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                           `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleIP]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                   `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                             `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                       `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                        `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                           `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                        `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                    `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                    `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                               `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                      `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                           `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                 `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                     `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                 `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyGetResponseEnvelope struct {
	Errors   []AccessAppPolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppPolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppPolicyGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppPolicyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppPolicyGetResponseEnvelopeJSON    `json:"-"`
}

// accessAppPolicyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseEnvelope]
type accessAppPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accessAppPolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppPolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppPolicyGetResponseEnvelopeErrors]
type accessAppPolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessAppPolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppPolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessAppPolicyGetResponseEnvelopeMessages]
type accessAppPolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyGetResponseEnvelopeSuccess bool

const (
	AccessAppPolicyGetResponseEnvelopeSuccessTrue AccessAppPolicyGetResponseEnvelopeSuccess = true
)

type AccessAppPolicyUpdateParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[AccessAppPolicyUpdateParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessAppPolicyUpdateParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]AccessAppPolicyUpdateParamsApprovalGroup] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessAppPolicyUpdateParamsExclude] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessAppPolicyUpdateParamsRequire] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyUpdateParamsDecision string

const (
	AccessAppPolicyUpdateParamsDecisionAllow       AccessAppPolicyUpdateParamsDecision = "allow"
	AccessAppPolicyUpdateParamsDecisionDeny        AccessAppPolicyUpdateParamsDecision = "deny"
	AccessAppPolicyUpdateParamsDecisionNonIdentity AccessAppPolicyUpdateParamsDecision = "non_identity"
	AccessAppPolicyUpdateParamsDecisionBypass      AccessAppPolicyUpdateParamsDecision = "bypass"
)

// Matches a specific email.
//
// Satisfied by [AccessAppPolicyUpdateParamsIncludeAccessEmailRule],
// [AccessAppPolicyUpdateParamsIncludeAccessEmailListRule],
// [AccessAppPolicyUpdateParamsIncludeAccessDomainRule],
// [AccessAppPolicyUpdateParamsIncludeAccessEveryoneRule],
// [AccessAppPolicyUpdateParamsIncludeAccessIPRule],
// [AccessAppPolicyUpdateParamsIncludeAccessIPListRule],
// [AccessAppPolicyUpdateParamsIncludeAccessCertificateRule],
// [AccessAppPolicyUpdateParamsIncludeAccessAccessGroupRule],
// [AccessAppPolicyUpdateParamsIncludeAccessAzureGroupRule],
// [AccessAppPolicyUpdateParamsIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateParamsIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateParamsIncludeAccessOktaGroupRule],
// [AccessAppPolicyUpdateParamsIncludeAccessSamlGroupRule],
// [AccessAppPolicyUpdateParamsIncludeAccessServiceTokenRule],
// [AccessAppPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateParamsIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateParamsIncludeAccessCountryRule],
// [AccessAppPolicyUpdateParamsIncludeAccessAuthenticationMethodRule],
// [AccessAppPolicyUpdateParamsIncludeAccessDevicePostureRule].
type AccessAppPolicyUpdateParamsInclude interface {
	implementsAccessAppPolicyUpdateParamsInclude()
}

// Matches a specific email.
type AccessAppPolicyUpdateParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessAppPolicyUpdateParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEmailRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessAppPolicyUpdateParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEmailListRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessAppPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessDomainRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessAppPolicyUpdateParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessEveryoneRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateParamsIncludeAccessIPRule struct {
	IP param.Field[AccessAppPolicyUpdateParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessIPRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessAppPolicyUpdateParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessIPListRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessCertificateRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

// Matches an Access group.
type AccessAppPolicyUpdateParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessAppPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAccessGroupRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessAppPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAzureGroupRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessAppPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessAppPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessAppPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessOktaGroupRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessAppPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessSamlGroupRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessAppPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessServiceTokenRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessAppPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessAppPolicyUpdateParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessAppPolicyUpdateParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessCountryRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessAppPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessAppPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsIncludeAccessDevicePostureRule) implementsAccessAppPolicyUpdateParamsInclude() {
}

type AccessAppPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessAppPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyUpdateParamsApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]interface{}] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid param.Field[string] `json:"email_list_uuid"`
}

func (r AccessAppPolicyUpdateParamsApprovalGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessAppPolicyUpdateParamsExcludeAccessEmailRule],
// [AccessAppPolicyUpdateParamsExcludeAccessEmailListRule],
// [AccessAppPolicyUpdateParamsExcludeAccessDomainRule],
// [AccessAppPolicyUpdateParamsExcludeAccessEveryoneRule],
// [AccessAppPolicyUpdateParamsExcludeAccessIPRule],
// [AccessAppPolicyUpdateParamsExcludeAccessIPListRule],
// [AccessAppPolicyUpdateParamsExcludeAccessCertificateRule],
// [AccessAppPolicyUpdateParamsExcludeAccessAccessGroupRule],
// [AccessAppPolicyUpdateParamsExcludeAccessAzureGroupRule],
// [AccessAppPolicyUpdateParamsExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateParamsExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateParamsExcludeAccessOktaGroupRule],
// [AccessAppPolicyUpdateParamsExcludeAccessSamlGroupRule],
// [AccessAppPolicyUpdateParamsExcludeAccessServiceTokenRule],
// [AccessAppPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateParamsExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateParamsExcludeAccessCountryRule],
// [AccessAppPolicyUpdateParamsExcludeAccessAuthenticationMethodRule],
// [AccessAppPolicyUpdateParamsExcludeAccessDevicePostureRule].
type AccessAppPolicyUpdateParamsExclude interface {
	implementsAccessAppPolicyUpdateParamsExclude()
}

// Matches a specific email.
type AccessAppPolicyUpdateParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessAppPolicyUpdateParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEmailRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessAppPolicyUpdateParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEmailListRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessAppPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessDomainRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessAppPolicyUpdateParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessEveryoneRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateParamsExcludeAccessIPRule struct {
	IP param.Field[AccessAppPolicyUpdateParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessIPRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessAppPolicyUpdateParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessIPListRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessCertificateRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

// Matches an Access group.
type AccessAppPolicyUpdateParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessAppPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAccessGroupRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessAppPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAzureGroupRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessAppPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessAppPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessAppPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessOktaGroupRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessAppPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessSamlGroupRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessAppPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessServiceTokenRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessAppPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessAppPolicyUpdateParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessAppPolicyUpdateParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessCountryRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessAppPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessAppPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsExcludeAccessDevicePostureRule) implementsAccessAppPolicyUpdateParamsExclude() {
}

type AccessAppPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessAppPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessAppPolicyUpdateParamsRequireAccessEmailRule],
// [AccessAppPolicyUpdateParamsRequireAccessEmailListRule],
// [AccessAppPolicyUpdateParamsRequireAccessDomainRule],
// [AccessAppPolicyUpdateParamsRequireAccessEveryoneRule],
// [AccessAppPolicyUpdateParamsRequireAccessIPRule],
// [AccessAppPolicyUpdateParamsRequireAccessIPListRule],
// [AccessAppPolicyUpdateParamsRequireAccessCertificateRule],
// [AccessAppPolicyUpdateParamsRequireAccessAccessGroupRule],
// [AccessAppPolicyUpdateParamsRequireAccessAzureGroupRule],
// [AccessAppPolicyUpdateParamsRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateParamsRequireAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateParamsRequireAccessOktaGroupRule],
// [AccessAppPolicyUpdateParamsRequireAccessSamlGroupRule],
// [AccessAppPolicyUpdateParamsRequireAccessServiceTokenRule],
// [AccessAppPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateParamsRequireAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateParamsRequireAccessCountryRule],
// [AccessAppPolicyUpdateParamsRequireAccessAuthenticationMethodRule],
// [AccessAppPolicyUpdateParamsRequireAccessDevicePostureRule].
type AccessAppPolicyUpdateParamsRequire interface {
	implementsAccessAppPolicyUpdateParamsRequire()
}

// Matches a specific email.
type AccessAppPolicyUpdateParamsRequireAccessEmailRule struct {
	Email param.Field[AccessAppPolicyUpdateParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessEmailRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessAppPolicyUpdateParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessEmailListRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessAppPolicyUpdateParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessDomainRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessAppPolicyUpdateParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessEveryoneRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateParamsRequireAccessIPRule struct {
	IP param.Field[AccessAppPolicyUpdateParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessIPRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessAppPolicyUpdateParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessIPListRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessCertificateRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

// Matches an Access group.
type AccessAppPolicyUpdateParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessAppPolicyUpdateParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessAccessGroupRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessAppPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessAzureGroupRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessAppPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessAppPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessAppPolicyUpdateParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessOktaGroupRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessAppPolicyUpdateParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessSamlGroupRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessAppPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessServiceTokenRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessAppPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessAppPolicyUpdateParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessAppPolicyUpdateParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessCountryRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessAppPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessAppPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyUpdateParamsRequireAccessDevicePostureRule) implementsAccessAppPolicyUpdateParamsRequire() {
}

type AccessAppPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessAppPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppPolicyUpdateResponseEnvelope struct {
	Errors   []AccessAppPolicyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppPolicyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppPolicyUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppPolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppPolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessAppPolicyUpdateResponseEnvelope]
type accessAppPolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessAppPolicyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppPolicyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessAppPolicyUpdateResponseEnvelopeErrors]
type accessAppPolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessAppPolicyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppPolicyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessAppPolicyUpdateResponseEnvelopeMessages]
type accessAppPolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyUpdateResponseEnvelopeSuccess bool

const (
	AccessAppPolicyUpdateResponseEnvelopeSuccessTrue AccessAppPolicyUpdateResponseEnvelopeSuccess = true
)

type AccessAppPolicyDeleteResponseEnvelope struct {
	Errors   []AccessAppPolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppPolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppPolicyDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppPolicyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessAppPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessAppPolicyDeleteResponseEnvelope]
type accessAppPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessAppPolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppPolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessAppPolicyDeleteResponseEnvelopeErrors]
type accessAppPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessAppPolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppPolicyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessAppPolicyDeleteResponseEnvelopeMessages]
type accessAppPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyDeleteResponseEnvelopeSuccess bool

const (
	AccessAppPolicyDeleteResponseEnvelopeSuccessTrue AccessAppPolicyDeleteResponseEnvelopeSuccess = true
)

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsApprovalGroup] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecision string

const (
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecisionAllow       AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecision = "allow"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecisionDeny        AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecision = "deny"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecisionNonIdentity AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecision = "non_identity"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecisionBypass      AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsDecision = "bypass"
)

// Matches a specific email.
//
// Satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAuthenticationMethodRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude()
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPRule struct {
	IP param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]interface{}] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid param.Field[string] `json:"email_list_uuid"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsApprovalGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAuthenticationMethodRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude()
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPRule struct {
	IP param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAuthenticationMethodRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire()
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule struct {
	Email param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPRule struct {
	IP param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelope struct {
	Errors   []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelope]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeErrors]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeMessages]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeSuccess bool

const (
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeSuccessTrue AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseEnvelopeSuccess = true
)

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelope struct {
	Errors   []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessAppPolicyAccessPoliciesListAccessPoliciesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelope]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeErrors]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeMessages]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeSuccess bool

const (
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeSuccessTrue AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeSuccess = true
)

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeResultInfo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
