// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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
func (r *AccessApplicationPolicyService) New(ctx context.Context, uuid string, params AccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *AccessPolicies, err error) {
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
func (r *AccessApplicationPolicyService) Update(ctx context.Context, uuid1 string, uuid string, params AccessApplicationPolicyUpdateParams, opts ...option.RequestOption) (res *AccessPolicies, err error) {
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
func (r *AccessApplicationPolicyService) List(ctx context.Context, uuid string, query AccessApplicationPolicyListParams, opts ...option.RequestOption) (res *[]AccessPolicies, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access policy.
func (r *AccessApplicationPolicyService) Get(ctx context.Context, uuid1 string, uuid string, query AccessApplicationPolicyGetParams, opts ...option.RequestOption) (res *AccessPolicies, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r accessPoliciesJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesApprovalGroupJSON) RawJSON() string {
	return r.raw
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
// Union satisfied by [zero_trust.AccessPoliciesExcludeAccessEmailRule],
// [zero_trust.AccessPoliciesExcludeAccessEmailListRule],
// [zero_trust.AccessPoliciesExcludeAccessDomainRule],
// [zero_trust.AccessPoliciesExcludeAccessEveryoneRule],
// [zero_trust.AccessPoliciesExcludeAccessIPRule],
// [zero_trust.AccessPoliciesExcludeAccessIPListRule],
// [zero_trust.AccessPoliciesExcludeAccessCertificateRule],
// [zero_trust.AccessPoliciesExcludeAccessAccessGroupRule],
// [zero_trust.AccessPoliciesExcludeAccessAzureGroupRule],
// [zero_trust.AccessPoliciesExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessPoliciesExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessPoliciesExcludeAccessOktaGroupRule],
// [zero_trust.AccessPoliciesExcludeAccessSamlGroupRule],
// [zero_trust.AccessPoliciesExcludeAccessServiceTokenRule],
// [zero_trust.AccessPoliciesExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessPoliciesExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessPoliciesExcludeAccessCountryRule],
// [zero_trust.AccessPoliciesExcludeAccessAuthenticationMethodRule] or
// [zero_trust.AccessPoliciesExcludeAccessDevicePostureRule].
type AccessPoliciesExclude interface {
	implementsZeroTrustAccessPoliciesExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessPoliciesExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesExcludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessPoliciesExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessEmailRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessEmailListRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessDomainRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessEveryoneRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessIPRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessIPListRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessCertificateRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessAccessGroupRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessAzureGroupRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessPoliciesExclude() {
}

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

func (r accessPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessOktaGroupRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessSamlGroupRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessServiceTokenRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessPoliciesExclude() {
}

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

func (r accessPoliciesExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessPoliciesExclude() {
}

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

func (r accessPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessCountryRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessPoliciesExclude() {
}

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

func (r accessPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesExcludeAccessDevicePostureRule) implementsZeroTrustAccessPoliciesExclude() {}

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

func (r accessPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.AccessPoliciesIncludeAccessEmailRule],
// [zero_trust.AccessPoliciesIncludeAccessEmailListRule],
// [zero_trust.AccessPoliciesIncludeAccessDomainRule],
// [zero_trust.AccessPoliciesIncludeAccessEveryoneRule],
// [zero_trust.AccessPoliciesIncludeAccessIPRule],
// [zero_trust.AccessPoliciesIncludeAccessIPListRule],
// [zero_trust.AccessPoliciesIncludeAccessCertificateRule],
// [zero_trust.AccessPoliciesIncludeAccessAccessGroupRule],
// [zero_trust.AccessPoliciesIncludeAccessAzureGroupRule],
// [zero_trust.AccessPoliciesIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessPoliciesIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessPoliciesIncludeAccessOktaGroupRule],
// [zero_trust.AccessPoliciesIncludeAccessSamlGroupRule],
// [zero_trust.AccessPoliciesIncludeAccessServiceTokenRule],
// [zero_trust.AccessPoliciesIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessPoliciesIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessPoliciesIncludeAccessCountryRule],
// [zero_trust.AccessPoliciesIncludeAccessAuthenticationMethodRule] or
// [zero_trust.AccessPoliciesIncludeAccessDevicePostureRule].
type AccessPoliciesInclude interface {
	implementsZeroTrustAccessPoliciesInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessPoliciesInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesIncludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessPoliciesIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessEmailRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessEmailListRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessDomainRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessEveryoneRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessIPRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessIPListRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessCertificateRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessAccessGroupRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessAzureGroupRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessPoliciesInclude() {
}

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

func (r accessPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessOktaGroupRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessSamlGroupRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessServiceTokenRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessPoliciesInclude() {
}

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

func (r accessPoliciesIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessPoliciesInclude() {
}

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

func (r accessPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessCountryRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessPoliciesInclude() {
}

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

func (r accessPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesIncludeAccessDevicePostureRule) implementsZeroTrustAccessPoliciesInclude() {}

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

func (r accessPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.AccessPoliciesRequireAccessEmailRule],
// [zero_trust.AccessPoliciesRequireAccessEmailListRule],
// [zero_trust.AccessPoliciesRequireAccessDomainRule],
// [zero_trust.AccessPoliciesRequireAccessEveryoneRule],
// [zero_trust.AccessPoliciesRequireAccessIPRule],
// [zero_trust.AccessPoliciesRequireAccessIPListRule],
// [zero_trust.AccessPoliciesRequireAccessCertificateRule],
// [zero_trust.AccessPoliciesRequireAccessAccessGroupRule],
// [zero_trust.AccessPoliciesRequireAccessAzureGroupRule],
// [zero_trust.AccessPoliciesRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessPoliciesRequireAccessGsuiteGroupRule],
// [zero_trust.AccessPoliciesRequireAccessOktaGroupRule],
// [zero_trust.AccessPoliciesRequireAccessSamlGroupRule],
// [zero_trust.AccessPoliciesRequireAccessServiceTokenRule],
// [zero_trust.AccessPoliciesRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessPoliciesRequireAccessExternalEvaluationRule],
// [zero_trust.AccessPoliciesRequireAccessCountryRule],
// [zero_trust.AccessPoliciesRequireAccessAuthenticationMethodRule] or
// [zero_trust.AccessPoliciesRequireAccessDevicePostureRule].
type AccessPoliciesRequire interface {
	implementsZeroTrustAccessPoliciesRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessPoliciesRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessPoliciesRequireAccessDevicePostureRule{}),
		},
	)
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

func (r accessPoliciesRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessEmailRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessEmailListRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessDomainRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessEveryoneRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessIPRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessIPListRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessCertificateRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessAccessGroupRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessAzureGroupRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessPoliciesRequire() {
}

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

func (r accessPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessGsuiteGroupRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessOktaGroupRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessSamlGroupRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessServiceTokenRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessPoliciesRequire() {
}

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

func (r accessPoliciesRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessExternalEvaluationRule) implementsZeroTrustAccessPoliciesRequire() {
}

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

func (r accessPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessCountryRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessPoliciesRequire() {
}

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

func (r accessPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessPoliciesRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessPoliciesRequireAccessDevicePostureRule) implementsZeroTrustAccessPoliciesRequire() {}

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

func (r accessPoliciesRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
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
// Satisfied by
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule],
// [zero_trust.AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule].
type AccessApplicationPolicyNewParamsInclude interface {
	implementsZeroTrustAccessApplicationPolicyNewParamsInclude()
}

// Matches a specific email.
type AccessApplicationPolicyNewParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyNewParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewParamsIncludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyNewParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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

func (r AccessApplicationPolicyNewParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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
// Satisfied by
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule],
// [zero_trust.AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule].
type AccessApplicationPolicyNewParamsExclude interface {
	implementsZeroTrustAccessApplicationPolicyNewParamsExclude()
}

// Matches a specific email.
type AccessApplicationPolicyNewParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyNewParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewParamsExcludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyNewParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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

func (r AccessApplicationPolicyNewParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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
// Satisfied by
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessDomainRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessIPRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessIPListRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessCountryRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule],
// [zero_trust.AccessApplicationPolicyNewParamsRequireAccessDevicePostureRule].
type AccessApplicationPolicyNewParamsRequire interface {
	implementsZeroTrustAccessApplicationPolicyNewParamsRequire()
}

// Matches a specific email.
type AccessApplicationPolicyNewParamsRequireAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyNewParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyNewParamsRequireAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyNewParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyNewParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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

func (r AccessApplicationPolicyNewParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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
	Result   AccessPolicies                                       `json:"result,required"`
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

func (r accessApplicationPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
// Satisfied by
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateParamsInclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude()
}

// Matches a specific email.
type AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateParamsIncludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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

func (r AccessApplicationPolicyUpdateParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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
// Satisfied by
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateParamsExclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude()
}

// Matches a specific email.
type AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateParamsExcludeAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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

func (r AccessApplicationPolicyUpdateParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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
// Satisfied by
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessDomainRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessIPRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessIPListRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessCountryRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule],
// [zero_trust.AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule].
type AccessApplicationPolicyUpdateParamsRequire interface {
	implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire()
}

// Matches a specific email.
type AccessApplicationPolicyUpdateParamsRequireAccessEmailRule struct {
	Email param.Field[AccessApplicationPolicyUpdateParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

// Matches an IP address block.
type AccessApplicationPolicyUpdateParamsRequireAccessIPRule struct {
	IP param.Field[AccessApplicationPolicyUpdateParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

// Matches an Access group.
type AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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

func (r AccessApplicationPolicyUpdateParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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
	Result   AccessPolicies                                          `json:"result,required"`
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

func (r accessApplicationPolicyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   []AccessPolicies                                      `json:"result,required,nullable"`
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

func (r accessApplicationPolicyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   AccessPolicies                                       `json:"result,required"`
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

func (r accessApplicationPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationPolicyGetResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyGetResponseEnvelopeSuccessTrue AccessApplicationPolicyGetResponseEnvelopeSuccess = true
)
