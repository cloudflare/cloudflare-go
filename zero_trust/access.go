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

// Matches any valid Access Service Token
type AnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                  `json:"any_valid_service_token,required"`
	JSON                 anyValidServiceTokenRuleJSON `json:"-"`
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

func (r AnyValidServiceTokenRule) implementsZeroTrustIncludeItem() {}

func (r AnyValidServiceTokenRule) implementsZeroTrustRule() {}

func (r AnyValidServiceTokenRule) implementsZeroTrustExcludeItem() {}

func (r AnyValidServiceTokenRule) implementsZeroTrustRequireItem() {}

// Matches any valid Access Service Token
type AnyValidServiceTokenRuleParam struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AnyValidServiceTokenRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AnyValidServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r AuthenticationMethodRule) implementsZeroTrustIncludeItem() {}

func (r AuthenticationMethodRule) implementsZeroTrustRule() {}

func (r AuthenticationMethodRule) implementsZeroTrustExcludeItem() {}

func (r AuthenticationMethodRule) implementsZeroTrustRequireItem() {}

type AuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
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

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AuthenticationMethodRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type AuthenticationMethodRuleAuthMethodParam struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AuthenticationMethodRuleAuthMethodParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AzureGroupRule struct {
	AzureAd AzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    azureGroupRuleJSON    `json:"-"`
}

// azureGroupRuleJSON contains the JSON metadata for the struct [AzureGroupRule]
type azureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AzureGroupRule) implementsZeroTrustIncludeItem() {}

func (r AzureGroupRule) implementsZeroTrustRule() {}

func (r AzureGroupRule) implementsZeroTrustExcludeItem() {}

func (r AzureGroupRule) implementsZeroTrustRequireItem() {}

type AzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                    `json:"connection_id,required"`
	JSON         azureGroupRuleAzureAdJSON `json:"-"`
}

// azureGroupRuleAzureAdJSON contains the JSON metadata for the struct
// [AzureGroupRuleAzureAd]
type azureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type AzureGroupRuleParam struct {
	AzureAd param.Field[AzureGroupRuleAzureAdParam] `json:"azureAD,required"`
}

func (r AzureGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r AzureGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type AzureGroupRuleAzureAdParam struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AzureGroupRuleAzureAdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type CertificateRule struct {
	Certificate interface{}         `json:"certificate,required"`
	JSON        certificateRuleJSON `json:"-"`
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

func (r CertificateRule) implementsZeroTrustIncludeItem() {}

func (r CertificateRule) implementsZeroTrustRule() {}

func (r CertificateRule) implementsZeroTrustExcludeItem() {}

func (r CertificateRule) implementsZeroTrustRequireItem() {}

// Matches any valid client certificate.
type CertificateRuleParam struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r CertificateRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CertificateRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r CountryRule) implementsZeroTrustIncludeItem() {}

func (r CountryRule) implementsZeroTrustRule() {}

func (r CountryRule) implementsZeroTrustExcludeItem() {}

func (r CountryRule) implementsZeroTrustRequireItem() {}

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

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r CountryRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type CountryRuleGeoParam struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r CountryRuleGeoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type DevicePostureRule struct {
	DevicePosture DevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          devicePostureRuleJSON          `json:"-"`
}

// devicePostureRuleJSON contains the JSON metadata for the struct
// [DevicePostureRule]
type devicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r DevicePostureRule) implementsZeroTrustIncludeItem() {}

func (r DevicePostureRule) implementsZeroTrustRule() {}

func (r DevicePostureRule) implementsZeroTrustExcludeItem() {}

func (r DevicePostureRule) implementsZeroTrustRequireItem() {}

type DevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                             `json:"integration_uid,required"`
	JSON           devicePostureRuleDevicePostureJSON `json:"-"`
}

// devicePostureRuleDevicePostureJSON contains the JSON metadata for the struct
// [DevicePostureRuleDevicePosture]
type devicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type DevicePostureRuleParam struct {
	DevicePosture param.Field[DevicePostureRuleDevicePostureParam] `json:"device_posture,required"`
}

func (r DevicePostureRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DevicePostureRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type DevicePostureRuleDevicePostureParam struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r DevicePostureRuleDevicePostureParam) MarshalJSON() (data []byte, err error) {
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

func (r DomainRule) implementsZeroTrustIncludeItem() {}

func (r DomainRule) implementsZeroTrustRule() {}

func (r DomainRule) implementsZeroTrustExcludeItem() {}

func (r DomainRule) implementsZeroTrustRequireItem() {}

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

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r DomainRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r EmailListRule) implementsZeroTrustIncludeItem() {}

func (r EmailListRule) implementsZeroTrustRule() {}

func (r EmailListRule) implementsZeroTrustExcludeItem() {}

func (r EmailListRule) implementsZeroTrustRequireItem() {}

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

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r EmailRule) implementsZeroTrustIncludeItem() {}

func (r EmailRule) implementsZeroTrustRule() {}

func (r EmailRule) implementsZeroTrustExcludeItem() {}

func (r EmailRule) implementsZeroTrustRequireItem() {}

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

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EmailRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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
	Everyone interface{}      `json:"everyone,required"`
	JSON     everyoneRuleJSON `json:"-"`
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

func (r EveryoneRule) implementsZeroTrustIncludeItem() {}

func (r EveryoneRule) implementsZeroTrustRule() {}

func (r EveryoneRule) implementsZeroTrustExcludeItem() {}

func (r EveryoneRule) implementsZeroTrustRequireItem() {}

// Matches everyone.
type EveryoneRuleParam struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r EveryoneRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r EveryoneRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r ExternalEvaluationRule) implementsZeroTrustIncludeItem() {}

func (r ExternalEvaluationRule) implementsZeroTrustRule() {}

func (r ExternalEvaluationRule) implementsZeroTrustExcludeItem() {}

func (r ExternalEvaluationRule) implementsZeroTrustRequireItem() {}

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

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ExternalEvaluationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r GitHubOrganizationRule) implementsZeroTrustIncludeItem() {}

func (r GitHubOrganizationRule) implementsZeroTrustRule() {}

func (r GitHubOrganizationRule) implementsZeroTrustExcludeItem() {}

func (r GitHubOrganizationRule) implementsZeroTrustRequireItem() {}

type GitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                       `json:"name,required"`
	JSON githubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// githubOrganizationRuleGitHubOrganizationJSON contains the JSON metadata for the
// struct [GitHubOrganizationRuleGitHubOrganization]
type githubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
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

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GitHubOrganizationRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type GitHubOrganizationRuleGitHubOrganizationParam struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
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

func (r GroupRule) implementsZeroTrustIncludeItem() {}

func (r GroupRule) implementsZeroTrustRule() {}

func (r GroupRule) implementsZeroTrustExcludeItem() {}

func (r GroupRule) implementsZeroTrustRequireItem() {}

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

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type GroupRuleGroupParam struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r GroupRuleGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type GsuiteGroupRule struct {
	Gsuite GsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   gsuiteGroupRuleJSON   `json:"-"`
}

// gsuiteGroupRuleJSON contains the JSON metadata for the struct [GsuiteGroupRule]
type gsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r GsuiteGroupRule) implementsZeroTrustIncludeItem() {}

func (r GsuiteGroupRule) implementsZeroTrustRule() {}

func (r GsuiteGroupRule) implementsZeroTrustExcludeItem() {}

func (r GsuiteGroupRule) implementsZeroTrustRequireItem() {}

type GsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                    `json:"email,required"`
	JSON  gsuiteGroupRuleGsuiteJSON `json:"-"`
}

// gsuiteGroupRuleGsuiteJSON contains the JSON metadata for the struct
// [GsuiteGroupRuleGsuite]
type gsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type GsuiteGroupRuleParam struct {
	Gsuite param.Field[GsuiteGroupRuleGsuiteParam] `json:"gsuite,required"`
}

func (r GsuiteGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r GsuiteGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type GsuiteGroupRuleGsuiteParam struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r GsuiteGroupRuleGsuiteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
type IncludeItem struct {
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
	JSON                 includeItemJSON `json:"-"`
	union                IncludeItemUnion
}

// includeItemJSON contains the JSON metadata for the struct [IncludeItem]
type includeItemJSON struct {
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

func (r includeItemJSON) RawJSON() string {
	return r.raw
}

func (r *IncludeItem) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r IncludeItem) AsUnion() IncludeItemUnion {
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
type IncludeItemUnion interface {
	implementsZeroTrustIncludeItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IncludeItemUnion)(nil)).Elem(),
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
type IncludeItemParam struct {
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

func (r IncludeItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IncludeItemParam) implementsZeroTrustIncludeItemUnionParam() {}

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
// [IncludeItemParam].
type IncludeItemUnionParam interface {
	implementsZeroTrustIncludeItemUnionParam()
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

func (r IPListRule) implementsZeroTrustIncludeItem() {}

func (r IPListRule) implementsZeroTrustRule() {}

func (r IPListRule) implementsZeroTrustExcludeItem() {}

func (r IPListRule) implementsZeroTrustRequireItem() {}

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

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPListRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r IPRule) implementsZeroTrustIncludeItem() {}

func (r IPRule) implementsZeroTrustRule() {}

func (r IPRule) implementsZeroTrustExcludeItem() {}

func (r IPRule) implementsZeroTrustRequireItem() {}

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

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r IPRuleParam) implementsZeroTrustRequireItemUnionParam() {}

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

func (r OktaGroupRule) implementsZeroTrustIncludeItem() {}

func (r OktaGroupRule) implementsZeroTrustRule() {}

func (r OktaGroupRule) implementsZeroTrustExcludeItem() {}

func (r OktaGroupRule) implementsZeroTrustRequireItem() {}

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

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r OktaGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type OktaGroupRuleOktaParam struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r OktaGroupRuleOktaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
type Rule struct {
	Email                interface{} `json:"email,required"`
	EmailList            interface{} `json:"email_list,required"`
	EmailDomain          interface{} `json:"email_domain,required"`
	Everyone             interface{} `json:"everyone,required"`
	IP                   interface{} `json:"ip,required"`
	IPList               interface{} `json:"ip_list,required"`
	Certificate          interface{} `json:"certificate,required"`
	Group                interface{} `json:"group,required"`
	AzureAd              interface{} `json:"azureAD,required"`
	GitHubOrganization   interface{} `json:"github-organization,required"`
	Gsuite               interface{} `json:"gsuite,required"`
	Okta                 interface{} `json:"okta,required"`
	Saml                 interface{} `json:"saml,required"`
	ServiceToken         interface{} `json:"service_token,required"`
	AnyValidServiceToken interface{} `json:"any_valid_service_token,required"`
	ExternalEvaluation   interface{} `json:"external_evaluation,required"`
	Geo                  interface{} `json:"geo,required"`
	AuthMethod           interface{} `json:"auth_method,required"`
	DevicePosture        interface{} `json:"device_posture,required"`
	JSON                 ruleJSON    `json:"-"`
	union                RuleUnion
}

// ruleJSON contains the JSON metadata for the struct [Rule]
type ruleJSON struct {
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

func (r ruleJSON) RawJSON() string {
	return r.raw
}

func (r *Rule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Rule) AsUnion() RuleUnion {
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
type RuleUnion interface {
	implementsZeroTrustRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleUnion)(nil)).Elem(),
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
type RuleParam struct {
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

func (r RuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleParam) implementsZeroTrustRuleUnionParam() {}

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
// [RuleParam].
type RuleUnionParam interface {
	implementsZeroTrustRuleUnionParam()
}

// Matches a SAML group. Requires a SAML identity provider.
type SamlGroupRule struct {
	Saml SamlGroupRuleSaml `json:"saml,required"`
	JSON samlGroupRuleJSON `json:"-"`
}

// samlGroupRuleJSON contains the JSON metadata for the struct [SamlGroupRule]
type samlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r SamlGroupRule) implementsZeroTrustIncludeItem() {}

func (r SamlGroupRule) implementsZeroTrustRule() {}

func (r SamlGroupRule) implementsZeroTrustExcludeItem() {}

func (r SamlGroupRule) implementsZeroTrustRequireItem() {}

type SamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                `json:"attribute_value,required"`
	JSON           samlGroupRuleSamlJSON `json:"-"`
}

// samlGroupRuleSamlJSON contains the JSON metadata for the struct
// [SamlGroupRuleSaml]
type samlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r samlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type SamlGroupRuleParam struct {
	Saml param.Field[SamlGroupRuleSamlParam] `json:"saml,required"`
}

func (r SamlGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r SamlGroupRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type SamlGroupRuleSamlParam struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r SamlGroupRuleSamlParam) MarshalJSON() (data []byte, err error) {
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

func (r ServiceTokenRule) implementsZeroTrustIncludeItem() {}

func (r ServiceTokenRule) implementsZeroTrustRule() {}

func (r ServiceTokenRule) implementsZeroTrustExcludeItem() {}

func (r ServiceTokenRule) implementsZeroTrustRequireItem() {}

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

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRuleUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustExcludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustIncludeItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

func (r ServiceTokenRuleParam) implementsZeroTrustRequireItemUnionParam() {}

type ServiceTokenRuleServiceTokenParam struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ServiceTokenRuleServiceTokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
