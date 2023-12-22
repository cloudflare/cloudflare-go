// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAccessGroupService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAccessGroupService] method
// instead.
type ZoneAccessGroupService struct {
	Options []option.RequestOption
}

// NewZoneAccessGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAccessGroupService(opts ...option.RequestOption) (r *ZoneAccessGroupService) {
	r = &ZoneAccessGroupService{}
	r.Options = opts
	return
}

// Creates a new Access group.
func (r *ZoneAccessGroupService) New(ctx context.Context, identifier string, body ZoneAccessGroupNewParams, opts ...option.RequestOption) (res *SingleResponseLef4sow9, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Access group.
func (r *ZoneAccessGroupService) Get(ctx context.Context, identifier string, uuid interface{}, opts ...option.RequestOption) (res *SingleResponseLef4sow9, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Access group.
func (r *ZoneAccessGroupService) Update(ctx context.Context, identifier string, uuid interface{}, body ZoneAccessGroupUpdateParams, opts ...option.RequestOption) (res *SingleResponseLef4sow9, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Access groups.
func (r *ZoneAccessGroupService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SchemasResponseCollectionF4F9gHlN, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Access group.
func (r *ZoneAccessGroupService) Delete(ctx context.Context, identifier string, uuid interface{}, opts ...option.RequestOption) (res *ZoneAccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneAccessGroupDeleteResponse struct {
	Errors   []ZoneAccessGroupDeleteResponseError   `json:"errors"`
	Messages []ZoneAccessGroupDeleteResponseMessage `json:"messages"`
	Result   ZoneAccessGroupDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessGroupDeleteResponseSuccess `json:"success"`
	JSON    zoneAccessGroupDeleteResponseJSON    `json:"-"`
}

// zoneAccessGroupDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneAccessGroupDeleteResponse]
type zoneAccessGroupDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneAccessGroupDeleteResponseErrorJSON `json:"-"`
}

// zoneAccessGroupDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneAccessGroupDeleteResponseError]
type zoneAccessGroupDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneAccessGroupDeleteResponseMessageJSON `json:"-"`
}

// zoneAccessGroupDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneAccessGroupDeleteResponseMessage]
type zoneAccessGroupDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupDeleteResponseResult struct {
	// The unique identifier for the Access group.
	ID   interface{}                             `json:"id"`
	JSON zoneAccessGroupDeleteResponseResultJSON `json:"-"`
}

// zoneAccessGroupDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneAccessGroupDeleteResponseResult]
type zoneAccessGroupDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessGroupDeleteResponseSuccess bool

const (
	ZoneAccessGroupDeleteResponseSuccessTrue ZoneAccessGroupDeleteResponseSuccess = true
)

type ZoneAccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]ZoneAccessGroupNewParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ZoneAccessGroupNewParamsExclude] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]ZoneAccessGroupNewParamsRequire] `json:"require"`
}

func (r ZoneAccessGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupNewParamsIncludeEmailRule],
// [ZoneAccessGroupNewParamsIncludeDomainRule],
// [ZoneAccessGroupNewParamsIncludeEveryoneRule],
// [ZoneAccessGroupNewParamsIncludeIPRule],
// [ZoneAccessGroupNewParamsIncludeIPListRule],
// [ZoneAccessGroupNewParamsIncludeCertificateRule],
// [ZoneAccessGroupNewParamsIncludeAccessGroupRule],
// [ZoneAccessGroupNewParamsIncludeAzureGroupRule],
// [ZoneAccessGroupNewParamsIncludeGitHubOrganizationRule],
// [ZoneAccessGroupNewParamsIncludeGsuiteGroupRule],
// [ZoneAccessGroupNewParamsIncludeOktaGroupRule],
// [ZoneAccessGroupNewParamsIncludeSamlGroupRule].
type ZoneAccessGroupNewParamsInclude interface {
	implementsZoneAccessGroupNewParamsInclude()
}

// Matches a specific email.
type ZoneAccessGroupNewParamsIncludeEmailRule struct {
	Email param.Field[ZoneAccessGroupNewParamsIncludeEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsIncludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeEmailRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupNewParamsIncludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupNewParamsIncludeDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupNewParamsIncludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupNewParamsIncludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeDomainRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupNewParamsIncludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupNewParamsIncludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupNewParamsIncludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeEveryoneRule) implementsZoneAccessGroupNewParamsInclude() {}

// Matches an IP address block.
type ZoneAccessGroupNewParamsIncludeIPRule struct {
	IP param.Field[ZoneAccessGroupNewParamsIncludeIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsIncludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeIPRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsIncludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewParamsIncludeIPListRule struct {
	IPList param.Field[ZoneAccessGroupNewParamsIncludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupNewParamsIncludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeIPListRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsIncludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewParamsIncludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupNewParamsIncludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeCertificateRule) implementsZoneAccessGroupNewParamsInclude() {}

// Matches an Access group.
type ZoneAccessGroupNewParamsIncludeAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupNewParamsIncludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupNewParamsIncludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeAccessGroupRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsIncludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewParamsIncludeAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupNewParamsIncludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupNewParamsIncludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeAzureGroupRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupNewParamsIncludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewParamsIncludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupNewParamsIncludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupNewParamsIncludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeGitHubOrganizationRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupNewParamsIncludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewParamsIncludeGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupNewParamsIncludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupNewParamsIncludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeGsuiteGroupRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsIncludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewParamsIncludeOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupNewParamsIncludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupNewParamsIncludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeOktaGroupRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsIncludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewParamsIncludeSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupNewParamsIncludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupNewParamsIncludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludeSamlGroupRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupNewParamsIncludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupNewParamsExcludeEmailRule],
// [ZoneAccessGroupNewParamsExcludeDomainRule],
// [ZoneAccessGroupNewParamsExcludeEveryoneRule],
// [ZoneAccessGroupNewParamsExcludeIPRule],
// [ZoneAccessGroupNewParamsExcludeIPListRule],
// [ZoneAccessGroupNewParamsExcludeCertificateRule],
// [ZoneAccessGroupNewParamsExcludeAccessGroupRule],
// [ZoneAccessGroupNewParamsExcludeAzureGroupRule],
// [ZoneAccessGroupNewParamsExcludeGitHubOrganizationRule],
// [ZoneAccessGroupNewParamsExcludeGsuiteGroupRule],
// [ZoneAccessGroupNewParamsExcludeOktaGroupRule],
// [ZoneAccessGroupNewParamsExcludeSamlGroupRule].
type ZoneAccessGroupNewParamsExclude interface {
	implementsZoneAccessGroupNewParamsExclude()
}

// Matches a specific email.
type ZoneAccessGroupNewParamsExcludeEmailRule struct {
	Email param.Field[ZoneAccessGroupNewParamsExcludeEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsExcludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeEmailRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupNewParamsExcludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupNewParamsExcludeDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupNewParamsExcludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupNewParamsExcludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeDomainRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupNewParamsExcludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupNewParamsExcludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupNewParamsExcludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeEveryoneRule) implementsZoneAccessGroupNewParamsExclude() {}

// Matches an IP address block.
type ZoneAccessGroupNewParamsExcludeIPRule struct {
	IP param.Field[ZoneAccessGroupNewParamsExcludeIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsExcludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeIPRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsExcludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewParamsExcludeIPListRule struct {
	IPList param.Field[ZoneAccessGroupNewParamsExcludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupNewParamsExcludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeIPListRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsExcludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewParamsExcludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupNewParamsExcludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeCertificateRule) implementsZoneAccessGroupNewParamsExclude() {}

// Matches an Access group.
type ZoneAccessGroupNewParamsExcludeAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupNewParamsExcludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupNewParamsExcludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeAccessGroupRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsExcludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewParamsExcludeAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupNewParamsExcludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupNewParamsExcludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeAzureGroupRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupNewParamsExcludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewParamsExcludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupNewParamsExcludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupNewParamsExcludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeGitHubOrganizationRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupNewParamsExcludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewParamsExcludeGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupNewParamsExcludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupNewParamsExcludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeGsuiteGroupRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsExcludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewParamsExcludeOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupNewParamsExcludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupNewParamsExcludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeOktaGroupRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsExcludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewParamsExcludeSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupNewParamsExcludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupNewParamsExcludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludeSamlGroupRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupNewParamsExcludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupNewParamsRequireEmailRule],
// [ZoneAccessGroupNewParamsRequireDomainRule],
// [ZoneAccessGroupNewParamsRequireEveryoneRule],
// [ZoneAccessGroupNewParamsRequireIPRule],
// [ZoneAccessGroupNewParamsRequireIPListRule],
// [ZoneAccessGroupNewParamsRequireCertificateRule],
// [ZoneAccessGroupNewParamsRequireAccessGroupRule],
// [ZoneAccessGroupNewParamsRequireAzureGroupRule],
// [ZoneAccessGroupNewParamsRequireGitHubOrganizationRule],
// [ZoneAccessGroupNewParamsRequireGsuiteGroupRule],
// [ZoneAccessGroupNewParamsRequireOktaGroupRule],
// [ZoneAccessGroupNewParamsRequireSamlGroupRule].
type ZoneAccessGroupNewParamsRequire interface {
	implementsZoneAccessGroupNewParamsRequire()
}

// Matches a specific email.
type ZoneAccessGroupNewParamsRequireEmailRule struct {
	Email param.Field[ZoneAccessGroupNewParamsRequireEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsRequireEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireEmailRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupNewParamsRequireEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupNewParamsRequireDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupNewParamsRequireDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupNewParamsRequireDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireDomainRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupNewParamsRequireDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupNewParamsRequireEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupNewParamsRequireEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireEveryoneRule) implementsZoneAccessGroupNewParamsRequire() {}

// Matches an IP address block.
type ZoneAccessGroupNewParamsRequireIPRule struct {
	IP param.Field[ZoneAccessGroupNewParamsRequireIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsRequireIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireIPRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsRequireIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewParamsRequireIPListRule struct {
	IPList param.Field[ZoneAccessGroupNewParamsRequireIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupNewParamsRequireIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireIPListRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsRequireIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewParamsRequireCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupNewParamsRequireCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireCertificateRule) implementsZoneAccessGroupNewParamsRequire() {}

// Matches an Access group.
type ZoneAccessGroupNewParamsRequireAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupNewParamsRequireAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupNewParamsRequireAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireAccessGroupRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsRequireAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewParamsRequireAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupNewParamsRequireAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupNewParamsRequireAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireAzureGroupRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupNewParamsRequireAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewParamsRequireGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupNewParamsRequireGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupNewParamsRequireGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireGitHubOrganizationRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequireGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupNewParamsRequireGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewParamsRequireGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupNewParamsRequireGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupNewParamsRequireGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireGsuiteGroupRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsRequireGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewParamsRequireOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupNewParamsRequireOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupNewParamsRequireOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireOktaGroupRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsRequireOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewParamsRequireSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupNewParamsRequireSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupNewParamsRequireSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequireSamlGroupRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequireSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupNewParamsRequireSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]ZoneAccessGroupUpdateParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ZoneAccessGroupUpdateParamsExclude] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]ZoneAccessGroupUpdateParamsRequire] `json:"require"`
}

func (r ZoneAccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupUpdateParamsIncludeEmailRule],
// [ZoneAccessGroupUpdateParamsIncludeDomainRule],
// [ZoneAccessGroupUpdateParamsIncludeEveryoneRule],
// [ZoneAccessGroupUpdateParamsIncludeIPRule],
// [ZoneAccessGroupUpdateParamsIncludeIPListRule],
// [ZoneAccessGroupUpdateParamsIncludeCertificateRule],
// [ZoneAccessGroupUpdateParamsIncludeAccessGroupRule],
// [ZoneAccessGroupUpdateParamsIncludeAzureGroupRule],
// [ZoneAccessGroupUpdateParamsIncludeGitHubOrganizationRule],
// [ZoneAccessGroupUpdateParamsIncludeGsuiteGroupRule],
// [ZoneAccessGroupUpdateParamsIncludeOktaGroupRule],
// [ZoneAccessGroupUpdateParamsIncludeSamlGroupRule].
type ZoneAccessGroupUpdateParamsInclude interface {
	implementsZoneAccessGroupUpdateParamsInclude()
}

// Matches a specific email.
type ZoneAccessGroupUpdateParamsIncludeEmailRule struct {
	Email param.Field[ZoneAccessGroupUpdateParamsIncludeEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeEmailRule) implementsZoneAccessGroupUpdateParamsInclude() {}

type ZoneAccessGroupUpdateParamsIncludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupUpdateParamsIncludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateParamsIncludeDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupUpdateParamsIncludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeDomainRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupUpdateParamsIncludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeEveryoneRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateParamsIncludeIPRule struct {
	IP param.Field[ZoneAccessGroupUpdateParamsIncludeIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeIPRule) implementsZoneAccessGroupUpdateParamsInclude() {}

type ZoneAccessGroupUpdateParamsIncludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateParamsIncludeIPListRule struct {
	IPList param.Field[ZoneAccessGroupUpdateParamsIncludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeIPListRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateParamsIncludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeCertificateRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateParamsIncludeAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupUpdateParamsIncludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeAccessGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateParamsIncludeAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupUpdateParamsIncludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeAzureGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateParamsIncludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupUpdateParamsIncludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeGitHubOrganizationRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateParamsIncludeGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupUpdateParamsIncludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeGsuiteGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateParamsIncludeOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupUpdateParamsIncludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeOktaGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateParamsIncludeSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupUpdateParamsIncludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludeSamlGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupUpdateParamsExcludeEmailRule],
// [ZoneAccessGroupUpdateParamsExcludeDomainRule],
// [ZoneAccessGroupUpdateParamsExcludeEveryoneRule],
// [ZoneAccessGroupUpdateParamsExcludeIPRule],
// [ZoneAccessGroupUpdateParamsExcludeIPListRule],
// [ZoneAccessGroupUpdateParamsExcludeCertificateRule],
// [ZoneAccessGroupUpdateParamsExcludeAccessGroupRule],
// [ZoneAccessGroupUpdateParamsExcludeAzureGroupRule],
// [ZoneAccessGroupUpdateParamsExcludeGitHubOrganizationRule],
// [ZoneAccessGroupUpdateParamsExcludeGsuiteGroupRule],
// [ZoneAccessGroupUpdateParamsExcludeOktaGroupRule],
// [ZoneAccessGroupUpdateParamsExcludeSamlGroupRule].
type ZoneAccessGroupUpdateParamsExclude interface {
	implementsZoneAccessGroupUpdateParamsExclude()
}

// Matches a specific email.
type ZoneAccessGroupUpdateParamsExcludeEmailRule struct {
	Email param.Field[ZoneAccessGroupUpdateParamsExcludeEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeEmailRule) implementsZoneAccessGroupUpdateParamsExclude() {}

type ZoneAccessGroupUpdateParamsExcludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupUpdateParamsExcludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateParamsExcludeDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupUpdateParamsExcludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeDomainRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupUpdateParamsExcludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeEveryoneRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateParamsExcludeIPRule struct {
	IP param.Field[ZoneAccessGroupUpdateParamsExcludeIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeIPRule) implementsZoneAccessGroupUpdateParamsExclude() {}

type ZoneAccessGroupUpdateParamsExcludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateParamsExcludeIPListRule struct {
	IPList param.Field[ZoneAccessGroupUpdateParamsExcludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeIPListRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateParamsExcludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeCertificateRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateParamsExcludeAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupUpdateParamsExcludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeAccessGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateParamsExcludeAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupUpdateParamsExcludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeAzureGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateParamsExcludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupUpdateParamsExcludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeGitHubOrganizationRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateParamsExcludeGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupUpdateParamsExcludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeGsuiteGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateParamsExcludeOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupUpdateParamsExcludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeOktaGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateParamsExcludeSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupUpdateParamsExcludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludeSamlGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupUpdateParamsRequireEmailRule],
// [ZoneAccessGroupUpdateParamsRequireDomainRule],
// [ZoneAccessGroupUpdateParamsRequireEveryoneRule],
// [ZoneAccessGroupUpdateParamsRequireIPRule],
// [ZoneAccessGroupUpdateParamsRequireIPListRule],
// [ZoneAccessGroupUpdateParamsRequireCertificateRule],
// [ZoneAccessGroupUpdateParamsRequireAccessGroupRule],
// [ZoneAccessGroupUpdateParamsRequireAzureGroupRule],
// [ZoneAccessGroupUpdateParamsRequireGitHubOrganizationRule],
// [ZoneAccessGroupUpdateParamsRequireGsuiteGroupRule],
// [ZoneAccessGroupUpdateParamsRequireOktaGroupRule],
// [ZoneAccessGroupUpdateParamsRequireSamlGroupRule].
type ZoneAccessGroupUpdateParamsRequire interface {
	implementsZoneAccessGroupUpdateParamsRequire()
}

// Matches a specific email.
type ZoneAccessGroupUpdateParamsRequireEmailRule struct {
	Email param.Field[ZoneAccessGroupUpdateParamsRequireEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireEmailRule) implementsZoneAccessGroupUpdateParamsRequire() {}

type ZoneAccessGroupUpdateParamsRequireEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupUpdateParamsRequireEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateParamsRequireDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupUpdateParamsRequireDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireDomainRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupUpdateParamsRequireEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireEveryoneRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateParamsRequireIPRule struct {
	IP param.Field[ZoneAccessGroupUpdateParamsRequireIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireIPRule) implementsZoneAccessGroupUpdateParamsRequire() {}

type ZoneAccessGroupUpdateParamsRequireIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateParamsRequireIPListRule struct {
	IPList param.Field[ZoneAccessGroupUpdateParamsRequireIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireIPListRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateParamsRequireCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireCertificateRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateParamsRequireAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupUpdateParamsRequireAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireAccessGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateParamsRequireAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupUpdateParamsRequireAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireAzureGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateParamsRequireGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupUpdateParamsRequireGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireGitHubOrganizationRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateParamsRequireGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupUpdateParamsRequireGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireGsuiteGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateParamsRequireOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupUpdateParamsRequireOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireOktaGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateParamsRequireSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupUpdateParamsRequireSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequireSamlGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequireSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupUpdateParamsRequireSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
