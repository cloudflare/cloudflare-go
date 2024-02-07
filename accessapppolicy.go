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
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured Access policy.
func (r *AccessAppPolicyService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, body AccessAppPolicyUpdateParams, opts ...option.RequestOption) (res *AccessAppPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete an Access policy.
func (r *AccessAppPolicyService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, opts ...option.RequestOption) (res *AccessAppPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Access policy for an application.
func (r *AccessAppPolicyService) AccessPoliciesNewAnAccessPolicy(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessAppPolicyAccessPoliciesNewAnAccessPolicyParams, opts ...option.RequestOption) (res *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists Access policies configured for an application.
func (r *AccessAppPolicyService) AccessPoliciesListAccessPolicies(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppPolicyAccessPoliciesListAccessPoliciesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessAppPolicyGetResponse struct {
	Errors   []AccessAppPolicyGetResponseError   `json:"errors"`
	Messages []AccessAppPolicyGetResponseMessage `json:"messages"`
	Result   AccessAppPolicyGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppPolicyGetResponseSuccess `json:"success"`
	JSON    accessAppPolicyGetResponseJSON    `json:"-"`
}

// accessAppPolicyGetResponseJSON contains the JSON metadata for the struct
// [AccessAppPolicyGetResponse]
type accessAppPolicyGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accessAppPolicyGetResponseErrorJSON `json:"-"`
}

// accessAppPolicyGetResponseErrorJSON contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseError]
type accessAppPolicyGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyGetResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accessAppPolicyGetResponseMessageJSON `json:"-"`
}

// accessAppPolicyGetResponseMessageJSON contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseMessage]
type accessAppPolicyGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyGetResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyGetResponseResultApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyGetResponseResultDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyGetResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyGetResponseResultInclude `json:"include"`
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
	Require []AccessAppPolicyGetResponseResultRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                               `json:"session_duration"`
	UpdatedAt       time.Time                            `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyGetResponseResultJSON `json:"-"`
}

// accessAppPolicyGetResponseResultJSON contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResult]
type accessAppPolicyGetResponseResultJSON struct {
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

func (r *AccessAppPolicyGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyGetResponseResultApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                            `json:"email_list_uuid"`
	JSON          accessAppPolicyGetResponseResultApprovalGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseResultApprovalGroupJSON contains the JSON metadata for
// the struct [AccessAppPolicyGetResponseResultApprovalGroup]
type accessAppPolicyGetResponseResultApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyGetResponseResultDecision string

const (
	AccessAppPolicyGetResponseResultDecisionAllow       AccessAppPolicyGetResponseResultDecision = "allow"
	AccessAppPolicyGetResponseResultDecisionDeny        AccessAppPolicyGetResponseResultDecision = "deny"
	AccessAppPolicyGetResponseResultDecisionNonIdentity AccessAppPolicyGetResponseResultDecision = "non_identity"
	AccessAppPolicyGetResponseResultDecisionBypass      AccessAppPolicyGetResponseResultDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyGetResponseResultExcludeAccessEmailRule],
// [AccessAppPolicyGetResponseResultExcludeAccessEmailListRule],
// [AccessAppPolicyGetResponseResultExcludeAccessDomainRule],
// [AccessAppPolicyGetResponseResultExcludeAccessEveryoneRule],
// [AccessAppPolicyGetResponseResultExcludeAccessIPRule],
// [AccessAppPolicyGetResponseResultExcludeAccessIPListRule],
// [AccessAppPolicyGetResponseResultExcludeAccessCertificateRule],
// [AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRule],
// [AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRule],
// [AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRule],
// [AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRule],
// [AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRule],
// [AccessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyGetResponseResultExcludeAccessCountryRule],
// [AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRule].
type AccessAppPolicyGetResponseResultExclude interface {
	implementsAccessAppPolicyGetResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyGetResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyGetResponseResultExcludeAccessEmailRule struct {
	Email AccessAppPolicyGetResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultExcludeAccessEmailRule]
type accessAppPolicyGetResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessEmailRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                          `json:"email,required" format:"email"`
	JSON  accessAppPolicyGetResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessEmailRuleEmail]
type accessAppPolicyGetResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyGetResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyGetResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyGetResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessEmailListRule]
type accessAppPolicyGetResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessEmailListRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                  `json:"id,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessEmailListRuleEmailList]
type accessAppPolicyGetResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyGetResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyGetResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyGetResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessDomainRule]
type accessAppPolicyGetResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessDomainRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                 `json:"domain,required"`
	JSON   accessAppPolicyGetResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyGetResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyGetResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                   `json:"everyone,required"`
	JSON     accessAppPolicyGetResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessEveryoneRule]
type accessAppPolicyGetResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessEveryoneRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

// Matches an IP address block.
type AccessAppPolicyGetResponseResultExcludeAccessIPRule struct {
	IP   AccessAppPolicyGetResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultExcludeAccessIPRule]
type accessAppPolicyGetResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessIPRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                    `json:"ip,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultExcludeAccessIPRuleIP]
type accessAppPolicyGetResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyGetResponseResultExcludeAccessIPListRule struct {
	IPList AccessAppPolicyGetResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyGetResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessIPListRule]
type accessAppPolicyGetResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessIPListRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                            `json:"id,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessIPListRuleIPList]
type accessAppPolicyGetResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyGetResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                      `json:"certificate,required"`
	JSON        accessAppPolicyGetResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessCertificateRule]
type accessAppPolicyGetResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessCertificateRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

// Matches an Access group.
type AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRule]
type accessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                `json:"id,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRule]
type accessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                 `json:"connection_id,required"`
	JSON         accessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRule]
type accessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                    `json:"name,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRule]
type accessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                 `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRule]
type accessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                             `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRule]
type accessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                             `json:"attribute_value,required"`
	JSON           accessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRule]
type accessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                        `json:"token_id,required"`
	JSON    accessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                               `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRule]
type accessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                    `json:"keys_url,required"`
	JSON    accessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyGetResponseResultExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyGetResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyGetResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessCountryRule]
type accessAppPolicyGetResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessCountryRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                          `json:"country_code,required"`
	JSON        accessAppPolicyGetResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessCountryRuleGeo]
type accessAppPolicyGetResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRule]
type accessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                              `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRule]
type accessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRule) implementsAccessAppPolicyGetResponseResultExclude() {
}

type AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                          `json:"integration_uid,required"`
	JSON           accessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyGetResponseResultIncludeAccessEmailRule],
// [AccessAppPolicyGetResponseResultIncludeAccessEmailListRule],
// [AccessAppPolicyGetResponseResultIncludeAccessDomainRule],
// [AccessAppPolicyGetResponseResultIncludeAccessEveryoneRule],
// [AccessAppPolicyGetResponseResultIncludeAccessIPRule],
// [AccessAppPolicyGetResponseResultIncludeAccessIPListRule],
// [AccessAppPolicyGetResponseResultIncludeAccessCertificateRule],
// [AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRule],
// [AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRule],
// [AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRule],
// [AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRule],
// [AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRule],
// [AccessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyGetResponseResultIncludeAccessCountryRule],
// [AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRule].
type AccessAppPolicyGetResponseResultInclude interface {
	implementsAccessAppPolicyGetResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyGetResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyGetResponseResultIncludeAccessEmailRule struct {
	Email AccessAppPolicyGetResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultIncludeAccessEmailRule]
type accessAppPolicyGetResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessEmailRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                          `json:"email,required" format:"email"`
	JSON  accessAppPolicyGetResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessEmailRuleEmail]
type accessAppPolicyGetResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyGetResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyGetResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyGetResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessEmailListRule]
type accessAppPolicyGetResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessEmailListRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                  `json:"id,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessEmailListRuleEmailList]
type accessAppPolicyGetResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyGetResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyGetResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyGetResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessDomainRule]
type accessAppPolicyGetResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessDomainRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                 `json:"domain,required"`
	JSON   accessAppPolicyGetResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyGetResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyGetResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                   `json:"everyone,required"`
	JSON     accessAppPolicyGetResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessEveryoneRule]
type accessAppPolicyGetResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessEveryoneRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

// Matches an IP address block.
type AccessAppPolicyGetResponseResultIncludeAccessIPRule struct {
	IP   AccessAppPolicyGetResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultIncludeAccessIPRule]
type accessAppPolicyGetResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessIPRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                    `json:"ip,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultIncludeAccessIPRuleIP]
type accessAppPolicyGetResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyGetResponseResultIncludeAccessIPListRule struct {
	IPList AccessAppPolicyGetResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyGetResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessIPListRule]
type accessAppPolicyGetResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessIPListRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                            `json:"id,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessIPListRuleIPList]
type accessAppPolicyGetResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyGetResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                      `json:"certificate,required"`
	JSON        accessAppPolicyGetResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessCertificateRule]
type accessAppPolicyGetResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessCertificateRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

// Matches an Access group.
type AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRule]
type accessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                `json:"id,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRule]
type accessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                 `json:"connection_id,required"`
	JSON         accessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRule]
type accessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                    `json:"name,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRule]
type accessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                 `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRule]
type accessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                             `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRule]
type accessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                             `json:"attribute_value,required"`
	JSON           accessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRule]
type accessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                        `json:"token_id,required"`
	JSON    accessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                               `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRule]
type accessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                    `json:"keys_url,required"`
	JSON    accessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyGetResponseResultIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyGetResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyGetResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessCountryRule]
type accessAppPolicyGetResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessCountryRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                          `json:"country_code,required"`
	JSON        accessAppPolicyGetResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessCountryRuleGeo]
type accessAppPolicyGetResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRule]
type accessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                              `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRule]
type accessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRule) implementsAccessAppPolicyGetResponseResultInclude() {
}

type AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                          `json:"integration_uid,required"`
	JSON           accessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyGetResponseResultRequireAccessEmailRule],
// [AccessAppPolicyGetResponseResultRequireAccessEmailListRule],
// [AccessAppPolicyGetResponseResultRequireAccessDomainRule],
// [AccessAppPolicyGetResponseResultRequireAccessEveryoneRule],
// [AccessAppPolicyGetResponseResultRequireAccessIPRule],
// [AccessAppPolicyGetResponseResultRequireAccessIPListRule],
// [AccessAppPolicyGetResponseResultRequireAccessCertificateRule],
// [AccessAppPolicyGetResponseResultRequireAccessAccessGroupRule],
// [AccessAppPolicyGetResponseResultRequireAccessAzureGroupRule],
// [AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRule],
// [AccessAppPolicyGetResponseResultRequireAccessOktaGroupRule],
// [AccessAppPolicyGetResponseResultRequireAccessSamlGroupRule],
// [AccessAppPolicyGetResponseResultRequireAccessServiceTokenRule],
// [AccessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRule],
// [AccessAppPolicyGetResponseResultRequireAccessCountryRule],
// [AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRule] or
// [AccessAppPolicyGetResponseResultRequireAccessDevicePostureRule].
type AccessAppPolicyGetResponseResultRequire interface {
	implementsAccessAppPolicyGetResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyGetResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyGetResponseResultRequireAccessEmailRule struct {
	Email AccessAppPolicyGetResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultRequireAccessEmailRule]
type accessAppPolicyGetResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessEmailRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                          `json:"email,required" format:"email"`
	JSON  accessAppPolicyGetResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessEmailRuleEmail]
type accessAppPolicyGetResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyGetResponseResultRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyGetResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyGetResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessEmailListRule]
type accessAppPolicyGetResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessEmailListRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                  `json:"id,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessEmailListRuleEmailList]
type accessAppPolicyGetResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyGetResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyGetResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyGetResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessDomainRule]
type accessAppPolicyGetResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessDomainRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                 `json:"domain,required"`
	JSON   accessAppPolicyGetResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessDomainRuleEmailDomain]
type accessAppPolicyGetResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyGetResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                   `json:"everyone,required"`
	JSON     accessAppPolicyGetResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessEveryoneRule]
type accessAppPolicyGetResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessEveryoneRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

// Matches an IP address block.
type AccessAppPolicyGetResponseResultRequireAccessIPRule struct {
	IP   AccessAppPolicyGetResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultRequireAccessIPRule]
type accessAppPolicyGetResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessIPRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                    `json:"ip,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessAppPolicyGetResponseResultRequireAccessIPRuleIP]
type accessAppPolicyGetResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyGetResponseResultRequireAccessIPListRule struct {
	IPList AccessAppPolicyGetResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyGetResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessIPListRule]
type accessAppPolicyGetResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessIPListRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                            `json:"id,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessIPListRuleIPList]
type accessAppPolicyGetResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyGetResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                      `json:"certificate,required"`
	JSON        accessAppPolicyGetResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessCertificateRule]
type accessAppPolicyGetResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessCertificateRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

// Matches an Access group.
type AccessAppPolicyGetResponseResultRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyGetResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyGetResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessAccessGroupRule]
type accessAppPolicyGetResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessAccessGroupRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                `json:"id,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessAccessGroupRuleGroup]
type accessAppPolicyGetResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyGetResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyGetResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyGetResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessAzureGroupRule]
type accessAppPolicyGetResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessAzureGroupRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                 `json:"connection_id,required"`
	JSON         accessAppPolicyGetResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyGetResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRule]
type accessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                    `json:"name,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRule]
type accessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                 `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyGetResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyGetResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessOktaGroupRule]
type accessAppPolicyGetResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessOktaGroupRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                             `json:"email,required"`
	JSON  accessAppPolicyGetResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessOktaGroupRuleOkta]
type accessAppPolicyGetResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyGetResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyGetResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessSamlGroupRule]
type accessAppPolicyGetResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessSamlGroupRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                             `json:"attribute_value,required"`
	JSON           accessAppPolicyGetResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessSamlGroupRuleSaml]
type accessAppPolicyGetResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyGetResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyGetResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyGetResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessServiceTokenRule]
type accessAppPolicyGetResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessServiceTokenRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                        `json:"token_id,required"`
	JSON    accessAppPolicyGetResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyGetResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                               `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRule]
type accessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                    `json:"keys_url,required"`
	JSON    accessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyGetResponseResultRequireAccessCountryRule struct {
	Geo  AccessAppPolicyGetResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyGetResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessCountryRule]
type accessAppPolicyGetResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessCountryRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                          `json:"country_code,required"`
	JSON        accessAppPolicyGetResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessCountryRuleGeo]
type accessAppPolicyGetResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRule]
type accessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                              `json:"auth_method,required"`
	JSON       accessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyGetResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyGetResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyGetResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessDevicePostureRule]
type accessAppPolicyGetResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyGetResponseResultRequireAccessDevicePostureRule) implementsAccessAppPolicyGetResponseResultRequire() {
}

type AccessAppPolicyGetResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                          `json:"integration_uid,required"`
	JSON           accessAppPolicyGetResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyGetResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyGetResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyGetResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyGetResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyGetResponseSuccess bool

const (
	AccessAppPolicyGetResponseSuccessTrue AccessAppPolicyGetResponseSuccess = true
)

type AccessAppPolicyUpdateResponse struct {
	Errors   []AccessAppPolicyUpdateResponseError   `json:"errors"`
	Messages []AccessAppPolicyUpdateResponseMessage `json:"messages"`
	Result   AccessAppPolicyUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppPolicyUpdateResponseSuccess `json:"success"`
	JSON    accessAppPolicyUpdateResponseJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponse]
type accessAppPolicyUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessAppPolicyUpdateResponseErrorJSON `json:"-"`
}

// accessAppPolicyUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseError]
type accessAppPolicyUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessAppPolicyUpdateResponseMessageJSON `json:"-"`
}

// accessAppPolicyUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccessAppPolicyUpdateResponseMessage]
type accessAppPolicyUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyUpdateResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyUpdateResponseResultApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyUpdateResponseResultDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyUpdateResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyUpdateResponseResultInclude `json:"include"`
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
	Require []AccessAppPolicyUpdateResponseResultRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                  `json:"session_duration"`
	UpdatedAt       time.Time                               `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyUpdateResponseResultJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccessAppPolicyUpdateResponseResult]
type accessAppPolicyUpdateResponseResultJSON struct {
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

func (r *AccessAppPolicyUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyUpdateResponseResultApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                               `json:"email_list_uuid"`
	JSON          accessAppPolicyUpdateResponseResultApprovalGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultApprovalGroupJSON contains the JSON metadata
// for the struct [AccessAppPolicyUpdateResponseResultApprovalGroup]
type accessAppPolicyUpdateResponseResultApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyUpdateResponseResultDecision string

const (
	AccessAppPolicyUpdateResponseResultDecisionAllow       AccessAppPolicyUpdateResponseResultDecision = "allow"
	AccessAppPolicyUpdateResponseResultDecisionDeny        AccessAppPolicyUpdateResponseResultDecision = "deny"
	AccessAppPolicyUpdateResponseResultDecisionNonIdentity AccessAppPolicyUpdateResponseResultDecision = "non_identity"
	AccessAppPolicyUpdateResponseResultDecisionBypass      AccessAppPolicyUpdateResponseResultDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyUpdateResponseResultExcludeAccessEmailRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessDomainRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessEveryoneRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessIPRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessIPListRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessCertificateRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessCountryRule],
// [AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRule].
type AccessAppPolicyUpdateResponseResultExclude interface {
	implementsAccessAppPolicyUpdateResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyUpdateResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyUpdateResponseResultExcludeAccessEmailRule struct {
	Email AccessAppPolicyUpdateResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessEmailRule]
type accessAppPolicyUpdateResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessEmailRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  accessAppPolicyUpdateResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessEmailRuleEmail]
type accessAppPolicyUpdateResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRule]
type accessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleEmailList]
type accessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyUpdateResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyUpdateResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessDomainRule]
type accessAppPolicyUpdateResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessDomainRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   accessAppPolicyUpdateResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyUpdateResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyUpdateResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     accessAppPolicyUpdateResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessEveryoneRule]
type accessAppPolicyUpdateResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessEveryoneRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateResponseResultExcludeAccessIPRule struct {
	IP   AccessAppPolicyUpdateResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseResultExcludeAccessIPRule]
type accessAppPolicyUpdateResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessIPRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessIPRuleIP]
type accessAppPolicyUpdateResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateResponseResultExcludeAccessIPListRule struct {
	IPList AccessAppPolicyUpdateResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyUpdateResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessIPListRule]
type accessAppPolicyUpdateResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessIPListRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessIPListRuleIPList]
type accessAppPolicyUpdateResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        accessAppPolicyUpdateResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessCertificateRule]
type accessAppPolicyUpdateResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessCertificateRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

// Matches an Access group.
type AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRule]
type accessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRule]
type accessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         accessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRule]
type accessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRule]
type accessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRule]
type accessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRule]
type accessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           accessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRule]
type accessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    accessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRule]
type accessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    accessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyUpdateResponseResultExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyUpdateResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyUpdateResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessCountryRule]
type accessAppPolicyUpdateResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessCountryRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        accessAppPolicyUpdateResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessCountryRuleGeo]
type accessAppPolicyUpdateResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRule]
type accessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRule]
type accessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRule) implementsAccessAppPolicyUpdateResponseResultExclude() {
}

type AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           accessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyUpdateResponseResultIncludeAccessEmailRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessDomainRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessEveryoneRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessIPRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessIPListRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessCertificateRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessCountryRule],
// [AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRule] or
// [AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRule].
type AccessAppPolicyUpdateResponseResultInclude interface {
	implementsAccessAppPolicyUpdateResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyUpdateResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyUpdateResponseResultIncludeAccessEmailRule struct {
	Email AccessAppPolicyUpdateResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessEmailRule]
type accessAppPolicyUpdateResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessEmailRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  accessAppPolicyUpdateResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessEmailRuleEmail]
type accessAppPolicyUpdateResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRule]
type accessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleEmailList]
type accessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyUpdateResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyUpdateResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessDomainRule]
type accessAppPolicyUpdateResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessDomainRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   accessAppPolicyUpdateResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyUpdateResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyUpdateResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     accessAppPolicyUpdateResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessEveryoneRule]
type accessAppPolicyUpdateResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessEveryoneRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateResponseResultIncludeAccessIPRule struct {
	IP   AccessAppPolicyUpdateResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseResultIncludeAccessIPRule]
type accessAppPolicyUpdateResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessIPRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessIPRuleIP]
type accessAppPolicyUpdateResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateResponseResultIncludeAccessIPListRule struct {
	IPList AccessAppPolicyUpdateResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyUpdateResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessIPListRule]
type accessAppPolicyUpdateResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessIPListRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessIPListRuleIPList]
type accessAppPolicyUpdateResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        accessAppPolicyUpdateResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessCertificateRule]
type accessAppPolicyUpdateResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessCertificateRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

// Matches an Access group.
type AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRule]
type accessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRule]
type accessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         accessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRule]
type accessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRule]
type accessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRule]
type accessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRule]
type accessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           accessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRule]
type accessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    accessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRule]
type accessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    accessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyUpdateResponseResultIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyUpdateResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyUpdateResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessCountryRule]
type accessAppPolicyUpdateResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessCountryRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        accessAppPolicyUpdateResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessCountryRuleGeo]
type accessAppPolicyUpdateResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRule]
type accessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRule]
type accessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRule) implementsAccessAppPolicyUpdateResponseResultInclude() {
}

type AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           accessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessAppPolicyUpdateResponseResultRequireAccessEmailRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessEmailListRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessDomainRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessEveryoneRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessIPRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessIPListRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessCertificateRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessCountryRule],
// [AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRule] or
// [AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRule].
type AccessAppPolicyUpdateResponseResultRequire interface {
	implementsAccessAppPolicyUpdateResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyUpdateResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyUpdateResponseResultRequireAccessEmailRule struct {
	Email AccessAppPolicyUpdateResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessEmailRule]
type accessAppPolicyUpdateResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessEmailRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  accessAppPolicyUpdateResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessEmailRuleEmail]
type accessAppPolicyUpdateResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyUpdateResponseResultRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyUpdateResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyUpdateResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessEmailListRule]
type accessAppPolicyUpdateResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessEmailListRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessEmailListRuleEmailList]
type accessAppPolicyUpdateResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyUpdateResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyUpdateResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyUpdateResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessDomainRule]
type accessAppPolicyUpdateResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessDomainRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   accessAppPolicyUpdateResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessDomainRuleEmailDomain]
type accessAppPolicyUpdateResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyUpdateResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     accessAppPolicyUpdateResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessEveryoneRule]
type accessAppPolicyUpdateResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessEveryoneRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

// Matches an IP address block.
type AccessAppPolicyUpdateResponseResultRequireAccessIPRule struct {
	IP   AccessAppPolicyUpdateResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessAppPolicyUpdateResponseResultRequireAccessIPRule]
type accessAppPolicyUpdateResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessIPRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessIPRuleIP]
type accessAppPolicyUpdateResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyUpdateResponseResultRequireAccessIPListRule struct {
	IPList AccessAppPolicyUpdateResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyUpdateResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessIPListRule]
type accessAppPolicyUpdateResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessIPListRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessIPListRuleIPList]
type accessAppPolicyUpdateResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyUpdateResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        accessAppPolicyUpdateResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessCertificateRule]
type accessAppPolicyUpdateResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessCertificateRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

// Matches an Access group.
type AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRule]
type accessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleGroup]
type accessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRule]
type accessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         accessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRule]
type accessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRule]
type accessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRule]
type accessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  accessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleOkta]
type accessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRule]
type accessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           accessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleSaml]
type accessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRule]
type accessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    accessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRule]
type accessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    accessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyUpdateResponseResultRequireAccessCountryRule struct {
	Geo  AccessAppPolicyUpdateResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyUpdateResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessCountryRule]
type accessAppPolicyUpdateResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessCountryRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        accessAppPolicyUpdateResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessCountryRuleGeo]
type accessAppPolicyUpdateResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRule]
type accessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       accessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRule]
type accessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRule) implementsAccessAppPolicyUpdateResponseResultRequire() {
}

type AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           accessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyUpdateResponseSuccess bool

const (
	AccessAppPolicyUpdateResponseSuccessTrue AccessAppPolicyUpdateResponseSuccess = true
)

type AccessAppPolicyDeleteResponse struct {
	Errors   []AccessAppPolicyDeleteResponseError   `json:"errors"`
	Messages []AccessAppPolicyDeleteResponseMessage `json:"messages"`
	Result   AccessAppPolicyDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppPolicyDeleteResponseSuccess `json:"success"`
	JSON    accessAppPolicyDeleteResponseJSON    `json:"-"`
}

// accessAppPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [AccessAppPolicyDeleteResponse]
type accessAppPolicyDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accessAppPolicyDeleteResponseErrorJSON `json:"-"`
}

// accessAppPolicyDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccessAppPolicyDeleteResponseError]
type accessAppPolicyDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessAppPolicyDeleteResponseMessageJSON `json:"-"`
}

// accessAppPolicyDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccessAppPolicyDeleteResponseMessage]
type accessAppPolicyDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyDeleteResponseResult struct {
	// UUID
	ID   string                                  `json:"id"`
	JSON accessAppPolicyDeleteResponseResultJSON `json:"-"`
}

// accessAppPolicyDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccessAppPolicyDeleteResponseResult]
type accessAppPolicyDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyDeleteResponseSuccess bool

const (
	AccessAppPolicyDeleteResponseSuccessTrue AccessAppPolicyDeleteResponseSuccess = true
)

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse struct {
	Errors   []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseError   `json:"errors"`
	Messages []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseMessage `json:"messages"`
	Result   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseSuccess `json:"success"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseJSON contains the JSON
// metadata for the struct [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseErrorJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseErrorJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseError]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseMessageJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseMessageJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseMessage]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude `json:"include"`
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
	Require []AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                                           `json:"session_duration"`
	UpdatedAt       time.Time                                                        `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResult]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultJSON struct {
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

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                                                        `json:"email_list_uuid"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultApprovalGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultApprovalGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultApprovalGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecision string

const (
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecisionAllow       AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecision = "allow"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecisionDeny        AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecision = "deny"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecisionNonIdentity AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecision = "non_identity"
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecisionBypass      AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                      `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                              `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                             `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                               `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                        `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                                  `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                            `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                             `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                             `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                         `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                         `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                    `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                           `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                      `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                          `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                      `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                      `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                              `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                             `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                               `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                        `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                                  `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                            `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                             `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                             `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                         `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                         `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                    `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                           `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                      `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                          `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                      `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRule],
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire interface {
	implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                      `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                              `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                             `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                               `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleIP]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                        `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                                                  `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                            `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                             `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                             `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                         `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                         `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                    `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                           `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                      `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                          `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                      `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseSuccess bool

const (
	AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseSuccessTrue AccessAppPolicyAccessPoliciesNewAnAccessPolicyResponseSuccess = true
)

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponse struct {
	Errors     []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseError    `json:"errors"`
	Messages   []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseMessage  `json:"messages"`
	Result     []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResult   `json:"result"`
	ResultInfo AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessAppPolicyAccessPoliciesListAccessPoliciesResponseSuccess `json:"success"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseJSON contains the JSON
// metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponse]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseError struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseErrorJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseErrorJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseError]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseMessage struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseMessageJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseMessageJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseMessage]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude `json:"include"`
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
	Require []AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                                            `json:"session_duration"`
	UpdatedAt       time.Time                                                         `json:"updated_at" format:"date-time"`
	JSON            accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultJSON contains the
// JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResult]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultJSON struct {
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

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                                                         `json:"email_list_uuid"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultApprovalGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultApprovalGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultApprovalGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecision string

const (
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecisionAllow       AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecision = "allow"
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecisionDeny        AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecision = "deny"
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecisionNonIdentity AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecision = "non_identity"
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecisionBypass      AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude interface {
	implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                       `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                               `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                              `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                 `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                                   `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                             `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                              `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                 `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                              `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                          `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                          `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                     `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                            `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                 `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                       `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                           `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                       `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude interface {
	implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                       `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                               `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                              `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                 `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleIP]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                                   `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                             `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                              `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                 `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                              `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                          `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                          `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                     `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                            `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                 `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                       `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                           `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInclude() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                       `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRule],
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRule]
// or
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRule].
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire interface {
	implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRule struct {
	Email AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                       `json:"email,required" format:"email"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleEmail]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRule struct {
	EmailList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                               `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleEmailList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                              `json:"domain,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleEmailDomain]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                `json:"everyone,required"`
	JSON     accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessEveryoneRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

// Matches an IP address block.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRule struct {
	IP   AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                 `json:"ip,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleIP]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRule struct {
	IPList AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleIPList]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                                                   `json:"certificate,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCertificateRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

// Matches an Access group.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRule struct {
	Group AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                             `json:"id,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleGroup]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                              `json:"connection_id,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                 `json:"name,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                              `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                          `json:"email,required"`
	JSON  accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleOkta]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                          `json:"attribute_value,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleSaml]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                     `json:"token_id,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                            `json:"any_valid_service_token,required"`
	JSON                 accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                 `json:"keys_url,required"`
	JSON    accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRule struct {
	Geo  AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                       `json:"country_code,required"`
	JSON        accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleGeo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                           `json:"auth_method,required"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRule]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRule) implementsAccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequire() {
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                       `json:"integration_uid,required"`
	JSON           accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInfoJSON `json:"-"`
}

// accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInfo]
type accessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppPolicyAccessPoliciesListAccessPoliciesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppPolicyAccessPoliciesListAccessPoliciesResponseSuccess bool

const (
	AccessAppPolicyAccessPoliciesListAccessPoliciesResponseSuccessTrue AccessAppPolicyAccessPoliciesListAccessPoliciesResponseSuccess = true
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
