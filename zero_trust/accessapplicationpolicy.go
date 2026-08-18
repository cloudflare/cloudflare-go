// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// AccessApplicationPolicyService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationPolicyService] method instead.
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

// Creates a policy applying exclusive to a single application that defines the
// users or groups who can reach it. We recommend creating a reusable policy
// instead and subsequently referencing its ID in the application's 'policies'
// array.
func (r *AccessApplicationPolicyService) New(ctx context.Context, appID string, params AccessApplicationPolicyNewParams, opts ...option.RequestOption) (res *AccessApplicationPolicyNewResponse, err error) {
	var env AccessApplicationPolicyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an Access policy specific to an application. To update a reusable
// policy, use the /account or zones/{account or zone_id}/policies/{uid} endpoint.
func (r *AccessApplicationPolicyService) Update(ctx context.Context, appID string, policyID string, params AccessApplicationPolicyUpdateParams, opts ...option.RequestOption) (res *AccessApplicationPolicyUpdateResponse, err error) {
	var env AccessApplicationPolicyUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, appID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists Access policies configured for an application. Returns both exclusively
// scoped and reusable policies used by the application.
func (r *AccessApplicationPolicyService) List(ctx context.Context, appID string, params AccessApplicationPolicyListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessApplicationPolicyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies", accountOrZone, accountOrZoneID, appID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists Access policies configured for an application. Returns both exclusively
// scoped and reusable policies used by the application.
func (r *AccessApplicationPolicyService) ListAutoPaging(ctx context.Context, appID string, params AccessApplicationPolicyListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessApplicationPolicyListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, appID, params, opts...))
}

// Deletes an Access policy specific to an application. To delete a reusable
// policy, use the /account or zones/{account or zone_id}/policies/{uid} endpoint.
func (r *AccessApplicationPolicyService) Delete(ctx context.Context, appID string, policyID string, body AccessApplicationPolicyDeleteParams, opts ...option.RequestOption) (res *AccessApplicationPolicyDeleteResponse, err error) {
	var env AccessApplicationPolicyDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, appID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches a single Access policy configured for an application. Returns both
// exclusively owned and reusable policies used by the application.
func (r *AccessApplicationPolicyService) Get(ctx context.Context, appID string, policyID string, query AccessApplicationPolicyGetParams, opts ...option.RequestOption) (res *AccessApplicationPolicyGetResponse, err error) {
	var env AccessApplicationPolicyGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/policies/%s", accountOrZone, accountOrZoneID, appID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Enforces a device posture rule has run successfully
type AccessDevicePostureRule struct {
	DevicePosture AccessDevicePostureRuleDevicePosture `json:"device_posture" api:"required"`
	JSON          accessDevicePostureRuleJSON          `json:"-"`
}

// accessDevicePostureRuleJSON contains the JSON metadata for the struct
// [AccessDevicePostureRule]
type accessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessDevicePostureRule) implementsAccessRule() {}

type AccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUID string                                   `json:"integration_uid" api:"required"`
	JSON           accessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessDevicePostureRuleDevicePostureJSON contains the JSON metadata for the
// struct [AccessDevicePostureRuleDevicePosture]
type accessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type AccessDevicePostureRuleParam struct {
	DevicePosture param.Field[AccessDevicePostureRuleDevicePostureParam] `json:"device_posture" api:"required"`
}

func (r AccessDevicePostureRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessDevicePostureRuleParam) implementsAccessRuleUnionParam() {}

type AccessDevicePostureRuleDevicePostureParam struct {
	// The ID of a device posture integration.
	IntegrationUID param.Field[string] `json:"integration_uid" api:"required"`
}

func (r AccessDevicePostureRuleDevicePostureParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Access group.
type AccessRule struct {
	// This field can have the runtime type of
	// [AnyValidServiceTokenRuleAnyValidServiceToken].
	AnyValidServiceToken interface{} `json:"any_valid_service_token"`
	// This field can have the runtime type of
	// [AccessRuleAccessAuthContextRuleAuthContext].
	AuthContext interface{} `json:"auth_context"`
	// This field can have the runtime type of [AuthenticationMethodRuleAuthMethod].
	AuthMethod interface{} `json:"auth_method"`
	// This field can have the runtime type of [AzureGroupRuleAzureAD].
	AzureAD interface{} `json:"azureAD"`
	// This field can have the runtime type of [CertificateRuleCertificate].
	Certificate interface{} `json:"certificate"`
	// This field can have the runtime type of
	// [AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMember].
	CloudflareAccountMember interface{} `json:"cloudflare_account_member"`
	// This field can have the runtime type of
	// [AccessRuleAccessCommonNameRuleCommonName].
	CommonName interface{} `json:"common_name"`
	// This field can have the runtime type of [AccessDevicePostureRuleDevicePosture].
	DevicePosture interface{} `json:"device_posture"`
	// This field can have the runtime type of [EmailRuleEmail].
	Email interface{} `json:"email"`
	// This field can have the runtime type of [DomainRuleEmailDomain].
	EmailDomain interface{} `json:"email_domain"`
	// This field can have the runtime type of [EmailListRuleEmailList].
	EmailList interface{} `json:"email_list"`
	// This field can have the runtime type of [EveryoneRuleEveryone].
	Everyone interface{} `json:"everyone"`
	// This field can have the runtime type of
	// [ExternalEvaluationRuleExternalEvaluation].
	ExternalEvaluation interface{} `json:"external_evaluation"`
	// This field can have the runtime type of [CountryRuleGeo].
	Geo interface{} `json:"geo"`
	// This field can have the runtime type of
	// [GitHubOrganizationRuleGitHubOrganization].
	GitHubOrganization interface{} `json:"github-organization"`
	// This field can have the runtime type of [GroupRuleGroup].
	Group interface{} `json:"group"`
	// This field can have the runtime type of [GSuiteGroupRuleGSuite].
	GSuite interface{} `json:"gsuite"`
	// This field can have the runtime type of [IPRuleIP].
	IP interface{} `json:"ip"`
	// This field can have the runtime type of [IPListRuleIPList].
	IPList interface{} `json:"ip_list"`
	// This field can have the runtime type of
	// [AccessRuleAccessLinkedAppTokenRuleLinkedAppToken].
	LinkedAppToken interface{} `json:"linked_app_token"`
	// This field can have the runtime type of
	// [AccessRuleAccessLoginMethodRuleLoginMethod].
	LoginMethod interface{} `json:"login_method"`
	// This field can have the runtime type of [AccessRuleAccessOIDCClaimRuleOIDC].
	OIDC interface{} `json:"oidc"`
	// This field can have the runtime type of [OktaGroupRuleOkta].
	Okta interface{} `json:"okta"`
	// This field can have the runtime type of [SAMLGroupRuleSAML].
	SAML interface{} `json:"saml"`
	// This field can have the runtime type of [ServiceTokenRuleServiceToken].
	ServiceToken interface{} `json:"service_token"`
	// This field can have the runtime type of
	// [AccessRuleAccessUserRiskScoreRuleUserRiskScore].
	UserRiskScore interface{}    `json:"user_risk_score"`
	JSON          accessRuleJSON `json:"-"`
	union         AccessRuleUnion
}

// accessRuleJSON contains the JSON metadata for the struct [AccessRule]
type accessRuleJSON struct {
	AnyValidServiceToken    apijson.Field
	AuthContext             apijson.Field
	AuthMethod              apijson.Field
	AzureAD                 apijson.Field
	Certificate             apijson.Field
	CloudflareAccountMember apijson.Field
	CommonName              apijson.Field
	DevicePosture           apijson.Field
	Email                   apijson.Field
	EmailDomain             apijson.Field
	EmailList               apijson.Field
	Everyone                apijson.Field
	ExternalEvaluation      apijson.Field
	Geo                     apijson.Field
	GitHubOrganization      apijson.Field
	Group                   apijson.Field
	GSuite                  apijson.Field
	IP                      apijson.Field
	IPList                  apijson.Field
	LinkedAppToken          apijson.Field
	LoginMethod             apijson.Field
	OIDC                    apijson.Field
	Okta                    apijson.Field
	SAML                    apijson.Field
	ServiceToken            apijson.Field
	UserRiskScore           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r accessRuleJSON) RawJSON() string {
	return r.raw
}

func (r *AccessRule) UnmarshalJSON(data []byte) (err error) {
	*r = AccessRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessRuleUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [GroupRule], [AnyValidServiceTokenRule],
// [AccessRuleAccessAuthContextRule], [AuthenticationMethodRule], [AzureGroupRule],
// [CertificateRule], [AccessRuleAccessCommonNameRule], [CountryRule],
// [AccessDevicePostureRule], [DomainRule], [EmailListRule], [EmailRule],
// [EveryoneRule], [ExternalEvaluationRule], [GitHubOrganizationRule],
// [GSuiteGroupRule], [AccessRuleAccessLoginMethodRule], [IPListRule], [IPRule],
// [OktaGroupRule], [SAMLGroupRule], [AccessRuleAccessOIDCClaimRule],
// [ServiceTokenRule], [AccessRuleAccessLinkedAppTokenRule],
// [AccessRuleAccessUserRiskScoreRule],
// [AccessRuleAccessCloudflareAccountMemberRule].
func (r AccessRule) AsUnion() AccessRuleUnion {
	return r.union
}

// Matches an Access group.
//
// Union satisfied by [GroupRule], [AnyValidServiceTokenRule],
// [AccessRuleAccessAuthContextRule], [AuthenticationMethodRule], [AzureGroupRule],
// [CertificateRule], [AccessRuleAccessCommonNameRule], [CountryRule],
// [AccessDevicePostureRule], [DomainRule], [EmailListRule], [EmailRule],
// [EveryoneRule], [ExternalEvaluationRule], [GitHubOrganizationRule],
// [GSuiteGroupRule], [AccessRuleAccessLoginMethodRule], [IPListRule], [IPRule],
// [OktaGroupRule], [SAMLGroupRule], [AccessRuleAccessOIDCClaimRule],
// [ServiceTokenRule], [AccessRuleAccessLinkedAppTokenRule],
// [AccessRuleAccessUserRiskScoreRule] or
// [AccessRuleAccessCloudflareAccountMemberRule].
type AccessRuleUnion interface {
	implementsAccessRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessRuleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAuthContextRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessCommonNameRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessDevicePostureRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GSuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessLoginMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SAMLGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessOIDCClaimRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessLinkedAppTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessUserRiskScoreRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessCloudflareAccountMemberRule{}),
		},
	)
}

// Matches an Azure Authentication Context. Requires an Azure identity provider.
type AccessRuleAccessAuthContextRule struct {
	AuthContext AccessRuleAccessAuthContextRuleAuthContext `json:"auth_context" api:"required"`
	JSON        accessRuleAccessAuthContextRuleJSON        `json:"-"`
}

// accessRuleAccessAuthContextRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessAuthContextRule]
type accessRuleAccessAuthContextRuleJSON struct {
	AuthContext apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAuthContextRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthContextRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAuthContextRule) implementsAccessRule() {}

type AccessRuleAccessAuthContextRuleAuthContext struct {
	// The ID of an Authentication context.
	ID string `json:"id" api:"required"`
	// The ACID of an Authentication context.
	AcID string `json:"ac_id" api:"required"`
	// The ID of your Azure identity provider.
	IdentityProviderID string                                         `json:"identity_provider_id" api:"required"`
	JSON               accessRuleAccessAuthContextRuleAuthContextJSON `json:"-"`
}

// accessRuleAccessAuthContextRuleAuthContextJSON contains the JSON metadata for
// the struct [AccessRuleAccessAuthContextRuleAuthContext]
type accessRuleAccessAuthContextRuleAuthContextJSON struct {
	ID                 apijson.Field
	AcID               apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessAuthContextRuleAuthContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthContextRuleAuthContextJSON) RawJSON() string {
	return r.raw
}

// Matches a specific common name.
type AccessRuleAccessCommonNameRule struct {
	CommonName AccessRuleAccessCommonNameRuleCommonName `json:"common_name" api:"required"`
	JSON       accessRuleAccessCommonNameRuleJSON       `json:"-"`
}

// accessRuleAccessCommonNameRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessCommonNameRule]
type accessRuleAccessCommonNameRuleJSON struct {
	CommonName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCommonNameRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCommonNameRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessCommonNameRule) implementsAccessRule() {}

type AccessRuleAccessCommonNameRuleCommonName struct {
	// The common name to match.
	CommonName string                                       `json:"common_name" api:"required"`
	JSON       accessRuleAccessCommonNameRuleCommonNameJSON `json:"-"`
}

// accessRuleAccessCommonNameRuleCommonNameJSON contains the JSON metadata for the
// struct [AccessRuleAccessCommonNameRuleCommonName]
type accessRuleAccessCommonNameRuleCommonNameJSON struct {
	CommonName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCommonNameRuleCommonName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCommonNameRuleCommonNameJSON) RawJSON() string {
	return r.raw
}

// Matches a specific identity provider id.
type AccessRuleAccessLoginMethodRule struct {
	LoginMethod AccessRuleAccessLoginMethodRuleLoginMethod `json:"login_method" api:"required"`
	JSON        accessRuleAccessLoginMethodRuleJSON        `json:"-"`
}

// accessRuleAccessLoginMethodRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessLoginMethodRule]
type accessRuleAccessLoginMethodRuleJSON struct {
	LoginMethod apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessLoginMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLoginMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessLoginMethodRule) implementsAccessRule() {}

type AccessRuleAccessLoginMethodRuleLoginMethod struct {
	// The ID of an identity provider.
	ID   string                                         `json:"id" api:"required"`
	JSON accessRuleAccessLoginMethodRuleLoginMethodJSON `json:"-"`
}

// accessRuleAccessLoginMethodRuleLoginMethodJSON contains the JSON metadata for
// the struct [AccessRuleAccessLoginMethodRuleLoginMethod]
type accessRuleAccessLoginMethodRuleLoginMethodJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessLoginMethodRuleLoginMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLoginMethodRuleLoginMethodJSON) RawJSON() string {
	return r.raw
}

// Matches an OIDC claim. Requires an OIDC identity provider.
type AccessRuleAccessOIDCClaimRule struct {
	OIDC AccessRuleAccessOIDCClaimRuleOIDC `json:"oidc" api:"required"`
	JSON accessRuleAccessOIDCClaimRuleJSON `json:"-"`
}

// accessRuleAccessOIDCClaimRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessOIDCClaimRule]
type accessRuleAccessOIDCClaimRuleJSON struct {
	OIDC        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessOIDCClaimRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessOIDCClaimRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessOIDCClaimRule) implementsAccessRule() {}

type AccessRuleAccessOIDCClaimRuleOIDC struct {
	// The name of the OIDC claim.
	ClaimName string `json:"claim_name" api:"required"`
	// The OIDC claim value to look for.
	ClaimValue string `json:"claim_value" api:"required"`
	// The ID of your OIDC identity provider.
	IdentityProviderID string                                `json:"identity_provider_id" api:"required"`
	JSON               accessRuleAccessOIDCClaimRuleOIDCJSON `json:"-"`
}

// accessRuleAccessOIDCClaimRuleOIDCJSON contains the JSON metadata for the struct
// [AccessRuleAccessOIDCClaimRuleOIDC]
type accessRuleAccessOIDCClaimRuleOIDCJSON struct {
	ClaimName          apijson.Field
	ClaimValue         apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessOIDCClaimRuleOIDC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessOIDCClaimRuleOIDCJSON) RawJSON() string {
	return r.raw
}

// Matches OAuth 2.0 access tokens issued by the specified Access OIDC SaaS
// application. Only compatible with non_identity and bypass decisions.
type AccessRuleAccessLinkedAppTokenRule struct {
	LinkedAppToken AccessRuleAccessLinkedAppTokenRuleLinkedAppToken `json:"linked_app_token" api:"required"`
	JSON           accessRuleAccessLinkedAppTokenRuleJSON           `json:"-"`
}

// accessRuleAccessLinkedAppTokenRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessLinkedAppTokenRule]
type accessRuleAccessLinkedAppTokenRuleJSON struct {
	LinkedAppToken apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessRuleAccessLinkedAppTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLinkedAppTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessLinkedAppTokenRule) implementsAccessRule() {}

type AccessRuleAccessLinkedAppTokenRuleLinkedAppToken struct {
	// The ID of an Access OIDC SaaS application
	AppUID string                                               `json:"app_uid" api:"required"`
	JSON   accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON `json:"-"`
}

// accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON contains the JSON metadata
// for the struct [AccessRuleAccessLinkedAppTokenRuleLinkedAppToken]
type accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON struct {
	AppUID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessLinkedAppTokenRuleLinkedAppToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessLinkedAppTokenRuleLinkedAppTokenJSON) RawJSON() string {
	return r.raw
}

// Matches a user's risk score.
type AccessRuleAccessUserRiskScoreRule struct {
	UserRiskScore AccessRuleAccessUserRiskScoreRuleUserRiskScore `json:"user_risk_score" api:"required"`
	JSON          accessRuleAccessUserRiskScoreRuleJSON          `json:"-"`
}

// accessRuleAccessUserRiskScoreRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessUserRiskScoreRule]
type accessRuleAccessUserRiskScoreRuleJSON struct {
	UserRiskScore apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleAccessUserRiskScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessUserRiskScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessUserRiskScoreRule) implementsAccessRule() {}

type AccessRuleAccessUserRiskScoreRuleUserRiskScore struct {
	// A list of risk score levels to match. Values can be low, medium, high, or
	// unscored.
	UserRiskScore []AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore `json:"user_risk_score" api:"required"`
	JSON          accessRuleAccessUserRiskScoreRuleUserRiskScoreJSON            `json:"-"`
}

// accessRuleAccessUserRiskScoreRuleUserRiskScoreJSON contains the JSON metadata
// for the struct [AccessRuleAccessUserRiskScoreRuleUserRiskScore]
type accessRuleAccessUserRiskScoreRuleUserRiskScoreJSON struct {
	UserRiskScore apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleAccessUserRiskScoreRuleUserRiskScore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessUserRiskScoreRuleUserRiskScoreJSON) RawJSON() string {
	return r.raw
}

type AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore string

const (
	AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreLow      AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore = "low"
	AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreMedium   AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore = "medium"
	AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreHigh     AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore = "high"
	AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreUnscored AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore = "unscored"
)

func (r AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore) IsKnown() bool {
	switch r {
	case AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreLow, AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreMedium, AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreHigh, AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScoreUnscored:
		return true
	}
	return false
}

// Matches users who are members of a specific Cloudflare account. Requires a
// Cloudflare identity provider.
type AccessRuleAccessCloudflareAccountMemberRule struct {
	CloudflareAccountMember AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMember `json:"cloudflare_account_member" api:"required"`
	JSON                    accessRuleAccessCloudflareAccountMemberRuleJSON                    `json:"-"`
}

// accessRuleAccessCloudflareAccountMemberRuleJSON contains the JSON metadata for
// the struct [AccessRuleAccessCloudflareAccountMemberRule]
type accessRuleAccessCloudflareAccountMemberRuleJSON struct {
	CloudflareAccountMember apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccessRuleAccessCloudflareAccountMemberRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCloudflareAccountMemberRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessCloudflareAccountMemberRule) implementsAccessRule() {}

type AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMember struct {
	// Identifier.
	AccountID string                                                                 `json:"account_id"`
	JSON      accessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMemberJSON `json:"-"`
}

// accessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMemberJSON contains
// the JSON metadata for the struct
// [AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMember]
type accessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMemberJSON struct {
	AccountID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMemberJSON) RawJSON() string {
	return r.raw
}

// Matches an Access group.
type AccessRuleParam struct {
	AnyValidServiceToken    param.Field[interface{}] `json:"any_valid_service_token"`
	AuthContext             param.Field[interface{}] `json:"auth_context"`
	AuthMethod              param.Field[interface{}] `json:"auth_method"`
	AzureAD                 param.Field[interface{}] `json:"azureAD"`
	Certificate             param.Field[interface{}] `json:"certificate"`
	CloudflareAccountMember param.Field[interface{}] `json:"cloudflare_account_member"`
	CommonName              param.Field[interface{}] `json:"common_name"`
	DevicePosture           param.Field[interface{}] `json:"device_posture"`
	Email                   param.Field[interface{}] `json:"email"`
	EmailDomain             param.Field[interface{}] `json:"email_domain"`
	EmailList               param.Field[interface{}] `json:"email_list"`
	Everyone                param.Field[interface{}] `json:"everyone"`
	ExternalEvaluation      param.Field[interface{}] `json:"external_evaluation"`
	Geo                     param.Field[interface{}] `json:"geo"`
	GitHubOrganization      param.Field[interface{}] `json:"github-organization"`
	Group                   param.Field[interface{}] `json:"group"`
	GSuite                  param.Field[interface{}] `json:"gsuite"`
	IP                      param.Field[interface{}] `json:"ip"`
	IPList                  param.Field[interface{}] `json:"ip_list"`
	LinkedAppToken          param.Field[interface{}] `json:"linked_app_token"`
	LoginMethod             param.Field[interface{}] `json:"login_method"`
	OIDC                    param.Field[interface{}] `json:"oidc"`
	Okta                    param.Field[interface{}] `json:"okta"`
	SAML                    param.Field[interface{}] `json:"saml"`
	ServiceToken            param.Field[interface{}] `json:"service_token"`
	UserRiskScore           param.Field[interface{}] `json:"user_risk_score"`
}

func (r AccessRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleParam) implementsAccessRuleUnionParam() {}

// Matches an Access group.
//
// Satisfied by [zero_trust.GroupRuleParam],
// [zero_trust.AnyValidServiceTokenRuleParam],
// [zero_trust.AccessRuleAccessAuthContextRuleParam],
// [zero_trust.AuthenticationMethodRuleParam], [zero_trust.AzureGroupRuleParam],
// [zero_trust.CertificateRuleParam],
// [zero_trust.AccessRuleAccessCommonNameRuleParam], [zero_trust.CountryRuleParam],
// [zero_trust.AccessDevicePostureRuleParam], [zero_trust.DomainRuleParam],
// [zero_trust.EmailListRuleParam], [zero_trust.EmailRuleParam],
// [zero_trust.EveryoneRuleParam], [zero_trust.ExternalEvaluationRuleParam],
// [zero_trust.GitHubOrganizationRuleParam], [zero_trust.GSuiteGroupRuleParam],
// [zero_trust.AccessRuleAccessLoginMethodRuleParam], [zero_trust.IPListRuleParam],
// [zero_trust.IPRuleParam], [zero_trust.OktaGroupRuleParam],
// [zero_trust.SAMLGroupRuleParam],
// [zero_trust.AccessRuleAccessOIDCClaimRuleParam],
// [zero_trust.ServiceTokenRuleParam],
// [zero_trust.AccessRuleAccessLinkedAppTokenRuleParam],
// [zero_trust.AccessRuleAccessUserRiskScoreRuleParam],
// [zero_trust.AccessRuleAccessCloudflareAccountMemberRuleParam],
// [AccessRuleParam].
type AccessRuleUnionParam interface {
	implementsAccessRuleUnionParam()
}

// Matches an Azure Authentication Context. Requires an Azure identity provider.
type AccessRuleAccessAuthContextRuleParam struct {
	AuthContext param.Field[AccessRuleAccessAuthContextRuleAuthContextParam] `json:"auth_context" api:"required"`
}

func (r AccessRuleAccessAuthContextRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAuthContextRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessAuthContextRuleAuthContextParam struct {
	// The ID of an Authentication context.
	ID param.Field[string] `json:"id" api:"required"`
	// The ACID of an Authentication context.
	AcID param.Field[string] `json:"ac_id" api:"required"`
	// The ID of your Azure identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id" api:"required"`
}

func (r AccessRuleAccessAuthContextRuleAuthContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific common name.
type AccessRuleAccessCommonNameRuleParam struct {
	CommonName param.Field[AccessRuleAccessCommonNameRuleCommonNameParam] `json:"common_name" api:"required"`
}

func (r AccessRuleAccessCommonNameRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCommonNameRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessCommonNameRuleCommonNameParam struct {
	// The common name to match.
	CommonName param.Field[string] `json:"common_name" api:"required"`
}

func (r AccessRuleAccessCommonNameRuleCommonNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific identity provider id.
type AccessRuleAccessLoginMethodRuleParam struct {
	LoginMethod param.Field[AccessRuleAccessLoginMethodRuleLoginMethodParam] `json:"login_method" api:"required"`
}

func (r AccessRuleAccessLoginMethodRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessLoginMethodRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessLoginMethodRuleLoginMethodParam struct {
	// The ID of an identity provider.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r AccessRuleAccessLoginMethodRuleLoginMethodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an OIDC claim. Requires an OIDC identity provider.
type AccessRuleAccessOIDCClaimRuleParam struct {
	OIDC param.Field[AccessRuleAccessOIDCClaimRuleOIDCParam] `json:"oidc" api:"required"`
}

func (r AccessRuleAccessOIDCClaimRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessOIDCClaimRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessOIDCClaimRuleOIDCParam struct {
	// The name of the OIDC claim.
	ClaimName param.Field[string] `json:"claim_name" api:"required"`
	// The OIDC claim value to look for.
	ClaimValue param.Field[string] `json:"claim_value" api:"required"`
	// The ID of your OIDC identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id" api:"required"`
}

func (r AccessRuleAccessOIDCClaimRuleOIDCParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches OAuth 2.0 access tokens issued by the specified Access OIDC SaaS
// application. Only compatible with non_identity and bypass decisions.
type AccessRuleAccessLinkedAppTokenRuleParam struct {
	LinkedAppToken param.Field[AccessRuleAccessLinkedAppTokenRuleLinkedAppTokenParam] `json:"linked_app_token" api:"required"`
}

func (r AccessRuleAccessLinkedAppTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessLinkedAppTokenRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessLinkedAppTokenRuleLinkedAppTokenParam struct {
	// The ID of an Access OIDC SaaS application
	AppUID param.Field[string] `json:"app_uid" api:"required"`
}

func (r AccessRuleAccessLinkedAppTokenRuleLinkedAppTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a user's risk score.
type AccessRuleAccessUserRiskScoreRuleParam struct {
	UserRiskScore param.Field[AccessRuleAccessUserRiskScoreRuleUserRiskScoreParam] `json:"user_risk_score" api:"required"`
}

func (r AccessRuleAccessUserRiskScoreRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessUserRiskScoreRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessUserRiskScoreRuleUserRiskScoreParam struct {
	// A list of risk score levels to match. Values can be low, medium, high, or
	// unscored.
	UserRiskScore param.Field[[]AccessRuleAccessUserRiskScoreRuleUserRiskScoreUserRiskScore] `json:"user_risk_score" api:"required"`
}

func (r AccessRuleAccessUserRiskScoreRuleUserRiskScoreParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches users who are members of a specific Cloudflare account. Requires a
// Cloudflare identity provider.
type AccessRuleAccessCloudflareAccountMemberRuleParam struct {
	CloudflareAccountMember param.Field[AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMemberParam] `json:"cloudflare_account_member" api:"required"`
}

func (r AccessRuleAccessCloudflareAccountMemberRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCloudflareAccountMemberRuleParam) implementsAccessRuleUnionParam() {}

type AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMemberParam struct {
	// Identifier.
	AccountID param.Field[string] `json:"account_id"`
}

func (r AccessRuleAccessCloudflareAccountMemberRuleCloudflareAccountMemberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken AnyValidServiceTokenRuleAnyValidServiceToken `json:"any_valid_service_token" api:"required"`
	JSON                 anyValidServiceTokenRuleJSON                 `json:"-"`
}

// anyValidServiceTokenRuleJSON contains the JSON metadata for the struct
// [AnyValidServiceTokenRule]
type anyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r anyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AnyValidServiceTokenRule) implementsAccessRule() {}

// An empty object which matches on all service tokens.
type AnyValidServiceTokenRuleAnyValidServiceToken struct {
	JSON anyValidServiceTokenRuleAnyValidServiceTokenJSON `json:"-"`
}

// anyValidServiceTokenRuleAnyValidServiceTokenJSON contains the JSON metadata for
// the struct [AnyValidServiceTokenRuleAnyValidServiceToken]
type anyValidServiceTokenRuleAnyValidServiceTokenJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnyValidServiceTokenRuleAnyValidServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r anyValidServiceTokenRuleAnyValidServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type AnyValidServiceTokenRuleParam struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[AnyValidServiceTokenRuleAnyValidServiceTokenParam] `json:"any_valid_service_token" api:"required"`
}

func (r AnyValidServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AnyValidServiceTokenRuleParam) implementsAccessRuleUnionParam() {}

// An empty object which matches on all service tokens.
type AnyValidServiceTokenRuleAnyValidServiceTokenParam struct {
}

func (r AnyValidServiceTokenRuleAnyValidServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AuthenticationMethodRule struct {
	AuthMethod AuthenticationMethodRuleAuthMethod `json:"auth_method" api:"required"`
	JSON       authenticationMethodRuleJSON       `json:"-"`
}

// authenticationMethodRuleJSON contains the JSON metadata for the struct
// [AuthenticationMethodRule]
type authenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AuthenticationMethodRule) implementsAccessRule() {}

type AuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method
	// https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	AuthMethod string                                 `json:"auth_method" api:"required"`
	JSON       authenticationMethodRuleAuthMethodJSON `json:"-"`
}

// authenticationMethodRuleAuthMethodJSON contains the JSON metadata for the struct
// [AuthenticationMethodRuleAuthMethod]
type authenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type AuthenticationMethodRuleParam struct {
	AuthMethod param.Field[AuthenticationMethodRuleAuthMethodParam] `json:"auth_method" api:"required"`
}

func (r AuthenticationMethodRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthenticationMethodRuleParam) implementsAccessRuleUnionParam() {}

type AuthenticationMethodRuleAuthMethodParam struct {
	// The type of authentication method
	// https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	AuthMethod param.Field[string] `json:"auth_method" api:"required"`
}

func (r AuthenticationMethodRuleAuthMethodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AzureGroupRule struct {
	AzureAD AzureGroupRuleAzureAD `json:"azureAD" api:"required"`
	JSON    azureGroupRuleJSON    `json:"-"`
}

// azureGroupRuleJSON contains the JSON metadata for the struct [AzureGroupRule]
type azureGroupRuleJSON struct {
	AzureAD     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AzureGroupRule) implementsAccessRule() {}

type AzureGroupRuleAzureAD struct {
	// The ID of an Azure group.
	ID string `json:"id" api:"required"`
	// The ID of your Azure identity provider.
	IdentityProviderID string                    `json:"identity_provider_id" api:"required"`
	JSON               azureGroupRuleAzureADJSON `json:"-"`
}

// azureGroupRuleAzureADJSON contains the JSON metadata for the struct
// [AzureGroupRuleAzureAD]
type azureGroupRuleAzureADJSON struct {
	ID                 apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AzureGroupRuleAzureAD) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureGroupRuleAzureADJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type AzureGroupRuleParam struct {
	AzureAD param.Field[AzureGroupRuleAzureADParam] `json:"azureAD" api:"required"`
}

func (r AzureGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AzureGroupRuleParam) implementsAccessRuleUnionParam() {}

type AzureGroupRuleAzureADParam struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id" api:"required"`
	// The ID of your Azure identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id" api:"required"`
}

func (r AzureGroupRuleAzureADParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type CertificateRule struct {
	Certificate CertificateRuleCertificate `json:"certificate" api:"required"`
	JSON        certificateRuleJSON        `json:"-"`
}

// certificateRuleJSON contains the JSON metadata for the struct [CertificateRule]
type certificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r CertificateRule) implementsAccessRule() {}

type CertificateRuleCertificate struct {
	JSON certificateRuleCertificateJSON `json:"-"`
}

// certificateRuleCertificateJSON contains the JSON metadata for the struct
// [CertificateRuleCertificate]
type certificateRuleCertificateJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateRuleCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateRuleCertificateJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type CertificateRuleParam struct {
	Certificate param.Field[CertificateRuleCertificateParam] `json:"certificate" api:"required"`
}

func (r CertificateRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CertificateRuleParam) implementsAccessRuleUnionParam() {}

type CertificateRuleCertificateParam struct {
}

func (r CertificateRuleCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type CountryRule struct {
	Geo  CountryRuleGeo  `json:"geo" api:"required"`
	JSON countryRuleJSON `json:"-"`
}

// countryRuleJSON contains the JSON metadata for the struct [CountryRule]
type countryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryRuleJSON) RawJSON() string {
	return r.raw
}

func (r CountryRule) implementsAccessRule() {}

type CountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string             `json:"country_code" api:"required"`
	JSON        countryRuleGeoJSON `json:"-"`
}

// countryRuleGeoJSON contains the JSON metadata for the struct [CountryRuleGeo]
type countryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type CountryRuleParam struct {
	Geo param.Field[CountryRuleGeoParam] `json:"geo" api:"required"`
}

func (r CountryRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CountryRuleParam) implementsAccessRuleUnionParam() {}

type CountryRuleGeoParam struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code" api:"required"`
}

func (r CountryRuleGeoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type DomainRule struct {
	EmailDomain DomainRuleEmailDomain `json:"email_domain" api:"required"`
	JSON        domainRuleJSON        `json:"-"`
}

// domainRuleJSON contains the JSON metadata for the struct [DomainRule]
type domainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainRuleJSON) RawJSON() string {
	return r.raw
}

func (r DomainRule) implementsAccessRule() {}

type DomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                    `json:"domain" api:"required"`
	JSON   domainRuleEmailDomainJSON `json:"-"`
}

// domainRuleEmailDomainJSON contains the JSON metadata for the struct
// [DomainRuleEmailDomain]
type domainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type DomainRuleParam struct {
	EmailDomain param.Field[DomainRuleEmailDomainParam] `json:"email_domain" api:"required"`
}

func (r DomainRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DomainRuleParam) implementsAccessRuleUnionParam() {}

type DomainRuleEmailDomainParam struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain" api:"required"`
}

func (r DomainRuleEmailDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type EmailListRule struct {
	EmailList EmailListRuleEmailList `json:"email_list" api:"required"`
	JSON      emailListRuleJSON      `json:"-"`
}

// emailListRuleJSON contains the JSON metadata for the struct [EmailListRule]
type emailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r EmailListRule) implementsAccessRule() {}

type EmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                     `json:"id" api:"required"`
	JSON emailListRuleEmailListJSON `json:"-"`
}

// emailListRuleEmailListJSON contains the JSON metadata for the struct
// [EmailListRuleEmailList]
type emailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type EmailListRuleParam struct {
	EmailList param.Field[EmailListRuleEmailListParam] `json:"email_list" api:"required"`
}

func (r EmailListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailListRuleParam) implementsAccessRuleUnionParam() {}

type EmailListRuleEmailListParam struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r EmailListRuleEmailListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
type EmailRule struct {
	Email EmailRuleEmail `json:"email" api:"required"`
	JSON  emailRuleJSON  `json:"-"`
}

// emailRuleJSON contains the JSON metadata for the struct [EmailRule]
type emailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleJSON) RawJSON() string {
	return r.raw
}

func (r EmailRule) implementsAccessRule() {}

type EmailRuleEmail struct {
	// The email of the user.
	Email string             `json:"email" api:"required" format:"email"`
	JSON  emailRuleEmailJSON `json:"-"`
}

// emailRuleEmailJSON contains the JSON metadata for the struct [EmailRuleEmail]
type emailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
type EmailRuleParam struct {
	Email param.Field[EmailRuleEmailParam] `json:"email" api:"required"`
}

func (r EmailRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailRuleParam) implementsAccessRuleUnionParam() {}

type EmailRuleEmailParam struct {
	// The email of the user.
	Email param.Field[string] `json:"email" api:"required" format:"email"`
}

func (r EmailRuleEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type EveryoneRule struct {
	// An empty object which matches on all users.
	Everyone EveryoneRuleEveryone `json:"everyone" api:"required"`
	JSON     everyoneRuleJSON     `json:"-"`
}

// everyoneRuleJSON contains the JSON metadata for the struct [EveryoneRule]
type everyoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r everyoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r EveryoneRule) implementsAccessRule() {}

// An empty object which matches on all users.
type EveryoneRuleEveryone struct {
	JSON everyoneRuleEveryoneJSON `json:"-"`
}

// everyoneRuleEveryoneJSON contains the JSON metadata for the struct
// [EveryoneRuleEveryone]
type everyoneRuleEveryoneJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EveryoneRuleEveryone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r everyoneRuleEveryoneJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type EveryoneRuleParam struct {
	// An empty object which matches on all users.
	Everyone param.Field[EveryoneRuleEveryoneParam] `json:"everyone" api:"required"`
}

func (r EveryoneRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EveryoneRuleParam) implementsAccessRuleUnionParam() {}

// An empty object which matches on all users.
type EveryoneRuleEveryoneParam struct {
}

func (r EveryoneRuleEveryoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ExternalEvaluationRule struct {
	ExternalEvaluation ExternalEvaluationRuleExternalEvaluation `json:"external_evaluation" api:"required"`
	JSON               externalEvaluationRuleJSON               `json:"-"`
}

// externalEvaluationRuleJSON contains the JSON metadata for the struct
// [ExternalEvaluationRule]
type externalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ExternalEvaluationRule) implementsAccessRule() {}

type ExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url" api:"required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                       `json:"keys_url" api:"required"`
	JSON    externalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// externalEvaluationRuleExternalEvaluationJSON contains the JSON metadata for the
// struct [ExternalEvaluationRuleExternalEvaluation]
type externalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ExternalEvaluationRuleParam struct {
	ExternalEvaluation param.Field[ExternalEvaluationRuleExternalEvaluationParam] `json:"external_evaluation" api:"required"`
}

func (r ExternalEvaluationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalEvaluationRuleParam) implementsAccessRuleUnionParam() {}

type ExternalEvaluationRuleExternalEvaluationParam struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url" api:"required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url" api:"required"`
}

func (r ExternalEvaluationRuleExternalEvaluationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type GitHubOrganizationRule struct {
	GitHubOrganization GitHubOrganizationRuleGitHubOrganization `json:"github-organization" api:"required"`
	JSON               githubOrganizationRuleJSON               `json:"-"`
}

// githubOrganizationRuleJSON contains the JSON metadata for the struct
// [GitHubOrganizationRule]
type githubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r githubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r GitHubOrganizationRule) implementsAccessRule() {}

type GitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	IdentityProviderID string `json:"identity_provider_id" api:"required"`
	// The name of the organization.
	Name string `json:"name" api:"required"`
	// The name of the team
	Team string                                       `json:"team"`
	JSON githubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// githubOrganizationRuleGitHubOrganizationJSON contains the JSON metadata for the
// struct [GitHubOrganizationRuleGitHubOrganization]
type githubOrganizationRuleGitHubOrganizationJSON struct {
	IdentityProviderID apijson.Field
	Name               apijson.Field
	Team               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r githubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type GitHubOrganizationRuleParam struct {
	GitHubOrganization param.Field[GitHubOrganizationRuleGitHubOrganizationParam] `json:"github-organization" api:"required"`
}

func (r GitHubOrganizationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GitHubOrganizationRuleParam) implementsAccessRuleUnionParam() {}

type GitHubOrganizationRuleGitHubOrganizationParam struct {
	// The ID of your Github identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id" api:"required"`
	// The name of the organization.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the team
	Team param.Field[string] `json:"team"`
}

func (r GitHubOrganizationRuleGitHubOrganizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Access group.
type GroupRule struct {
	Group GroupRuleGroup `json:"group" api:"required"`
	JSON  groupRuleJSON  `json:"-"`
}

// groupRuleJSON contains the JSON metadata for the struct [GroupRule]
type groupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupRuleJSON) RawJSON() string {
	return r.raw
}

func (r GroupRule) implementsAccessRule() {}

type GroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string             `json:"id" api:"required"`
	JSON groupRuleGroupJSON `json:"-"`
}

// groupRuleGroupJSON contains the JSON metadata for the struct [GroupRuleGroup]
type groupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Access group.
type GroupRuleParam struct {
	Group param.Field[GroupRuleGroupParam] `json:"group" api:"required"`
}

func (r GroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GroupRuleParam) implementsAccessRuleUnionParam() {}

type GroupRuleGroupParam struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r GroupRuleGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type GSuiteGroupRule struct {
	GSuite GSuiteGroupRuleGSuite `json:"gsuite" api:"required"`
	JSON   GSuiteGroupRuleJSON   `json:"-"`
}

// GSuiteGroupRuleJSON contains the JSON metadata for the struct [GSuiteGroupRule]
type GSuiteGroupRuleJSON struct {
	GSuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GSuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r GSuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r GSuiteGroupRule) implementsAccessRule() {}

type GSuiteGroupRuleGSuite struct {
	// The email of the Google Workspace group.
	Email string `json:"email" api:"required"`
	// The ID of your Google Workspace identity provider.
	IdentityProviderID string                    `json:"identity_provider_id" api:"required"`
	JSON               GSuiteGroupRuleGSuiteJSON `json:"-"`
}

// GSuiteGroupRuleGSuiteJSON contains the JSON metadata for the struct
// [GSuiteGroupRuleGSuite]
type GSuiteGroupRuleGSuiteJSON struct {
	Email              apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GSuiteGroupRuleGSuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r GSuiteGroupRuleGSuiteJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type GSuiteGroupRuleParam struct {
	GSuite param.Field[GSuiteGroupRuleGSuiteParam] `json:"gsuite" api:"required"`
}

func (r GSuiteGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GSuiteGroupRuleParam) implementsAccessRuleUnionParam() {}

type GSuiteGroupRuleGSuiteParam struct {
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email" api:"required"`
	// The ID of your Google Workspace identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id" api:"required"`
}

func (r GSuiteGroupRuleGSuiteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type IPListRule struct {
	IPList IPListRuleIPList `json:"ip_list" api:"required"`
	JSON   ipListRuleJSON   `json:"-"`
}

// ipListRuleJSON contains the JSON metadata for the struct [IPListRule]
type ipListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListRuleJSON) RawJSON() string {
	return r.raw
}

func (r IPListRule) implementsAccessRule() {}

type IPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string               `json:"id" api:"required"`
	JSON ipListRuleIPListJSON `json:"-"`
}

// ipListRuleIPListJSON contains the JSON metadata for the struct
// [IPListRuleIPList]
type ipListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type IPListRuleParam struct {
	IPList param.Field[IPListRuleIPListParam] `json:"ip_list" api:"required"`
}

func (r IPListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPListRuleParam) implementsAccessRuleUnionParam() {}

type IPListRuleIPListParam struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r IPListRuleIPListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address block.
type IPRule struct {
	IP   IPRuleIP   `json:"ip" api:"required"`
	JSON ipRuleJSON `json:"-"`
}

// ipRuleJSON contains the JSON metadata for the struct [IPRule]
type ipRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipRuleJSON) RawJSON() string {
	return r.raw
}

func (r IPRule) implementsAccessRule() {}

type IPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string       `json:"ip" api:"required"`
	JSON ipRuleIPJSON `json:"-"`
}

// ipRuleIPJSON contains the JSON metadata for the struct [IPRuleIP]
type ipRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address block.
type IPRuleParam struct {
	IP param.Field[IPRuleIPParam] `json:"ip" api:"required"`
}

func (r IPRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPRuleParam) implementsAccessRuleUnionParam() {}

type IPRuleIPParam struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip" api:"required"`
}

func (r IPRuleIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type OktaGroupRule struct {
	Okta OktaGroupRuleOkta `json:"okta" api:"required"`
	JSON oktaGroupRuleJSON `json:"-"`
}

// oktaGroupRuleJSON contains the JSON metadata for the struct [OktaGroupRule]
type oktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r OktaGroupRule) implementsAccessRule() {}

type OktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	IdentityProviderID string `json:"identity_provider_id" api:"required"`
	// The name of the Okta group.
	Name string                `json:"name" api:"required"`
	JSON oktaGroupRuleOktaJSON `json:"-"`
}

// oktaGroupRuleOktaJSON contains the JSON metadata for the struct
// [OktaGroupRuleOkta]
type oktaGroupRuleOktaJSON struct {
	IdentityProviderID apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type OktaGroupRuleParam struct {
	Okta param.Field[OktaGroupRuleOktaParam] `json:"okta" api:"required"`
}

func (r OktaGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OktaGroupRuleParam) implementsAccessRuleUnionParam() {}

type OktaGroupRuleOktaParam struct {
	// The ID of your Okta identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id" api:"required"`
	// The name of the Okta group.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r OktaGroupRuleOktaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SAMLGroupRule struct {
	SAML SAMLGroupRuleSAML `json:"saml" api:"required"`
	JSON samlGroupRuleJSON `json:"-"`
}

// samlGroupRuleJSON contains the JSON metadata for the struct [SAMLGroupRule]
type samlGroupRuleJSON struct {
	SAML        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SAMLGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r SAMLGroupRule) implementsAccessRule() {}

type SAMLGroupRuleSAML struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name" api:"required"`
	// The SAML attribute value to look for.
	AttributeValue string `json:"attribute_value" api:"required"`
	// The ID of your SAML identity provider.
	IdentityProviderID string                `json:"identity_provider_id" api:"required"`
	JSON               samlGroupRuleSAMLJSON `json:"-"`
}

// samlGroupRuleSAMLJSON contains the JSON metadata for the struct
// [SAMLGroupRuleSAML]
type samlGroupRuleSAMLJSON struct {
	AttributeName      apijson.Field
	AttributeValue     apijson.Field
	IdentityProviderID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SAMLGroupRuleSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlGroupRuleSAMLJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type SAMLGroupRuleParam struct {
	SAML param.Field[SAMLGroupRuleSAMLParam] `json:"saml" api:"required"`
}

func (r SAMLGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SAMLGroupRuleParam) implementsAccessRuleUnionParam() {}

type SAMLGroupRuleSAMLParam struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name" api:"required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value" api:"required"`
	// The ID of your SAML identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id" api:"required"`
}

func (r SAMLGroupRuleSAMLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ServiceTokenRule struct {
	ServiceToken ServiceTokenRuleServiceToken `json:"service_token" api:"required"`
	JSON         serviceTokenRuleJSON         `json:"-"`
}

// serviceTokenRuleJSON contains the JSON metadata for the struct
// [ServiceTokenRule]
type serviceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ServiceTokenRule) implementsAccessRule() {}

type ServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                           `json:"token_id" api:"required"`
	JSON    serviceTokenRuleServiceTokenJSON `json:"-"`
}

// serviceTokenRuleServiceTokenJSON contains the JSON metadata for the struct
// [ServiceTokenRuleServiceToken]
type serviceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ServiceTokenRuleParam struct {
	ServiceToken param.Field[ServiceTokenRuleServiceTokenParam] `json:"service_token" api:"required"`
}

func (r ServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceTokenRuleParam) implementsAccessRuleUnionParam() {}

type ServiceTokenRuleServiceTokenParam struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id" api:"required"`
}

func (r ServiceTokenRuleServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyNewResponse struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// This field can have the runtime type of [[]ApprovalGroup].
	ApprovalGroups interface{} `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyNewResponseObjectConnectionRules].
	ConnectionRules interface{} `json:"connection_rules"`
	CreatedAt       time.Time   `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// This field can have the runtime type of [[]AccessRule].
	Exclude interface{} `json:"exclude"`
	// This field can have the runtime type of [[]AccessRule].
	Include interface{} `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyNewResponseObjectMfaConfig].
	MfaConfig interface{} `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// This field can have the runtime type of [[]AccessRule].
	Require interface{} `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                 `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyNewResponseJSON `json:"-"`
	union           AccessApplicationPolicyNewResponseUnion
}

// accessApplicationPolicyNewResponseJSON contains the JSON metadata for the struct
// [AccessApplicationPolicyNewResponse]
type accessApplicationPolicyNewResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r accessApplicationPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationPolicyNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationPolicyNewResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccessApplicationPolicyNewResponseObject],
// [AccessApplicationPolicyNewResponseObject].
func (r AccessApplicationPolicyNewResponse) AsUnion() AccessApplicationPolicyNewResponseUnion {
	return r.union
}

// Union satisfied by [AccessApplicationPolicyNewResponseObject] or
// [AccessApplicationPolicyNewResponseObject].
type AccessApplicationPolicyNewResponseUnion interface {
	implementsAccessApplicationPolicyNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyNewResponseObject{}),
		},
	)
}

type AccessApplicationPolicyNewResponseObject struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// The rules that define how users may connect to targets secured by your
	// application.
	ConnectionRules AccessApplicationPolicyNewResponseObjectConnectionRules `json:"connection_rules"`
	CreatedAt       time.Time                                               `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// Configures multi-factor authentication (MFA) settings.
	MfaConfig AccessApplicationPolicyNewResponseObjectMfaConfig `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessRule `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                       `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                                    `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyNewResponseObjectJSON `json:"-"`
}

// accessApplicationPolicyNewResponseObjectJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyNewResponseObject]
type accessApplicationPolicyNewResponseObjectJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r *AccessApplicationPolicyNewResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyNewResponseObject) implementsAccessApplicationPolicyNewResponse() {}

// The rules that define how users may connect to targets secured by your
// application.
type AccessApplicationPolicyNewResponseObjectConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	RDP  AccessApplicationPolicyNewResponseObjectConnectionRulesRDP  `json:"rdp"`
	JSON accessApplicationPolicyNewResponseObjectConnectionRulesJSON `json:"-"`
}

// accessApplicationPolicyNewResponseObjectConnectionRulesJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseObjectConnectionRules]
type accessApplicationPolicyNewResponseObjectConnectionRulesJSON struct {
	RDP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseObjectConnectionRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseObjectConnectionRulesJSON) RawJSON() string {
	return r.raw
}

// The RDP-specific rules that define clipboard behavior for RDP connections.
type AccessApplicationPolicyNewResponseObjectConnectionRulesRDP struct {
	// Clipboard formats allowed when copying from local machine to remote RDP session.
	AllowedClipboardLocalToRemoteFormats []AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat `json:"allowed_clipboard_local_to_remote_formats"`
	// Clipboard formats allowed when copying from remote RDP session to local machine.
	AllowedClipboardRemoteToLocalFormats []AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat `json:"allowed_clipboard_remote_to_local_formats"`
	JSON                                 accessApplicationPolicyNewResponseObjectConnectionRulesRDPJSON                                  `json:"-"`
}

// accessApplicationPolicyNewResponseObjectConnectionRulesRDPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseObjectConnectionRulesRDP]
type accessApplicationPolicyNewResponseObjectConnectionRulesRDPJSON struct {
	AllowedClipboardLocalToRemoteFormats apijson.Field
	AllowedClipboardRemoteToLocalFormats apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseObjectConnectionRulesRDP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseObjectConnectionRulesRDPJSON) RawJSON() string {
	return r.raw
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat string

const (
	AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "text"
	AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "file"
)

func (r AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText, AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile:
		return true
	}
	return false
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat string

const (
	AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "text"
	AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "file"
)

func (r AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText, AccessApplicationPolicyNewResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile:
		return true
	}
	return false
}

// Configures multi-factor authentication (MFA) settings.
type AccessApplicationPolicyNewResponseObjectMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators []AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticator `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled bool `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration string                                                `json:"session_duration"`
	JSON            accessApplicationPolicyNewResponseObjectMfaConfigJSON `json:"-"`
}

// accessApplicationPolicyNewResponseObjectMfaConfigJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyNewResponseObjectMfaConfig]
type accessApplicationPolicyNewResponseObjectMfaConfigJSON struct {
	AllowedAuthenticators apijson.Field
	MfaDisabled           apijson.Field
	SessionDuration       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseObjectMfaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseObjectMfaConfigJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticatorTotp        AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticator = "totp"
	AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticatorBiometrics  AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticator = "biometrics"
	AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticatorSecurityKey AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticator = "security_key"
)

func (r AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticatorTotp, AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticatorBiometrics, AccessApplicationPolicyNewResponseObjectMfaConfigAllowedAuthenticatorSecurityKey:
		return true
	}
	return false
}

type AccessApplicationPolicyUpdateResponse struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// This field can have the runtime type of [[]ApprovalGroup].
	ApprovalGroups interface{} `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyUpdateResponseObjectConnectionRules].
	ConnectionRules interface{} `json:"connection_rules"`
	CreatedAt       time.Time   `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// This field can have the runtime type of [[]AccessRule].
	Exclude interface{} `json:"exclude"`
	// This field can have the runtime type of [[]AccessRule].
	Include interface{} `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyUpdateResponseObjectMfaConfig].
	MfaConfig interface{} `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// This field can have the runtime type of [[]AccessRule].
	Require interface{} `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                    `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                                 `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyUpdateResponseJSON `json:"-"`
	union           AccessApplicationPolicyUpdateResponseUnion
}

// accessApplicationPolicyUpdateResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyUpdateResponse]
type accessApplicationPolicyUpdateResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r accessApplicationPolicyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationPolicyUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationPolicyUpdateResponseUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccessApplicationPolicyUpdateResponseObject],
// [AccessApplicationPolicyUpdateResponseObject].
func (r AccessApplicationPolicyUpdateResponse) AsUnion() AccessApplicationPolicyUpdateResponseUnion {
	return r.union
}

// Union satisfied by [AccessApplicationPolicyUpdateResponseObject] or
// [AccessApplicationPolicyUpdateResponseObject].
type AccessApplicationPolicyUpdateResponseUnion interface {
	implementsAccessApplicationPolicyUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyUpdateResponseObject{}),
		},
	)
}

type AccessApplicationPolicyUpdateResponseObject struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// The rules that define how users may connect to targets secured by your
	// application.
	ConnectionRules AccessApplicationPolicyUpdateResponseObjectConnectionRules `json:"connection_rules"`
	CreatedAt       time.Time                                                  `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// Configures multi-factor authentication (MFA) settings.
	MfaConfig AccessApplicationPolicyUpdateResponseObjectMfaConfig `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessRule `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                          `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                                       `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyUpdateResponseObjectJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseObjectJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyUpdateResponseObject]
type accessApplicationPolicyUpdateResponseObjectJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r *AccessApplicationPolicyUpdateResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyUpdateResponseObject) implementsAccessApplicationPolicyUpdateResponse() {
}

// The rules that define how users may connect to targets secured by your
// application.
type AccessApplicationPolicyUpdateResponseObjectConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	RDP  AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDP  `json:"rdp"`
	JSON accessApplicationPolicyUpdateResponseObjectConnectionRulesJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseObjectConnectionRulesJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseObjectConnectionRules]
type accessApplicationPolicyUpdateResponseObjectConnectionRulesJSON struct {
	RDP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseObjectConnectionRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseObjectConnectionRulesJSON) RawJSON() string {
	return r.raw
}

// The RDP-specific rules that define clipboard behavior for RDP connections.
type AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDP struct {
	// Clipboard formats allowed when copying from local machine to remote RDP session.
	AllowedClipboardLocalToRemoteFormats []AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat `json:"allowed_clipboard_local_to_remote_formats"`
	// Clipboard formats allowed when copying from remote RDP session to local machine.
	AllowedClipboardRemoteToLocalFormats []AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat `json:"allowed_clipboard_remote_to_local_formats"`
	JSON                                 accessApplicationPolicyUpdateResponseObjectConnectionRulesRDPJSON                                  `json:"-"`
}

// accessApplicationPolicyUpdateResponseObjectConnectionRulesRDPJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDP]
type accessApplicationPolicyUpdateResponseObjectConnectionRulesRDPJSON struct {
	AllowedClipboardLocalToRemoteFormats apijson.Field
	AllowedClipboardRemoteToLocalFormats apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseObjectConnectionRulesRDPJSON) RawJSON() string {
	return r.raw
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat string

const (
	AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "text"
	AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "file"
)

func (r AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText, AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile:
		return true
	}
	return false
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat string

const (
	AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "text"
	AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "file"
)

func (r AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText, AccessApplicationPolicyUpdateResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile:
		return true
	}
	return false
}

// Configures multi-factor authentication (MFA) settings.
type AccessApplicationPolicyUpdateResponseObjectMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators []AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticator `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled bool `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration string                                                   `json:"session_duration"`
	JSON            accessApplicationPolicyUpdateResponseObjectMfaConfigJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseObjectMfaConfigJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyUpdateResponseObjectMfaConfig]
type accessApplicationPolicyUpdateResponseObjectMfaConfigJSON struct {
	AllowedAuthenticators apijson.Field
	MfaDisabled           apijson.Field
	SessionDuration       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseObjectMfaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseObjectMfaConfigJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticatorTotp        AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticator = "totp"
	AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticatorBiometrics  AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticator = "biometrics"
	AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticatorSecurityKey AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticator = "security_key"
)

func (r AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticatorTotp, AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticatorBiometrics, AccessApplicationPolicyUpdateResponseObjectMfaConfigAllowedAuthenticatorSecurityKey:
		return true
	}
	return false
}

type AccessApplicationPolicyListResponse struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// This field can have the runtime type of [[]ApprovalGroup].
	ApprovalGroups interface{} `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyListResponseObjectConnectionRules].
	ConnectionRules interface{} `json:"connection_rules"`
	CreatedAt       time.Time   `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// This field can have the runtime type of [[]AccessRule].
	Exclude interface{} `json:"exclude"`
	// This field can have the runtime type of [[]AccessRule].
	Include interface{} `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyListResponseObjectMfaConfig].
	MfaConfig interface{} `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// This field can have the runtime type of [[]AccessRule].
	Require interface{} `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                  `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                               `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyListResponseJSON `json:"-"`
	union           AccessApplicationPolicyListResponseUnion
}

// accessApplicationPolicyListResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyListResponse]
type accessApplicationPolicyListResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r accessApplicationPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationPolicyListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationPolicyListResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccessApplicationPolicyListResponseObject],
// [AccessApplicationPolicyListResponseObject].
func (r AccessApplicationPolicyListResponse) AsUnion() AccessApplicationPolicyListResponseUnion {
	return r.union
}

// Union satisfied by [AccessApplicationPolicyListResponseObject] or
// [AccessApplicationPolicyListResponseObject].
type AccessApplicationPolicyListResponseUnion interface {
	implementsAccessApplicationPolicyListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyListResponseObject{}),
		},
	)
}

type AccessApplicationPolicyListResponseObject struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// The rules that define how users may connect to targets secured by your
	// application.
	ConnectionRules AccessApplicationPolicyListResponseObjectConnectionRules `json:"connection_rules"`
	CreatedAt       time.Time                                                `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// Configures multi-factor authentication (MFA) settings.
	MfaConfig AccessApplicationPolicyListResponseObjectMfaConfig `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessRule `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                        `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                                     `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyListResponseObjectJSON `json:"-"`
}

// accessApplicationPolicyListResponseObjectJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyListResponseObject]
type accessApplicationPolicyListResponseObjectJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r *AccessApplicationPolicyListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyListResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyListResponseObject) implementsAccessApplicationPolicyListResponse() {}

// The rules that define how users may connect to targets secured by your
// application.
type AccessApplicationPolicyListResponseObjectConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	RDP  AccessApplicationPolicyListResponseObjectConnectionRulesRDP  `json:"rdp"`
	JSON accessApplicationPolicyListResponseObjectConnectionRulesJSON `json:"-"`
}

// accessApplicationPolicyListResponseObjectConnectionRulesJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyListResponseObjectConnectionRules]
type accessApplicationPolicyListResponseObjectConnectionRulesJSON struct {
	RDP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseObjectConnectionRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyListResponseObjectConnectionRulesJSON) RawJSON() string {
	return r.raw
}

// The RDP-specific rules that define clipboard behavior for RDP connections.
type AccessApplicationPolicyListResponseObjectConnectionRulesRDP struct {
	// Clipboard formats allowed when copying from local machine to remote RDP session.
	AllowedClipboardLocalToRemoteFormats []AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat `json:"allowed_clipboard_local_to_remote_formats"`
	// Clipboard formats allowed when copying from remote RDP session to local machine.
	AllowedClipboardRemoteToLocalFormats []AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat `json:"allowed_clipboard_remote_to_local_formats"`
	JSON                                 accessApplicationPolicyListResponseObjectConnectionRulesRDPJSON                                  `json:"-"`
}

// accessApplicationPolicyListResponseObjectConnectionRulesRDPJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyListResponseObjectConnectionRulesRDP]
type accessApplicationPolicyListResponseObjectConnectionRulesRDPJSON struct {
	AllowedClipboardLocalToRemoteFormats apijson.Field
	AllowedClipboardRemoteToLocalFormats apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseObjectConnectionRulesRDP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyListResponseObjectConnectionRulesRDPJSON) RawJSON() string {
	return r.raw
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat string

const (
	AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "text"
	AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "file"
)

func (r AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText, AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile:
		return true
	}
	return false
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat string

const (
	AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "text"
	AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "file"
)

func (r AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText, AccessApplicationPolicyListResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile:
		return true
	}
	return false
}

// Configures multi-factor authentication (MFA) settings.
type AccessApplicationPolicyListResponseObjectMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators []AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticator `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled bool `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration string                                                 `json:"session_duration"`
	JSON            accessApplicationPolicyListResponseObjectMfaConfigJSON `json:"-"`
}

// accessApplicationPolicyListResponseObjectMfaConfigJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyListResponseObjectMfaConfig]
type accessApplicationPolicyListResponseObjectMfaConfigJSON struct {
	AllowedAuthenticators apijson.Field
	MfaDisabled           apijson.Field
	SessionDuration       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessApplicationPolicyListResponseObjectMfaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyListResponseObjectMfaConfigJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticatorTotp        AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticator = "totp"
	AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticatorBiometrics  AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticator = "biometrics"
	AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticatorSecurityKey AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticator = "security_key"
)

func (r AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticatorTotp, AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticatorBiometrics, AccessApplicationPolicyListResponseObjectMfaConfigAllowedAuthenticatorSecurityKey:
		return true
	}
	return false
}

type AccessApplicationPolicyDeleteResponse struct {
	// UUID.
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
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// This field can have the runtime type of [[]ApprovalGroup].
	ApprovalGroups interface{} `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyGetResponseObjectConnectionRules].
	ConnectionRules interface{} `json:"connection_rules"`
	CreatedAt       time.Time   `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// This field can have the runtime type of [[]AccessRule].
	Exclude interface{} `json:"exclude"`
	// This field can have the runtime type of [[]AccessRule].
	Include interface{} `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// This field can have the runtime type of
	// [AccessApplicationPolicyGetResponseObjectMfaConfig].
	MfaConfig interface{} `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// This field can have the runtime type of [[]AccessRule].
	Require interface{} `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                 `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                              `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyGetResponseJSON `json:"-"`
	union           AccessApplicationPolicyGetResponseUnion
}

// accessApplicationPolicyGetResponseJSON contains the JSON metadata for the struct
// [AccessApplicationPolicyGetResponse]
type accessApplicationPolicyGetResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r accessApplicationPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AccessApplicationPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AccessApplicationPolicyGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessApplicationPolicyGetResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccessApplicationPolicyGetResponseObject],
// [AccessApplicationPolicyGetResponseObject].
func (r AccessApplicationPolicyGetResponse) AsUnion() AccessApplicationPolicyGetResponseUnion {
	return r.union
}

// Union satisfied by [AccessApplicationPolicyGetResponseObject] or
// [AccessApplicationPolicyGetResponseObject].
type AccessApplicationPolicyGetResponseUnion interface {
	implementsAccessApplicationPolicyGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationPolicyGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessApplicationPolicyGetResponseObject{}),
		},
	)
}

type AccessApplicationPolicyGetResponseObject struct {
	// The UUID of the policy
	ID string `json:"id"`
	// Identifier.
	AccountID string `json:"account_id"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroup `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// The rules that define how users may connect to targets secured by your
	// application.
	ConnectionRules AccessApplicationPolicyGetResponseObjectConnectionRules `json:"connection_rules"`
	CreatedAt       time.Time                                               `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision Decision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// Configures multi-factor authentication (MFA) settings.
	MfaConfig AccessApplicationPolicyGetResponseObjectMfaConfig `json:"mfa_config"`
	// The name of the Access policy.
	Name string `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require []AccessRule `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                                       `json:"session_duration" api:"nullable"`
	UpdatedAt       time.Time                                    `json:"updated_at" format:"date-time"`
	JSON            accessApplicationPolicyGetResponseObjectJSON `json:"-"`
}

// accessApplicationPolicyGetResponseObjectJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyGetResponseObject]
type accessApplicationPolicyGetResponseObjectJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	ConnectionRules              apijson.Field
	CreatedAt                    apijson.Field
	Decision                     apijson.Field
	Exclude                      apijson.Field
	Include                      apijson.Field
	IsolationRequired            apijson.Field
	MfaConfig                    apijson.Field
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

func (r *AccessApplicationPolicyGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccessApplicationPolicyGetResponseObject) implementsAccessApplicationPolicyGetResponse() {}

// The rules that define how users may connect to targets secured by your
// application.
type AccessApplicationPolicyGetResponseObjectConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	RDP  AccessApplicationPolicyGetResponseObjectConnectionRulesRDP  `json:"rdp"`
	JSON accessApplicationPolicyGetResponseObjectConnectionRulesJSON `json:"-"`
}

// accessApplicationPolicyGetResponseObjectConnectionRulesJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseObjectConnectionRules]
type accessApplicationPolicyGetResponseObjectConnectionRulesJSON struct {
	RDP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseObjectConnectionRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseObjectConnectionRulesJSON) RawJSON() string {
	return r.raw
}

// The RDP-specific rules that define clipboard behavior for RDP connections.
type AccessApplicationPolicyGetResponseObjectConnectionRulesRDP struct {
	// Clipboard formats allowed when copying from local machine to remote RDP session.
	AllowedClipboardLocalToRemoteFormats []AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat `json:"allowed_clipboard_local_to_remote_formats"`
	// Clipboard formats allowed when copying from remote RDP session to local machine.
	AllowedClipboardRemoteToLocalFormats []AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat `json:"allowed_clipboard_remote_to_local_formats"`
	JSON                                 accessApplicationPolicyGetResponseObjectConnectionRulesRDPJSON                                  `json:"-"`
}

// accessApplicationPolicyGetResponseObjectConnectionRulesRDPJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseObjectConnectionRulesRDP]
type accessApplicationPolicyGetResponseObjectConnectionRulesRDPJSON struct {
	AllowedClipboardLocalToRemoteFormats apijson.Field
	AllowedClipboardRemoteToLocalFormats apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseObjectConnectionRulesRDP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseObjectConnectionRulesRDPJSON) RawJSON() string {
	return r.raw
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat string

const (
	AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "text"
	AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "file"
)

func (r AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText, AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile:
		return true
	}
	return false
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat string

const (
	AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "text"
	AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "file"
)

func (r AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText, AccessApplicationPolicyGetResponseObjectConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile:
		return true
	}
	return false
}

// Configures multi-factor authentication (MFA) settings.
type AccessApplicationPolicyGetResponseObjectMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators []AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticator `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled bool `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration string                                                `json:"session_duration"`
	JSON            accessApplicationPolicyGetResponseObjectMfaConfigJSON `json:"-"`
}

// accessApplicationPolicyGetResponseObjectMfaConfigJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyGetResponseObjectMfaConfig]
type accessApplicationPolicyGetResponseObjectMfaConfigJSON struct {
	AllowedAuthenticators apijson.Field
	MfaDisabled           apijson.Field
	SessionDuration       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseObjectMfaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseObjectMfaConfigJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticatorTotp        AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticator = "totp"
	AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticatorBiometrics  AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticator = "biometrics"
	AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticatorSecurityKey AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticator = "security_key"
)

func (r AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticatorTotp, AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticatorBiometrics, AccessApplicationPolicyGetResponseObjectMfaConfigAllowedAuthenticatorSecurityKey:
		return true
	}
	return false
}

type AccessApplicationPolicyNewParams struct {
	Body AccessApplicationPolicyNewParamsBodyUnion `json:"body" api:"required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessApplicationPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessApplicationPolicyNewParamsBody struct {
	ApprovalGroups param.Field[interface{}] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool]        `json:"approval_required"`
	ConnectionRules  param.Field[interface{}] `json:"connection_rules"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude"`
	Include  param.Field[interface{}] `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool]        `json:"isolation_required"`
	MfaConfig         param.Field[interface{}] `json:"mfa_config"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsBody) implementsAccessApplicationPolicyNewParamsBodyUnion() {}

// Satisfied by
// [zero_trust.AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequest],
// [zero_trust.AccessApplicationPolicyNewParamsBodyObject],
// [AccessApplicationPolicyNewParamsBody].
type AccessApplicationPolicyNewParamsBodyUnion interface {
	implementsAccessApplicationPolicyNewParamsBodyUnion()
}

type AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequest struct {
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The rules that define how users may connect to targets secured by your
	// application.
	ConnectionRules param.Field[AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRules] `json:"connection_rules"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// Configures multi-factor authentication (MFA) settings.
	MfaConfig param.Field[AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfig] `json:"mfa_config"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequest) implementsAccessApplicationPolicyNewParamsBodyUnion() {
}

// The rules that define how users may connect to targets secured by your
// application.
type AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	RDP param.Field[AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDP] `json:"rdp"`
}

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRules) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The RDP-specific rules that define clipboard behavior for RDP connections.
type AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDP struct {
	// Clipboard formats allowed when copying from local machine to remote RDP session.
	AllowedClipboardLocalToRemoteFormats param.Field[[]AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat] `json:"allowed_clipboard_local_to_remote_formats"`
	// Clipboard formats allowed when copying from remote RDP session to local machine.
	AllowedClipboardRemoteToLocalFormats param.Field[[]AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat] `json:"allowed_clipboard_remote_to_local_formats"`
}

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat string

const (
	AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "text"
	AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "file"
)

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText, AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile:
		return true
	}
	return false
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat string

const (
	AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "text"
	AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "file"
)

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText, AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile:
		return true
	}
	return false
}

// Configures multi-factor authentication (MFA) settings.
type AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators param.Field[[]AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator] `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled param.Field[bool] `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorTotp        AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator = "totp"
	AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorBiometrics  AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator = "biometrics"
	AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorSecurityKey AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator = "security_key"
)

func (r AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorTotp, AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorBiometrics, AccessApplicationPolicyNewParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorSecurityKey:
		return true
	}
	return false
}

type AccessApplicationPolicyNewParamsBodyObject struct {
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision] `json:"decision" api:"required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include" api:"required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name" api:"required"`
	// The rules that define how users may connect to the targets secured by your
	// application.
	ConnectionRules param.Field[AccessApplicationPolicyNewParamsBodyObjectConnectionRules] `json:"connection_rules"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Configures multi-factor authentication (MFA) settings for infrastructure
	// applications.
	MfaConfig param.Field[AccessApplicationPolicyNewParamsBodyObjectMfaConfig] `json:"mfa_config"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r AccessApplicationPolicyNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyNewParamsBodyObject) implementsAccessApplicationPolicyNewParamsBodyUnion() {
}

// The rules that define how users may connect to the targets secured by your
// application.
type AccessApplicationPolicyNewParamsBodyObjectConnectionRules struct {
	// The SSH-specific rules that define how users may connect to the targets secured
	// by your application.
	SSH param.Field[AccessApplicationPolicyNewParamsBodyObjectConnectionRulesSSH] `json:"ssh"`
}

func (r AccessApplicationPolicyNewParamsBodyObjectConnectionRules) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The SSH-specific rules that define how users may connect to the targets secured
// by your application.
type AccessApplicationPolicyNewParamsBodyObjectConnectionRulesSSH struct {
	// Contains the Unix usernames that may be used when connecting over SSH.
	Usernames param.Field[[]string] `json:"usernames" api:"required"`
	// Enables using Identity Provider email alias as SSH username.
	AllowEmailAlias param.Field[bool] `json:"allow_email_alias"`
}

func (r AccessApplicationPolicyNewParamsBodyObjectConnectionRulesSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures multi-factor authentication (MFA) settings for infrastructure
// applications.
type AccessApplicationPolicyNewParamsBodyObjectMfaConfig struct {
	// Lists the MFA methods that users can authenticate with. For infrastructure
	// applications, supported values are `piv_key` and `ssh_fido2_key`.
	AllowedAuthenticators param.Field[[]AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticator] `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled param.Field[bool] `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples: `5m` or `24h`.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyNewParamsBodyObjectMfaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticatorPivKey      AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticator = "piv_key"
	AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticatorSSHFido2Key AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticator = "ssh_fido2_key"
)

func (r AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticatorPivKey, AccessApplicationPolicyNewParamsBodyObjectMfaConfigAllowedAuthenticatorSSHFido2Key:
		return true
	}
	return false
}

type AccessApplicationPolicyNewResponseEnvelope struct {
	Errors   []AccessApplicationPolicyNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessApplicationPolicyNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessApplicationPolicyNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessApplicationPolicyNewResponse                `json:"result"`
	JSON    accessApplicationPolicyNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyNewResponseEnvelope]
type accessApplicationPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessApplicationPolicyNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationPolicyNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyNewResponseEnvelopeErrors]
type accessApplicationPolicyNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessApplicationPolicyNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyNewResponseEnvelopeErrorsSource]
type accessApplicationPolicyNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyNewResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessApplicationPolicyNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationPolicyNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyNewResponseEnvelopeMessages]
type accessApplicationPolicyNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessApplicationPolicyNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationPolicyNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyNewResponseEnvelopeMessagesSource]
type accessApplicationPolicyNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	Body AccessApplicationPolicyUpdateParamsBodyUnion `json:"body" api:"required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r AccessApplicationPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessApplicationPolicyUpdateParamsBody struct {
	ApprovalGroups param.Field[interface{}] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool]        `json:"approval_required"`
	ConnectionRules  param.Field[interface{}] `json:"connection_rules"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision]    `json:"decision"`
	Exclude  param.Field[interface{}] `json:"exclude"`
	Include  param.Field[interface{}] `json:"include"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool]        `json:"isolation_required"`
	MfaConfig         param.Field[interface{}] `json:"mfa_config"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool]        `json:"purpose_justification_required"`
	Require                      param.Field[interface{}] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsBody) implementsAccessApplicationPolicyUpdateParamsBodyUnion() {
}

// Satisfied by
// [zero_trust.AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequest],
// [zero_trust.AccessApplicationPolicyUpdateParamsBodyObject],
// [AccessApplicationPolicyUpdateParamsBody].
type AccessApplicationPolicyUpdateParamsBodyUnion interface {
	implementsAccessApplicationPolicyUpdateParamsBodyUnion()
}

type AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequest struct {
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// The rules that define how users may connect to targets secured by your
	// application.
	ConnectionRules param.Field[AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRules] `json:"connection_rules"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// Configures multi-factor authentication (MFA) settings.
	MfaConfig param.Field[AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfig] `json:"mfa_config"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequest) implementsAccessApplicationPolicyUpdateParamsBodyUnion() {
}

// The rules that define how users may connect to targets secured by your
// application.
type AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRules struct {
	// The RDP-specific rules that define clipboard behavior for RDP connections.
	RDP param.Field[AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDP] `json:"rdp"`
}

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRules) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The RDP-specific rules that define clipboard behavior for RDP connections.
type AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDP struct {
	// Clipboard formats allowed when copying from local machine to remote RDP session.
	AllowedClipboardLocalToRemoteFormats param.Field[[]AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat] `json:"allowed_clipboard_local_to_remote_formats"`
	// Clipboard formats allowed when copying from remote RDP session to local machine.
	AllowedClipboardRemoteToLocalFormats param.Field[[]AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat] `json:"allowed_clipboard_remote_to_local_formats"`
}

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat string

const (
	AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "text"
	AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat = "file"
)

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatText, AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardLocalToRemoteFormatFile:
		return true
	}
	return false
}

// Clipboard format for RDP connections.
type AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat string

const (
	AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "text"
	AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat = "file"
)

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormat) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatText, AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestConnectionRulesRDPAllowedClipboardRemoteToLocalFormatFile:
		return true
	}
	return false
}

// Configures multi-factor authentication (MFA) settings.
type AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	AllowedAuthenticators param.Field[[]AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator] `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled param.Field[bool] `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorTotp        AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator = "totp"
	AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorBiometrics  AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator = "biometrics"
	AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorSecurityKey AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator = "security_key"
)

func (r AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorTotp, AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorBiometrics, AccessApplicationPolicyUpdateParamsBodyAccessAppPolicyRequestMfaConfigAllowedAuthenticatorSecurityKey:
		return true
	}
	return false
}

type AccessApplicationPolicyUpdateParamsBodyObject struct {
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision] `json:"decision" api:"required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include" api:"required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name" api:"required"`
	// The rules that define how users may connect to the targets secured by your
	// application.
	ConnectionRules param.Field[AccessApplicationPolicyUpdateParamsBodyObjectConnectionRules] `json:"connection_rules"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Configures multi-factor authentication (MFA) settings for infrastructure
	// applications.
	MfaConfig param.Field[AccessApplicationPolicyUpdateParamsBodyObjectMfaConfig] `json:"mfa_config"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r AccessApplicationPolicyUpdateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyUpdateParamsBodyObject) implementsAccessApplicationPolicyUpdateParamsBodyUnion() {
}

// The rules that define how users may connect to the targets secured by your
// application.
type AccessApplicationPolicyUpdateParamsBodyObjectConnectionRules struct {
	// The SSH-specific rules that define how users may connect to the targets secured
	// by your application.
	SSH param.Field[AccessApplicationPolicyUpdateParamsBodyObjectConnectionRulesSSH] `json:"ssh"`
}

func (r AccessApplicationPolicyUpdateParamsBodyObjectConnectionRules) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The SSH-specific rules that define how users may connect to the targets secured
// by your application.
type AccessApplicationPolicyUpdateParamsBodyObjectConnectionRulesSSH struct {
	// Contains the Unix usernames that may be used when connecting over SSH.
	Usernames param.Field[[]string] `json:"usernames" api:"required"`
	// Enables using Identity Provider email alias as SSH username.
	AllowEmailAlias param.Field[bool] `json:"allow_email_alias"`
}

func (r AccessApplicationPolicyUpdateParamsBodyObjectConnectionRulesSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures multi-factor authentication (MFA) settings for infrastructure
// applications.
type AccessApplicationPolicyUpdateParamsBodyObjectMfaConfig struct {
	// Lists the MFA methods that users can authenticate with. For infrastructure
	// applications, supported values are `piv_key` and `ssh_fido2_key`.
	AllowedAuthenticators param.Field[[]AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticator] `json:"allowed_authenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at
	// the application and policy level.
	MfaDisabled param.Field[bool] `json:"mfa_disabled"`
	// Defines the duration of an MFA session. Must be in minutes (m) or hours (h).
	// Minimum: 0m. Maximum: 720h (30 days). Examples: `5m` or `24h`.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyUpdateParamsBodyObjectMfaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticator string

const (
	AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticatorPivKey      AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticator = "piv_key"
	AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticatorSSHFido2Key AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticator = "ssh_fido2_key"
)

func (r AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticator) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticatorPivKey, AccessApplicationPolicyUpdateParamsBodyObjectMfaConfigAllowedAuthenticatorSSHFido2Key:
		return true
	}
	return false
}

type AccessApplicationPolicyUpdateResponseEnvelope struct {
	Errors   []AccessApplicationPolicyUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessApplicationPolicyUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessApplicationPolicyUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessApplicationPolicyUpdateResponse                `json:"result"`
	JSON    accessApplicationPolicyUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyUpdateResponseEnvelope]
type accessApplicationPolicyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessApplicationPolicyUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyUpdateResponseEnvelopeErrors]
type accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessApplicationPolicyUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyUpdateResponseEnvelopeErrorsSource]
type accessApplicationPolicyUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyUpdateResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccessApplicationPolicyUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyUpdateResponseEnvelopeMessages]
type accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accessApplicationPolicyUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationPolicyUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyUpdateResponseEnvelopeMessagesSource]
type accessApplicationPolicyUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccessApplicationPolicyListParams]'s query parameters as
// `url.Values`.
func (r AccessApplicationPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessApplicationPolicyDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationPolicyDeleteResponseEnvelope struct {
	Errors   []AccessApplicationPolicyDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessApplicationPolicyDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessApplicationPolicyDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessApplicationPolicyDeleteResponse                `json:"result"`
	JSON    accessApplicationPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyDeleteResponseEnvelope]
type accessApplicationPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessApplicationPolicyDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyDeleteResponseEnvelopeErrors]
type accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessApplicationPolicyDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyDeleteResponseEnvelopeErrorsSource]
type accessApplicationPolicyDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyDeleteResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccessApplicationPolicyDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyDeleteResponseEnvelopeMessages]
type accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accessApplicationPolicyDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationPolicyDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyDeleteResponseEnvelopeMessagesSource]
type accessApplicationPolicyDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	Errors   []AccessApplicationPolicyGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessApplicationPolicyGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessApplicationPolicyGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessApplicationPolicyGetResponse                `json:"result"`
	JSON    accessApplicationPolicyGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationPolicyGetResponseEnvelope]
type accessApplicationPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessApplicationPolicyGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationPolicyGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyGetResponseEnvelopeErrors]
type accessApplicationPolicyGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessApplicationPolicyGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyGetResponseEnvelopeErrorsSource]
type accessApplicationPolicyGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyGetResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessApplicationPolicyGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationPolicyGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyGetResponseEnvelopeMessages]
type accessApplicationPolicyGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessApplicationPolicyGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationPolicyGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyGetResponseEnvelopeMessagesSource]
type accessApplicationPolicyGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
