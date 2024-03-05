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

// ZeroTrustAccessApplicationPolicyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessApplicationPolicyService] method instead.
type ZeroTrustAccessApplicationPolicyService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessApplicationPolicyService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessApplicationPolicyService(opts ...option.RequestOption) (r *ZeroTrustAccessApplicationPolicyService) {
	r = &ZeroTrustAccessApplicationPolicyService{}
	r.Options = opts
	return
}

// Create a new Access policy for an application.
func (r *ZeroTrustAccessApplicationPolicyService) New(ctx context.Context, uuid string, params ZeroTrustAccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *AccessPolicies, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationPolicyNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured Access policy.
func (r *ZeroTrustAccessApplicationPolicyService) Update(ctx context.Context, uuid1 string, uuid string, params ZeroTrustAccessApplicationPolicyUpdateParams, opts ...option.RequestOption) (res *AccessPolicies, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationPolicyUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Access policies configured for an application.
func (r *ZeroTrustAccessApplicationPolicyService) List(ctx context.Context, uuid string, query ZeroTrustAccessApplicationPolicyListParams, opts ...option.RequestOption) (res *[]AccessPolicies, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationPolicyListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an Access policy.
func (r *ZeroTrustAccessApplicationPolicyService) Delete(ctx context.Context, uuid1 string, uuid string, body ZeroTrustAccessApplicationPolicyDeleteParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationPolicyDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access policy.
func (r *ZeroTrustAccessApplicationPolicyService) Get(ctx context.Context, uuid1 string, uuid string, query ZeroTrustAccessApplicationPolicyGetParams, opts ...option.RequestOption) (res *AccessPolicies, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationPolicyGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, uuid1, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessPolicies struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessPoliciesApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessPoliciesDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessPoliciesExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessPoliciesInclude `json:"include"`
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
	Require []AccessPoliciesRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string             `json:"session_duration"`
	UpdatedAt       time.Time          `json:"updated_at" format:"date-time"`
	JSON            accessPoliciesJSON `json:"-"`
}

// accessPoliciesJSON contains the JSON metadata for the struct [AccessPolicies]
type accessPoliciesJSON struct {
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

func (r *AccessPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessPoliciesApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                          `json:"email_list_uuid"`
	JSON          accessPoliciesApprovalGroupJSON `json:"-"`
}

// accessPoliciesApprovalGroupJSON contains the JSON metadata for the struct
// [AccessPoliciesApprovalGroup]
type accessPoliciesApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessPoliciesApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessPoliciesDecision string

const (
	AccessPoliciesDecisionAllow       AccessPoliciesDecision = "allow"
	AccessPoliciesDecisionDeny        AccessPoliciesDecision = "deny"
	AccessPoliciesDecisionNonIdentity AccessPoliciesDecision = "non_identity"
	AccessPoliciesDecisionBypass      AccessPoliciesDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by [AccessPoliciesExcludeAccessEmailRule],
// [AccessPoliciesExcludeAccessEmailListRule],
// [AccessPoliciesExcludeAccessDomainRule],
// [AccessPoliciesExcludeAccessEveryoneRule], [AccessPoliciesExcludeAccessIPRule],
// [AccessPoliciesExcludeAccessIPListRule],
// [AccessPoliciesExcludeAccessCertificateRule],
// [AccessPoliciesExcludeAccessAccessGroupRule],
// [AccessPoliciesExcludeAccessAzureGroupRule],
// [AccessPoliciesExcludeAccessGitHubOrganizationRule],
// [AccessPoliciesExcludeAccessGsuiteGroupRule],
// [AccessPoliciesExcludeAccessOktaGroupRule],
// [AccessPoliciesExcludeAccessSamlGroupRule],
// [AccessPoliciesExcludeAccessServiceTokenRule],
// [AccessPoliciesExcludeAccessAnyValidServiceTokenRule],
// [AccessPoliciesExcludeAccessExternalEvaluationRule],
// [AccessPoliciesExcludeAccessCountryRule],
// [AccessPoliciesExcludeAccessAuthenticationMethodRule] or
// [AccessPoliciesExcludeAccessDevicePostureRule].
type AccessPoliciesExclude interface {
	implementsAccessPoliciesExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessPoliciesExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessPoliciesExcludeAccessEmailRule struct {
	Email AccessPoliciesExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessPoliciesExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessPoliciesExcludeAccessEmailRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessEmailRule]
type accessPoliciesExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessEmailRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                        `json:"email,required" format:"email"`
	JSON  accessPoliciesExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessPoliciesExcludeAccessEmailRuleEmailJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessEmailRuleEmail]
type accessPoliciesExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessPoliciesExcludeAccessEmailListRule struct {
	EmailList AccessPoliciesExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessPoliciesExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessPoliciesExcludeAccessEmailListRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessEmailListRule]
type accessPoliciesExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessEmailListRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                `json:"id,required"`
	JSON accessPoliciesExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessPoliciesExcludeAccessEmailListRuleEmailListJSON contains the JSON metadata
// for the struct [AccessPoliciesExcludeAccessEmailListRuleEmailList]
type accessPoliciesExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessPoliciesExcludeAccessDomainRule struct {
	EmailDomain AccessPoliciesExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessPoliciesExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessPoliciesExcludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessDomainRule]
type accessPoliciesExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessDomainRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                               `json:"domain,required"`
	JSON   accessPoliciesExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessPoliciesExcludeAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [AccessPoliciesExcludeAccessDomainRuleEmailDomain]
type accessPoliciesExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessPoliciesExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                 `json:"everyone,required"`
	JSON     accessPoliciesExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessPoliciesExcludeAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessEveryoneRule]
type accessPoliciesExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessEveryoneRule) implementsAccessPoliciesExclude() {}

// Matches an IP address block.
type AccessPoliciesExcludeAccessIPRule struct {
	IP   AccessPoliciesExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessPoliciesExcludeAccessIPRuleJSON `json:"-"`
}

// accessPoliciesExcludeAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessPoliciesExcludeAccessIPRule]
type accessPoliciesExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessIPRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                  `json:"ip,required"`
	JSON accessPoliciesExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessPoliciesExcludeAccessIPRuleIPJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessIPRuleIP]
type accessPoliciesExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessPoliciesExcludeAccessIPListRule struct {
	IPList AccessPoliciesExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessPoliciesExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessPoliciesExcludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessIPListRule]
type accessPoliciesExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessIPListRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                          `json:"id,required"`
	JSON accessPoliciesExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessPoliciesExcludeAccessIPListRuleIPListJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessIPListRuleIPList]
type accessPoliciesExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessPoliciesExcludeAccessCertificateRule struct {
	Certificate interface{}                                    `json:"certificate,required"`
	JSON        accessPoliciesExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessPoliciesExcludeAccessCertificateRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessCertificateRule]
type accessPoliciesExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessCertificateRule) implementsAccessPoliciesExclude() {}

// Matches an Access group.
type AccessPoliciesExcludeAccessAccessGroupRule struct {
	Group AccessPoliciesExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessPoliciesExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessPoliciesExcludeAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessAccessGroupRule]
type accessPoliciesExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessAccessGroupRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                              `json:"id,required"`
	JSON accessPoliciesExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessPoliciesExcludeAccessAccessGroupRuleGroupJSON contains the JSON metadata
// for the struct [AccessPoliciesExcludeAccessAccessGroupRuleGroup]
type accessPoliciesExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessPoliciesExcludeAccessAzureGroupRule struct {
	AzureAd AccessPoliciesExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessPoliciesExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessPoliciesExcludeAccessAzureGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessAzureGroupRule]
type accessPoliciesExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessAzureGroupRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                               `json:"connection_id,required"`
	JSON         accessPoliciesExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessPoliciesExcludeAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [AccessPoliciesExcludeAccessAzureGroupRuleAzureAd]
type accessPoliciesExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessPoliciesExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessPoliciesExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessPoliciesExcludeAccessGitHubOrganizationRuleJSON contains the JSON metadata
// for the struct [AccessPoliciesExcludeAccessGitHubOrganizationRule]
type accessPoliciesExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessGitHubOrganizationRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                  `json:"name,required"`
	JSON accessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON contains
// the JSON metadata for the struct
// [AccessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessPoliciesExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessPoliciesExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessPoliciesExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessPoliciesExcludeAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessGsuiteGroupRule]
type accessPoliciesExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessGsuiteGroupRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                               `json:"email,required"`
	JSON  accessPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [AccessPoliciesExcludeAccessGsuiteGroupRuleGsuite]
type accessPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessPoliciesExcludeAccessOktaGroupRule struct {
	Okta AccessPoliciesExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessPoliciesExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessPoliciesExcludeAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessOktaGroupRule]
type accessPoliciesExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessOktaGroupRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                           `json:"email,required"`
	JSON  accessPoliciesExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessPoliciesExcludeAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessOktaGroupRuleOkta]
type accessPoliciesExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessPoliciesExcludeAccessSamlGroupRule struct {
	Saml AccessPoliciesExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessPoliciesExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessPoliciesExcludeAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessSamlGroupRule]
type accessPoliciesExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessSamlGroupRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                           `json:"attribute_value,required"`
	JSON           accessPoliciesExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessPoliciesExcludeAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessSamlGroupRuleSaml]
type accessPoliciesExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessPoliciesExcludeAccessServiceTokenRule struct {
	ServiceToken AccessPoliciesExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessPoliciesExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessPoliciesExcludeAccessServiceTokenRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessServiceTokenRule]
type accessPoliciesExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessServiceTokenRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                      `json:"token_id,required"`
	JSON    accessPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [AccessPoliciesExcludeAccessServiceTokenRuleServiceToken]
type accessPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessPoliciesExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                             `json:"any_valid_service_token,required"`
	JSON                 accessPoliciesExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessPoliciesExcludeAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessPoliciesExcludeAccessAnyValidServiceTokenRule]
type accessPoliciesExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessAnyValidServiceTokenRule) implementsAccessPoliciesExclude() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessPoliciesExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessPoliciesExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessPoliciesExcludeAccessExternalEvaluationRuleJSON contains the JSON metadata
// for the struct [AccessPoliciesExcludeAccessExternalEvaluationRule]
type accessPoliciesExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessExternalEvaluationRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                  `json:"keys_url,required"`
	JSON    accessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON contains
// the JSON metadata for the struct
// [AccessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessPoliciesExcludeAccessCountryRule struct {
	Geo  AccessPoliciesExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessPoliciesExcludeAccessCountryRuleJSON `json:"-"`
}

// accessPoliciesExcludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessCountryRule]
type accessPoliciesExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessCountryRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                        `json:"country_code,required"`
	JSON        accessPoliciesExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessPoliciesExcludeAccessCountryRuleGeoJSON contains the JSON metadata for the
// struct [AccessPoliciesExcludeAccessCountryRuleGeo]
type accessPoliciesExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessPoliciesExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessPoliciesExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessPoliciesExcludeAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [AccessPoliciesExcludeAccessAuthenticationMethodRule]
type accessPoliciesExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessAuthenticationMethodRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                            `json:"auth_method,required"`
	JSON       accessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [AccessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessPoliciesExcludeAccessDevicePostureRule struct {
	DevicePosture AccessPoliciesExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessPoliciesExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessPoliciesExcludeAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesExcludeAccessDevicePostureRule]
type accessPoliciesExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesExcludeAccessDevicePostureRule) implementsAccessPoliciesExclude() {}

type AccessPoliciesExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                        `json:"integration_uid,required"`
	JSON           accessPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [AccessPoliciesExcludeAccessDevicePostureRuleDevicePosture]
type accessPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessPoliciesExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessPoliciesIncludeAccessEmailRule],
// [AccessPoliciesIncludeAccessEmailListRule],
// [AccessPoliciesIncludeAccessDomainRule],
// [AccessPoliciesIncludeAccessEveryoneRule], [AccessPoliciesIncludeAccessIPRule],
// [AccessPoliciesIncludeAccessIPListRule],
// [AccessPoliciesIncludeAccessCertificateRule],
// [AccessPoliciesIncludeAccessAccessGroupRule],
// [AccessPoliciesIncludeAccessAzureGroupRule],
// [AccessPoliciesIncludeAccessGitHubOrganizationRule],
// [AccessPoliciesIncludeAccessGsuiteGroupRule],
// [AccessPoliciesIncludeAccessOktaGroupRule],
// [AccessPoliciesIncludeAccessSamlGroupRule],
// [AccessPoliciesIncludeAccessServiceTokenRule],
// [AccessPoliciesIncludeAccessAnyValidServiceTokenRule],
// [AccessPoliciesIncludeAccessExternalEvaluationRule],
// [AccessPoliciesIncludeAccessCountryRule],
// [AccessPoliciesIncludeAccessAuthenticationMethodRule] or
// [AccessPoliciesIncludeAccessDevicePostureRule].
type AccessPoliciesInclude interface {
	implementsAccessPoliciesInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessPoliciesInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessPoliciesIncludeAccessEmailRule struct {
	Email AccessPoliciesIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessPoliciesIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessPoliciesIncludeAccessEmailRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessEmailRule]
type accessPoliciesIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessEmailRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                        `json:"email,required" format:"email"`
	JSON  accessPoliciesIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessPoliciesIncludeAccessEmailRuleEmailJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessEmailRuleEmail]
type accessPoliciesIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessPoliciesIncludeAccessEmailListRule struct {
	EmailList AccessPoliciesIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessPoliciesIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessPoliciesIncludeAccessEmailListRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessEmailListRule]
type accessPoliciesIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessEmailListRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                `json:"id,required"`
	JSON accessPoliciesIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessPoliciesIncludeAccessEmailListRuleEmailListJSON contains the JSON metadata
// for the struct [AccessPoliciesIncludeAccessEmailListRuleEmailList]
type accessPoliciesIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessPoliciesIncludeAccessDomainRule struct {
	EmailDomain AccessPoliciesIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessPoliciesIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessPoliciesIncludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessDomainRule]
type accessPoliciesIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessDomainRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                               `json:"domain,required"`
	JSON   accessPoliciesIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessPoliciesIncludeAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [AccessPoliciesIncludeAccessDomainRuleEmailDomain]
type accessPoliciesIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessPoliciesIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                 `json:"everyone,required"`
	JSON     accessPoliciesIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessPoliciesIncludeAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessEveryoneRule]
type accessPoliciesIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessEveryoneRule) implementsAccessPoliciesInclude() {}

// Matches an IP address block.
type AccessPoliciesIncludeAccessIPRule struct {
	IP   AccessPoliciesIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessPoliciesIncludeAccessIPRuleJSON `json:"-"`
}

// accessPoliciesIncludeAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessPoliciesIncludeAccessIPRule]
type accessPoliciesIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessIPRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                  `json:"ip,required"`
	JSON accessPoliciesIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessPoliciesIncludeAccessIPRuleIPJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessIPRuleIP]
type accessPoliciesIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessPoliciesIncludeAccessIPListRule struct {
	IPList AccessPoliciesIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessPoliciesIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessPoliciesIncludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessIPListRule]
type accessPoliciesIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessIPListRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                          `json:"id,required"`
	JSON accessPoliciesIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessPoliciesIncludeAccessIPListRuleIPListJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessIPListRuleIPList]
type accessPoliciesIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessPoliciesIncludeAccessCertificateRule struct {
	Certificate interface{}                                    `json:"certificate,required"`
	JSON        accessPoliciesIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessPoliciesIncludeAccessCertificateRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessCertificateRule]
type accessPoliciesIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessCertificateRule) implementsAccessPoliciesInclude() {}

// Matches an Access group.
type AccessPoliciesIncludeAccessAccessGroupRule struct {
	Group AccessPoliciesIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessPoliciesIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessPoliciesIncludeAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessAccessGroupRule]
type accessPoliciesIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessAccessGroupRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                              `json:"id,required"`
	JSON accessPoliciesIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessPoliciesIncludeAccessAccessGroupRuleGroupJSON contains the JSON metadata
// for the struct [AccessPoliciesIncludeAccessAccessGroupRuleGroup]
type accessPoliciesIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessPoliciesIncludeAccessAzureGroupRule struct {
	AzureAd AccessPoliciesIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessPoliciesIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessPoliciesIncludeAccessAzureGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessAzureGroupRule]
type accessPoliciesIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessAzureGroupRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                               `json:"connection_id,required"`
	JSON         accessPoliciesIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessPoliciesIncludeAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [AccessPoliciesIncludeAccessAzureGroupRuleAzureAd]
type accessPoliciesIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessPoliciesIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessPoliciesIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessPoliciesIncludeAccessGitHubOrganizationRuleJSON contains the JSON metadata
// for the struct [AccessPoliciesIncludeAccessGitHubOrganizationRule]
type accessPoliciesIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessGitHubOrganizationRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                  `json:"name,required"`
	JSON accessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON contains
// the JSON metadata for the struct
// [AccessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessPoliciesIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessPoliciesIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessPoliciesIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessPoliciesIncludeAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessGsuiteGroupRule]
type accessPoliciesIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessGsuiteGroupRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                               `json:"email,required"`
	JSON  accessPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [AccessPoliciesIncludeAccessGsuiteGroupRuleGsuite]
type accessPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessPoliciesIncludeAccessOktaGroupRule struct {
	Okta AccessPoliciesIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessPoliciesIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessPoliciesIncludeAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessOktaGroupRule]
type accessPoliciesIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessOktaGroupRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                           `json:"email,required"`
	JSON  accessPoliciesIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessPoliciesIncludeAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessOktaGroupRuleOkta]
type accessPoliciesIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessPoliciesIncludeAccessSamlGroupRule struct {
	Saml AccessPoliciesIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessPoliciesIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessPoliciesIncludeAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessSamlGroupRule]
type accessPoliciesIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessSamlGroupRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                           `json:"attribute_value,required"`
	JSON           accessPoliciesIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessPoliciesIncludeAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessSamlGroupRuleSaml]
type accessPoliciesIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessPoliciesIncludeAccessServiceTokenRule struct {
	ServiceToken AccessPoliciesIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessPoliciesIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessPoliciesIncludeAccessServiceTokenRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessServiceTokenRule]
type accessPoliciesIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessServiceTokenRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                      `json:"token_id,required"`
	JSON    accessPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [AccessPoliciesIncludeAccessServiceTokenRuleServiceToken]
type accessPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessPoliciesIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                             `json:"any_valid_service_token,required"`
	JSON                 accessPoliciesIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessPoliciesIncludeAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessPoliciesIncludeAccessAnyValidServiceTokenRule]
type accessPoliciesIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessAnyValidServiceTokenRule) implementsAccessPoliciesInclude() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessPoliciesIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessPoliciesIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessPoliciesIncludeAccessExternalEvaluationRuleJSON contains the JSON metadata
// for the struct [AccessPoliciesIncludeAccessExternalEvaluationRule]
type accessPoliciesIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessExternalEvaluationRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                  `json:"keys_url,required"`
	JSON    accessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON contains
// the JSON metadata for the struct
// [AccessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessPoliciesIncludeAccessCountryRule struct {
	Geo  AccessPoliciesIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessPoliciesIncludeAccessCountryRuleJSON `json:"-"`
}

// accessPoliciesIncludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessCountryRule]
type accessPoliciesIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessCountryRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                        `json:"country_code,required"`
	JSON        accessPoliciesIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessPoliciesIncludeAccessCountryRuleGeoJSON contains the JSON metadata for the
// struct [AccessPoliciesIncludeAccessCountryRuleGeo]
type accessPoliciesIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessPoliciesIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessPoliciesIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessPoliciesIncludeAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [AccessPoliciesIncludeAccessAuthenticationMethodRule]
type accessPoliciesIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessAuthenticationMethodRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                            `json:"auth_method,required"`
	JSON       accessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [AccessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessPoliciesIncludeAccessDevicePostureRule struct {
	DevicePosture AccessPoliciesIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessPoliciesIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessPoliciesIncludeAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesIncludeAccessDevicePostureRule]
type accessPoliciesIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesIncludeAccessDevicePostureRule) implementsAccessPoliciesInclude() {}

type AccessPoliciesIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                        `json:"integration_uid,required"`
	JSON           accessPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [AccessPoliciesIncludeAccessDevicePostureRuleDevicePosture]
type accessPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessPoliciesIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessPoliciesRequireAccessEmailRule],
// [AccessPoliciesRequireAccessEmailListRule],
// [AccessPoliciesRequireAccessDomainRule],
// [AccessPoliciesRequireAccessEveryoneRule], [AccessPoliciesRequireAccessIPRule],
// [AccessPoliciesRequireAccessIPListRule],
// [AccessPoliciesRequireAccessCertificateRule],
// [AccessPoliciesRequireAccessAccessGroupRule],
// [AccessPoliciesRequireAccessAzureGroupRule],
// [AccessPoliciesRequireAccessGitHubOrganizationRule],
// [AccessPoliciesRequireAccessGsuiteGroupRule],
// [AccessPoliciesRequireAccessOktaGroupRule],
// [AccessPoliciesRequireAccessSamlGroupRule],
// [AccessPoliciesRequireAccessServiceTokenRule],
// [AccessPoliciesRequireAccessAnyValidServiceTokenRule],
// [AccessPoliciesRequireAccessExternalEvaluationRule],
// [AccessPoliciesRequireAccessCountryRule],
// [AccessPoliciesRequireAccessAuthenticationMethodRule] or
// [AccessPoliciesRequireAccessDevicePostureRule].
type AccessPoliciesRequire interface {
	implementsAccessPoliciesRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessPoliciesRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessPoliciesRequireAccessEmailRule struct {
	Email AccessPoliciesRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessPoliciesRequireAccessEmailRuleJSON  `json:"-"`
}

// accessPoliciesRequireAccessEmailRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessEmailRule]
type accessPoliciesRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessEmailRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                        `json:"email,required" format:"email"`
	JSON  accessPoliciesRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessPoliciesRequireAccessEmailRuleEmailJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessEmailRuleEmail]
type accessPoliciesRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessPoliciesRequireAccessEmailListRule struct {
	EmailList AccessPoliciesRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessPoliciesRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessPoliciesRequireAccessEmailListRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessEmailListRule]
type accessPoliciesRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessEmailListRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                `json:"id,required"`
	JSON accessPoliciesRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessPoliciesRequireAccessEmailListRuleEmailListJSON contains the JSON metadata
// for the struct [AccessPoliciesRequireAccessEmailListRuleEmailList]
type accessPoliciesRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessPoliciesRequireAccessDomainRule struct {
	EmailDomain AccessPoliciesRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessPoliciesRequireAccessDomainRuleJSON        `json:"-"`
}

// accessPoliciesRequireAccessDomainRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessDomainRule]
type accessPoliciesRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessDomainRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                               `json:"domain,required"`
	JSON   accessPoliciesRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessPoliciesRequireAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [AccessPoliciesRequireAccessDomainRuleEmailDomain]
type accessPoliciesRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessPoliciesRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                 `json:"everyone,required"`
	JSON     accessPoliciesRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessPoliciesRequireAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessEveryoneRule]
type accessPoliciesRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessEveryoneRule) implementsAccessPoliciesRequire() {}

// Matches an IP address block.
type AccessPoliciesRequireAccessIPRule struct {
	IP   AccessPoliciesRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessPoliciesRequireAccessIPRuleJSON `json:"-"`
}

// accessPoliciesRequireAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessPoliciesRequireAccessIPRule]
type accessPoliciesRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessIPRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                  `json:"ip,required"`
	JSON accessPoliciesRequireAccessIPRuleIPJSON `json:"-"`
}

// accessPoliciesRequireAccessIPRuleIPJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessIPRuleIP]
type accessPoliciesRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessPoliciesRequireAccessIPListRule struct {
	IPList AccessPoliciesRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessPoliciesRequireAccessIPListRuleJSON   `json:"-"`
}

// accessPoliciesRequireAccessIPListRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessIPListRule]
type accessPoliciesRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessIPListRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                          `json:"id,required"`
	JSON accessPoliciesRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessPoliciesRequireAccessIPListRuleIPListJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessIPListRuleIPList]
type accessPoliciesRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessPoliciesRequireAccessCertificateRule struct {
	Certificate interface{}                                    `json:"certificate,required"`
	JSON        accessPoliciesRequireAccessCertificateRuleJSON `json:"-"`
}

// accessPoliciesRequireAccessCertificateRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessCertificateRule]
type accessPoliciesRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessCertificateRule) implementsAccessPoliciesRequire() {}

// Matches an Access group.
type AccessPoliciesRequireAccessAccessGroupRule struct {
	Group AccessPoliciesRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessPoliciesRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessPoliciesRequireAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessAccessGroupRule]
type accessPoliciesRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessAccessGroupRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                              `json:"id,required"`
	JSON accessPoliciesRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessPoliciesRequireAccessAccessGroupRuleGroupJSON contains the JSON metadata
// for the struct [AccessPoliciesRequireAccessAccessGroupRuleGroup]
type accessPoliciesRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessPoliciesRequireAccessAzureGroupRule struct {
	AzureAd AccessPoliciesRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessPoliciesRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessPoliciesRequireAccessAzureGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessAzureGroupRule]
type accessPoliciesRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessAzureGroupRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                               `json:"connection_id,required"`
	JSON         accessPoliciesRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessPoliciesRequireAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [AccessPoliciesRequireAccessAzureGroupRuleAzureAd]
type accessPoliciesRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessPoliciesRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessPoliciesRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessPoliciesRequireAccessGitHubOrganizationRuleJSON contains the JSON metadata
// for the struct [AccessPoliciesRequireAccessGitHubOrganizationRule]
type accessPoliciesRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessGitHubOrganizationRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                  `json:"name,required"`
	JSON accessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON contains
// the JSON metadata for the struct
// [AccessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessPoliciesRequireAccessGsuiteGroupRule struct {
	Gsuite AccessPoliciesRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessPoliciesRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessPoliciesRequireAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessGsuiteGroupRule]
type accessPoliciesRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessGsuiteGroupRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                               `json:"email,required"`
	JSON  accessPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [AccessPoliciesRequireAccessGsuiteGroupRuleGsuite]
type accessPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessPoliciesRequireAccessOktaGroupRule struct {
	Okta AccessPoliciesRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessPoliciesRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessPoliciesRequireAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessOktaGroupRule]
type accessPoliciesRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessOktaGroupRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                           `json:"email,required"`
	JSON  accessPoliciesRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessPoliciesRequireAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessOktaGroupRuleOkta]
type accessPoliciesRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessPoliciesRequireAccessSamlGroupRule struct {
	Saml AccessPoliciesRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessPoliciesRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessPoliciesRequireAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessSamlGroupRule]
type accessPoliciesRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessSamlGroupRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                           `json:"attribute_value,required"`
	JSON           accessPoliciesRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessPoliciesRequireAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessSamlGroupRuleSaml]
type accessPoliciesRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessPoliciesRequireAccessServiceTokenRule struct {
	ServiceToken AccessPoliciesRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessPoliciesRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessPoliciesRequireAccessServiceTokenRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessServiceTokenRule]
type accessPoliciesRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessServiceTokenRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                      `json:"token_id,required"`
	JSON    accessPoliciesRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessPoliciesRequireAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [AccessPoliciesRequireAccessServiceTokenRuleServiceToken]
type accessPoliciesRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessPoliciesRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                             `json:"any_valid_service_token,required"`
	JSON                 accessPoliciesRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessPoliciesRequireAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessPoliciesRequireAccessAnyValidServiceTokenRule]
type accessPoliciesRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessAnyValidServiceTokenRule) implementsAccessPoliciesRequire() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessPoliciesRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessPoliciesRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessPoliciesRequireAccessExternalEvaluationRuleJSON contains the JSON metadata
// for the struct [AccessPoliciesRequireAccessExternalEvaluationRule]
type accessPoliciesRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessExternalEvaluationRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                  `json:"keys_url,required"`
	JSON    accessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON contains
// the JSON metadata for the struct
// [AccessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessPoliciesRequireAccessCountryRule struct {
	Geo  AccessPoliciesRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessPoliciesRequireAccessCountryRuleJSON `json:"-"`
}

// accessPoliciesRequireAccessCountryRuleJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessCountryRule]
type accessPoliciesRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessCountryRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                        `json:"country_code,required"`
	JSON        accessPoliciesRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessPoliciesRequireAccessCountryRuleGeoJSON contains the JSON metadata for the
// struct [AccessPoliciesRequireAccessCountryRuleGeo]
type accessPoliciesRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessPoliciesRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessPoliciesRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessPoliciesRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessPoliciesRequireAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [AccessPoliciesRequireAccessAuthenticationMethodRule]
type accessPoliciesRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessAuthenticationMethodRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                            `json:"auth_method,required"`
	JSON       accessPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [AccessPoliciesRequireAccessAuthenticationMethodRuleAuthMethod]
type accessPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessPoliciesRequireAccessDevicePostureRule struct {
	DevicePosture AccessPoliciesRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessPoliciesRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessPoliciesRequireAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [AccessPoliciesRequireAccessDevicePostureRule]
type accessPoliciesRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessPoliciesRequireAccessDevicePostureRule) implementsAccessPoliciesRequire() {}

type AccessPoliciesRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                        `json:"integration_uid,required"`
	JSON           accessPoliciesRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessPoliciesRequireAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [AccessPoliciesRequireAccessDevicePostureRuleDevicePosture]
type accessPoliciesRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessPoliciesRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyDeleteResponse struct {
	// UUID
	ID   string                                             `json:"id"`
	JSON zeroTrustAccessApplicationPolicyDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyDeleteResponseJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationPolicyDeleteResponse]
type zeroTrustAccessApplicationPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyNewParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[ZeroTrustAccessApplicationPolicyNewParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]ZeroTrustAccessApplicationPolicyNewParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ZeroTrustAccessApplicationPolicyNewParamsApprovalGroup] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ZeroTrustAccessApplicationPolicyNewParamsExclude] `json:"exclude"`
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
	Require param.Field[[]ZeroTrustAccessApplicationPolicyNewParamsRequire] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r ZeroTrustAccessApplicationPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type ZeroTrustAccessApplicationPolicyNewParamsDecision string

const (
	ZeroTrustAccessApplicationPolicyNewParamsDecisionAllow       ZeroTrustAccessApplicationPolicyNewParamsDecision = "allow"
	ZeroTrustAccessApplicationPolicyNewParamsDecisionDeny        ZeroTrustAccessApplicationPolicyNewParamsDecision = "deny"
	ZeroTrustAccessApplicationPolicyNewParamsDecisionNonIdentity ZeroTrustAccessApplicationPolicyNewParamsDecision = "non_identity"
	ZeroTrustAccessApplicationPolicyNewParamsDecisionBypass      ZeroTrustAccessApplicationPolicyNewParamsDecision = "bypass"
)

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyNewParamsInclude interface {
	implementsZeroTrustAccessApplicationPolicyNewParamsInclude()
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of email addresses that can approve a temporary authentication request.
type ZeroTrustAccessApplicationPolicyNewParamsApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]interface{}] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID param.Field[string] `json:"email_list_uuid"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsApprovalGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyNewParamsExclude interface {
	implementsZeroTrustAccessApplicationPolicyNewParamsExclude()
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule],
// [ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyNewParamsRequire interface {
	implementsZeroTrustAccessApplicationPolicyNewParamsRequire()
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPRule struct {
	IP param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessApplicationPolicyNewParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessApplicationPolicyNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationPolicyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationPolicyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessPolicies                                                `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationPolicyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationPolicyNewResponseEnvelope]
type zeroTrustAccessApplicationPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyNewResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseEnvelopeErrors]
type zeroTrustAccessApplicationPolicyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyNewResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseEnvelopeMessages]
type zeroTrustAccessApplicationPolicyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationPolicyNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationPolicyNewResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationPolicyNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationPolicyUpdateParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]ZeroTrustAccessApplicationPolicyUpdateParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ZeroTrustAccessApplicationPolicyUpdateParamsApprovalGroup] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ZeroTrustAccessApplicationPolicyUpdateParamsExclude] `json:"exclude"`
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
	Require param.Field[[]ZeroTrustAccessApplicationPolicyUpdateParamsRequire] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type ZeroTrustAccessApplicationPolicyUpdateParamsDecision string

const (
	ZeroTrustAccessApplicationPolicyUpdateParamsDecisionAllow       ZeroTrustAccessApplicationPolicyUpdateParamsDecision = "allow"
	ZeroTrustAccessApplicationPolicyUpdateParamsDecisionDeny        ZeroTrustAccessApplicationPolicyUpdateParamsDecision = "deny"
	ZeroTrustAccessApplicationPolicyUpdateParamsDecisionNonIdentity ZeroTrustAccessApplicationPolicyUpdateParamsDecision = "non_identity"
	ZeroTrustAccessApplicationPolicyUpdateParamsDecisionBypass      ZeroTrustAccessApplicationPolicyUpdateParamsDecision = "bypass"
)

// Matches a specific email.
//
// Satisfied by
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyUpdateParamsInclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude()
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of email addresses that can approve a temporary authentication request.
type ZeroTrustAccessApplicationPolicyUpdateParamsApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]interface{}] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID param.Field[string] `json:"email_list_uuid"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsApprovalGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyUpdateParamsExclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude()
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule],
// [ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyUpdateParamsRequire interface {
	implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire()
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPRule struct {
	IP param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessApplicationPolicyUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessPolicies                                                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationPolicyUpdateResponseEnvelope]
type zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrors]
type zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessages]
type zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationPolicyUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationPolicyListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationPolicyListResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationPolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationPolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessPolicies                                               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessApplicationPolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessApplicationPolicyListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationPolicyListResponseEnvelope]
type zeroTrustAccessApplicationPolicyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyListResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseEnvelopeErrors]
type zeroTrustAccessApplicationPolicyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyListResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseEnvelopeMessages]
type zeroTrustAccessApplicationPolicyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationPolicyListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationPolicyListResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationPolicyListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                            `json:"total_count"`
	JSON       zeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfo]
type zeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationPolicyDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessApplicationPolicyDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationPolicyDeleteResponseEnvelope]
type zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrors]
type zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessages]
type zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationPolicyDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationPolicyGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationPolicyGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationPolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationPolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessPolicies                                                `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationPolicyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationPolicyGetResponseEnvelope]
type zeroTrustAccessApplicationPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyGetResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseEnvelopeErrors]
type zeroTrustAccessApplicationPolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessApplicationPolicyGetResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseEnvelopeMessages]
type zeroTrustAccessApplicationPolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessApplicationPolicyGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationPolicyGetResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationPolicyGetResponseEnvelopeSuccess = true
)
