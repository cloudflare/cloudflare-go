// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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

func (r accessApplicationPolicyNewResponseJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
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
// Union satisfied by
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyNewResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyNewResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyNewResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
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

func (r accessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyNewResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyNewResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyNewResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
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

func (r accessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessEmailRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessDomainRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessIPRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessIPListRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessCountryRule],
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyNewResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyNewResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyNewResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyNewResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
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

func (r accessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
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
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyUpdateResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
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

func (r accessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyUpdateResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyUpdateResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
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

func (r accessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessEmailRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessDomainRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessIPRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessIPListRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessCountryRule],
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyUpdateResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyUpdateResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
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

func (r accessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
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
// Union satisfied by
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyListResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyListResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyListResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyListResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
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

func (r accessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyListResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyListResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyListResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyListResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
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

func (r accessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessEmailRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessDomainRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessIPRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessIPListRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessCountryRule],
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyListResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyListResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyListResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyListResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseRequireAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyListResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
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

func (r accessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
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

func (r accessApplicationPolicyGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
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
// Union satisfied by
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule].
type AccessApplicationPolicyGetResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyGetResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyGetResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
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

func (r accessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessEmailRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessDomainRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessIPRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessIPListRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessCountryRule],
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule].
type AccessApplicationPolicyGetResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyGetResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyGetResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
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

func (r accessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessEmailRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessEmailListRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessDomainRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessEveryoneRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessIPRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessIPListRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessCertificateRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessCountryRule],
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule]
// or
// [zero_trust.AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule].
type AccessApplicationPolicyGetResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyGetResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyGetResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule{}),
		},
	)
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

func (r accessApplicationPolicyGetResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
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

func (r accessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
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
