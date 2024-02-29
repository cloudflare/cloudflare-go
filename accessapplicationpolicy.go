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
func (r *AccessApplicationPolicyService) New(ctx context.Context, uuid string, params AccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *AccessApplicationPolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyNewResponseEnvelope
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
func (r *AccessApplicationPolicyService) Update(ctx context.Context, uuid1 string, uuid string, params AccessApplicationPolicyUpdateParams, opts ...option.RequestOption) (res *AccessApplicationPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyUpdateResponseEnvelope
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
func (r *AccessApplicationPolicyService) List(ctx context.Context, uuid string, query AccessApplicationPolicyListParams, opts ...option.RequestOption) (res *[]AccessApplicationPolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyListResponseEnvelope
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
func (r *AccessApplicationPolicyService) Delete(ctx context.Context, uuid1 string, uuid string, body AccessApplicationPolicyDeleteParams, opts ...option.RequestOption) (res *AccessApplicationPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyDeleteResponseEnvelope
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
func (r *AccessApplicationPolicyService) Get(ctx context.Context, uuid1 string, uuid string, query AccessApplicationPolicyGetParams, opts ...option.RequestOption) (res *AccessApplicationPolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationPolicyGetResponseEnvelope
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
	EmailListUUID string                                              `json:"email_list_uuid"`
	JSON          accessApplicationPolicyNewResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyNewResponseApprovalGroupJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyNewResponseApprovalGroup]
type accessApplicationPolicyNewResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
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

type AccessApplicationPolicyUpdateResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []AccessApplicationPolicyUpdateResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision AccessApplicationPolicyUpdateResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessApplicationPolicyUpdateResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessApplicationPolicyUpdateResponseInclude `json:"include"`
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
	Require []AccessApplicationPolicyUpdateResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                    `json:"session_duration"`
	UpdatedAt       time.Time                                 `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyUpdateResponseJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyUpdateResponse]
type accessApplicationPolicyUpdateResponseJSON struct {
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

func (r *AccessApplicationPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyUpdateResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                                                 `json:"email_list_uuid"`
	JSON          accessApplicationPolicyUpdateResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseApprovalGroupJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyUpdateResponseApprovalGroup]
type accessApplicationPolicyUpdateResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyUpdateResponseDecision string

const (
	AccessApplicationPolicyUpdateResponseDecisionAllow       AccessApplicationPolicyUpdateResponseDecision = "allow"
	AccessApplicationPolicyUpdateResponseDecisionDeny        AccessApplicationPolicyUpdateResponseDecision = "deny"
	AccessApplicationPolicyUpdateResponseDecisionNonIdentity AccessApplicationPolicyUpdateResponseDecision = "non_identity"
	AccessApplicationPolicyUpdateResponseDecisionBypass      AccessApplicationPolicyUpdateResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessIPRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule],
// [AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateResponseExclude interface {
	implementsAccessApplicationPolicyUpdateResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyUpdateResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule struct {
	Email AccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule]
type accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail]
type accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule]
type accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList]
type accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule]
type accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     accessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule]
type accessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateResponseExcludeAccessIPRule struct {
	IP   AccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessIPRule]
type accessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessIPRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP]
type accessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule struct {
	IPList AccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule]
type accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList]
type accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        accessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule]
type accessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule]
type accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule]
type accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule]
type accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule]
type accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule]
type accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule]
type accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule]
type accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule]
type accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule]
type accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo]
type accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule]
type accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule]
type accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyUpdateResponseExclude() {
}

type AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessIPRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule],
// [AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateResponseInclude interface {
	implementsAccessApplicationPolicyUpdateResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyUpdateResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule struct {
	Email AccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule]
type accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail]
type accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule struct {
	EmailList AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule]
type accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList]
type accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule]
type accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain]
type accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     accessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule]
type accessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateResponseIncludeAccessIPRule struct {
	IP   AccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessIPRule]
type accessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessIPRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP]
type accessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule struct {
	IPList AccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule]
type accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList]
type accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        accessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule]
type accessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule struct {
	Group AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule]
type accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup]
type accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule]
type accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule]
type accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule]
type accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule]
type accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta]
type accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule]
type accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml]
type accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule]
type accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule]
type accessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule]
type accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule struct {
	Geo  AccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule]
type accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo]
type accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule]
type accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule]
type accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyUpdateResponseInclude() {
}

type AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessApplicationPolicyUpdateResponseRequireAccessEmailRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessDomainRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessIPRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessIPListRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessCountryRule],
// [AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule] or
// [AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyUpdateResponseRequire interface {
	implementsAccessApplicationPolicyUpdateResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessApplicationPolicyUpdateResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessApplicationPolicyUpdateResponseRequireAccessEmailRule struct {
	Email AccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessEmailRule]
type accessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessEmailRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  accessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail]
type accessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule struct {
	EmailList AccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule]
type accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList]
type accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessApplicationPolicyUpdateResponseRequireAccessDomainRule struct {
	EmailDomain AccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessDomainRule]
type accessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessDomainRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   accessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain]
type accessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     accessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule]
type accessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateResponseRequireAccessIPRule struct {
	IP   AccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessIPRule]
type accessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessIPRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP]
type accessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyUpdateResponseRequireAccessIPListRule struct {
	IPList AccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessIPListRule]
type accessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessIPListRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList]
type accessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        accessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule]
type accessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule struct {
	Group AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule]
type accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup]
type accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule]
type accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd]
type accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule]
type accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule]
type accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite]
type accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule struct {
	Okta AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule]
type accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta]
type accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule struct {
	Saml AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule]
type accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml]
type accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule]
type accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken]
type accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule]
type accessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule]
type accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessApplicationPolicyUpdateResponseRequireAccessCountryRule struct {
	Geo  AccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessCountryRule]
type accessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessCountryRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        accessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo]
type accessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule]
type accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule]
type accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule) implementsAccessApplicationPolicyUpdateResponseRequire() {
}

type AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture]
type accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
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
	EmailListUUID string                                               `json:"email_list_uuid"`
	JSON          accessApplicationPolicyListResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyListResponseApprovalGroupJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyListResponseApprovalGroup]
type accessApplicationPolicyListResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
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
	EmailListUUID string                                              `json:"email_list_uuid"`
	JSON          accessApplicationPolicyGetResponseApprovalGroupJSON `json:"-"`
}

// accessApplicationPolicyGetResponseApprovalGroupJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyGetResponseApprovalGroup]
type accessApplicationPolicyGetResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
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

type AccessApplicationPolicyNewParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[AccessApplicationPolicyNewParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessApplicationPolicyNewParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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
	EmailListUUID param.Field[string] `json:"email_list_uuid"`
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

type AccessApplicationPolicyUpdateParams struct {
	// The action Access will take if a user matches this policy.
	Decision param.Field[AccessApplicationPolicyUpdateParamsDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessApplicationPolicyUpdateParamsInclude] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]AccessApplicationPolicyUpdateParamsApprovalGroup] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessApplicationPolicyUpdateParamsExclude] `json:"exclude"`
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
	Require param.Field[[]AccessApplicationPolicyUpdateParamsRequire] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action Access will take if a user matches this policy.
type AccessApplicationPolicyUpdateParamsDecision string

const (
	AccessApplicationPolicyUpdateParamsDecisionAllow       AccessApplicationPolicyUpdateParamsDecision = "allow"
	AccessApplicationPolicyUpdateParamsDecisionDeny        AccessApplicationPolicyUpdateParamsDecision = "deny"
	AccessApplicationPolicyUpdateParamsDecisionNonIdentity AccessApplicationPolicyUpdateParamsDecision = "non_identity"
	AccessApplicationPolicyUpdateParamsDecisionBypass      AccessApplicationPolicyUpdateParamsDecision = "bypass"
)

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessDomainRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessIPRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessIPListRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessCountryRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule],
// [AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateParamsInclude interface {
	implementsAccessApplicationPolicyUpdateParamsInclude()
}

// Matches a specific email.
type AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyUpdateParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDomainRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateParamsIncludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyUpdateParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPListRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyUpdateParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessCountryRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule) implementsAccessApplicationPolicyUpdateParamsInclude() {
}

type AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of email addresses that can approve a temporary authentication request.
type AccessApplicationPolicyUpdateParamsApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]interface{}] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID param.Field[string] `json:"email_list_uuid"`
}

func (r AccessApplicationPolicyUpdateParamsApprovalGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessDomainRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessIPRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessIPListRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessCountryRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule],
// [AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateParamsExclude interface {
	implementsAccessApplicationPolicyUpdateParamsExclude()
}

// Matches a specific email.
type AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyUpdateParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDomainRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateParamsExcludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyUpdateParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPListRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyUpdateParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessCountryRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule) implementsAccessApplicationPolicyUpdateParamsExclude() {
}

type AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessApplicationPolicyUpdateParamsRequireAccessEmailRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessEmailListRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessDomainRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessIPRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessIPListRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessCertificateRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessCountryRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule],
// [AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule].
type AccessApplicationPolicyUpdateParamsRequire interface {
	implementsAccessApplicationPolicyUpdateParamsRequire()
}

// Matches a specific email.
type AccessApplicationPolicyUpdateParamsRequireAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessApplicationPolicyUpdateParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessApplicationPolicyUpdateParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailListRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessApplicationPolicyUpdateParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessApplicationPolicyUpdateParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessDomainRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateParamsRequireAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyUpdateParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessApplicationPolicyUpdateParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessApplicationPolicyUpdateParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPListRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessApplicationPolicyUpdateParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessCertificateRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessApplicationPolicyUpdateParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessApplicationPolicyUpdateParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessCountryRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule) implementsAccessApplicationPolicyUpdateParamsRequire() {
}

type AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyUpdateResponseEnvelope struct {
	Errors   []AccessApplicationPolicyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationPolicyUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationPolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationPolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyUpdateResponseEnvelope]
type accessApplicationPolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyUpdateResponseEnvelopeErrors]
type accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationPolicyUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyUpdateResponseEnvelopeMessages]
type accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationPolicyUpdateResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyUpdateResponseEnvelopeSuccessTrue AccessApplicationPolicyUpdateResponseEnvelopeSuccess = true
)

type AccessApplicationPolicyListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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

type AccessApplicationPolicyDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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

type AccessApplicationPolicyGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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
