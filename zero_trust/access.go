// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AccessService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessService] method instead.
type AccessService struct {
	Options       []option.RequestOption
	Applications  *AccessApplicationService
	Certificates  *AccessCertificateService
	Groups        *AccessGroupService
	ServiceTokens *AccessServiceTokenService
	Bookmarks     *AccessBookmarkService
	Keys          *AccessKeyService
	Logs          *AccessLogService
	Users         *AccessUserService
	CustomPages   *AccessCustomPageService
	Tags          *AccessTagService
}

// NewAccessService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccessService(opts ...option.RequestOption) (r *AccessService) {
	r = &AccessService{}
	r.Options = opts
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
	return
}

// Matches a specific email.
type AccessRule struct {
	Email                interface{}    `json:"email,required"`
	EmailList            interface{}    `json:"email_list,required"`
	EmailDomain          interface{}    `json:"email_domain,required"`
	Everyone             interface{}    `json:"everyone,required"`
	IP                   interface{}    `json:"ip,required"`
	IPList               interface{}    `json:"ip_list,required"`
	Certificate          interface{}    `json:"certificate,required"`
	Group                interface{}    `json:"group,required"`
	AzureAD              interface{}    `json:"azureAD,required"`
	GitHubOrganization   interface{}    `json:"github-organization,required"`
	GSuite               interface{}    `json:"gsuite,required"`
	Okta                 interface{}    `json:"okta,required"`
	SAML                 interface{}    `json:"saml,required"`
	ServiceToken         interface{}    `json:"service_token,required"`
	AnyValidServiceToken interface{}    `json:"any_valid_service_token,required"`
	ExternalEvaluation   interface{}    `json:"external_evaluation,required"`
	Geo                  interface{}    `json:"geo,required"`
	AuthMethod           interface{}    `json:"auth_method,required"`
	DevicePosture        interface{}    `json:"device_posture,required"`
	JSON                 accessRuleJSON `json:"-"`
	union                AccessRuleUnion
}

// accessRuleJSON contains the JSON metadata for the struct [AccessRule]
type accessRuleJSON struct {
	Email                apijson.Field
	EmailList            apijson.Field
	EmailDomain          apijson.Field
	Everyone             apijson.Field
	IP                   apijson.Field
	IPList               apijson.Field
	Certificate          apijson.Field
	Group                apijson.Field
	AzureAD              apijson.Field
	GitHubOrganization   apijson.Field
	GSuite               apijson.Field
	Okta                 apijson.Field
	SAML                 apijson.Field
	ServiceToken         apijson.Field
	AnyValidServiceToken apijson.Field
	ExternalEvaluation   apijson.Field
	Geo                  apijson.Field
	AuthMethod           apijson.Field
	DevicePosture        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r accessRuleJSON) RawJSON() string {
	return r.raw
}

func (r *AccessRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r AccessRule) AsUnion() AccessRuleUnion {
	return r.union
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.AccessRuleAccessEmailRule],
// [zero_trust.AccessRuleAccessEmailListRule],
// [zero_trust.AccessRuleAccessDomainRule],
// [zero_trust.AccessRuleAccessEveryoneRule], [zero_trust.AccessRuleAccessIPRule],
// [zero_trust.AccessRuleAccessIPListRule],
// [zero_trust.AccessRuleAccessCertificateRule],
// [zero_trust.AccessRuleAccessAccessGroupRule],
// [zero_trust.AccessRuleAccessAzureGroupRule],
// [zero_trust.AccessRuleAccessGitHubOrganizationRule],
// [zero_trust.AccessRuleAccessGSuiteGroupRule], [zero_trust.OktaGroupRule],
// [zero_trust.AccessRuleAccessSAMLGroupRule],
// [zero_trust.AccessRuleAccessServiceTokenRule],
// [zero_trust.AccessRuleAccessAnyValidServiceTokenRule],
// [zero_trust.AccessRuleAccessExternalEvaluationRule],
// [zero_trust.AccessRuleAccessCountryRule],
// [zero_trust.AccessRuleAccessAuthenticationMethodRule] or
// [zero_trust.AccessRuleAccessDevicePostureRule].
type AccessRuleUnion interface {
	implementsZeroTrustAccessRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessRuleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessGSuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessSAMLGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type AccessRuleAccessEmailRule struct {
	Email AccessRuleAccessEmailRuleEmail `json:"email,required"`
	JSON  accessRuleAccessEmailRuleJSON  `json:"-"`
}

// accessRuleAccessEmailRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessEmailRule]
type accessRuleAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessEmailRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                             `json:"email,required" format:"email"`
	JSON  accessRuleAccessEmailRuleEmailJSON `json:"-"`
}

// accessRuleAccessEmailRuleEmailJSON contains the JSON metadata for the struct
// [AccessRuleAccessEmailRuleEmail]
type accessRuleAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type AccessRuleAccessEmailListRule struct {
	EmailList AccessRuleAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessRuleAccessEmailListRuleJSON      `json:"-"`
}

// accessRuleAccessEmailListRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessEmailListRule]
type accessRuleAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessEmailListRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                     `json:"id,required"`
	JSON accessRuleAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessRuleAccessEmailListRuleEmailListJSON contains the JSON metadata for the
// struct [AccessRuleAccessEmailListRuleEmailList]
type accessRuleAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type AccessRuleAccessDomainRule struct {
	EmailDomain AccessRuleAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessRuleAccessDomainRuleJSON        `json:"-"`
}

// accessRuleAccessDomainRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessDomainRule]
type accessRuleAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessDomainRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                    `json:"domain,required"`
	JSON   accessRuleAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessRuleAccessDomainRuleEmailDomainJSON contains the JSON metadata for the
// struct [AccessRuleAccessDomainRuleEmailDomain]
type accessRuleAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type AccessRuleAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                      `json:"everyone,required"`
	JSON     accessRuleAccessEveryoneRuleJSON `json:"-"`
}

// accessRuleAccessEveryoneRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessEveryoneRule]
type accessRuleAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessEveryoneRule) implementsZeroTrustAccessRule() {}

// Matches an IP address block.
type AccessRuleAccessIPRule struct {
	IP   AccessRuleAccessIPRuleIP   `json:"ip,required"`
	JSON accessRuleAccessIPRuleJSON `json:"-"`
}

// accessRuleAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPRule]
type accessRuleAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessIPRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                       `json:"ip,required"`
	JSON accessRuleAccessIPRuleIPJSON `json:"-"`
}

// accessRuleAccessIPRuleIPJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPRuleIP]
type accessRuleAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type AccessRuleAccessIPListRule struct {
	IPList AccessRuleAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessRuleAccessIPListRuleJSON   `json:"-"`
}

// accessRuleAccessIPListRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPListRule]
type accessRuleAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessIPListRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                               `json:"id,required"`
	JSON accessRuleAccessIPListRuleIPListJSON `json:"-"`
}

// accessRuleAccessIPListRuleIPListJSON contains the JSON metadata for the struct
// [AccessRuleAccessIPListRuleIPList]
type accessRuleAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type AccessRuleAccessCertificateRule struct {
	Certificate interface{}                         `json:"certificate,required"`
	JSON        accessRuleAccessCertificateRuleJSON `json:"-"`
}

// accessRuleAccessCertificateRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessCertificateRule]
type accessRuleAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessCertificateRule) implementsZeroTrustAccessRule() {}

// Matches an Access group.
type AccessRuleAccessAccessGroupRule struct {
	Group AccessRuleAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessRuleAccessAccessGroupRuleJSON  `json:"-"`
}

// accessRuleAccessAccessGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessAccessGroupRule]
type accessRuleAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAccessGroupRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                   `json:"id,required"`
	JSON accessRuleAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessRuleAccessAccessGroupRuleGroupJSON contains the JSON metadata for the
// struct [AccessRuleAccessAccessGroupRuleGroup]
type accessRuleAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessRuleAccessAzureGroupRule struct {
	AzureAD AccessRuleAccessAzureGroupRuleAzureAD `json:"azureAD,required"`
	JSON    accessRuleAccessAzureGroupRuleJSON    `json:"-"`
}

// accessRuleAccessAzureGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessAzureGroupRule]
type accessRuleAccessAzureGroupRuleJSON struct {
	AzureAD     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAzureGroupRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessAzureGroupRuleAzureAD struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                    `json:"connection_id,required"`
	JSON         accessRuleAccessAzureGroupRuleAzureADJSON `json:"-"`
}

// accessRuleAccessAzureGroupRuleAzureADJSON contains the JSON metadata for the
// struct [AccessRuleAccessAzureGroupRuleAzureAD]
type accessRuleAccessAzureGroupRuleAzureADJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessRuleAccessAzureGroupRuleAzureAD) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAzureGroupRuleAzureADJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type AccessRuleAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessRuleAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessRuleAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessRuleAccessGitHubOrganizationRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessGitHubOrganizationRule]
type accessRuleAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessGitHubOrganizationRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                       `json:"name,required"`
	JSON accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON contains the JSON
// metadata for the struct
// [AccessRuleAccessGitHubOrganizationRuleGitHubOrganization]
type accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessRuleAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessRuleAccessGSuiteGroupRule struct {
	GSuite AccessRuleAccessGSuiteGroupRuleGSuite `json:"gsuite,required"`
	JSON   accessRuleAccessGSuiteGroupRuleJSON   `json:"-"`
}

// accessRuleAccessGSuiteGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessGSuiteGroupRule]
type accessRuleAccessGSuiteGroupRuleJSON struct {
	GSuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessGSuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGSuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessGSuiteGroupRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessGSuiteGroupRuleGSuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                    `json:"email,required"`
	JSON  accessRuleAccessGSuiteGroupRuleGSuiteJSON `json:"-"`
}

// accessRuleAccessGSuiteGroupRuleGSuiteJSON contains the JSON metadata for the
// struct [AccessRuleAccessGSuiteGroupRuleGSuite]
type accessRuleAccessGSuiteGroupRuleGSuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessRuleAccessGSuiteGroupRuleGSuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessGSuiteGroupRuleGSuiteJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessRuleAccessSAMLGroupRule struct {
	SAML AccessRuleAccessSAMLGroupRuleSAML `json:"saml,required"`
	JSON accessRuleAccessSAMLGroupRuleJSON `json:"-"`
}

// accessRuleAccessSAMLGroupRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessSAMLGroupRule]
type accessRuleAccessSAMLGroupRuleJSON struct {
	SAML        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessSAMLGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessSAMLGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessSAMLGroupRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessSAMLGroupRuleSAML struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                `json:"attribute_value,required"`
	JSON           accessRuleAccessSAMLGroupRuleSAMLJSON `json:"-"`
}

// accessRuleAccessSAMLGroupRuleSAMLJSON contains the JSON metadata for the struct
// [AccessRuleAccessSAMLGroupRuleSAML]
type accessRuleAccessSAMLGroupRuleSAMLJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessRuleAccessSAMLGroupRuleSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessSAMLGroupRuleSAMLJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type AccessRuleAccessServiceTokenRule struct {
	ServiceToken AccessRuleAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessRuleAccessServiceTokenRuleJSON         `json:"-"`
}

// accessRuleAccessServiceTokenRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessServiceTokenRule]
type accessRuleAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessRuleAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessServiceTokenRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                           `json:"token_id,required"`
	JSON    accessRuleAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessRuleAccessServiceTokenRuleServiceTokenJSON contains the JSON metadata for
// the struct [AccessRuleAccessServiceTokenRuleServiceToken]
type accessRuleAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type AccessRuleAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                  `json:"any_valid_service_token,required"`
	JSON                 accessRuleAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessRuleAccessAnyValidServiceTokenRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessAnyValidServiceTokenRule]
type accessRuleAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessRuleAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAnyValidServiceTokenRule) implementsZeroTrustAccessRule() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessRuleAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessRuleAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessRuleAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessRuleAccessExternalEvaluationRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessExternalEvaluationRule]
type accessRuleAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessRuleAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessExternalEvaluationRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                       `json:"keys_url,required"`
	JSON    accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON contains the JSON
// metadata for the struct
// [AccessRuleAccessExternalEvaluationRuleExternalEvaluation]
type accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type AccessRuleAccessCountryRule struct {
	Geo  AccessRuleAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessRuleAccessCountryRuleJSON `json:"-"`
}

// accessRuleAccessCountryRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessCountryRule]
type accessRuleAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessCountryRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                             `json:"country_code,required"`
	JSON        accessRuleAccessCountryRuleGeoJSON `json:"-"`
}

// accessRuleAccessCountryRuleGeoJSON contains the JSON metadata for the struct
// [AccessRuleAccessCountryRuleGeo]
type accessRuleAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type AccessRuleAccessAuthenticationMethodRule struct {
	AuthMethod AccessRuleAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessRuleAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessRuleAccessAuthenticationMethodRuleJSON contains the JSON metadata for the
// struct [AccessRuleAccessAuthenticationMethodRule]
type accessRuleAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessAuthenticationMethodRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                 `json:"auth_method,required"`
	JSON       accessRuleAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessRuleAccessAuthenticationMethodRuleAuthMethodJSON contains the JSON
// metadata for the struct [AccessRuleAccessAuthenticationMethodRuleAuthMethod]
type accessRuleAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type AccessRuleAccessDevicePostureRule struct {
	DevicePosture AccessRuleAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessRuleAccessDevicePostureRuleJSON          `json:"-"`
}

// accessRuleAccessDevicePostureRuleJSON contains the JSON metadata for the struct
// [AccessRuleAccessDevicePostureRule]
type accessRuleAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleAccessDevicePostureRule) implementsZeroTrustAccessRule() {}

type AccessRuleAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                             `json:"integration_uid,required"`
	JSON           accessRuleAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessRuleAccessDevicePostureRuleDevicePostureJSON contains the JSON metadata
// for the struct [AccessRuleAccessDevicePostureRuleDevicePosture]
type accessRuleAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessRuleAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
type AccessRuleParam struct {
	Email                param.Field[interface{}] `json:"email,required"`
	EmailList            param.Field[interface{}] `json:"email_list,required"`
	EmailDomain          param.Field[interface{}] `json:"email_domain,required"`
	Everyone             param.Field[interface{}] `json:"everyone,required"`
	IP                   param.Field[interface{}] `json:"ip,required"`
	IPList               param.Field[interface{}] `json:"ip_list,required"`
	Certificate          param.Field[interface{}] `json:"certificate,required"`
	Group                param.Field[interface{}] `json:"group,required"`
	AzureAD              param.Field[interface{}] `json:"azureAD,required"`
	GitHubOrganization   param.Field[interface{}] `json:"github-organization,required"`
	GSuite               param.Field[interface{}] `json:"gsuite,required"`
	Okta                 param.Field[interface{}] `json:"okta,required"`
	SAML                 param.Field[interface{}] `json:"saml,required"`
	ServiceToken         param.Field[interface{}] `json:"service_token,required"`
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
	ExternalEvaluation   param.Field[interface{}] `json:"external_evaluation,required"`
	Geo                  param.Field[interface{}] `json:"geo,required"`
	AuthMethod           param.Field[interface{}] `json:"auth_method,required"`
	DevicePosture        param.Field[interface{}] `json:"device_posture,required"`
}

func (r AccessRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

// Matches a specific email.
//
// Satisfied by [zero_trust.AccessRuleAccessEmailRuleParam],
// [zero_trust.AccessRuleAccessEmailListRuleParam],
// [zero_trust.AccessRuleAccessDomainRuleParam],
// [zero_trust.AccessRuleAccessEveryoneRuleParam],
// [zero_trust.AccessRuleAccessIPRuleParam],
// [zero_trust.AccessRuleAccessIPListRuleParam],
// [zero_trust.AccessRuleAccessCertificateRuleParam],
// [zero_trust.AccessRuleAccessAccessGroupRuleParam],
// [zero_trust.AccessRuleAccessAzureGroupRuleParam],
// [zero_trust.AccessRuleAccessGitHubOrganizationRuleParam],
// [zero_trust.AccessRuleAccessGSuiteGroupRuleParam],
// [zero_trust.OktaGroupRuleParam],
// [zero_trust.AccessRuleAccessSAMLGroupRuleParam],
// [zero_trust.AccessRuleAccessServiceTokenRuleParam],
// [zero_trust.AccessRuleAccessAnyValidServiceTokenRuleParam],
// [zero_trust.AccessRuleAccessExternalEvaluationRuleParam],
// [zero_trust.AccessRuleAccessCountryRuleParam],
// [zero_trust.AccessRuleAccessAuthenticationMethodRuleParam],
// [zero_trust.AccessRuleAccessDevicePostureRuleParam], [AccessRuleParam].
type AccessRuleUnionParam interface {
	implementsZeroTrustAccessRuleUnionParam()
}

// Matches a specific email.
type AccessRuleAccessEmailRuleParam struct {
	Email param.Field[AccessRuleAccessEmailRuleEmailParam] `json:"email,required"`
}

func (r AccessRuleAccessEmailRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessEmailRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessEmailRuleEmailParam struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessRuleAccessEmailRuleEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessRuleAccessEmailListRuleParam struct {
	EmailList param.Field[AccessRuleAccessEmailListRuleEmailListParam] `json:"email_list,required"`
}

func (r AccessRuleAccessEmailListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessEmailListRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessEmailListRuleEmailListParam struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessRuleAccessEmailListRuleEmailListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessRuleAccessDomainRuleParam struct {
	EmailDomain param.Field[AccessRuleAccessDomainRuleEmailDomainParam] `json:"email_domain,required"`
}

func (r AccessRuleAccessDomainRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessDomainRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessDomainRuleEmailDomainParam struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessRuleAccessDomainRuleEmailDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessRuleAccessEveryoneRuleParam struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessRuleAccessEveryoneRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessEveryoneRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

// Matches an IP address block.
type AccessRuleAccessIPRuleParam struct {
	IP param.Field[AccessRuleAccessIPRuleIPParam] `json:"ip,required"`
}

func (r AccessRuleAccessIPRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessIPRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessIPRuleIPParam struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessRuleAccessIPRuleIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessRuleAccessIPListRuleParam struct {
	IPList param.Field[AccessRuleAccessIPListRuleIPListParam] `json:"ip_list,required"`
}

func (r AccessRuleAccessIPListRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessIPListRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessIPListRuleIPListParam struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessRuleAccessIPListRuleIPListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessRuleAccessCertificateRuleParam struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessRuleAccessCertificateRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCertificateRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

// Matches an Access group.
type AccessRuleAccessAccessGroupRuleParam struct {
	Group param.Field[AccessRuleAccessAccessGroupRuleGroupParam] `json:"group,required"`
}

func (r AccessRuleAccessAccessGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAccessGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessAccessGroupRuleGroupParam struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessRuleAccessAccessGroupRuleGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessRuleAccessAzureGroupRuleParam struct {
	AzureAD param.Field[AccessRuleAccessAzureGroupRuleAzureADParam] `json:"azureAD,required"`
}

func (r AccessRuleAccessAzureGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAzureGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessAzureGroupRuleAzureADParam struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessRuleAccessAzureGroupRuleAzureADParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessRuleAccessGitHubOrganizationRuleParam struct {
	GitHubOrganization param.Field[AccessRuleAccessGitHubOrganizationRuleGitHubOrganizationParam] `json:"github-organization,required"`
}

func (r AccessRuleAccessGitHubOrganizationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessGitHubOrganizationRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessGitHubOrganizationRuleGitHubOrganizationParam struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessRuleAccessGitHubOrganizationRuleGitHubOrganizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessRuleAccessGSuiteGroupRuleParam struct {
	GSuite param.Field[AccessRuleAccessGSuiteGroupRuleGSuiteParam] `json:"gsuite,required"`
}

func (r AccessRuleAccessGSuiteGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessGSuiteGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessGSuiteGroupRuleGSuiteParam struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessRuleAccessGSuiteGroupRuleGSuiteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessRuleAccessSAMLGroupRuleParam struct {
	SAML param.Field[AccessRuleAccessSAMLGroupRuleSAMLParam] `json:"saml,required"`
}

func (r AccessRuleAccessSAMLGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessSAMLGroupRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessSAMLGroupRuleSAMLParam struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessRuleAccessSAMLGroupRuleSAMLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessRuleAccessServiceTokenRuleParam struct {
	ServiceToken param.Field[AccessRuleAccessServiceTokenRuleServiceTokenParam] `json:"service_token,required"`
}

func (r AccessRuleAccessServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessServiceTokenRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessServiceTokenRuleServiceTokenParam struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessRuleAccessServiceTokenRuleServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessRuleAccessAnyValidServiceTokenRuleParam struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessRuleAccessAnyValidServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAnyValidServiceTokenRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessRuleAccessExternalEvaluationRuleParam struct {
	ExternalEvaluation param.Field[AccessRuleAccessExternalEvaluationRuleExternalEvaluationParam] `json:"external_evaluation,required"`
}

func (r AccessRuleAccessExternalEvaluationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessExternalEvaluationRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessExternalEvaluationRuleExternalEvaluationParam struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessRuleAccessExternalEvaluationRuleExternalEvaluationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessRuleAccessCountryRuleParam struct {
	Geo param.Field[AccessRuleAccessCountryRuleGeoParam] `json:"geo,required"`
}

func (r AccessRuleAccessCountryRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessCountryRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessCountryRuleGeoParam struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessRuleAccessCountryRuleGeoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessRuleAccessAuthenticationMethodRuleParam struct {
	AuthMethod param.Field[AccessRuleAccessAuthenticationMethodRuleAuthMethodParam] `json:"auth_method,required"`
}

func (r AccessRuleAccessAuthenticationMethodRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessAuthenticationMethodRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessAuthenticationMethodRuleAuthMethodParam struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessRuleAccessAuthenticationMethodRuleAuthMethodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessRuleAccessDevicePostureRuleParam struct {
	DevicePosture param.Field[AccessRuleAccessDevicePostureRuleDevicePostureParam] `json:"device_posture,required"`
}

func (r AccessRuleAccessDevicePostureRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleAccessDevicePostureRuleParam) implementsZeroTrustAccessRuleUnionParam() {}

type AccessRuleAccessDevicePostureRuleDevicePostureParam struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessRuleAccessDevicePostureRuleDevicePostureParam) MarshalJSON() (data []byte, err error) {
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
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                `json:"email,required"`
	JSON  oktaGroupRuleOktaJSON `json:"-"`
}

// oktaGroupRuleOktaJSON contains the JSON metadata for the struct
// [OktaGroupRuleOkta]
type oktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
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
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r OktaGroupRuleOktaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
