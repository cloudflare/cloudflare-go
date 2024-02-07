// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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

// Fetches a single Access group.
func (r *AccessGroupService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Access group.
func (r *AccessGroupService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessGroupUpdateParams, opts ...option.RequestOption) (res *AccessGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an Access group.
func (r *AccessGroupService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new Access group.
func (r *AccessGroupService) AccessGroupsNewAnAccessGroup(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessGroupAccessGroupsNewAnAccessGroupParams, opts ...option.RequestOption) (res *AccessGroupAccessGroupsNewAnAccessGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access groups.
func (r *AccessGroupService) AccessGroupsListAccessGroups(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessGroupAccessGroupsListAccessGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupAccessGroupsListAccessGroupsResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessGroupGetResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupGetResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupGetResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupGetResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupGetResponseRequire `json:"require"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      accessGroupGetResponseJSON      `json:"-"`
}

// accessGroupGetResponseJSON contains the JSON metadata for the struct
// [AccessGroupGetResponse]
type accessGroupGetResponseJSON struct {
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

func (r *AccessGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseExcludeAccessEmailRule],
// [AccessGroupGetResponseExcludeAccessEmailListRule],
// [AccessGroupGetResponseExcludeAccessDomainRule],
// [AccessGroupGetResponseExcludeAccessEveryoneRule],
// [AccessGroupGetResponseExcludeAccessIPRule],
// [AccessGroupGetResponseExcludeAccessIPListRule],
// [AccessGroupGetResponseExcludeAccessCertificateRule],
// [AccessGroupGetResponseExcludeAccessAccessGroupRule],
// [AccessGroupGetResponseExcludeAccessAzureGroupRule],
// [AccessGroupGetResponseExcludeAccessGitHubOrganizationRule],
// [AccessGroupGetResponseExcludeAccessGsuiteGroupRule],
// [AccessGroupGetResponseExcludeAccessOktaGroupRule],
// [AccessGroupGetResponseExcludeAccessSamlGroupRule],
// [AccessGroupGetResponseExcludeAccessServiceTokenRule],
// [AccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseExcludeAccessExternalEvaluationRule],
// [AccessGroupGetResponseExcludeAccessCountryRule],
// [AccessGroupGetResponseExcludeAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseExcludeAccessDevicePostureRule].
type AccessGroupGetResponseExclude interface {
	implementsAccessGroupGetResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseExcludeAccessEmailRule struct {
	Email AccessGroupGetResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseExcludeAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseExcludeAccessEmailRule]
type accessGroupGetResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessEmailRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessEmailRuleEmailJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessEmailRuleEmail]
type accessGroupGetResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseExcludeAccessEmailListRule struct {
	EmailList AccessGroupGetResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseExcludeAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessEmailListRule]
type accessGroupGetResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessEmailListRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupGetResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseExcludeAccessEmailListRuleEmailList]
type accessGroupGetResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseExcludeAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseExcludeAccessDomainRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseExcludeAccessDomainRule]
type accessGroupGetResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessDomainRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   accessGroupGetResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseExcludeAccessDomainRuleEmailDomain]
type accessGroupGetResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     accessGroupGetResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessEveryoneRule]
type accessGroupGetResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessEveryoneRule) implementsAccessGroupGetResponseExclude() {}

// Matches an IP address block.
type AccessGroupGetResponseExcludeAccessIPRule struct {
	IP   AccessGroupGetResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessIPRuleJSON contains the JSON metadata for the
// struct [AccessGroupGetResponseExcludeAccessIPRule]
type accessGroupGetResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessIPRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON accessGroupGetResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseExcludeAccessIPRuleIP]
type accessGroupGetResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseExcludeAccessIPListRule struct {
	IPList AccessGroupGetResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseExcludeAccessIPListRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseExcludeAccessIPListRule]
type accessGroupGetResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessIPListRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON accessGroupGetResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessIPListRuleIPList]
type accessGroupGetResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        accessGroupGetResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessCertificateRule]
type accessGroupGetResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessCertificateRule) implementsAccessGroupGetResponseExclude() {
}

// Matches an Access group.
type AccessGroupGetResponseExcludeAccessAccessGroupRule struct {
	Group AccessGroupGetResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessAccessGroupRule]
type accessGroupGetResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessAccessGroupRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupGetResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseExcludeAccessAccessGroupRuleGroup]
type accessGroupGetResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseExcludeAccessAzureGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessAzureGroupRule]
type accessGroupGetResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessAzureGroupRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         accessGroupGetResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseExcludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseExcludeAccessGitHubOrganizationRule]
type accessGroupGetResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessGitHubOrganizationRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON accessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessGsuiteGroupRule]
type accessGroupGetResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessGsuiteGroupRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseExcludeAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessOktaGroupRule]
type accessGroupGetResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessOktaGroupRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  accessGroupGetResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessOktaGroupRuleOkta]
type accessGroupGetResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseExcludeAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessSamlGroupRule]
type accessGroupGetResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessSamlGroupRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           accessGroupGetResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessSamlGroupRuleSaml]
type accessGroupGetResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessServiceTokenRule]
type accessGroupGetResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessServiceTokenRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                              `json:"token_id,required"`
	JSON    accessGroupGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                     `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule]
type accessGroupGetResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseExcludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseExcludeAccessExternalEvaluationRule]
type accessGroupGetResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessExternalEvaluationRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                          `json:"keys_url,required"`
	JSON    accessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseExcludeAccessCountryRule struct {
	Geo  AccessGroupGetResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessCountryRule]
type accessGroupGetResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessCountryRule) implementsAccessGroupGetResponseExclude() {}

type AccessGroupGetResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                `json:"country_code,required"`
	JSON        accessGroupGetResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessCountryRuleGeoJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseExcludeAccessCountryRuleGeo]
type accessGroupGetResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseExcludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseExcludeAccessAuthenticationMethodRule]
type accessGroupGetResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessAuthenticationMethodRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                    `json:"auth_method,required"`
	JSON       accessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseExcludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseExcludeAccessDevicePostureRule]
type accessGroupGetResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseExcludeAccessDevicePostureRule) implementsAccessGroupGetResponseExclude() {
}

type AccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                `json:"integration_uid,required"`
	JSON           accessGroupGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseIncludeAccessEmailRule],
// [AccessGroupGetResponseIncludeAccessEmailListRule],
// [AccessGroupGetResponseIncludeAccessDomainRule],
// [AccessGroupGetResponseIncludeAccessEveryoneRule],
// [AccessGroupGetResponseIncludeAccessIPRule],
// [AccessGroupGetResponseIncludeAccessIPListRule],
// [AccessGroupGetResponseIncludeAccessCertificateRule],
// [AccessGroupGetResponseIncludeAccessAccessGroupRule],
// [AccessGroupGetResponseIncludeAccessAzureGroupRule],
// [AccessGroupGetResponseIncludeAccessGitHubOrganizationRule],
// [AccessGroupGetResponseIncludeAccessGsuiteGroupRule],
// [AccessGroupGetResponseIncludeAccessOktaGroupRule],
// [AccessGroupGetResponseIncludeAccessSamlGroupRule],
// [AccessGroupGetResponseIncludeAccessServiceTokenRule],
// [AccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseIncludeAccessExternalEvaluationRule],
// [AccessGroupGetResponseIncludeAccessCountryRule],
// [AccessGroupGetResponseIncludeAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseIncludeAccessDevicePostureRule].
type AccessGroupGetResponseInclude interface {
	implementsAccessGroupGetResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseIncludeAccessEmailRule struct {
	Email AccessGroupGetResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseIncludeAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseIncludeAccessEmailRule]
type accessGroupGetResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessEmailRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessEmailRuleEmailJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessEmailRuleEmail]
type accessGroupGetResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseIncludeAccessEmailListRule struct {
	EmailList AccessGroupGetResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseIncludeAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessEmailListRule]
type accessGroupGetResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessEmailListRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupGetResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIncludeAccessEmailListRuleEmailList]
type accessGroupGetResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseIncludeAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseIncludeAccessDomainRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseIncludeAccessDomainRule]
type accessGroupGetResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessDomainRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   accessGroupGetResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIncludeAccessDomainRuleEmailDomain]
type accessGroupGetResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     accessGroupGetResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessEveryoneRule]
type accessGroupGetResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessEveryoneRule) implementsAccessGroupGetResponseInclude() {}

// Matches an IP address block.
type AccessGroupGetResponseIncludeAccessIPRule struct {
	IP   AccessGroupGetResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessIPRuleJSON contains the JSON metadata for the
// struct [AccessGroupGetResponseIncludeAccessIPRule]
type accessGroupGetResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessIPRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON accessGroupGetResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseIncludeAccessIPRuleIP]
type accessGroupGetResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseIncludeAccessIPListRule struct {
	IPList AccessGroupGetResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseIncludeAccessIPListRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseIncludeAccessIPListRule]
type accessGroupGetResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessIPListRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON accessGroupGetResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessIPListRuleIPList]
type accessGroupGetResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        accessGroupGetResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessCertificateRule]
type accessGroupGetResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessCertificateRule) implementsAccessGroupGetResponseInclude() {
}

// Matches an Access group.
type AccessGroupGetResponseIncludeAccessAccessGroupRule struct {
	Group AccessGroupGetResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessAccessGroupRule]
type accessGroupGetResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessAccessGroupRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupGetResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIncludeAccessAccessGroupRuleGroup]
type accessGroupGetResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseIncludeAccessAzureGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessAzureGroupRule]
type accessGroupGetResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessAzureGroupRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         accessGroupGetResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseIncludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIncludeAccessGitHubOrganizationRule]
type accessGroupGetResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessGitHubOrganizationRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON accessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessGsuiteGroupRule]
type accessGroupGetResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessGsuiteGroupRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseIncludeAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessOktaGroupRule]
type accessGroupGetResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessOktaGroupRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  accessGroupGetResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessOktaGroupRuleOkta]
type accessGroupGetResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseIncludeAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessSamlGroupRule]
type accessGroupGetResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessSamlGroupRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           accessGroupGetResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessSamlGroupRuleSaml]
type accessGroupGetResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessServiceTokenRule]
type accessGroupGetResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessServiceTokenRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                              `json:"token_id,required"`
	JSON    accessGroupGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                     `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule]
type accessGroupGetResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseIncludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIncludeAccessExternalEvaluationRule]
type accessGroupGetResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessExternalEvaluationRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                          `json:"keys_url,required"`
	JSON    accessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseIncludeAccessCountryRule struct {
	Geo  AccessGroupGetResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessCountryRule]
type accessGroupGetResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessCountryRule) implementsAccessGroupGetResponseInclude() {}

type AccessGroupGetResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                `json:"country_code,required"`
	JSON        accessGroupGetResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessCountryRuleGeoJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIncludeAccessCountryRuleGeo]
type accessGroupGetResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseIncludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIncludeAccessAuthenticationMethodRule]
type accessGroupGetResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessAuthenticationMethodRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                    `json:"auth_method,required"`
	JSON       accessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseIncludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIncludeAccessDevicePostureRule]
type accessGroupGetResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIncludeAccessDevicePostureRule) implementsAccessGroupGetResponseInclude() {
}

type AccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                `json:"integration_uid,required"`
	JSON           accessGroupGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseIsDefaultAccessEmailRule],
// [AccessGroupGetResponseIsDefaultAccessEmailListRule],
// [AccessGroupGetResponseIsDefaultAccessDomainRule],
// [AccessGroupGetResponseIsDefaultAccessEveryoneRule],
// [AccessGroupGetResponseIsDefaultAccessIPRule],
// [AccessGroupGetResponseIsDefaultAccessIPListRule],
// [AccessGroupGetResponseIsDefaultAccessCertificateRule],
// [AccessGroupGetResponseIsDefaultAccessAccessGroupRule],
// [AccessGroupGetResponseIsDefaultAccessAzureGroupRule],
// [AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupGetResponseIsDefaultAccessGsuiteGroupRule],
// [AccessGroupGetResponseIsDefaultAccessOktaGroupRule],
// [AccessGroupGetResponseIsDefaultAccessSamlGroupRule],
// [AccessGroupGetResponseIsDefaultAccessServiceTokenRule],
// [AccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseIsDefaultAccessExternalEvaluationRule],
// [AccessGroupGetResponseIsDefaultAccessCountryRule],
// [AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseIsDefaultAccessDevicePostureRule].
type AccessGroupGetResponseIsDefault interface {
	implementsAccessGroupGetResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseIsDefaultAccessEmailRule struct {
	Email AccessGroupGetResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIsDefaultAccessEmailRule]
type accessGroupGetResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessEmailRule) implementsAccessGroupGetResponseIsDefault() {}

type AccessGroupGetResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                  `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessEmailRuleEmail]
type accessGroupGetResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessEmailListRule]
type accessGroupGetResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessEmailListRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                          `json:"id,required"`
	JSON accessGroupGetResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList]
type accessGroupGetResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIsDefaultAccessDomainRule]
type accessGroupGetResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessDomainRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                         `json:"domain,required"`
	JSON   accessGroupGetResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain]
type accessGroupGetResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                           `json:"everyone,required"`
	JSON     accessGroupGetResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIsDefaultAccessEveryoneRule]
type accessGroupGetResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessEveryoneRule) implementsAccessGroupGetResponseIsDefault() {
}

// Matches an IP address block.
type AccessGroupGetResponseIsDefaultAccessIPRule struct {
	IP   AccessGroupGetResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseIsDefaultAccessIPRule]
type accessGroupGetResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessIPRule) implementsAccessGroupGetResponseIsDefault() {}

type AccessGroupGetResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                            `json:"ip,required"`
	JSON accessGroupGetResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseIsDefaultAccessIPRuleIP]
type accessGroupGetResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseIsDefaultAccessIPListRule struct {
	IPList AccessGroupGetResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIsDefaultAccessIPListRule]
type accessGroupGetResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessIPListRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                    `json:"id,required"`
	JSON accessGroupGetResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessIPListRuleIPList]
type accessGroupGetResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                              `json:"certificate,required"`
	JSON        accessGroupGetResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessCertificateRule]
type accessGroupGetResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessCertificateRule) implementsAccessGroupGetResponseIsDefault() {
}

// Matches an Access group.
type AccessGroupGetResponseIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessAccessGroupRule]
type accessGroupGetResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessAccessGroupRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupGetResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup]
type accessGroupGetResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessAzureGroupRule]
type accessGroupGetResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessAzureGroupRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                         `json:"connection_id,required"`
	JSON         accessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule]
type accessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                            `json:"name,required"`
	JSON accessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessGsuiteGroupRule]
type accessGroupGetResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessGsuiteGroupRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                         `json:"email,required"`
	JSON  accessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessOktaGroupRule]
type accessGroupGetResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessOktaGroupRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                     `json:"email,required"`
	JSON  accessGroupGetResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta]
type accessGroupGetResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessSamlGroupRule]
type accessGroupGetResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessSamlGroupRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                     `json:"attribute_value,required"`
	JSON           accessGroupGetResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml]
type accessGroupGetResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessServiceTokenRule]
type accessGroupGetResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessServiceTokenRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                `json:"token_id,required"`
	JSON    accessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                       `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessExternalEvaluationRule]
type accessGroupGetResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessExternalEvaluationRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                            `json:"keys_url,required"`
	JSON    accessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseIsDefaultAccessCountryRule struct {
	Geo  AccessGroupGetResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseIsDefaultAccessCountryRule]
type accessGroupGetResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessCountryRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                  `json:"country_code,required"`
	JSON        accessGroupGetResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessCountryRuleGeo]
type accessGroupGetResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule]
type accessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                      `json:"auth_method,required"`
	JSON       accessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseIsDefaultAccessDevicePostureRule]
type accessGroupGetResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseIsDefaultAccessDevicePostureRule) implementsAccessGroupGetResponseIsDefault() {
}

type AccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                  `json:"integration_uid,required"`
	JSON           accessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseRequireAccessEmailRule],
// [AccessGroupGetResponseRequireAccessEmailListRule],
// [AccessGroupGetResponseRequireAccessDomainRule],
// [AccessGroupGetResponseRequireAccessEveryoneRule],
// [AccessGroupGetResponseRequireAccessIPRule],
// [AccessGroupGetResponseRequireAccessIPListRule],
// [AccessGroupGetResponseRequireAccessCertificateRule],
// [AccessGroupGetResponseRequireAccessAccessGroupRule],
// [AccessGroupGetResponseRequireAccessAzureGroupRule],
// [AccessGroupGetResponseRequireAccessGitHubOrganizationRule],
// [AccessGroupGetResponseRequireAccessGsuiteGroupRule],
// [AccessGroupGetResponseRequireAccessOktaGroupRule],
// [AccessGroupGetResponseRequireAccessSamlGroupRule],
// [AccessGroupGetResponseRequireAccessServiceTokenRule],
// [AccessGroupGetResponseRequireAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseRequireAccessExternalEvaluationRule],
// [AccessGroupGetResponseRequireAccessCountryRule],
// [AccessGroupGetResponseRequireAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseRequireAccessDevicePostureRule].
type AccessGroupGetResponseRequire interface {
	implementsAccessGroupGetResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseRequireAccessEmailRule struct {
	Email AccessGroupGetResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseRequireAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseRequireAccessEmailRule]
type accessGroupGetResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessEmailRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessEmailRuleEmailJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessEmailRuleEmail]
type accessGroupGetResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseRequireAccessEmailListRule struct {
	EmailList AccessGroupGetResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseRequireAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessEmailListRule]
type accessGroupGetResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessEmailListRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupGetResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseRequireAccessEmailListRuleEmailList]
type accessGroupGetResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseRequireAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseRequireAccessDomainRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseRequireAccessDomainRule]
type accessGroupGetResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessDomainRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   accessGroupGetResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseRequireAccessDomainRuleEmailDomain]
type accessGroupGetResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     accessGroupGetResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessEveryoneRule]
type accessGroupGetResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessEveryoneRule) implementsAccessGroupGetResponseRequire() {}

// Matches an IP address block.
type AccessGroupGetResponseRequireAccessIPRule struct {
	IP   AccessGroupGetResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessIPRuleJSON contains the JSON metadata for the
// struct [AccessGroupGetResponseRequireAccessIPRule]
type accessGroupGetResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessIPRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON accessGroupGetResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseRequireAccessIPRuleIP]
type accessGroupGetResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseRequireAccessIPListRule struct {
	IPList AccessGroupGetResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseRequireAccessIPListRuleJSON contains the JSON metadata for
// the struct [AccessGroupGetResponseRequireAccessIPListRule]
type accessGroupGetResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessIPListRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON accessGroupGetResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessIPListRuleIPList]
type accessGroupGetResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        accessGroupGetResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessCertificateRule]
type accessGroupGetResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessCertificateRule) implementsAccessGroupGetResponseRequire() {
}

// Matches an Access group.
type AccessGroupGetResponseRequireAccessAccessGroupRule struct {
	Group AccessGroupGetResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessAccessGroupRule]
type accessGroupGetResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessAccessGroupRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupGetResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseRequireAccessAccessGroupRuleGroup]
type accessGroupGetResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseRequireAccessAzureGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessAzureGroupRule]
type accessGroupGetResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessAzureGroupRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         accessGroupGetResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseRequireAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseRequireAccessGitHubOrganizationRule]
type accessGroupGetResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessGitHubOrganizationRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON accessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessGsuiteGroupRule]
type accessGroupGetResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessGsuiteGroupRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseRequireAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessOktaGroupRule]
type accessGroupGetResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessOktaGroupRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  accessGroupGetResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessOktaGroupRuleOkta]
type accessGroupGetResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseRequireAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessSamlGroupRule]
type accessGroupGetResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessSamlGroupRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           accessGroupGetResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessSamlGroupRuleSaml]
type accessGroupGetResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessServiceTokenRule]
type accessGroupGetResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessServiceTokenRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                              `json:"token_id,required"`
	JSON    accessGroupGetResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessServiceTokenRuleServiceTokenJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                     `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseRequireAccessAnyValidServiceTokenRule]
type accessGroupGetResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseRequireAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseRequireAccessExternalEvaluationRule]
type accessGroupGetResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessExternalEvaluationRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                          `json:"keys_url,required"`
	JSON    accessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseRequireAccessCountryRule struct {
	Geo  AccessGroupGetResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessCountryRule]
type accessGroupGetResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessCountryRule) implementsAccessGroupGetResponseRequire() {}

type AccessGroupGetResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                `json:"country_code,required"`
	JSON        accessGroupGetResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessCountryRuleGeoJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseRequireAccessCountryRuleGeo]
type accessGroupGetResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseRequireAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseRequireAccessAuthenticationMethodRule]
type accessGroupGetResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessAuthenticationMethodRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                    `json:"auth_method,required"`
	JSON       accessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseRequireAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseRequireAccessDevicePostureRule]
type accessGroupGetResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseRequireAccessDevicePostureRule) implementsAccessGroupGetResponseRequire() {
}

type AccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                `json:"integration_uid,required"`
	JSON           accessGroupGetResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseRequireAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupUpdateResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupUpdateResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupUpdateResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupUpdateResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupUpdateResponseRequire `json:"require"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      accessGroupUpdateResponseJSON      `json:"-"`
}

// accessGroupUpdateResponseJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponse]
type accessGroupUpdateResponseJSON struct {
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

func (r *AccessGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseExcludeAccessEmailRule],
// [AccessGroupUpdateResponseExcludeAccessEmailListRule],
// [AccessGroupUpdateResponseExcludeAccessDomainRule],
// [AccessGroupUpdateResponseExcludeAccessEveryoneRule],
// [AccessGroupUpdateResponseExcludeAccessIPRule],
// [AccessGroupUpdateResponseExcludeAccessIPListRule],
// [AccessGroupUpdateResponseExcludeAccessCertificateRule],
// [AccessGroupUpdateResponseExcludeAccessAccessGroupRule],
// [AccessGroupUpdateResponseExcludeAccessAzureGroupRule],
// [AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseExcludeAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseExcludeAccessOktaGroupRule],
// [AccessGroupUpdateResponseExcludeAccessSamlGroupRule],
// [AccessGroupUpdateResponseExcludeAccessServiceTokenRule],
// [AccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseExcludeAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseExcludeAccessCountryRule],
// [AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseExcludeAccessDevicePostureRule].
type AccessGroupUpdateResponseExclude interface {
	implementsAccessGroupUpdateResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseExcludeAccessEmailRule struct {
	Email AccessGroupUpdateResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseExcludeAccessEmailRule]
type accessGroupUpdateResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessEmailRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                   `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessEmailRuleEmail]
type accessGroupUpdateResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseExcludeAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessEmailListRule]
type accessGroupUpdateResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessEmailListRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupUpdateResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList]
type accessGroupUpdateResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseExcludeAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseExcludeAccessDomainRule]
type accessGroupUpdateResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessDomainRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                          `json:"domain,required"`
	JSON   accessGroupUpdateResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                            `json:"everyone,required"`
	JSON     accessGroupUpdateResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessEveryoneRule]
type accessGroupUpdateResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessEveryoneRule) implementsAccessGroupUpdateResponseExclude() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseExcludeAccessIPRule struct {
	IP   AccessGroupUpdateResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupUpdateResponseExcludeAccessIPRule]
type accessGroupUpdateResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessIPRule) implementsAccessGroupUpdateResponseExclude() {}

type AccessGroupUpdateResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                             `json:"ip,required"`
	JSON accessGroupUpdateResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseExcludeAccessIPRuleIP]
type accessGroupUpdateResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseExcludeAccessIPListRule struct {
	IPList AccessGroupUpdateResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseExcludeAccessIPListRule]
type accessGroupUpdateResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessIPListRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                     `json:"id,required"`
	JSON accessGroupUpdateResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessIPListRuleIPList]
type accessGroupUpdateResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                               `json:"certificate,required"`
	JSON        accessGroupUpdateResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessCertificateRule]
type accessGroupUpdateResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessCertificateRule) implementsAccessGroupUpdateResponseExclude() {
}

// Matches an Access group.
type AccessGroupUpdateResponseExcludeAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessAccessGroupRule]
type accessGroupUpdateResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessAccessGroupRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                         `json:"id,required"`
	JSON accessGroupUpdateResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessAzureGroupRule]
type accessGroupUpdateResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessAzureGroupRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule]
type accessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                             `json:"name,required"`
	JSON accessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessGsuiteGroupRule]
type accessGroupUpdateResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                          `json:"email,required"`
	JSON  accessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseExcludeAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessOktaGroupRule]
type accessGroupUpdateResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessOktaGroupRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                      `json:"email,required"`
	JSON  accessGroupUpdateResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseExcludeAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessSamlGroupRule]
type accessGroupUpdateResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessSamlGroupRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                      `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessServiceTokenRule]
type accessGroupUpdateResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessServiceTokenRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                 `json:"token_id,required"`
	JSON    accessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                        `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessExternalEvaluationRule]
type accessGroupUpdateResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                             `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseExcludeAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseExcludeAccessCountryRule]
type accessGroupUpdateResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessCountryRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                   `json:"country_code,required"`
	JSON        accessGroupUpdateResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseExcludeAccessCountryRuleGeo]
type accessGroupUpdateResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule]
type accessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                       `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessDevicePostureRule]
type accessGroupUpdateResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseExcludeAccessDevicePostureRule) implementsAccessGroupUpdateResponseExclude() {
}

type AccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                   `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseIncludeAccessEmailRule],
// [AccessGroupUpdateResponseIncludeAccessEmailListRule],
// [AccessGroupUpdateResponseIncludeAccessDomainRule],
// [AccessGroupUpdateResponseIncludeAccessEveryoneRule],
// [AccessGroupUpdateResponseIncludeAccessIPRule],
// [AccessGroupUpdateResponseIncludeAccessIPListRule],
// [AccessGroupUpdateResponseIncludeAccessCertificateRule],
// [AccessGroupUpdateResponseIncludeAccessAccessGroupRule],
// [AccessGroupUpdateResponseIncludeAccessAzureGroupRule],
// [AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseIncludeAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseIncludeAccessOktaGroupRule],
// [AccessGroupUpdateResponseIncludeAccessSamlGroupRule],
// [AccessGroupUpdateResponseIncludeAccessServiceTokenRule],
// [AccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseIncludeAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseIncludeAccessCountryRule],
// [AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseIncludeAccessDevicePostureRule].
type AccessGroupUpdateResponseInclude interface {
	implementsAccessGroupUpdateResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseIncludeAccessEmailRule struct {
	Email AccessGroupUpdateResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIncludeAccessEmailRule]
type accessGroupUpdateResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessEmailRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                   `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessEmailRuleEmail]
type accessGroupUpdateResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseIncludeAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessEmailListRule]
type accessGroupUpdateResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessEmailListRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupUpdateResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList]
type accessGroupUpdateResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseIncludeAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIncludeAccessDomainRule]
type accessGroupUpdateResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessDomainRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                          `json:"domain,required"`
	JSON   accessGroupUpdateResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                            `json:"everyone,required"`
	JSON     accessGroupUpdateResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessEveryoneRule]
type accessGroupUpdateResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessEveryoneRule) implementsAccessGroupUpdateResponseInclude() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseIncludeAccessIPRule struct {
	IP   AccessGroupUpdateResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupUpdateResponseIncludeAccessIPRule]
type accessGroupUpdateResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessIPRule) implementsAccessGroupUpdateResponseInclude() {}

type AccessGroupUpdateResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                             `json:"ip,required"`
	JSON accessGroupUpdateResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIncludeAccessIPRuleIP]
type accessGroupUpdateResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseIncludeAccessIPListRule struct {
	IPList AccessGroupUpdateResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIncludeAccessIPListRule]
type accessGroupUpdateResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessIPListRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                     `json:"id,required"`
	JSON accessGroupUpdateResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessIPListRuleIPList]
type accessGroupUpdateResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                               `json:"certificate,required"`
	JSON        accessGroupUpdateResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessCertificateRule]
type accessGroupUpdateResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessCertificateRule) implementsAccessGroupUpdateResponseInclude() {
}

// Matches an Access group.
type AccessGroupUpdateResponseIncludeAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessAccessGroupRule]
type accessGroupUpdateResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessAccessGroupRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                         `json:"id,required"`
	JSON accessGroupUpdateResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessAzureGroupRule]
type accessGroupUpdateResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessAzureGroupRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule]
type accessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                             `json:"name,required"`
	JSON accessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessGsuiteGroupRule]
type accessGroupUpdateResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                          `json:"email,required"`
	JSON  accessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseIncludeAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessOktaGroupRule]
type accessGroupUpdateResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessOktaGroupRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                      `json:"email,required"`
	JSON  accessGroupUpdateResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseIncludeAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessSamlGroupRule]
type accessGroupUpdateResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessSamlGroupRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                      `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessServiceTokenRule]
type accessGroupUpdateResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessServiceTokenRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                 `json:"token_id,required"`
	JSON    accessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                        `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessExternalEvaluationRule]
type accessGroupUpdateResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                             `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseIncludeAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIncludeAccessCountryRule]
type accessGroupUpdateResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessCountryRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                   `json:"country_code,required"`
	JSON        accessGroupUpdateResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIncludeAccessCountryRuleGeo]
type accessGroupUpdateResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule]
type accessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                       `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessDevicePostureRule]
type accessGroupUpdateResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIncludeAccessDevicePostureRule) implementsAccessGroupUpdateResponseInclude() {
}

type AccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                   `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseIsDefaultAccessEmailRule],
// [AccessGroupUpdateResponseIsDefaultAccessEmailListRule],
// [AccessGroupUpdateResponseIsDefaultAccessDomainRule],
// [AccessGroupUpdateResponseIsDefaultAccessEveryoneRule],
// [AccessGroupUpdateResponseIsDefaultAccessIPRule],
// [AccessGroupUpdateResponseIsDefaultAccessIPListRule],
// [AccessGroupUpdateResponseIsDefaultAccessCertificateRule],
// [AccessGroupUpdateResponseIsDefaultAccessAccessGroupRule],
// [AccessGroupUpdateResponseIsDefaultAccessAzureGroupRule],
// [AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseIsDefaultAccessOktaGroupRule],
// [AccessGroupUpdateResponseIsDefaultAccessSamlGroupRule],
// [AccessGroupUpdateResponseIsDefaultAccessServiceTokenRule],
// [AccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseIsDefaultAccessCountryRule],
// [AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseIsDefaultAccessDevicePostureRule].
type AccessGroupUpdateResponseIsDefault interface {
	implementsAccessGroupUpdateResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseIsDefaultAccessEmailRule struct {
	Email AccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIsDefaultAccessEmailRule]
type accessGroupUpdateResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessEmailRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                     `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail]
type accessGroupUpdateResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessEmailListRule]
type accessGroupUpdateResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessEmailListRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                             `json:"id,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList]
type accessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessDomainRule]
type accessGroupUpdateResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessDomainRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                            `json:"domain,required"`
	JSON   accessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                              `json:"everyone,required"`
	JSON     accessGroupUpdateResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessEveryoneRule]
type accessGroupUpdateResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessEveryoneRule) implementsAccessGroupUpdateResponseIsDefault() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseIsDefaultAccessIPRule struct {
	IP   AccessGroupUpdateResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIsDefaultAccessIPRule]
type accessGroupUpdateResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessIPRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                               `json:"ip,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseIsDefaultAccessIPRuleIP]
type accessGroupUpdateResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseIsDefaultAccessIPListRule struct {
	IPList AccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessIPListRule]
type accessGroupUpdateResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessIPListRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                       `json:"id,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList]
type accessGroupUpdateResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                 `json:"certificate,required"`
	JSON        accessGroupUpdateResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessCertificateRule]
type accessGroupUpdateResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessCertificateRule) implementsAccessGroupUpdateResponseIsDefault() {
}

// Matches an Access group.
type AccessGroupUpdateResponseIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessAccessGroupRule]
type accessGroupUpdateResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessAccessGroupRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessAzureGroupRule]
type accessGroupUpdateResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessAzureGroupRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                            `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule]
type accessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                               `json:"name,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule]
type accessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                            `json:"email,required"`
	JSON  accessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessOktaGroupRule]
type accessGroupUpdateResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessOktaGroupRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                        `json:"email,required"`
	JSON  accessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessSamlGroupRule]
type accessGroupUpdateResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessSamlGroupRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                        `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessServiceTokenRule]
type accessGroupUpdateResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessServiceTokenRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                   `json:"token_id,required"`
	JSON    accessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                          `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule]
type accessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                               `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseIsDefaultAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessCountryRule]
type accessGroupUpdateResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessCountryRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                     `json:"country_code,required"`
	JSON        accessGroupUpdateResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo]
type accessGroupUpdateResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule]
type accessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                         `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessDevicePostureRule]
type accessGroupUpdateResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseIsDefaultAccessDevicePostureRule) implementsAccessGroupUpdateResponseIsDefault() {
}

type AccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                     `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseRequireAccessEmailRule],
// [AccessGroupUpdateResponseRequireAccessEmailListRule],
// [AccessGroupUpdateResponseRequireAccessDomainRule],
// [AccessGroupUpdateResponseRequireAccessEveryoneRule],
// [AccessGroupUpdateResponseRequireAccessIPRule],
// [AccessGroupUpdateResponseRequireAccessIPListRule],
// [AccessGroupUpdateResponseRequireAccessCertificateRule],
// [AccessGroupUpdateResponseRequireAccessAccessGroupRule],
// [AccessGroupUpdateResponseRequireAccessAzureGroupRule],
// [AccessGroupUpdateResponseRequireAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseRequireAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseRequireAccessOktaGroupRule],
// [AccessGroupUpdateResponseRequireAccessSamlGroupRule],
// [AccessGroupUpdateResponseRequireAccessServiceTokenRule],
// [AccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseRequireAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseRequireAccessCountryRule],
// [AccessGroupUpdateResponseRequireAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseRequireAccessDevicePostureRule].
type AccessGroupUpdateResponseRequire interface {
	implementsAccessGroupUpdateResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseRequireAccessEmailRule struct {
	Email AccessGroupUpdateResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseRequireAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseRequireAccessEmailRule]
type accessGroupUpdateResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessEmailRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                   `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessEmailRuleEmail]
type accessGroupUpdateResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseRequireAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessEmailListRule]
type accessGroupUpdateResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessEmailListRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupUpdateResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessEmailListRuleEmailList]
type accessGroupUpdateResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseRequireAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseRequireAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseRequireAccessDomainRule]
type accessGroupUpdateResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessDomainRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                          `json:"domain,required"`
	JSON   accessGroupUpdateResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                            `json:"everyone,required"`
	JSON     accessGroupUpdateResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessEveryoneRule]
type accessGroupUpdateResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessEveryoneRule) implementsAccessGroupUpdateResponseRequire() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseRequireAccessIPRule struct {
	IP   AccessGroupUpdateResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupUpdateResponseRequireAccessIPRule]
type accessGroupUpdateResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessIPRule) implementsAccessGroupUpdateResponseRequire() {}

type AccessGroupUpdateResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                             `json:"ip,required"`
	JSON accessGroupUpdateResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseRequireAccessIPRuleIP]
type accessGroupUpdateResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseRequireAccessIPListRule struct {
	IPList AccessGroupUpdateResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseRequireAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseRequireAccessIPListRule]
type accessGroupUpdateResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessIPListRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                     `json:"id,required"`
	JSON accessGroupUpdateResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessIPListRuleIPList]
type accessGroupUpdateResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                               `json:"certificate,required"`
	JSON        accessGroupUpdateResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessCertificateRule]
type accessGroupUpdateResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessCertificateRule) implementsAccessGroupUpdateResponseRequire() {
}

// Matches an Access group.
type AccessGroupUpdateResponseRequireAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessAccessGroupRule]
type accessGroupUpdateResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessAccessGroupRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                         `json:"id,required"`
	JSON accessGroupUpdateResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessAzureGroupRule]
type accessGroupUpdateResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessAzureGroupRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseRequireAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessGitHubOrganizationRule]
type accessGroupUpdateResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                             `json:"name,required"`
	JSON accessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessGsuiteGroupRule]
type accessGroupUpdateResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                          `json:"email,required"`
	JSON  accessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseRequireAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessOktaGroupRule]
type accessGroupUpdateResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessOktaGroupRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                      `json:"email,required"`
	JSON  accessGroupUpdateResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseRequireAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessSamlGroupRule]
type accessGroupUpdateResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessSamlGroupRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                      `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessServiceTokenRule]
type accessGroupUpdateResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessServiceTokenRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                 `json:"token_id,required"`
	JSON    accessGroupUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                        `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseRequireAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessExternalEvaluationRule]
type accessGroupUpdateResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                             `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseRequireAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupUpdateResponseRequireAccessCountryRule]
type accessGroupUpdateResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessCountryRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                   `json:"country_code,required"`
	JSON        accessGroupUpdateResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseRequireAccessCountryRuleGeo]
type accessGroupUpdateResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseRequireAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessAuthenticationMethodRule]
type accessGroupUpdateResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                       `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseRequireAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseRequireAccessDevicePostureRule]
type accessGroupUpdateResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseRequireAccessDevicePostureRule) implementsAccessGroupUpdateResponseRequire() {
}

type AccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                   `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type AccessGroupAccessGroupsNewAnAccessGroupResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupAccessGroupsNewAnAccessGroupResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupAccessGroupsNewAnAccessGroupResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupAccessGroupsNewAnAccessGroupResponseRequire `json:"require"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseJSON contains the JSON metadata
// for the struct [AccessGroupAccessGroupsNewAnAccessGroupResponse]
type accessGroupAccessGroupsNewAnAccessGroupResponseJSON struct {
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

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseExclude interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                         `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                  `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                           `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                     `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                               `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                   `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                            `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                            `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                       `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                   `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                         `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                             `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                         `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseInclude interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                         `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                  `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                           `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                     `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                               `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                   `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                            `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                            `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                       `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                   `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                         `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                             `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                         `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                           `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                   `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                  `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                    `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                     `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                             `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                                       `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                  `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                     `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                              `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                              `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                         `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                     `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                           `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                               `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                           `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequire interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                         `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                  `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                           `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                     `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                               `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                   `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                            `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                            `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                       `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                   `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                         `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                             `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                         `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupAccessGroupsListAccessGroupsResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupAccessGroupsListAccessGroupsResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupAccessGroupsListAccessGroupsResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupAccessGroupsListAccessGroupsResponseRequire `json:"require"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseJSON contains the JSON metadata
// for the struct [AccessGroupAccessGroupsListAccessGroupsResponse]
type accessGroupAccessGroupsListAccessGroupsResponseJSON struct {
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

func (r *AccessGroupAccessGroupsListAccessGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseExclude interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                         `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                  `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                           `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                     `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                               `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                   `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                            `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                            `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                       `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                   `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                         `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                             `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                         `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseInclude interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                         `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                  `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                           `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                     `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                               `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                   `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                            `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                            `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                       `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                   `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                         `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                             `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                         `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefault interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                           `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                   `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                  `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                    `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                     `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                             `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                                       `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                  `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                     `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                              `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                              `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                         `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                     `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                           `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                               `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                           `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseRequire interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                         `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                  `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                           `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                                     `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                               `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                   `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                            `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                            `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                       `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                   `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                         `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                             `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                         `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupGetResponseEnvelope struct {
	Errors   []AccessGroupGetResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessGroupGetResponseEnvelopeMessages `json:"messages"`
	Result   AccessGroupGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessGroupGetResponseEnvelopeSuccess `json:"success"`
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

type AccessGroupGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessGroupGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupGetResponseEnvelopeErrors]
type accessGroupGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessGroupGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupGetResponseEnvelopeMessages]
type accessGroupGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupGetResponseEnvelopeSuccess bool

const (
	AccessGroupGetResponseEnvelopeSuccessTrue AccessGroupGetResponseEnvelopeSuccess = true
)

type AccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessGroupUpdateParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessGroupUpdateParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessGroupUpdateParamsRequire] `json:"require"`
}

func (r AccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupUpdateParamsIncludeAccessEmailRule],
// [AccessGroupUpdateParamsIncludeAccessEmailListRule],
// [AccessGroupUpdateParamsIncludeAccessDomainRule],
// [AccessGroupUpdateParamsIncludeAccessEveryoneRule],
// [AccessGroupUpdateParamsIncludeAccessIPRule],
// [AccessGroupUpdateParamsIncludeAccessIPListRule],
// [AccessGroupUpdateParamsIncludeAccessCertificateRule],
// [AccessGroupUpdateParamsIncludeAccessAccessGroupRule],
// [AccessGroupUpdateParamsIncludeAccessAzureGroupRule],
// [AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule],
// [AccessGroupUpdateParamsIncludeAccessGsuiteGroupRule],
// [AccessGroupUpdateParamsIncludeAccessOktaGroupRule],
// [AccessGroupUpdateParamsIncludeAccessSamlGroupRule],
// [AccessGroupUpdateParamsIncludeAccessServiceTokenRule],
// [AccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule],
// [AccessGroupUpdateParamsIncludeAccessCountryRule],
// [AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule],
// [AccessGroupUpdateParamsIncludeAccessDevicePostureRule].
type AccessGroupUpdateParamsInclude interface {
	implementsAccessGroupUpdateParamsInclude()
}

// Matches a specific email.
type AccessGroupUpdateParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessGroupUpdateParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessEmailRule) implementsAccessGroupUpdateParamsInclude() {}

type AccessGroupUpdateParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupUpdateParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupUpdateParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupUpdateParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessEmailListRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupUpdateParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupUpdateParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessDomainRule) implementsAccessGroupUpdateParamsInclude() {}

type AccessGroupUpdateParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupUpdateParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessEveryoneRule) implementsAccessGroupUpdateParamsInclude() {
}

// Matches an IP address block.
type AccessGroupUpdateParamsIncludeAccessIPRule struct {
	IP param.Field[AccessGroupUpdateParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessIPRule) implementsAccessGroupUpdateParamsInclude() {}

type AccessGroupUpdateParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupUpdateParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessGroupUpdateParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessIPListRule) implementsAccessGroupUpdateParamsInclude() {}

type AccessGroupUpdateParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupUpdateParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessCertificateRule) implementsAccessGroupUpdateParamsInclude() {
}

// Matches an Access group.
type AccessGroupUpdateParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupUpdateParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessAccessGroupRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupUpdateParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessAzureGroupRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupUpdateParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessGsuiteGroupRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupUpdateParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessOktaGroupRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupUpdateParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessSamlGroupRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupUpdateParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessServiceTokenRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupUpdateParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessGroupUpdateParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessCountryRule) implementsAccessGroupUpdateParamsInclude() {}

type AccessGroupUpdateParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupUpdateParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessDevicePostureRule) implementsAccessGroupUpdateParamsInclude() {
}

type AccessGroupUpdateParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupUpdateParamsExcludeAccessEmailRule],
// [AccessGroupUpdateParamsExcludeAccessEmailListRule],
// [AccessGroupUpdateParamsExcludeAccessDomainRule],
// [AccessGroupUpdateParamsExcludeAccessEveryoneRule],
// [AccessGroupUpdateParamsExcludeAccessIPRule],
// [AccessGroupUpdateParamsExcludeAccessIPListRule],
// [AccessGroupUpdateParamsExcludeAccessCertificateRule],
// [AccessGroupUpdateParamsExcludeAccessAccessGroupRule],
// [AccessGroupUpdateParamsExcludeAccessAzureGroupRule],
// [AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule],
// [AccessGroupUpdateParamsExcludeAccessGsuiteGroupRule],
// [AccessGroupUpdateParamsExcludeAccessOktaGroupRule],
// [AccessGroupUpdateParamsExcludeAccessSamlGroupRule],
// [AccessGroupUpdateParamsExcludeAccessServiceTokenRule],
// [AccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule],
// [AccessGroupUpdateParamsExcludeAccessCountryRule],
// [AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule],
// [AccessGroupUpdateParamsExcludeAccessDevicePostureRule].
type AccessGroupUpdateParamsExclude interface {
	implementsAccessGroupUpdateParamsExclude()
}

// Matches a specific email.
type AccessGroupUpdateParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessGroupUpdateParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessEmailRule) implementsAccessGroupUpdateParamsExclude() {}

type AccessGroupUpdateParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupUpdateParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupUpdateParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupUpdateParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessEmailListRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupUpdateParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupUpdateParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessDomainRule) implementsAccessGroupUpdateParamsExclude() {}

type AccessGroupUpdateParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupUpdateParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessEveryoneRule) implementsAccessGroupUpdateParamsExclude() {
}

// Matches an IP address block.
type AccessGroupUpdateParamsExcludeAccessIPRule struct {
	IP param.Field[AccessGroupUpdateParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessIPRule) implementsAccessGroupUpdateParamsExclude() {}

type AccessGroupUpdateParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupUpdateParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessGroupUpdateParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessIPListRule) implementsAccessGroupUpdateParamsExclude() {}

type AccessGroupUpdateParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupUpdateParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessCertificateRule) implementsAccessGroupUpdateParamsExclude() {
}

// Matches an Access group.
type AccessGroupUpdateParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupUpdateParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessAccessGroupRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupUpdateParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessAzureGroupRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupUpdateParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessGsuiteGroupRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupUpdateParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessOktaGroupRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupUpdateParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessSamlGroupRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupUpdateParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessServiceTokenRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupUpdateParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessGroupUpdateParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessCountryRule) implementsAccessGroupUpdateParamsExclude() {}

type AccessGroupUpdateParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupUpdateParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessDevicePostureRule) implementsAccessGroupUpdateParamsExclude() {
}

type AccessGroupUpdateParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupUpdateParamsRequireAccessEmailRule],
// [AccessGroupUpdateParamsRequireAccessEmailListRule],
// [AccessGroupUpdateParamsRequireAccessDomainRule],
// [AccessGroupUpdateParamsRequireAccessEveryoneRule],
// [AccessGroupUpdateParamsRequireAccessIPRule],
// [AccessGroupUpdateParamsRequireAccessIPListRule],
// [AccessGroupUpdateParamsRequireAccessCertificateRule],
// [AccessGroupUpdateParamsRequireAccessAccessGroupRule],
// [AccessGroupUpdateParamsRequireAccessAzureGroupRule],
// [AccessGroupUpdateParamsRequireAccessGitHubOrganizationRule],
// [AccessGroupUpdateParamsRequireAccessGsuiteGroupRule],
// [AccessGroupUpdateParamsRequireAccessOktaGroupRule],
// [AccessGroupUpdateParamsRequireAccessSamlGroupRule],
// [AccessGroupUpdateParamsRequireAccessServiceTokenRule],
// [AccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateParamsRequireAccessExternalEvaluationRule],
// [AccessGroupUpdateParamsRequireAccessCountryRule],
// [AccessGroupUpdateParamsRequireAccessAuthenticationMethodRule],
// [AccessGroupUpdateParamsRequireAccessDevicePostureRule].
type AccessGroupUpdateParamsRequire interface {
	implementsAccessGroupUpdateParamsRequire()
}

// Matches a specific email.
type AccessGroupUpdateParamsRequireAccessEmailRule struct {
	Email param.Field[AccessGroupUpdateParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupUpdateParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessEmailRule) implementsAccessGroupUpdateParamsRequire() {}

type AccessGroupUpdateParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupUpdateParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupUpdateParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessGroupUpdateParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupUpdateParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessEmailListRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupUpdateParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupUpdateParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupUpdateParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessDomainRule) implementsAccessGroupUpdateParamsRequire() {}

type AccessGroupUpdateParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupUpdateParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupUpdateParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupUpdateParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessEveryoneRule) implementsAccessGroupUpdateParamsRequire() {
}

// Matches an IP address block.
type AccessGroupUpdateParamsRequireAccessIPRule struct {
	IP param.Field[AccessGroupUpdateParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessIPRule) implementsAccessGroupUpdateParamsRequire() {}

type AccessGroupUpdateParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupUpdateParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessGroupUpdateParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupUpdateParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessIPListRule) implementsAccessGroupUpdateParamsRequire() {}

type AccessGroupUpdateParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupUpdateParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupUpdateParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessCertificateRule) implementsAccessGroupUpdateParamsRequire() {
}

// Matches an Access group.
type AccessGroupUpdateParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessGroupUpdateParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessAccessGroupRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupUpdateParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessAzureGroupRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupUpdateParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessGitHubOrganizationRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupUpdateParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupUpdateParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessGsuiteGroupRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupUpdateParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupUpdateParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupUpdateParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessOktaGroupRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupUpdateParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupUpdateParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupUpdateParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessSamlGroupRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupUpdateParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupUpdateParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupUpdateParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessServiceTokenRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupUpdateParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupUpdateParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessExternalEvaluationRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupUpdateParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessGroupUpdateParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupUpdateParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessCountryRule) implementsAccessGroupUpdateParamsRequire() {}

type AccessGroupUpdateParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupUpdateParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupUpdateParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessAuthenticationMethodRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupUpdateParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupUpdateParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessDevicePostureRule) implementsAccessGroupUpdateParamsRequire() {
}

type AccessGroupUpdateParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupUpdateParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupUpdateResponseEnvelope struct {
	Errors   []AccessGroupUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessGroupUpdateResponseEnvelopeMessages `json:"messages"`
	Result   AccessGroupUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessGroupUpdateResponseEnvelopeSuccess `json:"success"`
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

type AccessGroupUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessGroupUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupUpdateResponseEnvelopeErrors]
type accessGroupUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessGroupUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupUpdateResponseEnvelopeMessages]
type accessGroupUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupUpdateResponseEnvelopeSuccess bool

const (
	AccessGroupUpdateResponseEnvelopeSuccessTrue AccessGroupUpdateResponseEnvelopeSuccess = true
)

type AccessGroupDeleteResponseEnvelope struct {
	Errors   []AccessGroupDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessGroupDeleteResponseEnvelopeMessages `json:"messages"`
	Result   AccessGroupDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessGroupDeleteResponseEnvelopeSuccess `json:"success"`
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

type AccessGroupDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessGroupDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupDeleteResponseEnvelopeErrors]
type accessGroupDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessGroupDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupDeleteResponseEnvelopeMessages]
type accessGroupDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupDeleteResponseEnvelopeSuccess bool

const (
	AccessGroupDeleteResponseEnvelopeSuccessTrue AccessGroupDeleteResponseEnvelopeSuccess = true
)

type AccessGroupAccessGroupsNewAnAccessGroupParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessGroupAccessGroupsNewAnAccessGroupParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessGroupAccessGroupsNewAnAccessGroupParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessGroupAccessGroupsNewAnAccessGroupParamsRequire] `json:"require"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAuthenticationMethodRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupParamsInclude interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude()
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPRule struct {
	IP param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAuthenticationMethodRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupParamsExclude interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude()
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPRule struct {
	IP param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAuthenticationMethodRule],
// [AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequire interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire()
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailRule struct {
	Email param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPRule struct {
	IP param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelope struct {
	Errors   []AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeMessages `json:"messages"`
	Result   AccessGroupAccessGroupsNewAnAccessGroupResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeSuccess `json:"success"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelope]
type accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeErrors]
type accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeMessages]
type accessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeSuccess bool

const (
	AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeSuccessTrue AccessGroupAccessGroupsNewAnAccessGroupResponseEnvelopeSuccess = true
)

type AccessGroupAccessGroupsListAccessGroupsResponseEnvelope struct {
	Errors     []AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeMessages `json:"messages"`
	Result     []AccessGroupAccessGroupsListAccessGroupsResponse                 `json:"result"`
	ResultInfo AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeSuccess `json:"success"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseEnvelopeJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseEnvelope]
type accessGroupAccessGroupsListAccessGroupsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeErrors]
type accessGroupAccessGroupsListAccessGroupsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeMessages]
type accessGroupAccessGroupsListAccessGroupsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeResultInfo]
type accessGroupAccessGroupsListAccessGroupsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeSuccess bool

const (
	AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeSuccessTrue AccessGroupAccessGroupsListAccessGroupsResponseEnvelopeSuccess = true
)
