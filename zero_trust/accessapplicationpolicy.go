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
func (r *AccessApplicationPolicyService) New(ctx context.Context, uuid string, params AccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
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
func (r *AccessApplicationPolicyService) Update(ctx context.Context, uuid1 string, uuid string, params AccessApplicationPolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
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
func (r *AccessApplicationPolicyService) List(ctx context.Context, uuid string, query AccessApplicationPolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Policy], err error) {
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
func (r *AccessApplicationPolicyService) ListAutoPaging(ctx context.Context, uuid string, query AccessApplicationPolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Policy] {
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
func (r *AccessApplicationPolicyService) Get(ctx context.Context, uuid1 string, uuid string, query AccessApplicationPolicyGetParams, opts ...option.RequestOption) (res *Policy, err error) {
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

// A group of email addresses that can approve a temporary authentication request.
type ApprovalGroupItem struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []string `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID string                `json:"email_list_uuid"`
	JSON          approvalGroupItemJSON `json:"-"`
}

// approvalGroupItemJSON contains the JSON metadata for the struct
// [ApprovalGroupItem]
type approvalGroupItemJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUUID   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ApprovalGroupItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalGroupItemJSON) RawJSON() string {
	return r.raw
}

// A group of email addresses that can approve a temporary authentication request.
type ApprovalGroupItemParam struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]string] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUUID param.Field[string] `json:"email_list_uuid"`
}

func (r ApprovalGroupItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
type ExcludeItem struct {
	Email                interface{}     `json:"email,required"`
	EmailList            interface{}     `json:"email_list,required"`
	EmailDomain          interface{}     `json:"email_domain,required"`
	Everyone             interface{}     `json:"everyone,required"`
	IP                   interface{}     `json:"ip,required"`
	IPList               interface{}     `json:"ip_list,required"`
	Certificate          interface{}     `json:"certificate,required"`
	Group                interface{}     `json:"group,required"`
	AzureAd              interface{}     `json:"azureAD,required"`
	GitHubOrganization   interface{}     `json:"github-organization,required"`
	Gsuite               interface{}     `json:"gsuite,required"`
	Okta                 interface{}     `json:"okta,required"`
	Saml                 interface{}     `json:"saml,required"`
	ServiceToken         interface{}     `json:"service_token,required"`
	AnyValidServiceToken interface{}     `json:"any_valid_service_token,required"`
	ExternalEvaluation   interface{}     `json:"external_evaluation,required"`
	Geo                  interface{}     `json:"geo,required"`
	AuthMethod           interface{}     `json:"auth_method,required"`
	DevicePosture        interface{}     `json:"device_posture,required"`
	JSON                 excludeItemJSON `json:"-"`
	union                ExcludeItemUnion
}

// excludeItemJSON contains the JSON metadata for the struct [ExcludeItem]
type excludeItemJSON struct {
	Email                apijson.Field
	EmailList            apijson.Field
	EmailDomain          apijson.Field
	Everyone             apijson.Field
	IP                   apijson.Field
	IPList               apijson.Field
	Certificate          apijson.Field
	Group                apijson.Field
	AzureAd              apijson.Field
	GitHubOrganization   apijson.Field
	Gsuite               apijson.Field
	Okta                 apijson.Field
	Saml                 apijson.Field
	ServiceToken         apijson.Field
	AnyValidServiceToken apijson.Field
	ExternalEvaluation   apijson.Field
	Geo                  apijson.Field
	AuthMethod           apijson.Field
	DevicePosture        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r excludeItemJSON) RawJSON() string {
	return r.raw
}

func (r *ExcludeItem) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ExcludeItem) AsUnion() ExcludeItemUnion {
	return r.union
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.EmailRule], [zero_trust.EmailListRule],
// [zero_trust.DomainRule], [zero_trust.EveryoneRule], [zero_trust.IPRule],
// [zero_trust.IPListRule], [zero_trust.CertificateRule], [zero_trust.GroupRule],
// [zero_trust.AzureGroupRule], [zero_trust.GitHubOrganizationRule],
// [zero_trust.GsuiteGroupRule], [zero_trust.OktaGroupRule],
// [zero_trust.SamlGroupRule], [zero_trust.ServiceTokenRule],
// [zero_trust.AnyValidServiceTokenRule], [zero_trust.ExternalEvaluationRule],
// [zero_trust.CountryRule], [zero_trust.AuthenticationMethodRule] or
// [zero_trust.DevicePostureRule].
type ExcludeItemUnion interface {
	implementsZeroTrustExcludeItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExcludeItemUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ExcludeItemParam struct {
	Email                param.Field[interface{}] `json:"email,required"`
	EmailList            param.Field[interface{}] `json:"email_list,required"`
	EmailDomain          param.Field[interface{}] `json:"email_domain,required"`
	Everyone             param.Field[interface{}] `json:"everyone,required"`
	IP                   param.Field[interface{}] `json:"ip,required"`
	IPList               param.Field[interface{}] `json:"ip_list,required"`
	Certificate          param.Field[interface{}] `json:"certificate,required"`
	Group                param.Field[interface{}] `json:"group,required"`
	AzureAd              param.Field[interface{}] `json:"azureAD,required"`
	GitHubOrganization   param.Field[interface{}] `json:"github-organization,required"`
	Gsuite               param.Field[interface{}] `json:"gsuite,required"`
	Okta                 param.Field[interface{}] `json:"okta,required"`
	Saml                 param.Field[interface{}] `json:"saml,required"`
	ServiceToken         param.Field[interface{}] `json:"service_token,required"`
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
	ExternalEvaluation   param.Field[interface{}] `json:"external_evaluation,required"`
	Geo                  param.Field[interface{}] `json:"geo,required"`
	AuthMethod           param.Field[interface{}] `json:"auth_method,required"`
	DevicePosture        param.Field[interface{}] `json:"device_posture,required"`
}

func (r ExcludeItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExcludeItemParam) implementsZeroTrustExcludeItemUnionParam() {}

// Matches a specific email.
//
// Satisfied by [zero_trust.EmailRuleParam], [zero_trust.EmailListRuleParam],
// [zero_trust.DomainRuleParam], [zero_trust.EveryoneRuleParam],
// [zero_trust.IPRuleParam], [zero_trust.IPListRuleParam],
// [zero_trust.CertificateRuleParam], [zero_trust.GroupRuleParam],
// [zero_trust.AzureGroupRuleParam], [zero_trust.GitHubOrganizationRuleParam],
// [zero_trust.GsuiteGroupRuleParam], [zero_trust.OktaGroupRuleParam],
// [zero_trust.SamlGroupRuleParam], [zero_trust.ServiceTokenRuleParam],
// [zero_trust.AnyValidServiceTokenRuleParam],
// [zero_trust.ExternalEvaluationRuleParam], [zero_trust.CountryRuleParam],
// [zero_trust.AuthenticationMethodRuleParam], [zero_trust.DevicePostureRuleParam],
// [ExcludeItemParam].
type ExcludeItemUnionParam interface {
	implementsZeroTrustExcludeItemUnionParam()
}

type Policy struct {
	// UUID
	ID string `json:"id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroupItem `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool      `json:"approval_required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy.
	Decision PolicyDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ExcludeItem `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []IncludeItem `json:"include"`
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
	Require []RequireItem `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string     `json:"session_duration"`
	UpdatedAt       time.Time  `json:"updated_at" format:"date-time"`
	JSON            policyJSON `json:"-"`
}

// policyJSON contains the JSON metadata for the struct [Policy]
type policyJSON struct {
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

func (r *Policy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyJSON) RawJSON() string {
	return r.raw
}

// The action Access will take if a user matches this policy.
type PolicyDecision string

const (
	PolicyDecisionAllow       PolicyDecision = "allow"
	PolicyDecisionDeny        PolicyDecision = "deny"
	PolicyDecisionNonIdentity PolicyDecision = "non_identity"
	PolicyDecisionBypass      PolicyDecision = "bypass"
)

func (r PolicyDecision) IsKnown() bool {
	switch r {
	case PolicyDecisionAllow, PolicyDecisionDeny, PolicyDecisionNonIdentity, PolicyDecisionBypass:
		return true
	}
	return false
}

type PolicyParam struct {
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupItemParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The action Access will take if a user matches this policy.
	Decision param.Field[PolicyDecision] `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ExcludeItemUnionParam] `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]IncludeItemUnionParam] `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// The order of execution for this policy. Must be unique for each policy.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]RequireItemUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r PolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
type RequireItem struct {
	Email                interface{}     `json:"email,required"`
	EmailList            interface{}     `json:"email_list,required"`
	EmailDomain          interface{}     `json:"email_domain,required"`
	Everyone             interface{}     `json:"everyone,required"`
	IP                   interface{}     `json:"ip,required"`
	IPList               interface{}     `json:"ip_list,required"`
	Certificate          interface{}     `json:"certificate,required"`
	Group                interface{}     `json:"group,required"`
	AzureAd              interface{}     `json:"azureAD,required"`
	GitHubOrganization   interface{}     `json:"github-organization,required"`
	Gsuite               interface{}     `json:"gsuite,required"`
	Okta                 interface{}     `json:"okta,required"`
	Saml                 interface{}     `json:"saml,required"`
	ServiceToken         interface{}     `json:"service_token,required"`
	AnyValidServiceToken interface{}     `json:"any_valid_service_token,required"`
	ExternalEvaluation   interface{}     `json:"external_evaluation,required"`
	Geo                  interface{}     `json:"geo,required"`
	AuthMethod           interface{}     `json:"auth_method,required"`
	DevicePosture        interface{}     `json:"device_posture,required"`
	JSON                 requireItemJSON `json:"-"`
	union                RequireItemUnion
}

// requireItemJSON contains the JSON metadata for the struct [RequireItem]
type requireItemJSON struct {
	Email                apijson.Field
	EmailList            apijson.Field
	EmailDomain          apijson.Field
	Everyone             apijson.Field
	IP                   apijson.Field
	IPList               apijson.Field
	Certificate          apijson.Field
	Group                apijson.Field
	AzureAd              apijson.Field
	GitHubOrganization   apijson.Field
	Gsuite               apijson.Field
	Okta                 apijson.Field
	Saml                 apijson.Field
	ServiceToken         apijson.Field
	AnyValidServiceToken apijson.Field
	ExternalEvaluation   apijson.Field
	Geo                  apijson.Field
	AuthMethod           apijson.Field
	DevicePosture        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r requireItemJSON) RawJSON() string {
	return r.raw
}

func (r *RequireItem) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RequireItem) AsUnion() RequireItemUnion {
	return r.union
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.EmailRule], [zero_trust.EmailListRule],
// [zero_trust.DomainRule], [zero_trust.EveryoneRule], [zero_trust.IPRule],
// [zero_trust.IPListRule], [zero_trust.CertificateRule], [zero_trust.GroupRule],
// [zero_trust.AzureGroupRule], [zero_trust.GitHubOrganizationRule],
// [zero_trust.GsuiteGroupRule], [zero_trust.OktaGroupRule],
// [zero_trust.SamlGroupRule], [zero_trust.ServiceTokenRule],
// [zero_trust.AnyValidServiceTokenRule], [zero_trust.ExternalEvaluationRule],
// [zero_trust.CountryRule], [zero_trust.AuthenticationMethodRule] or
// [zero_trust.DevicePostureRule].
type RequireItemUnion interface {
	implementsZeroTrustRequireItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RequireItemUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type RequireItemParam struct {
	Email                param.Field[interface{}] `json:"email,required"`
	EmailList            param.Field[interface{}] `json:"email_list,required"`
	EmailDomain          param.Field[interface{}] `json:"email_domain,required"`
	Everyone             param.Field[interface{}] `json:"everyone,required"`
	IP                   param.Field[interface{}] `json:"ip,required"`
	IPList               param.Field[interface{}] `json:"ip_list,required"`
	Certificate          param.Field[interface{}] `json:"certificate,required"`
	Group                param.Field[interface{}] `json:"group,required"`
	AzureAd              param.Field[interface{}] `json:"azureAD,required"`
	GitHubOrganization   param.Field[interface{}] `json:"github-organization,required"`
	Gsuite               param.Field[interface{}] `json:"gsuite,required"`
	Okta                 param.Field[interface{}] `json:"okta,required"`
	Saml                 param.Field[interface{}] `json:"saml,required"`
	ServiceToken         param.Field[interface{}] `json:"service_token,required"`
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
	ExternalEvaluation   param.Field[interface{}] `json:"external_evaluation,required"`
	Geo                  param.Field[interface{}] `json:"geo,required"`
	AuthMethod           param.Field[interface{}] `json:"auth_method,required"`
	DevicePosture        param.Field[interface{}] `json:"device_posture,required"`
}

func (r RequireItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RequireItemParam) implementsZeroTrustRequireItemUnionParam() {}

// Matches a specific email.
//
// Satisfied by [zero_trust.EmailRuleParam], [zero_trust.EmailListRuleParam],
// [zero_trust.DomainRuleParam], [zero_trust.EveryoneRuleParam],
// [zero_trust.IPRuleParam], [zero_trust.IPListRuleParam],
// [zero_trust.CertificateRuleParam], [zero_trust.GroupRuleParam],
// [zero_trust.AzureGroupRuleParam], [zero_trust.GitHubOrganizationRuleParam],
// [zero_trust.GsuiteGroupRuleParam], [zero_trust.OktaGroupRuleParam],
// [zero_trust.SamlGroupRuleParam], [zero_trust.ServiceTokenRuleParam],
// [zero_trust.AnyValidServiceTokenRuleParam],
// [zero_trust.ExternalEvaluationRuleParam], [zero_trust.CountryRuleParam],
// [zero_trust.AuthenticationMethodRuleParam], [zero_trust.DevicePostureRuleParam],
// [RequireItemParam].
type RequireItemUnionParam interface {
	implementsZeroTrustRequireItemUnionParam()
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
	Include param.Field[[]IncludeItemUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupItemParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ExcludeItemUnionParam] `json:"exclude"`
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
	Require param.Field[[]RequireItemUnionParam] `json:"require"`
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

type AccessApplicationPolicyNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Policy                                                    `json:"result,required"`
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
	Include param.Field[[]IncludeItemUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupItemParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ExcludeItemUnionParam] `json:"exclude"`
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
	Require param.Field[[]RequireItemUnionParam] `json:"require"`
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

type AccessApplicationPolicyUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Policy                                                    `json:"result,required"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   AccessApplicationPolicyDeleteResponse                     `json:"result,required"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Policy                                                    `json:"result,required"`
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
