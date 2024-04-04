// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *AccessApplicationPolicyService) New(ctx context.Context, uuid string, params AccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *ZeroTrustPolicies, err error) {
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
func (r *AccessApplicationPolicyService) Update(ctx context.Context, uuid1 string, uuid string, params AccessApplicationPolicyUpdateParams, opts ...option.RequestOption) (res *ZeroTrustPolicies, err error) {
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
func (r *AccessApplicationPolicyService) List(ctx context.Context, uuid string, query AccessApplicationPolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustPolicies], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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

// Lists Access policies configured for an application.
func (r *AccessApplicationPolicyService) ListAutoPaging(ctx context.Context, uuid string, query AccessApplicationPolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustPolicies] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, uuid, query, opts...))
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
func (r *AccessApplicationPolicyService) Get(ctx context.Context, uuid1 string, uuid string, query AccessApplicationPolicyGetParams, opts ...option.RequestOption) (res *ZeroTrustPolicies, err error) {
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

type ZeroTrustPolicies struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ZeroTrustPoliciesApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision ZeroTrustPoliciesDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustPoliciesExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustPoliciesInclude `json:"include"`
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
	Require []ZeroTrustPoliciesRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                `json:"session_duration"`
	UpdatedAt       time.Time             `json:"updated_at" format:"date-time"`
	JSON            zeroTrustPoliciesJSON `json:"-"`
}

// zeroTrustPoliciesJSON contains the JSON metadata for the struct
// [ZeroTrustPolicies]
type zeroTrustPoliciesJSON struct {
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

func (r *ZeroTrustPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesJSON) RawJSON() string {
	return r.raw
}

// A group of email addresses that can approve a temporary authentication request.
type ZeroTrustPoliciesApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []string `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                             `json:"email_list_uuid"`
	JSON          zeroTrustPoliciesApprovalGroupJSON `json:"-"`
}

// zeroTrustPoliciesApprovalGroupJSON contains the JSON metadata for the struct
// [ZeroTrustPoliciesApprovalGroup]
type zeroTrustPoliciesApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustPoliciesApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesApprovalGroupJSON) RawJSON() string {
	return r.raw
}

// The action Access will take if a user matches this policy.
type ZeroTrustPoliciesDecision string

const (
	ZeroTrustPoliciesDecisionAllow       ZeroTrustPoliciesDecision = "allow"
	ZeroTrustPoliciesDecisionDeny        ZeroTrustPoliciesDecision = "deny"
	ZeroTrustPoliciesDecisionNonIdentity ZeroTrustPoliciesDecision = "non_identity"
	ZeroTrustPoliciesDecisionBypass      ZeroTrustPoliciesDecision = "bypass"
)

func (r ZeroTrustPoliciesDecision) IsKnown() bool {
	switch r {
	case ZeroTrustPoliciesDecisionAllow, ZeroTrustPoliciesDecisionDeny, ZeroTrustPoliciesDecisionNonIdentity, ZeroTrustPoliciesDecisionBypass:
		return true
	}
	return false
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.ZeroTrustPoliciesExcludeAccessEmailRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessEmailListRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessDomainRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessEveryoneRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessIPRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessIPListRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessCertificateRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessAccessGroupRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessAzureGroupRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessGitHubOrganizationRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessGsuiteGroupRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessOktaGroupRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessSamlGroupRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessServiceTokenRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessExternalEvaluationRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessCountryRule],
// [zero_trust.ZeroTrustPoliciesExcludeAccessAuthenticationMethodRule] or
// [zero_trust.ZeroTrustPoliciesExcludeAccessDevicePostureRule].
type ZeroTrustPoliciesExclude interface {
	implementsZeroTrustZeroTrustPoliciesExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustPoliciesExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesExcludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustPoliciesExcludeAccessEmailRule struct {
	Email ZeroTrustPoliciesExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustPoliciesExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustPoliciesExcludeAccessEmailRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesExcludeAccessEmailRule]
type zeroTrustPoliciesExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessEmailRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                           `json:"email,required" format:"email"`
	JSON  zeroTrustPoliciesExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessEmailRuleEmailJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessEmailRuleEmail]
type zeroTrustPoliciesExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustPoliciesExcludeAccessEmailListRule struct {
	EmailList ZeroTrustPoliciesExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustPoliciesExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustPoliciesExcludeAccessEmailListRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessEmailListRule]
type zeroTrustPoliciesExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessEmailListRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                   `json:"id,required"`
	JSON zeroTrustPoliciesExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessEmailListRuleEmailList]
type zeroTrustPoliciesExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustPoliciesExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustPoliciesExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustPoliciesExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustPoliciesExcludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesExcludeAccessDomainRule]
type zeroTrustPoliciesExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessDomainRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                  `json:"domain,required"`
	JSON   zeroTrustPoliciesExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessDomainRuleEmailDomain]
type zeroTrustPoliciesExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustPoliciesExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                    `json:"everyone,required"`
	JSON     zeroTrustPoliciesExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessEveryoneRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessEveryoneRule]
type zeroTrustPoliciesExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessEveryoneRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

// Matches an IP address block.
type ZeroTrustPoliciesExcludeAccessIPRule struct {
	IP   ZeroTrustPoliciesExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustPoliciesExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessIPRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesExcludeAccessIPRule]
type zeroTrustPoliciesExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessIPRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                     `json:"ip,required"`
	JSON zeroTrustPoliciesExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessIPRuleIPJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesExcludeAccessIPRuleIP]
type zeroTrustPoliciesExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustPoliciesExcludeAccessIPListRule struct {
	IPList ZeroTrustPoliciesExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustPoliciesExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustPoliciesExcludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesExcludeAccessIPListRule]
type zeroTrustPoliciesExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessIPListRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                             `json:"id,required"`
	JSON zeroTrustPoliciesExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessIPListRuleIPListJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesExcludeAccessIPListRuleIPList]
type zeroTrustPoliciesExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustPoliciesExcludeAccessCertificateRule struct {
	Certificate interface{}                                       `json:"certificate,required"`
	JSON        zeroTrustPoliciesExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessCertificateRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessCertificateRule]
type zeroTrustPoliciesExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessCertificateRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

// Matches an Access group.
type ZeroTrustPoliciesExcludeAccessAccessGroupRule struct {
	Group shared.UnnamedSchemaRef131                        `json:"group,required"`
	JSON  zeroTrustPoliciesExcludeAccessAccessGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessAccessGroupRule]
type zeroTrustPoliciesExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessAccessGroupRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustPoliciesExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustPoliciesExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustPoliciesExcludeAccessAzureGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessAzureGroupRule]
type zeroTrustPoliciesExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessAzureGroupRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                  `json:"connection_id,required"`
	JSON         zeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustPoliciesExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessGitHubOrganizationRule]
type zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessGitHubOrganizationRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

type ZeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                     `json:"name,required"`
	JSON zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustPoliciesExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustPoliciesExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustPoliciesExcludeAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessGsuiteGroupRule]
type zeroTrustPoliciesExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessGsuiteGroupRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

type ZeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                  `json:"email,required"`
	JSON  zeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustPoliciesExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustPoliciesExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustPoliciesExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessOktaGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessOktaGroupRule]
type zeroTrustPoliciesExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessOktaGroupRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                              `json:"email,required"`
	JSON  zeroTrustPoliciesExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessOktaGroupRuleOktaJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesExcludeAccessOktaGroupRuleOkta]
type zeroTrustPoliciesExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustPoliciesExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustPoliciesExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustPoliciesExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessSamlGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessSamlGroupRule]
type zeroTrustPoliciesExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessSamlGroupRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                              `json:"attribute_value,required"`
	JSON           zeroTrustPoliciesExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessSamlGroupRuleSamlJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesExcludeAccessSamlGroupRuleSaml]
type zeroTrustPoliciesExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustPoliciesExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustPoliciesExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustPoliciesExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustPoliciesExcludeAccessServiceTokenRuleJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesExcludeAccessServiceTokenRule]
type zeroTrustPoliciesExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessServiceTokenRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

type ZeroTrustPoliciesExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                         `json:"token_id,required"`
	JSON    zeroTrustPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [ZeroTrustPoliciesExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustPoliciesExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                `json:"any_valid_service_token,required"`
	JSON                 zeroTrustPoliciesExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessAnyValidServiceTokenRule]
type zeroTrustPoliciesExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustPoliciesExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustPoliciesExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustPoliciesExcludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessExternalEvaluationRule]
type zeroTrustPoliciesExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessExternalEvaluationRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

type ZeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                     `json:"keys_url,required"`
	JSON    zeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustPoliciesExcludeAccessCountryRule struct {
	Geo  ZeroTrustPoliciesExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustPoliciesExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesExcludeAccessCountryRule]
type zeroTrustPoliciesExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessCountryRule) implementsZeroTrustZeroTrustPoliciesExclude() {}

type ZeroTrustPoliciesExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                           `json:"country_code,required"`
	JSON        zeroTrustPoliciesExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessCountryRuleGeoJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesExcludeAccessCountryRuleGeo]
type zeroTrustPoliciesExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustPoliciesExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesExcludeAccessAuthenticationMethodRule]
type zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessAuthenticationMethodRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

type ZeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                               `json:"auth_method,required"`
	JSON       zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON contains
// the JSON metadata for the struct
// [ZeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustPoliciesExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustPoliciesExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustPoliciesExcludeAccessDevicePostureRuleJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesExcludeAccessDevicePostureRule]
type zeroTrustPoliciesExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesExcludeAccessDevicePostureRule) implementsZeroTrustZeroTrustPoliciesExclude() {
}

type ZeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                           `json:"integration_uid,required"`
	JSON           zeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON contains the
// JSON metadata for the struct
// [ZeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.ZeroTrustPoliciesIncludeAccessEmailRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessEmailListRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessDomainRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessEveryoneRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessIPRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessIPListRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessCertificateRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessAccessGroupRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessAzureGroupRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessGitHubOrganizationRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessGsuiteGroupRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessOktaGroupRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessSamlGroupRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessServiceTokenRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessExternalEvaluationRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessCountryRule],
// [zero_trust.ZeroTrustPoliciesIncludeAccessAuthenticationMethodRule] or
// [zero_trust.ZeroTrustPoliciesIncludeAccessDevicePostureRule].
type ZeroTrustPoliciesInclude interface {
	implementsZeroTrustZeroTrustPoliciesInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustPoliciesInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesIncludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustPoliciesIncludeAccessEmailRule struct {
	Email ZeroTrustPoliciesIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustPoliciesIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustPoliciesIncludeAccessEmailRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesIncludeAccessEmailRule]
type zeroTrustPoliciesIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessEmailRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                           `json:"email,required" format:"email"`
	JSON  zeroTrustPoliciesIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessEmailRuleEmailJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessEmailRuleEmail]
type zeroTrustPoliciesIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustPoliciesIncludeAccessEmailListRule struct {
	EmailList ZeroTrustPoliciesIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustPoliciesIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustPoliciesIncludeAccessEmailListRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessEmailListRule]
type zeroTrustPoliciesIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessEmailListRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                   `json:"id,required"`
	JSON zeroTrustPoliciesIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessEmailListRuleEmailList]
type zeroTrustPoliciesIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustPoliciesIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustPoliciesIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustPoliciesIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustPoliciesIncludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesIncludeAccessDomainRule]
type zeroTrustPoliciesIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessDomainRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                  `json:"domain,required"`
	JSON   zeroTrustPoliciesIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessDomainRuleEmailDomain]
type zeroTrustPoliciesIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustPoliciesIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                    `json:"everyone,required"`
	JSON     zeroTrustPoliciesIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessEveryoneRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessEveryoneRule]
type zeroTrustPoliciesIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessEveryoneRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

// Matches an IP address block.
type ZeroTrustPoliciesIncludeAccessIPRule struct {
	IP   ZeroTrustPoliciesIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustPoliciesIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessIPRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesIncludeAccessIPRule]
type zeroTrustPoliciesIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessIPRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                     `json:"ip,required"`
	JSON zeroTrustPoliciesIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessIPRuleIPJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesIncludeAccessIPRuleIP]
type zeroTrustPoliciesIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustPoliciesIncludeAccessIPListRule struct {
	IPList ZeroTrustPoliciesIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustPoliciesIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustPoliciesIncludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesIncludeAccessIPListRule]
type zeroTrustPoliciesIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessIPListRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                             `json:"id,required"`
	JSON zeroTrustPoliciesIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessIPListRuleIPListJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesIncludeAccessIPListRuleIPList]
type zeroTrustPoliciesIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustPoliciesIncludeAccessCertificateRule struct {
	Certificate interface{}                                       `json:"certificate,required"`
	JSON        zeroTrustPoliciesIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessCertificateRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessCertificateRule]
type zeroTrustPoliciesIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessCertificateRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

// Matches an Access group.
type ZeroTrustPoliciesIncludeAccessAccessGroupRule struct {
	Group shared.UnnamedSchemaRef131                        `json:"group,required"`
	JSON  zeroTrustPoliciesIncludeAccessAccessGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessAccessGroupRule]
type zeroTrustPoliciesIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessAccessGroupRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustPoliciesIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustPoliciesIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustPoliciesIncludeAccessAzureGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessAzureGroupRule]
type zeroTrustPoliciesIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessAzureGroupRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                  `json:"connection_id,required"`
	JSON         zeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustPoliciesIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessGitHubOrganizationRule]
type zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessGitHubOrganizationRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

type ZeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                     `json:"name,required"`
	JSON zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustPoliciesIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustPoliciesIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustPoliciesIncludeAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessGsuiteGroupRule]
type zeroTrustPoliciesIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessGsuiteGroupRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

type ZeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                  `json:"email,required"`
	JSON  zeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustPoliciesIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustPoliciesIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustPoliciesIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessOktaGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessOktaGroupRule]
type zeroTrustPoliciesIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessOktaGroupRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                              `json:"email,required"`
	JSON  zeroTrustPoliciesIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessOktaGroupRuleOktaJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesIncludeAccessOktaGroupRuleOkta]
type zeroTrustPoliciesIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustPoliciesIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustPoliciesIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustPoliciesIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessSamlGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessSamlGroupRule]
type zeroTrustPoliciesIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessSamlGroupRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                              `json:"attribute_value,required"`
	JSON           zeroTrustPoliciesIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessSamlGroupRuleSamlJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesIncludeAccessSamlGroupRuleSaml]
type zeroTrustPoliciesIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustPoliciesIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustPoliciesIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustPoliciesIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustPoliciesIncludeAccessServiceTokenRuleJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesIncludeAccessServiceTokenRule]
type zeroTrustPoliciesIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessServiceTokenRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

type ZeroTrustPoliciesIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                         `json:"token_id,required"`
	JSON    zeroTrustPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [ZeroTrustPoliciesIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustPoliciesIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                `json:"any_valid_service_token,required"`
	JSON                 zeroTrustPoliciesIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessAnyValidServiceTokenRule]
type zeroTrustPoliciesIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustPoliciesIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustPoliciesIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustPoliciesIncludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessExternalEvaluationRule]
type zeroTrustPoliciesIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessExternalEvaluationRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

type ZeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                     `json:"keys_url,required"`
	JSON    zeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustPoliciesIncludeAccessCountryRule struct {
	Geo  ZeroTrustPoliciesIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustPoliciesIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesIncludeAccessCountryRule]
type zeroTrustPoliciesIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessCountryRule) implementsZeroTrustZeroTrustPoliciesInclude() {}

type ZeroTrustPoliciesIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                           `json:"country_code,required"`
	JSON        zeroTrustPoliciesIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessCountryRuleGeoJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesIncludeAccessCountryRuleGeo]
type zeroTrustPoliciesIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustPoliciesIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesIncludeAccessAuthenticationMethodRule]
type zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessAuthenticationMethodRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

type ZeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                               `json:"auth_method,required"`
	JSON       zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON contains
// the JSON metadata for the struct
// [ZeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustPoliciesIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustPoliciesIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustPoliciesIncludeAccessDevicePostureRuleJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesIncludeAccessDevicePostureRule]
type zeroTrustPoliciesIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesIncludeAccessDevicePostureRule) implementsZeroTrustZeroTrustPoliciesInclude() {
}

type ZeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                           `json:"integration_uid,required"`
	JSON           zeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON contains the
// JSON metadata for the struct
// [ZeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.ZeroTrustPoliciesRequireAccessEmailRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessEmailListRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessDomainRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessEveryoneRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessIPRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessIPListRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessCertificateRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessAccessGroupRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessAzureGroupRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessGitHubOrganizationRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessGsuiteGroupRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessOktaGroupRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessSamlGroupRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessServiceTokenRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessAnyValidServiceTokenRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessExternalEvaluationRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessCountryRule],
// [zero_trust.ZeroTrustPoliciesRequireAccessAuthenticationMethodRule] or
// [zero_trust.ZeroTrustPoliciesRequireAccessDevicePostureRule].
type ZeroTrustPoliciesRequire interface {
	implementsZeroTrustZeroTrustPoliciesRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustPoliciesRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustPoliciesRequireAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustPoliciesRequireAccessEmailRule struct {
	Email ZeroTrustPoliciesRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustPoliciesRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustPoliciesRequireAccessEmailRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesRequireAccessEmailRule]
type zeroTrustPoliciesRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessEmailRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                           `json:"email,required" format:"email"`
	JSON  zeroTrustPoliciesRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessEmailRuleEmailJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessEmailRuleEmail]
type zeroTrustPoliciesRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustPoliciesRequireAccessEmailListRule struct {
	EmailList ZeroTrustPoliciesRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustPoliciesRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustPoliciesRequireAccessEmailListRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessEmailListRule]
type zeroTrustPoliciesRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessEmailListRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                   `json:"id,required"`
	JSON zeroTrustPoliciesRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessEmailListRuleEmailList]
type zeroTrustPoliciesRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustPoliciesRequireAccessDomainRule struct {
	EmailDomain ZeroTrustPoliciesRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustPoliciesRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustPoliciesRequireAccessDomainRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesRequireAccessDomainRule]
type zeroTrustPoliciesRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessDomainRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                  `json:"domain,required"`
	JSON   zeroTrustPoliciesRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessDomainRuleEmailDomain]
type zeroTrustPoliciesRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustPoliciesRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                    `json:"everyone,required"`
	JSON     zeroTrustPoliciesRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessEveryoneRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessEveryoneRule]
type zeroTrustPoliciesRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessEveryoneRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

// Matches an IP address block.
type ZeroTrustPoliciesRequireAccessIPRule struct {
	IP   ZeroTrustPoliciesRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustPoliciesRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessIPRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesRequireAccessIPRule]
type zeroTrustPoliciesRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessIPRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                     `json:"ip,required"`
	JSON zeroTrustPoliciesRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessIPRuleIPJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesRequireAccessIPRuleIP]
type zeroTrustPoliciesRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustPoliciesRequireAccessIPListRule struct {
	IPList ZeroTrustPoliciesRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustPoliciesRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustPoliciesRequireAccessIPListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesRequireAccessIPListRule]
type zeroTrustPoliciesRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessIPListRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                             `json:"id,required"`
	JSON zeroTrustPoliciesRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessIPListRuleIPListJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesRequireAccessIPListRuleIPList]
type zeroTrustPoliciesRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustPoliciesRequireAccessCertificateRule struct {
	Certificate interface{}                                       `json:"certificate,required"`
	JSON        zeroTrustPoliciesRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessCertificateRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessCertificateRule]
type zeroTrustPoliciesRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessCertificateRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

// Matches an Access group.
type ZeroTrustPoliciesRequireAccessAccessGroupRule struct {
	Group shared.UnnamedSchemaRef131                        `json:"group,required"`
	JSON  zeroTrustPoliciesRequireAccessAccessGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessAccessGroupRule]
type zeroTrustPoliciesRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessAccessGroupRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustPoliciesRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustPoliciesRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustPoliciesRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustPoliciesRequireAccessAzureGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessAzureGroupRule]
type zeroTrustPoliciesRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessAzureGroupRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                  `json:"connection_id,required"`
	JSON         zeroTrustPoliciesRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessAzureGroupRuleAzureAd]
type zeroTrustPoliciesRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustPoliciesRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustPoliciesRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustPoliciesRequireAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessGitHubOrganizationRule]
type zeroTrustPoliciesRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessGitHubOrganizationRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

type ZeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                     `json:"name,required"`
	JSON zeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustPoliciesRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustPoliciesRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustPoliciesRequireAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessGsuiteGroupRule]
type zeroTrustPoliciesRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessGsuiteGroupRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

type ZeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                  `json:"email,required"`
	JSON  zeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustPoliciesRequireAccessOktaGroupRule struct {
	Okta ZeroTrustPoliciesRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustPoliciesRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessOktaGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessOktaGroupRule]
type zeroTrustPoliciesRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessOktaGroupRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                              `json:"email,required"`
	JSON  zeroTrustPoliciesRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessOktaGroupRuleOktaJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesRequireAccessOktaGroupRuleOkta]
type zeroTrustPoliciesRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustPoliciesRequireAccessSamlGroupRule struct {
	Saml ZeroTrustPoliciesRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustPoliciesRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessSamlGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessSamlGroupRule]
type zeroTrustPoliciesRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessSamlGroupRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                              `json:"attribute_value,required"`
	JSON           zeroTrustPoliciesRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessSamlGroupRuleSamlJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesRequireAccessSamlGroupRuleSaml]
type zeroTrustPoliciesRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustPoliciesRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustPoliciesRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustPoliciesRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustPoliciesRequireAccessServiceTokenRuleJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesRequireAccessServiceTokenRule]
type zeroTrustPoliciesRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessServiceTokenRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

type ZeroTrustPoliciesRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                         `json:"token_id,required"`
	JSON    zeroTrustPoliciesRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [ZeroTrustPoliciesRequireAccessServiceTokenRuleServiceToken]
type zeroTrustPoliciesRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustPoliciesRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                `json:"any_valid_service_token,required"`
	JSON                 zeroTrustPoliciesRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessAnyValidServiceTokenRule]
type zeroTrustPoliciesRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessAnyValidServiceTokenRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustPoliciesRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustPoliciesRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustPoliciesRequireAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessExternalEvaluationRule]
type zeroTrustPoliciesRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessExternalEvaluationRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

type ZeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                     `json:"keys_url,required"`
	JSON    zeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustPoliciesRequireAccessCountryRule struct {
	Geo  ZeroTrustPoliciesRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustPoliciesRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessCountryRuleJSON contains the JSON metadata for the
// struct [ZeroTrustPoliciesRequireAccessCountryRule]
type zeroTrustPoliciesRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessCountryRule) implementsZeroTrustZeroTrustPoliciesRequire() {}

type ZeroTrustPoliciesRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                           `json:"country_code,required"`
	JSON        zeroTrustPoliciesRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessCountryRuleGeoJSON contains the JSON metadata for
// the struct [ZeroTrustPoliciesRequireAccessCountryRuleGeo]
type zeroTrustPoliciesRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustPoliciesRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustPoliciesRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustPoliciesRequireAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [ZeroTrustPoliciesRequireAccessAuthenticationMethodRule]
type zeroTrustPoliciesRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessAuthenticationMethodRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

type ZeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                               `json:"auth_method,required"`
	JSON       zeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON contains
// the JSON metadata for the struct
// [ZeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustPoliciesRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustPoliciesRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustPoliciesRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustPoliciesRequireAccessDevicePostureRuleJSON contains the JSON metadata
// for the struct [ZeroTrustPoliciesRequireAccessDevicePostureRule]
type zeroTrustPoliciesRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustPoliciesRequireAccessDevicePostureRule) implementsZeroTrustZeroTrustPoliciesRequire() {
}

type ZeroTrustPoliciesRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                           `json:"integration_uid,required"`
	JSON           zeroTrustPoliciesRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustPoliciesRequireAccessDevicePostureRuleDevicePostureJSON contains the
// JSON metadata for the struct
// [ZeroTrustPoliciesRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustPoliciesRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustPoliciesRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustPoliciesRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
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

func (r AccessApplicationPolicyNewParamsDecision) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewParamsDecisionAllow, AccessApplicationPolicyNewParamsDecisionDeny, AccessApplicationPolicyNewParamsDecisionNonIdentity, AccessApplicationPolicyNewParamsDecisionBypass:
		return true
	}
	return false
}

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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsInclude() {
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
	EmailAddresses param.Field[[]string] `json:"email_addresses"`
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsExclude() {
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewParamsRequire() {
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   ZeroTrustPolicies            `json:"result,required"`
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

// Whether the API call was successful
type AccessApplicationPolicyNewResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyNewResponseEnvelopeSuccessTrue AccessApplicationPolicyNewResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r AccessApplicationPolicyUpdateParamsDecision) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateParamsDecisionAllow, AccessApplicationPolicyUpdateParamsDecisionDeny, AccessApplicationPolicyUpdateParamsDecisionNonIdentity, AccessApplicationPolicyUpdateParamsDecisionBypass:
		return true
	}
	return false
}

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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsInclude() {
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
	EmailAddresses param.Field[[]string] `json:"email_addresses"`
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsExclude() {
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateParamsRequire() {
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   ZeroTrustPolicies            `json:"result,required"`
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

// Whether the API call was successful
type AccessApplicationPolicyUpdateResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyUpdateResponseEnvelopeSuccessTrue AccessApplicationPolicyUpdateResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationPolicyListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationPolicyDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationPolicyDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172          `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172          `json:"messages,required"`
	Result   AccessApplicationPolicyDeleteResponse `json:"result,required"`
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

// Whether the API call was successful
type AccessApplicationPolicyDeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyDeleteResponseEnvelopeSuccessTrue AccessApplicationPolicyDeleteResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationPolicyGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationPolicyGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   ZeroTrustPolicies            `json:"result,required"`
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

// Whether the API call was successful
type AccessApplicationPolicyGetResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyGetResponseEnvelopeSuccessTrue AccessApplicationPolicyGetResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
