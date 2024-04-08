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

// AccessGroupService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessGroupService] method
// instead.
type AccessGroupService struct {
	Options []option.RequestOption
}

// NewAccessGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessGroupService(opts ...option.RequestOption) (r *AccessGroupService) {
	r = &AccessGroupService{}
	r.Options = opts
	return
}

// Creates a new Access group.
func (r *AccessGroupService) New(ctx context.Context, params AccessGroupNewParams, opts ...option.RequestOption) (res *ZeroTrustGroup, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Access group.
func (r *AccessGroupService) Update(ctx context.Context, uuid string, params AccessGroupUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGroup, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access groups.
func (r *AccessGroupService) List(ctx context.Context, query AccessGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustGroup], err error) {
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
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
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

// Lists all Access groups.
func (r *AccessGroupService) ListAutoPaging(ctx context.Context, query AccessGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustGroup] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an Access group.
func (r *AccessGroupService) Delete(ctx context.Context, uuid string, body AccessGroupDeleteParams, opts ...option.RequestOption) (res *AccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access group.
func (r *AccessGroupService) Get(ctx context.Context, uuid string, query AccessGroupGetParams, opts ...option.RequestOption) (res *ZeroTrustGroup, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type ZeroTrustGroup struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ExcludeItem `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []IncludeItem `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []RequireItem `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []RequireItem      `json:"require"`
	UpdatedAt time.Time          `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGroupJSON `json:"-"`
}

// zeroTrustGroupJSON contains the JSON metadata for the struct [ZeroTrustGroup]
type zeroTrustGroupJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupJSON) RawJSON() string {
	return r.raw
}

type AccessGroupDeleteResponse struct {
	// UUID
	ID   string                        `json:"id"`
	JSON accessGroupDeleteResponseJSON `json:"-"`
}

// accessGroupDeleteResponseJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponse]
type accessGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]IncludeItemUnionParam] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ExcludeItemUnionParam] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]RequireItemUnionParam] `json:"require"`
}

func (r AccessGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ZeroTrustGroup                                            `json:"result,required"`
	// Whether the API call was successful
	Success AccessGroupNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessGroupNewResponseEnvelopeJSON    `json:"-"`
}

// accessGroupNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupNewResponseEnvelope]
type accessGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupNewResponseEnvelopeSuccess bool

const (
	AccessGroupNewResponseEnvelopeSuccessTrue AccessGroupNewResponseEnvelopeSuccess = true
)

func (r AccessGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]IncludeItemUnionParam] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ExcludeItemUnionParam] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]RequireItemUnionParam] `json:"require"`
}

func (r AccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ZeroTrustGroup                                            `json:"result,required"`
	// Whether the API call was successful
	Success AccessGroupUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessGroupUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessGroupUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponseEnvelope]
type accessGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupUpdateResponseEnvelopeSuccess bool

const (
	AccessGroupUpdateResponseEnvelopeSuccessTrue AccessGroupUpdateResponseEnvelopeSuccess = true
)

func (r AccessGroupUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   AccessGroupDeleteResponse                                 `json:"result,required"`
	// Whether the API call was successful
	Success AccessGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponseEnvelope]
type accessGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupDeleteResponseEnvelopeSuccess bool

const (
	AccessGroupDeleteResponseEnvelopeSuccessTrue AccessGroupDeleteResponseEnvelopeSuccess = true
)

func (r AccessGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ZeroTrustGroup                                            `json:"result,required"`
	// Whether the API call was successful
	Success AccessGroupGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessGroupGetResponseEnvelopeJSON    `json:"-"`
}

// accessGroupGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupGetResponseEnvelope]
type accessGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupGetResponseEnvelopeSuccess bool

const (
	AccessGroupGetResponseEnvelopeSuccessTrue AccessGroupGetResponseEnvelopeSuccess = true
)

func (r AccessGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
