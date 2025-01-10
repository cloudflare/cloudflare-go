// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/tidwall/gjson"
)

// AccessService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessService] method instead.
type AccessService struct {
	Options        []option.RequestOption
	GatewayCA      *AccessGatewayCAService
	Infrastructure *AccessInfrastructureService
	Applications   *AccessApplicationService
	Certificates   *AccessCertificateService
	Groups         *AccessGroupService
	ServiceTokens  *AccessServiceTokenService
	Bookmarks      *AccessBookmarkService
	Keys           *AccessKeyService
	Logs           *AccessLogService
	Users          *AccessUserService
	CustomPages    *AccessCustomPageService
	Tags           *AccessTagService
	Policies       *AccessPolicyService
}

// NewAccessService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccessService(opts ...option.RequestOption) (r *AccessService) {
	r = &AccessService{}
	r.Options = opts
	r.GatewayCA = NewAccessGatewayCAService(opts...)
	r.Infrastructure = NewAccessInfrastructureService(opts...)
	r.Applications = NewAccessApplicationService(opts...)
	r.Certificates = NewAccessCertificateService(opts...)
	r.Groups = NewAccessGroupService(opts...)
	r.ServiceTokens = NewAccessServiceTokenService(opts...)
	r.Bookmarks = NewAccessBookmarkService(opts...)
	r.Keys = NewAccessKeyService(opts...)
	r.Logs = NewAccessLogService(opts...)
	r.Users = NewAccessUserService(opts...)
	r.CustomPages = NewAccessCustomPageService(opts...)
	r.Tags = NewAccessTagService(opts...)
	r.Policies = NewAccessPolicyService(opts...)
	return
}

// Enforces a device posture rule has run successfully
type AccessDevicePostureRule struct {
	DevicePosture AccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
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

func (r AccessDevicePostureRule) implementsZeroTrustAccessRule() {}

type AccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUID string                                   `json:"integration_uid,required"`
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
	DevicePosture param.Field[AccessDevicePostureRuleDevicePostureParam] `json:"device_posture,required"`
}

func (r AccessDevicePostureRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessDevicePostureRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessDevicePostureRuleDevicePostureParam struct {
	// The ID of a device posture integration.
	IntegrationUID param.Field[string] `json:"integration_uid,required"`
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
	// This field can have the runtime type of [OktaGroupRuleOkta].
	Okta interface{} `json:"okta"`
	// This field can have the runtime type of [SAMLGroupRuleSAML].
	SAML interface{} `json:"saml"`
	// This field can have the runtime type of [ServiceTokenRuleServiceToken].
	ServiceToken interface{}    `json:"service_token"`
	JSON         accessRuleJSON `json:"-"`
	union        AccessRuleUnion
}

// accessRuleJSON contains the JSON metadata for the struct [AccessRule]
type accessRuleJSON struct {
	AnyValidServiceToken apijson.Field
	AuthContext          apijson.Field
	AuthMethod           apijson.Field
	AzureAD              apijson.Field
	Certificate          apijson.Field
	CommonName           apijson.Field
	DevicePosture        apijson.Field
	Email                apijson.Field
	EmailDomain          apijson.Field
	EmailList            apijson.Field
	Everyone             apijson.Field
	ExternalEvaluation   apijson.Field
	Geo                  apijson.Field
	GitHubOrganization   apijson.Field
	Group                apijson.Field
	GSuite               apijson.Field
	IP                   apijson.Field
	IPList               apijson.Field
	Okta                 apijson.Field
	SAML                 apijson.Field
	ServiceToken         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
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
// Possible runtime types of the union are [zero_trust.GroupRule],
// [zero_trust.AnyValidServiceTokenRule],
// [zero_trust.AccessRuleAccessAuthContextRule],
// [zero_trust.AuthenticationMethodRule], [zero_trust.AzureGroupRule],
// [zero_trust.CertificateRule], [zero_trust.AccessRuleAccessCommonNameRule],
// [zero_trust.CountryRule], [zero_trust.AccessDevicePostureRule],
// [zero_trust.DomainRule], [zero_trust.EmailListRule], [zero_trust.EmailRule],
// [zero_trust.EveryoneRule], [zero_trust.ExternalEvaluationRule],
// [zero_trust.GitHubOrganizationRule], [zero_trust.GSuiteGroupRule],
// [zero_trust.IPListRule], [zero_trust.IPRule], [zero_trust.OktaGroupRule],
// [zero_trust.SAMLGroupRule], [zero_trust.ServiceTokenRule].
func (r AccessRule) AsUnion() AccessRuleUnion {
	return r.union
}

// Matches an Access group.
//
// Union satisfied by [zero_trust.GroupRule],
// [zero_trust.AnyValidServiceTokenRule],
// [zero_trust.AccessRuleAccessAuthContextRule],
// [zero_trust.AuthenticationMethodRule], [zero_trust.AzureGroupRule],
// [zero_trust.CertificateRule], [zero_trust.AccessRuleAccessCommonNameRule],
// [zero_trust.CountryRule], [zero_trust.AccessDevicePostureRule],
// [zero_trust.DomainRule], [zero_trust.EmailListRule], [zero_trust.EmailRule],
// [zero_trust.EveryoneRule], [zero_trust.ExternalEvaluationRule],
// [zero_trust.GitHubOrganizationRule], [zero_trust.GSuiteGroupRule],
// [zero_trust.IPListRule], [zero_trust.IPRule], [zero_trust.OktaGroupRule],
// [zero_trust.SAMLGroupRule] or [zero_trust.ServiceTokenRule].
type AccessRuleUnion interface {
	implementsZeroTrustAccessRule()
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
			Type:       reflect.TypeOf(ServiceTokenRule{}),
		},
	)
}

// Matches an Azure Authentication Context. Requires an Azure identity provider.
type AccessRuleAccessAuthContextRule struct {
	AuthContext AccessRuleAccessAuthContextRuleAuthContext `json:"auth_context,required"`
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

func (r AccessRuleAccessAuthContextRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessAuthContextRuleAuthContext struct {
	// The ID of an Authentication context.
	ID string `json:"id,required"`
	// The ACID of an Authentication context.
	AcID string `json:"ac_id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID string                                         `json:"identity_provider_id,required"`
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
	CommonName AccessRuleAccessCommonNameRuleCommonName `json:"common_name,required"`
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

func (r AccessRuleAccessCommonNameRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessCommonNameRuleCommonName struct {
	// The common name to match.
	CommonName string                                       `json:"common_name,required"`
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

// Matches an Access group.
type AccessRuleParam struct {
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token"`
	AuthContext          param.Field[interface{}] `json:"auth_context"`
	AuthMethod           param.Field[interface{}] `json:"auth_method"`
	AzureAD              param.Field[interface{}] `json:"azureAD"`
	Certificate          param.Field[interface{}] `json:"certificate"`
	CommonName           param.Field[interface{}] `json:"common_name"`
	DevicePosture        param.Field[interface{}] `json:"device_posture"`
	Email                param.Field[interface{}] `json:"email"`
	EmailDomain          param.Field[interface{}] `json:"email_domain"`
	EmailList            param.Field[interface{}] `json:"email_list"`
	Everyone             param.Field[interface{}] `json:"everyone"`
	ExternalEvaluation   param.Field[interface{}] `json:"external_evaluation"`
	Geo                  param.Field[interface{}] `json:"geo"`
	GitHubOrganization   param.Field[interface{}] `json:"github-organization"`
	Group                param.Field[interface{}] `json:"group"`
	GSuite               param.Field[interface{}] `json:"gsuite"`
	IP                   param.Field[interface{}] `json:"ip"`
	IPList               param.Field[interface{}] `json:"ip_list"`
	Okta                 param.Field[interface{}] `json:"okta"`
	SAML                 param.Field[interface{}] `json:"saml"`
	ServiceToken         param.Field[interface{}] `json:"service_token"`
}

func (r AccessRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

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
// [zero_trust.IPListRuleParam], [zero_trust.IPRuleParam],
// [zero_trust.OktaGroupRuleParam], [zero_trust.SAMLGroupRuleParam],
// [zero_trust.ServiceTokenRuleParam], [AccessRuleParam].
type AccessRuleUnionParam interface {
	implementsZeroTrustAccessRuleUnionParam()
}

// Matches an Azure Authentication Context. Requires an Azure identity provider.
type AccessRuleAccessAuthContextRuleParam struct {
	AuthContext param.Field[AccessRuleAccessAuthContextRuleAuthContextParam] `json:"auth_context,required"`
}

func (r AccessRuleAccessAuthContextRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAuthContextRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessAuthContextRuleAuthContextParam struct {
	// The ID of an Authentication context.
	ID param.Field[string] `json:"id,required"`
	// The ACID of an Authentication context.
	AcID param.Field[string] `json:"ac_id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r AccessRuleAccessAuthContextRuleAuthContextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific common name.
type AccessRuleAccessCommonNameRuleParam struct {
	CommonName param.Field[AccessRuleAccessCommonNameRuleCommonNameParam] `json:"common_name,required"`
}

func (r AccessRuleAccessCommonNameRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCommonNameRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessCommonNameRuleCommonNameParam struct {
	// The common name to match.
	CommonName param.Field[string] `json:"common_name,required"`
}

func (r AccessRuleAccessCommonNameRuleCommonNameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken AnyValidServiceTokenRuleAnyValidServiceToken `json:"any_valid_service_token,required"`
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

func (r AnyValidServiceTokenRule) implementsZeroTrustAccessRule() {}

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
	AnyValidServiceToken param.Field[AnyValidServiceTokenRuleAnyValidServiceTokenParam] `json:"any_valid_service_token,required"`
}

func (r AnyValidServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

// An empty object which matches on all service tokens.
type AnyValidServiceTokenRuleAnyValidServiceTokenParam struct {
}

func (r AnyValidServiceTokenRuleAnyValidServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AuthenticationMethodRule struct {
	AuthMethod AuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
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

func (r AuthenticationMethodRule) implementsZeroTrustAccessRule() {}

type AuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method
	// https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	AuthMethod string                                 `json:"auth_method,required"`
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
	AuthMethod param.Field[AuthenticationMethodRuleAuthMethodParam] `json:"auth_method,required"`
}

func (r AuthenticationMethodRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthenticationMethodRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AuthenticationMethodRuleAuthMethodParam struct {
	// The type of authentication method
	// https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AuthenticationMethodRuleAuthMethodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AzureGroupRule struct {
	AzureAD AzureGroupRuleAzureAD `json:"azureAD,required"`
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

func (r AzureGroupRule) implementsZeroTrustAccessRule() {}

type AzureGroupRuleAzureAD struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID string                    `json:"identity_provider_id,required"`
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
	AzureAD param.Field[AzureGroupRuleAzureADParam] `json:"azureAD,required"`
}

func (r AzureGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AzureGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AzureGroupRuleAzureADParam struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r AzureGroupRuleAzureADParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type CertificateRule struct {
	Certificate CertificateRuleCertificate `json:"certificate,required"`
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

func (r CertificateRule) implementsZeroTrustAccessRule() {}

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
	Certificate param.Field[CertificateRuleCertificateParam] `json:"certificate,required"`
}

func (r CertificateRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CertificateRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type CertificateRuleCertificateParam struct {
}

func (r CertificateRuleCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type CountryRule struct {
	Geo  CountryRuleGeo  `json:"geo,required"`
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

func (r CountryRule) implementsZeroTrustAccessRule() {}

type CountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string             `json:"country_code,required"`
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
	Geo param.Field[CountryRuleGeoParam] `json:"geo,required"`
}

func (r CountryRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CountryRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type CountryRuleGeoParam struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r CountryRuleGeoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type DomainRule struct {
	EmailDomain DomainRuleEmailDomain `json:"email_domain,required"`
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

func (r DomainRule) implementsZeroTrustAccessRule() {}

type DomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                    `json:"domain,required"`
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
	EmailDomain param.Field[DomainRuleEmailDomainParam] `json:"email_domain,required"`
}

func (r DomainRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DomainRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type DomainRuleEmailDomainParam struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r DomainRuleEmailDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type EmailListRule struct {
	EmailList EmailListRuleEmailList `json:"email_list,required"`
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

func (r EmailListRule) implementsZeroTrustAccessRule() {}

type EmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                     `json:"id,required"`
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
	EmailList param.Field[EmailListRuleEmailListParam] `json:"email_list,required"`
}

func (r EmailListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailListRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type EmailListRuleEmailListParam struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r EmailListRuleEmailListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
type EmailRule struct {
	Email EmailRuleEmail `json:"email,required"`
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

func (r EmailRule) implementsZeroTrustAccessRule() {}

type EmailRuleEmail struct {
	// The email of the user.
	Email string             `json:"email,required" format:"email"`
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
	Email param.Field[EmailRuleEmailParam] `json:"email,required"`
}

func (r EmailRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type EmailRuleEmailParam struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r EmailRuleEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type EveryoneRule struct {
	// An empty object which matches on all users.
	Everyone EveryoneRuleEveryone `json:"everyone,required"`
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

func (r EveryoneRule) implementsZeroTrustAccessRule() {}

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
	Everyone param.Field[EveryoneRuleEveryoneParam] `json:"everyone,required"`
}

func (r EveryoneRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EveryoneRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

// An empty object which matches on all users.
type EveryoneRuleEveryoneParam struct {
}

func (r EveryoneRuleEveryoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ExternalEvaluationRule struct {
	ExternalEvaluation ExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
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

func (r ExternalEvaluationRule) implementsZeroTrustAccessRule() {}

type ExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                       `json:"keys_url,required"`
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
	ExternalEvaluation param.Field[ExternalEvaluationRuleExternalEvaluationParam] `json:"external_evaluation,required"`
}

func (r ExternalEvaluationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExternalEvaluationRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type ExternalEvaluationRuleExternalEvaluationParam struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ExternalEvaluationRuleExternalEvaluationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type GitHubOrganizationRule struct {
	GitHubOrganization GitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
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

func (r GitHubOrganizationRule) implementsZeroTrustAccessRule() {}

type GitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	IdentityProviderID string `json:"identity_provider_id,required"`
	// The name of the organization.
	Name string `json:"name,required"`
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
	GitHubOrganization param.Field[GitHubOrganizationRuleGitHubOrganizationParam] `json:"github-organization,required"`
}

func (r GitHubOrganizationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GitHubOrganizationRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type GitHubOrganizationRuleGitHubOrganizationParam struct {
	// The ID of your Github identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
	// The name of the team
	Team param.Field[string] `json:"team"`
}

func (r GitHubOrganizationRuleGitHubOrganizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Access group.
type GroupRule struct {
	Group GroupRuleGroup `json:"group,required"`
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

func (r GroupRule) implementsZeroTrustAccessRule() {}

type GroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string             `json:"id,required"`
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
	Group param.Field[GroupRuleGroupParam] `json:"group,required"`
}

func (r GroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type GroupRuleGroupParam struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r GroupRuleGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type GSuiteGroupRule struct {
	GSuite GSuiteGroupRuleGSuite `json:"gsuite,required"`
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

func (r GSuiteGroupRule) implementsZeroTrustAccessRule() {}

type GSuiteGroupRuleGSuite struct {
	// The email of the Google Workspace group.
	Email string `json:"email,required"`
	// The ID of your Google Workspace identity provider.
	IdentityProviderID string                    `json:"identity_provider_id,required"`
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
	GSuite param.Field[GSuiteGroupRuleGSuiteParam] `json:"gsuite,required"`
}

func (r GSuiteGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GSuiteGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type GSuiteGroupRuleGSuiteParam struct {
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
	// The ID of your Google Workspace identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r GSuiteGroupRuleGSuiteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type IPListRule struct {
	IPList IPListRuleIPList `json:"ip_list,required"`
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

func (r IPListRule) implementsZeroTrustAccessRule() {}

type IPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string               `json:"id,required"`
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
	IPList param.Field[IPListRuleIPListParam] `json:"ip_list,required"`
}

func (r IPListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPListRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type IPListRuleIPListParam struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r IPListRuleIPListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address block.
type IPRule struct {
	IP   IPRuleIP   `json:"ip,required"`
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

func (r IPRule) implementsZeroTrustAccessRule() {}

type IPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string       `json:"ip,required"`
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
	IP param.Field[IPRuleIPParam] `json:"ip,required"`
}

func (r IPRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type IPRuleIPParam struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r IPRuleIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type OktaGroupRule struct {
	Okta OktaGroupRuleOkta `json:"okta,required"`
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

func (r OktaGroupRule) implementsZeroTrustAccessRule() {}

type OktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	IdentityProviderID string `json:"identity_provider_id,required"`
	// The name of the Okta group.
	Name string                `json:"name,required"`
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
	Okta param.Field[OktaGroupRuleOktaParam] `json:"okta,required"`
}

func (r OktaGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OktaGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type OktaGroupRuleOktaParam struct {
	// The ID of your Okta identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
	// The name of the Okta group.
	Name param.Field[string] `json:"name,required"`
}

func (r OktaGroupRuleOktaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SAMLGroupRule struct {
	SAML SAMLGroupRuleSAML `json:"saml,required"`
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

func (r SAMLGroupRule) implementsZeroTrustAccessRule() {}

type SAMLGroupRuleSAML struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string `json:"attribute_value,required"`
	// The ID of your SAML identity provider.
	IdentityProviderID string                `json:"identity_provider_id,required"`
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
	SAML param.Field[SAMLGroupRuleSAMLParam] `json:"saml,required"`
}

func (r SAMLGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SAMLGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type SAMLGroupRuleSAMLParam struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
	// The ID of your SAML identity provider.
	IdentityProviderID param.Field[string] `json:"identity_provider_id,required"`
}

func (r SAMLGroupRuleSAMLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ServiceTokenRule struct {
	ServiceToken ServiceTokenRuleServiceToken `json:"service_token,required"`
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

func (r ServiceTokenRule) implementsZeroTrustAccessRule() {}

type ServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                           `json:"token_id,required"`
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
	ServiceToken param.Field[ServiceTokenRuleServiceTokenParam] `json:"service_token,required"`
}

func (r ServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceTokenRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type ServiceTokenRuleServiceTokenParam struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ServiceTokenRuleServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
