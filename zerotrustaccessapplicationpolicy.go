// File generated from our OpenAPI spec by Stainless.

package cloudflare

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
func (r *ZeroTrustAccessApplicationPolicyService) New(ctx context.Context, uuid string, params ZeroTrustAccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationPolicyNewResponse, err error) {
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
func (r *ZeroTrustAccessApplicationPolicyService) Update(ctx context.Context, uuid1 string, uuid string, params ZeroTrustAccessApplicationPolicyUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationPolicyUpdateResponse, err error) {
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
func (r *ZeroTrustAccessApplicationPolicyService) List(ctx context.Context, uuid string, query ZeroTrustAccessApplicationPolicyListParams, opts ...option.RequestOption) (res *[]ZeroTrustAccessApplicationPolicyListResponse, err error) {
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
func (r *ZeroTrustAccessApplicationPolicyService) Get(ctx context.Context, uuid1 string, uuid string, query ZeroTrustAccessApplicationPolicyGetParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationPolicyGetResponse, err error) {
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

type ZeroTrustAccessApplicationPolicyNewResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ZeroTrustAccessApplicationPolicyNewResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision ZeroTrustAccessApplicationPolicyNewResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessApplicationPolicyNewResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessApplicationPolicyNewResponseInclude `json:"include"`
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
	Require []ZeroTrustAccessApplicationPolicyNewResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                          `json:"session_duration"`
	UpdatedAt       time.Time                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationPolicyNewResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessApplicationPolicyNewResponse]
type zeroTrustAccessApplicationPolicyNewResponseJSON struct {
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

func (r *ZeroTrustAccessApplicationPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

// A group of email addresses that can approve a temporary authentication request.
type ZeroTrustAccessApplicationPolicyNewResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                                                       `json:"email_list_uuid"`
	JSON          zeroTrustAccessApplicationPolicyNewResponseApprovalGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseApprovalGroupJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseApprovalGroup]
type zeroTrustAccessApplicationPolicyNewResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
}

// The action Access will take if a user matches this policy.
type ZeroTrustAccessApplicationPolicyNewResponseDecision string

const (
	ZeroTrustAccessApplicationPolicyNewResponseDecisionAllow       ZeroTrustAccessApplicationPolicyNewResponseDecision = "allow"
	ZeroTrustAccessApplicationPolicyNewResponseDecisionDeny        ZeroTrustAccessApplicationPolicyNewResponseDecision = "deny"
	ZeroTrustAccessApplicationPolicyNewResponseDecisionNonIdentity ZeroTrustAccessApplicationPolicyNewResponseDecision = "non_identity"
	ZeroTrustAccessApplicationPolicyNewResponseDecisionBypass      ZeroTrustAccessApplicationPolicyNewResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyNewResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyNewResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyNewResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                     `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                             `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                            `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                              `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                               `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                       `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                 `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                           `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                               `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                            `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                        `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                        `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                   `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                          `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                               `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                     `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                         `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                     `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyNewResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyNewResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyNewResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                     `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                             `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                            `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                              `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                               `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                       `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                 `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                           `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                               `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                            `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                        `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                        `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                   `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                          `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                               `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                     `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                         `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                     `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyNewResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyNewResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyNewResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                     `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                             `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                            `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                              `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                               `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                       `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                 `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                           `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                               `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                            `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                        `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                        `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                   `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                          `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                               `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                     `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                         `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyNewResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                     `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyNewResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationPolicyUpdateResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ZeroTrustAccessApplicationPolicyUpdateResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision ZeroTrustAccessApplicationPolicyUpdateResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessApplicationPolicyUpdateResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessApplicationPolicyUpdateResponseInclude `json:"include"`
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
	Require []ZeroTrustAccessApplicationPolicyUpdateResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                             `json:"session_duration"`
	UpdatedAt       time.Time                                          `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationPolicyUpdateResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationPolicyUpdateResponse]
type zeroTrustAccessApplicationPolicyUpdateResponseJSON struct {
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

func (r *ZeroTrustAccessApplicationPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A group of email addresses that can approve a temporary authentication request.
type ZeroTrustAccessApplicationPolicyUpdateResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                                                          `json:"email_list_uuid"`
	JSON          zeroTrustAccessApplicationPolicyUpdateResponseApprovalGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseApprovalGroupJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseApprovalGroup]
type zeroTrustAccessApplicationPolicyUpdateResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
}

// The action Access will take if a user matches this policy.
type ZeroTrustAccessApplicationPolicyUpdateResponseDecision string

const (
	ZeroTrustAccessApplicationPolicyUpdateResponseDecisionAllow       ZeroTrustAccessApplicationPolicyUpdateResponseDecision = "allow"
	ZeroTrustAccessApplicationPolicyUpdateResponseDecisionDeny        ZeroTrustAccessApplicationPolicyUpdateResponseDecision = "deny"
	ZeroTrustAccessApplicationPolicyUpdateResponseDecisionNonIdentity ZeroTrustAccessApplicationPolicyUpdateResponseDecision = "non_identity"
	ZeroTrustAccessApplicationPolicyUpdateResponseDecisionBypass      ZeroTrustAccessApplicationPolicyUpdateResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule]
// or
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyUpdateResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyUpdateResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                        `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                               `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                 `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                  `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                          `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                    `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                              `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                               `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                  `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                               `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                           `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                           `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                      `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                             `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                  `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                        `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                            `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                        `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule]
// or
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyUpdateResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyUpdateResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                        `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                               `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                 `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                  `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                          `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                    `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                              `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                               `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                  `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                               `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                           `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                           `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                      `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                             `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                  `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                        `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                            `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                        `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule]
// or
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyUpdateResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyUpdateResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                        `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                               `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                 `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                  `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                          `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                    `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                              `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                               `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                  `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                               `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                           `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                           `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                      `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                             `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                  `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                        `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                            `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyUpdateResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                        `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationPolicyListResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ZeroTrustAccessApplicationPolicyListResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision ZeroTrustAccessApplicationPolicyListResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessApplicationPolicyListResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessApplicationPolicyListResponseInclude `json:"include"`
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
	Require []ZeroTrustAccessApplicationPolicyListResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                           `json:"session_duration"`
	UpdatedAt       time.Time                                        `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationPolicyListResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessApplicationPolicyListResponse]
type zeroTrustAccessApplicationPolicyListResponseJSON struct {
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

func (r *ZeroTrustAccessApplicationPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// A group of email addresses that can approve a temporary authentication request.
type ZeroTrustAccessApplicationPolicyListResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                                                        `json:"email_list_uuid"`
	JSON          zeroTrustAccessApplicationPolicyListResponseApprovalGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseApprovalGroupJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseApprovalGroup]
type zeroTrustAccessApplicationPolicyListResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
}

// The action Access will take if a user matches this policy.
type ZeroTrustAccessApplicationPolicyListResponseDecision string

const (
	ZeroTrustAccessApplicationPolicyListResponseDecisionAllow       ZeroTrustAccessApplicationPolicyListResponseDecision = "allow"
	ZeroTrustAccessApplicationPolicyListResponseDecisionDeny        ZeroTrustAccessApplicationPolicyListResponseDecision = "deny"
	ZeroTrustAccessApplicationPolicyListResponseDecisionNonIdentity ZeroTrustAccessApplicationPolicyListResponseDecision = "non_identity"
	ZeroTrustAccessApplicationPolicyListResponseDecisionBypass      ZeroTrustAccessApplicationPolicyListResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyListResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyListResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyListResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                      `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                              `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                             `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                               `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                        `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                  `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                            `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                             `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                             `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                         `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                         `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                    `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                           `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                      `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                          `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyListResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                      `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyListResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyListResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyListResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                      `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                              `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                             `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                               `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                        `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                  `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                            `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                             `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                             `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                         `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                         `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                    `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                           `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                      `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                          `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyListResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                      `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyListResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyListResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyListResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                      `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                              `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                             `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                               `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                        `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                  `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                            `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                             `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                             `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                         `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                         `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                    `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                           `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                      `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                          `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyListResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                      `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyListResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationPolicyGetResponse struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ZeroTrustAccessApplicationPolicyGetResponseApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision ZeroTrustAccessApplicationPolicyGetResponseDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessApplicationPolicyGetResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessApplicationPolicyGetResponseInclude `json:"include"`
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
	Require []ZeroTrustAccessApplicationPolicyGetResponseRequire `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                          `json:"session_duration"`
	UpdatedAt       time.Time                                       `json:"updated_at" format:"date-time"`
	JSON            zeroTrustAccessApplicationPolicyGetResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessApplicationPolicyGetResponse]
type zeroTrustAccessApplicationPolicyGetResponseJSON struct {
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

func (r *ZeroTrustAccessApplicationPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

// A group of email addresses that can approve a temporary authentication request.
type ZeroTrustAccessApplicationPolicyGetResponseApprovalGroup struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []interface{} `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                                                       `json:"email_list_uuid"`
	JSON          zeroTrustAccessApplicationPolicyGetResponseApprovalGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseApprovalGroupJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseApprovalGroup]
type zeroTrustAccessApplicationPolicyGetResponseApprovalGroupJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseApprovalGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseApprovalGroupJSON) RawJSON() string {
	return r.raw
}

// The action Access will take if a user matches this policy.
type ZeroTrustAccessApplicationPolicyGetResponseDecision string

const (
	ZeroTrustAccessApplicationPolicyGetResponseDecisionAllow       ZeroTrustAccessApplicationPolicyGetResponseDecision = "allow"
	ZeroTrustAccessApplicationPolicyGetResponseDecisionDeny        ZeroTrustAccessApplicationPolicyGetResponseDecision = "deny"
	ZeroTrustAccessApplicationPolicyGetResponseDecisionNonIdentity ZeroTrustAccessApplicationPolicyGetResponseDecision = "non_identity"
	ZeroTrustAccessApplicationPolicyGetResponseDecisionBypass      ZeroTrustAccessApplicationPolicyGetResponseDecision = "bypass"
)

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyGetResponseExclude interface {
	implementsZeroTrustAccessApplicationPolicyGetResponseExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyGetResponseExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                     `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                             `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                            `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                              `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                               `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                       `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                 `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                           `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                               `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                            `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                        `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                        `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                   `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                          `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                               `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                     `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                         `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyGetResponseExclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                     `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyGetResponseInclude interface {
	implementsZeroTrustAccessApplicationPolicyGetResponseInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyGetResponseInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                     `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                             `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                            `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                              `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                               `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                       `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                 `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                           `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                               `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                            `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                        `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                        `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                   `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                          `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                               `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                     `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                         `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyGetResponseInclude() {
}

type ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                     `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRule],
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule]
// or [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessApplicationPolicyGetResponseRequire interface {
	implementsZeroTrustAccessApplicationPolicyGetResponseRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationPolicyGetResponseRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                     `json:"email,required" format:"email"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                             `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                            `json:"domain,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                              `json:"everyone,required"`
	JSON     zeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessEveryoneRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                               `json:"ip,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIP]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                       `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                 `json:"certificate,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCertificateRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                           `json:"id,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                            `json:"connection_id,required"`
	JSON         zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                               `json:"name,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                            `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                        `json:"email,required"`
	JSON  zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                        `json:"attribute_value,required"`
	JSON           zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                   `json:"token_id,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                          `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                               `json:"keys_url,required"`
	JSON    zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                     `json:"country_code,required"`
	JSON        zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                         `json:"auth_method,required"`
	JSON       zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRule]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessApplicationPolicyGetResponseRequire() {
}

type ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                     `json:"integration_uid,required"`
	JSON           zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationPolicyGetResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustAccessApplicationPolicyNewResponse                   `json:"result,required"`
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

func (r zeroTrustAccessApplicationPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustAccessApplicationPolicyUpdateResponse                   `json:"result,required"`
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

func (r zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   []ZeroTrustAccessApplicationPolicyListResponse                 `json:"result,required,nullable"`
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

func (r zeroTrustAccessApplicationPolicyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustAccessApplicationPolicyGetResponse                   `json:"result,required"`
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

func (r zeroTrustAccessApplicationPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessApplicationPolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessApplicationPolicyGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationPolicyGetResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationPolicyGetResponseEnvelopeSuccess = true
)
