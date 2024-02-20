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

// AccessApplicationPolicyService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessApplicationPolicyService] method instead.
type AccessApplicationPolicyService struct {
	Options []option.RequestOption
}

// NewAccessApplicationPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessApplicationPolicyService(opts ...option.RequestOption) (r *AccessApplicationPolicyService) {
	r = &AccessApplicationPolicyService{}
	r.Options = opts
	return
}

// Create a new Access policy for an application.
func (r *AccessApplicationPolicyService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *AccessApplicationPolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Access policies configured for an application.
func (r *AccessApplicationPolicyService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *[]AccessApplicationPolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an Access policy.
func (r *AccessApplicationPolicyService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, opts ...option.RequestOption) (res *AccessApplicationPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access policy.
func (r *AccessApplicationPolicyService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, opts ...option.RequestOption) (res *AccessApplicationPolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured Access policy.
func (r *AccessApplicationPolicyService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid1 string, uuid string, body AccessApplicationPolicyReplaceParams, opts ...option.RequestOption) (res *AccessApplicationPolicyReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessApplicationPolicyNewResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessApplicationPolicyNewResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessApplicationPolicyNewResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessApplicationPolicyNewResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessApplicationPolicyNewResponseInclude `json:"include"`
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
	Require []AccessApplicationPolicyNewResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                 `json:"session_duration"`
	UpdatedAt       time.Time                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyNewResponseJSON `json:"-"`
}

// accessApplicationPolicyNewResponseJSON contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponse]
type accessApplicationPolicyNewResponseJSON struct {
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

func (r *AccessApplicationPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyNewResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                              `json:"email_list_uuid"`
	JSON          accessApplicationPolicyNewResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyNewResponseApprovalGroupJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyNewResponseApprovalGroup]
type accessApplicationPolicyNewResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyNewResponseDecision string

const (
	AccessApplicationPolicyNewResponseDecisionAllow       AccessApplicationPolicyNewResponseDecision = "allow"
	AccessApplicationPolicyNewResponseDecisionDeny        AccessApplicationPolicyNewResponseDecision = "deny"
	AccessApplicationPolicyNewResponseDecisionNonIdentity AccessApplicationPolicyNewResponseDecision = "non_identity"
	AccessApplicationPolicyNewResponseDecisionBypass      AccessApplicationPolicyNewResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyNewResponseExcludeAccessEmailRule],
// [AccessApplicationPolicyNewResponseExcludeAccessEmailListRule],
// [AccessApplicationPolicyNewResponseExcludeAccessDomainRule],
// [AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule],
// [AccessApplicationPolicyNewResponseExcludeAccessIPRule],
// [AccessApplicationPolicyNewResponseExcludeAccessIPListRule],
// [AccessApplicationPolicyNewResponseExcludeAccessCertificateRule],
// [AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyNewResponseExcludeAccessCountryRule],
// [AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyNewResponseExclude interface {
	implementsAccessApplicationPolicyNewResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyNewResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyNewResponseExcludeAccessEmailRule struct {
	Email AccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessEmailRule]
type accessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessEmailRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail]
type accessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyNewResponseExcludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessEmailListRule]
type accessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessEmailListRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList]
type accessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyNewResponseExcludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessDomainRule]
type accessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessDomainRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   accessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     accessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule]
type accessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewResponseExcludeAccessIPRule struct {
	IP   AccessApplicationPolicyNewResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyNewResponseExcludeAccessIPRule]
type accessApplicationPolicyNewResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessIPRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessIPRuleIP]
type accessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyNewResponseExcludeAccessIPListRule struct {
	IPList AccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessIPListRule]
type accessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessIPListRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList]
type accessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyNewResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        accessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessCertificateRule]
type accessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessCertificateRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule]
type accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule]
type accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule]
type accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule]
type accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule]
type accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule]
type accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule]
type accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule]
type accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyNewResponseExcludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessCountryRule]
type accessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessCountryRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        accessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo]
type accessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule]
type accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule]
type accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyNewResponseExclude() {
}

type AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyNewResponseIncludeAccessEmailRule],
// [AccessApplicationPolicyNewResponseIncludeAccessEmailListRule],
// [AccessApplicationPolicyNewResponseIncludeAccessDomainRule],
// [AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule],
// [AccessApplicationPolicyNewResponseIncludeAccessIPRule],
// [AccessApplicationPolicyNewResponseIncludeAccessIPListRule],
// [AccessApplicationPolicyNewResponseIncludeAccessCertificateRule],
// [AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyNewResponseIncludeAccessCountryRule],
// [AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyNewResponseInclude interface {
	implementsAccessApplicationPolicyNewResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyNewResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyNewResponseIncludeAccessEmailRule struct {
	Email AccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessEmailRule]
type accessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessEmailRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail]
type accessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyNewResponseIncludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessEmailListRule]
type accessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessEmailListRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList]
type accessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyNewResponseIncludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessDomainRule]
type accessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessDomainRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   accessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     accessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule]
type accessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewResponseIncludeAccessIPRule struct {
	IP   AccessApplicationPolicyNewResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyNewResponseIncludeAccessIPRule]
type accessApplicationPolicyNewResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessIPRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessIPRuleIP]
type accessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyNewResponseIncludeAccessIPListRule struct {
	IPList AccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessIPListRule]
type accessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessIPListRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList]
type accessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyNewResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        accessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessCertificateRule]
type accessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessCertificateRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule]
type accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule]
type accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule]
type accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule]
type accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule]
type accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule]
type accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule]
type accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule]
type accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyNewResponseIncludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessCountryRule]
type accessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessCountryRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        accessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo]
type accessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule]
type accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule]
type accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyNewResponseInclude() {
}

type AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyNewResponseRequireAccessEmailRule],
// [AccessApplicationPolicyNewResponseRequireAccessEmailListRule],
// [AccessApplicationPolicyNewResponseRequireAccessDomainRule],
// [AccessApplicationPolicyNewResponseRequireAccessEveryoneRule],
// [AccessApplicationPolicyNewResponseRequireAccessIPRule],
// [AccessApplicationPolicyNewResponseRequireAccessIPListRule],
// [AccessApplicationPolicyNewResponseRequireAccessCertificateRule],
// [AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule],
// [AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule],
// [AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule],
// [AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule],
// [AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule],
// [AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyNewResponseRequireAccessCountryRule],
// [AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyNewResponseRequire interface {
	implementsAccessApplicationPolicyNewResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyNewResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyNewResponseRequireAccessEmailRule struct {
	Email AccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessEmailRule]
type accessApplicationPolicyNewResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessEmailRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail]
type accessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyNewResponseRequireAccessEmailListRule struct {
	EmailList AccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessEmailListRule]
type accessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessEmailListRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList]
type accessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyNewResponseRequireAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyNewResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessDomainRule]
type accessApplicationPolicyNewResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessDomainRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   accessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain]
type accessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyNewResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     accessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessEveryoneRule]
type accessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessEveryoneRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewResponseRequireAccessIPRule struct {
	IP   AccessApplicationPolicyNewResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyNewResponseRequireAccessIPRule]
type accessApplicationPolicyNewResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessIPRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessIPRuleIP]
type accessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyNewResponseRequireAccessIPListRule struct {
	IPList AccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyNewResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessIPListRule]
type accessApplicationPolicyNewResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessIPListRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList]
type accessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyNewResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        accessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessCertificateRule]
type accessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessCertificateRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule struct {
	Group AccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule]
type accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup]
type accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule]
type accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule]
type accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule]
type accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule]
type accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta]
type accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule]
type accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml]
type accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule]
type accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule]
type accessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule]
type accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyNewResponseRequireAccessCountryRule struct {
	Geo  AccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyNewResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessCountryRule]
type accessApplicationPolicyNewResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessCountryRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        accessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo]
type accessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule]
type accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule]
type accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule) implementsAccessApplicationPolicyNewResponseRequire() {
}

type AccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyListResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessApplicationPolicyListResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessApplicationPolicyListResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessApplicationPolicyListResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessApplicationPolicyListResponseInclude `json:"include"`
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
	Require []AccessApplicationPolicyListResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                  `json:"session_duration"`
	UpdatedAt       time.Time                               `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyListResponseJSON `json:"-"`
}

// accessApplicationPolicyListResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyListResponse]
type accessApplicationPolicyListResponseJSON struct {
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

func (r *AccessApplicationPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyListResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                               `json:"email_list_uuid"`
	JSON          accessApplicationPolicyListResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyListResponseApprovalGroupJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyListResponseApprovalGroup]
type accessApplicationPolicyListResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyListResponseDecision string

const (
	AccessApplicationPolicyListResponseDecisionAllow       AccessApplicationPolicyListResponseDecision = "allow"
	AccessApplicationPolicyListResponseDecisionDeny        AccessApplicationPolicyListResponseDecision = "deny"
	AccessApplicationPolicyListResponseDecisionNonIdentity AccessApplicationPolicyListResponseDecision = "non_identity"
	AccessApplicationPolicyListResponseDecisionBypass      AccessApplicationPolicyListResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyListResponseExcludeAccessEmailRule],
// [AccessApplicationPolicyListResponseExcludeAccessEmailListRule],
// [AccessApplicationPolicyListResponseExcludeAccessDomainRule],
// [AccessApplicationPolicyListResponseExcludeAccessEveryoneRule],
// [AccessApplicationPolicyListResponseExcludeAccessIPRule],
// [AccessApplicationPolicyListResponseExcludeAccessIPListRule],
// [AccessApplicationPolicyListResponseExcludeAccessCertificateRule],
// [AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyListResponseExcludeAccessCountryRule],
// [AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyListResponseExclude interface {
	implementsAccessApplicationPolicyListResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyListResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyListResponseExcludeAccessEmailRule struct {
	Email AccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyListResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessEmailRule]
type accessApplicationPolicyListResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessEmailRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail]
type accessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyListResponseExcludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessEmailListRule]
type accessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessEmailListRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList]
type accessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyListResponseExcludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyListResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessDomainRule]
type accessApplicationPolicyListResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessDomainRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   accessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyListResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     accessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessEveryoneRule]
type accessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessEveryoneRule) implementsAccessApplicationPolicyListResponseExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyListResponseExcludeAccessIPRule struct {
	IP   AccessApplicationPolicyListResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyListResponseExcludeAccessIPRule]
type accessApplicationPolicyListResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessIPRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessIPRuleIP]
type accessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyListResponseExcludeAccessIPListRule struct {
	IPList AccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyListResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessIPListRule]
type accessApplicationPolicyListResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessIPListRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList]
type accessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyListResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        accessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessCertificateRule]
type accessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessCertificateRule) implementsAccessApplicationPolicyListResponseExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule]
type accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule]
type accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule]
type accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule]
type accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule]
type accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule]
type accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule]
type accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyListResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule]
type accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyListResponseExcludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyListResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessCountryRule]
type accessApplicationPolicyListResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessCountryRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        accessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo]
type accessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule]
type accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule]
type accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyListResponseExclude() {
}

type AccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyListResponseIncludeAccessEmailRule],
// [AccessApplicationPolicyListResponseIncludeAccessEmailListRule],
// [AccessApplicationPolicyListResponseIncludeAccessDomainRule],
// [AccessApplicationPolicyListResponseIncludeAccessEveryoneRule],
// [AccessApplicationPolicyListResponseIncludeAccessIPRule],
// [AccessApplicationPolicyListResponseIncludeAccessIPListRule],
// [AccessApplicationPolicyListResponseIncludeAccessCertificateRule],
// [AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyListResponseIncludeAccessCountryRule],
// [AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyListResponseInclude interface {
	implementsAccessApplicationPolicyListResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyListResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyListResponseIncludeAccessEmailRule struct {
	Email AccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyListResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessEmailRule]
type accessApplicationPolicyListResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessEmailRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail]
type accessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyListResponseIncludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessEmailListRule]
type accessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessEmailListRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList]
type accessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyListResponseIncludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyListResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessDomainRule]
type accessApplicationPolicyListResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessDomainRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   accessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyListResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     accessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessEveryoneRule]
type accessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessEveryoneRule) implementsAccessApplicationPolicyListResponseInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyListResponseIncludeAccessIPRule struct {
	IP   AccessApplicationPolicyListResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyListResponseIncludeAccessIPRule]
type accessApplicationPolicyListResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessIPRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessIPRuleIP]
type accessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyListResponseIncludeAccessIPListRule struct {
	IPList AccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyListResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessIPListRule]
type accessApplicationPolicyListResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessIPListRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList]
type accessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyListResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        accessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessCertificateRule]
type accessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessCertificateRule) implementsAccessApplicationPolicyListResponseInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule]
type accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule]
type accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule]
type accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule]
type accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule]
type accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule]
type accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule]
type accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyListResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule]
type accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyListResponseIncludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyListResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessCountryRule]
type accessApplicationPolicyListResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessCountryRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        accessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo]
type accessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule]
type accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule]
type accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyListResponseInclude() {
}

type AccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyListResponseRequireAccessEmailRule],
// [AccessApplicationPolicyListResponseRequireAccessEmailListRule],
// [AccessApplicationPolicyListResponseRequireAccessDomainRule],
// [AccessApplicationPolicyListResponseRequireAccessEveryoneRule],
// [AccessApplicationPolicyListResponseRequireAccessIPRule],
// [AccessApplicationPolicyListResponseRequireAccessIPListRule],
// [AccessApplicationPolicyListResponseRequireAccessCertificateRule],
// [AccessApplicationPolicyListResponseRequireAccessAccessGroupRule],
// [AccessApplicationPolicyListResponseRequireAccessAzureGroupRule],
// [AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyListResponseRequireAccessOktaGroupRule],
// [AccessApplicationPolicyListResponseRequireAccessSamlGroupRule],
// [AccessApplicationPolicyListResponseRequireAccessServiceTokenRule],
// [AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyListResponseRequireAccessCountryRule],
// [AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyListResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyListResponseRequire interface {
	implementsAccessApplicationPolicyListResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyListResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyListResponseRequireAccessEmailRule struct {
	Email AccessApplicationPolicyListResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyListResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessEmailRule]
type accessApplicationPolicyListResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessEmailRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessEmailRuleEmail]
type accessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyListResponseRequireAccessEmailListRule struct {
	EmailList AccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyListResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessEmailListRule]
type accessApplicationPolicyListResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessEmailListRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON accessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList]
type accessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyListResponseRequireAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyListResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessDomainRule]
type accessApplicationPolicyListResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessDomainRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   accessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain]
type accessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyListResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     accessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessEveryoneRule]
type accessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessEveryoneRule) implementsAccessApplicationPolicyListResponseRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyListResponseRequireAccessIPRule struct {
	IP   AccessApplicationPolicyListResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyListResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyListResponseRequireAccessIPRule]
type accessApplicationPolicyListResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessIPRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON accessApplicationPolicyListResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessIPRuleIP]
type accessApplicationPolicyListResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyListResponseRequireAccessIPListRule struct {
	IPList AccessApplicationPolicyListResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyListResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessIPListRule]
type accessApplicationPolicyListResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessIPListRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON accessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessIPListRuleIPList]
type accessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyListResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        accessApplicationPolicyListResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessCertificateRule]
type accessApplicationPolicyListResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessCertificateRule) implementsAccessApplicationPolicyListResponseRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyListResponseRequireAccessAccessGroupRule struct {
	Group AccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessAccessGroupRule]
type accessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessAccessGroupRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON accessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup]
type accessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyListResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessAzureGroupRule]
type accessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessAzureGroupRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         accessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule]
type accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule]
type accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyListResponseRequireAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessOktaGroupRule]
type accessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessOktaGroupRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  accessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta]
type accessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyListResponseRequireAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessSamlGroupRule]
type accessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessSamlGroupRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           accessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml]
type accessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyListResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessServiceTokenRule]
type accessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessServiceTokenRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    accessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule]
type accessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyListResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule]
type accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyListResponseRequireAccessCountryRule struct {
	Geo  AccessApplicationPolicyListResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyListResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessCountryRule]
type accessApplicationPolicyListResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessCountryRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        accessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessCountryRuleGeo]
type accessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule]
type accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyListResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessDevicePostureRule]
type accessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyListResponseRequireAccessDevicePostureRule) implementsAccessApplicationPolicyListResponseRequire() {
}

type AccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           accessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyDeleteResponse struct {
	// UUID
	ID   string                                    `json:"id"`
	JSON accessApplicationPolicyDeleteResponseJSON `json:"-"`
}

// accessApplicationPolicyDeleteResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyDeleteResponse]
type accessApplicationPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyGetResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessApplicationPolicyGetResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessApplicationPolicyGetResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessApplicationPolicyGetResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessApplicationPolicyGetResponseInclude `json:"include"`
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
	Require []AccessApplicationPolicyGetResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                 `json:"session_duration"`
	UpdatedAt       time.Time                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyGetResponseJSON `json:"-"`
}

// accessApplicationPolicyGetResponseJSON contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponse]
type accessApplicationPolicyGetResponseJSON struct {
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

func (r *AccessApplicationPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyGetResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                              `json:"email_list_uuid"`
	JSON          accessApplicationPolicyGetResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyGetResponseApprovalGroupJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyGetResponseApprovalGroup]
type accessApplicationPolicyGetResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyGetResponseDecision string

const (
	AccessApplicationPolicyGetResponseDecisionAllow       AccessApplicationPolicyGetResponseDecision = "allow"
	AccessApplicationPolicyGetResponseDecisionDeny        AccessApplicationPolicyGetResponseDecision = "deny"
	AccessApplicationPolicyGetResponseDecisionNonIdentity AccessApplicationPolicyGetResponseDecision = "non_identity"
	AccessApplicationPolicyGetResponseDecisionBypass      AccessApplicationPolicyGetResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyGetResponseExcludeAccessEmailRule],
// [AccessApplicationPolicyGetResponseExcludeAccessEmailListRule],
// [AccessApplicationPolicyGetResponseExcludeAccessDomainRule],
// [AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule],
// [AccessApplicationPolicyGetResponseExcludeAccessIPRule],
// [AccessApplicationPolicyGetResponseExcludeAccessIPListRule],
// [AccessApplicationPolicyGetResponseExcludeAccessCertificateRule],
// [AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyGetResponseExcludeAccessCountryRule],
// [AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyGetResponseExclude interface {
	implementsAccessApplicationPolicyGetResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyGetResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyGetResponseExcludeAccessEmailRule struct {
	Email AccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessEmailRule]
type accessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessEmailRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail]
type accessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyGetResponseExcludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessEmailListRule]
type accessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessEmailListRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList]
type accessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyGetResponseExcludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessDomainRule]
type accessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessDomainRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   accessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     accessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule]
type accessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyGetResponseExcludeAccessIPRule struct {
	IP   AccessApplicationPolicyGetResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyGetResponseExcludeAccessIPRule]
type accessApplicationPolicyGetResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessIPRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessIPRuleIP]
type accessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyGetResponseExcludeAccessIPListRule struct {
	IPList AccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessIPListRule]
type accessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessIPListRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList]
type accessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyGetResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        accessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessCertificateRule]
type accessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessCertificateRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule]
type accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule]
type accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule]
type accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule]
type accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule]
type accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule]
type accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule]
type accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule]
type accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyGetResponseExcludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessCountryRule]
type accessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessCountryRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        accessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo]
type accessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule]
type accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule]
type accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyGetResponseExclude() {
}

type AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyGetResponseIncludeAccessEmailRule],
// [AccessApplicationPolicyGetResponseIncludeAccessEmailListRule],
// [AccessApplicationPolicyGetResponseIncludeAccessDomainRule],
// [AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule],
// [AccessApplicationPolicyGetResponseIncludeAccessIPRule],
// [AccessApplicationPolicyGetResponseIncludeAccessIPListRule],
// [AccessApplicationPolicyGetResponseIncludeAccessCertificateRule],
// [AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyGetResponseIncludeAccessCountryRule],
// [AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyGetResponseInclude interface {
	implementsAccessApplicationPolicyGetResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyGetResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyGetResponseIncludeAccessEmailRule struct {
	Email AccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessEmailRule]
type accessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessEmailRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail]
type accessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyGetResponseIncludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessEmailListRule]
type accessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessEmailListRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList]
type accessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyGetResponseIncludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessDomainRule]
type accessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessDomainRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   accessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     accessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule]
type accessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyGetResponseIncludeAccessIPRule struct {
	IP   AccessApplicationPolicyGetResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyGetResponseIncludeAccessIPRule]
type accessApplicationPolicyGetResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessIPRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessIPRuleIP]
type accessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyGetResponseIncludeAccessIPListRule struct {
	IPList AccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessIPListRule]
type accessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessIPListRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList]
type accessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyGetResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        accessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessCertificateRule]
type accessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessCertificateRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule]
type accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule]
type accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule]
type accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule]
type accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule]
type accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule]
type accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule]
type accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule]
type accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyGetResponseIncludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessCountryRule]
type accessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessCountryRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        accessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo]
type accessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule]
type accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule]
type accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyGetResponseInclude() {
}

type AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessApplicationPolicyGetResponseRequireAccessEmailRule],
// [AccessApplicationPolicyGetResponseRequireAccessEmailListRule],
// [AccessApplicationPolicyGetResponseRequireAccessDomainRule],
// [AccessApplicationPolicyGetResponseRequireAccessEveryoneRule],
// [AccessApplicationPolicyGetResponseRequireAccessIPRule],
// [AccessApplicationPolicyGetResponseRequireAccessIPListRule],
// [AccessApplicationPolicyGetResponseRequireAccessCertificateRule],
// [AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule],
// [AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule],
// [AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule],
// [AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule],
// [AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule],
// [AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyGetResponseRequireAccessCountryRule],
// [AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyGetResponseRequire interface {
	implementsAccessApplicationPolicyGetResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyGetResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyGetResponseRequireAccessEmailRule struct {
	Email AccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessEmailRule]
type accessApplicationPolicyGetResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessEmailRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail]
type accessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyGetResponseRequireAccessEmailListRule struct {
	EmailList AccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessEmailListRule]
type accessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessEmailListRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList]
type accessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyGetResponseRequireAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyGetResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessDomainRule]
type accessApplicationPolicyGetResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessDomainRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   accessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain]
type accessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyGetResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     accessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessEveryoneRule]
type accessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessEveryoneRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyGetResponseRequireAccessIPRule struct {
	IP   AccessApplicationPolicyGetResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyGetResponseRequireAccessIPRule]
type accessApplicationPolicyGetResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessIPRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessIPRuleIP]
type accessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyGetResponseRequireAccessIPListRule struct {
	IPList AccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyGetResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessIPListRule]
type accessApplicationPolicyGetResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessIPListRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList]
type accessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyGetResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        accessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessCertificateRule]
type accessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessCertificateRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule struct {
	Group AccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule]
type accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup]
type accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule]
type accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule]
type accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule]
type accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule]
type accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta]
type accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule]
type accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml]
type accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule]
type accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule]
type accessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule]
type accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyGetResponseRequireAccessCountryRule struct {
	Geo  AccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyGetResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessCountryRule]
type accessApplicationPolicyGetResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessCountryRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        accessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo]
type accessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule]
type accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule]
type accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule) implementsAccessApplicationPolicyGetResponseRequire() {
}

type AccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyReplaceResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessApplicationPolicyReplaceResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessApplicationPolicyReplaceResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessApplicationPolicyReplaceResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessApplicationPolicyReplaceResponseInclude `json:"include"`
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
	Require []AccessApplicationPolicyReplaceResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                     `json:"session_duration"`
	UpdatedAt       time.Time                                  `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyReplaceResponseJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyReplaceResponse]
type accessApplicationPolicyReplaceResponseJSON struct {
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

func (r *AccessApplicationPolicyReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyReplaceResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                                                  `json:"email_list_uuid"`
	JSON          accessApplicationPolicyReplaceResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseApprovalGroupJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyReplaceResponseApprovalGroup]
type accessApplicationPolicyReplaceResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyReplaceResponseDecision string

const (
	AccessApplicationPolicyReplaceResponseDecisionAllow       AccessApplicationPolicyReplaceResponseDecision = "allow"
	AccessApplicationPolicyReplaceResponseDecisionDeny        AccessApplicationPolicyReplaceResponseDecision = "deny"
	AccessApplicationPolicyReplaceResponseDecisionNonIdentity AccessApplicationPolicyReplaceResponseDecision = "non_identity"
	AccessApplicationPolicyReplaceResponseDecisionBypass      AccessApplicationPolicyReplaceResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [AccessApplicationPolicyReplaceResponseExcludeAccessEmailRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessDomainRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessEveryoneRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessIPRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessIPListRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessCertificateRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessCountryRule],
// [AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyReplaceResponseExclude interface {
	implementsAccessApplicationPolicyReplaceResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyReplaceResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyReplaceResponseExcludeAccessEmailRule struct {
	Email AccessApplicationPolicyReplaceResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessEmailRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessEmailRule]
type accessApplicationPolicyReplaceResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessEmailRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyReplaceResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessEmailRuleEmail]
type accessApplicationPolicyReplaceResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRule]
type accessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                        `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleEmailList]
type accessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyReplaceResponseExcludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyReplaceResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyReplaceResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessDomainRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessDomainRule]
type accessApplicationPolicyReplaceResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessDomainRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                       `json:"domain,required"`
	JSON   accessApplicationPolicyReplaceResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyReplaceResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyReplaceResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                         `json:"everyone,required"`
	JSON     accessApplicationPolicyReplaceResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessEveryoneRule]
type accessApplicationPolicyReplaceResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessEveryoneRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyReplaceResponseExcludeAccessIPRule struct {
	IP   AccessApplicationPolicyReplaceResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessIPRule]
type accessApplicationPolicyReplaceResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessIPRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                          `json:"ip,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessIPRuleIP]
type accessApplicationPolicyReplaceResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyReplaceResponseExcludeAccessIPListRule struct {
	IPList AccessApplicationPolicyReplaceResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyReplaceResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessIPListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessIPListRule]
type accessApplicationPolicyReplaceResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessIPListRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessIPListRuleIPList]
type accessApplicationPolicyReplaceResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyReplaceResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                            `json:"certificate,required"`
	JSON        accessApplicationPolicyReplaceResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessCertificateRule]
type accessApplicationPolicyReplaceResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessCertificateRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRule]
type accessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                      `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRule]
type accessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                       `json:"connection_id,required"`
	JSON         accessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRule]
type accessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                          `json:"name,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRule]
type accessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                       `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRule]
type accessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRule]
type accessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                   `json:"attribute_value,required"`
	JSON           accessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRule]
type accessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                              `json:"token_id,required"`
	JSON    accessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                     `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRule]
type accessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                          `json:"keys_url,required"`
	JSON    accessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyReplaceResponseExcludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyReplaceResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyReplaceResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessCountryRule]
type accessApplicationPolicyReplaceResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessCountryRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                `json:"country_code,required"`
	JSON        accessApplicationPolicyReplaceResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessCountryRuleGeo]
type accessApplicationPolicyReplaceResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRule]
type accessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                    `json:"auth_method,required"`
	JSON       accessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRule]
type accessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyReplaceResponseExclude() {
}

type AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                `json:"integration_uid,required"`
	JSON           accessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessApplicationPolicyReplaceResponseIncludeAccessEmailRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessDomainRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessEveryoneRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessIPRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessIPListRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessCertificateRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessCountryRule],
// [AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyReplaceResponseInclude interface {
	implementsAccessApplicationPolicyReplaceResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyReplaceResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyReplaceResponseIncludeAccessEmailRule struct {
	Email AccessApplicationPolicyReplaceResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessEmailRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessEmailRule]
type accessApplicationPolicyReplaceResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessEmailRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyReplaceResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessEmailRuleEmail]
type accessApplicationPolicyReplaceResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRule]
type accessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                        `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleEmailList]
type accessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyReplaceResponseIncludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyReplaceResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyReplaceResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessDomainRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessDomainRule]
type accessApplicationPolicyReplaceResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessDomainRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                       `json:"domain,required"`
	JSON   accessApplicationPolicyReplaceResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyReplaceResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyReplaceResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                         `json:"everyone,required"`
	JSON     accessApplicationPolicyReplaceResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessEveryoneRule]
type accessApplicationPolicyReplaceResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessEveryoneRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyReplaceResponseIncludeAccessIPRule struct {
	IP   AccessApplicationPolicyReplaceResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessIPRule]
type accessApplicationPolicyReplaceResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessIPRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                          `json:"ip,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessIPRuleIP]
type accessApplicationPolicyReplaceResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyReplaceResponseIncludeAccessIPListRule struct {
	IPList AccessApplicationPolicyReplaceResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyReplaceResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessIPListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessIPListRule]
type accessApplicationPolicyReplaceResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessIPListRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessIPListRuleIPList]
type accessApplicationPolicyReplaceResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyReplaceResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                            `json:"certificate,required"`
	JSON        accessApplicationPolicyReplaceResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessCertificateRule]
type accessApplicationPolicyReplaceResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessCertificateRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRule]
type accessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                      `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRule]
type accessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                       `json:"connection_id,required"`
	JSON         accessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRule]
type accessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                          `json:"name,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRule]
type accessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                       `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRule]
type accessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRule]
type accessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                   `json:"attribute_value,required"`
	JSON           accessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRule]
type accessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                              `json:"token_id,required"`
	JSON    accessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                     `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRule]
type accessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                          `json:"keys_url,required"`
	JSON    accessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyReplaceResponseIncludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyReplaceResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyReplaceResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessCountryRule]
type accessApplicationPolicyReplaceResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessCountryRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                `json:"country_code,required"`
	JSON        accessApplicationPolicyReplaceResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessCountryRuleGeo]
type accessApplicationPolicyReplaceResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRule]
type accessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                    `json:"auth_method,required"`
	JSON       accessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRule]
type accessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyReplaceResponseInclude() {
}

type AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                `json:"integration_uid,required"`
	JSON           accessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessApplicationPolicyReplaceResponseRequireAccessEmailRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessEmailListRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessDomainRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessEveryoneRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessIPRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessIPListRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessCertificateRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessCountryRule],
// [AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyReplaceResponseRequire interface {
	implementsAccessApplicationPolicyReplaceResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyReplaceResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyReplaceResponseRequireAccessEmailRule struct {
	Email AccessApplicationPolicyReplaceResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessEmailRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessEmailRule]
type accessApplicationPolicyReplaceResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessEmailRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyReplaceResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessEmailRuleEmail]
type accessApplicationPolicyReplaceResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyReplaceResponseRequireAccessEmailListRule struct {
	EmailList AccessApplicationPolicyReplaceResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyReplaceResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessEmailListRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessEmailListRule]
type accessApplicationPolicyReplaceResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessEmailListRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                        `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessEmailListRuleEmailList]
type accessApplicationPolicyReplaceResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyReplaceResponseRequireAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyReplaceResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyReplaceResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessDomainRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessDomainRule]
type accessApplicationPolicyReplaceResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessDomainRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                       `json:"domain,required"`
	JSON   accessApplicationPolicyReplaceResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessDomainRuleEmailDomain]
type accessApplicationPolicyReplaceResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyReplaceResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                         `json:"everyone,required"`
	JSON     accessApplicationPolicyReplaceResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessEveryoneRule]
type accessApplicationPolicyReplaceResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessEveryoneRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyReplaceResponseRequireAccessIPRule struct {
	IP   AccessApplicationPolicyReplaceResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessIPRule]
type accessApplicationPolicyReplaceResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessIPRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                          `json:"ip,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessIPRuleIP]
type accessApplicationPolicyReplaceResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyReplaceResponseRequireAccessIPListRule struct {
	IPList AccessApplicationPolicyReplaceResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyReplaceResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessIPListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessIPListRule]
type accessApplicationPolicyReplaceResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessIPListRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                  `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessIPListRuleIPList]
type accessApplicationPolicyReplaceResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyReplaceResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                            `json:"certificate,required"`
	JSON        accessApplicationPolicyReplaceResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessCertificateRule]
type accessApplicationPolicyReplaceResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessCertificateRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRule struct {
	Group AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRule]
type accessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                      `json:"id,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleGroup]
type accessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRule]
type accessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                       `json:"connection_id,required"`
	JSON         accessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRule]
type accessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                          `json:"name,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRule]
type accessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                       `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRule]
type accessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                   `json:"email,required"`
	JSON  accessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleOkta]
type accessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRule]
type accessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                   `json:"attribute_value,required"`
	JSON           accessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleSaml]
type accessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRule]
type accessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                              `json:"token_id,required"`
	JSON    accessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                     `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRule]
type accessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRule]
type accessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                          `json:"keys_url,required"`
	JSON    accessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyReplaceResponseRequireAccessCountryRule struct {
	Geo  AccessApplicationPolicyReplaceResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyReplaceResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessCountryRule]
type accessApplicationPolicyReplaceResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessCountryRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                `json:"country_code,required"`
	JSON        accessApplicationPolicyReplaceResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessCountryRuleGeo]
type accessApplicationPolicyReplaceResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRule]
type accessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                    `json:"auth_method,required"`
	JSON       accessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRule]
type accessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRule) implementsAccessApplicationPolicyReplaceResponseRequire() {
}

type AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                `json:"integration_uid,required"`
	JSON           accessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyNewParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[AccessApplicationPolicyNewParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessApplicationPolicyNewParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]AccessApplicationPolicyNewParamsApprovalGroup] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessApplicationPolicyNewParamsExclude] `json:"exclude"`
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
	Require param.Field[[]AccessApplicationPolicyNewParamsRequire] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyNewParamsDecision string

const (
	AccessApplicationPolicyNewParamsDecisionAllow       AccessApplicationPolicyNewParamsDecision = "allow"
	AccessApplicationPolicyNewParamsDecisionDeny        AccessApplicationPolicyNewParamsDecision = "deny"
	AccessApplicationPolicyNewParamsDecisionNonIdentity AccessApplicationPolicyNewParamsDecision = "non_identity"
	AccessApplicationPolicyNewParamsDecisionBypass      AccessApplicationPolicyNewParamsDecision = "bypass"
)

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyNewParamsIncludeAccessEmailRule],
// [AccessApplicationPolicyNewParamsIncludeAccessEmailListRule],
// [AccessApplicationPolicyNewParamsIncludeAccessDomainRule],
// [AccessApplicationPolicyNewParamsIncludeAccessEveryoneRule],
// [AccessApplicationPolicyNewParamsIncludeAccessIPRule],
// [AccessApplicationPolicyNewParamsIncludeAccessIPListRule],
// [AccessApplicationPolicyNewParamsIncludeAccessCertificateRule],
// [AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyNewParamsIncludeAccessCountryRule],
// [AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule],
// [AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule].
type AccessApplicationPolicyNewParamsInclude interface {
	implementsAccessApplicationPolicyNewParamsInclude()
}

// Matches a specific email.
type AccessApplicationPolicyNewParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyNewParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyNewParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailListRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyNewParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyNewParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessDomainRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyNewParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEveryoneRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewParamsIncludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyNewParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyNewParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyNewParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPListRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyNewParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessCertificateRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyNewParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyNewParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessCountryRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyNewParamsInclude() {
}

type AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyNewParamsApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]interface{}] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid param.Field[string] `json:"email_list_uuid"`
}

func (r AccessApplicationPolicyNewParamsApprovalGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyNewParamsExcludeAccessEmailRule],
// [AccessApplicationPolicyNewParamsExcludeAccessEmailListRule],
// [AccessApplicationPolicyNewParamsExcludeAccessDomainRule],
// [AccessApplicationPolicyNewParamsExcludeAccessEveryoneRule],
// [AccessApplicationPolicyNewParamsExcludeAccessIPRule],
// [AccessApplicationPolicyNewParamsExcludeAccessIPListRule],
// [AccessApplicationPolicyNewParamsExcludeAccessCertificateRule],
// [AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyNewParamsExcludeAccessCountryRule],
// [AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule],
// [AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule].
type AccessApplicationPolicyNewParamsExclude interface {
	implementsAccessApplicationPolicyNewParamsExclude()
}

// Matches a specific email.
type AccessApplicationPolicyNewParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyNewParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyNewParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailListRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyNewParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyNewParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessDomainRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyNewParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEveryoneRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewParamsExcludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyNewParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyNewParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyNewParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPListRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyNewParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessCertificateRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyNewParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyNewParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessCountryRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyNewParamsExclude() {
}

type AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyNewParamsRequireAccessEmailRule],
// [AccessApplicationPolicyNewParamsRequireAccessEmailListRule],
// [AccessApplicationPolicyNewParamsRequireAccessDomainRule],
// [AccessApplicationPolicyNewParamsRequireAccessEveryoneRule],
// [AccessApplicationPolicyNewParamsRequireAccessIPRule],
// [AccessApplicationPolicyNewParamsRequireAccessIPListRule],
// [AccessApplicationPolicyNewParamsRequireAccessCertificateRule],
// [AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule],
// [AccessApplicationPolicyNewParamsRequireAccessAzureGroupRule],
// [AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyNewParamsRequireAccessOktaGroupRule],
// [AccessApplicationPolicyNewParamsRequireAccessSamlGroupRule],
// [AccessApplicationPolicyNewParamsRequireAccessServiceTokenRule],
// [AccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyNewParamsRequireAccessCountryRule],
// [AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule],
// [AccessApplicationPolicyNewParamsRequireAccessDevicePostureRule].
type AccessApplicationPolicyNewParamsRequire interface {
	implementsAccessApplicationPolicyNewParamsRequire()
}

// Matches a specific email.
type AccessApplicationPolicyNewParamsRequireAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyNewParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyNewParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailListRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyNewParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyNewParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessDomainRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyNewParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessEveryoneRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewParamsRequireAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyNewParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyNewParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyNewParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPListRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyNewParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessCertificateRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyNewParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyNewParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyNewParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessAzureGroupRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyNewParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyNewParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessOktaGroupRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyNewParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyNewParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessSamlGroupRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyNewParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyNewParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessServiceTokenRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyNewParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyNewParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessCountryRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyNewParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyNewParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessDevicePostureRule) implementsAccessApplicationPolicyNewParamsRequire() {
}

type AccessApplicationPolicyNewParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyNewResponseEnvelope struct {
	Errors   []AccessApplicationPolicyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationPolicyNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationPolicyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationPolicyNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyNewResponseEnvelope]
type accessApplicationPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessApplicationPolicyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyNewResponseEnvelopeErrors]
type accessApplicationPolicyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accessApplicationPolicyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyNewResponseEnvelopeMessages]
type accessApplicationPolicyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationPolicyNewResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyNewResponseEnvelopeSuccessTrue AccessApplicationPolicyNewResponseEnvelopeSuccess = true
)

type AccessApplicationPolicyListResponseEnvelope struct {
	Errors   []AccessApplicationPolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessApplicationPolicyListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessApplicationPolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessApplicationPolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessApplicationPolicyListResponseEnvelopeJSON       `json:"-"`
}

// accessApplicationPolicyListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyListResponseEnvelope]
type accessApplicationPolicyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessApplicationPolicyListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationPolicyListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyListResponseEnvelopeErrors]
type accessApplicationPolicyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accessApplicationPolicyListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationPolicyListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyListResponseEnvelopeMessages]
type accessApplicationPolicyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationPolicyListResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyListResponseEnvelopeSuccessTrue AccessApplicationPolicyListResponseEnvelopeSuccess = true
)

type AccessApplicationPolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       accessApplicationPolicyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessApplicationPolicyListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyListResponseEnvelopeResultInfo]
type accessApplicationPolicyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyDeleteResponseEnvelope struct {
	Errors   []AccessApplicationPolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationPolicyDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationPolicyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyDeleteResponseEnvelope]
type accessApplicationPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyDeleteResponseEnvelopeErrors]
type accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyDeleteResponseEnvelopeMessages]
type accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationPolicyDeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyDeleteResponseEnvelopeSuccessTrue AccessApplicationPolicyDeleteResponseEnvelopeSuccess = true
)

type AccessApplicationPolicyGetResponseEnvelope struct {
	Errors   []AccessApplicationPolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationPolicyGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationPolicyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationPolicyGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyGetResponseEnvelope]
type accessApplicationPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accessApplicationPolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyGetResponseEnvelopeErrors]
type accessApplicationPolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accessApplicationPolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyGetResponseEnvelopeMessages]
type accessApplicationPolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationPolicyGetResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyGetResponseEnvelopeSuccessTrue AccessApplicationPolicyGetResponseEnvelopeSuccess = true
)

type AccessApplicationPolicyReplaceParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[AccessApplicationPolicyReplaceParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessApplicationPolicyReplaceParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]AccessApplicationPolicyReplaceParamsApprovalGroup] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessApplicationPolicyReplaceParamsExclude] `json:"exclude"`
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
	Require param.Field[[]AccessApplicationPolicyReplaceParamsRequire] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyReplaceParamsDecision string

const (
	AccessApplicationPolicyReplaceParamsDecisionAllow       AccessApplicationPolicyReplaceParamsDecision = "allow"
	AccessApplicationPolicyReplaceParamsDecisionDeny        AccessApplicationPolicyReplaceParamsDecision = "deny"
	AccessApplicationPolicyReplaceParamsDecisionNonIdentity AccessApplicationPolicyReplaceParamsDecision = "non_identity"
	AccessApplicationPolicyReplaceParamsDecisionBypass      AccessApplicationPolicyReplaceParamsDecision = "bypass"
)

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessEmailListRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessDomainRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessEveryoneRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessIPRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessIPListRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessCertificateRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessCountryRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessAuthenticationMethodRule],
// [AccessApplicationPolicyReplaceParamsIncludeAccessDevicePostureRule].
type AccessApplicationPolicyReplaceParamsInclude interface {
	implementsAccessApplicationPolicyReplaceParamsInclude()
}

// Matches a specific email.
type AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEmailRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyReplaceParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEmailListRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyReplaceParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessDomainRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyReplaceParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessEveryoneRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyReplaceParamsIncludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessIPRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyReplaceParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessIPListRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyReplaceParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessCertificateRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyReplaceParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyReplaceParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyReplaceParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyReplaceParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyReplaceParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyReplaceParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyReplaceParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyReplaceParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyReplaceParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyReplaceParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessCountryRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyReplaceParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyReplaceParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyReplaceParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyReplaceParamsInclude() {
}

type AccessApplicationPolicyReplaceParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyReplaceParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyReplaceParamsApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]interface{}] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid param.Field[string] `json:"email_list_uuid"`
}

func (r AccessApplicationPolicyReplaceParamsApprovalGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessEmailListRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessDomainRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessEveryoneRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessIPRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessIPListRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessCertificateRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessCountryRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessAuthenticationMethodRule],
// [AccessApplicationPolicyReplaceParamsExcludeAccessDevicePostureRule].
type AccessApplicationPolicyReplaceParamsExclude interface {
	implementsAccessApplicationPolicyReplaceParamsExclude()
}

// Matches a specific email.
type AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEmailRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyReplaceParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEmailListRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyReplaceParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessDomainRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyReplaceParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessEveryoneRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyReplaceParamsExcludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessIPRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyReplaceParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessIPListRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyReplaceParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessCertificateRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyReplaceParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyReplaceParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyReplaceParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyReplaceParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyReplaceParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyReplaceParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyReplaceParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyReplaceParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyReplaceParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyReplaceParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessCountryRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyReplaceParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyReplaceParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyReplaceParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyReplaceParamsExclude() {
}

type AccessApplicationPolicyReplaceParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyReplaceParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyReplaceParamsRequireAccessEmailRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessEmailListRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessDomainRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessEveryoneRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessIPRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessIPListRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessCertificateRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessAccessGroupRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessAzureGroupRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessOktaGroupRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessSamlGroupRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessServiceTokenRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessCountryRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessAuthenticationMethodRule],
// [AccessApplicationPolicyReplaceParamsRequireAccessDevicePostureRule].
type AccessApplicationPolicyReplaceParamsRequire interface {
	implementsAccessApplicationPolicyReplaceParamsRequire()
}

// Matches a specific email.
type AccessApplicationPolicyReplaceParamsRequireAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyReplaceParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEmailRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyReplaceParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyReplaceParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEmailListRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyReplaceParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyReplaceParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessDomainRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyReplaceParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessEveryoneRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyReplaceParamsRequireAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyReplaceParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessIPRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyReplaceParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyReplaceParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessIPListRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyReplaceParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessCertificateRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyReplaceParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyReplaceParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAccessGroupRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyReplaceParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyReplaceParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAzureGroupRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyReplaceParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyReplaceParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyReplaceParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyReplaceParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyReplaceParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyReplaceParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessOktaGroupRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyReplaceParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyReplaceParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessSamlGroupRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyReplaceParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyReplaceParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessServiceTokenRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyReplaceParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyReplaceParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyReplaceParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyReplaceParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyReplaceParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessCountryRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyReplaceParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyReplaceParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyReplaceParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyReplaceParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessDevicePostureRule) implementsAccessApplicationPolicyReplaceParamsRequire() {
}

type AccessApplicationPolicyReplaceParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyReplaceParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyReplaceResponseEnvelope struct {
	Errors   []AccessApplicationPolicyReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationPolicyReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationPolicyReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationPolicyReplaceResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyReplaceResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyReplaceResponseEnvelope]
type accessApplicationPolicyReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyReplaceResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessApplicationPolicyReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyReplaceResponseEnvelopeErrors]
type accessApplicationPolicyReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyReplaceResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accessApplicationPolicyReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationPolicyReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyReplaceResponseEnvelopeMessages]
type accessApplicationPolicyReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationPolicyReplaceResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyReplaceResponseEnvelopeSuccessTrue AccessApplicationPolicyReplaceResponseEnvelopeSuccess = true
)
